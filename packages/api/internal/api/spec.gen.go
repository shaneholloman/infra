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
