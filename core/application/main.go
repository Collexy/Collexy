package application

import(
    "log"
    "net/http"
    "html/template"
    "collexy/core/application/controllers"
    applicationglobals "collexy/core/application/globals"
    "collexy/core/api/models"
    corehelpers "collexy/core/helpers"
    "github.com/gorilla/mux"
    //"collexy/globals"
    "fmt"
    "os"
    coreglobals "collexy/core/globals"
    "encoding/json"
    "io/ioutil"
    "path/filepath"
    //"collexy/helpers"
)

func executeDatabaseInstallScript(site_title,username,password,email string,privacy bool) (err error) {
    db := coreglobals.Db
    _, err1 := db.Exec(coreglobals.DbCreateScriptDML + coreglobals.DbCreateScriptDDL, site_title, username, password, email)
    err = err1
    return
}

func installPostHandler(w http.ResponseWriter, r *http.Request){
    if _, err := os.Stat("./../../config/config.json"); err != nil {
        if os.IsNotExist(err) {
            // file does not exist
            log.Println("Config file does not exist")
            // create file
            coreglobals.Conf = coreglobals.Config{r.PostFormValue("db_name"), r.PostFormValue("db_user"), r.PostFormValue("db_password"), "", r.PostFormValue("db_ssl_mode")}
            res, err3 := json.Marshal(coreglobals.Conf)
            if(err3 != nil){

            } else {
                // write whole the body
                absPath, _ := filepath.Abs("./config/config.json")
                err4 := ioutil.WriteFile(absPath, res, 0644)
                if err4 != nil {
                    panic(err4)
                }
                installHandler(w,r)
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
        err5 := executeDatabaseInstallScript(site_title,username,password,email,privacy)
        if(err5 != nil){
            log.Println("ERROR INSTALLING DATABASE SCRIPT")
        } else{
            log.Println("DATABASE SCRIPT INSTALLED SUCCESSFULLY")
        }
    }
}

func installHandler(w http.ResponseWriter, r *http.Request){
    //stepStr := r.URL.Query().Post("isPostBack")
    r.ParseForm()
    step := r.PostFormValue("step")
    if(step == "2"){
        fmt.Println("POST VALUE STEP = 2:::::::::::::::::")
        if _, err := os.Stat("./../../config/config.json"); err != nil {
            if os.IsNotExist(err) {
                // file does not exist
                log.Println("Config file does not exist")
            } else {
                // other error
            }
        } else {

            configFile, err1 := os.Open("./../..//config/config.json")
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

    coreglobals.Db = coreglobals.SetupDB()
    if(corehelpers.CheckIfDbInstalled()){
        htmlStr := `<html>
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
        fmt.Fprintf(w, htmlStr)
    } else {
        if _, err := os.Stat("./config/config.json"); err != nil {
            if os.IsNotExist(err) {
                // file does not exist
                log.Println("Config file does not exist")
                htmlStr := `<html >
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
                                        <td><input type="text" name="db_password"/></td>
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
            htmlStr := `<html>
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
                                        <input type="password">
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
            fmt.Fprintf(w, htmlStr)
        }
    }
    
    
}

func adminHandler(w http.ResponseWriter, r *http.Request) {

    sid := corehelpers.CheckCookie(w,r)
    u, _ := models.GetUser(sid)

    models.SetLoggedInUser(r,u)


    cc := controllers.ContentController{}
    content := models.Content{}
    if(r.URL.String() == "/admin/login"){
        //cc.RenderTemplate(w, "admin.tmpl", &content, &user)
        if user := models.GetLoggedInUser(r); user != nil {
            //cc.RenderTemplate(w, "admin.tmpl", &content, user)
            http.Redirect(w, r, "/admin", 301)
        } else {
            cc.RenderAdminTemplate(w, "admin.tmpl", &content, nil)
        }
    } else {
        if user := models.GetLoggedInUser(r); user != nil {
            cc.RenderAdminTemplate(w, "admin.tmpl", &content, user)
        } else {
            http.Redirect(w, r, "/admin/login", 301)
        }
    }
}


// func init is run before main. All packages can have an init function

func init(){
    test()
}

func test() {
    // var routes []globals.IRoute
    // var menus []interface{}
    /**
    * Menu
    */
    am := models.AdminMenu{"main", nil}

    
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
    rDashboard := models.AdminRoute{"index", nil, "/admin", components, "", nil}

    // Menu items
    miDashboard := models.AdminMenuItem {"Dashboard", "fa fa-dashboard fa-fw", &rDashboard, nil}
    am.AddItem(miDashboard)

    /**
    * Settings
    */

    // Sub menus
    smSettings := models.AdminMenu{"Settings", nil}

    // Routes
    components = nil
    component = map[string]interface{}{
        "single":"public/views/settings/index.html",
    }
    components = append(components, component)
    //components = [0]map[string]interface{}{"single":"public/views/settings/index.html"}
    rSettings := models.AdminRoute{"settings", nil, "/admin/settings", components, "", nil}

    //
    components = nil
    component = map[string]interface{}{
        "single":"public/views/settings/content-type/index.html",
    }
    components = append(components, component)
    //components = [0]map[string]interface{}{"single":"public/views/settings/content-type/index.html"}
    rSettingsContentTypes := models.AdminRoute{"contentType", nil, "/content-type", components, "", nil}

    components = nil
    component = map[string]interface{}{
        "single":"public/views/settings/content-type/new.html",
    }
    components = append(components, component)
    //components = [0]map[string]interface{}{"single":"public/views/settings/content-type/new.html"}
    rSettingsContentTypesNew := models.AdminRoute{"new", nil, "/new?type&parent", components, "", nil}

    components = nil
    component = map[string]interface{}{
        "single":"public/views/settings/content-type/edit.html",
    }
    components = append(components, component)
    //components = [0]map[string]interface{}{"single":"public/views/settings/content-type/edit.html"}
    rSettingsContentTypesEdit := models.AdminRoute{"edit", nil, "/edit/:nodeId", components, "", nil}


    // Menu items

    

    miSettingsContentTypes := models.AdminMenuItem {"Content Types", "fa fa-newspaper-o fa-fw", &rSettingsContentTypes, nil}

    smSettings.AddItem(miSettingsContentTypes)

    miSettings := models.AdminMenuItem {"Settings", "fa fa-gear fa-fw", &rSettings, &smSettings}

    am.AddItem(miSettings)

    coreglobals.Menus = append(coreglobals.Menus, am)

    // rSettingsContentTypes.Children = append(rSettingsContentTypes.Children,rSettingsContentTypesNew)
    // rSettingsContentTypes.Children = append(rSettingsContentTypes.Children,rSettingsContentTypesEdit)
    // rSettings.Children = append(rSettings.Children,rSettingsContentTypes)

    rSettingsContentTypes.AddChildren(&rSettingsContentTypesNew)
    rSettingsContentTypes.AddChildren(&rSettingsContentTypesEdit)
    rSettings.AddChildren(&rSettingsContentTypes)

    coreglobals.Routes = append(coreglobals.Routes, &rDashboard)
    coreglobals.Routes = append(coreglobals.Routes, &rSettings)

    //globals.Routes = routes
}

func Main(){
    applicationglobals.Templates["admin.tmpl"] = template.Must(template.ParseFiles("core/application/views/includes/admin.tmpl", "core/application/views/layouts/base.tmpl"))

    m := mux.NewRouter()

    contentController := controllers.ContentController{}

    // Entity routes
    // m.Get("/api/entity/{nodeId:.*}") ?node-type=2&section=myplugin ???????????? l8r

    m.HandleFunc("/admin", adminHandler).Methods("GET")
    m.HandleFunc("/admin/install", installPostHandler).Methods("POST")
    m.HandleFunc("/admin/install", installHandler).Methods("GET")
    m.HandleFunc("/admin/{_dummy:.*}", adminHandler).Methods("GET")

    m.HandleFunc("/{url:.*}", http.HandlerFunc(contentController.RenderContent)).Methods("GET")


    http.Handle("/media/",  http.StripPrefix("/", http.FileServer(http.Dir("./"))))
    http.Handle("/public/", http.FileServer(http.Dir("./core/application"))) 

    log.Println("Registered a handler for static files.")
    
    http.Handle("/", m)
}