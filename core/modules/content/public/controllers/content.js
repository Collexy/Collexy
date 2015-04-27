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
        $scope.tree = tree;
    });
}
/**
 * @ngdoc controller
 * @name ContentEditCtrl
 * @function
 * @description
 * The controller for editing content
 */
function ContentEditCtrl($scope, $stateParams, Content, Template, ContentType) {
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
    }
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