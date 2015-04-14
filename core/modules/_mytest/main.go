package mytest

/*
* The purpose of a module is to let you extend and customize Collexy to to your needs, without touching the core files
* You can add custom sections, trees, routes, views, controllers, services, directives, assets, custom db tables etc.
*
* Todo:
* - Context Menu per tree
*/

import
(
	"log"
	"net/http"
	"collexy/core/lib"
)

func init(){
	// setup routes
	// params: state, url, templateUrl, isAbstract

	rTestSection := lib.Route{"admin.test", "/admin/test", "public/views/test/index.html", false}
	
	rTestTreeMethodEdit := lib.Route{"admin.test.edit", "/admin/test/edit/:id", "public/views/test/edit.html", false}
	rTestTreeMethodNew := lib.Route{"admin.test.new", "/admin/test/new", "public/views/test/new.html", false}

	// Setup actions
	// params: function+params, isDialog
	// aTestTree := lib.Action{"delete(item)", true}

	// setup methods
	// mTestEdit := lib.Method{"Edit", "edit", rTestTreeMethodEdit}

	// setup tree
	// params: name, alias, routes
	routesTestTree := []lib.Route{rTestTreeMethodEdit, rTestTreeMethodNew}
	
	tTest := lib.Tree{"Test Tree", "testTree", routesTestTree}

	// setup section
	treesTestSection := []*lib.Tree{&tTest}
	// params: name, alias, icon, route, trees, iscontainer, parent
	sTest := lib.Section{"Test Section", "testSection", "fa fa-newspaper-o fa-fw", &rTestSection, treesTestSection, false, nil, nil}

	// setup module
	sections := []lib.Section{sTest}
	// params: name, alias, description, sections
	moduleTest := lib.Module{"Test Module", "testModule", "Just a test module", sections}

	// register module
	lib.RegisterModule(moduleTest)
	// log.Println(lib.Modules)

	// FileServer
	log.Println("Registered a handler for static files. [mytest::module]")
	http.Handle("/core/modules/mytest/public/", http.FileServer(http.Dir("./")))
}