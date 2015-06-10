package models

import (
	//"code.google.com/p/go.crypto/bcrypt"
	//"crypto"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	//"github.com/dgrijalva/jwt-go"
	"time"
	//"collexy/helpers"
	coreglobals "collexy/core/globals"
	"log"
	// "io"
	// "io/ioutil"
	// "os"
	//"labix.org/v2/mgo/bson"
	"database/sql"
	"encoding/base32"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/context"
	"github.com/gorilla/securecookie"
	"net/http"
	"strings"
)

type User struct {
	Id           int          `json:"id"`
	Username     string       `json:"username"`
	FirstName    string       `json:"first_name,omitempty"`
	LastName     string       `json:"last_name,omitempty"`
	Password     []byte       `json:"-"`
	Email        string       `json:"email"`
	CreatedDate  *time.Time   `json:"created_date"`
	UpdatedDate  *time.Time   `json:"updated_date,omitempty"`
	LoginDate    *time.Time   `json:"login_date,omitempty"`
	AccessedDate *time.Time   `json:"accessed_date,omitempty"`
	Status       uint8        `json:"status"`
	Sid          string       `json:"sid,omitempty"`
	UserGroupIds []int        `json:"user_group_ids,omitempty"`
	UserGroups   []*UserGroup `json:"user_groups,omitempty"`
	Permissions  []string     `json:"permissions,omitempty"`
}

func (u *User) Post() {

	//meta, err := json.Marshal(d.Meta)
	//corehelpers.PanicIf(err)

	db := coreglobals.Db

	userGroupIds, _ := coreglobals.IntSlice(u.UserGroupIds).Value()
	permissions, _ := coreglobals.StringSlice(u.Permissions).Value()

	// sqlStr := `INSERT INTO data_type (name, alias, created_by, html, editor_alias, meta)
	// VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	// err1 := db.QueryRow(sqlStr, d.Name, d.Alias, d.CreatedBy, d.Html, d.EditorAlias, meta).Scan(&id)
	sqlStr := `INSERT INTO "user" (username, first_name, last_name, "password", email, status, 
		user_group_ids, permissions) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Todo:
	// Level: not important
	// Difficulty: easy
	// Time: quick
	// Description:
	// u.SetPassword(string(u.Password)) seems kinda stupid.
	// the SetPassword functions should just compute the hash using the struft password field
	// instead of a parameter
	u.SetPassword(string(u.Password))

	_, err1 := db.Exec(sqlStr, u.Username, u.FirstName, u.LastName, u.Password, u.Email,
		u.Status, userGroupIds, permissions)

	if err1 != nil {
		panic(err1)
	}

	log.Println("user created successfully")
}

func (u *User) Put() {

	userGroupIds, _ := coreglobals.IntSlice(u.UserGroupIds).Value()
	permissions, _ := coreglobals.StringSlice(u.Permissions).Value()

	db := coreglobals.Db

	sqlStr := `UPDATE "user" 
	SET username=$1, first_name=$2, last_name=$3, "password"=$4, email=$5, status=$6, 
		user_group_ids=$7, permissions=$8 
		WHERE id=$9`

	u.SetPassword(string(u.Password))

	_, err1 := db.Exec(sqlStr, u.Username, u.FirstName, u.LastName, u.Password, u.Email,
		u.Status, userGroupIds, permissions, u.Id)

	if err1 != nil {
		panic(err1)
	}

	log.Println("user updated successfully")
}

func DeleteUser(id int) {

	db := coreglobals.Db

	sqlStr := `delete FROM "user" 
	WHERE id=$1`

	_, err := db.Exec(sqlStr, id)

	if err != nil {
		panic(err)
	}

	log.Printf("user with id %d was successfully deleted", id)
}

//SetPassword takes a plaintext password and hashes it with bcrypt and sets the
//password field to the hash.
func (u *User) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //this is a panic because bcrypt errors on invalid costs
	}
	u.Password = hpass
	fmt.Println("hpass: " + string(hpass))
}

//Login validates and returns a user object if they exist in the database.
// func Login(password string) (u *User, err error) {
// 	// err = ctx.C("users").Find(bson.M{"username": username}).One(&u)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
// 	if err != nil {
// 		u = nil
// 	}
// 	return
// }

// func (u *User) Login(password string) (cookie *http.Cookie, err error) {
func (u *User) Login(password string) (cookie *http.Cookie, err error) {
	err = nil
	fmt.Println([]byte(password))

	// Hashing the password with the cost of 10

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		//fmt.Printf("[%s] != \n[%s]\n", u.Password, []byte(password))
		//tokenString = ""
		log.Println("login failed, try again.")
	} else {
		sid := u.GenerateSessionId(32)
		// u.InsertSessionId(sid)
		t, err := u.CreateCookie(sid)
		//helpers.PanicIf(err)
		log.Println(err)
		cookie = t

		//
		db := coreglobals.Db

		querystr := `UPDATE "user" SET login_date=$1, sid=$2 where id=$3`
		loggedInTime := time.Now()
		timeFormatted := loggedInTime.Format("2006-01-02 15:04:05.000")

		_, err1 := db.Exec(querystr, timeFormatted, sid, u.Id)
		log.Println(err1)
		//

	}
	//fmt.Println(err)
	return
}

func (u *User) GenerateSessionId(length int) (sid string) {
	//sessionid := strings.TrimRight(base32.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32)), "=")
	sid = strings.TrimRight(
		base32.StdEncoding.EncodeToString(
			securecookie.GenerateRandomKey(length)), "=")
	return
}

func (u *User) CreateCookie(sid string) (cookie *http.Cookie, err error) {
	value := map[string]string{
		"sid": sid,
	}

	// Store the session ID key in a cookie so it can be looked up in DB later.
	encoded, err := coreglobals.S.Encode("sessionauth", value)

	cookie = &http.Cookie{
		Name:  "sessionauth",
		Value: encoded,
		Path:  "/",
	}
	return
}

func GetUsers() (users []*User) {
	db := coreglobals.Db

	rows, err := db.Query(`SELECT id, username, first_name, last_name, password, email,
      created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids 
      FROM "user"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, email string
		var password []byte
		//var meta []byte
		var created_date, updated_date, login_date, accessed_date *time.Time
		var status uint8
		// coreglobals.StringSlice is temporarily defined in template.go model
		var user_group_ids coreglobals.IntSlice

		var sid, first_name, last_name sql.NullString

		if err := rows.Scan(&id, &username, &first_name, &last_name, &password, &email, &created_date, &updated_date, &login_date, &accessed_date, &status, &sid, &user_group_ids); err != nil {
			log.Fatal(err)
		}

		// var metaMap map[string]interface{}
		// err1 := json.Unmarshal(meta, &metaMap)
		// if err1 != nil {
		//   fmt.Println("error:", err)
		// }

		var user_sid string
		if sid.Valid {
			// use s.String
			user_sid = sid.String
		} else {
			// NULL value
		}

		var user_first_name string
		if first_name.Valid {
			// use s.String
			user_first_name = first_name.String
		} else {
			// NULL value
		}

		var user_last_name string
		if last_name.Valid {
			// use s.String
			user_last_name = last_name.String
		} else {
			// NULL value
		}

		user := &User{id, username, user_first_name, user_last_name, password, email, created_date, updated_date, login_date, accessed_date, status, user_sid, user_group_ids, nil, nil}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func GetUserById(id int) (user *User) {
	db := coreglobals.Db

	var username, email string
	var password []byte
	//var meta []byte
	var created_date, updated_date, login_date, accessed_date *time.Time
	var status uint8
	// coreglobals.IntSlice is temporarily defined in template.go model
	var user_group_ids coreglobals.IntSlice

	var sid, first_name, last_name sql.NullString

	err := db.QueryRow(`SELECT username, first_name, last_name, password, email,
      created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids 
      FROM "user" WHERE id=$1`, id).Scan(&username, &first_name, &last_name, &password, &email, &created_date, &updated_date, &login_date, &accessed_date, &status, &sid, &user_group_ids)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No member with that ID.")
	case err != nil:
		log.Println("lolol")
		log.Fatal(err)
	default:
		var user_sid string
		if sid.Valid {
			// use s.String
			user_sid = sid.String
		} else {
			// NULL value
		}

		var user_first_name string
		if first_name.Valid {
			// use s.String
			user_first_name = first_name.String
		} else {
			// NULL value
		}

		var user_last_name string
		if last_name.Valid {
			// use s.String
			user_last_name = last_name.String
		} else {
			// NULL value
		}

		user = &User{id, username, user_first_name, user_last_name, password, email, created_date, updated_date, login_date, accessed_date, status, user_sid, user_group_ids, nil, nil}
	}
	return
}

func GetUser(sid string) (u *User, err error) {
	// check memcache first
	fmt.Println("Memcache is searching for session ID: " + sid)
	it, err := coreglobals.Mc.Get(sid)

	if err == nil {
		fmt.Println("Memcache did'nt report any errors")
		if it != nil {
			fmt.Println("item is not nil")
			fmt.Println("it value: ", it.Value)
			err1 := json.Unmarshal(it.Value, &u)
			if err1 != nil {
				fmt.Println("error1:", err1.Error())
			}

			fmt.Println("Memcache found the following user in cache lookup: " + u.Username)
			return
		}
	} else {
		fmt.Println("Memcache error: " + err.Error())
	}

	// if not in memcache, look in db
	db := coreglobals.Db
	querystr := `SELECT "user".id, "user".username, "user".first_name, "user".last_name, "user".password, 
  "user".email, "user".created_date, "user".updated_date, "user".login_date, "user".accessed_date, "user".status, "user".sid, "user".user_group_ids, user_groups.user_groups 
FROM "user",
LATERAL (
  SELECT array_to_json(array_agg(user_group_agg)) AS user_groups
  FROM (
    SELECT * 
    FROM user_group
    WHERE user_group.id = ANY ("user".user_group_ids)
    --WHERE "user".user_group_ids @> user_group.id::text
  ) user_group_agg
) user_groups
WHERE sid=$1`
	// querystr := `SELECT * FROM "user" where sid=$1`
	//   querystr := `SELECT "user".id, "user".username, "user".first_name, "user".last_name, "user".password,
	//   "user".email, "user".created_date, "user".updated_date, "user".login_date, "user".accessed_date, "user".status, "user".sid, "user".user_group_ids, user_groups.user_groups
	// FROM "user",
	// LATERAL (
	//   SELECT array_to_json(array_agg(user_group_agg)) AS user_groups
	//   FROM (
	//     SELECT user_group.*, inner1.permissions
	//     FROM user_group,
	//     LATERAL (
	//       SELECT array_to_json(array_agg(permission_agg)) AS permissions
	//       FROM (
	//         SELECT * FROM permission
	//         WHERE id = ANY (user_group.permission_ids)
	//       ) permission_agg
	//     ) inner1
	//     WHERE user_group.id = ANY ("user".user_group_ids)
	//   ) user_group_agg
	// ) user_groups
	// WHERE sid=$1`

	var id int
	var status uint8
	var username, email string
	var password []byte
	var created_date, accessed_date, updated_date, login_date *time.Time

	// potential nulls
	// var user_group_ids []int // doesn't work with scan
	// coreglobals.IntSlice custom type is right now located in models.Template
	var user_group_ids coreglobals.IntSlice
	var user_groups []byte
	var first_name, last_name sql.NullString

	err = db.QueryRow(querystr, sid).Scan(&id, &username, &first_name, &last_name, &password, &email, &created_date, &updated_date, &login_date, &accessed_date, &status, &sid, &user_group_ids, &user_groups)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that Session ID.")
	case err != nil:
		log.Fatal(err)
	default:
		var fname, lname string
		if first_name.Valid {
			fname = first_name.String
		} else {
			fname = ""
		}

		if last_name.Valid {
			lname = first_name.String
		} else {
			lname = ""
		}
		//fmt.Printf("Username is %s\n", username)
		var user_groupsSlice []*UserGroup
		json.Unmarshal(user_groups, &user_groupsSlice)
		u = &User{id, username, fname, lname, password, email, created_date, updated_date,
			login_date, accessed_date, status, sid, user_group_ids, user_groupsSlice, nil}

		uByteArr, err := json.Marshal(u)
		if err != nil {
			fmt.Println("error:", err)
		}
		//mc.Set(&memcache.Item{Key: "users", Value: []byte("my value")})
		coreglobals.Mc.Set(&memcache.Item{Key: sid, Value: uByteArr})
	}
	return
}

// MEMCACHE

type key int

const loggedInUser key = 0

// GetLoggedInUser returns a value for this package from the request values.
func GetLoggedInUser(r *http.Request) *User {
	if rv := context.Get(r, loggedInUser); rv != nil {
		return rv.(*User)
	}
	return nil
}

// SetLoggedInUser sets a value for this package in the request values.
func SetLoggedInUser(r *http.Request, val *User) {
	context.Set(r, loggedInUser, val)
}

// func (u *User) HasNodePermissions (permissions []string) (hasPermissions bool){
//   return
// }

func (u *User) HasPermissions(permissions []string) (hasPermissions bool) {
	permFound := false
	hasPermissions = false

	// First check if a the currently logged in user has specific permissions per user-level
	if u.Permissions != nil {
		if len(u.Permissions) > 0 {

		}
	} else if u.UserGroups != nil { // If first check fails, check permissions for each group if any
		if len(u.UserGroups) > 0 {
			for i := 0; i < len(permissions); i++ {
				permFound = false
				for j := 0; j < len(u.UserGroups); j++ {
					if permFound {
						break
					}
					for k := 0; k < len(u.UserGroups[j].Permissions); k++ {
						if permFound {
							break
						}
						if permissions[i] == u.UserGroups[j].Permissions[k] {
							permFound = true
						}
					}
				}
			}
		}
	}
	hasPermissions = permFound
	return
}
