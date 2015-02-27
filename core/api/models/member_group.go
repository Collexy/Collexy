package models

type MemberGroup struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Description []int `json:"description,omitempty"`
}