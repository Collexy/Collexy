func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("middleware", r.URL)

        sid := corehelpers.CheckCookie(w,r)
        u, _ := models.GetUser(sid)

        models.SetLoggedInUser(r,u)

        h.ServeHTTP(w, r)
    })
}

func FrontendMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("frontend middleware", r.URL)

        sid := corehelpers.CheckMemberCookie(w,r)
        m, _ := models.GetMember(sid)

        models.SetLoggedInMember(r,m)

        fmt.Println("YAY: logged in as - " + m.Username)

        h.ServeHTTP(w, r)
    })
}