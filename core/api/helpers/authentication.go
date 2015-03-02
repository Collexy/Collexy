package helpers

import(
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "encoding/json"
    corehelpers "collexy/core/helpers"
    "collexy/core/api/models"
    coreglobals "collexy/core/globals"
)

func AngularAuth(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    encodedSid := params["sid"]
    var sid string
    value := make(map[string]string)
    if err := coreglobals.S.Decode("sessionauth", encodedSid, &value); err == nil {
        sid = value["sid"]
        fmt.Println("corehelpers.CheckCookie returns sid (string): " + sid)
    }
    
    u, _ := models.GetUser(sid)
    res, err := json.Marshal(u)
    corehelpers.PanicIf(err)
    fmt.Fprintf(w,"%s",res)
}