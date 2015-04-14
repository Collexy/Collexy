package content

import
(
	"log"
	"net/http"
	"collexy/core/lib"
	coreglobals "collexy/core/globals"
	coremodulecontentcontrollers "collexy/core/modules/content/controllers"
)

func init(){
	log.Println("Registered module: Content")

	privateApiRouter := coreglobals.PrivateApiRouter
    subr := privateApiRouter.PathPrefix("/").Subrouter()

    contentApiController := coremodulecontentcontrollers.ContentApiController{}
    
    // Content
    subr.HandleFunc("/api/content", http.HandlerFunc(contentApiController.Get)).Methods("GET")
    subr.HandleFunc("/api/content/{id:.*}/children", http.HandlerFunc(contentApiController.GetByIdChildren)).Methods("GET")
    //privateApiRouter.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.Delete)).Methods("DELETE")
	//privateApiRouter.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.Post)).Methods("POST")
    subr.HandleFunc("/api/content/{id:.*}", http.HandlerFunc(contentApiController.GetBackendContentById)).Methods("GET")
    //privateApiRouter.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.PutContent)).Methods("PUT")

    subr.HandleFunc("/api/media/{id:.*}", http.HandlerFunc(contentApiController.GetBackendContentById)).Methods("GET")

	// Setup FileServer for the settings module
	log.Println("Registered a handler for static files. [content::module]")
	http.Handle("/core/modules/content/public/", http.FileServer(http.Dir("./")))





	//////





	rContentSection := lib.Route{"content", "/admin/content", "core/modules/content/public/views/content/index.html", false}
	rMediaSection := lib.Route{"media", "/admin/media", "core/modules/content/public/views/media/index.html", false}
	
	rContentTreeMethodEdit := lib.Route{"content.edit", "/edit/:id", "core/modules/content/public/views/content/edit.html", false}
	rContentTreeMethodNew := lib.Route{"content.new", "/new?type_id&content_type_id&parent_id", "core/modules/content/public/views/content/new.html", false}

	rMediaTreeMethodEdit := lib.Route{"media.edit", "/edit/:id", "core/modules/content/public/views/media/edit.html", false}
	rMediaTreeMethodNew := lib.Route{"media.new", "/new?type_id&content_type_id&parent_id", "core/modules/content/public/views/media/new.html", false}

	// setup trees
	routesContentTree := []lib.Route{rContentTreeMethodEdit, rContentTreeMethodNew}
	routesMediaTree := []lib.Route{rMediaTreeMethodEdit, rMediaTreeMethodNew}
	
	tContent := lib.Tree{"Content", "content", routesContentTree}
	tMedia := lib.Tree{"Media", "media", routesMediaTree}

	treesContentSection := []*lib.Tree{&tContent}
	treesMediaSection := []*lib.Tree{&tMedia}

	// params: name, alias, icon, route, trees, iscontainer, parent
	sContent := lib.Section{"Content Section", "contentSection", "fa fa-newspaper-o fa-fw", &rContentSection, treesContentSection, true, nil,nil}
	sMedia := lib.Section{"Media Section", "mediaSection", "fa fa-file-image-o fa-fw", &rMediaSection, treesMediaSection, true, nil,nil}

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
	sections := []lib.Section{sContent,sMedia}
	// params: name, alias, description, sections
	moduleSettings := lib.Module{"Content Module", "contentModule", "Just a content module", sections}

	// register module
	lib.RegisterModule(moduleSettings)
}