package settings

import (
	"collexy/core/lib"
	"log"
	"net/http"
	// "reflect"
	//"encoding/json"
	coreglobals "collexy/core/globals"
	coremodulesettingscontrollers "collexy/core/modules/settings/controllers"
)

func init() {

	// register API route endpoints, controllers and handlers

	contentTypeApiController := coremodulesettingscontrollers.ContentTypeApiController{}
	mediaTypeApiController := coremodulesettingscontrollers.MediaTypeApiController{}
	dataTypeApiController := coremodulesettingscontrollers.DataTypeApiController{}
	templateApiController := coremodulesettingscontrollers.TemplateApiController{}
	directoryApiController := coremodulesettingscontrollers.DirectoryApiController{}

	contentTypeTreeController := coremodulesettingscontrollers.ContentTypeTreeController{}
	mediaTypeTreeController := coremodulesettingscontrollers.MediaTypeTreeController{}
	dataTypeTreeController := coremodulesettingscontrollers.DataTypeTreeController{}
	templateTreeController := coremodulesettingscontrollers.TemplateTreeController{}
	directoryTreeController := coremodulesettingscontrollers.DirectoryTreeController{}

	privateApiRouter := coreglobals.PrivateApiRouter
	subrPrivate := privateApiRouter.PathPrefix("/").Subrouter()

	// m.Get("/api/content-type/extended/{nodeId:.*}", http.HandlerFunc(contentTypeApiController.GetContentTypeExtendedByNodeId))
	subrPrivate.HandleFunc("/api/content-type/{id:.*}/contextmenu", http.HandlerFunc(contentTypeTreeController.GetMenuForContentType)).Methods("GET")
	subrPrivate.HandleFunc("/api/content-type/{id:.*}/children", http.HandlerFunc(contentTypeApiController.GetByIdChildren)).Methods("GET")
	subrPrivate.HandleFunc("/api/content-type/{id:.*}", http.HandlerFunc(contentTypeApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/content-type", http.HandlerFunc(contentTypeApiController.Get)).Methods("GET")

	subrPrivate.HandleFunc("/api/media-type/{id:.*}/contextmenu", http.HandlerFunc(mediaTypeTreeController.GetMenuForMediaType)).Methods("GET")
	subrPrivate.HandleFunc("/api/media-type/{id:.*}/children", http.HandlerFunc(mediaTypeApiController.GetByIdChildren)).Methods("GET")
	subrPrivate.HandleFunc("/api/media-type/{id:.*}", http.HandlerFunc(mediaTypeApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/media-type", http.HandlerFunc(mediaTypeApiController.Get)).Methods("GET")
	//m.Get("/api/content-type/", http.HandlerFunc(contentTypeApiController.GetContentTypeByNodeId)) // not sure about this
	//m.Get("/api/content-type/", http.HandlerFunc(contentTypeApiController.GetContentTypes)) // not sure about this
	//privateApiRouter.HandleFunc("/api/content-type/{nodeId:.*}", http.HandlerFunc(contentTypeApiController.PutContentType)).Methods("PUT")
	//privateApiRouter.HandleFunc("/api/content-type/{nodeId:.*}", http.HandlerFunc(contentTypeApiController.Post)).Methods("POST")

	subrPrivate.HandleFunc("/api/data-type/{id:.*}/contextmenu", http.HandlerFunc(dataTypeTreeController.GetMenuForDataType)).Methods("GET")
	subrPrivate.HandleFunc("/api/data-type/{id:.*}", http.HandlerFunc(dataTypeApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/data-type", http.HandlerFunc(dataTypeApiController.Get)).Methods("GET") // not sure about this
	subrPrivate.HandleFunc("/api/data-type/{id:.*}", http.HandlerFunc(dataTypeApiController.Put)).Methods("PUT")
	subrPrivate.HandleFunc("/api/data-type/{id:.*}", http.HandlerFunc(dataTypeApiController.Post)).Methods("POST")
	subrPrivate.HandleFunc("/api/data-type/{id:.*}", http.HandlerFunc(dataTypeApiController.Delete)).Methods("DELETE")

	subrPrivate.HandleFunc("/api/template/{id:.*}/contextmenu", http.HandlerFunc(templateTreeController.GetMenuForTemplate)).Methods("GET")
	subrPrivate.HandleFunc("/api/template/{id:.*}/children", http.HandlerFunc(templateApiController.GetByIdChildren)).Methods("GET")
	subrPrivate.HandleFunc("/api/template/{id:.*}", http.HandlerFunc(templateApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/template", http.HandlerFunc(templateApiController.Get)).Methods("GET") // not sure about this
	subrPrivate.HandleFunc("/api/template/{id:.*}", http.HandlerFunc(templateApiController.Put)).Methods("PUT")
	subrPrivate.HandleFunc("/api/template/{id:.*}", http.HandlerFunc(templateApiController.Post)).Methods("POST")
	subrPrivate.HandleFunc("/api/template/{id:.*}", http.HandlerFunc(templateApiController.Delete)).Methods("DELETE")

	// Directory
	subrPrivate.HandleFunc("/api/directory/{rootdir:.*}/{name:.*}/{is_dir:.*}/contextmenu", http.HandlerFunc(directoryTreeController.GetMenuForDirectory)).Methods("GET")
	subrPrivate.HandleFunc("/api/directory/upload-file-test", http.HandlerFunc(directoryApiController.UploadFileTest)).Methods("POST")
	subrPrivate.HandleFunc("/api/directory/{rootdir:.*}/{name:.*}", http.HandlerFunc(directoryApiController.Post)).Methods("POST")
	subrPrivate.HandleFunc("/api/directory/{rootdir:.*}/{name:.*}", http.HandlerFunc(directoryApiController.Put)).Methods("PUT")

	subrPrivate.HandleFunc("/api/directory/{rootdir:.*}/{name:.*}", http.HandlerFunc(directoryApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/directory/{rootdir:.*}", http.HandlerFunc(directoryApiController.Get)).Methods("GET")

	// API END

	// setup routes
	rSettingsSection := lib.Route{"settings", "/admin/settings", "core/modules/settings/public/views/settings/index.html", true}

	rContentTypeSection := lib.Route{"settings.contentType", "/content-type", "core/modules/settings/public/views/content-type/index.html", false}
	rContentTypeTreeMethodEdit := lib.Route{"settings.contentType.edit", "/edit/:id", "core/modules/settings/public/views/content-type/edit.html", false}
	rContentTypeTreeMethodNew := lib.Route{"settings.contentType.new", "/new?type_id&parent_id", "core/modules/settings/public/views/content-type/new.html", false}

	rMediaTypeSection := lib.Route{"settings.mediaType", "/media-type", "core/modules/settings/public/views/media-type/index.html", false}
	rMediaTypeTreeMethodEdit := lib.Route{"settings.mediaType.edit", "/edit/:id", "core/modules/settings/public/views/media-type/edit.html", false}
	rMediaTypeTreeMethodNew := lib.Route{"settings.mediaType.new", "/new?type_id&parent_id", "core/modules/settings/public/views/media-type/new.html", false}

	rDataTypeSection := lib.Route{"settings.dataType", "/data-type", "core/modules/settings/public/views/data-type/index.html", false}
	rDataTypeTreeMethodEdit := lib.Route{"settings.dataType.edit", "/edit/:id", "core/modules/settings/public/views/data-type/edit.html", false}
	rDataTypeTreeMethodNew := lib.Route{"settings.dataType.new", "/new", "core/modules/settings/public/views/data-type/new.html", false}

	rTemplateSection := lib.Route{"settings.template", "/template", "core/modules/settings/public/views/template/index.html", false}
	rTemplateTreeMethodEdit := lib.Route{"settings.template.edit", "/edit/:id", "core/modules/settings/public/views/template/edit.html", false}
	rTemplateTreeMethodNew := lib.Route{"settings.template.new", "/new?parent_id", "core/modules/settings/public/views/template/new.html", false}

	rScriptSection := lib.Route{"settings.script", "/script", "core/modules/settings/public/views/script/index.html", false}
	rScriptTreeMethodEdit := lib.Route{"settings.script.edit", "/edit/:name", "core/modules/settings/public/views/script/edit.html", false}
	rScriptTreeMethodNew := lib.Route{"settings.script.new", "/new?type&parent", "core/modules/settings/public/views/script/new.html", false}

	rStylesheetSection := lib.Route{"settings.stylesheet", "/stylesheet", "core/modules/settings/public/views/stylesheet/index.html", false}
	rStylesheetTreeMethodEdit := lib.Route{"settings.stylesheet.edit", "/edit/:name", "core/modules/settings/public/views/stylesheet/edit.html", false}
	rStylesheetTreeMethodNew := lib.Route{"settings.stylesheet.new", "/new?type&parent", "core/modules/settings/public/views/stylesheet/new.html", false}

	// setup trees
	routesContentTypeTree := []lib.Route{rContentTypeTreeMethodEdit, rContentTypeTreeMethodNew}
	routesMediaTypeTree := []lib.Route{rMediaTypeTreeMethodEdit, rMediaTypeTreeMethodNew}
	routesDataTypeTree := []lib.Route{rDataTypeTreeMethodEdit, rDataTypeTreeMethodNew}
	routesTemplateTree := []lib.Route{rTemplateTreeMethodEdit, rTemplateTreeMethodNew}
	routesScriptTree := []lib.Route{rScriptTreeMethodEdit, rScriptTreeMethodNew}
	routesStylesheetTree := []lib.Route{rStylesheetTreeMethodEdit, rStylesheetTreeMethodNew}

	tContentType := lib.Tree{"Content Types", "contentTypes", routesContentTypeTree}
	tMediaType := lib.Tree{"Media Types", "mediaTypes", routesMediaTypeTree}
	tDataType := lib.Tree{"DataTypes", "dataTypes", routesDataTypeTree}
	tTemplate := lib.Tree{"Templates", "templates", routesTemplateTree}
	tScript := lib.Tree{"Scripts", "scripts", routesScriptTree}
	tStylesheet := lib.Tree{"Stylesheets", "stylesheets", routesStylesheetTree}

	treesContentTypeSection := []*lib.Tree{&tContentType}
	treesMediaTypeSection := []*lib.Tree{&tMediaType}
	treesDataTypeSection := []*lib.Tree{&tDataType}
	treesTemplateSection := []*lib.Tree{&tTemplate}
	treesScriptSection := []*lib.Tree{&tScript}
	treesStylesheetSection := []*lib.Tree{&tStylesheet}

	// params: name, alias, icon, route, trees, iscontainer, parent
	sSettings := lib.Section{"Settings Section", "settingsSection", "fa fa-gear fa-fw", &rSettingsSection, nil, true, nil, nil, []string{"settings_section"}}

	sContentType := lib.Section{"Content Type Section", "contentTypeSection", "fa fa-newspaper-o fa-fw", &rContentTypeSection, treesContentTypeSection, false, nil, nil, []string{"content_type_section"}}
	sMediaType := lib.Section{"Media Type Section", "mediaTypeSection", "fa fa-files-o fa-fw", &rMediaTypeSection, treesMediaTypeSection, false, nil, nil, []string{"media_type_section"}}
	sDataType := lib.Section{"Data Type Section", "dataTypeSection", "fa fa-check-square-o fa-fw", &rDataTypeSection, treesDataTypeSection, false, nil, nil, []string{"data_type_section"}}
	sTemplate := lib.Section{"Template Section", "templateSection", "fa fa-eye fa-fw", &rTemplateSection, treesTemplateSection, false, nil, nil, []string{"template_section"}}
	sScript := lib.Section{"Script Section", "scriptSection", "fa fa-file-code-o fa-fw", &rScriptSection, treesScriptSection, false, nil, nil, []string{"script_section"}}
	sStylesheet := lib.Section{"Stylesheet Type Section", "stylesheetSection", "fa fa-desktop fa-fw", &rStylesheetSection, treesStylesheetSection, false, nil, nil, []string{"stylesheet_section"}}

	lol := []lib.Section{sContentType, sMediaType, sDataType, sTemplate, sScript, sStylesheet}
	sSettings.Children = lol
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
	sections := []lib.Section{sSettings}
	// params: name, alias, description, sections
	moduleSettings := lib.Module{"Settings Module", "settingsModule", "Just a settings module", sections, 500}

	// register module
	lib.RegisterModule(moduleSettings)

	// Setup FileServer for the settings module
	log.Println("Registered a handler for static files. [settings::module]")
	http.Handle("/core/modules/settings/public/", http.FileServer(http.Dir("./")))
}

// func main(){
// 	log.Println("MAIN ------------------------------------------------------------------------")
// }
