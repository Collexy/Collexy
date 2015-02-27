package controllers

import (
    "fmt"
    "net/http"
    //"time"
    //"database/sql"
    _ "github.com/lib/pq"
    "collexy/helpers"
    coreglobals "collexy/core/globals"
    "strconv"
    "log"
    "encoding/json"
    "github.com/gorilla/schema"
    "collexy/core/api/models"
    "github.com/gorilla/mux"
    //"github.com/dgrijalva/jwt-go"
    //"encoding/json"
)

type UserApiController struct {}

func (this *UserApiController) Get(w http.ResponseWriter, r *http.Request) {
    db := coreglobals.Db
    rows, err := db.Query("SELECT id, username, first_name, last_name, password FROM \"user\"")
    helpers.PanicIf(err)
    defer rows.Close()

    // var id, created_by, User_type int
    // var path, label string
    // var created_date time.Time
    var id int
    var username, first_name, last_name, password string

    for rows.Next(){
        err := rows.Scan(&id, &username, &first_name, &last_name, &password)
        helpers.PanicIf(err)
        fmt.Fprintf(w, "Id: %d, Username: %s, First: %s, Last: %s, Password:%s\n", id, username, first_name, last_name, password)
    }

}

// func (this *UserApiController) GetById(w http.ResponseWriter, r *http.Request) {

//     var parm_id, id, created_by, User_type int
//     var path, label string
//     var created_date time.Time

//     templol := r.URL.Query().Get(":id")
//     rofl,err1 := strconv.Atoi(templol)
//     helpers.PanicIf(err1)

//     parm_id = rofl

//     db := coreglobals.Db

//     row := db.QueryRow("SELECT id, path, created_by, label, User_type, created_date FROM User WHERE id=$1", parm_id)
//     err:= row.Scan(&id, &path, &created_by, &label, &User_type, &created_date)
    
//     //helpers.PanicIf(err)
//     switch {
//         case err == sql.ErrNoRows:
//                 log.Printf("No User with that ID.")
//         case err != nil:
//                 log.Fatal(err)
//         default:
//                 fmt.Fprintf(w, "Id: %d, Path: %s, Created by: %d, Label: %s, User type: %d, Created date: %s\n", id, path, created_by, label, User_type, created_date)
//     }


// }

// type test_struct struct {
//     Id int
//     Path string
//     Created_by int
//     Label string
//     User_type int
//     Created_date time.Time
// }

func (this *UserApiController) Post(w http.ResponseWriter, r *http.Request) {
    user := new(models.User)

    err := r.ParseForm()

    helpers.PanicIf(err)

    decoder := schema.NewDecoder()
    // r.PostForm is a map of our POST form values
    decoder.Decode(user, r.PostForm)


    fmt.Println(r.PostForm)
    fmt.Println(user.FirstName)
    fmt.Println(user.Password)
    fmt.Println(r.FormValue("Password"))

    db := coreglobals.Db

    // http://stackoverflow.com/questions/244243/how-to-reset-postgres-primary-key-sequence-when-it-falls-out-of-sync
    //fmt.Println(fmt.Sprintf("path: %s, created_by: %d, label: %s, User type: %d", t.Path, t.Created_by, t.Label, t.User_type))
    lol := string(r.FormValue("Password"))
    user.SetPassword(lol)
    
    // password := user.Password
    fmt.Println(fmt.Sprintf("username: %s, first name: %s, last name: %s, password: %s", user.Username, user.FirstName, user.LastName, user.Password))

    querystr := fmt.Sprintf("INSERT INTO \"user\" (username, first_name, last_name, password) VALUES ('%s','%s','%s','%s')", user.Username, user.FirstName, user.LastName, user.Password)
    fmt.Println("querystring: " + querystr)
    res, err := db.Exec(querystr)
    helpers.PanicIf(err)
    fmt.Println(res)
    
}

// func (this *UserApiController) Put(w http.ResponseWriter, r *http.Request) {
//     // path := r.FormValue("path")
//     // created_by,err := strconv.Atoi(r.FormValue("created_by"))
//     // label := r.FormValue("label")
//     // User_type,err := strconv.Atoi(r.FormValue("User_type"))


//     templol := r.URL.Query().Get(":id")
//     rofl,err1 := strconv.Atoi(templol)
//     helpers.PanicIf(err1)

//     parm_id := rofl

//     t := new(test_struct)

//     err := r.ParseForm()

//     helpers.PanicIf(err)

//     decoder := schema.NewDecoder()
//     // r.PostForm is a map of our POST form values
//     decoder.Decode(t, r.PostForm)


//     fmt.Println(r.PostForm)
//     fmt.Println(t.Path)

//     db := coreglobals.Db

//     // http://stackoverflow.com/questions/244243/how-to-reset-postgres-primary-key-sequence-when-it-falls-out-of-sync
//     fmt.Println(fmt.Sprintf("path: %s, created_by: %d, label: %s, User type: %d", t.Path, t.Created_by, t.Label, t.User_type))


//     querystr := fmt.Sprintf("UPDATE User SET (path, created_by, label, User_type) = ('%s', %d, '%s', %d) WHERE id=%d", t.Path, t.Created_by, t.Label, t.User_type, parm_id)
//     res, err := db.Exec(querystr)
//     helpers.PanicIf(err)
//     fmt.Println(res)
    
//     // JSON(w, r.Body)
// }

func (this *UserApiController) Delete(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    idStr := params["id"]
    id, _ := strconv.Atoi(idStr)

    parm_id := id

   
    db := coreglobals.Db
    
    querystr := fmt.Sprintf("DELETE FROM \"user\" WHERE id=%d", parm_id)
    res, err := db.Exec(querystr)
    helpers.PanicIf(err)
    fmt.Println(res)

}

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
        helpers.PanicIf(err)
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