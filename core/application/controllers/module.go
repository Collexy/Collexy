package controllers

import
(
	"net/http"
	"fmt"
	"os"
	"log"
	"collexy/core/application/models/xml_models"
	"io/ioutil"
	"encoding/xml"
	//coremodulesettingsmodels "collexy/core/modules/settings/models"
	coremodulesettingsmodelsxmlmodels "collexy/core/modules/settings/models/xml_models"
	//coreglobals "collexy/core/globals"
)

type ModuleApiController struct{
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

func (this *ModuleApiController) ExportModuleToDatabase(){
	// this.Module.DataTypes.Post()
	// make sure Post() is completed before continuing
	//dataTypes := coremodulesettingsmodels.DataTypes.Get()
	// make sure Get() is completed before continuing

	//templates := coremodulesettingsmodels.Templates.Get()

	for _, t := range this.Module.Templates{
		// params: parent
		// t.Post(nil,existingTemplatesFromDb)
		t.Post(nil)
	}

	var flatTemplatesSlice []*coremodulesettingsmodelsxmlmodels.Template

	coremodulesettingsmodelsxmlmodels.Walk(this.Module.Templates, func(t *coremodulesettingsmodelsxmlmodels.Template) bool {
		flatTemplatesSlice = append(flatTemplatesSlice, t)
		return true
	})
	
	//templates = coremodulesettingsmodels.Templates.Get()
	
	for _, ct := range this.Module.ContentTypes{
		// params: parent, parentContentTypes, dataTypes, templates
		ct.Post(nil, nil, flatTemplatesSlice)
		
	}
	//contentTypes := this.Module.ContentType.Get()

	// this.Module.MediaTypes.Post(dataTypes)
	// mediaTypes := coremodulesettingsmodels.MediaTypes.Get()

	// this.Module.MimeTypes.Post(mediaTypes)
	
	// this.Module.Content.Post(contentTypes)

	// this.Module.Media.Post(mediaTypes)

	// import assets
}
