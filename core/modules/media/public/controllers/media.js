angular.module("myApp").controller("MediaTreeCtrl", MediaTreeCtrl);
angular.module("myApp").controller("MediaEditCtrl", MediaEditCtrl);
angular.module("myApp").controller("MediaDeleteCtrl", MediaDeleteCtrl);
/**
 * @ngdoc controller
 * @name MediaTreeCtrl
 * @function
 * @description
 * The controller for the media tree
 */
function MediaTreeCtrl($scope, Media) {
    $scope.ContextMenuServiceName = "MediaContextMenu"
    $scope.EntityChildrenServiceName = "MediaChildren"
    Media.query({
        'levels': '1'
    }, {}, function(tree) {
        $scope.tree = tree;
    });
}
/**
 * @ngdoc controller
 * @name MediaTreeCtrl
 * @function
 * @description
 * The controller for editing media
 */
function MediaEditCtrl($scope, $http, $stateParams, Media, Template, MediaType, MediaParents, Member, MemberGroup, User, UserGroup, Permission) {
    // Tabs
    var tabs = [];
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        // Edit
        Media.get({
            id: $stateParams.id
        }, function(data) {
            if (data.media_type.tabs != null) {
                tabs = data.media_type.tabs;
            }
            if (data.media_type.parent_media_types != null) {
                for (var i = 0; i < data.media_type.parent_media_types.length; i++) {
                    if (data.media_type.parent_media_types[i].tabs != null) {
                        tabs = tabs.concat(data.media_type.parent_media_types[i].tabs).unique();
                    }
                }
            }
            if (data.media_type.composite_media_types != null) {
                for (var i = 0; i < data.media_type.composite_media_types.length; i++) {
                    if (data.media_type.composite_media_types[i].tabs != null) {
                        tabs = tabs.concat(data.media_type.composite_media_types[i].tabs).unique();
                    }
                }
            }
            console.log(tabs);
            $scope.tabs = tabs;
            $scope.currentTab = tabs[0].name;
            $scope.data = data;
            console.log("lol")
            console.log(data.path)
            MediaParents.query({
                "id": data.parent_id
            }, function() {}).$promise.then(function(mediaParents) {
                var location = "media\\";
                for (var i = 0; i < mediaParents.length; i++) {
                    location = location + mediaParents[i].name;
                    if (i != mediaParents.length - 1) {
                        location = location + "\\"
                    }
                }
                $scope.location = location;
                $scope.location_url = pathToUrl(location)
                console.log(location)
            }, function() {
                //error
                var location = "media";
                $scope.location = location;
                $scope.location_url = pathToUrl(location)
                console.log(location)
            })
            $scope.originalData = angular.copy(data);
            $scope.latestData = angular.copy(data);
            Member.query({}, function() {}).$promise.then(function(members) {
                $scope.allMembers = members;
                var availableMembers = [];
                var selectedMembers = [];
                if (typeof data.public_access_members != 'undefined') {
                    for (var i = 0; i < members.length; i++) {
                        var memberId = members[i].id
                        if (typeof data.public_access_members["" + memberId + ""] === 'undefined') {
                            availableMembers.push(members[i])
                        } else {
                            selectedMembers.push(members[i])
                        }
                    }
                }
                availableMembers.unique();
                selectedMembers.unique();
                $scope.availableMembers = availableMembers;
                $scope.selectedMembers = selectedMembers;
                if (selectedMembers.length == 0) {
                    $scope.availableMembers = members;
                }
            }, function() {
                // error
            })
            MemberGroup.query({}, function() {}).$promise.then(function(memberGroups) {
                $scope.allMemberGroups = memberGroups;
                var availableMemberGroups = [];
                var selectedMemberGroups = [];
                if (typeof data.public_access_member_groups != 'undefined') {
                    for (var i = 0; i < memberGroups.length; i++) {
                        var memberGroupId = memberGroups[i].id
                        if (typeof data.public_access_member_groups["" + memberGroupId + ""] === 'undefined') {
                            availableMemberGroups.push(memberGroups[i])
                        } else {
                            selectedMemberGroups.push(memberGroups[i])
                        }
                    }
                    // for (var i = 0; i < memberGroups.length; i++) {
                    //     var memberGroupId = memberGroups[i].id
                    //     if (data.public_access_member_groups[""+memberGroupId+""] === 'undefined') {
                    //         availableMemberGroups.push(memberGroups[i])
                    //     } else {
                    //         selectedMemberGroups.push(memberGroups[i])
                    //     }
                    // }
                }
                availableMemberGroups.unique();
                selectedMemberGroups.unique();
                $scope.availableMemberGroups = availableMemberGroups;
                $scope.selectedMemberGroups = selectedMemberGroups;
                if (selectedMemberGroups.length == 0) {
                    $scope.availableMemberGroups = memberGroups;
                }
            }, function() {
                // error
            })
            User.query({}, function() {}).$promise.then(function(users) {
                $scope.allUsers = users;
            });
            UserGroup.query({}, function() {}).$promise.then(function(userGroups) {
                $scope.allUserGroups = userGroups;
            });
            Permission.query({}, function() {}).$promise.then(function(permissions) {
                $scope.allPermissions = permissions;
            });
            console.log($scope.data)
        });
    } else {
        // New
        if ($scope.stateParams.media_type_id) {
            var ct = MediaType.getExtended({
                extended: true
            }, {
                id: $scope.stateParams.media_type_id
            }, function(c) {
                if (c.tabs != null) {
                    tabs = c.tabs;
                }
                if (c.parent_media_types != null) {
                    for (var i = 0; i < c.parent_media_types.length; i++) {
                        if (c.parent_media_types[i].tabs != null) {
                            tabs = tabs.concat(c.parent_media_types[i].tabs).unique();
                        }
                    }
                }
                if (c.composite_media_types != null) {
                    for (var i = 0; i < c.composite_media_types.length; i++) {
                        if (c.composite_media_types[i].tabs != null) {
                            tabs = tabs.concat(c.composite_media_types[i].tabs).unique();
                        }
                    }
                }
                console.log(tabs);
                $scope.tabs = tabs;
                $scope.currentTab = tabs[0].name;
            });
            $scope.data = {
                media_type: ct,
                meta: {}
            }
        }
        if ($scope.stateParams.parent_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["parent_id"] = parseInt($scope.stateParams.parent_id);
            } else {
                $scope.data = {
                    parent_id: parseInt($scope.stateParams.parent_id)
                }
            }
        }
        if ($scope.stateParams.media_type_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["media_type_id"] = parseInt($scope.stateParams.media_type_id);
            } else {
                $scope.data = {
                    media_type_id: parseInt($scope.stateParams.media_type_id)
                }
            }
        }
        if ($scope.stateParams.type_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["type_id"] = parseInt($scope.stateParams.type_id);
            } else {
                $scope.data = {
                    type_id: parseInt($scope.stateParams.type_id)
                }
            }
        }
        Member.query({}, function() {}).$promise.then(function(members) {
            $scope.allMembers = members;
            var availableMembers = [];
            var selectedMembers = [];
            availableMembers.unique();
            selectedMembers.unique();
            $scope.availableMembers = availableMembers;
            $scope.selectedMembers = selectedMembers;
            if (selectedMembers.length == 0) {
                $scope.availableMembers = members;
            }
        }, function() {
            // error
        })
        MemberGroup.query({}, function() {}).$promise.then(function(memberGroups) {
            $scope.allMemberGroups = memberGroups;
            var availableMemberGroups = [];
            var selectedMemberGroups = [];
            availableMemberGroups.unique();
            selectedMemberGroups.unique();
            $scope.availableMemberGroups = availableMemberGroups;
            $scope.selectedMemberGroups = selectedMemberGroups;
            if (selectedMemberGroups.length == 0) {
                $scope.availableMemberGroups = memberGroups;
            }
        }, function() {
            // error
        })
        User.query({}, function() {}).$promise.then(function(users) {
            $scope.allUsers = users;
        });
        UserGroup.query({}, function() {}).$promise.then(function(userGroups) {
            $scope.allUserGroups = userGroups;
        });
        Permission.query({}, function() {}).$promise.then(function(permissions) {
            $scope.allPermissions = permissions;
        });
        $scope.originalData = angular.copy($scope.data);
        console.log($scope.originalData)
        $scope.latestData = angular.copy($scope.data);
    }
    // $scope.filesChanged = function(elm){
    //   $scope.files=elm.files
    //   $scope.$apply();
    //   console.log("mediaControllers scope: ")
    // console.log($scope.files);
    // }
    // $scope.files = [];
    // $scope.persistedFiles = [pathToUrl("media\\Sample Images\\TXT\\pic04.jpg")];
    // $scope.test = {
    //         files: undefined
    //     }
    // $scope.submit = function() {
    //     $scope.$emit("formSubmit"); 
    // }
    $scope.submit = function() {
        console.log("submit")

        function success(response) {
            console.log("success", response)
            var escapedPath = replaceAll($scope.location, '\\', '%5C');
            $scope.$broadcast("formSubmitSuccess");
            // if ($scope.files.length > 0) {
            //     $scope.upload(escapedPath);
            // }
        }

        function failure(response) {
            console.log("failure", response);
            // _.each(response.data, function(errors, key) {
            //   if (errors.length > 0) {
            //     _.each(errors, function(e) {
            //       $scope.form[key].$dirty = true;
            //       $scope.form[key].$setValidity(e, false);
            //     });
            //   }
            // });
        }
        if ($stateParams.id) {
            console.log("update");
            Media.update({
                id: $stateParams.id
            }, $scope.data, success, failure);
            console.log($scope.data)
                //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            Media.create($scope.data, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
    $scope.moveMember = function(item, from, to) {
        //alert("moveMember")
        var idx = from.indexOf(item);
        if (idx != -1) {
            from.splice(idx, 1);
            to.push(item);
        }
        var member_ids = {};
        for (var i = 0; i < $scope.selectedMembers.length; i++) {
            member_ids["" + $scope.selectedMembers[i].id + ""] = true;
        }
        $scope.data.public_access_members = member_ids;
        console.log($scope.data)
    };
    /** object instead of array */
    $scope.moveMemberGroup = function(item, from, to) {
        //alert("moveMemberGroup")
        console.log(from)
        var idx = from.indexOf(item);
        if (idx != -1) {
            from.splice(idx, 1);
            to.push(item);
        }
        var member_group_ids = {};
        for (var i = 0; i < $scope.selectedMemberGroups.length; i++) {
            //member_group_ids.push($scope.selectedMemberGroups[i].id);
            member_group_ids["" + $scope.selectedMemberGroups[i].id + ""] = true;
        }
        $scope.data.public_access_member_groups = member_group_ids;
        console.log($scope.data)
    };
    $scope.moveAll = function(from, to) {
        angular.forEach(from, function(item) {
            to.push(item);
        });
        from.length = 0;
    };
    // $scope.moveMemberGroup = function(item, from, to) {
    //     //alert("moveMemberGroup")
    //     var idx = from.indexOf(item);
    //     if (idx != -1) {
    //         from.splice(idx, 1);
    //         to.push(item);
    //     }
    //     var member_group_ids = [];
    //     for (var i = 0; i < $scope.selectedMemberGroups.length; i++) {
    //         member_group_ids.push($scope.selectedMemberGroups[i].id);
    //     }
    //     $scope.data.public_access_member_groups = member_group_ids;
    //     console.log($scope.data)
    // };
    // $scope.moveAll = function(from, to) {
    //     angular.forEach(from, function(item) {
    //         to.push(item);
    //     });
    //     from.length = 0;
    // };
}
/**
 * @ngdoc controller
 * @name MediaDeleteCtrl
 * @function
 * @description
 * The controller for deleting media
 */
function MediaDeleteCtrl($scope, $stateParams, Media) {
    $scope.delete = function(item) {
        console.log(item)
        Media.delete({
            id: item.id
        }, function() {
            console.log("media record with id: " + item.id + " deleted")
        })
    };
}