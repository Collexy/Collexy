package controllers

import (
    "fmt"
    "net/http"
    //"time"
    //"database/sql"
    _ "github.com/lib/pq"
    //"collexy/helpers"
    "collexy/core/api/models"
    //"strconv"
    //"log"
    //"github.com/gorilla/schema"
    //"encoding/json"
    "log"
    //"io/ioutil"
    //"path/filepath"
    "strings"
    "html/template"
    "github.com/gorilla/mux"
    applicationglobals "collexy/core/application/globals"
    corehelpers "collexy/core/helpers"
    //"github.com/gorilla/context"
    
)

type ContentController struct {}

type TestData struct {
    Data *TestStruct
    HasUser bool
}

type TestStruct struct {
    User *models.User
    Content *models.Content
}

type TestDataMember struct {
    Data *TestStructMember
    HasMember bool
}

type TestStructMember struct {
    Member *models.Member
    Content *models.Content
}

//var Templates map[string]*template.Template

func (this *ContentController) RenderTemplate(w http.ResponseWriter, name string, content *models.Content, member *models.Member) error {
    // Ensure the template exists in the map.
    tmpl, ok := applicationglobals.Templates[name]
    if !ok {
        return fmt.Errorf("The template %s does not exist.", name)
    }
    fmt.Print(applicationglobals.Templates)
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    if(member == nil){
        fmt.Println("controller.content.RenderTemplate(): user is nil")
        //tmpl.ExecuteTemplate(w, "base", content)
        test := &TestStructMember{nil, content}
        tmpl.ExecuteTemplate(w, "base", TestDataMember{test, false})
    } else {
        fmt.Println("controller.content.RenderTemplate(): username is: " + member.Username)
        test := &TestStructMember{member, content}
        fmt.Println("is this working? username: " + test.Member.Username)
        if err := tmpl.ExecuteTemplate(w, "base", TestDataMember{test, true}); err == nil{
            fmt.Println("member & data structs has been passed on to the template")
        } else{
            // handle error
            log.Println("Error in controllers.content.RenderTemplate(): " + err.Error())
        }
    }

    return nil
}


func (this *ContentController) RenderAdminTemplate(w http.ResponseWriter, name string, content *models.Content, user *models.User) error {
    // Ensure the template exists in the map.
    tmpl, ok := applicationglobals.Templates[name]
    if !ok {
        return fmt.Errorf("The template %s does not exist.", name)
    }
    fmt.Print(applicationglobals.Templates)
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    if(user == nil){
        fmt.Println("controller.content.RenderTemplate(): user is nil")
        //tmpl.ExecuteTemplate(w, "base", content)
        test := &TestStruct{nil, content}
        tmpl.ExecuteTemplate(w, "base", TestData{test, false})
    } else {
        fmt.Println("controller.content.RenderTemplate(): username is: " + user.Username)
        test := &TestStruct{user, content}
        fmt.Println("is this working? username: " + test.User.Username)
        if err := tmpl.ExecuteTemplate(w, "base", TestData{test, true}); err == nil{
            fmt.Println("user & data structs has been passed on to the template")
        } else{
            // handle error
            log.Println("Error in controllers.content.RenderTemplate(): " + err.Error())
        }
    }

    return nil
}


func (this *ContentController) RenderContent(w http.ResponseWriter, r *http.Request) {

    sid := corehelpers.CheckMemberCookie(w,r)
    m, _ := models.GetMember(sid)

    models.SetLoggedInMember(r,m)

    fmt.Println("RENDERCONTENT")

    // idStr := r.URL.Query().Get(":nodeId")

    params := mux.Vars(r)
    // idStr := params["nodeId"]

    // nodeId, _ := strconv.Atoi(idStr)
    // fmt.Println(nodeId)

    // content := models.GetFrontendContentByNodeId(nodeId)

    url := params["url"]
    s := strings.Split(url, "/")
    name := strings.Replace(strings.ToLower(s[len(s)-1]), "-", " ",-1)
    var content *models.Content
    if(name==""){
        content = models.GetFrontendContentByDomain(r.Host)
    } else {
        content = models.GetFrontendContentByUrl(name, r.Host + r.URL.String())
    }
    


    if(content == nil){
        fmt.Println("content is null!!")
        applicationglobals.Templates["404.tmpl"] = template.Must(template.ParseFiles("views/Layout.tmpl","views/404.tmpl"))
        this.RenderTemplate(w, "404.tmpl", nil, nil)
    } else{
        var templateName string = content.Template.Node.Name + ".tmpl"
        //templateName := strings.Replace(content["template_name"].(string), " ", "-", -1) + ".tmpl"
        if(templateName !=".tmpl"){
            if(applicationglobals.Templates[templateName] != nil){

            } else{        
                if(content.Template.Node.ParentNodes != nil || content.Template.PartialTemplates != nil){
                    templateArray := []string{"views/" + templateName}

                    if(content.Template.Node.ParentNodes != nil){
      
                        parentTemplateNodes := content.Template.Node.ParentNodes

                        v := make([]string, 0, len(parentTemplateNodes))

                        for  _, value := range parentTemplateNodes {
                            tplName := "views/" + value.Name + ".tmpl"
                            v = append(v, tplName)
                        }
                        templateArray = append(templateArray, v...)
                        
                    }

                    if(content.Template.PartialTemplateNodes != nil){
                        partialTemplateNodes := content.Template.PartialTemplateNodes
                        
                        x := make([]string, 0, len(partialTemplateNodes))

                        for  _, value := range partialTemplateNodes {
                            tplName := "views/" + value.Name + ".tmpl"
                            x = append(x, tplName)
                        }
                        templateArray = append(templateArray, x...)
                    }

                    applicationglobals.Templates[templateName] = template.Must(template.ParseFiles(templateArray...))
                    //.Delims("{@","@}")
                } else {
                    applicationglobals.Templates[templateName] = template.Must(template.ParseFiles("views/" + templateName))
                }
            }

            // if user := models.GetLoggedInUser(r); user != nil {
            //     this.RenderTemplate(w, templateName, content, user)
            // } else {
            //     this.RenderTemplate(w, templateName, content, nil)
            // }

            if member := models.GetLoggedInMember(r); member != nil {
                if(content.PublicAccess != nil){
                    if(corehelpers.IntInSlice(member.Id, content.PublicAccess.Members)){
                        this.RenderTemplate(w, templateName, content, member)
                    } else if(member.Groups2PublicAccess(content.PublicAccess.Groups)){
                        this.RenderTemplate(w, templateName, content, member)
                    } else{
                        fmt.Println("Member do have access to this content")
                        applicationglobals.Templates["Unauthorized.tmpl"] = template.Must(template.ParseFiles("views/Layout.tmpl","views/Unauthorized.tmpl"))
                        this.RenderTemplate(w, "Unauthorized.tmpl", nil, nil)
                    }
                } else {
                    this.RenderTemplate(w, templateName, content, member)
                }
            } else {
                if(content.PublicAccess != nil){
                    fmt.Println("Member do have access to this content")
                    applicationglobals.Templates["Unauthorized.tmpl"] = template.Must(template.ParseFiles("views/Layout.tmpl","views/Unauthorized.tmpl"))
                    this.RenderTemplate(w, "Unauthorized.tmpl", nil, nil)
                } else{
                    this.RenderTemplate(w, templateName, content, nil)
                }
                
            }

            //this.RenderTemplate(w, templateName, &content, nil)
        }
    }

    
}