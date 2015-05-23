package controllers

import (
	"fmt"
	"net/http"
	//"time"
	//"database/sql"
	_ "github.com/lib/pq"
	//"collexy/helpers"
	"collexy/core/modules/content/models"
	"strconv"
	//"log"
	//"github.com/gorilla/schema"
	//"encoding/json"
	"log"
	//"io/ioutil"
	//"path/filepath"
	corehelpers "collexy/core/helpers"
	"github.com/gorilla/mux"
	"html/template"
	"strings"
	//"github.com/gorilla/context"
	coreglobals "collexy/core/globals"
	coremodulemembermodels "collexy/core/modules/member/models"
	coremodulesettingsmodels "collexy/core/modules/settings/models"
	coremoduleuser "collexy/core/modules/user/models"
	"regexp"
)

type ContentController struct{}

// type TestData struct {
//     Data *TestStruct
//     HasUser bool
// }

// type TestStruct struct {
//     User *coremoduleuser.User
//     Content *models.Content
// }

// type TestDataMember struct {
//     Data *TestStructMember
//     HasMember bool
// }

// type TestStructMember struct {
//     Member *coremodulemembermodels.Member
//     Content *models.Content
// }

//var Templates map[string]*template.Template

func (this *ContentController) RenderTemplate(w http.ResponseWriter, name string, content *models.Content, member *coremodulemembermodels.Member) error {
	defer fmt.Println("RenderTemplate finished!")
	// Ensure the template exists in the map.

	tmpl, ok := coreglobals.Templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	// fmt.Print(coreglobals.Templates)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if member == nil {
		fmt.Println("controller.content.RenderTemplate(): user is nil")
		//tmpl.ExecuteTemplate(w, "base", content)
		test := &TestStructMember{nil, content}
		err := tmpl.ExecuteTemplate(w, "base", TestDataMember{test, false})
		corehelpers.PanicIf(err)

	} else {
		fmt.Println("controller.content.RenderTemplate(): username is: " + member.Username)
		test := &TestStructMember{member, content}
		fmt.Println("is this working? username: " + test.Member.Username)
		if err := tmpl.ExecuteTemplate(w, "base", TestDataMember{test, true}); err == nil {
			fmt.Println("member & data structs has been passed on to the template")
		} else {
			// handle error
			log.Println("Error in controllers.content.RenderTemplate(): " + err.Error())
		}
	}

	return nil
}

func (this *ContentController) RenderAdminTemplate(w http.ResponseWriter, name string, content *models.Content, user *coremoduleuser.User) error {
	fmt.Println("RenderAdminTemplate")
	// Ensure the template exists in the map.
	tmpl, ok := coreglobals.Templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}
	// fmt.Print(coreglobals.Templates)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if user == nil {
		fmt.Println("controller.content.RenderTemplate(): user is nil")
		//tmpl.ExecuteTemplate(w, "base", content)
		test := &TestStruct{nil, content}
		tmpl.ExecuteTemplate(w, "base", TestData{test, false})
	} else {
		fmt.Println("controller.content.RenderTemplate(): username is: " + user.Username)
		test := &TestStruct{user, content}
		fmt.Println("is this working? username: " + test.User.Username)
		if err := tmpl.ExecuteTemplate(w, "base", TestData{test, true}); err == nil {
			fmt.Println("user & data structs has been passed on to the template")
		} else {
			// handle error
			log.Println("Error in controllers.content.RenderTemplate(): " + err.Error())
		}
	}

	return nil
}

func (this *ContentController) RenderContent(w http.ResponseWriter, r *http.Request) {
	defer corehelpers.Un(corehelpers.Trace("SOME_ARBITRARY_STRING_SO_YOU_CAN_KEEP_TRACK"))
	sid := corehelpers.CheckMemberCookie(w, r)
	m, _ := coremodulemembermodels.GetMember(sid)

	coremodulemembermodels.SetLoggedInMember(r, m)

	fmt.Println("RENDERCONTENT")

	// idStr := r.URL.Query().Get(":nodeId")

	params := mux.Vars(r)
	// idStr := params["nodeId"]

	// nodeId, _ := strconv.Atoi(idStr)
	// fmt.Println(nodeId)

	// content := models.GetFrontendContentById(nodeId)

	url := params["url"]

	if strings.HasPrefix(url, "admin/") {
		//fmt.Println("lol")
	} else {
		fmt.Println(url)
		s := strings.Split(url, "/")
		name := strings.Replace(strings.ToLower(s[len(s)-1]), "-", " ", -1)

		var content *models.Content
		if name == "" {
			content = models.GetFrontendContentByDomain(r.Host)
		} else {
			content = models.GetFrontendContentByUrl(name, r.Host+r.URL.String())
		}

		if content == nil {
			fmt.Println("content is null!!")
			//First, check if 404 node id has been set in config
			if coreglobals.Conf.Id404 == -1 {
				//turn off 404 functionality
			} else if coreglobals.Conf.Id404 >= 0 {
				if coreglobals.Conf.Id404 == 0 {
					// use custom handler (a little slower, but works for multiple domains/sites)
					// //GET home page by domain from http request
					// urlSlice := strings.Split(url, "/")
					// homeContent := GetHomeContentByDomainAlternate(urlSlice[0])
					// // if home has set a 404 page
					// if(homeContent.Meta.id_404){
					//     // get content by id
					//     content = GetContentById(homeContent.Meta.404_node_id)
					//     // parse templates
					//     // render template
					// }
				} else {
					// use id_404 from /confic/config.json
					content = models.GetFrontendContentById(coreglobals.Conf.Id404)
				}
			}

			// coreglobals.Templates["404.tmpl"] = template.Must(template.ParseFiles("views/404.tmpl"))
			// // coreglobals.Templates["404.tmpl"] = template.Must(template.ParseFiles("views/Layout.tmpl","views/404.tmpl"))
			// this.RenderTemplate(w, "404.tmpl", content, nil)

			var templateName string = content.Template.Name + ".tmpl"
			fmt.Println("Content node name: " + templateName)
			//templateName := strings.Replace(content["template_name"].(string), " ", "-", -1) + ".tmpl"
			if templateName != ".tmpl" {
				if coreglobals.Templates[templateName] != nil {

				} else {

					templateArray := []string{"views/" + templateName}
					if content.Template.ParentTemplates != nil {

						tplFile := coremodulesettingsmodels.GetFilesystemNodeById("./views", templateName)
						rp1, _ := regexp.Compile("template \".*\"")

						lol := rp1.FindAllString(tplFile.Contents, -1) // ["abc", "def"]
						// fmt.Println(tplFile.Contents)
						//fmt.Println("regex return")
						//fmt.Println(lol)

						for _, value := range lol {
							concatStr := "views/" + value + ".tmpl"
							concatStr = strings.Replace(concatStr, "\"", "", -1)
							concatStr = strings.Replace(concatStr, "template ", "", -1)
							templateArray = append(templateArray, concatStr)
						}

						fmt.Println(templateArray)
						if content.Template.ParentTemplates != nil {

							parentTemplateNodes := content.Template.ParentTemplates

							v := make([]string, 0, len(parentTemplateNodes))

							for _, value := range parentTemplateNodes {
								tplName := "views/" + value.Name + ".tmpl"
								v = append(v, tplName)
							}
							templateArray = append(templateArray, v...)

							for _, value := range templateArray {
								if value != "views/"+templateName {
									dirViewSlice := strings.Split(value, "/")

									tplFile := coremodulesettingsmodels.GetFilesystemNodeById(dirViewSlice[0]+"/", dirViewSlice[1])
									rp1, _ := regexp.Compile("template \".*\"")

									lol := rp1.FindAllString(tplFile.Contents, -1) // ["abc", "def"]
									// fmt.Println(tplFile.Contents)
									//fmt.Println("regex return")
									//fmt.Println(lol)

									for _, value1 := range lol {
										//
										concatStr := strings.Replace(value1, "\"", "", -1)
										concatStr = strings.Replace(concatStr, "template ", "", -1)
										//character := dirViewSlice[1]

										if corehelpers.IsFirstNumber(concatStr) {
											fmt.Println("character is numric")
										} else {
											if corehelpers.IsFirstUpper(concatStr) {
												//fmt.Println("upper case true")

												concatStr = "views/" + concatStr + ".tmpl"
												if !corehelpers.StringInSlice(concatStr, templateArray) {
													templateArray = append(templateArray, concatStr)
												}
											}
										}
										//

									}
								}
							}
						}

						coreglobals.Templates[templateName] = template.Must(template.ParseFiles(templateArray...))
						//.Delims("{@","@}")

					} else {
						coreglobals.Templates[templateName] = template.Must(template.ParseFiles("views/" + templateName))
					}

				}
				this.RenderTemplate(w, templateName, content, nil)

			}

		} else {
			var templateName string = content.Template.Name + ".tmpl"
			//templateName := strings.Replace(content["template_name"].(string), " ", "-", -1) + ".tmpl"
			if templateName != ".tmpl" {
				if coreglobals.Templates[templateName] != nil {

				} else {
					if content.Template.ParentTemplates != nil {
						templateArray := []string{"views/" + templateName}

						tplFile := coremodulesettingsmodels.GetFilesystemNodeById("./views", templateName)
						rp1, _ := regexp.Compile("template \".*\"")

						lol := rp1.FindAllString(tplFile.Contents, -1) // ["abc", "def"]
						// fmt.Println(tplFile.Contents)
						//fmt.Println("regex return")
						//fmt.Println(lol)

						for _, value := range lol {
							concatStr := "views/" + value + ".tmpl"
							concatStr = strings.Replace(concatStr, "\"", "", -1)
							concatStr = strings.Replace(concatStr, "template ", "", -1)
							templateArray = append(templateArray, concatStr)
						}

						// fin, err := models.GetFilesystemNodes("./views")
						// corehelpers.PanicIf(err)
						// //fmt.Println(fin)
						// for _, value := range fin.Children {
						//     //fmt.Println(value.Info.Name)
						//     if(value.Info.Name != templateName && value.Info.Name != "Layout.tmpl"){
						//         tplName := "views/" + value.Info.Name
						//         templateArray = append(templateArray, tplName)
						//     }
						// }
						// for _, value := range templateArray {
						//     fmt.Println(value)
						// }

						if content.Template.ParentTemplates != nil {

							parentTemplateNodes := content.Template.ParentTemplates

							v := make([]string, 0, len(parentTemplateNodes))

							for _, value := range parentTemplateNodes {
								tplName := "views/" + value.Name + ".tmpl"
								v = append(v, tplName)
							}
							templateArray = append(templateArray, v...)

							for _, value := range templateArray {
								if value != "views/"+templateName {
									dirViewSlice := strings.Split(value, "/")

									tplFile := coremodulesettingsmodels.GetFilesystemNodeById(dirViewSlice[0]+"/", dirViewSlice[1])
									rp1, _ := regexp.Compile("template \".*\"")

									lol := rp1.FindAllString(tplFile.Contents, -1) // ["abc", "def"]
									// fmt.Println(tplFile.Contents)
									//fmt.Println("regex return")
									//fmt.Println(lol)

									for _, value1 := range lol {
										//
										concatStr := strings.Replace(value1, "\"", "", -1)
										concatStr = strings.Replace(concatStr, "template ", "", -1)
										//character := dirViewSlice[1]

										if corehelpers.IsFirstNumber(concatStr) {
											fmt.Println("character is numric")
										} else {
											if corehelpers.IsFirstUpper(concatStr) {
												//fmt.Println("upper case true")

												concatStr = "views/" + concatStr + ".tmpl"
												if !corehelpers.StringInSlice(concatStr, templateArray) {
													templateArray = append(templateArray, concatStr)
												}
											}
										}
										//

									}
								}
							}
						}

						// if(content.Template.PartialTemplateNodes != nil){
						//     partialTemplateNodes := content.Template.PartialTemplateNodes

						//     x := make([]string, 0, len(partialTemplateNodes))

						//     for  _, value := range partialTemplateNodes {
						//         tplName := "views/" + value.Name + ".tmpl"
						//         x = append(x, tplName)
						//     }
						//     templateArray = append(templateArray, x...)
						// }

						// templateArray = append(templateArray, "views/Top Navigation.tmpl")
						// templateArray = append(templateArray, "views/About Widget.tmpl")
						// templateArray = append(templateArray, "views/Social.tmpl")

						// if(content.Template.ParentTemplates != nil){
						//     for _, value := range content.Template.ParentTemplates {
						//         for _, pn := range value.
						//     }
						// }

						coreglobals.Templates[templateName] = template.Must(template.ParseFiles(templateArray...))
						//.Delims("{@","@}")

					} else {
						coreglobals.Templates[templateName] = template.Must(template.ParseFiles("views/" + templateName))
					}

				}

				// if user := models.GetLoggedInUser(r); user != nil {
				//     this.RenderTemplate(w, templateName, content, user)
				// } else {
				//     this.RenderTemplate(w, templateName, content, nil)
				// }




				if member := coremodulemembermodels.GetLoggedInMember(r); member != nil {
					if content.PublicAccessMembers != nil || content.PublicAccessMemberGroups != nil {
						memberIdstr := strconv.Itoa(member.Id)
						if content.PublicAccessMembers[memberIdstr] != nil {
							this.RenderTemplate(w, templateName, content, member)
						} else if member.Groups2PublicAccess(content.PublicAccessMemberGroups) {
							this.RenderTemplate(w, templateName, content, member)
						} else {
							fmt.Println("You do not seem to have access to this content")
							coreglobals.Templates["Unauthorized.tmpl"] = template.Must(template.ParseFiles("views/Layout.tmpl", "views/Unauthorized.tmpl"))
							this.RenderTemplate(w, "Unauthorized.tmpl", nil, nil)
						}
						// if corehelpers.IntInSlice(member.Id, content.PublicAccess.Members) {
						// 	this.RenderTemplate(w, templateName, content, member)
						// } else if member.Groups2PublicAccess(content.PublicAccess.Groups) {
						// 	this.RenderTemplate(w, templateName, content, member)
						// } else {
						// 	fmt.Println("You do not seem to have access to this content")
						// 	coreglobals.Templates["Unauthorized.tmpl"] = template.Must(template.ParseFiles("views/Layout.tmpl", "views/Unauthorized.tmpl"))
						// 	this.RenderTemplate(w, "Unauthorized.tmpl", nil, nil)
						// }
					} else {
						this.RenderTemplate(w, templateName, content, member)
					}
				} else {
					if content.PublicAccessMembers != nil || content.PublicAccessMemberGroups != nil {
						fmt.Println("You need to be a member and be logged in to have access to this content")
						templateName := "Unauthorized.tmpl"









				// if member := coremodulemembermodels.GetLoggedInMember(r); member != nil {
				// 	if content.PublicAccess != nil {
				// 		if corehelpers.IntInSlice(member.Id, content.PublicAccess.Members) {
				// 			this.RenderTemplate(w, templateName, content, member)
				// 		} else if member.Groups2PublicAccess(content.PublicAccess.Groups) {
				// 			this.RenderTemplate(w, templateName, content, member)
				// 		} else {
				// 			fmt.Println("You do not seem to have access to this content")
				// 			coreglobals.Templates["Unauthorized.tmpl"] = template.Must(template.ParseFiles("views/Layout.tmpl", "views/Unauthorized.tmpl"))
				// 			this.RenderTemplate(w, "Unauthorized.tmpl", nil, nil)
				// 		}
				// 	} else {
				// 		this.RenderTemplate(w, templateName, content, member)
				// 	}
				// } else {
				// 	if content.PublicAccess != nil {
				// 		fmt.Println("You need to be a member and be logged in to have access to this content")
				// 		templateName := "Unauthorized.tmpl"









						// templateArray := []string{"views/" + templateName}
						// templateArray = append(templateArray, "views/Layout.tmpl")
						// tempTemplateArray := templateArray

						// for _, value := range tempTemplateArray {
						//     //if value != "views/" + templateName {
						//         tplArrSlice := strings.Split(value,"/")
						//         fmt.Println(tplArrSlice)
						//         tplFile := models.GetFilesystemNodeById("./views",tplArrSlice[1])
						//         rp1, _ := regexp.Compile("template \".*\"")

						//         lol := rp1.FindAllString(tplFile.Contents, -1) // ["abc", "def"]
						//         // fmt.Println(tplFile.Contents)
						//         //fmt.Println("regex return")
						//         //fmt.Println(lol)

						//         for _, value1 := range lol {
						//             //if value1 != "views/" + templateName {
						//                 fmt.Println("LOL: " + value1)

						//                 concatStr := strings.Replace(value1, "\"", "", -1)
						//                 concatStr = strings.Replace(concatStr, "template ", "", -1)
						//                 fmt.Println(concatStr)
						//                 if (corehelpers.IsFirstNumber(concatStr)){
						//                     fmt.Println("character is numric")
						//                 }else{
						//                     if (corehelpers.IsFirstUpper(concatStr)) {
						//                         fmt.Println("upper case true")

						//                         concatStr = "views/" + concatStr + ".tmpl"
						//                         if(!corehelpers.StringInSlice(concatStr, templateArray)){
						//                             templateArray = append(templateArray, concatStr)
						//                         }
						//                     }
						//                 }
						//             //}
						//         }
						//     //}
						// }

						//coreglobals.Templates["Unauthorized.tmpl"] = template.Must(template.ParseFiles("views/Layout.tmpl","views/Unauthorized.tmpl"))
						coreglobals.Templates["Unauthorized.tmpl"] = template.Must(template.ParseFiles("views/Unauthorized.tmpl", "views/Login Widget.tmpl"))
						this.RenderTemplate(w, templateName, nil, nil)
					} else {
						this.RenderTemplate(w, templateName, content, nil)
					}

				}

				//this.RenderTemplate(w, templateName, &content, nil)
			}
		}
	}

}
