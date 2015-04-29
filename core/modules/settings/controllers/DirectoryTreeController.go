package controllers

import (
	corehelpers "collexy/core/helpers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	//"reflect"
)

type DirectoryTreeController struct{}

func (this *DirectoryTreeController) GetMenuForDirectory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	rootdir := params["rootdir"]
	name := params["name"]
	isDir, _ := strconv.ParseBool(params["is_dir"])

	rootdir1 := rootdir[:len(rootdir)-1]
	//name, _ := strconv.Atoi(idStr)

	var cmItems []ContextMenuItem

	//Name, Path, Url, IsDialog, Items
	cmiNew := ContextMenuItem{"Create", "", "", "", false, nil, "node_create"}

	// TODO
	// Content types should have added an AllowAtRoot field, to help determine ContextMenuItems at root level
	// Also a IsContainer field needs to be added so container content types will not appear in context menu?????
	// Permissions should be added again
	if name == "root" {
		pathFile := fmt.Sprintf("settings.%s.new({type:'file', parent:null})", rootdir1)
		itemFile := ContextMenuItem{"file", pathFile, "fa fa-file-code-o fa-fw", "", false, nil, ""}
		pathFolder := fmt.Sprintf("settings.%s.new({type:'folder', parent:null})", rootdir1)
		itemFolder := ContextMenuItem{"folder", pathFolder, "fa fa-folder-o fa-fw", "", false, nil, ""}
		cmiNew.Items = append(cmiNew.Items, itemFile)
		cmiNew.Items = append(cmiNew.Items, itemFolder)
		if len(cmiNew.Items) > 0 {
			cmItems = append(cmItems, cmiNew)
		}
	} else {
		pathFile := fmt.Sprintf("settings.%s.new({type:'file', parent:%q})", rootdir1, name)
		itemFile := ContextMenuItem{"file", pathFile, "fa fa-file-code-o fa-fw", "", false, nil, ""}
		pathFolder := fmt.Sprintf("settings.%s.new({type:'folder', parent:null})", rootdir1)
		itemFolder := ContextMenuItem{"folder", pathFolder, "fa fa-folder-o fa-fw", "", false, nil, ""}
		cmiNew.Items = append(cmiNew.Items, itemFile)
		cmiNew.Items = append(cmiNew.Items, itemFolder)

		cmiDelPath := fmt.Sprintf("core/modules/settings/public/views/%s/delete.html", rootdir1)
		cmiDel := ContextMenuItem{"Delete", "", "", cmiDelPath, true, nil, ""}
		if len(cmiNew.Items) > 0 && isDir {
			cmItems = append(cmItems, cmiNew)
		}
		cmItems = append(cmItems, cmiDel)
	}

	//Name, Path, Url, IsDialog, Items

	// TODO
	// Content types should have added an AllowAtRoot field, to help determine ContextMenuItems at root level
	// Also a IsContainer field needs to be added so container content types will not appear in context menu?????
	// Permissions should be added again

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
