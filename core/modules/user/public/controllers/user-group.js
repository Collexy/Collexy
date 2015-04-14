var userGroupControllers = angular.module('userGroupControllers', []);

userGroupControllers.controller('UserGroupListCtrl', ['$scope', '$stateParams', 'UserGroup', function ($scope, $stateParams, UserGroup) {

  var userGroups = UserGroup.query();
  $scope.tree = userGroups;

}]);



userGroupControllers.controller('UserGroupEditCtrl', ['$scope', '$stateParams', 'UserGroup', function ($scope, $stateParams, UserGroup) {
  $scope.currentTab = 'properties';
  UserGroup.get({id:$stateParams.id}, function(){}).$promise.then(function(data){
    $scope.data = data;
  });
}]);