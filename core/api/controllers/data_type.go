package controllers

import (
    "fmt"
    "net/http"
    corehelpers "collexy/core/helpers"
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
    corehelpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}

func (this *DataTypeApiController) GetById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    idStr := params["id"]

    id, _ := strconv.Atoi(idStr)

    dataType := models.GetDataTypeById(id)

    res, err := json.Marshal(dataType)
    corehelpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}

// func (this *DataTypeApiController) Post(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     dataType := models.DataType{}

//     err := json.NewDecoder(r.Body).Decode(&dataType)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     dataType.Post()

// }

// func (this *DataTypeApiController) Put(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     dataType := models.DataType{}

//     err := json.NewDecoder(r.Body).Decode(&dataType)

//     if err != nil {
//         http.Error(w, "Bad Request", 400)
//     }

//     dataType.Update()
// }