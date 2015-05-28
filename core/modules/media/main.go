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

// func MediaProtectHandler(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(*r.URL)
// 		var urlStr string = ""

// 		if _, err := os.Stat("./config/media-access.json"); err != nil {
// 			if os.IsNotExist(err) {
// 				// file does not exist
// 				log.Println("media-access.json config file does not exist")
// 			} else {
// 				// other error
// 			}
// 		} else {

// 			configFile, err1 := os.Open("./config/media-access.json")
// 			defer configFile.Close()
// 			if err1 != nil {
// 				log.Println("Error opening media-access.json config file")
// 				//printError("opening config file", err1.Error())
// 			}

// 			jsonParser := json.NewDecoder(configFile)
// 			if err1 = jsonParser.Decode(&coreglobals.MediaAccessConf); err1 != nil {
// 				log.Println("Error parsing media-access.json config file")
// 				log.Println(err1.Error())
// 				//printError("parsing config file", err1.Error())
// 			}

// 			urlStr = r.URL.Path
// 			// urlStrTemp, err2 := r.URL.Path()
// 			// if err2!=nil{
// 			// 	log.Println(err2.Error())

// 			// } else{
// 			// 	urlStrTemp = urlStrTemp
// 			// }

// 			log.Println("urlStr: " + urlStr)
// 			log.Println(coreglobals.MediaAccessConf[urlStr])
// 			//log.Println(url.QueryEscape(urlStr))
// 			log.Println("111!!!!!1")
// 			// log.Println(coreglobals.Maccess.Items[0].Domains[0])
// 			// log.Println(coreglobals.Maccess.Items[0].Url)
// 			// fmt.Println(coreglobals.Maccess.Items[0].MemberGroups)
// 		}

// 		// fmt.Println(r.URL.Path)
// 		// fmt.Println(coreglobals.Maccess.Domains[0] + "/" + coreglobals.Maccess.Url)
// 		// fmt.Println(r.Host)

// 		isProtected := false
// 		hasAccess := false
// 		var protectedItem *coreglobals.MediaAccessItem

// 		if val, ok := coreglobals.MediaAccessConf[urlStr]; ok {
// 			isProtected = true
// 			protectedItem = val
// 		}

// 		// var protectedItem *coreglobals.MediaAccessItem = nil

// 		// for _, maItem := range coreglobals.Maccess.Items {
// 		// 	if isProtected {
// 		// 		break;
// 		// 	}
// 		// 	for _, domain := range maItem.Domains {
// 		// 		if isProtected {
// 		// 			break;
// 		// 		}
// 		// 		if domain + "/" + maItem.Url == r.Host + "/" + r.URL.Path{
// 		// 			if isProtected {
// 		// 				break;
// 		// 			}
// 		// 			isProtected = true;
// 		// 			protectedItem = &maItem
// 		// 			// fmt.Fprintf(w, "loldalolselol")

// 		// 		}
// 		// 	}
// 		// }
// 		if isProtected {
			
// 			SetProtectedMediaKey(r,protectedItem)

// 			sid := corehelpers.CheckMemberCookie(w, r)

// 			m, err := coremodulemembermodels.GetMember(sid)

// 			if m == nil || err == sql.ErrNoRows {
// 				fmt.Println(err)
// 				// hasAccess = false //already set when var was initialized

// 			} else {
// 				coremodulemembermodels.SetLoggedInMember(r, m)
// 				for _, mg := range m.Groups {
// 					if hasAccess {
// 						break
// 					}
// 					// fmt.Println("MEMBER :::::: ")
// 					// fmt.Println(mg)
// 					// fmt.Println(protectedItem.MemberGroups[0])
// 					if mg.Id == protectedItem.MemberGroups[0] {
// 						// fmt.Println("workz?")
// 						hasAccess = true
// 					}
// 				}

// 			}
// 			if !hasAccess {
// 				fmt.Fprintf(w, "You need to be logged in to access this media item.")
// 			} else {
// 				h.ServeHTTP(w, r)
// 			}

// 		} else {
// 			h.ServeHTTP(w, r)
// 		}

// 		//sid := corehelpers.CheckCookie(w, r)

// 		//m, err := coremodulemembermodels.GetMember(sid)

// 		// if m == nil || err == sql.ErrNoRows {
// 		// 	fmt.Println(err)
// 		// 	fmt.Fprintf(w, "You need to be logged in to access the API.")

// 		// } else {
// 		// 	coremodulemembermodels.SetLoggedInMember(r, m)

// 		// 	h.ServeHTTP(w, r)
// 		// }
// 	})
// }

// // // type key int // already defined in user module.models

// // const myProtectedMediakey key = 1

// // GetProtectedMedia returns a value for this package from the request values.
// func GetProtectedMedia(r *http.Request) *coreglobals.MediaAccessItem {
// 	if rv := context.Get(r, coreglobals.MyProtectedMediakey); rv != nil {
// 		return rv.(*coreglobals.MediaAccessItem)
// 	}
// 	return nil
// }

// func SetProtectedMediaKey(r *http.Request, val *coreglobals.MediaAccessItem) {
//     context.Set(r, coreglobals.MyProtectedMediakey, val)
// }