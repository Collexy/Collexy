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
	"github.com/gorilla/mux"
	//"github.com/dgrijalva/jwt-go"
	//"encoding/json"
)

type MemberApiController struct{}

func (this *MemberApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	members := models.GetMembers()
	res, err := json.Marshal(members)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *MemberApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]

	memberId, _ := strconv.Atoi(idStr)

	member := models.GetMemberById(memberId)
	res, err := json.Marshal(member)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
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

// func (this *MemberApiController) Post(w http.ResponseWriter, r *http.Request) {
//     user := new(models.User)

//     err := r.ParseForm()

//     corehelpers.PanicIf(err)

//     decoder := schema.NewDecoder()
//     // r.PostForm is a map of our POST form values
//     decoder.Decode(user, r.PostForm)

//     fmt.Println(r.PostForm)
//     fmt.Println(user.FirstName)
//     fmt.Println(user.Password)
//     fmt.Println(r.FormValue("Password"))

//     db := corehelpers.Db

//     // http://stackoverflow.com/questions/244243/how-to-reset-postgres-primary-key-sequence-when-it-falls-out-of-sync
//     //fmt.Println(fmt.Sprintf("path: %s, created_by: %d, label: %s, User type: %d", t.Path, t.Created_by, t.Label, t.User_type))
//     lol := string(r.FormValue("Password"))
//     user.SetPassword(lol)

//     // password := user.Password
//     fmt.Println(fmt.Sprintf("username: %s, first name: %s, last name: %s, password: %s", user.Username, user.FirstName, user.LastName, user.Password))

//     querystr := fmt.Sprintf("INSERT INTO \"user\" (username, first_name, last_name, password) VALUES ('%s','%s','%s','%s')", user.Username, user.FirstName, user.LastName, user.Password)
//     fmt.Println("querystring: " + querystr)
//     res, err := db.Exec(querystr)
//     corehelpers.PanicIf(err)
//     fmt.Println(res)

// }

// func (this *UserController) Delete(w http.ResponseWriter, r *http.Request) {

//     params := mux.Vars(r)
//     idStr := params["id"]
//     id, _ := strconv.Atoi(idStr)

//     parm_id := id

//     db := corehelpers.Db

//     querystr := fmt.Sprintf("DELETE FROM \"user\" WHERE id=%d", parm_id)
//     res, err := db.Exec(querystr)
//     corehelpers.PanicIf(err)
//     fmt.Println(res)

// }

// type User struct {
//   Username string `json:"username,omitempty"`
//   Password string `json:"password"`
// }

// func (this *MemberApiController) Login(w http.ResponseWriter, r *http.Request) {}

// // func (this *UserController) ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
// //     if cookie, err := r.Cookie("cookie-name-test"); err == nil {
// //         value := make(map[string]string)
// //         // if err = s2.Decode("cookie-name-test", cookie.Value, &value); err == nil {
// //         //     fmt.Fprintf(w, "The value of foo is %q", value["foo"])
// //         // }
// //     }
// // }
