angular.module("myApp").controller("MediaTypeTreeCtrl", MediaTypeTreeCtrl);
angular.module("myApp").controller("MediaTypeEditCtrl", MediaTypeEditCtrl);
angular.module("myApp").controller("MediaTypeDeleteCtrl", MediaTypeDeleteCtrl);
/**
 * @ngdoc controller
 * @name MediaTypeTreeCtrl
 * @function
 * @description
 * The controller for the media type tree
 */
function MediaTypeTreeCtrl($scope, $stateParams, ContentTypeChildren, ContentType, sessionService, $interpolate, ngDialog) {
    $scope.ContextMenuServiceName = "MediaTypeContextMenu"
    $scope.EntityChildrenServiceName = "ContentTypeChildren"
    ContentType.query({
        'type-id': '2',
        'levels': '1'
    }, {}, function(tree) {
        $scope.tree = tree;
    });
}
/**
 * @ngdoc controller
 * @name MediaTypeEditCtrl
 * @function
 * @description
 * The controller for editing a media type
 */
function MediaTypeEditCtrl($scope, $stateParams, ContentType, DataType) {
    $scope.currentTab = 'content-type';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        $scope.node = ContentType.get({
            extended: true
        }, {
            id: $stateParams.id
        }, function(node) {
            console.log(node);
        });
    } else if ($stateParams.parent_id) {
        $scope.node = {
            "parent_content_type_id": parseInt($stateParams.parent_id)
        }
    } else {
        $scope.node = {}
    }
    if ($scope.stateParams.type_id) {
        if (typeof $scope.node !== 'undefined') {
            $scope.node["type_id"] = parseInt($scope.stateParams.type_id);
        } else {
            $scope["node"] = {
                type_id: parseInt($scope.stateParams.type_id)
            }
        }
    }
    $scope.allMediaTypes = ContentType.query({
        'type-id': '2'
    }, {}, function(allMediaTypes) {
        var availableCompositeMediaTypes = []
        for (var i = 0; i < allMediaTypes.length; i++) {
            if (typeof $scope.node.parent_content_types !== 'undefined' && $scope.node.parent_content_types.length > 0) {
                if ($scope.node.parent_content_types.containsId(allMediaTypes[i].id)) {} else {
                    availableCompositeMediaTypes.push(allMediaTypes[i])
                }
            } else {
                if(allMediaTypes[i].id != $stateParams.id){
                    availableCompositeMediaTypes.push(allMediaTypes[i])
                }
                
            }
        }
        $scope.availableCompositeMediaTypes = availableCompositeMediaTypes;
        console.log(availableCompositeMediaTypes)
    });
    $scope.allDataTypes = DataType.query({}, {}, function(node) {});
    console.log($scope.allDataTypes)
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
            ContentType.update({
                id: $stateParams.id
            }, $scope.node, success, failure);
            console.log($scope.node)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            ContentType.create($scope.node, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
    $scope.addTab = function() {
        if ('tabs' in $scope.node) {} else {
            $scope.node["tabs"] = [];
        }
        tab = {
            "name": "mytab",
            "properties": []
        }
        $scope.node.tabs.push(tab);
    }
    $scope.addProp = function(tab) {
        if ('tabs' in $scope.node) {
            var tabs = $scope.node.tabs;
            if (tabs.length > 0) {
                for (var i = 0; i < tabs.length; i++) {
                    if (tabs[i].name == tab) {
                        if ('properties' in tabs[i]) {} else {
                            tabs[i].properties = [];
                        }
                        tabs[i].properties.push({
                            "name": "property name",
                            "order": 1,
                            "data_type_node_id": 2,
                            "help_text": "prop help text",
                            "description": "prop description"
                        });
                    }
                }
                $scope.node.tabs = tabs;
            }
        }
    }
}
/**
 * @ngdoc controller
 * @name MediaTypeDeleteCtrl
 * @function
 * @description
 * The controller for deleting content type
 */
function MediaTypeDeleteCtrl($scope, $stateParams, ContentType) {
    $scope.delete = function(item) {
        console.log(item)
        ContentType.delete({
            id: item.id
        }, function() {
            console.log("media type record with id: " + item.id + " deleted")
        })
    };
}