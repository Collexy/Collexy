angular.module("myApp").controller("AdminContentCtrl", AdminContentCtrl);
angular.module("myApp").controller("AdminMenuCtrl", AdminMenuCtrl);

function AdminContentCtrl($scope, $interpolate, sessionService, $state, $document) {
    //$scope.users = User.query();
    $scope.state = $state;
    $scope.userSession = sessionService.userSession;
    $scope.mySessionService = sessionService;
    console.log("adminmenustrl")
    console.log(sessionService.userSession)
    $scope.$watch("mySessionService.getUser()", function(newValue, oldValue) {
        $scope.userSession = newValue;
        console.log($scope.userSession)
    }, true);
    $scope.isDefined = function(obj, prop) {
        if (obj != null || obj != undefined) {
            if (prop in obj) {
                //alert("obj: " + obj.name + ", prop: " + prop + ", :::: true")
                return true;
            } else {
                //alert("obj: " + obj.name + ", prop: " + prop + ", :::: false")
                return false;
            }
        } else {
            //alert("obj: " + obj.name + ", prop: " + prop + ", :::: false")
            return false;
        }
    }
    $scope.userHasPermission = function(permissionsString) {
        //alert(permissionsString)
        permissions = permissionsString.split(",");
        var permFound = false;
        var hasPermissions = false;
        var user = $scope.userSession;
        // First check if a the currently logged in user has specific permissions per user-level
        if ($scope.isDefined(user, "permissions")) {
            if (user.permissions.length > 0) {}
        } else if ($scope.isDefined(user, "user_groups")) { // If first check fails, check permissions for each group if any
            if (user.user_groups.length > 0) {
                i_loop: for (var i = 0; i < permissions.length; i++) {
                    permFound = false;
                    j_loop: for (var j = 0; j < user.user_groups.length; j++) {
                        if (permFound) {
                            break j_loop;
                        }
                        k_loop: for (var k = 0; k < user.user_groups[j].permissions.length; k++) {
                            if (permFound) {
                                break k_loop;
                            }
                            if (permissions[i] == user.user_groups[j].permissions[k]) {
                                permFound = true;
                            }
                        }
                    }
                }
            }
        }
        hasPermissions = permFound;
        //console.log(hasPermissions)
        //alert(hasPermissions)
        return hasPermissions;
    }
    $scope.interpolate = function(value) {
        return $interpolate(value)($scope);
    };
    $scope.aliasOrName = function aliasOrName(alias, name) {
        if (alias != null && alias != "") {
            return alias;
        }
        return name;
    }
    $scope.filteredArray = function(allArray, allowedArray, allCompareOn, allowedCompareOn) {
        var filteredArray = [];
        if (allCompareOn != null) {
            if (allowedCompareOn != null) {
                for (var i = 0; i < allArray.length; i++) {
                    for (var j = 0; j < allowedArray.length; j++) {
                        if (allowedArray[j][allowedCompareOn] == allArray[i][allCompareOn]) {
                            filteredArray.push(allArray[i])
                        }
                    }
                }
            } else {
                for (var i = 0; i < allArray.length; i++) {
                    if (allowedArray.contains(allArray[i][allCompareOn])) {
                        filteredArray.push(allArray[i])
                    }
                }
            }
        } else if (allowedCompareOn != null) {
            for (var i = 0; i < allArray.length; i++) {
                for (var j = 0; j < allowedArray.length; j++) {
                    if (allowedArray[j][allowedCompareOn] == allArray[i]) {
                        filteredArray.push(allArray[i])
                    }
                }
            }
        } else {
            for (var i = 0; i < allArray.length; i++) {
                if (allowedArray.contains(allArray[i])) {
                    filteredArray.push(allArray[i])
                }
            }
        }
        return filteredArray;
    };
    // lookin, for item, filter by collection
    // $scope.filteredTemplates = function(allItemsCollection, item, allowedCollection) {
    //     return allItemsCollection.filter(function(item) {
    //         return allowedCollection.indexOf(item.id) !== -1;
    //     });
    // };
    // $scope.filteredTemplates = function() {
    //     return $scope.allTemplates.filter(function(template) {
    //         return $scope.data.content_type.meta.allowed_template_ids.indexOf(template.id) !== -1;
    //     });
    // };
    $document.bind('click', handleClickEvent);
}

function AdminMenuCtrl($scope, $state, Section) {
    //$scope.sections = AngularRoute.query({'type': '1'},{}, function(section){})
    $scope.sections = Section.query({}, {}, function(section) {})
    $scope.currentSectionAlias = "sectionContent";
    $scope.toggleSubMenu = function(alias) {
        //alert($scope.sections.length)
        var subMenuItems = [];
        for (var i = 0; i < $scope.sections.length; i++) {
            // if('parent_id' in $scope.sections[i]){
            //  if($scope.sections[i].parent_id == id){
            //      subMenuItems.push($scope.sections[i]);
            //  }
            // }
            if ('alias' in $scope.sections[i]) {
                if ($scope.sections[i].alias == alias) {
                    if ('children' in $scope.sections[i]) {
                        subMenuItems = $scope.sections[i].children;
                        // for (var j = 0; j < subMenuItems.length; j++) {
                        //     if ('children' in subMenuItems[j]) {
                        //         for (var k = 0; k < subMenuItems[j].children.length; k++) {
                        //             subMenuItems.push(subMenuItems[j].children[k])
                        //         };
                        //     }
                        // };
                    } else {
                        //alert("lol")
                        subMenuItems = [];
                    }
                }
            }
        }
        $scope.subMenuItems = subMenuItems;
        // console.log($scope.subMenuItems)
        if ($scope.currentSectionAlias == "sectionContent") {
            $scope.currentSectionAlias = alias;
        }
        if (angular.element('#adminsubmenucontainer').hasClass("collapse1")) {
            if ($scope.subMenuItems.length > 0) {
                if ($scope.currentSectionAlias == alias) {
                    angular.element('#adminsubmenucontainer').removeClass("collapse1");
                    angular.element('#adminsubmenucontainer').addClass("expanded1");
                    angular.forEach(angular.element(".nosubmenu-margin-top"), function(value, key) {
                        var a = angular.element(value);
                        a.removeClass('nosubmenu-margin-top');
                        a.addClass('submenu-margin-top');
                    });
                } else {
                    angular.element('#adminsubmenucontainer').removeClass("collapse1");
                    angular.element('#adminsubmenucontainer').addClass("expanded1");
                    angular.forEach(angular.element(".nosubmenu-margin-top"), function(value, key) {
                        var a = angular.element(value);
                        a.removeClass('nosubmenu-margin-top');
                        a.addClass('submenu-margin-top');
                    });
                    $scope.currentSectionAlias = alias;
                }
            }
        } else {
            if ($scope.currentSectionAlias == alias) {
                angular.element('#adminsubmenucontainer').removeClass("expanded1");
                angular.element('#adminsubmenucontainer').addClass("collapse1");
                angular.forEach(angular.element(".submenu-margin-top"), function(value, key) {
                    var a = angular.element(value);
                    a.removeClass('submenu-margin-top');
                    a.addClass('nosubmenu-margin-top');
                });
                $scope.currentSectionAlias = "sectionContent";
            } else {
                angular.element('#adminsubmenucontainer').removeClass("expanded1");
                angular.element('#adminsubmenucontainer').addClass("collapse1");
                angular.forEach(angular.element(".submenu-margin-top"), function(value, key) {
                    var a = angular.element(value);
                    a.removeClass('submenu-margin-top');
                    a.addClass('nosubmenu-margin-top');
                });
                // var hasSubs = false;
                // for (var i = 0; i < subMenuItems.length; i++) {
                //     if (id == subMenuItems[i].parent_id) {
                //         hasSubs = true;
                //         break;
                //     }
                // }
                // if (!hasSubs) {
                //     angular.element('#adminsubmenucontainer').removeClass("expanded1");
                //     angular.element('#adminsubmenucontainer').addClass("collapse1");
                //     angular.forEach(angular.element(".submenu-margin-top"), function(value, key) {
                //         var a = angular.element(value);
                //         a.removeClass('submenu-margin-top');
                //         a.addClass('nosubmenu-margin-top');
                //     });
                // }
                // $scope.currentSectionAlias = alias;
            }
        }
    }

    $scope.toggleSubMenu2 = function(alias) {
        //alert($$scope.subMenuItems.length)
        var subMenuItems2 = [];
        for (var i = 0; i < $scope.subMenuItems.length; i++) {
            // if('parent_id' in $scope.subMenuItems[i]){
            //  if($scope.subMenuItems[i].parent_id == id){
            //      subMenuItems2.push($scope.subMenuItems[i]);
            //  }
            // }
            if ('alias' in $scope.subMenuItems[i]) {
                if ($scope.subMenuItems[i].alias == alias) {
                    if ('children' in $scope.subMenuItems[i]) {
                        subMenuItems2 = $scope.subMenuItems[i].children;
                        for (var j = 0; j < subMenuItems2.length; j++) {
                            if ('children' in subMenuItems2[j]) {
                                for (var k = 0; k < subMenuItems2[j].children.length; k++) {
                                    subMenuItems2.push(subMenuItems2[j].children[k])
                                };
                            }
                        };
                    } else {
                        //alert("lol")
                        subMenuItems2 = [];
                    }
                }
            }
        }
        //alert(subMenuItems2.length)
        $scope.subMenuItems2 = subMenuItems2;
        // console.log($scope.subMenuItems2)
        if ($scope.currentSectionAlias == "sectionContent") {
            $scope.currentSectionAlias = alias;
        }
        if (angular.element('#adminsubmenucontainer2').hasClass("collapse1")) {
            if ($scope.subMenuItems2.length > 0) {
                // alert("lol")
                // alert($scope.currentSectionAlias)
                // alert(alias)
                if ($scope.currentSectionAlias == alias) {
                    angular.element('#adminsubmenucontainer2').removeClass("collapse1");
                    angular.element('#adminsubmenucontainer2').addClass("expanded1");
                    if(angular.element(".nosubmenu-margin-top").length>0){
                        angular.forEach(angular.element(".nosubmenu-margin-top"), function(value, key) {
                            var a = angular.element(value);
                            a.removeClass('nosubmenu-margin-top');
                            a.addClass('submenu-margin-top2');
                        });
                    } else {
                        angular.forEach(angular.element(".submenu-margin-top"), function(value, key) {
                            var a = angular.element(value);
                            a.removeClass('submenu-margin-top');
                            a.addClass('submenu-margin-top2');
                        });
                    }
                    
                } else {
                    // alert("lol1")
                    angular.element('#adminsubmenucontainer2').removeClass("collapse1");
                    angular.element('#adminsubmenucontainer2').addClass("expanded1");
                    angular.forEach(angular.element(".submenu-margin-top"), function(value, key) {
                        var a = angular.element(value);
                        a.removeClass('submenu-margin-top');
                        a.addClass('submenu-margin-top2');
                    });
                    $scope.currentSectionAlias = alias;
                }
            }
        } else {
            // alert("lol2")
            if ($scope.currentSectionAlias == alias) {
                angular.element('#adminsubmenucontainer2').removeClass("expanded1");
                angular.element('#adminsubmenucontainer2').addClass("collapse1");
                angular.forEach(angular.element(".submenu-margin-top2"), function(value, key) {
                    var a = angular.element(value);
                    a.removeClass('submenu-margin-top2');
                    a.addClass('submenu-margin-top');
                });
                $scope.currentSectionAlias = "sectionContent";
            } else {
               //alert("lol")
                angular.element('#adminsubmenucontainer2').removeClass("expanded1");
                angular.element('#adminsubmenucontainer2').addClass("collapse1");
                angular.forEach(angular.element(".submenu-margin-top2"), function(value, key) {
                    var a = angular.element(value);
                    a.removeClass('submenu-margin-top2');
                    a.addClass('submenu-margin-top');
                });
                // var hasSubs = false;
                // for (var i = 0; i < subMenuItems2.length; i++) {
                //     if (id == subMenuItems2[i].parent_id) {
                //         hasSubs = true;
                //         break;
                //     }
                // }
                // if (!hasSubs) {
                //     angular.element('#adminsubmenucontainer2').removeClass("expanded1");
                //     angular.element('#adminsubmenucontainer2').addClass("collapse1");
                //     angular.forEach(angular.element(".submenu-margin-top"), function(value, key) {
                //         var a = angular.element(value);
                //         a.removeClass('submenu-margin-top');
                //         a.addClass('nosubmenu-margin-top');
                //     });
                // }
                // $scope.currentSectionAlias = alias;
            }
        }
        console.log(subMenuItems2)
    }
}
// adminControllers.controller('AdminTopUserInfoCtrl', ['$scope', 'sessionService', '$state', function ($scope, sessionService, $state) {
//  $scope.userSession = sessionService.userSession;
//  $scope.mySessionService = sessionService;
//  console.log("adminmenustrl")
//  console.log(sessionService.userSession)
//  $scope.$watch("mySessionService.getUser()", function(newValue, oldValue) {
//      $scope.userSession = newValue;
//     },true);
// }]);
function handleClickEvent(event) {
    $oLay = angular.element(document.getElementById('overlay'));
    //scope.currentItem = null;
    $oLay.css({
        display: 'none'
    })
}