package controllers

import (
	"fmt"
	"net/http"
	//"time"
	//"database/sql"
	corehelpers "collexy/core/helpers"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	//"github.com/gorilla/schema"
	"collexy/core/modules/member/models"
	coremoduleusermodels "collexy/core/modules/user/models"
	"github.com/gorilla/mux"
	//"github.com/dgrijalva/jwt-go"
	//"encoding/json"
)

type MemberApiController struct{}

func (this *MemberApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_browse", "member_all"})
		if hasPermission {

			members := models.GetMembers()
			res, err := json.Marshal(members)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse members")
		}
	}
}

func (this *MemberApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_browse", "member_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]

			memberId, _ := strconv.Atoi(idStr)

			member := models.GetMemberById(memberId)
			res, err := json.Marshal(member)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse members")
		}
	}
}

func (this *MemberApiController) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	member := models.GetMemberByUsername(username)
	cookie, err := member.Login(password)
	switch {
	case err != nil:
		log.Println(err)
	default:
		fmt.Println(member.Username + " successfully logged in")
		http.SetCookie(w, cookie)
		//fmt.Fprintf(w,"%s",tokenString)
	}
	//fmt.Fprintf(w,"username: %s\n password: ",username, password)
}

func (this *MemberApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_create", "member_all"})
		if hasPermission {

			u := models.Member{}

			err := json.NewDecoder(r.Body).Decode(&u)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			u.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create members")
		}
	}

}

func (this *MemberApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_update", "member_all"})
		if hasPermission {

			u := models.Member{}

			err := json.NewDecoder(r.Body).Decode(&u)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			u.Put()
		} else {
			fmt.Fprintf(w, "You do not have permission to update members")
		}
	}
}

func (this *MemberApiController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := coremoduleusermodels.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"member_delete", "member_all"})
		if hasPermission {
			params := mux.Vars(r)

			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			models.DeleteMember(id)
		} else {
			fmt.Fprintf(w, "You do not have permission to delete members")
		}

	}
}
