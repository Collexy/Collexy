package controllers

import (
	corehelpers "collexy/core/helpers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	//"reflect"
	coremodulemembermodels "collexy/core/modules/member/models"
)

type MemberTreeController struct{}

func (this *MemberTreeController) GetMenuForMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	var cmItems []ContextMenuItem

	cmiNew := ContextMenuItem{"Create", "", "", "", false, nil, "node_create"}

	// TODO
	// Content types should have added an AllowAtRoot field, to help determine ContextMenuItems at root level
	// Also a IsContainer field needs to be added so container content types will not appear in context menu?????
	// Permissions should be added again
	
		memberTypes := coremodulemembermodels.GetMemberTypes()

		for _, mt := range memberTypes {
			//if mt.TypeId == 1 {
			// if mt.TypeId == 1 && mt.AllowAtRoot{
				path := fmt.Sprintf("member.new({member_type_id:%d})", mt.Id)
				//tempIdStr := strconv.Itoa(ctId)
				item := ContextMenuItem{mt.Name, path, mt.Icon, "", false, nil, ""}
				cmiNew.Items = append(cmiNew.Items, item)
			//}
		}
		if len(cmiNew.Items) > 0 {
			cmItems = append(cmItems, cmiNew)
		}

	if id != 0 {
		cmiDel := ContextMenuItem{"Delete", "", "", "core/modules/member/public/views/member/delete.html", true, nil, ""}
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

type ContextMenuItem struct {
	Name       string            `json:"name,omitempty"`
	Path       string            `json:"path,omitempty"`
	CssClass   string            `json:"css_class,omitempty"`
	Url        string            `json:"url,omitempty"`
	IsDialog   bool              `json:"is_dialog"`
	Items      []ContextMenuItem `json:"items,omitempty"`
	Permission string            `json:"permission,omitempty"`
}
