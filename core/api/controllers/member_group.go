package controllers

import (
    "fmt"
    "net/http"
    _ "github.com/lib/pq"
    corehelpers "collexy/core/helpers"
    "collexy/core/api/models"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
)

type MemberGroupApiController struct {}

func (this *MemberGroupApiController) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    memberGroups := models.GetMemberGroups()
    res, err := json.Marshal(memberGroups)
    corehelpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

func (this *MemberGroupApiController) GetById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    idStr := params["id"]

    id, _ := strconv.Atoi(idStr)

    memberGroup := models.GetMemberGroupById(id)
    res, err := json.Marshal(memberGroup)
    corehelpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
    
}

func (this *MemberGroupApiController) Post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    memberGroup := models.MemberGroup{}

    err := json.NewDecoder(r.Body).Decode(&memberGroup)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    memberGroup.Post()

}

func (this *MemberGroupApiController) Put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    memberGroup := models.MemberGroup{}

    err := json.NewDecoder(r.Body).Decode(&memberGroup)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    memberGroup.Put()

}