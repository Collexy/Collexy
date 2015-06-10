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

function MemberEditCtrl($scope, $stateParams, Member, MemberType, MemberGroup) {
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

            MemberGroup.query().$promise.then(function(memberGroups) {
                $scope.allMemberGroups = memberGroups;
                var availableMemberGroups = [];
                var selectedMemberGroups = [];
                for (var i = 0; i < data.member_group_ids.length; i++) {
                    for (var j = 0; j < memberGroups.length; j++) {
                        //console.log("[i] = " + data.member_group_ids[i] + ", [j]: " +memberGroups[j].id)
                        if (data.member_group_ids[i] != memberGroups[j].id) {
                            availableMemberGroups.push(memberGroups[j])
                        } else {
                            selectedMemberGroups.push(memberGroups[j])
                        }
                    }
                }
                availableMemberGroups.unique();
                selectedMemberGroups.unique();
                $scope.availableMemberGroups = availableMemberGroups;
                $scope.selectedMemberGroups = selectedMemberGroups;
            }, function() {
                //ERR
            })

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
        MemberGroup.query().$promise.then(function(memberGroups) {
            $scope.allMemberGroups = memberGroups;
            var availableMemberGroups = [];
            var selectedMemberGroups = [];
            
            availableMemberGroups = $scope.allMemberGroups;

            availableMemberGroups.unique();
            //selectedMemberGroups.unique();
            $scope.availableMemberGroups = availableMemberGroups;
            $scope.selectedMemberGroups = selectedMemberGroups;
        }, function() {
            //ERR
        })
        $scope.data["created_by"] = $scope.userSession.id;
    }

    $scope.moveItem = function(item, from, to) {
        alert("moveitem")
        var idx = from.indexOf(item);
        if (idx != -1) {
            from.splice(idx, 1);
            to.push(item);
        }
        var member_group_ids = [];
        for (var i = 0; i < $scope.selectedMemberGroups.length; i++) {
            member_group_ids.push($scope.selectedMemberGroups[i].id);
        }
        $scope.data.member_group_ids = member_group_ids;
    };
    $scope.moveAll = function(from, to) {
        angular.forEach(from, function(item) {
            to.push(item);
        });
        from.length = 0;
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
            Member.update({
                id: $stateParams.id
            }, $scope.data, success, failure);
            console.log($scope.data)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            console.log($scope.data)
            Member.create($scope.data, success, failure);
            //User.create($scope.user, success, failure);
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