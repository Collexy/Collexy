angular.module("myApp").controller("UserGroupListCtrl", UserGroupListCtrl);
angular.module("myApp").controller("UserGroupEditCtrl", UserGroupEditCtrl);
angular.module("myApp").controller('UserGroupDeleteCtrl', UserGroupDeleteCtrl);
//var userGroupControllers = angular.module('userGroupControllers', []);
function UserGroupListCtrl($scope, $stateParams, UserGroup) {
	$scope.ContextMenuServiceName = "UserGroupContextMenu"
    $scope.tree = UserGroup.query();
}

function UserGroupEditCtrl($scope, $stateParams, UserGroup, Permission) {
    $scope.currentTab = 'properties';
    if($stateParams.id){
        UserGroup.get({
            id: $stateParams.id
        }, function() {}).$promise.then(function(data) {
            $scope.data = data;
            Permission.query().$promise.then(function(permissions) {
                $scope.allPermissions = permissions;
                var availablePermissions = [];
                var selectedPermissions = [];
                if(typeof data.permissions != 'undefined'){
                    for (var i = 0; i < permissions.length; i++) {
                        if(data.permissions.containsName(permissions[i].name)){
                            selectedPermissions.push(permissions[i])
                        } else {
                           availablePermissions.push(permissions[i]) 
                        }
                    }
                } else {
                    availablePermissions = permissions;
                }
                
                // for (var i = 0; i < data.permissions.length; i++) {
                //     for (var j = 0; j < permissions.length; j++) {
                //         //console.log("[i] = " + data.permissions[i] + ", [j]: " +permission[j].id)
                //         if (data.permissions[i] != permissions[j].name) {
                //             availablePermissions.push(permissions[j])
                //         } else {
                //             selectedPermissions.push(permissions[j])
                //         }
                //     }
                // }
                availablePermissions.unique();
                selectedPermissions.unique();
                $scope.availablePermissions = availablePermissions;
                $scope.selectedPermissions = selectedPermissions;
            }, function() {
                //ERR
            })
        });
    } else {
        $scope.data = {}
        Permission.query().$promise.then(function(permissions) {
            $scope.allPermissions = permissions;
            var availablePermissions = [];
            var selectedPermissions = [];
            
            availablePermissions = permissions;
            
            
            // for (var i = 0; i < data.permissions.length; i++) {
            //     for (var j = 0; j < permissions.length; j++) {
            //         //console.log("[i] = " + data.permissions[i] + ", [j]: " +permission[j].id)
            //         if (data.permissions[i] != permissions[j].name) {
            //             availablePermissions.push(permissions[j])
            //         } else {
            //             selectedPermissions.push(permissions[j])
            //         }
            //     }
            // }
            availablePermissions.unique();
            selectedPermissions.unique();
            $scope.availablePermissions = availablePermissions;
            $scope.selectedPermissions = selectedPermissions;
        }, function() {
            //ERR
        })
    }
    

    $scope.moveItem = function(item, from, to) {
        alert("moveitem")
        var idx = from.indexOf(item);
        if (idx != -1) {
            from.splice(idx, 1);
            to.push(item);
        }
        var permissions = [];
        for (var i = 0; i < $scope.selectedPermissions.length; i++) {
            permissions.push($scope.selectedPermissions[i].name);
        }
        $scope.data.permissions = permissions;
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
            UserGroup.update({
                id: $stateParams.id
            }, $scope.data, success, failure);
            console.log($scope.data)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            console.log($scope.data)
            UserGroup.create($scope.data, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
}
/**
 * @ngdoc controller
 * @name UserGroupDeleteCtrl
 * @function
 * @description
 * The controller for deleting a user group
 */
function UserGroupDeleteCtrl($scope, $stateParams, UserGroup) {
    $scope.delete = function(item) {
        console.log(item)
        UserGroup.delete({
            id: item.id
        }, function() {
            console.log("user group record with id: " + item.id + " deleted")
        })
    };
}