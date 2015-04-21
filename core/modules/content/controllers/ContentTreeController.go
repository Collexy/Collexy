package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/content/models"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	//"reflect"
)

type ContentTreeController struct{}

func (this *ContentTreeController) GetMenuForContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	var cmItems []ContextMenuItem

	//Name, Path, Url, IsDialog, Items
	cmiNew := ContextMenuItem{"Create", "", "", "", false, nil, "node_create"}

	// TODO
	// Content types should have added an AllowAtRoot field, to help determine ContextMenuItems at root level
	// Also a IsContainer field needs to be added so container content types will not appear in context menu?????
	// Permissions should be added again
	if id == 0 {
		contentTypes := coremodulesettingsmodels.GetContentTypes(nil)

		for _, ct := range contentTypes {
			if ct.TypeId == 1 && ct.AllowAtRoot{
				path := fmt.Sprintf("content.new({type_id:%d, content_type_id:%d})", 1, ct.Id)
				//tempIdStr := strconv.Itoa(ctId)
				item := ContextMenuItem{ct.Name, path, ct.Icon, "", false, nil, ""}
				cmiNew.Items = append(cmiNew.Items, item)
			}
		}

		cmItems = append(cmItems, cmiNew)
	} else {
		c := models.GetContentById(id)

		// var myIntSlice []int

		// s := reflect.ValueOf(c.ContentType.Meta["allowed_content_type_ids"])

		// for i := 0; i < s.Len(); i++ {
		//  myinterface := s.Index(i).Interface()
		//  lol := reflect.ValueOf(myinterface).Float()
		//     fmt.Println(lol)

		//     myIntSlice = append(myIntSlice,int(lol))
		// }

		for _, ct := range c.ContentType.AllowedContentTypes {
			path := fmt.Sprintf("content.new({type_id:%d, content_type_id:%d, parent_id:%d})", c.TypeId, ct.Id, c.Id)
			//tempIdStr := strconv.Itoa(ctId)
			item := ContextMenuItem{ct.Name, path, ct.Icon, "", false, nil, ""}
			cmiNew.Items = append(cmiNew.Items, item)
		}
		cmiDel := ContextMenuItem{"Delete", "", "", "core/modules/content/public/views/content/delete.html", true, nil, ""}
		cmItems = append(cmItems, cmiNew)
		cmItems = append(cmItems, cmiDel)
	}

	res, err := json.Marshal(cmItems)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

type ContextMenuItem struct {
	Name       string            `json:"name,omitempty"`
	Path       string            `json:"path,omitempty"`
	CssClass   string            `json:"css_class,omitempty"`
	Url        string            `json:"url,omitempty"`
	IsDialog   bool              `json:"is_dialog"`
	Items      []ContextMenuItem `json:"items,omitempty"`
	Permission string            `json:"permission,omitempty"`
}
