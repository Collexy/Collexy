// Package directory_tree provides a way to generate a directory tree.
//
// Example usage:
//
//	tree, err := directory_tree.NewTree("/home/me")
//
// I did my best to keep it OS-independent but truth be told I only tested it
// on OS X and Debian Linux so mileage may vary. You've been warned.
package models

import (
	corehelpers "collexy/core/helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// FileInfo is a struct created from os.FileInfo interface for serialization.
type FileInfo struct {
	Name    string      `json:"name,omitempty"`
	Size    int64       `json:"size,omitempty"`
	Mode    os.FileMode `json:"mode,omitempty"`
	ModTime time.Time   `json:"mod_time,omitempty"`
	IsDir   bool        `json:"is_dir,omitempty"`
}

// Helper function to create a local FileInfo struct from os.FileInfo interface.
func fileInfoFromInterface(v os.FileInfo) *FileInfo {
	return &FileInfo{v.Name(), v.Size(), v.Mode(), v.ModTime(), v.IsDir()}
}

// FileNode represents a node in a directory tree.
type FileNode struct {
	FullPath string      `json:"path,omitempty"`
	OldPath  string      `json:"old_path,omitempty"`
	Info     *FileInfo   `json:"info,omitempty"`
	Children []*FileNode `json:"children,omitempty"`
	Contents string      `json:"contents,omitempty"`
	Show     bool        `json:"show,omitempty"`
	Parent   string      `json:"parent,omitempty"`
}

// Helper function to get a path's parent path (OS-specific).
func getParentPath(path string) string {
	els := strings.Split(path, string(os.PathSeparator))
	return strings.Join(els[:len(els)-1], string(os.PathSeparator))
}

// Create directory hierarchy.
func NewTree(root string) (result *FileNode, err error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return
	}
	parents := make(map[string]*FileNode)
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		parents[path] = &FileNode{path, "", fileInfoFromInterface(info), []*FileNode{}, "", true, ""}
		return nil
	}
	if err = filepath.Walk(absRoot, walkFunc); err != nil {
		return
	}
	for path, node := range parents {
		parentPath := getParentPath(path)
		parent, exists := parents[parentPath]
		if !exists { // If a parent does not exist, this is the root.
			result = node
		} else {
			parent.Children = append(parent.Children, node)
		}
	}
	return
}

/* NEW STUFF */

func (t *FileNode) Post() (err error) {
	fmt.Println("t.parent:::: " + t.Parent + "\n")
	fmt.Println("t.path:::: " + t.FullPath + "\n")
	fmt.Println("t.info.name:::: " + t.Info.Name + "\n")
	fmt.Println(t.Info.IsDir)
	fmt.Println("rgr \n")
	tm, err := json.Marshal(t)
	corehelpers.PanicIf(err)
	fmt.Println("tm:::: ")
	fmt.Println(string(tm))
	//db := corehelpers.Db

	//tplNodeName := t.Info.Name + ".tmpl"
	absPath := t.FullPath

	if t.Info.IsDir {
		// create directory 0777 permission too liberal?
		fmt.Println("creating directory: " + t.Info.Name + "with path: " + t.FullPath)
		err = os.Mkdir(absPath, 0644)

	} else {
		fmt.Println("creating file...")
		// write whole the body - maybe use bufio/os/io packages for buffered read/write on big files
		err = ioutil.WriteFile(absPath, []byte(t.Contents), 0644)
	}
	return
}

func (t *FileNode) Update() {
	fmt.Println("::: FileNode Update Initiated :::")
	//tplNodeName := t.Info.Name + ".tmpl"

	// equivalent to Python's `if os.path.exists(filename)`
	if _, err := os.Stat(t.OldPath); err == nil {
		newAbsPath := t.FullPath
		if t.Info.IsDir {
			fmt.Println("DIRECTORY RENAME:::")
			// create directory 0777 permission too liberal?
			//os.Mkdir(newAbsPath,0644)
			err := os.Rename(t.OldPath, newAbsPath)

			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("old: " + t.OldPath + "new: ")
			fmt.Println(newAbsPath)

			//fmt.Println("FILE EXISTS ::: RENAMING FILE :::")
			//Add old/original file path attribute in FileNode
			err := os.Rename(t.OldPath, newAbsPath)

			if err != nil {
				fmt.Println(err)
				return
			}

			// write whole the body - maybe use bufio/os/io packages for buffered read/write on big files
			err = ioutil.WriteFile(newAbsPath, []byte(t.Contents), 0644)
			if err != nil {
				panic(err)
			}

		}
	}
}

func GetFilesystemNodes(rootdir string) (tree *FileNode, err error) {
	tree, err = NewTree(rootdir) // maybe try prepending with slash /
	return
}

func GetFilesystemNodes1(rootdir string) (tree *FileNode) {
	tree, _ = NewTree(rootdir) // maybe try prepending with slash /
	return
}

func GetFilesystemNodeById(rootdir, filename string) (fileNode FileNode) {
	filepath.Walk(rootdir, func(path string, fi os.FileInfo, err error) (e error) {
		//if !fi.IsDir() {
		if fi.Name() == filename {
			//fmt.Println("AWESOME... WORKS!!!")
			//fmt.Println(fi.Name())
			//fmt.Println("AWESOME... WORKS!!!")
			finfo := FileInfo{fi.Name(), fi.Size(), fi.Mode(), fi.ModTime(), fi.IsDir()}
			//finfoInterface := *FileInfo
			//f, _ := json.Marshal(finfo)
			if !fi.IsDir() {
				bytes, err1 := ioutil.ReadFile(path) // path is the path to the file.
				corehelpers.PanicIf(err1)

				fileNode = FileNode{path, path, &finfo, nil, string(bytes), true, ""}
				return
			} else {

				var children []*FileNode = nil

				// Channel c, is for getting the parent template
				// We need to append the id of the newly created template to the path of the parent id to create the new path
				c := make(chan *FileNode)

				var wg sync.WaitGroup

				wg.Add(1)

				go func() {
					defer wg.Done()
					c <- GetFilesystemNodes1(path)
				}()

				go func() {
					for i := range c {
						fmt.Println(i)
						children = append(children, i)
					}
				}()

				wg.Wait()

				fileNode = FileNode{path, path, &finfo, children[0].Children, "", true, ""}
				return
				// finod, _ := json.Marshal(fin)
				// fmt.Fprintf(w,"%s",finod)
			}
		}
		if err != nil {
			fmt.Println("Fail")
		}
		return
	})
	return
}

func DeleteFileSystemNode(path string) (err error) {

	err = os.RemoveAll(path)

	if err != nil {
		fmt.Println(err)

	}
	return
}
