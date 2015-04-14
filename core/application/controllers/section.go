package controllers

import
(
	"fmt"
	//"collexy/core/api/models"
	"encoding/json"
	"net/http"
	corehelpers "collexy/core/helpers"
    // "collexy/globals"
    coreglobals "collexy/core/globals"
)

type SectionApiController struct{}

func (this *SectionApiController) Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    res, err := json.Marshal(coreglobals.Sections)
    corehelpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}



// stuff from the old menu_link.go file

// func (this *MenuLinkApiController) GetByName(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
    
//     if user := coremoduleuser.GetLoggedInUser(r); user != nil {
//         params := mux.Vars(r)
//         menuName := params["name"]

//         fmt.Println("TEST TEST 123321")
//         fmt.Println(menuName)
        
//         menuLinks := models.GetMenuLinks(menuName)
//         //fmt.Println(menuLinks[0].Permissions[0])
//         var allowedMenuLinks []models.MenuLink

//         //var userGroupPermissions []string

//         for i:=0; i<len(menuLinks); i++ {
//             for j:=0; j < len(menuLinks[i].Permissions); j++{
//                 for k:=0; k < len(user.UserGroupIds); k++ {
//                     for l:=0; l < len(user.UserGroups[k].Permissions); l++{
//                         for m:=0; m < len(menuLinks[i].Permissions); m++{
//                             if(menuLinks[i].Permissions[m] == user.UserGroups[k].Permissions[l]){
//                                 allowedMenuLinks = append(allowedMenuLinks,menuLinks[i])
//                             }
//                         }
                        
//                     }
                    
//                 }
//             }     
//         }

//         // for k:=0; k < len(user.UserGroups); k++ { // for each user  group
//         //     for l:=0; l < len(user.UserGroups[k].Permissions); l++{ // for each permission in user group
//         //         userGroupPermissions = append(userGroupPermissions,user.UserGroups[k].Permissions[l])
//         //     }
            
//         // }

//         // globals.RemoveDuplicatesStringSlice(&userGroupPermissions)

//         // for i:=0; i<len(menuLinks); i++ { // for each menu link
//         //     var linkFound bool = false
//         //     for j:=0; j < len(menuLinks[i].Permissions); j++{ // for each permission in menu link
//         //         if(linkFound){
//         //             break
//         //         }
//         //         for m:=0; m < len(userGroupPermissions); m++{
//         //             if(linkFound){
//         //                 break
//         //             }
//         //             //fmt.Println(userGroupPermissions[m])
//         //             //fmt.Println(menuLinks[i].Permissions[j])
//         //             if(userGroupPermissions[m] == menuLinks[i].Permissions[j]){
//         //                 allowedMenuLinks = append(allowedMenuLinks, menuLinks[i])
//         //                 linkFound = true
//         //             }
//         //         }
//         //     }     
//         // }

//         res, err := json.Marshal(allowedMenuLinks)
//         corehelpers.PanicIf(err)

//         fmt.Fprintf(w,"%s",res)
//     }
// }

// func RemoveDuplicatesMenuLinkSlice(xs *[]models.MenuLink) {
// found := make(map[string]bool)
// j := 0
// for i, x := range *xs {
// if !found[x.Name] {
// found[x.Name] = true
// (*xs)[j] = (*xs)[i]
// j++
// }
// }
// *xs = (*xs)[:j]
// }