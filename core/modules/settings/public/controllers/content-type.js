angular.module("myApp").controller("ContentTypeTreeCtrl", ContentTypeTreeCtrl);
angular.module("myApp").controller("ContentTypeEditCtrl", ContentTypeEditCtrl);
angular.module("myApp").controller("ContentTypeDeleteCtrl", ContentTypeDeleteCtrl);
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for the content type tree
 */
function ContentTypeTreeCtrl($scope, $stateParams, ContentType) {
    $scope.ContextMenuServiceName = "ContentTypeContextMenu"
    $scope.EntityChildrenServiceName = "ContentTypeChildren"
    ContentType.query({
        'type-id': '1',
        'levels': '1'
    }, {}, function(tree) {
        $scope.tree = tree;
    });
}
/**
 * @ngdoc controller
 * @name ContentTypeEditCtrl
 * @function
 * @description
 * The controller for editing a content type
 */
function ContentTypeEditCtrl($scope, $stateParams, ContentType, DataType, Template) {
    $scope.currentTab = 'content-type';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        $scope.node = ContentType.get({
            extended: true
        }, {
            id: $stateParams.id
        }, function(node) {
            console.log(node)
        });
    } else if ($stateParams.parent_id) {
        $scope.node = {
            "parent_id": parseInt($stateParams.parent_id),
            "created_by": $scope.userSession.id
        }
    } else {
        $scope.node = {
            "created_by": $scope.userSession.id
        }
    }
    if ($scope.stateParams.type_id) {
        if (typeof $scope.node !== 'undefined') {
            $scope.node["type_id"] = parseInt($scope.stateParams.type_id);
        } else {
            $scope.node = {
                type_id: parseInt($scope.stateParams.type_id)
            }
        }
    }
    $scope.allTemplates = Template.query({}, {}, function(node) {});
    $scope.allContentTypes = ContentType.query({
        'type-id': '1'
    }, {}, function(allContentTypes) {
        var availableCompositeContentTypes = []
        for (var i = 0; i < allContentTypes.length; i++) {
            if (typeof $scope.node.parent_content_types !== 'undefined' && $scope.node.parent_content_types.length > 0) {
                if ($scope.node.parent_content_types.containsId(allContentTypes[i].id)) {} else {
                    availableCompositeContentTypes.push(allContentTypes[i])
                }
            } else {
                if(allContentTypes[i].id != $stateParams.id){
                    availableCompositeContentTypes.push(allContentTypes[i])
                }
                
            }
        }
        $scope.availableCompositeContentTypes = availableCompositeContentTypes;
        console.log(availableCompositeContentTypes)
    });
    $scope.allDataTypes = DataType.query({}, {}, function(node) {});
    $scope.checkAll = function() {
        $scope.node.ct.meta.allowed_template_ids = $scope.allTemplates.map(function(item) {
            return item.id;
        });
    };
    $scope.uncheckAll = function() {
        $scope.node.ct.meta.allowed_template_ids = [];
    };
    $scope.submit = function() {
        console.log("submit")

        function success(response) {
            console.log("success", response)
            //$location.path("/admin/users");
        }

        function failure(response) {
            console.log("failure", response);
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
 * @name ContentTypeDeleteCtrl
 * @function
 * @description
 * The controller for deleting content type
 */
function ContentTypeDeleteCtrl($scope, $stateParams, ContentType) {
    $scope.delete = function(item) {
        console.log(item)
        ContentType.delete({
            id: item.id
        }, function() {
            console.log("content type record with id: " + item.id + " deleted")
        })
    };
}