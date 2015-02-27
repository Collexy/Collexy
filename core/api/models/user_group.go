package models

type UserGroup struct {
  Id int `json:"id"`
  Name string `json:"name, omitempty"`
  PermissionIds []int `json:"permission_ids,omitempty"`
  AngularRouteIds []int `json:"angular_route_ids,omitempty"`
  Permissions []string `json:"permissions,omitempty"`
}