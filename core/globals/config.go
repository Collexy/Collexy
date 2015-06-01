package globals

import

(
	"encoding/xml"
)

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

// type MediaAccessItem struct {
// 	MediaId int `json:"media_id,omitempty"`
// 	//Domains     []string `json:"domains,omitempty"`
// 	//Url     string `json:"url,omitempty"`
// 	LoginPage        int   `json:"login_page,omitempty"`
// 	AccessDeniedPage int   `json:"access_denied_page,omitempty"`
// 	MemberGroups     []int `json:"member_groups,omitempty"`
// }

type MediaAccessItems struct {
	XMLName xml.Name     `xml:mediaItems"`
	Items   []*MediaAccessItem `xml:"item"`
	//Items map[string]*MediaItem `xml:"Item"` //does not work - how would you do it with maps?
}

type MediaAccessItem struct {
	XMLName xml.Name     `xml:item"`
	MediaId int `xml:"id,attr" json:"media_id,omitempty"`
	//Domains     []string `json:"domains,omitempty"`
	Url     string `xml:"url,attr" json:"url,omitempty"`
	LoginPage        int   `xml:"loginPage,attr,omitempty" json:"login_page,omitempty"`
	AccessDeniedPage int   `xml:"accessDeniedPage,attr,omitempty" json:"access_denied_page,omitempty"`
	Members     []int		`xml:"members>member" json:"members,omitempty"` // https://github.com/golang/go/issues/3688
	MemberGroups     []int `xml:"memberGroups>group" json:"member_groups,omitempty"` // https://github.com/golang/go/issues/3688
	// Members     []int		`xml:"members>member,id,attr" json:"members,omitempty"`
	// MemberGroups     []int `xml:"memberGroups>group,id,attr" json:"member_groups,omitempty"`
}

// type MediaAccessItemMember int {
// 	XMLName xml.Name     `xml:member"`
// }

// type MediaAccessItemGroup int {
// 	XMLName xml.Name     `xml:"group"`
// }

// type MediaAccessItemMembers struct {
// 	XMLName xml.Name     `xml:members"`
// 	Members []int		`xml: "id,attr"`
// }

// type MediaAccessItemGroups struct {
// 	XMLName xml.Name     `xml:"group"`
// 	Groups []int		`xml: "id,attr"`
// }





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
