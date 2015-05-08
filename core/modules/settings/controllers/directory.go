package controllers

import (
	"fmt"
	"net/http"
	//"time"
	//"database/sql"
	corehelpers "collexy/core/helpers"
	"collexy/core/modules/settings/models"
	_ "github.com/lib/pq"
	//"strconv"
	//"log"
	//"github.com/gorilla/schema"
	"encoding/json"
	//"log"
	"io/ioutil"
	//"path/filepath"
	//"strings"
	//"html/fileNode"
	"github.com/gorilla/mux"
)

type DirectoryApiController struct{}

func (this *DirectoryApiController) UploadFileTest(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "multipart/mixed; boundary=frontier")
	fmt.Println("UPLOAD FILE TEST:::")

	queryStrParams := r.URL.Query();

	path := queryStrParams.Get("path")

	fmt.Println("path: " + path)
	
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(path + "\\" + handler.Filename, data, 0777)

	if err != nil {
		fmt.Println(err)
	}

}

func (this *DirectoryApiController) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fileNode := models.FileNode{}

	// var lol map[string]interface{}
	// json.NewDecoder(r.Body).Decode(&lol)
	// fmt.Println(lol)

	err := json.NewDecoder(r.Body).Decode(&fileNode)

	if err != nil {
		http.Error(w, "Bad Request1", 400)
	}

	fileNode.Post()

}

func (this *DirectoryApiController) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fileNode := models.FileNode{}

	err := json.NewDecoder(r.Body).Decode(&fileNode)

	if err != nil {
		http.Error(w, "Bad Request", 400)
	}

	fileNode.Update()

}

func (this *DirectoryApiController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	rootdir := params["rootdir"]

	tree, err := models.GetFilesystemNodes(rootdir)
	corehelpers.PanicIf(err)

	b, err := json.Marshal(tree)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Fprintf(w, "%s", b)

}

func (this *DirectoryApiController) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	rootdir := params["rootdir"]

	filename := params["name"]

	fileNode := models.GetFilesystemNodeById(rootdir, filename)

	finod, _ := json.Marshal(fileNode)
	fmt.Fprintf(w, "%s", finod)

}
