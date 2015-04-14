package globals

import(
	"github.com/gorilla/mux"
)

var PrivateApiRouter *mux.Router = mux.NewRouter()
var PublicApiRouter *mux.Router = mux.NewRouter()

type IRoute interface {
    AddChildren(child IRoute)
}

//var Routes []IRoute