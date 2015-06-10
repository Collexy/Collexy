angular.module("myApp").controller("MemberTypeListCtrl", MemberTypeListCtrl);
angular.module("myApp").controller("MemberTypeEditCtrl", MemberTypeEditCtrl);
angular.module("myApp").controller("MemberTypeDeleteCtrl", MemberTypeDeleteCtrl);
/**
 * @ngdoc controller
 * @name MemberTypeListCtrl
 * @function
 * @description
 * The controller for the member type tree
 */
function MemberTypeListCtrl($scope, MemberType) {
    $scope.ContextMenuServiceName = "MemberTypeContextMenu"
    $scope.tree = MemberType.query();
}
/**
 * @ngdoc controller
 * @name MemberTypeEditCtrl
 * @function
 * @description
 * The controller for editing a member type
 */
function MemberTypeEditCtrl($scope, $stateParams, MemberType, DataType) {
    $scope.currentTab = 'member-type';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        MemberType.getExtended({
            extended: true
        }, {
            id: $stateParams.id
        }, function() {}).$promise.then(function(entity) {
            $scope.entity = entity;
        }, function() {
            console.log("Database error: Error fetching MemberType")
        });
    } else {
        $scope.entity = {
            "created_by" : $scope.userSession.id
        }
    }

    DataType.query().$promise.then(function(allDataTypes) {
        $scope.allDataTypes = allDataTypes;
    }, function() {
        console.log("Database error: Error fetching Data Types")
    });
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
            MemberType.update({
                id: $stateParams.id
            }, $scope.entity, success, failure);
            console.log($scope.entity)
        } else {
            console.log("create");
            MemberType.create($scope.entity, success, failure);
        }
    }
    $scope.addTab = function() {
        if ('tabs' in $scope.entity) {} else {
            $scope.entity["tabs"] = [];
        }
        tab = {
            "name": "mytab",
            "properties": []
        }
        $scope.entity.tabs.push(tab);
    }
    $scope.addProp = function(tab) {
        if ('tabs' in $scope.entity) {
            var tabs = $scope.entity.tabs;
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
                $scope.entity.tabs = tabs;
            }
        }
    }
}
/**
 * @ngdoc controller
 * @name MemberTypeDeleteCtrl
 * @function
 * @description
 * The controller for deleting member type
 */
function MemberTypeDeleteCtrl($scope, $stateParams, MemberType) {
    $scope.delete = function(item) {
        console.log(item)
        MemberType.delete({
            id: item.id
        }, function() {
            console.log("member type record with id: " + item.id + " deleted")
        })
    };
}