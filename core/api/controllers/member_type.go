package controllers

import (
    "fmt"
    "net/http"
    //"time"
    //"database/sql"
    _ "github.com/lib/pq"
    corehelpers "collexy/core/helpers"
    "collexy/core/api/models"
    "strconv"
    //"log"
    //"github.com/gorilla/schema"
    "encoding/json"
    //"log"
    //"io/ioutil"
    //"path/filepath"
    //"strings"
    //"html/template"
    "github.com/gorilla/mux"
)

type MemberTypeApiController struct {}

// func (this *MemberTypeApiController) Post(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Member-Type", "application/json")

//     memberType := models.MemberType{}

//     err := json.NewDecoder(r.Body).Decode(&memberType)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     memberType.Post()

// }

// func (this *MemberTypeApiController) GetMemberTypes(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Member-Type", "application/json")

//     memberTypes := models.GetMemberTypes()
//     res, err := json.Marshal(memberTypes)
//     corehelpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)
// }

func (this *MemberTypeApiController) GetMemberTypeByNodeId(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    idStr := params["nodeId"]

    nodeId, _ := strconv.Atoi(idStr)

    var extended bool = false
    extended, _ = strconv.ParseBool(r.URL.Query().Get("extended"))

    //extended, _ := strconv.Atoi(extendedStr)

    if(!extended){
        memberType := models.GetMemberTypeByNodeId(nodeId)
        res, err := json.Marshal(memberType)
        corehelpers.PanicIf(err)

        fmt.Fprintf(w,"%s",res)
    } else {
        memberType := models.GetMemberTypeExtendedByNodeId(nodeId)
        res, err := json.Marshal(memberType)
        corehelpers.PanicIf(err)

        fmt.Fprintf(w,"%s",res)
    }
    
}


// func (this *MemberTypeApiController) PutMemberType(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Member-Type", "application/json")

//     memberType := models.MemberType{}

//     err := json.NewDecoder(r.Body).Decode(&memberType)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }
    
//     memberType.Update()
// }