package controllers

import (
    "fmt"
    "net/http"
    "collexy/helpers"
    "collexy/core/api/models"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
)

type DataTypeApiController struct {}

func (this *DataTypeApiController) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    dataTypes := models.GetDataTypes()

    res, err := json.Marshal(dataTypes)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}

func (this *DataTypeApiController) GetByNodeId(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    idStr := params["nodeId"]

    nodeId, _ := strconv.Atoi(idStr)

    dataType := models.GetDataTypeByNodeId(nodeId)

    res, err := json.Marshal(dataType)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}

func (this *DataTypeApiController) Post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    dataType := models.DataType{}

    err := json.NewDecoder(r.Body).Decode(&dataType)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    dataType.Post()

}

func (this *DataTypeApiController) Put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    dataType := models.DataType{}

    err := json.NewDecoder(r.Body).Decode(&dataType)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    dataType.Update()
}