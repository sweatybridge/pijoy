// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
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

// Defines values for Button.
const (
	Down  Button = "down"
	Left  Button = "left"
	Right Button = "right"
	Up    Button = "up"
)

// Defines values for HealthStatus.
const (
	Draining HealthStatus = "draining"
	Ready    HealthStatus = "ready"
)

// Button defines model for Button.
type Button string

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Health defines model for Health.
type Health struct {
	Status  HealthStatus `json:"status"`
	Version string       `json:"version"`
}

// HealthStatus defines model for Health.Status.
type HealthStatus string

// InternalError defines model for InternalError.
type InternalError = Error

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6yUzY7jNgzHX0Vge/RY6W4PC5/6gUE3RVts27kt5qDItK2sLakknVlj4HcvJDsfM0nR",
	"BdpTHJH8kfxT1DPYMMTg0QtD9QyEHINnzH+2XpC86e+JAqUDG7ygl/RpYuydNeKC13sOPp2x7XAw6etr",
	"wgYq+Eqf6Xqxsl5o8zwXUCNbcjFBoILR4+eIVrBWuPoUKzOX88MosiRCPw5QfYQxQgF1ePJQQI+NQAHk",
	"2k7gsQCZIkIFLOR8C3MBpy4ihYgkbmnShhrTbxNoMAIVOC9v38AJ4Lxgi5QIAzKbNnu/os8FEP41OsI6",
	"1ZWZZ/9zNWG3RyuJ9R5NL911OSxGRr5sktDUU+qTjPMp263mDkjsVnE+myH2yXz4ptyUGyj+pdo15Rly",
	"XW8Kcb4Jx0tgbL4EOBjXJ/ATGpl25OoWv2vTYWnDAAV4MyTM784E9d6kMY2UAjqRyJXWrZNu3CVnfcmA",
	"q7vx/YetagKp6PYhidE7i57zKNYcv24fvhSvM0Xv+rDTg3Fe/7L98f63P+9TWnGStcsudya6C10qWASd",
	"CwgRfTJW8HbVOBrp8tx0dxpti1mml638gTKSZ8VIB2dRrXRlfK2WULVMpISch/KSbWuo4CeU9d4ULzf1",
	"zWbzv+3nmuHGgq7VHTMvU2rM2Ms/MU9F6pdvyZzpeh8mFmc/6edd3u0570PgG6p9IGTOGhH2aBiVUUuQ",
	"Cl5Jh+oIu5Ytx/68mvOoyAwoSAzVx9eJHqaIKjRHuAQVU3iiumRPcz7f7MULLhdKaMTiC7Ven7R5frwa",
	"6LfXGjx0eCzryfBSF9aKR2uRuRn7fir/41DSi4t0OEpzXqdK6z5Y03eBpXq3eZe24JWdDMcdEk3Rldn3",
	"ll9aS+4CST/dcWfSg3TXBWIsfUvh011DiKWJEebH+e8AAAD//z7lMS6dBgAA",
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
