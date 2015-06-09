angular.module("myApp").controller("ContentTreeCtrl", ContentTreeCtrl);
angular.module("myApp").controller("ContentEditCtrl", ContentEditCtrl);
angular.module("myApp").controller("ContentDeleteCtrl", ContentDeleteCtrl);
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for the content tree
 */
function ContentTreeCtrl($scope, Content) {
    $scope.ContextMenuServiceName = "ContentContextMenu"
    $scope.EntityChildrenServiceName = "ContentChildren"
    Content.query({
        'type-id': '1',
        'levels': '1'
    }, {}, function(tree) {
        // for (var i = 0; i < tree.length; i++) {
        //     tree[i]["show"] = false;
        // }
        $scope.tree = tree;
    });
    // $scope.$watch("tree", function(newValue, oldValue) {
    //         $scope.tree = newValue;
    //         console.log(newValue);
    //     });
}
/**
 * @ngdoc controller
 * @name ContentEditCtrl
 * @function
 * @description
 * The controller for editing content
 */
function ContentEditCtrl($scope, $stateParams, Content, Template, ContentType, Member, MemberGroup, User, UserGroup, Permission) {
    // Content.query({
    //     'type-id': '1',
    //     'content-type': 9
    // }, {}, function(contentNodes) {
    //     $scope.contentNodes = contentNodes;
    // });
    
    Template.query({}, {}, function(allTemplates) {
        $scope.allTemplates = allTemplates;
    });
    // Tabs
    var tabs = [];
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        // Edit
        $scope.data = Content.get({
            id: $stateParams.id
        }, function(data) {
            if (data.content_type.tabs != null) {
                tabs = data.content_type.tabs;
            }
            if (data.content_type.parent_content_types != null) {
                for (var i = 0; i < data.content_type.parent_content_types.length; i++) {
                    if (data.content_type.parent_content_types[i].tabs != null) {
                        tabs = tabs.concat(data.content_type.parent_content_types[i].tabs).unique();
                    }
                }
            }
            if (data.content_type.composite_content_types != null) {
                for (var i = 0; i < data.content_type.composite_content_types.length; i++) {
                    if (data.content_type.composite_content_types[i].tabs != null) {
                        tabs = tabs.concat(data.content_type.composite_content_types[i].tabs).unique();
                    }
                }
            }
            console.log(tabs);
            $scope.tabs = tabs;
            $scope.currentTab = tabs[0].name;
            Member.query({}, function() {}).$promise.then(function(members) {
                $scope.allMembers = members;
                var availableMembers = [];
                var selectedMembers = [];
                if (typeof data.public_access_members != 'undefined') {
                    for (var i = 0; i < members.length; i++) {
                        var memberId = members[i].id
                        if (data.public_access_members["" + memberId + ""] === 'undefined') {
                            availableMembers.push(members[i])
                        } else {
                            selectedMembers.push(members[i])
                        }
                    }
                    // if(typeof data.public_access_members != 'undefined'){
                    //     for (var i = 0; i < data.public_access.members.length; i++) {
                    //         for (var j = 0; j < members.length; j++) {
                    //             //console.log("[i] = " + data.user_group_ids[i] + ", [j]: " +userGroups[j].id)
                    //             if (data.public_access.members[i] != members[j].id) {
                    //                 availableMembers.push(members[j])
                    //             } else {
                    //                 selectedMembers.push(members[j])
                    //             }
                    //         }
                    //     }
                    // }
                }
                // if(typeof data.public_access != 'undefined'){
                //     if(typeof data.public_access.members != 'undefined'){
                //         for (var i = 0; i < data.public_access.members.length; i++) {
                //             for (var j = 0; j < members.length; j++) {
                //                 //console.log("[i] = " + data.user_group_ids[i] + ", [j]: " +userGroups[j].id)
                //                 if (data.public_access.members[i] != members[j].id) {
                //                     availableMembers.push(members[j])
                //                 } else {
                //                     selectedMembers.push(members[j])
                //                 }
                //             }
                //         }
                //     }
                // }
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
                        if (data.public_access_member_groups["" + memberGroupId + ""] === 'undefined') {
                            availableMemberGroups.push(memberGroups[i])
                        } else {
                            selectedMemberGroups.push(memberGroups[i])
                        }
                    }
                }
                // if(typeof data.public_access != 'undefined'){
                //     if(typeof data.public_access.groups != 'undefined'){
                //         for (var i = 0; i < data.public_access.groups.length; i++) {
                //             for (var j = 0; j < memberGroups.length; j++) {
                //                 //console.log("[i] = " + data.user_group_ids[i] + ", [j]: " +userGroups[j].id)
                //                 if (data.public_access.groups[i] != memberGroups[j].id) {
                //                     availableMemberGroups.push(memberGroups[j])
                //                 } else {
                //                     selectedMemberGroups.push(memberGroups[j])
                //                 }
                //             }
                //         }
                //     }
                // }
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
                // var usersWithoutPermissions = [];
                // var usersWithPermissions = [];
                // if(typeof data.user_permissions !== 'undefined'){
                //     for (var i = 0; i < data.user_permissions.length; i++) {
                //         for (var j = 0; j < users.length; j++) {
                //             //console.log("[i] = " + data.user_group_ids[i] + ", [j]: " +userGroups[j].id)
                //             if (data.user_permissions[i].id != users[j].id) {
                //                 usersWithoutPermissions.push(users[j])
                //             } else {
                //                 usersWithPermissions.push(users[j])
                //             }
                //         }
                //     }
                //     usersWithoutPermissions.unique();
                //     usersWithPermissions.unique();
                //     $scope.usersWithoutPermissions = usersWithoutPermissions;
                //     $scope.usersWithPermissions = usersWithPermissions;
                // } else {
                //     $scope.usersWithoutPermissions = users;
                // }
            });
            UserGroup.query({}, function() {}).$promise.then(function(userGroups) {
                $scope.allUserGroups = userGroups;
                // var userGroupsWithoutPermissions = [];
                // var userGroupsWithPermissions = [];
                // if(typeof data.user_group_permissions !== 'undefined'){
                //     for (var i = 0; i < data.user_group_permissions.length; i++) {
                //         for (var j = 0; j < userGroups.length; j++) {
                //             //console.log("[i] = " + data.user_group_ids[i] + ", [j]: " +userGroups[j].id)
                //             if (data.user_group_permissions[i].id == userGroups[j].id) {
                //                 userGroupsWithPermissions.push(userGroups[j])
                //                 if(userGroupsWithoutPermissions.indexOf(userGroups[j]) > -1){
                //                     userGroupsWithoutPermissions.splice(userGroupsWithoutPermissions.indexOf(userGroups[j]),1)
                //                 }
                //                 break;
                //             } else {
                //                 userGroupsWithoutPermissions.push(userGroups[j])
                //             }
                //         }
                //     }
                //     userGroupsWithoutPermissions.unique();
                //     userGroupsWithPermissions.unique();
                //     $scope.userGroupsWithoutPermissions = userGroupsWithoutPermissions.unique();
                //     $scope.userGroupsWithPermissions = userGroupsWithPermissions.unique();
                // } else {
                //     $scope.userGroupsWithoutPermissions = userGroups.unique();
                // }
            });
            Permission.query({}, function() {}).$promise.then(function(permissions) {
                $scope.allPermissions = permissions;
            });
            // UserGroup.query({}, function(){
            // }).$promise.then(function(userGroups){
            //     $scope.allUserGroups = userGroups;
            //     var availableUserGroups = [];
            //     var selectedUserGroups = [];
            //     for (var i = 0; i < data.user_group_permissions.length; i++) {
            //         for (var j = 0; j < userGroups.length; j++) {
            //             //console.log("[i] = " + data.user_group_ids[i] + ", [j]: " +userGroups[j].id)
            //             if (data.user_group_permissions[i] != userGroups[j].id) {
            //                 availableUserGroups.push(userGroups[j])
            //             } else {
            //                 selectedUserGroups.push(userGroups[j])
            //             }
            //         }
            //     }
            //     availableUserGroups.unique();
            //     selectedUserGroups.unique();
            //     $scope.availableUserGroups = availableUserGroups;
            //     $scope.selectedUserGroups = selectedUserGroups;
            //     if(selectedUserGroups.length == 0){
            //         $scope.availableUserGroups = allUserGroups;
            //     }
            // }, function(){
            //     // error
            // })
        });
    } else {
        // New
        if ($scope.stateParams.content_type_id) {
            var ct = ContentType.getExtended({
                extended: true
            }, {
                id: $scope.stateParams.content_type_id
            }, function(c) {
                if (c.tabs != null) {
                    tabs = c.tabs;
                }
                if (c.parent_content_types != null) {
                    for (var i = 0; i < c.parent_content_types.length; i++) {
                        if (c.parent_content_types[i].tabs != null) {
                            tabs = tabs.concat(c.parent_content_types[i].tabs).unique();
                        }
                    }
                }
                if (c.composite_content_types != null) {
                    for (var i = 0; i < c.composite_content_types.length; i++) {
                        if (c.composite_content_types[i].tabs != null) {
                            tabs = tabs.concat(c.composite_content_types[i].tabs).unique();
                        }
                    }
                }
                console.log(tabs);
                $scope.tabs = tabs;
                $scope.currentTab = tabs[0].name;
            });
            $scope.data = {
                content_type: ct
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
        if ($scope.stateParams.content_type_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["content_type_id"] = parseInt($scope.stateParams.content_type_id);
            } else {
                $scope.data = {
                    content_type_id: parseInt($scope.stateParams.content_type_id)
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
        $scope.data["created_by"] = $scope.userSession.id;
    }
    $scope.$watch("data", function(newValue, oldValue) {
        if (typeof $scope.data.user_group_permissions != 'undefined') {
            //alert($scope.data.user_group_permissions.length)
            for (var k in $scope.data.user_group_permissions) {
                if (newValue.user_group_permissions[k].permissions.length == 0) {
                    delete newValue.user_group_permissions[k];
                }
            }
            $scope.data = newValue;
        }
        if (typeof $scope.data.user_permissions != 'undefined') {
            //alert($scope.data.user_group_permissions.length)
            for (var k in $scope.data.user_permissions) {
                if (newValue.user_permissions[k].permissions.length == 0) {
                    delete newValue.user_permissions[k];
                }
            }
            $scope.data = newValue;
        }
    }, true);
    // User.get({
    //     id: $stateParams.id
    // }, function() {}).$promise.then(function(data) {
    //     $scope.data = data;
    //     $scope.currentTab = data.username;
    //     UserGroup.query().$promise.then(function(userGroups) {
    //         $scope.allUserGroups = userGroups;
    //         var availableUserGroups = [];
    //         var selectedUserGroups = [];
    //         for (var i = 0; i < data.user_group_ids.length; i++) {
    //             for (var j = 0; j < userGroups.length; j++) {
    //                 //console.log("[i] = " + data.user_group_ids[i] + ", [j]: " +userGroups[j].id)
    //                 if (data.user_group_ids[i] != userGroups[j].id) {
    //                     availableUserGroups.push(userGroups[j])
    //                 } else {
    //                     selectedUserGroups.push(userGroups[j])
    //                 }
    //             }
    //         }
    //         availableUserGroups.unique();
    //         selectedUserGroups.unique();
    //         $scope.availableUserGroups = availableUserGroups;
    //         $scope.selectedUserGroups = selectedUserGroups;
    //     }, function() {
    //         //ERR
    //     })
    // }, function() {
    //     // ERR
    // });
    // $scope.moveItem = function(item, from, to) {
    //     alert("moveitem")
    //     var idx = from.indexOf(item);
    //     if (idx != -1) {
    //         from.splice(idx, 1);
    //         to.push(item);
    //     }
    //     var user_group_ids = [];
    //     for (var i = 0; i < $scope.selectedUserGroups.length; i++) {
    //         user_group_ids.push($scope.selectedUserGroups[i].id);
    //     }
    //     $scope.data.user_group_ids = user_group_ids;
    // };
    // $scope.moveAll = function(from, to) {
    //     angular.forEach(from, function(item) {
    //         to.push(item);
    //     });
    //     from.length = 0;
    // };
    //
    // Update / Create
    $scope.submit = function() {
        console.log("submit")

        function success(response) {
            console.log("success", response)
            //$location.path("/admin/users");
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
            Content.update({
                id: $stateParams.id
            }, $scope.data, success, failure);
            console.log($scope.data)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            Content.create($scope.data, success, failure);
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
            member_ids[""+$scope.selectedMembers[i].id+""]= true;
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
            member_group_ids[""+$scope.selectedMemberGroups[i].id+""]= true;
        }
        $scope.data.public_access_member_groups = member_group_ids;
        console.log($scope.data)
    };
}
/**
 * @ngdoc controller
 * @name ContentDeleteCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function ContentDeleteCtrl($scope, $stateParams, Content) {
    $scope.delete = function(item) {
        console.log(item)
        Content.delete({
            id: item.id
        }, function() {
            console.log("content record with id: " + item.id + " deleted")
        })
    };
}