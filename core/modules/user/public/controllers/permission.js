angular.module("myApp").controller("PermissionListCtrl", PermissionListCtrl);
angular.module("myApp").controller("PermissionEditCtrl", PermissionEditCtrl);
angular.module("myApp").controller('PermissionDeleteCtrl', PermissionDeleteCtrl);
//var permissionControllers = angular.module('permissionControllers', []);
function PermissionListCtrl($scope, $stateParams, Permission) {
	$scope.ContextMenuServiceName = "PermissionContextMenu"
    Permission.query({},{}, function(permissions){
        // for (var i = 0; i < permissions.length; i++) {
        //     permissions[i]["id"] = permissions[i].name
        // };
        $scope.tree = permissions;
    });
}

function PermissionEditCtrl($scope, $stateParams, Permission) {
    $scope.currentTab = 'properties';
    Permission.get({
        id: $stateParams.id
    }, function() {}).$promise.then(function(data) {
        $scope.data = data;
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
            Permission.update({
                id: $stateParams.id
            }, $scope.data, success, failure);
            console.log($scope.data)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            console.log($scope.data)
            Permission.create($scope.data, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
}
/**
 * @ngdoc controller
 * @name PermissionDeleteCtrl
 * @function
 * @description
 * The controller for deleting a user group
 */
function PermissionDeleteCtrl($scope, $stateParams, Permission) {
    $scope.delete = function(item) {
        console.log(item)
        Permission.delete({
            id: item.id
        }, function() {
            console.log("user group record with id: " + item.id + " deleted")
        })
    };
}