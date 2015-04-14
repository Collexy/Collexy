package globals

import (
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

var hashKey = securecookie.GenerateRandomKey(64)
var blockKey = securecookie.GenerateRandomKey(32)

var S = securecookie.New(hashKey, blockKey)

func SetPassword(password string) (hPass string) {
    b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err) //this is a panic because bcrypt errors on invalid costs
    }
    hPass = string(b)
    return
}