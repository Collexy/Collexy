package controllers

import (
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/user/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserGroupApiController struct{}

func (this *UserGroupApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := models.GetLoggedInUser(r)

	userGroups := models.GetUserGroups(user)

	res, err := json.Marshal(userGroups)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}

func (this *UserGroupApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	user := models.GetLoggedInUser(r)

	userGroup := models.GetUserGroupById(id, user)

	res, err := json.Marshal(userGroup)
	corehelpers.PanicIf(err)

	fmt.Fprintf(w, "%s", res)
}
