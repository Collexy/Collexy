package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/member/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

type MemberGroupApiController struct{}

func (this *MemberGroupApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	memberGroups := models.GetMemberGroups()
	res, err := json.Marshal(memberGroups)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *MemberGroupApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]

	id, _ := strconv.Atoi(idStr)

	memberGroup := models.GetMemberGroupById(id)
	res, err := json.Marshal(memberGroup)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)

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
