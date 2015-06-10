angular.module("myApp").controller("UserGroupListCtrl", UserGroupListCtrl);
angular.module("myApp").controller("UserGroupEditCtrl", UserGroupEditCtrl);
angular.module("myApp").controller('UserGroupDeleteCtrl', UserGroupDeleteCtrl);
//var userGroupControllers = angular.module('userGroupControllers', []);
function UserGroupListCtrl($scope, $stateParams, UserGroup) {
	$scope.ContextMenuServiceName = "UserGroupContextMenu"
    $scope.tree = UserGroup.query();
}

function UserGroupEditCtrl($scope, $stateParams, UserGroup) {
    $scope.currentTab = 'properties';
    UserGroup.get({
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