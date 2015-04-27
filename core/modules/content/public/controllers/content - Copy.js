angular.module("myApp").controller("ContentTreeCtrl", ContentTreeCtrl);
angular.module("myApp").controller("ContentEditCtrl", ContentEditCtrl);
angular.module("myApp").controller("ContentDeleteCtrl", ContentDeleteCtrl);
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function ContentTreeCtrl($scope, $stateParams, Content, ContentType, sessionService, ngDialog) {
    $scope.ContextMenuServiceName = "ContentContextMenu"
    $scope.EntityChildrenService = "ContentChildren"
    

    Content.query({
        'type-id': '1',
        'levels': '1'
    }, {}, function(tree) {
        $scope.tree = tree;
    });

    // $scope.clickToOpen = function(item) {
    //     ngDialog.open({
    //         template: item.url,
    //         scope: $scope
    //     });
    // };
    
    // $scope.expand_collapse = function(data) {
    //     if (!data.show) {
    //         if (data.nodes == undefined) {
    //             data.nodes = [];
    //         }
    //         if (data.nodes.length == 0) {
    //             // REST API call to fetch the current node's immediate children
    //             data.nodes = ContentChildren.query({
    //                 id: data.id
    //             }, function(node) {
    //                 //console.log(node)
    //             });
    //         }
    //     }
    //     data.show = !data.show;
    // }
    
    // $scope.getEntityInfo = function(currentItem) {
    //     console.log("getentityinfo")
    //     if (currentItem == undefined) {
    //         currentItem = {id:0}
    //         $scope.currentItem = currentItem;
    //     }
    //     ContentContextMenu.query({
    //         id: currentItem.id
    //     }, function() {}).$promise.then(function(data) {
    //         $scope.contextMenu = data;
    //         $scope.currentItem = currentItem;
    //     });
    // }
    
}

function ContentEditCtrl($scope, $stateParams, Content, Template, ContentType, $interpolate) {
    //$scope._ = _;
    var tabs = [];
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
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
        //User.get({ userId: $stateParams.userId} , function(phone) {
    } else {
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
                console.log(tabs);
                $scope.tabs = tabs;
                $scope.currentTab = tabs[0].name;
            });
            $scope.data = {
                content_type: ct
            }
        }
        if ($scope.stateParams.parent_id) {
            if (typeof $scope.data.node !== 'undefined') {
                $scope.data.node["parent_id"] = parseInt($scope.stateParams.parent_id);
            } else {
                $scope.data["node"] = {
                    parent_id: parseInt($scope.stateParams.parent_id)
                }
            }
        }
        if ($scope.stateParams.content_type_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["content_type_id"] = parseInt($scope.stateParams.content_type_id);
            }
        }
        if ($scope.stateParams.node_type) {
            if (typeof $scope.data.node !== 'undefined') {
                $scope.data.node["node_type"] = parseInt($scope.stateParams.node_type);
            } else {
                $scope.data["node"] = {
                    node_type: parseInt($scope.stateParams.node_type)
                }
            }
        }
    }
    $scope.allTemplates = Template.query({}, {}, function(node) {});
    $scope.aliasOrNodeName = function(alias, node_name) {
        if (alias != null && alias != "") {
            return alias;
        }
        return node_name;
    }
    $scope.filteredTemplates = function() {
        return $scope.allTemplates.filter(function(template) {
            return $scope.data.content_type.meta.allowed_template_ids.indexOf(template.id) !== -1;
        });
    };
    //console.log($scope.node)
    $scope.toggleTab = function(item, $event) {
        $scope.currentTab = item;
    }
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

function ContentDeleteCtrl($scope, $stateParams, Content){
    $scope.delete = function(item) {
        console.log(item)
        Content.delete({
            id: item.id
        }, function() {
            console.log("content record with id: " + item.entity.node.id + " deleted")
        })
    };
}