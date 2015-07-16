package application

import (
	"collexy/core/application/controllers"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	//"net/url"
	//"collexy/core/application/controllers"
	//"collexy/core/api/models"
	corehelpers "collexy/core/helpers"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	//"collexy/globals"
	coreglobals "collexy/core/globals"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	//"collexy/helpers"
	"collexy/core/lib"
	"strings"
	// _ "collexy/core/modules/mytest"
	_ "collexy/modules/collexyecommerce"
	_ "collexy/core/modules/content"
	coremodulecontentcontrollers "collexy/core/modules/content/controllers"
	coremodulecontentmodels "collexy/core/modules/content/models"
	_ "collexy/core/modules/media"
	_ "collexy/core/modules/member"
	coremodulemembermodels "collexy/core/modules/member/models"
	_ "collexy/core/modules/settings"
	// coremodulesettingsmodels "collexy/core/modules/settings/models"
	_ "collexy/core/modules/user"
	coremoduleusermodels "collexy/core/modules/user/models"
	// "github.com/clbanning/mxj"
	"sort"
	"strconv"
	//"sync"
	//coreapplicationmodelsxmlmodels "collexy/core/application/models/xml_models"
	coreapplicationcontrollers "collexy/core/application/controllers"

)

func executeDatabaseInstallScript(site_title, username, password, email string, privacy bool) (err error) {
	db := coreglobals.Db
	tx, err1 := db.Begin()

	// fmt.Printf("title: %s, username: %s, password: %s, email: %s, privacy: %v\n", site_title, username, password, email, privacy)
	// rofl := `INSERT INTO "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids, permissions) VALUES (1, '%s', 'Admin', 'Demo', '%s', '%s', '2014-11-15 16:51:00.215', NULL, '2015-06-26 18:45:51.652', NULL, 1, 'Q6BZMGO6GTBLVLFW2Z2PHZSVLISDGGH6GVKNLS2V2ZS44LXT4ZZA', '{1}', NULL);`
	// lol := fmt.Sprintf(rofl, username, password, email)
	// fmt.Println(lol)

	sqlDDL := fmt.Sprintf(coreglobals.DbCreateScriptDDL, username, password, email)
	sqlStr := coreglobals.DbCreateScriptDML + sqlDDL
	fmt.Println(sqlStr)
	sqlStmtSlice := strings.Split(sqlStr, ";\r")

	if err != nil {
		return err
	}

	defer func() {
		// Rollback the transaction after the function returns.
		// If the transaction was already commited, this will do nothing.
		_ = tx.Rollback()
	}()

	for _, q := range sqlStmtSlice {
		//fmt.Println(q)
		// Execute the query in the transaction.
		_, err := tx.Exec(q)

		if err != nil {
			return err
		}
	}

	err1 = tx.Commit()
	err = err1
	return
}

// type ModuleItems struct {
// 	XMLName xml.Name           `xml:modules"`
// 	Modules   []*Module `xml:"module"`
// 	//Items map[string]*MediaItem `xml:"Item"` //does not work - how would you do it with maps?
// }

// type ModuleXML struct {
// 	XMLName xml.Name
// 	Module *Module `xml:"module"`
// }

// type XMLModel interface {
// 	func Post()
// }

// func doPost(i XMLModel){
// 	i.Post()
// 	if i.Children == nil {

// 	} else {
// 		if len(i.Children) == 0 {

// 		} else {
// 			for _, c := range i.Children {
// 				doPost(i)
// 			}
// 		}
// 	}
// }

// func executeModuleInstallScript(moduleDirectoryPath string) (err error) {
// 	moduleXMLFile, err1 := os.Open(moduleDirectoryPath + "/module.xml")
// 	defer moduleXMLFile.Close()
// 	if err1 != nil {
// 		log.Println("Error opening module.xml file")
// 		//printError("opening config file", err1.Error())
// 	}

// 	XMLdata, err2 := ioutil.ReadAll(moduleXMLFile) // use bufio intead since the xml can scale big

// 	// fmt.Println(string(XMLdata))

// 	if err2 != nil {
// 		log.Println("Error reading from module.xml file")
// 		fmt.Printf("error: %v", err2)
// 	}

// 	var v coreapplicationmodelsxmlmodels.Module

// 	err3 := xml.Unmarshal(XMLdata, &v)
// 	if err3 != nil {
// 		fmt.Printf("error: %v", err3)
// 	}

// 	// use recursive method / closures
// 	// for _, t := range v.Templates {
// 	// 	fmt.Println(t.Name)
// 	// 	for _, child1 := range t.Children {
// 	// 		fmt.Println(child1.Name)
// 	// 	}
// 	// }

// 	//dataTypes := v.DataTypes

// 	// for _, dt := range v.DataTypes {
// 	// 	dt.Post()
// 	// }

// 	// GET ALL DATATYPES FROM DB
// 	// use module controller

// 	for _, t := range v.Templates {
// 		t.Post(nil)
// 	}

// 	// GET ALL TEMPLATES FROM DB
// 	// use module controller

// 	var flatTemplatesSlice []*coremodulesettingsmodelsxmlmodels.Template

// 	coremodulesettingsmodelsxmlmodels.Walk(v.Templates, func(t *coremodulesettingsmodelsxmlmodels.Template) bool {
// 		flatTemplatesSlice = append(flatTemplatesSlice, t)
// 		return true
// 	})

// 	fmt.Printf("flatTemplatesSlice length: %d\n", len(flatTemplatesSlice))

// 	for _, ct := range v.ContentTypes {
// 		//fmt.Printf("sdflksjdkfjsdklf : %s\n", ct.Alias)
// 		ct.Post(nil, nil, flatTemplatesSlice)
// 	}

// 	// mimeTypes := v.MimeTypes

// 	// for _, mt := range v.MimeTypes {
// 	// 	mt.Post(nil, nil)
// 	// }

// 	// for _, mt := range v.MediaTypes {
// 	// 	mt.Post(nil, nil, mimeTypes)
// 	// }

// 	// var v coreglobals.MediaAccessItems
// 	// err := xml.Unmarshal(XMLdata, &v)
// 	// if err != nil {
// 	// 	fmt.Printf("error: %v", err)
// 	// 	return
// 	// }
// 	return
// }

// func testPostHandler(w http.ResponseWriter, r *http.Request) {
// 	//str := r.PostFormValue("collexy_module")
// 	if _, err := os.Stat("./_temporary_module_library/" + r.PostFormValue("collexy_module") + "/module.xml"); err != nil {
// 		if os.IsNotExist(err) {
// 			log.Println("XML file does not exist")
// 			fmt.Fprintf(w, "XML file does not exist")
// 		} else {
// 			// other error
// 		}
// 	} else {
// 		err5 := executeModuleInstallScript("./_temporary_module_library/" + r.PostFormValue("collexy_module"))
// 		if err5 != nil {
// 			log.Println("ERROR INSTALLING MODULE SCRIPT")
// 			log.Fatal(err5)
// 		} else {
// 			log.Println("MODULE SCRIPT INSTALLED SUCCESSFULLY")
// 			adminHandler(w, r)
// 		}
// 	}
// 	fmt.Fprintf(w, "testPostHandler")
// }

// func testHandler(w http.ResponseWriter, r *http.Request) {
// 	htmlStr := `<html >
//                 <head>
//                     <title>Collexy Installation</title>
//                 </head>
//                 <body>
//                     <div>
//                         <h1 style="text-align:center;">Collexy Logo</h1>
//                         <p>Below you should select your desired module to install.</p>
//                         <form method="POST" action="?step=2">
//                             <table>
//                                 <tr>
//                                     <td><strong>Modules</strong></td>
//                                     <td><label><input type="radio" name="collexy_module" value="TXT Starter Kit"/> TXT Starter Kit</label></td>
//                                 </tr>
//                                 <tr>
//                                     <td><input type="submit" value="Submit"></td>
//                                     <td></td>
//                                 </tr>
//                             </table>
//                             <input type="hidden" name="step" value="2"/>
//                         </form>
//                     </div>
//                 </body>
//             </html>`
// 	fmt.Fprintf(w, htmlStr)
// }

func installPostHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("./config/config.json"); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			log.Println("Config file does not exist")
			// create file
			// var bool hide_from_search_engines = false
			// if(r.PostFormValue("hide_from_search_engines")){
			//     hide_from_search_engines = r.PostFormValue("hide_from_search_engines"
			// }
			coreglobals.Conf = coreglobals.Config{r.PostFormValue("db_name"), r.PostFormValue("db_user"), r.PostFormValue("db_password"), "", r.PostFormValue("db_ssl_mode"), -1}
			res, err3 := json.Marshal(coreglobals.Conf)
			if err3 != nil {

			} else {
				// write whole the body
				absPath, _ := filepath.Abs("./config/config.json")
				err4 := ioutil.WriteFile(absPath, res, 0644)
				if err4 != nil {
					panic(err4)
				}
				installHandler(w, r)
			}
		} else {
			// other error
		}
	} else {
		// run DB create script
		var site_title, username, password, email string
		var privacy bool
		site_title = r.PostFormValue("site_title")
		username = r.PostFormValue("username")
		password = r.PostFormValue("password")
		email = r.PostFormValue("email")
		privacy = false //r.PostFormValue("privacy")
		err5 := executeDatabaseInstallScript(site_title, username, coreglobals.SetPassword(password), email, privacy)
		if err5 != nil {
			log.Println("ERROR INSTALLING DATABASE SCRIPT")
			log.Fatal(err5)
		} else {
			log.Println("DATABASE SCRIPT INSTALLED SUCCESSFULLY")
			adminHandler(w, r)
		}
	}
}

func installHandler(w http.ResponseWriter, r *http.Request) {
	//stepStr := r.URL.Query().Post("isPostBack")
	var htmlStr string
	r.ParseForm()
	step := r.PostFormValue("step")
	if step == "2" {
		fmt.Println("POST VALUE STEP = 2:::::::::::::::::")
		if _, err := os.Stat("./config/config.json"); err != nil {
			if os.IsNotExist(err) {
				// file does not exist
				log.Println("Config file does not exist")
			} else {
				// other error
			}
		} else {

			configFile, err1 := os.Open("./config/config.json")
			defer configFile.Close()
			if err1 != nil {
				log.Println("Error opening config file")
				//printError("opening config file", err1.Error())
			}

			jsonParser := json.NewDecoder(configFile)
			if err1 = jsonParser.Decode(&coreglobals.Conf); err1 != nil {
				log.Println("Error parsing config file")
				//printError("parsing config file", err1.Error())
			}
			// log.Println(coreglobals.Conf.DbName)
			// log.Println(coreglobals.Conf.DbUser)
			// log.Println(coreglobals.Conf.DbPassword)
			// log.Println(coreglobals.Conf.DbHost)
			// log.Println(coreglobals.Conf.SslMode)
			coreglobals.Db = coreglobals.SetupDB()

		}
	}

	if _, err := os.Stat("./config/config.json"); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			log.Println("Config file does not exist")
			htmlStr = `<html >
                <head>
                    <title>Collexy Installation</title>
                </head>
                <body>
                    <div>
                        <h1 style="text-align:center;">Collexy Logo</h1>
                        <p>Below you should enter your database connection details. If you're
                        not sure about theese, you should contact your host.</p>
                        <form method="POST" action="?step=2">
                            <table>
                                <tr>
                                    <td><strong>Database Name</strong></td>
                                    <td><input type="text"/ name="db_name"></td>
                                </tr>
                                <tr>
                                    <td><strong>Database username</strong></td>
                                    <td><input type="text" name="db_user"/></td>
                                </tr>
                                <tr>
                                    <td><strong>Database user password</strong></td>
                                    <td><input type="password" name="db_password"/></td>
                                </tr>
                                <tr>
                                    <td><strong>Database SSL mode</strong></td>
                                    <td><input type="text" name="db_ssl_mode"/></td>
                                </tr>
                                <tr>
                                    <td><strong>Database host (not implemented yet)</strong></td>
                                    <td><input type="text" name="db_host"/></td>
                                </tr>
                                <tr>
                                    <td><strong>Database table prefix (not implemented yet)</strong></td>
                                    <td><input type="text" name="db_table_prefix"/></td>
                                </tr>
                                <tr>
                                    <td><input type="submit" value="Submit"></td>
                                    <td></td>
                                </tr>
                            </table>
                            <input type="hidden" name="step" value="2"/>
                        </form>
                    </div>
                </body>
            </html>`
			fmt.Fprintf(w, htmlStr)
		} else {
			// other error
		}
	} else {
		log.Println(coreglobals.Conf.DbName)
		coreglobals.Db = coreglobals.SetupDB()
		if corehelpers.CheckIfDbInstalled() {
			htmlStr = `<html>
                    <head>
                        <title>Collexy Installation</title>
                    </head>
                    <body>
                        <div>
                            <h1 style="text-align:center;">Collexy Logo</h1>
                            <h2>Already Installed</h2>
                            <hr>
                            <p>You appear to have already installed Collexy. To reinstall please clear your old database tables first.</p>
                        <div>
                    </body
                </html>`
		} else {
			htmlStr = `<html>
                <head>
                    <title>Collexy Installation</title>
                </head>
                <body>
                    <div>
                        <h1 style="text-align:center;">Collexy Logo</h1>
                        <h2>Welcome</h2>
                        <hr>
                        <p>Welcome to the famous five minute Collexy installation process!
                        Just fill in the information below and you'll be on your ay to using 
                        the most extendable and powerful CMS in the world!</p>
                        <h2>Information needed</h2>
                        <hr>
                        <p>Please provide the following information. Don't worry, you can
                        always change these settings later.</p>
                        <form method="POST" action="">
                            <table>
                                <tr>
                                    <td><strong>Site title</strong></td>
                                    <td><input type="text" name="site_title"/></td>
                                </tr>
                                <tr>
                                    <td><strong>Username</strong></td>
                                    <td><input type="text" name="username"/></td>
                                </tr>
                                <tr>
                                    <td><strong>Password, twice</strong><br>
                                        <small>A password will be automatically be
                                        generated for you if you leave this blank</small>
                                    </td>
                                    <td>
                                        <input type="password" name="password"/><br>
                                        <input type="password"/>
                                    </td>
                                </tr>
                                <tr>
                                    <td><strong>Your E-mail</strong></td>
                                    <td><input type="text" name="email"/></td>
                                </tr>
                                <tr>
                                    <td><strong>Privacy</strong></td>
                                    <td>
                                        <label><input type="checkbox"/> Allow my site to appear
                                        in the search engines like Google and Technorati</label>
                                    </td>
                                </tr>
                                <tr>
                                    <td><input type="submit" value="Install Collexy"></td>
                                    <td></td>
                                </tr>
                            </table>
                        </form>
                    </div>
                </body>
            </html>`
		}

		fmt.Fprintf(w, htmlStr)
	}

}

func adminHandler(w http.ResponseWriter, r *http.Request) {

	sid := corehelpers.CheckCookie(w, r)
	u, _ := coremoduleusermodels.GetUser(sid)

	coremoduleusermodels.SetLoggedInUser(r, u)

	cc := coremodulecontentcontrollers.ContentController{}
	content := coremodulecontentmodels.Content{}
	if r.URL.String() == "/admin/login" {
		fmt.Println("FSLSO LOOOL ;;::: :: LOOL")
		//cc.RenderTemplate(w, "admin.tmpl", &content, &user)
		if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
			//cc.RenderTemplate(w, "admin.tmpl", &content, user)
			http.Redirect(w, r, "/admin", 301)
		} else {
			cc.RenderAdminTemplate(w, "admin.tmpl", &content, nil)
		}
	} else {
		if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
			cc.RenderAdminTemplate(w, "admin.tmpl", &content, user)
		} else {
			http.Redirect(w, r, "/admin/login", 301)
		}
	}
}

func GetSections(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	// for _, s := range Sections{
	//     fmt.Fprintf(w,"Section: " + s.Name + "<br>")
	// }

	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(coreglobals.Sections)
	if err != nil {
		panic(err)
	}
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)

}

func GetRoutes(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	// for _, s := range Sections{
	//     fmt.Fprintf(w,"Section: " + s.Name + "<br>")
	// }

	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(coreglobals.Routes)
	if err != nil {
		panic(err)
	}
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)

}

// func GetSectionById(w http.ResponseWriter, r *http.Request){
//     w.Header().Set("Content-Type", "application/json")
//     params := mux.Vars(r)
//     idStr := params["id"]
//     id, _ := strconv.Atoi(idStr)
// }

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)

		sid := corehelpers.CheckCookie(w, r)

		u, err := coremoduleusermodels.GetUser(sid)

		if u == nil || err == sql.ErrNoRows {
			fmt.Println(err)
			fmt.Fprintf(w, "You need to be logged in to access the API.")
		} else {
			coremoduleusermodels.SetLoggedInUser(r, u)
			h.ServeHTTP(w, r)
		}

	})
}

func APIstuff() {
	// Setup API controllers
	// nodeApiController := controllers.NodeApiController{}

	//angularRouteApiController := controllers.AngularRouteApiController{}
	sectionApiController := controllers.SectionApiController{}
	routeApiController := controllers.RouteApiController{}

	dataTypeEditorApiController := controllers.DataTypeEditorApiController{}

	fileApiController := controllers.FileApiController{}

	//menuApiController := controllers.MenuApiController{}

	//menuLinkApiController := controllers.MenuLinkApiController{}

	// Setup API routes
	privateApiRouter := coreglobals.PrivateApiRouter
	publicApiRouter := coreglobals.PublicApiRouter

	privateApiRouter.HandleFunc("/api/auth/{sid:.*}", http.HandlerFunc(corehelpers.AngularAuth)).Methods("GET")

	privateApiRouter.HandleFunc("/api/section", http.HandlerFunc(sectionApiController.Get)).Methods("GET")

	// privateApiRouter.HandleFunc("/api/section/{name:.*}", http.HandlerFunc(sectionApiController.GetByName)).Methods("GET")
	privateApiRouter.HandleFunc("/api/route", http.HandlerFunc(routeApiController.Get)).Methods("GET")
	//privateApiRouter.HandleFunc("/api/menu-link/{name:.*}", http.HandlerFunc(menuLinkApiController.GetByName)).Methods("GET")

	//publicApiRouter.HandleFunc("/api/public/contextmenutest/{nodeType:.*}", http.HandlerFunc(models.CmTest)).Methods("GET")

	// temp only
	privateApiRouter.HandleFunc("/api/data-type-editor/{alias:.*}", http.HandlerFunc(dataTypeEditorApiController.GetByAlias)).Methods("GET")
	privateApiRouter.HandleFunc("/api/data-type-editor", http.HandlerFunc(dataTypeEditorApiController.Get)).Methods("GET")

	privateApiRouter.HandleFunc("/api/file", http.HandlerFunc(fileApiController.Delete)).Methods("DELETE")

	http.Handle("/api/public/", publicApiRouter)
	http.Handle("/api/", context.ClearHandler(Middleware(privateApiRouter)))
}

func Main() {

	// temp DataTypeEditor stuff
	// should eventually be a field in each module
	dteContentPicker := lib.DataTypeEditor{"Collexy Content Picker Data Type Editor", "Collexy.DataTypeEditor.ContentPicker", "core/modules/settings/public/views/data-type/editor/content-picker.html"}
	dteRadioButtonList := lib.DataTypeEditor{"Collexy Radio Button List Data Type Editor", "Collexy.DataTypeEditor.RadioButtonList", "core/modules/settings/public/views/data-type/editor/radio-button-list.html"}

	coreglobals.DataTypeEditors = append(coreglobals.DataTypeEditors, []lib.DataTypeEditor{dteContentPicker, dteRadioButtonList}...)

	// temp end

	APIstuff()

	// var sortedModules lib.Modules

	// for _, m := range lib.Modules {

	// }
	

	sort.Sort(lib.Modules)
	m := mux.NewRouter()
	n := mux.NewRouter()

	for _, module := range lib.Modules {
		for _, s := range module.Sections {
			var r *lib.Route = s.Route
			var r1 lib.Route = *r
			coreglobals.Routes = append(coreglobals.Routes, r1)
			coreglobals.Sections = append(coreglobals.Sections, s)
			for _, t := range s.Trees {
				coreglobals.Routes = append(coreglobals.Routes, t.Routes...)
			}
			for _, cs := range s.Children {
				var r *lib.Route = cs.Route
				var r1 lib.Route = *r
				coreglobals.Routes = append(coreglobals.Routes, r1)
				// Sections = append(Sections, s)
				for _, t := range cs.Trees {
					coreglobals.Routes = append(coreglobals.Routes, t.Routes...)
				}
				for _, cs2 := range cs.Children {
					var r *lib.Route = cs2.Route
					var r1 lib.Route = *r
					coreglobals.Routes = append(coreglobals.Routes, r1)
					// Sections = append(Sections, s)
					for _, t := range cs2.Trees {
						coreglobals.Routes = append(coreglobals.Routes, t.Routes...)
					}
				}
			}

		}
		if module.ServerRoutes != nil && len(module.ServerRoutes) > 0{
			for _, sr := range module.ServerRoutes {
				m.HandleFunc(sr.Path, sr.HandlerFunc).Methods(sr.Methods...)
			}
		}
	}

	// fmt.Println("Routes:")
	// // fmt.Println(len(Routes))
	// for _, r := range coreglobals.Routes {
	// 	mystr, _ := json.Marshal(r)
	// 	fmt.Println(string(mystr))
	// }

	coreglobals.Templates["admin.tmpl"] = template.Must(template.ParseFiles("core/application/views/includes/admin.tmpl", "core/application/views/layouts/base.tmpl"))

	

	n.HandleFunc("/test/section", GetSections).Methods("GET")
	n.HandleFunc("/test/routes", GetRoutes).Methods("GET")

	contentController := coremodulecontentcontrollers.ContentController{}

	moduleController := coreapplicationcontrollers.ModuleApiController{}

	// Entity routes
	// m.Get("/api/entity/{nodeId:.*}") ?node-type=2&section=myplugin ???????????? l8r

	m.HandleFunc("/admin/install", installPostHandler).Methods("POST")
	m.HandleFunc("/admin/install", installHandler).Methods("GET")
	m.HandleFunc("/admin/test", moduleController.ModuleHandler).Methods("GET")
	m.HandleFunc("/admin/test", moduleController.ModulePostHandler).Methods("POST")
	//m.HandleFunc("/admin/{_dummy:^((?!install).)*$}", adminHandler).Methods("GET")
	//`BBB([^B]*)EEE`
	//m.HandleFunc("/admin/{([^install])*}/{*}", adminHandler).Methods("GET")
	m.HandleFunc(`/{_dummy:admin\/([^install]*).*}`, adminHandler).Methods("GET")
	//m.HandleFunc("/admin/{^((?!install).)*$}", adminHandler).Methods("GET")
	m.HandleFunc("/admin", adminHandler).Methods("GET")

	// or use "/url:.*" for all
	
	m.HandleFunc("/{url:.*}", http.HandlerFunc(contentController.RenderContent)).Methods("GET")

	http.Handle("/stylesheets/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.Handle("/scripts/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.Handle("/media/", http.StripPrefix("/", MediaProtectHandler(http.FileServer(http.Dir("./")))))
	http.Handle("/assets/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.Handle("/public/", http.FileServer(http.Dir("./core/application")))

	log.Println("Registered a handler for static files.")

	http.Handle("/test/", n)
	http.Handle("/", m)
}


func buildMap(mySlice ...*coreglobals.MediaAccessItem) (myMap map[string]*coreglobals.MediaAccessItem) {
	myMap = make(map[string]*coreglobals.MediaAccessItem)
	for _, item := range mySlice {

		myMap[item.Url] = item
		itemKeyIdStr := strconv.Itoa(item.MediaId)
		myMap[itemKeyIdStr] = item
	}
	return
}

func MediaProtectHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r.URL)
		var urlStr string = ""

		if _, err := os.Stat("./config/media-access.xml"); err != nil {
			if os.IsNotExist(err) {
				// file does not exist
				log.Println("media-access.xml XML file does not exist")
			} else {
				// other error
			}
		} else {

			configFile, err1 := os.Open("./config/media-access.xml")
			defer configFile.Close()
			if err1 != nil {
				log.Println("Error opening media-access.xml config file")
				//printError("opening config file", err1.Error())
			}

			XMLdata, err2 := ioutil.ReadAll(configFile)

			fmt.Println(string(XMLdata))

			if err2 != nil {
				log.Println("Error reading from media-access.xml config file")
				fmt.Printf("error: %v", err2)
			}

			var v coreglobals.MediaAccessItems
			err := xml.Unmarshal(XMLdata, &v)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			//fmt.Printf("%#v\n", v)

			coreglobals.MediaAccessConf = buildMap(v.Items...)

			fmt.Println("test MEdia acceSS")
			fmt.Println(v)

			// jsonParser := json.NewDecoder(configFile)
			// if err1 = jsonParser.Decode(&coreglobals.MediaAccessConf); err1 != nil {
			// 	log.Println("Error parsing media-access.xml config file")
			// 	log.Println(err1.Error())
			// 	//printError("parsing config file", err1.Error())
			// }

			urlStr = r.URL.Path
			// urlStrTemp, err2 := r.URL.Path()
			// if err2!=nil{
			// 	log.Println(err2.Error())

			// } else{
			// 	urlStrTemp = urlStrTemp
			// }

			log.Println("urlStr: " + urlStr)
			log.Println(coreglobals.MediaAccessConf[urlStr])
			//log.Println(url.QueryEscape(urlStr))
			log.Println("111!!!!!1")
			// log.Println(coreglobals.Maccess.Items[0].Domains[0])
			// log.Println(coreglobals.Maccess.Items[0].Url)
			// fmt.Println(coreglobals.Maccess.Items[0].MemberGroups)
		}

		// fmt.Println(r.URL.Path)
		// fmt.Println(coreglobals.Maccess.Domains[0] + "/" + coreglobals.Maccess.Url)
		// fmt.Println(r.Host)

		isProtected := false
		hasAccess := false
		var protectedItem *coreglobals.MediaAccessItem

		if val, ok := coreglobals.MediaAccessConf[urlStr]; ok {
			isProtected = true
			protectedItem = val
		}

		// fmt.Println(protectedItem)
		// fmt.Println(protectedItem.MemberGroups)

		// var protectedItem *coreglobals.MediaAccessItem = nil

		// for _, maItem := range coreglobals.Maccess.Items {
		// 	if isProtected {
		// 		break;
		// 	}
		// 	for _, domain := range maItem.Domains {
		// 		if isProtected {
		// 			break;
		// 		}
		// 		if domain + "/" + maItem.Url == r.Host + "/" + r.URL.Path{
		// 			if isProtected {
		// 				break;
		// 			}
		// 			isProtected = true;
		// 			protectedItem = &maItem
		// 			// fmt.Fprintf(w, "loldalolselol")

		// 		}
		// 	}
		// }
		if isProtected {
			sid := corehelpers.CheckMemberCookie(w, r)

			m, err := coremodulemembermodels.GetMember(sid)

			if m == nil || err == sql.ErrNoRows {
				fmt.Println(err)
				// hasAccess = false //already set when var was initialized

			} else {

				coremodulemembermodels.SetLoggedInMember(r, m)

				for _, mg := range m.Groups {
					if hasAccess {
						break
					}
					// fmt.Println("MEMBER :::::: ")
					// fmt.Println(mg)
					// fmt.Println(protectedItem.MemberGroups[0])
					if mg.Id == protectedItem.MemberGroups[0] {
						// fmt.Println("workz?")
						hasAccess = true
					}
				}

			}
			if !hasAccess {
				fmt.Fprintf(w, "You need to be logged in to access this media item.")
			} else {
				h.ServeHTTP(w, r)
			}

		} else {
			h.ServeHTTP(w, r)
		}

		//sid := corehelpers.CheckCookie(w, r)

		//m, err := coremodulemembermodels.GetMember(sid)

		// if m == nil || err == sql.ErrNoRows {
		// 	fmt.Println(err)
		// 	fmt.Fprintf(w, "You need to be logged in to access the API.")

		// } else {
		// 	coremodulemembermodels.SetLoggedInMember(r, m)

		// 	h.ServeHTTP(w, r)
		// }
	})
}

/* ALL THE STRUCTS BELOW SHOULD NOT BE NECESSARY
* A custom UnmarshalXML function should be implemented instead!!
* For testing purposes only!
 */

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

// func MediaProtectHandler(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(*r.URL)

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
// 			if err1 = jsonParser.Decode(&coreglobals.Maccess); err1 != nil {
// 				log.Println("Error parsing media-access.json config file")
// 				//printError("parsing config file", err1.Error())
// 			}
// 			log.Println(coreglobals.Maccess.Items[0].Domains[0])
// 			log.Println(coreglobals.Maccess.Items[0].Url)
// 			fmt.Println(coreglobals.Maccess.Items[0].MemberGroups)
// 		}

// 		// fmt.Println(r.URL.Path)
// 		// fmt.Println(coreglobals.Maccess.Domains[0] + "/" + coreglobals.Maccess.Url)
// 		// fmt.Println(r.Host)

// 		isProtected := false;
// 		hasAccess := false;
// 		var protectedItem *coreglobals.MediaAccessItem = nil

// 		for _, maItem := range coreglobals.Maccess.Items {
// 			if isProtected {
// 				break;
// 			}
// 			for _, domain := range maItem.Domains {
// 				if isProtected {
// 					break;
// 				}
// 				if domain + "/" + maItem.Url == r.Host + "/" + r.URL.Path{
// 					if isProtected {
// 						break;
// 					}
// 					isProtected = true;
// 					protectedItem = &maItem
// 					// fmt.Fprintf(w, "loldalolselol")

// 				}
// 			}
// 		}
// 		if isProtected{
// 			sid := corehelpers.CheckMemberCookie(w, r)

// 			m, err := coremodulemembermodels.GetMember(sid)

// 			if m == nil || err == sql.ErrNoRows {
// 				fmt.Println(err)
// 				// hasAccess = false //already set when var was initialized

// 			} else {
// 				coremodulemembermodels.SetLoggedInMember(r, m)

// 				for _, mg := range m.Groups {
// 					if hasAccess {
// 						break;
// 					}
// 					fmt.Println("MEMBER :::::: ")
// 					fmt.Println(mg)
// 					fmt.Println(protectedItem.MemberGroups[0])
// 					if mg.Id == protectedItem.MemberGroups[0]{
// 						fmt.Println("workz?")
// 						hasAccess = true;
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
