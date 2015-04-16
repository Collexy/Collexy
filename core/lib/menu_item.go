package lib

type MenuItem struct {
	Name        string
	Icon        string
	Permissions []string
	Route       *Route
	//Action *Action
	SubMenu *Menu
}
