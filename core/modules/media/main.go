package media

import (
	coreglobals "collexy/core/globals"
	"collexy/core/lib"
	coremodulemediacontrollers "collexy/core/modules/media/controllers"
	"log"
	"net/http"
)

func init() {
	log.Println("Registered module: Media")

	privateApiRouter := coreglobals.PrivateApiRouter
	subr := privateApiRouter.PathPrefix("/").Subrouter()

	mediaApiController := coremodulemediacontrollers.MediaApiController{}
	mediaTreeController := coremodulemediacontrollers.MediaTreeController{}

	// Content
	subr.HandleFunc("/api/media/{id:.*}/contextmenu", http.HandlerFunc(mediaTreeController.GetMenu)).Methods("GET")
	subr.HandleFunc("/api/media", http.HandlerFunc(mediaApiController.Get)).Methods("GET")
	
	subr.HandleFunc("/api/media/{id:.*}/children", http.HandlerFunc(mediaApiController.GetByIdChildren)).Methods("GET")
	subr.HandleFunc("/api/media/{id:.*}/parents", http.HandlerFunc(mediaApiController.GetByIdParents)).Methods("GET")
	
	subr.HandleFunc("/api/media/{id:.*}", http.HandlerFunc(mediaApiController.GetBackendMediaById)).Methods("GET")
	//privateApiRouter.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.Delete)).Methods("DELETE")
	//privateApiRouter.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.Post)).Methods("POST")
	
	privateApiRouter.HandleFunc("/api/media/{id:.*}", http.HandlerFunc(mediaApiController.Put)).Methods("PUT")

	// Setup FileServer for the settings module
	log.Println("Registered a handler for static files. [media::module]")
	http.Handle("/core/modules/media/public/", http.FileServer(http.Dir("./")))

	//////

	rMediaSection := lib.Route{"media", "/admin/media", "core/modules/media/public/views/media/index.html", false}

	rMediaTreeMethodEdit := lib.Route{"media.edit", "/edit/:id", "core/modules/media/public/views/media/edit.html", false}
	rMediaTreeMethodNew := lib.Route{"media.new", "/new?media_type_id&parent_id", "core/modules/media/public/views/media/new.html", false}

	// setup trees
	routesMediaTree := []lib.Route{rMediaTreeMethodEdit, rMediaTreeMethodNew}

	tMedia := lib.Tree{"Media", "media", routesMediaTree}

	treesMediaSection := []*lib.Tree{&tMedia}

	// params: name, alias, icon, route, trees, iscontainer, parent
	sMedia := lib.Section{"Media Section", "mediaSection", "fa fa-file-image-o fa-fw", &rMediaSection, treesMediaSection, true, nil, nil, []string{"media_section"}}

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
	sections := []lib.Section{sMedia}
	// params: name, alias, description, sections
	moduleSettings := lib.Module{"Media Module", "mediaModule", "Just a media module", sections}

	// register module
	lib.RegisterModule(moduleSettings)
}
