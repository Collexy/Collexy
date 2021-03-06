package content

import (
	coreglobals "collexy/core/globals"
	"collexy/core/lib"
	coremodulecontentcontrollers "collexy/core/modules/content/controllers"
	"log"
	"net/http"
)

func init() {
	log.Println("Registered module: Content")

	privateApiRouter := coreglobals.PrivateApiRouter
	subr := privateApiRouter.PathPrefix("/").Subrouter()

	contentApiController := coremodulecontentcontrollers.ContentApiController{}
	contentTreeController := coremodulecontentcontrollers.ContentTreeController{}

	// Content
	subr.HandleFunc("/api/content/{id:.*}/contextmenu", http.HandlerFunc(contentTreeController.GetMenuForContent)).Methods("GET")
	subr.HandleFunc("/api/content", http.HandlerFunc(contentApiController.Get)).Methods("GET")
	subr.HandleFunc("/api/content/{id:.*}/children", http.HandlerFunc(contentApiController.GetByIdChildren)).Methods("GET")
	subr.HandleFunc("/api/content/{id:.*}/parents", http.HandlerFunc(contentApiController.GetByIdParents)).Methods("GET")
	subr.HandleFunc("/api/content/{id:.*}", http.HandlerFunc(contentApiController.GetBackendContentById)).Methods("GET")

	subr.HandleFunc("/api/content/{id:.*}", http.HandlerFunc(contentApiController.Post)).Methods("POST")
	subr.HandleFunc("/api/content/{id:.*}", http.HandlerFunc(contentApiController.Put)).Methods("PUT")
	subr.HandleFunc("/api/content/{id:.*}", http.HandlerFunc(contentApiController.Delete)).Methods("DELETE")

	// Setup FileServer for the settings module
	log.Println("Registered a handler for static files. [content::module]")
	http.Handle("/core/modules/content/public/", http.FileServer(http.Dir("./")))

	//////

	rContentSection := lib.Route{"content", "/admin/content", "core/modules/content/public/views/content/index.html", false}

	rContentTreeMethodEdit := lib.Route{"content.edit", "/edit/:id", "core/modules/content/public/views/content/edit.html", false}
	rContentTreeMethodNew := lib.Route{"content.new", "/new?type_id&content_type_id&parent_id", "core/modules/content/public/views/content/new.html", false}

	// setup trees
	routesContentTree := []lib.Route{rContentTreeMethodEdit, rContentTreeMethodNew}

	tContent := lib.Tree{"Content", "content", routesContentTree}

	treesContentSection := []*lib.Tree{&tContent}

	// params: name, alias, icon, route, trees, iscontainer, parent
	sContent := lib.Section{"Content Section", "contentSection", "fa fa-newspaper-o fa-fw", &rContentSection, treesContentSection, true, nil, nil, []string{"content_section"}}

	//reflect.ValueOf(&sSettings).Elem().FieldByName("Children").Set(reflect.ValueOf(lol))
	// log.Println(sSettings.Children[0].Name + ":FDSF:SDF:DS:F:")
	// res, err := json.Marshal(sSettings)
	// if err!=nil{
	// 	panic(err)
	// }
	// log.Println(res)
	// log.Println("__-----------------")
	//sSettings.SetChildren(lol)
	//sSettings.Children =

	// setup settings section
	// maybe add IsContainer bool?
	// section parent by name or section children - or both? For subsections ofc

	// setup subsections

	// setup module
	sections := []lib.Section{sContent}
	// params: name, alias, description, sections
	moduleSettings := lib.Module{"Content Module", "contentModule", "Just a content module", sections, nil, 100}

	// register module
	lib.RegisterModule(moduleSettings)
}
