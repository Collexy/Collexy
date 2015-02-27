package helpers

import(
    //"github.com/dgrijalva/jwt-go"
    "fmt"
    coreglobals "collexy/core/globals"
    // "collexy/admin/models" // Avoid circular dependencies, go compiler however will tell you!
    "net/http"
)

// func SetupSecureCookie() (sc *securecookie.SecureCookie){
//     var hashKey = securecookie.GenerateRandomKey(64)
//     var blockKey = securecookie.GenerateRandomKey(32)

//     sc = securecookie.New(hashKey, blockKey)
//     return
// }

func CheckCookie(w http.ResponseWriter, r *http.Request) (sid string){
    if cookie, err := r.Cookie("sessionauth"); err == nil {
        value := make(map[string]string)
        if err = coreglobals.S.Decode("sessionauth", cookie.Value, &value); err == nil {
            sid = value["sid"]
            fmt.Println("helpers.CheckCookie returns sid (string): " + sid)
        }
    }
    return
}

func CheckMemberCookie(w http.ResponseWriter, r *http.Request) (sid string){
    if cookie, err := r.Cookie("membersessionauth"); err == nil {
        value := make(map[string]string)
        if err = coreglobals.S.Decode("membersessionauth", cookie.Value, &value); err == nil {
            sid = value["sid"]
            fmt.Println("helpers.CheckMemberCookie returns sid (string): " + sid)
        }
    }
    return
}