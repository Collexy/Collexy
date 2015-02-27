package models

type PublicAccess struct {
  Members []int `json:"members,omitempty"`
  Groups []int `json:"groups,omitempty"`
}