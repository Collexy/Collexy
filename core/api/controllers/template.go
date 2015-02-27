package controllers

import (
    "fmt"
    "net/http"
    //"time"
    //"database/sql"
    _ "github.com/lib/pq"
    "collexy/helpers"
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

type TemplateApiController struct {}

func (this *TemplateApiController) PostTemplate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    template := models.Template{}

    err := json.NewDecoder(r.Body).Decode(&template)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    template.Post()

}


func (this *TemplateApiController) PutTemplate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    template := models.Template{}

    err := json.NewDecoder(r.Body).Decode(&template)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    template.Update()

}

func (this *TemplateApiController) GetTemplates(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
fmt.Println("GETtemplates")
    templates := models.GetTemplates()
    
    res, err := json.Marshal(templates)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}

func (this *TemplateApiController) GetTemplateByNodeId(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    idStr := params["nodeId"]

    nodeId, _ := strconv.Atoi(idStr)

    template := models.GetTemplateByNodeId(nodeId)

    res, err := json.Marshal(template)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}