package user

import (
	coreglobals "collexy/core/globals"
	"collexy/core/lib"
	coremoduleusercontrollers "collexy/core/modules/user/controllers"
	"log"
	"net/http"
)

func init() {

	privateApiRouter := coreglobals.PrivateApiRouter
	publicApiRouter := coreglobals.PublicApiRouter
	subrPrivate := privateApiRouter.PathPrefix("/").Subrouter()
	subrPublic := publicApiRouter.PathPrefix("/").Subrouter()

	userApiController := coremoduleusercontrollers.UserApiController{}
	userGroupApiController := coremoduleusercontrollers.UserGroupApiController{}

	permissionApiController := coremoduleusercontrollers.PermissionApiController{}

	userTreeController := coremoduleusercontrollers.UserTreeController{}
	userGroupTreeController := coremoduleusercontrollers.UserGroupTreeController{}
	permissionTreeController := coremoduleusercontrollers.PermissionTreeController{}

	// User
	subrPublic.HandleFunc("/api/public/user/login", http.HandlerFunc(userApiController.Login)).Methods("POST")

	subrPrivate.HandleFunc("/api/user/{id:.*}/contextmenu", http.HandlerFunc(userTreeController.GetMenuForUser)).Methods("GET")
	subrPrivate.HandleFunc("/api/user/{id:.*}", http.HandlerFunc(userApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/user", http.HandlerFunc(userApiController.Get)).Methods("GET")
	subrPrivate.HandleFunc("/api/user", http.HandlerFunc(userApiController.Post)).Methods("POST")
	subrPrivate.HandleFunc("/api/user/{id:.*}", http.HandlerFunc(userApiController.Put)).Methods("PUT")
	subrPrivate.HandleFunc("/api/user/{id:.*}", http.HandlerFunc(userApiController.Delete)).Methods("DELETE")

	subrPrivate.HandleFunc("/api/user-group/{id:.*}/contextmenu", http.HandlerFunc(userGroupTreeController.GetMenuForUserGroup)).Methods("GET")
	subrPrivate.HandleFunc("/api/user-group", http.HandlerFunc(userGroupApiController.Get)).Methods("GET")
	subrPrivate.HandleFunc("/api/user-group/{id:.*}", http.HandlerFunc(userGroupApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/user-group", http.HandlerFunc(userGroupApiController.Post)).Methods("POST")
	subrPrivate.HandleFunc("/api/user-group/{id:.*}", http.HandlerFunc(userGroupApiController.Put)).Methods("PUT")
	subrPrivate.HandleFunc("/api/user-group/{id:.*}", http.HandlerFunc(userGroupApiController.Delete)).Methods("DELETE")

	subrPrivate.HandleFunc("/api/permission/{id:.*}/contextmenu", http.HandlerFunc(permissionTreeController.GetMenuForPermission)).Methods("GET")
	subrPrivate.HandleFunc("/api/permission/{id:.*}", http.HandlerFunc(permissionApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/permission", http.HandlerFunc(permissionApiController.Get)).Methods("GET")
	subrPrivate.HandleFunc("/api/permission/{id:.*}", http.HandlerFunc(permissionApiController.Post)).Methods("POST")
	subrPrivate.HandleFunc("/api/permission/{id:.*}", http.HandlerFunc(permissionApiController.Put)).Methods("PUT")
	subrPrivate.HandleFunc("/api/permission/{id:.*}", http.HandlerFunc(permissionApiController.Delete)).Methods("DELETE")

	///

	// setup routes
	rUserSection := lib.Route{"user", "/admin/user", "core/modules/user/public/views/user/index.html", false}
	rUserTreeMethodEdit := lib.Route{"user.edit", "/edit/:id", "core/modules/user/public/views/user/edit.html", false}
	rUserTreeMethodNew := lib.Route{"user.new", "/new", "core/modules/user/public/views/user/new.html", false}

	rUserGroupSection := lib.Route{"user.userGroup", "/user-group", "core/modules/user/public/views/user-group/index.html", false}
	rUserGroupTreeMethodEdit := lib.Route{"user.userGroup.edit", "/edit/:id", "core/modules/user/public/views/user-group/edit.html", false}
	rUserGroupTreeMethodNew := lib.Route{"user.userGroup.new", "/new", "core/modules/user/public/views/user-group/new.html", false}

	rPermissionSection := lib.Route{"user.permission", "/permission", "core/modules/user/public/views/permission/index.html", false}
	rPermissionTreeMethodEdit := lib.Route{"user.permission.edit", "/edit/:id", "core/modules/user/public/views/permission/edit.html", false}
	rPermissionTreeMethodNew := lib.Route{"user.permission.new", "/new", "core/modules/user/public/views/permission/new.html", false}

	// setup trees
	routesUserTree := []lib.Route{rUserTreeMethodEdit, rUserTreeMethodNew}
	routesUserGroupTree := []lib.Route{rUserGroupTreeMethodEdit, rUserGroupTreeMethodNew}
	routesPermissionTree := []lib.Route{rPermissionTreeMethodEdit, rPermissionTreeMethodNew}

	tUser := lib.Tree{"Users", "users", routesUserTree}
	tUserGroup := lib.Tree{"User Groups", "userGroups", routesUserGroupTree}
	tPermission := lib.Tree{"Permissions", "permissions", routesPermissionTree}

	treesUserSection := []*lib.Tree{&tUser}
	treesUserGroupSection := []*lib.Tree{&tUserGroup}
	treesPermissionSection := []*lib.Tree{&tPermission}

	// params: name, alias, icon, route, trees, iscontainer, parent
	sUsers := lib.Section{"Users Section", "usersSection", "fa fa-user fa-fw", &rUserSection, treesUserSection, false, nil, nil, []string{"user_section"}}

	sUserGroup := lib.Section{"User Group Section", "userGroupSection", "fa fa-smile-o fa-fw", &rUserGroupSection, treesUserGroupSection, false, nil, nil, []string{"user_group_section"}}

	sPermission := lib.Section{"Permission Section", "permissionSection", "fa fa-smile-o fa-fw", &rPermissionSection, treesPermissionSection, false, nil, nil, []string{"permission_section"}}

	lol := []lib.Section{sUserGroup, sPermission}
	sUsers.Children = lol
	//reflect.ValueOf(&sUsers).Elem().FieldByName("Children").Set(reflect.ValueOf(lol))
	// log.Println(sUsers.Children[0].Name + ":FDSF:SDF:DS:F:")
	// res, err := json.Marshal(sUsers)
	// if err!=nil{
	// 	panic(err)
	// }
	// log.Println(res)
	// log.Println("__-----------------")
	//sUsers.SetChildren(lol)
	//sUsers.Children =

	// setup users section
	// maybe add IsContainer bool?
	// section parent by name or section children - or both? For subsections ofc

	// setup subsections

	// setup module
	sections := []lib.Section{sUsers}
	// params: name, alias, description, sections
	moduleUsers := lib.Module{"Users Module", "usersModule", "Just a users module", sections, nil, 400}

	// register module
	lib.RegisterModule(moduleUsers)

	// FileServer
	log.Println("Registered a handler for static files. [user::module]")
	http.Handle("/core/modules/user/public/", http.FileServer(http.Dir("./")))
}
