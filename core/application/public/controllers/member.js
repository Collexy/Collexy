var memberControllers = angular.module('memberControllers', []);

memberControllers.controller('MemberCtrl', ['$scope', '$state', function ($scope, $state) {
	$scope.state = $state;
	console.log($state)
}]);

memberControllers.controller('MemberListCtrl', ['$scope', 'Member', function ($scope, Member) {
	$scope.members = Member.query({}, function(members){});
	//console.log($state.current.name)
}]);

memberControllers.controller('MemberEditCtrl', ['$scope', '$stateParams', 'Member', 'MemberType', function ($scope, $stateParams, Member, MemberType) {
  //$scope._ = _;
  $scope.currentTab = 'Membership';

  var tabs = [];

  $scope.data = Member.get({id:$stateParams.id}, function(data){
  	$scope.member_type = MemberType.getExtended({extended: true},{id:data.member_type_id}, function(member_type){});
  });

  $scope.toggleTab = function (item,$event) {
    $scope.currentTab = item;
  }



  
}]);