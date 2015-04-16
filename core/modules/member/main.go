package member

import (
	coreglobals "collexy/core/globals"
	"collexy/core/lib"
	coremodulemembercontrollers "collexy/core/modules/member/controllers"
	"log"
	"net/http"
)

func init() {

	privateApiRouter := coreglobals.PrivateApiRouter
	publicApiRouter := coreglobals.PublicApiRouter
	subrPrivate := privateApiRouter.PathPrefix("/").Subrouter()
	subrPublic := publicApiRouter.PathPrefix("/").Subrouter()

	memberApiController := coremodulemembercontrollers.MemberApiController{}
	memberGroupApiController := coremodulemembercontrollers.MemberGroupApiController{}
	memberTypeApiController := coremodulemembercontrollers.MemberTypeApiController{}

	// Member
	subrPublic.HandleFunc("/api/public/member/login", http.HandlerFunc(memberApiController.Login)).Methods("POST")

	subrPrivate.HandleFunc("/api/member", http.HandlerFunc(memberApiController.Get)).Methods("GET")
	subrPrivate.HandleFunc("/api/member/{id:.*}", http.HandlerFunc(memberApiController.GetById)).Methods("GET")

	// Member Group
	subrPrivate.HandleFunc("/api/member-group", http.HandlerFunc(memberGroupApiController.Get)).Methods("GET")
	subrPrivate.HandleFunc("/api/member-group/{id:.*}", http.HandlerFunc(memberGroupApiController.GetById)).Methods("GET")
	subrPrivate.HandleFunc("/api/member-group", http.HandlerFunc(memberGroupApiController.Post)).Methods("POST")
	subrPrivate.HandleFunc("/api/member-group/{id:.*}", http.HandlerFunc(memberGroupApiController.Put)).Methods("PUT")

	// Member type
	subrPrivate.HandleFunc("/api/member-type", http.HandlerFunc(memberTypeApiController.Get)).Methods("GET")
	subrPrivate.HandleFunc("/api/member-type/{id:.*}", http.HandlerFunc(memberTypeApiController.GetById)).Methods("GET")
	// privateApiRouter.HandleFunc("/api/member-type", http.HandlerFunc(memberTypeApiController.Post)).Methods("POST")
	// privateApiRouter.HandleFunc("/api/member-type/{id:.*}", http.HandlerFunc(memberTypeApiController.Put)).Methods("PUT")

	///

	// setup routes
	rMemberSection := lib.Route{"member", "/admin/member", "core/modules/member/public/views/member/index.html", false}
	rMemberTreeMethodEdit := lib.Route{"member.edit", "/edit/:id", "core/modules/member/public/views/member/edit.html", false}
	rMemberTreeMethodNew := lib.Route{"member.new", "/new", "core/modules/member/public/views/member/new.html", false}

	rMemberGroupSection := lib.Route{"member.memberGroup", "/member-group", "core/modules/member/public/views/member-group/index.html", false}
	rMemberGroupTreeMethodEdit := lib.Route{"member.memberGroup.edit", "/edit/:id", "core/modules/member/public/views/member-group/edit.html", false}
	rMemberGroupTreeMethodNew := lib.Route{"member.memberGroup.new", "/new", "core/modules/member/public/views/member-group/new.html", false}

	rMemberTypeSection := lib.Route{"member.memberType", "/member-type", "core/modules/member/public/views/member-type/index.html", false}
	rMemberTypeTreeMethodEdit := lib.Route{"member.memberType.edit", "/edit/:id", "core/modules/member/public/views/member-type/edit.html", false}
	rMemberTypeTreeMethodNew := lib.Route{"member.memberType.new", "/new", "core/modules/member/public/views/member-type/new.html", false}

	// setup trees
	routesMemberTree := []lib.Route{rMemberTreeMethodEdit, rMemberTreeMethodNew}
	routesMemberGroupTree := []lib.Route{rMemberGroupTreeMethodEdit, rMemberGroupTreeMethodNew}
	routesMemberTypeTree := []lib.Route{rMemberTypeTreeMethodEdit, rMemberTypeTreeMethodNew}

	tMember := lib.Tree{"Members", "members", routesMemberTree}
	tMemberGroup := lib.Tree{"Member Groups", "memberGroups", routesMemberGroupTree}
	tMemberType := lib.Tree{"Member Types", "memberTypes", routesMemberTypeTree}

	treesMemberSection := []*lib.Tree{&tMember}
	treesMemberGroupSection := []*lib.Tree{&tMemberGroup}
	treesMemberTypeSection := []*lib.Tree{&tMemberType}

	// params: name, alias, icon, route, trees, iscontainer, parent
	sMembers := lib.Section{"Members Section", "membersSection", "fa fa-users fa-fw", &rMemberSection, treesMemberSection, false, nil, nil, []string{"member_section"}}

	sMemberGroup := lib.Section{"Member Group Section", "memberGroupSection", "fa fa-smile-o fa-fw", &rMemberGroupSection, treesMemberGroupSection, false, nil, nil, []string{"member_group_section"}}

	sMemberType := lib.Section{"Member Type Section", "memberTypeSection", "fa fa-smile-o fa-fw", &rMemberTypeSection, treesMemberTypeSection, false, nil, nil, []string{"member_type_section"}}

	lol := []lib.Section{sMemberGroup, sMemberType}
	sMembers.Children = lol
	//reflect.ValueOf(&sMembers).Elem().FieldByName("Children").Set(reflect.ValueOf(lol))
	// log.Println(sMembers.Children[0].Name + ":FDSF:SDF:DS:F:")
	// res, err := json.Marshal(sMembers)
	// if err!=nil{
	//  panic(err)
	// }
	// log.Println(res)
	// log.Println("__-----------------")
	//sMembers.SetChildren(lol)
	//sMembers.Children =

	// setup members section
	// maybe add IsContainer bool?
	// section parent by name or section children - or both? For subsections ofc

	// setup subsections

	// setup module
	sections := []lib.Section{sMembers}
	// params: name, alias, description, sections
	moduleMembers := lib.Module{"Members Module", "membersModule", "Just a members module", sections}

	// register module
	lib.RegisterModule(moduleMembers)

	// FileServer
	log.Println("Registered a handler for static files. [member::module]")
	http.Handle("/core/modules/member/public/", http.FileServer(http.Dir("./")))
}
