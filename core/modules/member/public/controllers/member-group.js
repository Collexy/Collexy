
angular.module("myApp").controller("MemberGroupListCtrl", MemberGroupListCtrl);
angular.module("myApp").controller("MemberGroupEditCtrl", MemberGroupEditCtrl);

/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function MemberGroupListCtrl($scope, MemberGroup){
  $scope.tree = MemberGroup.query();
}

function MemberGroupEditCtrl($scope, $stateParams, MemberGroup){
  $scope.currentTab = 'properties';
  $scope.stateParams = $stateParams;
  if ($stateParams.id) {
    MemberGroup.get({id: $stateParams.id}, function(){}).$promise.then(function(entity){
      $scope.entity = entity;
    }, function(){
      console.log("Database error: Error fetching membergroup")
    });
  }

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
      MemberGroup.update({id: $stateParams.id}, $scope.entity, success, failure);
      console.log($scope.entity)
    } else {
      console.log("create");
      MemberGroup.create($scope.entity, success, failure);
    }

  }
  $scope.aliasOrName = function(alias, name){
    if(alias != null && alias != ""){
      return alias;
    }
    return name;
  }
}




var memberGroupControllers = angular.module('memberGroupControllers', []);



