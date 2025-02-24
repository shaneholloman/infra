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
	"aHhZAAA=",
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
