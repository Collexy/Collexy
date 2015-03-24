package models

import
(
	"fmt"
	"encoding/json"
	"net/http"
	corehelpers "collexy/core/helpers"
	"github.com/gorilla/mux"
	"strconv"
	coreglobals "collexy/core/globals"
)

type AdminMenu struct {
	// Id int `json:"id"`
	Name string `json:"name,omitempty"`
	Items []AdminMenuItem `json:"items,omitempty"`
}

func (menu *AdminMenu) AddItem (item AdminMenuItem){
	menu.Items = append(menu.Items, item)
}

type AdminMenuItem struct {
	Name string `json:"name"`
	CssClass string `json:"css_class,omitempty"`
	Route *AdminRoute `json:"route,omitempty"`
	SubMenu *AdminMenu `json:"sub_menu,omitempty"`
}

type AdminRoute struct {
	//Name string `json:"name"`
	Alias string `json:"alias"`
	// Path string `json:"name"`
	//Parent *AdminRoute `json:"parent,omitempty"`
	Children []*AdminRoute `json:"children,omitempty"`
	// Type int `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Components []map[string]interface{} `json:"components,omitempty"`  
	RedirectTo string `json:"redirect_to,omitempty"`  
	Data map[string]interface{} `json:"data,omitempty"`  
	// Ref string `json:"ref,omitempty"`
}

func (this *AdminRoute) AddChildren (child coreglobals.IRoute){
	c := child.(*AdminRoute)
	this.Children = append(this.Children, c)
}

func Test(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content-Type", "application/json")

	var routes []coreglobals.IRoute 
	/**
	* Menu
	*/
	am := AdminMenu{"Main", nil}

	
	/**
	* Dashboard
	*/

	// Routes
	var components []map[string]interface{}
	component := map[string]interface{}{
		"single":"public/views/admin/dashboard.html",
	}
	//components = [0]map[string]interface{}{"single":"public/views/admin/dashboard.html"}
	components = append(components, component)
	rDashboard := AdminRoute{"index", nil, "/admin", components, "", nil}

	// Menu items
	miDashboard := AdminMenuItem {"Dashboard", "fa fa-dashboard fa-fw", &rDashboard, nil}
	am.AddItem(miDashboard)

	/**
	* Settings
	*/

	// Sub menus
	smSettings := AdminMenu{"Settings", nil}

	// Routes
	components = nil
	component = map[string]interface{}{
		"single":"public/views/settings/index.html",
	}
	components = append(components, component)
	//components = [0]map[string]interface{}{"single":"public/views/settings/index.html"}
	rSettings := AdminRoute{"settings", nil, "/admin/settings", components, "", nil}

	//
	components = nil
	component = map[string]interface{}{
		"single":"public/views/settings/content-type/index.html",
	}
	components = append(components, component)
	//components = [0]map[string]interface{}{"single":"public/views/settings/content-type/index.html"}
	rSettingsContentTypes := AdminRoute{"contentType", nil, "/content-type", components, "", nil}

	components = nil
	component = map[string]interface{}{
		"single":"public/views/settings/content-type/new.html",
	}
	components = append(components, component)
	//components = [0]map[string]interface{}{"single":"public/views/settings/content-type/new.html"}
	rSettingsContentTypesNew := AdminRoute{"new", nil, "/new?type&parent", components, "", nil}

	components = nil
	component = map[string]interface{}{
		"single":"public/views/settings/content-type/edit.html",
	}
	components = append(components, component)
	//components = [0]map[string]interface{}{"single":"public/views/settings/content-type/edit.html"}
	rSettingsContentTypesEdit := AdminRoute{"edit", nil, "/edit/:nodeId", components, "", nil}


	// Menu items

	

	miSettingsContentTypes := AdminMenuItem {"Content Types", "fa fa-newspaper-o fa-fw", &rSettingsContentTypes, nil}

	smSettings.AddItem(miSettingsContentTypes)

	miSettings := AdminMenuItem {"Settings", "fa fa-gear fa-fw", &rSettings, &smSettings}

	am.AddItem(miSettings)

	rSettingsContentTypes.AddChildren(&rSettingsContentTypesNew)
	rSettingsContentTypes.AddChildren(&rSettingsContentTypesEdit)
	rSettings.AddChildren(&rSettingsContentTypes)

	routes = append(routes, &rDashboard)
	routes = append(routes, &rSettings)

	res, err := json.Marshal(am)
    corehelpers.PanicIf(err)

    //coreglobals.Routes = routes

    fmt.Fprintf(w,"%s",res)
}

func CmTest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
    typeStr := params["nodeType"]

    nodeType, _ := strconv.Atoi(typeStr)

	var contentTypeItems []ContextMenuItem
	//contentTypeItem := ContextMenuItem{"content.new({node_type:1, content_type_node_id: ct.node.id, parent_id:{{currentItem.id}}})", "", nil, nil}
	
	
	//contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/content/delete.html", "", nil,"node_delete", true}
	//contentTypeItems = append(contentTypeItems, contentTypeItemDelete)

	// this could be refactored away easily, AND SHOULD!!!!
	if(nodeType == 1){
		cmNew := ContextMenu{nil}
		contentTypeItemNew := ContextMenuItem{"New", "", "", "",&cmNew,"node_create", false}
		contentTypeItems = append(contentTypeItems, contentTypeItemNew)
		contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/content/delete.html", "", nil,"node_delete", true}
		// contentTypeItemDelete := ContextMenuItem{"Delete", "", "content.delete({nodeId: {{currentItem.id}}})", "", nil,2, true}
		contentTypeItems = append(contentTypeItems, contentTypeItemDelete)
		
	} else if(nodeType == 2){ 
		cmNew := ContextMenu{nil}
		contentTypeItemNew := ContextMenuItem{"New", "", "", "",&cmNew,"node_create", false}
		contentTypeItems = append(contentTypeItems, contentTypeItemNew)
		//fmt.Fprintf(w,"nodetype is: %d", nodeType)
		contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/media/delete.html", "", nil,"node_delete", true}
		// contentTypeItemDelete := ContextMenuItem{"Delete", "", "content.delete({nodeId: {{currentItem.id}}})", "", nil,2, true}
		contentTypeItems = append(contentTypeItems, contentTypeItemDelete)

	} else if(nodeType == 3){
		contentTypeItemNew := ContextMenuItem{"New", "", "settings.templates.new({type:3, parent:{{currentItem.id}}})", "fa fa-folder fa-fw",nil,"node_create", false}
		contentTypeItems = append(contentTypeItems, contentTypeItemNew)
		//fmt.Fprintf(w,"nodetype is: %d", nodeType)
		contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/settings/template/delete.html", "", nil,"node_delete", true}
		// contentTypeItemDelete := ContextMenuItem{"Delete", "", "content.delete({nodeId: {{currentItem.id}}})", "", nil,2, true}
		contentTypeItems = append(contentTypeItems, contentTypeItemDelete)

	} else if(nodeType == 4){ 
		
		contentTypeItemNew := ContextMenuItem{"New", "", "settings.contentTypes.new({type:4, parent:{{currentItem.id}}})", "fa fa-folder fa-fw",nil,"node_create", false}
		contentTypeItems = append(contentTypeItems, contentTypeItemNew)
		//fmt.Fprintf(w,"nodetype is: %d", nodeType)
		contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/settings/content-type/delete.html", "", nil,"node_delete", true}
		// contentTypeItemDelete := ContextMenuItem{"Delete", "", "content.delete({nodeId: {{currentItem.id}}})", "", nil,2, true}
		contentTypeItems = append(contentTypeItems, contentTypeItemDelete)

	} else if(nodeType == 5){ 
		// ROOT NODE
	} else if(nodeType == 7){ 
		contentTypeItemNew := ContextMenuItem{"New", "", "settings.mediaTypes.new({type:7, parent:{{currentItem.id}}})", "fa fa-folder fa-fw",nil,"node_create", false}
		contentTypeItems = append(contentTypeItems, contentTypeItemNew)
		//fmt.Fprintf(w,"nodetype is: %d", nodeType)
		contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/settings/media-type/delete.html", "", nil,"node_delete", true}
		// contentTypeItemDelete := ContextMenuItem{"Delete", "", "content.delete({nodeId: {{currentItem.id}}})", "", nil,2, true}
		contentTypeItems = append(contentTypeItems, contentTypeItemDelete)

	} else if(nodeType == 9){ 
		cmNew := ContextMenu{nil}
		contentTypeItemNew := ContextMenuItem{"New", "", "", "",&cmNew,"node_create", false}
		contentTypeItems = append(contentTypeItems, contentTypeItemNew)
		//fmt.Fprintf(w,"nodetype is: %d", nodeType)
		contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/settings/stylesheet/delete.html", "", nil,"node_delete", true}
		// contentTypeItemDelete := ContextMenuItem{"Delete", "", "content.delete({nodeId: {{currentItem.id}}})", "", nil,2, true}
		contentTypeItems = append(contentTypeItems, contentTypeItemDelete)

	} else if(nodeType == 10){ 
		
		contentTypeItemNewFolder := ContextMenuItem{"Folder", "", "settings.scripts.new({type:'folder', parent:currentItem.path})", "fa fa-folder fa-fw",nil,"node_create", false}
		contentTypeItemNewFile := ContextMenuItem{"File", "", "settings.scripts.new({type:'file', parent:currentItem.path})", "fa fa-file fa-fw",nil,"node_create", false}
		var directoryItems []ContextMenuItem
		directoryItems = append(directoryItems, contentTypeItemNewFolder)
		directoryItems = append(directoryItems, contentTypeItemNewFile)

		cmNew := ContextMenu{directoryItems}

		contentTypeItemNew := ContextMenuItem{"New", "", "", "",&cmNew,"node_create", false}
		contentTypeItems = append(contentTypeItems, contentTypeItemNew)
		//fmt.Fprintf(w,"nodetype is: %d", nodeType)
		contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/settings/script/delete.html", "", nil,"node_delete", true}
		// contentTypeItemDelete := ContextMenuItem{"Delete", "", "content.delete({nodeId: {{currentItem.id}}})", "", nil,2, true}
		contentTypeItems = append(contentTypeItems, contentTypeItemDelete)

	} else if(nodeType == 11){ 
		contentTypeItemNew := ContextMenuItem{"New", "", "settings.dataTypes.new({type:11, parent:{{currentItem.id}}})", "fa fa-folder fa-fw",nil,"node_create", false}
		contentTypeItems = append(contentTypeItems, contentTypeItemNew)
		//fmt.Fprintf(w,"nodetype is: %d", nodeType)
		contentTypeItemDelete := ContextMenuItem{"Delete", "", "public/views/settings/data-type/delete.html", "", nil,"node_delete", true}
		// contentTypeItemDelete := ContextMenuItem{"Delete", "", "content.delete({nodeId: {{currentItem.id}}})", "", nil,2, true}
		contentTypeItems = append(contentTypeItems, contentTypeItemDelete)
	}
	res, err := json.Marshal(contentTypeItems)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w,"%s",res)
}

type ContextMenu struct {
	ContextMenuItems []ContextMenuItem `json:"items,omitempty"`
}

// <a prevent-default="" ng-click="executeMenuItem(action)">
// getallowedchildren?contentId=x

type ContextMenuItem struct {
	Name string `json:"name"`
	Alias string `json:"alias"`
	Url string `json:"url,omitempty"`
	CssClass string `json:"css_class,omitempty"`
	SubMenu *ContextMenu `json:"sub_menu,omitempty"`
	Permission string `json:"permission,omitempty"`
	IsDialog bool `json:"is_dialog"`
}

func GetMenu(nodeId int, nodeType int){
	// get node from db where id=nodeId}
	// if node.type = x
	// 		load default type items slice
	// 			foreach node permissions
	// 				foreach default items slice
	//					if default_item[j].permissions
	//						add to allowedSlice
	// cm := ContextMenu{allowedSlice}
	// return
}

type Action struct {
	
}