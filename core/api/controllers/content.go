package controllers

import (
    "fmt"
    "net/http"
    //"time"
    //"database/sql"
    _ "github.com/lib/pq"
    //"collexy/helpers"
    "collexy/core/api/models"
    "strconv"
    //"log"
    //"github.com/gorilla/schema"
    "encoding/json"
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

type ContentApiController struct {}

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

func (this *ContentApiController) RenderTemplate(w http.ResponseWriter, name string, content *models.Content, member *models.Member) error {
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


func (this *ContentApiController) RenderAdminTemplate(w http.ResponseWriter, name string, content *models.Content, user *models.User) error {
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

// func (this *ContentApiController) Post(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     if user := models.GetLoggedInUser(r); user != nil {
//         var hasPermission bool = false
//         hasPermission = user.HasPermissions([]string{"content_create"})
//         if(hasPermission){

//             content := models.Content{}

//             err := json.NewDecoder(r.Body).Decode(&content)

//             if err != nil {
//                 http.Error(w, "Bad Request", 400)
//             }

//             content.Post()
//         } else {
//             fmt.Fprintf(w,"You do not have permission to create content")
//         }
//     }
// }

// func (this *ContentApiController) PutContent(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     if user := models.GetLoggedInUser(r); user != nil {
//         var hasPermission bool = false
//         hasPermission = user.HasPermissions([]string{"content_update"})
//         if(hasPermission){
//             content := models.Content{}

//             err := json.NewDecoder(r.Body).Decode(&content)

//             if err != nil {
//                 http.Error(w, "Bad Request", 400)
//             }

//             content.Update()
//         } else {
//             fmt.Fprintf(w,"You do not have permission to update content")
//         }
        
//     } 
// }

// func (this *ContentApiController) Delete(w http.ResponseWriter, r *http.Request){
//     w.Header().Set("Content-Type", "application/json")
//     if user := models.GetLoggedInUser(r); user != nil {
//         var hasPermission bool = false
//         hasPermission = user.HasPermissions([]string{"content_delete"})
//         if(hasPermission){
//             params := mux.Vars(r)

//             idStr := params["id"]
//             id, _ := strconv.Atoi(idStr)

//             models.DeleteContent(id)
//         } else {
//             fmt.Fprintf(w,"You do not have permission to delete content")
//         }
        
//     } 
// }

func (this *ContentApiController) RenderContent(w http.ResponseWriter, r *http.Request) {

    sid := corehelpers.CheckMemberCookie(w,r)
    m, _ := models.GetMember(sid)

    models.SetLoggedInMember(r,m)

    fmt.Println("RENDERCONTENT")

    // idStr := r.URL.Query().Get(":id")

    params := mux.Vars(r)
    // idStr := params["id"]

    // id, _ := strconv.Atoi(idStr)
    // fmt.Println(id)

    // content := models.GetFrontendContentById(id)

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
        var templateName string = content.Template.Name + ".tmpl"
        //templateName := strings.Replace(content["template_name"].(string), " ", "-", -1) + ".tmpl"
        if(templateName !=".tmpl"){
            if(applicationglobals.Templates[templateName] != nil){

            } else{        
                if(content.Template.ParentTemplates != nil){
                    templateArray := []string{"views/" + templateName}

                    if(content.Template.ParentTemplates != nil){
      
                        parentTemplateNodes := content.Template.ParentTemplates

                        v := make([]string, 0, len(parentTemplateNodes))

                        for  _, value := range parentTemplateNodes {
                            tplName := "views/" + value.Name + ".tmpl"
                            v = append(v, tplName)
                        }
                        templateArray = append(templateArray, v...)
                        
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


func (this *ContentApiController) GetBackendContentById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    idStr := params["id"]

    id, _ := strconv.Atoi(idStr)

    // Content object including the content' Node, the content type object. 
    // Note: Inside the content type object is an array of parent ContentTypes
    content := models.GetBackendContentById(id)

    res, err := json.Marshal(content)
    corehelpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

    /*
    1) 
    SELECT * FROM content, node 
    JOIN node ON node.id = content.node_id 
    WHERE node.id = $1
    2) 
    Create node
    3) 
    SELECT my_content_type.*, ffgd.*
    FROM content_type as my_content_type,
    LATERAL 
    (
        SELECT array_to_json(array_agg(c)) as parent_content_types
        from content_type as c, node
        where path @> subpath('1.3.5',0,nlevel('1.3.5')-1) and c.node_id = node.id
    ) ffgd
    WHERE my_content_type.node_id = 5
    4) create array of parent conten types
    5) create content type with parent content types inside
    6) create content with content type and node
    */

//     sql := `SELECT my_content_type.*,ffgd.*
// FROM content_type as my_content_type,
// LATERAL 
// (
//     SELECT array_to_json(array_agg(okidoki)) as parent_content_types
//     FROM (
//         SELECT c.id, c.node_id, c.name, c.description, c.icon, c.thumbnail, c.master_content_type_node_id, c.meta, gf.* as tabs
//         FROM content_type as c, node,
//         LATERAL (
//             select json_agg(row1) as tabs from((
//             select y.name, ss.properties
//             from json_to_recordset(
//             (
//                 select * 
//                 from json_to_recordset(
//                 (
//                     SELECT json_agg(ggg)
//                     from(
//                     SELECT tabs
//                     FROM 
//                     (   
//                         SELECT *
//                         FROM content_type as ct
//                         WHERE ct.id=c.id
//                     ) dsfds

//                     )ggg
//                 )
//                 ) as x(tabs json)
//             )
//             ) as y(name text, properties json),
//             LATERAL (
//             select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id',row.data_type, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
//             from(
//                 select name, "order", data_type, data_type.html as data_type_html, help_text, description
//                 from json_to_recordset(properties) 
//                 as k(name text, "order" int, data_type int, help_text text, description text)
//                 JOIN data_type
//                 ON data_type.id = k.data_type
//                 )row
//             ) ss
//             ))row1
//         ) gf
//         where path @> subpath('1.3.5',0,nlevel('1.3.5')-1) and c.node_id = node.id
//     )okidoki
// ) ffgd
// WHERE my_content_type.node_id = 5`















//     querystr := `SELECT *
// FROM 
// (
//     SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
//     content.id as content_id, content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
//     ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
//     ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.meta::json as ct_meta,
//     ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.name as ctm_name,
//     ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.meta::json as ctm_meta
//     FROM node
//     JOIN content
//     ON content.node_id = node.id
//     JOIN content_type as ct
//     ON ct.node_id = content.content_type_node_id
//     JOIN content_type as ctm
//     ON ctm.node_id = ct.master_content_type_node_id
//     WHERE node.id=$1
// )noden,
// (
// --select json_agg(row2) from((
// SELECT content_type.id as content_type_id, content_type.master_content_type_node_id as master_content_type_node_id, gf.json_agg as ct_tabs, gf2.json_agg as ctm_tabs
// FROM content_type,
// LATERAL (
//     select json_agg(row1) from((
//     select y.name, ss.properties
//     from json_to_recordset(
//         (
//             select * 
//             from json_to_recordset(
//                 (
//                     SELECT json_agg(ggg)
//                     from(
//                         SELECT tabs
//                         FROM 
//                         (   
//                             SELECT *
//                             FROM content_type as ct
//                             WHERE ct.id=content_type.id
//                         ) dsfds

//                     )ggg
//                 )
//             ) as x(tabs json)
//         )
//     ) as y(name text, properties json),
//     LATERAL (
//         --select json_agg(row) as properties
//         select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id',row.data_type, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
//         from(
//             select name, "order", data_type, data_type.html as data_type_html, help_text, description
//             from json_to_recordset(properties) 
//             as k(name text, "order" int, data_type int, help_text text, description text)
//             JOIN data_type
//             ON data_type.id = k.data_type
//             )row
//     ) ss
//     ))row1
// ) gf,
// LATERAL (
//     select json_agg(row2) from((
//     select p.name, ss2.properties
//     from json_to_recordset(
//         (
//             select * 
//             from json_to_recordset(
//                 (
//                     SELECT json_agg(ggg)
//                     from(
//                         SELECT tabs
//                         FROM 
//                         (   
//                             SELECT *
//                             FROM content_type as ctm
//                             WHERE ctm.node_id=content_type.master_content_type_node_id
//                         ) dsfds

//                     )ggg
//                 )
//             ) as x(tabs json)
//         )
//     ) as p(name text, properties json),
//     LATERAL (
//         --select json_agg(row) properties
//         select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id',row.data_type, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
//         from(
//             select name, "order", data_type, data_type.html as data_type_html, help_text, description
//             from json_to_recordset(properties) 
//             as k(name text, "order" int, data_type int, help_text text, description text)
//             JOIN data_type
//             ON data_type.id = k.data_type
//             )row
//     ) ss2
//     ))row2
//     )gf2
// ) rgr
// --JOIN content_type as lollo
// --ON lollo.node_id = rgr.master_content_type_node_id
// where rgr.content_type_id = ct_id`

//     w.Header().Set("Content-Type", "application/json")

//     // node
//     var parm_id, node_id, node_created_by, node_type int
//     var node_path, node_name string
//     var node_created_date time.Time

//     var content_id, content_node_id, content_content_type_node_id int
//     var content_meta []byte

//     var ct_id, ct_node_id, ct_master_content_type_node_id int
//     var ct_name, ct_description, ct_icon, ct_thumbnail string
//     var ct_tabs, ct_meta []byte

//     var ctm_id, ctm_node_id, ctm_master_content_type_node_id int
//     var ctm_name, ctm_description, ctm_icon, ctm_thumbnail string
//     var ctm_tabs, ctm_meta []byte
    
//     var content_type_id, master_content_type_node_id int

//     templol := r.URL.Query().Get(":id")
//     //fmt.Println(templol)
//     rofl,err1 := strconv.Atoi(templol)
//     corehelpers.PanicIf(err1)

//     parm_id = rofl

//     db := corehelpers.Db

//     row := db.QueryRow(querystr, parm_id)

//     err:= row.Scan(
//         &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
//         &content_id, &content_node_id, &content_content_type_node_id, &content_meta,
//         &ct_id, &ct_node_id, &ct_master_content_type_node_id, &ct_name, &ct_description, &ct_icon, &ct_thumbnail, &ct_meta,
//         &ctm_id, &ctm_node_id, &ctm_name, &ctm_description, &ctm_icon, &ctm_thumbnail, &ctm_meta, &content_type_id, &master_content_type_node_id, &ct_tabs, &ctm_tabs)

//     fmt.Println("content_type_id: ")
//     fmt.Println(content_type_id)

//     ct_tabs_str := string(ct_tabs)
//     //fmt.Println(ct_tabs_str + " dsfjldskfj skdf")
//     ct_meta_str := string(ct_meta)
//     //fmt.Println(ct_meta_str + " dsfjldskfj skdf")
//     ctm_tabs_str := string(ctm_tabs)
//     //fmt.Println(ctm_tabs_str + " dsfjldskfj skdf")
//     ctm_meta_str := string(ctm_meta)
//     //fmt.Println(ctm_meta_str + " dsfjldskfj skdf")

//     var ct_meta_map map[string]interface{}
//     json.Unmarshal([]byte(string(ct_meta_str)), &ct_meta_map)

//     var ctm_meta_map map[string]interface{}
//     json.Unmarshal([]byte(string(ctm_meta_str)), &ctm_meta_map)


// // Decode the json object

//     var ctTabs []models.Tab
//     var ctmTabs []models.Tab
//     //var tab Tab

//     json.Unmarshal([]byte(ct_tabs_str), &ctTabs)
//     json.Unmarshal([]byte(ctm_tabs_str), &ctmTabs)
//     fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)
//     //fmt.Println(ct_tabs_str)


//     var x map[string]interface{}
//     json.Unmarshal([]byte(string(content_meta)), &x)
//     fmt.Println(x)
    
//     fmt.Println("ksjdflk sdfkj: " + node_name)
//     //fmt.Println(string(content_meta))
//     node := models.Node{node_id,node_path,node_created_by, node_name, node_type, node_created_date, nil, nil, false}
//     content := models.Content{content_id,content_node_id,content_content_type_node_id, x}
//     ct := models.ContentType{ct_id, ct_node_id, ct_name, ct_description, ct_icon, ct_thumbnail, ct_master_content_type_node_id, ctTabs, ct_meta_map}
//     ctm := models.ContentType{ctm_id, ctm_node_id, ctm_name, ctm_description, ctm_icon, ctm_thumbnail, ctm_master_content_type_node_id, ctmTabs, ctm_meta_map}


//     //helpers.PanicIf(err)
//     switch {
//         case err == sql.ErrNoRows:
//                 log.Printf("No node with that ID.")
//         case err != nil:
//                 log.Fatal(err)
//         default:
//                 node_str,err := json.Marshal(node)
//                 corehelpers.PanicIf(err)
//                 content_str,err := json.Marshal(content)
//                 corehelpers.PanicIf(err)
//                 ct_str,err := json.Marshal(ct)
//                 corehelpers.PanicIf(err)
//                 ctm_str,err := json.Marshal(ctm)
//                 corehelpers.PanicIf(err)
//                 combined_res := fmt.Sprintf("{\"node\": %s, \"content\": %s, \"ct\": %s, \"ctm\": %s}",node_str, content_str, ct_str, ctm_str)
//                 fmt.Fprintf(w, "%s", combined_res)
//     }


// func (this *ContentApiController) Get(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     content := models.GetContent()

//     res, err := json.Marshal(content)
//     corehelpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)
// }

// func (this *ContentApiController) GetContentOfType(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     idStr := r.URL.Query().Get(":id")
//     id, _ := strconv.Atoi(idStr)

//     content := models.GetContentOfType(id)

//     res, err := json.Marshal(content)
//     corehelpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)
// }