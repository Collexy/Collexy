package models

// this needs to be moved to lib in app core since, this is used both for content and media modules

import (
	coreglobals "collexy/core/globals"
)

type Permission struct {
	Name string `json:"name,omitempty"`
}

type PermissionsContainer struct {
	Id          int                     `json:"id"`
	Permissions coreglobals.StringSlice `json:"permissions"` //map[string]struct{} `json:"permissions"`
}

type PermissionTest struct {
	Permissions coreglobals.StringSlice `json:"permissions"` //map[string]struct{} `json:"permissions"`
}