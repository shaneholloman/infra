package server

import (
	"context"
	"log"

	"cloud.google.com/go/storage"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/filters"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/dns"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox"
	localStorage "github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/local_storage"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox/network"
	"github.com/e2b-dev/infra/packages/shared/pkg/grpc/orchestrator"
	"github.com/e2b-dev/infra/packages/shared/pkg/smap"
	templateStorage "github.com/e2b-dev/infra/packages/shared/pkg/storage"
)

const ServiceName = "orchestrator"

type server struct {
	orchestrator.UnimplementedSandboxServiceServer
	sandboxes     *smap.Map[*sandbox.Sandbox]
	dns           *dns.DNS
	tracer        trace.Tracer
	networkPool   *network.SlotPool
	templateCache *localStorage.TemplateCache
}

func New() *grpc.Server {
	log.Println("Initializing orchestrator")

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler(otelgrpc.WithInterceptorFilter(filters.Not(filters.HealthCheck())))),
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
		),
	)

	ctx := context.Background()

	dnsServer := dns.New()
	go dnsServer.Start("127.0.0.1:53")

	tracer := otel.Tracer(ServiceName)

	client, err := storage.NewClient(ctx, storage.WithJSONReads())
	if err != nil {
		log.Fatalf("failed to create GCS client: %v", err)
	}

	templateCache := localStorage.NewTemplateCache(
		ctx,
		client.Bucket(templateStorage.BucketName),
	)

	networkPool := network.NewSlotPool()

	// We start the pool last to avoid allocation network slots if the other components fail to initialize.
	go func() {
		poolErr := networkPool.Populate(ctx)
		if poolErr != nil {
			log.Fatalf("network pool error: %v\n", poolErr)
		}
	}()

	orchestrator.RegisterSandboxServiceServer(s, &server{
		tracer:        tracer,
		dns:           dnsServer,
		sandboxes:     smap.New[*sandbox.Sandbox](),
		networkPool:   networkPool,
		templateCache: templateCache,
	})

	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	return s
}
