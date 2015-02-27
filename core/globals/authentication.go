package globals

import (
	//corehelpers "collexy/core/helpers"
	"github.com/gorilla/securecookie"
)

var hashKey = securecookie.GenerateRandomKey(64)
var blockKey = securecookie.GenerateRandomKey(32)

var S = securecookie.New(hashKey, blockKey)

