package controllers

import (
    "fmt"
    "net/http"
    _ "github.com/lib/pq"
    "collexy/helpers"
    coreglobals "collexy/core/globals"
    "collexy/core/api/models"
    "strconv"
    "github.com/gorilla/schema"
    "github.com/gorilla/mux"
    "encoding/json"
    //"strings"
)

type NodeApiController struct {}

// func (this *NodeApiController) GetTest(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     if user := models.GetLoggedInUser(r); user != nil {
//         nodes := models.TempGetUserAllowedNodes(user)

//         res, err := json.Marshal(nodes)
//         helpers.PanicIf(err)

//         fmt.Fprintf(w,"%s",res)
//     } else {
//         fmt.Println("GetTest ERROR")
//     }
// }

func (this *NodeApiController) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    queryStrParams := r.URL.Query()

    user := models.GetLoggedInUser(r)

    nodes := models.GetNodes(queryStrParams,user)

    res, err := json.Marshal(nodes)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

func (this *NodeApiController) GetById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // idStr := r.URL.Query().Get(":id")
    // id, _ := strconv.Atoi(idStr)
    params := mux.Vars(r)
    idStr := params["id"]
    id, _ := strconv.Atoi(idStr)

    node := models.GetNodeById(id)

    res, err := json.Marshal(node)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

func (this *NodeApiController) GetByIdChildren(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    idStr := params["id"]
    id, _ := strconv.Atoi(idStr)

    nodes := models.GetNodeByIdChildren(id)

    res, err := json.Marshal(nodes)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

func (this *NodeApiController) Post(w http.ResponseWriter, r *http.Request) {
    path := r.FormValue("path")
    created_by,err := strconv.Atoi(r.FormValue("created_by"))
    name := r.FormValue("name")
    node_type,err := strconv.Atoi(r.FormValue("node_type"))

    db := coreglobals.Db

    // http://stackoverflow.com/questions/244243/how-to-reset-postgres-primary-key-sequence-when-it-falls-out-of-sync
    querystr := fmt.Sprintf("INSERT INTO node (path, created_by, name, node_type) VALUES ('%s', %d, '%s', %d)", path, created_by, name, node_type)
    res, err := db.Exec(querystr)
    helpers.PanicIf(err)
    fmt.Println(res)
}

func (this *NodeApiController) Put(w http.ResponseWriter, r *http.Request) {

    params := mux.Vars(r)
    idStr := params["id"]
    id, _ := strconv.Atoi(idStr)

    parm_id := id

    t := new(models.Node)

    err := r.ParseForm()

    helpers.PanicIf(err)

    decoder := schema.NewDecoder()
    // r.PostForm is a map of our POST form values
    decoder.Decode(t, r.PostForm)

    db := coreglobals.Db
    // http://stackoverflow.com/questions/244243/how-to-reset-postgres-primary-key-sequence-when-it-falls-out-of-sync
    fmt.Println(fmt.Sprintf("path: %s, created_by: %d, name: %s, node type: %d", t.Path, t.CreatedBy, t.Name, t.NodeType))


    querystr := fmt.Sprintf("UPDATE node SET (path, created_by, name, node_type) = ('%s', %d, '%s', %d) WHERE id=%d", t.Path, t.CreatedBy, t.Name, t.NodeType, parm_id)
    res, err := db.Exec(querystr)
    helpers.PanicIf(err)
    fmt.Println(res)
}

func (this *NodeApiController) Delete(w http.ResponseWriter, r *http.Request) {

    templol := r.URL.Query().Get(":id")
    rofl,err1 := strconv.Atoi(templol)
    helpers.PanicIf(err1)

    parm_id := rofl

   
    db := coreglobals.Db
    
    querystr := fmt.Sprintf("DELETE FROM node WHERE id=%d", parm_id)
    _, err := db.Exec(querystr)
    helpers.PanicIf(err)
}