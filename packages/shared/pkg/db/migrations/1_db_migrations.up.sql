BEGIN;
ALTER TABLE IF EXISTS "public"."schema_migrations" ENABLE ROW LEVEL SECURITY;
COMMIT;
