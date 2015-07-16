package xml_models

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	XMLName xml.Name `xml:"file"`
	Name    string   `xml:"name"`
	Path    string   `xml:"path"`
	// Id       *int     `xml:"id,omitempty"`
	// Path     string   `xml:path,omitempty`
	// ParentId *int     `xml:"parentId,omitempty"`
	// Name     string   `xml:"name"`
	// Alias    string   `xml:"alias"`
	// // Parent string			`xml:"parent"`
	// IsPartial bool        `xml:"isPartial"`
	// Content   string      `xml:"content"`
	// Children  []*Template `xml:"children>template,omitempty"`
}

func (this *File) Copy(module string) {
	// check if dir exists
	// if true
	// copy file
	// else
	// create dir
	// copy file
	absPath, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "\\" + this.Path)
	if _, err := os.Stat(absPath); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			log.Println("Directory does not exist")
			// err = os.MkdirAll(absPath, 0644)
			err = os.MkdirAll(absPath, 0755)
			log.Println("Directory created with path: " + absPath)
		} else {
			// other error

		}
	} else {

	}
	src, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "\\_temporary_module_library\\" + module + "\\" + this.Path + "\\" + this.Name)
	dst := absPath + "\\" + this.Name
	CopyFile(src, dst)
	// // module path = reader path
	// // temporary hardcode it
	// myAbsPath, _ := filepath.Abs(filepath.Dir(os.Args[0]) + "\\_temporary_module_library\\TXT Starter Kit\\" + this.Path)
	// fmt.Println(myAbsPath)

	// //myAbsPath := "" + absPath2
	// if reader, err1 := os.Open(myAbsPath); err1 == nil {
	// 	defer reader.Close()

	// 	// http://golang.org/pkg/os/#FileInfo
	// 	statinfo, err2 := reader.Stat()

	// 	if err2 != nil {
	// 		fmt.Println(err2)
	// 		//return nil
	// 	}

	// 	fmt.Println()
	// 	fmt.Println(statinfo.Size())

	// 	// Directory exists and is writable
	// 	err3 := CopyFile(reader, absPath+"\\"+this.Name)
	// 	if err3 != nil {
	// 		fmt.Println(err3)
	// 		//return nil
	// 	}
	// } else {
	// 	log.Println(err1)
	// }
}

func CopyFile(src, dst string) {
	// open files r and w
	r, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// do the actual work
	n, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Copied %v bytes\n", n)
}

// func CopyFile(in io.Reader, dst string) (err error) {

// 	// Does file already exist? Skip
// 	if _, err := os.Stat(dst); err == nil {
// 		return nil
// 	}

// 	err = nil

// 	out, err := os.Create(dst)
// 	if err != nil {
// 		fmt.Println("Error creating file", err)
// 		return
// 	}

// 	defer func() {
// 		cerr := out.Close()
// 		if err == nil {
// 			err = cerr
// 		}
// 	}()

// 	var bytes int64
// 	if bytes, err = io.Copy(out, in); err != nil {
// 		fmt.Println("io.Copy error")
// 		return
// 	}
// 	fmt.Println(bytes)

// 	err = out.Sync()
// 	return
// }
