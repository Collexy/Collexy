angular.module("myApp").controller("AdminContentCtrl", AdminContentCtrl);
angular.module("myApp").controller("AdminMenuCtrl", AdminMenuCtrl);

function AdminContentCtrl($scope, $interpolate, sessionService, $state) {
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
                var hasSubs = false;
                for (var i = 0; i < subMenuItems.length; i++) {
                    if (id == subMenuItems[i].parent_id) {
                        hasSubs = true;
                        break;
                    }
                }
                if (!hasSubs) {
                    angular.element('#adminsubmenucontainer').removeClass("expanded1");
                    angular.element('#adminsubmenucontainer').addClass("collapse1");
                    angular.forEach(angular.element(".submenu-margin-top"), function(value, key) {
                        var a = angular.element(value);
                        a.removeClass('submenu-margin-top');
                        a.addClass('nosubmenu-margin-top');
                    });
                }
                $scope.currentSectionAlias = alias;
            }
        }
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