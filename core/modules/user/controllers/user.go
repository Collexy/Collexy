package controllers

import (
	"fmt"
	"net/http"
	//"time"
	//"database/sql"
	coreglobals "collexy/core/globals"
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/user/models"
	"encoding/json"
	"github.com/gorilla/mux"
	//"github.com/gorilla/schema"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	//"github.com/dgrijalva/jwt-go"
	//"encoding/json"
)

type UserApiController struct{}

func (this *UserApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_update", "user_all"})
		if hasPermission {
			users := models.GetUsers()
			res, err := json.Marshal(users)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse users")
		}
	}	
}

func (this *UserApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_update", "user_all"})
		if hasPermission {

			params := mux.Vars(r)
			idStr := params["id"]

			userId, _ := strconv.Atoi(idStr)

			user := models.GetUserById(userId)
			res, err := json.Marshal(user)
			corehelpers.PanicIf(err)

			fmt.Fprintf(w, "%s", res)
		} else {
			fmt.Fprintf(w, "You do not have permission to browse users")
		}
	}	
}

func (this *UserApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_create", "user_all"})
		if hasPermission {

			u := models.User{}

			err := json.NewDecoder(r.Body).Decode(&u)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			u.Post()
		} else {
			fmt.Fprintf(w, "You do not have permission to create users")
		}
	}

}

func (this *UserApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_update", "user_all"})
		if hasPermission {

			u := models.User{}

			err := json.NewDecoder(r.Body).Decode(&u)

			if err != nil {
				http.Error(w, "Bad Request", 400)
			}

			u.Put()
		} else {
			fmt.Fprintf(w, "You do not have permission to update users")
		}
	}
}

func (this *UserApiController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if user := models.GetLoggedInUser(r); user != nil {
		var hasPermission bool = false
		hasPermission = user.HasPermissions([]string{"user_delete", "user_all"})
		if hasPermission {
			params := mux.Vars(r)

			idStr := params["id"]
			id, _ := strconv.Atoi(idStr)

			models.DeleteUser(id)
		} else {
			fmt.Fprintf(w, "You do not have permission to delete users")
		}

	}
}

// func (this *UserApiController) Post(w http.ResponseWriter, r *http.Request) {
// 	user := new(models.User)

// 	err := r.ParseForm()

// 	corehelpers.PanicIf(err)

// 	decoder := schema.NewDecoder()
// 	// r.PostForm is a map of our POST form values
// 	decoder.Decode(user, r.PostForm)

// 	fmt.Println(r.PostForm)
// 	fmt.Println(user.FirstName)
// 	fmt.Println(user.Password)
// 	fmt.Println(r.FormValue("Password"))

// 	db := coreglobals.Db

// 	// http://stackoverflow.com/questions/244243/how-to-reset-postgres-primary-key-sequence-when-it-falls-out-of-sync
// 	//fmt.Println(fmt.Sprintf("path: %s, created_by: %d, label: %s, User type: %d", t.Path, t.Created_by, t.Label, t.User_type))
// 	lol := string(r.FormValue("Password"))
// 	user.SetPassword(lol)

// 	// password := user.Password
// 	fmt.Println(fmt.Sprintf("username: %s, first name: %s, last name: %s, password: %s", user.Username, user.FirstName, user.LastName, user.Password))

// 	querystr := fmt.Sprintf("INSERT INTO \"user\" (username, first_name, last_name, password) VALUES ('%s','%s','%s','%s')", user.Username, user.FirstName, user.LastName, user.Password)
// 	fmt.Println("querystring: " + querystr)
// 	res, err := db.Exec(querystr)
// 	corehelpers.PanicIf(err)
// 	fmt.Println(res)

// }



type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

func (this *UserApiController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//r.ParseForm()
	defer r.Body.Close()

	v := new(User)

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		// error
	}

	fmt.Println(v)

	var id int
	var username, first_name, last_name string
	var password []byte
	db := coreglobals.Db

	stmt, err := db.Prepare("SELECT id, username, first_name, last_name, password FROM \"user\" WHERE username=$1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// rows, err := stmt.Query(r.FormValue("Username"))
	rows, err := stmt.Query(v.Username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// ...
		err := rows.Scan(&id, &username, &first_name, &last_name, &password)
		corehelpers.PanicIf(err)
	}

	user := models.User{Id: id, Username: username, FirstName: first_name, LastName: last_name, Password: password}

	tokenString, err := user.Login(v.Password)
	switch {
	case err != nil:
		log.Println(err)
	default:
		fmt.Println(user.Username + " successfully logged in")
		http.SetCookie(w, tokenString)
		//fmt.Fprintf(w,"%s",tokenString)
	}
}

// func (this *UserApiController) ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
//     if cookie, err := r.Cookie("cookie-name-test"); err == nil {
//         value := make(map[string]string)
//         // if err = s2.Decode("cookie-name-test", cookie.Value, &value); err == nil {
//         //     fmt.Fprintf(w, "The value of foo is %q", value["foo"])
//         // }
//     }
// }
