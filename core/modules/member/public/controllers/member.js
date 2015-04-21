angular.module("myApp").controller('MemberCtrl', MemberCtrl);
angular.module("myApp").controller('MemberListCtrl', MemberListCtrl);
angular.module("myApp").controller('MemberEditCtrl', MemberEditCtrl);

function MemberCtrl($scope, $state) {
    $scope.state = $state;
    console.log($state)
}

function MemberListCtrl($scope, Member) {
    $scope.members = Member.query({}, function(members) {});
    //console.log($state.current.name)
}

function MemberEditCtrl($scope, $stateParams, Member, MemberType) {
    //$scope._ = _;
    $scope.currentTab = 'Membership';
    var tabs = [];
    $scope.data = Member.get({
        id: $stateParams.id
    }, function(data) {
        $scope.member_type = MemberType.getExtended({
            extended: true
        }, {
            id: data.member_type_id
        }, function(member_type) {});
    });
    $scope.toggleTab = function(item, $event) {
        $scope.currentTab = item;
    }
}