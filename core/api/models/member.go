package models

import (
	//"code.google.com/p/go.crypto/bcrypt"
	//"crypto"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"encoding/json"
	//"github.com/dgrijalva/jwt-go"
	"time"
	//"collexy/helpers"
	coreglobals "collexy/core/globals"
	"log"
	// "io"
	// "io/ioutil"
	// "os"
	//"labix.org/v2/mgo/bson"
	"net/http"
	"github.com/gorilla/securecookie"
	"strings"
	"encoding/base32"
	"database/sql"
	"github.com/gorilla/context"
	"github.com/bradfitz/gomemcache/memcache"
)

type Member struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"-"`
	Email string `json:"email"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	CreatedDate *time.Time `json:"created_date"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
	LoginDate *time.Time `json:"login_date,omitempty"`
	AccessedDate *time.Time `json:"accessed_date,omitempty"`
	Status uint8 `json:"status"`
	Sid string `json:"sid,omitempty"`
	MemberTypeNodeId int `json:"member_type_node_id,omitempty"`
	GroupNodeIds []int `json:"member_group_node_ids,omitempty"`
	Groups []*MemberGroup `json:"groups,omitempty"`
}

func (m *Member) Groups2PublicAccess(contentGroups []int) bool{
    for _, contentGroup := range contentGroups {
        for _, memberGroup := range m.Groups {
            if contentGroup == memberGroup.Id {
                return true
            }
        }
    }
    return false
}

//SetPassword takes a plaintext password and hashes it with bcrypt and sets the
//password field to the hash.
func (m *Member) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //this is a panic because bcrypt errors on invalid costs
	}
	m.Password = hpass
	fmt.Println("hpass: " + string(hpass))
}

func (m *Member) GenerateSessionId(length int) (sid string){
  //sessionid := strings.TrimRight(base32.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32)), "=")
  sid = strings.TrimRight(
base32.StdEncoding.EncodeToString(
securecookie.GenerateRandomKey(length)), "=")
  return
}

func (m *Member) CreateCookie(sid string) (cookie *http.Cookie, err error){
  value := map[string]string{
        "sid": sid,
    }

  // Store the session ID key in a cookie so it can be looked up in DB later.
  encoded, err := coreglobals.S.Encode("membersessionauth", value)

  cookie = &http.Cookie{
            Name:  "membersessionauth",
            Value: encoded,
            Path:  "/",
        }
  return
}


func (m *Member) Login(password string) (cookie *http.Cookie, err error){
	err = nil
  	fmt.Println([]byte(password))

  	// Hashing the password with the cost of 10
    // Comparing the password with the hash
    err = bcrypt.CompareHashAndPassword(m.Password, []byte(password))
    if err != nil {
		//fmt.Printf("[%s] != \n[%s]\n", u.Password, []byte(password))
		//tokenString = ""
		log.Println("login failed, try again.")
    } else {
		sid := m.GenerateSessionId(32)
		// u.InsertSessionId(sid)
		t,err := m.CreateCookie(sid)
		//helpers.PanicIf(err)
		log.Println(err)
		cookie = t

		//
		db := coreglobals.Db

		querystr := `UPDATE member SET login_date=$1, sid=$2 where id=$3`
		loggedInTime := time.Now()
		timeFormatted := loggedInTime.Format("2006-01-02 15:04:05.000")

		_, err1 := db.Exec(querystr, timeFormatted, sid, m.Id)
		log.Println(err1)
		//
      
    }
    //fmt.Println(err)
    return
}

func GetMemberByUsername(username string)(member *Member){
	db := coreglobals.Db

	var id, member_type_node_id int
	var email string
	var password []byte
	var meta []byte
	var created_date, updated_date, login_date, accessed_date *time.Time
	var status uint8
	// IntArray is temporarily defined in template.go model
	var member_group_node_ids IntArray

	var sid sql.NullString

    err := db.QueryRow(`SELECT id, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_node_id, member_group_node_ids 
    	FROM member WHERE username=$1`,username).Scan(&id,&password,&email,&meta,&created_date,&updated_date,&login_date,&accessed_date,&status,&sid,&member_type_node_id,&member_group_node_ids)
    switch {
	    case err == sql.ErrNoRows:
	        log.Printf("No member with that ID.")
	    case err != nil:
	        log.Fatal(err)
	    default:
			var metaMap map[string]interface{}
			err := json.Unmarshal(meta, &metaMap)
			if err != nil {
				fmt.Println("error:", err)
			}

			var member_sid string
			if sid.Valid {
				// use s.String
				member_sid = sid.String
			} else { 
				// NULL value 
			}
			member = &Member{id,username,password,email,metaMap,created_date,updated_date,login_date,accessed_date,status,member_sid,member_type_node_id,member_group_node_ids,nil}
	}
	return
}

func GetMembers() (members []*Member){
	db := coreglobals.Db

    rows, err := db.Query("SELECT * FROM member")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var id, member_type_node_id int
		var username, email string
		var password []byte
		var meta []byte
		var created_date, updated_date, login_date, accessed_date *time.Time
		var status uint8
		// IntArray is temporarily defined in template.go model
		var member_group_node_ids IntArray

		var sid sql.NullString

        if err := rows.Scan(&id,&username,&password,&email,&meta,&created_date,&updated_date,&login_date,&accessed_date,&status,&sid,&member_type_node_id,&member_group_node_ids); err != nil {
                log.Fatal(err)
        }

        var metaMap map[string]interface{}
		err1 := json.Unmarshal(meta, &metaMap)
		if err1 != nil {
			fmt.Println("error:", err)
		}

		var member_sid string
		if sid.Valid {
			// use s.String
			member_sid = sid.String
		} else { 
			// NULL value 
		}
		member := &Member{id,username,password,email,metaMap,created_date,updated_date,login_date,accessed_date,status,member_sid,member_type_node_id,member_group_node_ids,nil}
		members = append(members, member)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    return
}

func GetMemberById(id int) (member *Member){
	db := coreglobals.Db

    var member_type_node_id int
	var username, email string
	var password []byte
	var meta []byte
	var created_date, updated_date, login_date, accessed_date *time.Time
	var status uint8
	// IntArray is temporarily defined in template.go model
	var member_group_node_ids IntArray

	var sid sql.NullString

    err := db.QueryRow(`SELECT username, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_node_id, member_group_node_ids 
    	FROM member WHERE id=$1`,id).Scan(&username,&password,&email,&meta,&created_date,&updated_date,&login_date,&accessed_date,&status,&sid,&member_type_node_id,&member_group_node_ids)
    switch {
    case err == sql.ErrNoRows:
        log.Printf("No member with that ID.")
    case err != nil:
        log.Fatal(err)
    default:
		var metaMap map[string]interface{}
		err := json.Unmarshal(meta, &metaMap)
		if err != nil {
			fmt.Println("error:", err)
		}

		var member_sid string
		if sid.Valid {
			// use s.String
			member_sid = sid.String
		} else { 
			// NULL value 
		}
		member = &Member{id,username,password,email,metaMap,created_date,updated_date,login_date,accessed_date,status,member_sid,member_type_node_id,member_group_node_ids,nil}
        fmt.Printf("Username is %s\n", username)
    }
    return
}

func GetMember (sid string) (m *Member, err error){
  // check memcache first
  fmt.Println("Memcache is searching for member session ID: " + sid)
  it, err := coreglobals.Mc.Get(sid)

  if(err == nil){
    fmt.Println("Memcache did'nt report any errors")
    if(it != nil){
      fmt.Println("item is not nil")
      fmt.Println("it value: ", it.Value)
      err1 := json.Unmarshal(it.Value, &m)
      if err1 != nil {
        fmt.Println("error1:", err1.Error())
      }
      
      fmt.Println("Memcache found the following user in cache lookup: " + m.Username)
      return
    }
  } else {
    fmt.Println("Memcache error: " + err.Error())
  }

  // if not in memcache, look in db
  db := coreglobals.Db
  // querystr := `SELECT * FROM member where sid=$1`
//   querystr := `SELECT member.id, member.username, member.password, 
//   member.email, member.created_date, member.updated_date, member.login_date, member.accessed_date, member.status, member.sid, member.role_ids, roles.roles 
// FROM member,
// LATERAL (
//   SELECT array_to_json(array_agg(role_agg)) AS roles
//   FROM (
//     SELECT member_role.*, inner1.permissions 
//     FROM member_role,
//     LATERAL (
//       SELECT array_to_json(array_agg(permission_agg)) AS permissions
//       FROM ( 
//         SELECT * FROM member_permission
//         WHERE id = ANY (member_role.permission_ids)
//       ) permission_agg
//     ) inner1
//     WHERE member_role.id = ANY (member.role_ids)
//   ) role_agg
// ) roles
// WHERE sid=$1`

  querystr := `SELECT member.id, member.username, member.password, 
  member.email, member.created_date, member.updated_date, member.login_date, member.accessed_date, member.status, member.sid, member.member_group_node_ids, groups.groups 
FROM member,
LATERAL (
  SELECT array_to_json(array_agg(group_agg)) AS groups
  FROM (
    SELECT node.id, node.name
    FROM node
    WHERE node.id = ANY (member.member_group_node_ids) and node.node_type=13
  ) group_agg
) groups
WHERE sid=$1`

  var id int
  var status uint8
  var username, email string
  var password []byte
  var created_date, accessed_date, updated_date, login_date *time.Time

  // potential nulls
  // var role_ids []int // doesn't work with scan
  // IntArray custom type is right now located in models.Template
  var member_group_node_ids IntArray
  var groups []byte

  err = db.QueryRow(querystr, sid).Scan(&id, &username, &password, &email, &created_date, &updated_date, &login_date, &accessed_date, &status, &sid, &member_group_node_ids, &groups)

  switch {
    case err == sql.ErrNoRows:
      log.Printf("No user with that Session ID.")
    case err != nil:
      log.Fatal(err)
    default:
      
      //fmt.Printf("Username is %s\n", username)
      var groupsSlice []*MemberGroup
      json.Unmarshal(groups, &groupsSlice)
      m = &Member{id, username, password, email, nil, created_date, updated_date,
        login_date, accessed_date, status, sid, 0, member_group_node_ids, groupsSlice}

      uByteArr, err := json.Marshal(m)
      if err != nil {
        fmt.Println("error:", err)
      }
      //mc.Set(&memcache.Item{Key: "users", Value: []byte("my value")})
      coreglobals.Mc.Set(&memcache.Item{Key: sid, Value: uByteArr})
  }
  return
}

// MEMCACHE

//type key int

const loggedInMember key = 0

// GetLoggedInUser returns a value for this package from the request values.
func GetLoggedInMember(r *http.Request) *Member {
    if rv := context.Get(r, loggedInMember); rv != nil {
        return rv.(*Member)
    }
    return nil
}

// SetLoggedInUser sets a value for this package in the request values.
func SetLoggedInMember(r *http.Request, val *Member) {
    context.Set(r, loggedInMember, val)
}