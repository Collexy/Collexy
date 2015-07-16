package lib

type Module struct {
	Name        string
	Alias       string
	Description string
	Sections    []Section
	ServerRoutes []ServerRoute
	Order       int
}
