package globals

type IRoute interface {
    AddChildren(child IRoute)
}

var Routes []IRoute