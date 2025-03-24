// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for EntryInfoType.
const (
	File EntryInfoType = "file"
)

// EntryInfo defines model for EntryInfo.
type EntryInfo struct {
	// Name Name of the file
	Name string `json:"name"`

	// Path Path to the file
	Path string `json:"path"`

	// Type Type of the file
	Type EntryInfoType `json:"type"`
}

// EntryInfoType Type of the file
type EntryInfoType string

// EnvVars Environment variables to set
type EnvVars map[string]string

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// Metrics Resource usage metrics
type Metrics struct {
	// CpuUsedPct CPU usage percentage
	CpuUsedPct *float32 `json:"cpu_used_pct,omitempty"`

	// MemBytes Total virtual memory usage in bytes
	MemBytes *int `json:"mem_bytes,omitempty"`
}

// FilePath defines model for FilePath.
type FilePath = string

// User defines model for User.
type User = string

// FileNotFound defines model for FileNotFound.
type FileNotFound = Error

// InternalServerError defines model for InternalServerError.
type InternalServerError = Error

// InvalidPath defines model for InvalidPath.
type InvalidPath = Error

// InvalidUser defines model for InvalidUser.
type InvalidUser = Error

// NotEnoughDiskSpace defines model for NotEnoughDiskSpace.
type NotEnoughDiskSpace = Error

// UploadSuccess defines model for UploadSuccess.
type UploadSuccess = []EntryInfo

// GetFilesParams defines parameters for GetFiles.
type GetFilesParams struct {
	// Path Path to the file, URL encoded. Can be relative to user's home directory.
	Path *FilePath `form:"path,omitempty" json:"path,omitempty"`

	// Username User used for setting the owner, or resolving relative paths.
	Username User `form:"username" json:"username"`
}

// PostFilesMultipartBody defines parameters for PostFiles.
type PostFilesMultipartBody struct {
	File *openapi_types.File `json:"file,omitempty"`
}

// PostFilesParams defines parameters for PostFiles.
type PostFilesParams struct {
	// Path Path to the file, URL encoded. Can be relative to user's home directory.
	Path *FilePath `form:"path,omitempty" json:"path,omitempty"`

	// Username User used for setting the owner, or resolving relative paths.
	Username User `form:"username" json:"username"`
}

// PostInitJSONBody defines parameters for PostInit.
type PostInitJSONBody struct {
	// EnvVars Environment variables to set
	EnvVars *EnvVars `json:"envVars,omitempty"`
}

// PostFilesMultipartRequestBody defines body for PostFiles for multipart/form-data ContentType.
type PostFilesMultipartRequestBody PostFilesMultipartBody

// PostInitJSONRequestBody defines body for PostInit for application/json ContentType.
type PostInitJSONRequestBody PostInitJSONBody
