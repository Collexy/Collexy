package controllers

import (
	corehelpers "collexy/core/helpers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	//"strconv"
	//"reflect"
)

type PermissionTreeController struct{}

func (this *PermissionTreeController) GetMenuForPermission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	//id, _ := strconv.Atoi(idStr)

	fmt.Println(idStr)

	var cmItems []ContextMenuItem

	//Name, Path, Url, IsDialog, Items

	// TODO
	// Content types should have added an AllowAtRoot field, to help determine ContextMenuItems at root level
	// Also a IsContainer field needs to be added so container content types will not appear in context menu?????
	// Permissions should be added again
	if idStr == "0" {
		cmiNew := ContextMenuItem{"Create", "user.permission.new", "", "", false, nil, "node_create"}
		cmItems = append(cmItems, cmiNew)
	} else {
		cmiDel := ContextMenuItem{"Delete", "", "", "core/modules/user/public/views/permission/delete.html", true, nil, ""}
		cmiNew := ContextMenuItem{"Create", "user.permission.new", "", "", false, nil, "node_create"}
		cmItems = append(cmItems, cmiNew)
		cmItems = append(cmItems, cmiDel)
	}

	res, err := json.Marshal(cmItems)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

// type ContextMenuItem struct {
// 	Name       string            `json:"name,omitempty"`
// 	Path       string            `json:"path,omitempty"`
// 	CssClass   string            `json:"css_class,omitempty"`
// 	Url        string            `json:"url,omitempty"`
// 	IsDialog   bool              `json:"is_dialog"`
// 	Items      []ContextMenuItem `json:"items,omitempty"`
// 	Permission string            `json:"permission,omitempty"`
// }
