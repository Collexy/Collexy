angular.module("myApp").controller('MemberCtrl', MemberCtrl);
angular.module("myApp").controller('MemberListCtrl', MemberListCtrl);
angular.module("myApp").controller('MemberEditCtrl', MemberEditCtrl);
angular.module("myApp").controller('MemberDeleteCtrl', MemberDeleteCtrl);

function MemberCtrl($scope, $state) {
    $scope.state = $state;
    console.log($state)
}

function MemberListCtrl($scope, Member) {
    $scope.ContextMenuServiceName = "MemberContextMenu"
    $scope.members = Member.query();
    //console.log($state.current.name)
}

function MemberEditCtrl($scope, $stateParams, Member, MemberType) {
    //$scope._ = _;
    $scope.currentTab = 'Membership';
    var tabs = [];
    if ($stateParams.id) {
        $scope.data = Member.get({
            id: $stateParams.id
        }, function(data) {
            $scope.member_type = MemberType.getExtended({
                extended: true
            }, {
                id: data.member_type_id
            }, function(member_type) {});
        });
    } else {
        if ($stateParams.member_type_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["member_type_id"] = parseInt($stateParams.member_type_id);
            } else {
                $scope.data = {
                    member_type_id: parseInt($stateParams.member_type_id)
                }
                $scope.member_type = MemberType.getExtended({
                    extended: true
                }, {
                    id: $scope.data.member_type_id
                }, function(member_type) {})
            }
        }
    }
}
/**
 * @ngdoc controller
 * @name MemberDeleteCtrl
 * @function
 * @description
 * The controller for deleting a member
 */
function MemberDeleteCtrl($scope, $stateParams, Member) {
    $scope.delete = function(item) {
        console.log(item)
        Member.delete({
            id: item.id
        }, function() {
            console.log("member record with id: " + item.id + " deleted")
        })
    };
}