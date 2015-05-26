package globals

var Conf Config

type Config struct {
	DbName     string `json:"db_name,omitempty"`
	DbUser     string `json:"db_user,omitempty"`
	DbPassword string `json:"db_password,omitempty"`
	DbHost     string `json:"db_host,omitempty"`
	SslMode    string `json:"ssl_mode,omitempty"`
	Id404      int    `json:"id_404,omitempty"`
}

var MediaAccessConf map[string]*MediaAccessItem

type key int // already defined in user module.models

const MyProtectedMediakey key = 1

type MediaAccessItem struct {
	MediaId int `json:"media_id,omitempty"`
	//Domains     []string `json:"domains,omitempty"`
	//Url     string `json:"url,omitempty"`
	LoginPage        int   `json:"login_page,omitempty"`
	AccessDeniedPage int   `json:"access_denied_page,omitempty"`
	MemberGroups     []int `json:"member_groups,omitempty"`
}

// var Maccess MediaAccess

// type MediaAccess struct {
// 	Items     []MediaAccessItem `json:"items,omitempty"`
// }

// type MediaAccessItem struct {
// 	ContentId	int	`json:"content_id,omitempty"`
// 	Domains     []string `json:"domains,omitempty"`
// 	Url     string `json:"url,omitempty"`
// 	LoginPage int `json:"login_page,omitempty"`
// 	AccessDeniedPage     int `json:"access_denied_page,omitempty"`
// 	MemberGroups    []int `json:"member_groups,omitempty"`
// }
