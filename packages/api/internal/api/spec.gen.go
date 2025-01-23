// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"bD40d+bV1ep/AQAA//8Gb+UCIlIAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
