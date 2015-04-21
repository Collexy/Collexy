angular.module("myApp").controller("UserGroupListCtrl", UserGroupListCtrl);
angular.module("myApp").controller("UserGroupEditCtrl", UserGroupEditCtrl);
//var userGroupControllers = angular.module('userGroupControllers', []);
function UserGroupListCtrl($scope, $stateParams, UserGroup) {
    var userGroups = UserGroup.query();
    $scope.tree = userGroups;
}

function UserGroupEditCtrl($scope, $stateParams, UserGroup) {
    $scope.currentTab = 'properties';
    UserGroup.get({
        id: $stateParams.id
    }, function() {}).$promise.then(function(data) {
        $scope.data = data;
    });
}