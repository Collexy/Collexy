package models

import (
//"fmt"
//"encoding/json"
//"time"
)

type Tab struct {
	//Id int `json:"id"`
	Name       string         `json:"name"`
	Properties []*TabProperty `json:"properties,omitempty"`
}

// type tab Tab

// func (t *Tab) UnmarshalJSON(b []byte) (err error) {
// 	j := tab{}
// 	var n []*TabProperty
// 	if err = json.Unmarshal(b, &j); err == nil {
// 		*t = Tab(j)
// 		return
// 	}
// 	if err = json.Unmarshal(b, &n); err == nil {
// 		t.Properties = n
// 		return
// 	}
// 	return
// }

// func (t *Tab) MarshalJSON() ([]byte, error) {
//     if t.Name != "" && t.Properties != nil {
//         return json.Marshal(map[string]interface{}{
//             "name": t.Name,
//             "properties": t.Properties,
//         })
//     }
//     if t.Name != "" {
//         return json.Marshal(t.Name)
//     }
//     if t.Properties != nil {
//         return json.Marshal(t.Properties)
//     }
//     return json.Marshal(nil)
// }
