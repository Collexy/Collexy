package models

import (
  //"fmt"
  //"encoding/json"
  //"time"
)

type TabProperty struct {
  //Id int `json:"id"`
  Name string `json:"name,omitempty"`
  Order int `json:"order,omitempty"`
  DataTypeId int `json:"data_type_id,omitempty"`
  DataType *DataType `json:"data_type,omitempty"`
  HelpText string `json:"help_text,omitempty"`
  Description string `json:"description,omitempty"`
}

// type tabProperty TabProperty
// type dataType DataType

// func (t *TabProperty) UnmarshalJSON(b []byte) (err error) {
// 	j, n, o, d, h, de := tabProperty{}, "", 0, DataType{}, "", ""
// 	if err = json.Unmarshal(b, &j); err == nil {
// 		*t = TabProperty(j)
// 		return
// 	}
// 	if err = json.Unmarshal(b, &n); err == nil {
// 		t.Name = n
// 		return
// 	}
// 	if err = json.Unmarshal(b, &o); err == nil {
// 		t.Order = o
// 		return
// 	}
// 	if err = json.Unmarshal(b, &d); err == nil {
// 		t.DataType = d
// 	}
// 	if err = json.Unmarshal(b, &h); err == nil {
// 		t.HelpText = h
// 	}
// 	if err = json.Unmarshal(b, &de); err == nil {
// 		t.Description = de
// 	}
// 	return
// }

// func (t *TabProperty) MarshalJSON() ([]byte, error) {
//   newTp := TabProperty{}
//   //newDt := DataType{}
//   if(t.Name != ""){
//     newTp.Name = t.Name
//   } 
//   if(t.Order != 0){
//     newTp.Order = t.Order
//   } 
//   if(t.DataTypeId != 0){
//     newTp.DataTypeId = t.DataTypeId
//   } 
//   if(t.DataType.NodeId != 0){
//     newTp.DataType = t.DataType
//   } else{
//     newTp.DataType = DataType{}
//   }
//   if(t.HelpText != ""){
//     newTp.HelpText = t.HelpText
//   }
//   if(t.Description != ""){
//     newTp.Description = t.Description
//   }

//   return json.Marshal(newTp)
// }