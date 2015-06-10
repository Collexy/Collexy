angular.module("myApp").controller("MemberGroupListCtrl", MemberGroupListCtrl);
angular.module("myApp").controller("MemberGroupEditCtrl", MemberGroupEditCtrl);
angular.module("myApp").controller('MemberGroupDeleteCtrl', MemberGroupDeleteCtrl);
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function MemberGroupListCtrl($scope, MemberGroup) {
    $scope.ContextMenuServiceName = "MemberGroupContextMenu"
    $scope.tree = MemberGroup.query();
}

function MemberGroupEditCtrl($scope, $stateParams, MemberGroup) {
    $scope.currentTab = 'properties';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        MemberGroup.get({
            id: $stateParams.id
        }, function() {}).$promise.then(function(entity) {
            $scope.entity = entity;
        }, function() {
            console.log("Database error: Error fetching membergroup")
        });
    } else {
        $scope.entity = {
            "created_by" : $scope.userSession.id
        }
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
            MemberGroup.update({
                id: $stateParams.id
            }, $scope.entity, success, failure);
            console.log($scope.entity)
        } else {
            console.log("create");
            MemberGroup.create($scope.entity, success, failure);
        }
    }

}
/**
 * @ngdoc controller
 * @name MemberGroupDeleteCtrl
 * @function
 * @description
 * The controller for deleting a member group
 */
function MemberGroupDeleteCtrl($scope, $stateParams, MemberGroup) {
    $scope.delete = function(item) {
        console.log(item)
        MemberGroup.delete({
            id: item.id
        }, function() {
            console.log("member group record with id: " + item.id + " deleted")
        })
    };
}