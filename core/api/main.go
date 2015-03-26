package api

import(
	"github.com/gorilla/mux"
    "github.com/gorilla/context"
	"net/http"
	"collexy/core/api/controllers"
	apihelpers "collexy/core/api/helpers"
    "log"
    "fmt"
    "database/sql"
    "collexy/core/api/models"
    corehelpers "collexy/core/helpers"
)


func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("middleware", r.URL)

        sid := corehelpers.CheckCookie(w,r)

        u, err := models.GetUser(sid)
        
        if(u == nil || err==sql.ErrNoRows){
            fmt.Println(err)
            fmt.Fprintf(w, "You need to be logged in to access the API.")
        } else {
            models.SetLoggedInUser(r,u)
            h.ServeHTTP(w, r)
        }

        
    })
}

func Main(){

	// Setup API controllers
    nodeApiController := controllers.NodeApiController{}
    contentApiController := controllers.ContentApiController{}
    contentTypeApiController := controllers.ContentTypeApiController{}
    dataTypeApiController := controllers.DataTypeApiController{}
    templateApiController := controllers.TemplateApiController{}
    directoryApiController := controllers.DirectoryApiController{}

    userApiController := controllers.UserApiController{}
    userGroupApiController := controllers.UserGroupApiController{}

    memberApiController := controllers.MemberApiController{}
    memberTypeApiController := controllers.MemberTypeApiController{}

    angularRouteApiController := controllers.AngularRouteApiController{}
    routeApiController := controllers.RouteApiController{}
    menuApiController := controllers.MenuApiController{}

    menuLinkApiController := controllers.MenuLinkApiController{}

	// Setup API routes
	m := mux.NewRouter()
    publicApiRouter := mux.NewRouter()

    m.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.Delete)).Methods("DELETE")
	m.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.Post)).Methods("POST")
    m.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.GetBackendContentByNodeId)).Methods("GET")
    m.HandleFunc("/api/content/{nodeId:.*}", http.HandlerFunc(contentApiController.PutContent)).Methods("PUT")
    m.HandleFunc("/api/media/{nodeId:.*}", http.HandlerFunc(contentApiController.GetBackendContentByNodeId)).Methods("GET")

    // m.Get("/api/content-type/extended/{nodeId:.*}", http.HandlerFunc(contentTypeApiController.GetContentTypeExtendedByNodeId))
    m.HandleFunc("/api/content-type/{nodeId:.*}", http.HandlerFunc(contentTypeApiController.GetContentTypeByNodeId)).Methods("GET")
    //m.Get("/api/content-type/", http.HandlerFunc(contentTypeApiController.GetContentTypeByNodeId)) // not sure about this
    //m.Get("/api/content-type/", http.HandlerFunc(contentTypeApiController.GetContentTypes)) // not sure about this
    m.HandleFunc("/api/content-type/{nodeId:.*}", http.HandlerFunc(contentTypeApiController.PutContentType)).Methods("PUT")
    m.HandleFunc("/api/content-type/{nodeId:.*}", http.HandlerFunc(contentTypeApiController.Post)).Methods("POST")
    // m.Put("/api/content-type/{nodeId:.*}", http.HandlerFunc(contentTypeApiController.PutContentType))
    // m.Post("/api/content-type", http.HandlerFunc(contentTypeApiController.PostContentType))
    // m.Delete("/api/content-type/nodeId", http.HandlerFunc(contentTypeApiController.DeleteContentType))

    m.HandleFunc("/api/data-type/{nodeId:.*}", http.HandlerFunc(dataTypeApiController.GetByNodeId)).Methods("GET")
    m.HandleFunc("/api/data-type", http.HandlerFunc(dataTypeApiController.Get)).Methods("GET") // not sure about this
    m.HandleFunc("/api/data-type/{nodeId:.*}", http.HandlerFunc(dataTypeApiController.Put)).Methods("PUT")
    m.HandleFunc("/api/data-type/{nodeId:.*}", http.HandlerFunc(dataTypeApiController.Post)).Methods("POST")
    // m.Post("/api/data-type", http.HandlerFunc(dataTypeController.PostDataType))
    // m.Delete("/api/data-type/nodeId", http.HandlerFunc(dataTypeController.DeleteDataType))

    m.HandleFunc("/api/template", http.HandlerFunc(templateApiController.GetTemplates)).Methods("GET") // not sure about this
    m.HandleFunc("/api/template/{nodeId:.*}", http.HandlerFunc(templateApiController.GetTemplateByNodeId)).Methods("GET")
    
    m.HandleFunc("/api/template/{nodeId:.*}", http.HandlerFunc(templateApiController.PutTemplate)).Methods("PUT")
    m.HandleFunc("/api/template/{nodeId:.*}", http.HandlerFunc(templateApiController.PostTemplate)).Methods("POST")

    m.HandleFunc("/api/auth/{sid:.*}", http.HandlerFunc(apihelpers.AngularAuth)).Methods("GET")
    

    // Node
    //m.HandleFunc("/api/test/allowednode", http.HandlerFunc(nodeApiController.GetTest)).Methods("GET")
    m.HandleFunc("/api/node", http.HandlerFunc(nodeApiController.Get)).Methods("GET")
    m.HandleFunc("/api/node/{id:.*}/children", http.HandlerFunc(nodeApiController.GetByIdChildren)).Methods("GET")
    m.HandleFunc("/api/node/{id:.*}", http.HandlerFunc(nodeApiController.GetById)).Methods("GET")
    
    m.HandleFunc("/api/node", http.HandlerFunc(nodeApiController.Post)).Methods("POST")
    m.HandleFunc("/api/node/{id:.*}", http.HandlerFunc(nodeApiController.Put)).Methods("PUT")
    m.HandleFunc("/api/node/{id:.*}", http.HandlerFunc(nodeApiController.Delete)).Methods("DEL")

    // Directory
    m.HandleFunc("/api/directory/upload-file-test", http.HandlerFunc(directoryApiController.UploadFileTest)).Methods("POST")
    m.HandleFunc("/api/directory/{rootdir:.*}/{name:.*}", http.HandlerFunc(directoryApiController.Post)).Methods("POST")
    m.HandleFunc("/api/directory/{rootdir:.*}/{name:.*}", http.HandlerFunc(directoryApiController.Put)).Methods("PUT")

    m.HandleFunc("/api/directory/{rootdir:.*}/{name:.*}", http.HandlerFunc(directoryApiController.GetById)).Methods("GET")
    m.HandleFunc("/api/directory/{rootdir:.*}", http.HandlerFunc(directoryApiController.Get)).Methods("GET")
    
    // User
    publicApiRouter.HandleFunc("/api/public/user/login", http.HandlerFunc(userApiController.Login)).Methods("POST")
    m.HandleFunc("/api/user", http.HandlerFunc(userApiController.Get)).Methods("GET")
    m.HandleFunc("/api/user/{id:.*}", http.HandlerFunc(userApiController.GetById)).Methods("GET")
    m.HandleFunc("/api/user", http.HandlerFunc(userApiController.Post)).Methods("POST")

    m.HandleFunc("/api/user-group", http.HandlerFunc(userGroupApiController.Get)).Methods("GET")
    m.HandleFunc("/api/user-group/{id:.*}", http.HandlerFunc(userGroupApiController.GetById)).Methods("GET")

    // Member
    publicApiRouter.HandleFunc("/api/public/member/login", http.HandlerFunc(memberApiController.Login)).Methods("POST")

    m.HandleFunc("/api/member", http.HandlerFunc(memberApiController.Get)).Methods("GET")
    m.HandleFunc("/api/member/{id:.*}", http.HandlerFunc(memberApiController.GetById)).Methods("GET")

    // Member type
    m.HandleFunc("/api/member-type/{nodeId:.*}", http.HandlerFunc(memberTypeApiController.GetMemberTypeByNodeId)).Methods("GET")

    m.HandleFunc("/api/angular-route", http.HandlerFunc(angularRouteApiController.Get)).Methods("GET")
    m.HandleFunc("/api/route", http.HandlerFunc(routeApiController.Get)).Methods("GET")
    m.HandleFunc("/api/menu-link/{name:.*}", http.HandlerFunc(menuLinkApiController.GetByName)).Methods("GET")

    //publicApiRouter.HandleFunc("/api/public/route", http.HandlerFunc(routeApiController.Get)).Methods("GET")
    m.HandleFunc("/api/menu/{name:.*}", http.HandlerFunc(menuApiController.GetByName)).Methods("GET")
    publicApiRouter.HandleFunc("/api/public/testing", http.HandlerFunc(models.Test)).Methods("GET")
    publicApiRouter.HandleFunc("/api/public/contextmenutest/{nodeType:.*}", http.HandlerFunc(models.CmTest)).Methods("GET")
    
    http.Handle("/api/public/", publicApiRouter)
    http.Handle("/api/", context.ClearHandler(Middleware(m)))
}