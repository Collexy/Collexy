package lib

// type Route struct {
// 	Path string `json:"path"`
// 	Url string `json:"url,omitempty"`
// 	Components []RouteComponent `json:"components,omitempty"`
// 	IsAbstract bool `json:"is_abstract"`
// }

type Route struct {
	State       string `json:"state"`
	Url         string `json:"url,omitempty"`
	TemplateUrl string `json:"template_url,omitempty"`
	IsAbstract  bool   `json:"is_abstract"`
}
