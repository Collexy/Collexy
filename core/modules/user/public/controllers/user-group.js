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