package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/media/models"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	//"reflect"
)

type MediaTreeController struct{}

func (this *MediaTreeController) GetMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	var cmItems []ContextMenuItem

	//Name, Path, Url, IsDialog, Items
	cmiNew := ContextMenuItem{"Create", "", "", "", false, nil, "node_create"}

	// TODO
	// Media types should have added an AllowAtRoot field, to help determine ContextMenuItems at root level
	// Also a IsContainer field needs to be added so container media types will not appear in context menu?????
	// Permissions should be added again
	if id == 0 {
		mediaTypes := coremodulesettingsmodels.GetMediaTypes(nil)

		for _, ct := range mediaTypes {
			if ct.AllowAtRoot {
				path := fmt.Sprintf("media.new({media_type_id:%d, parent_id:null})", ct.Id)
				//tempIdStr := strconv.Itoa(ctId)
				item := ContextMenuItem{ct.Name, path, ct.Icon, "", false, nil, ""}
				cmiNew.Items = append(cmiNew.Items, item)
			}
		}
		if len(cmiNew.Items) > 0 {
			cmItems = append(cmItems, cmiNew)
		}

	} else {
		c := models.GetMediaById(id)

		// var myIntSlice []int

		// s := reflect.ValueOf(c.MediaType.Meta["allowed_media_type_ids"])

		// for i := 0; i < s.Len(); i++ {
		//  myinterface := s.Index(i).Interface()
		//  lol := reflect.ValueOf(myinterface).Float()
		//     fmt.Println(lol)

		//     myIntSlice = append(myIntSlice,int(lol))
		// }

		for _, ct := range c.MediaType.AllowedMediaTypes {
			path := fmt.Sprintf("media.new({media_type_id:%d, parent_id:%d})", ct.Id, c.Id)
			//tempIdStr := strconv.Itoa(ctId)
			item := ContextMenuItem{ct.Name, path, ct.Icon, "", false, nil, ""}
			cmiNew.Items = append(cmiNew.Items, item)
		}
		cmiDel := ContextMenuItem{"Delete", "", "", "core/modules/media/public/views/media/delete.html", true, nil, ""}
		if len(cmiNew.Items) > 0 {
			cmItems = append(cmItems, cmiNew)
		}
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
