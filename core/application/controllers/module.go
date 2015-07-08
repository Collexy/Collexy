package controllers

import (
	"collexy/core/application/models/xml_models"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	coremodulesettingsmodelsxmlmodels "collexy/core/modules/settings/models/xml_models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//coreglobals "collexy/core/globals"
)

type ModuleApiController struct {
	Module xml_models.Module
}

func (this *ModuleApiController) ModuleHandler(w http.ResponseWriter, r *http.Request) {
	htmlStr := `<html >
                <head>
                    <title>Collexy Installation</title>
                </head>
                <body>
                    <div>
                        <h1 style="text-align:center;">Collexy Logo</h1>
                        <p>Below you should select your desired module to install.</p>
                        <form method="POST" action="?step=2">
                            <table>
                                <tr>
                                    <td><strong>Modules</strong></td>
                                    <td><label><input type="radio" name="collexy_module" value="TXT Starter Kit"/> TXT Starter Kit</label></td>
                                </tr>
                                <tr>
                                    <td><input type="submit" value="Submit"></td>
                                    <td></td>
                                </tr>
                            </table>
                            <input type="hidden" name="step" value="2"/>
                        </form>
                    </div>
                </body>
            </html>`
	fmt.Fprintf(w, htmlStr)
}

func (this *ModuleApiController) ModulePostHandler(w http.ResponseWriter, r *http.Request) {
	//str := r.PostFormValue("collexy_module")
	if _, err := os.Stat("./_temporary_module_library/" + r.PostFormValue("collexy_module") + "/module.xml"); err != nil {
		if os.IsNotExist(err) {
			log.Println("XML file does not exist")
			fmt.Fprintf(w, "XML file does not exist")
		} else {
			// other error
		}
	} else {
		err5 := this.ImportModuleFromXML("./_temporary_module_library/" + r.PostFormValue("collexy_module"))
		if err5 != nil {
			log.Println("ERROR INSTALLING MODULE SCRIPT")
			log.Fatal(err5)
		} else {
			log.Println("MODULE SCRIPT INSTALLED SUCCESSFULLY")
			//adminHandler(w, r)
		}
	}
	fmt.Fprintf(w, "testPostHandler")
}

func (this *ModuleApiController) ImportModuleFromXML(moduleDirectoryPath string) (err error) {
	moduleXMLFile, err1 := os.Open(moduleDirectoryPath + "/module.xml")
	defer moduleXMLFile.Close()
	if err1 != nil {
		log.Println("Error opening module.xml file")
	}

	XMLdata, err2 := ioutil.ReadAll(moduleXMLFile) // use bufio intead since the xml can scale big

	if err2 != nil {
		log.Println("Error reading from module.xml file")
		fmt.Printf("error: %v", err2)
	}

	err3 := xml.Unmarshal(XMLdata, &this.Module)
	if err3 != nil {
		fmt.Printf("error: %v", err3)
	}

	this.ExportModuleToDatabase()
	return
}

func (this *ModuleApiController) ExportModuleToDatabase() {
	// this.Module.DataTypes.Post()
	// make sure Post() is completed before continuing
	//dataTypes := coremodulesettingsmodels.DataType.Get()
	dataTypes := coremodulesettingsmodelsxmlmodels.GetDataTypes()
	// make sure Get() is completed before continuing

	for _, t := range this.Module.Templates {
		// params: parent
		// t.Post(nil,existingTemplatesFromDb)
		t.Post(nil)
	}

	var flatTemplatesSlice []*coremodulesettingsmodelsxmlmodels.Template

	coremodulesettingsmodelsxmlmodels.Walk(this.Module.Templates, func(t *coremodulesettingsmodelsxmlmodels.Template) bool {
		flatTemplatesSlice = append(flatTemplatesSlice, t)
		return true
	})

	templates := coremodulesettingsmodels.GetTemplates(nil)

	//templates = coremodulesettingsmodels.Templates.Get()

	for _, ct := range this.Module.ContentTypes {
		// params: parent, parentContentTypes, dataTypes, templates
		ct.Post(nil, nil, flatTemplatesSlice, dataTypes)

	}
	contentTypes := coremodulesettingsmodels.GetContentTypes(nil)

	// fmt.Printf("this.Module.ContentItems length is %d\n", len(this.Module.ContentItems))
	for _, c := range this.Module.ContentItems {
		fmt.Println(c.Name)
		c.Post(nil, contentTypes, templates)
	}

	for _, mt := range this.Module.MediaTypes {
		// params: parent, parentContentTypes, dataTypes, templates
		mt.Post(nil, nil, dataTypes)

	}

	mediaTypes := coremodulesettingsmodels.GetMediaTypes(nil)

	for _, mt := range this.Module.MimeTypes {
		// params: parent, parentContentTypes, dataTypes, templates
		mt.Post(mediaTypes)

	}

	fmt.Printf("this.Module.MediaItems length: %d\n", len(this.Module.MediaItems))

	for _, m := range this.Module.MediaItems {
		// params: parent, parentContentTypes, dataTypes, templates
		m.Post(nil, mediaTypes)

	}

	for _, f := range this.Module.Files {
		// params: parent, parentContentTypes, dataTypes, templates
		f.Copy()

	}

	// import assets
}
