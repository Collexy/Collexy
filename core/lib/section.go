package lib

import (
// "reflect"
)

type Section struct {
	Name        string    `json:"name,omitempty"`
	Alias       string    `json:"alias,omitempty"`
	Icon        string    `json:"icon,omitempty"`
	Route       *Route    `json:"route,omitempty"`
	Trees       []*Tree   `json:"trees,omitempty"`
	IsContainer bool      `json:"is_container,omitempty"`
	Parent      *Section  `json:"parent,omitempty"`
	Children    []Section `json:"children,omitempty"`
	Permissions []string  `json:"permissions,omitempty"`
	//ContextMenu *ContextMenu
	//SubSections []*Section
}

// func (this *Section) SetChildren(s []Section){
// 	//f := &Section{}

// 	fVal := reflect.ValueOf(this).Elem()
// 	sVal := reflect.ValueOf(s)

// 	fVal.FieldByName("Children").Set(sVal)

// 	//this.Children = s
// }

// func (this *Section) SetupTree(){}
