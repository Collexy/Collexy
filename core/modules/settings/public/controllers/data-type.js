angular.module("myApp").controller("DataTypeTreeCtrl", DataTypeTreeCtrl);
angular.module("myApp").controller("DataTypeEditCtrl", DataTypeEditCtrl);
angular.module("myApp").controller("DataTypeDeleteCtrl", DataTypeDeleteCtrl);
/**
 * @ngdoc controller
 * @name DataTypeTreeCtrl
 * @function
 * @description
 * The controller for the data type tree
 */
function DataTypeTreeCtrl($scope, DataType) {
    $scope.ContextMenuServiceName = "DataTypeContextMenu"
    $scope.tree = DataType.query();
}
/**
 * @ngdoc controller
 * @name DataTypeEditCtrl
 * @function
 * @description
 * The controller for editing a data type
 */
function DataTypeEditCtrl($scope, $stateParams, DataType) {
    $scope.editorOptions = {
        lineWrapping: true,
        lineNumbers: true,
        readOnly: 'nocursor',
        mode: 'htmlmixed',
    };
    $scope.currentTab = 'data-type';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        $scope.entity = DataType.get({
            id: $stateParams.id
        }, function(node) {});
    }
    
    // $scope.readFile = function() {
    //   var file = $scope.entity.dt.Path;
    //   for(var i = 0; i < $scope.node.tmpl.length; i++){
    //     if($scope.node.tmpl[i].Path == $scope.node.dt.Path){
    //       $scope.node.dt.Html = $scope.node.tmpl[i].Html;
    //       $scope.node.dt.Name = $scope.node.tmpl[i].Name;
    //     }
    //   }
    //   //$scope.node.dt.Html = 
    //   // Create a new FileReader Object
    //   //my_parser('http://localhost:8080/public/views/settings/data-type/tmpl/text-input.html');
    // };
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
            DataType.update({
                id: $stateParams.id
            }, $scope.entity, success, failure);
            console.log($scope.entity)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            DataType.create($scope.entity, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
}
/**
 * @ngdoc controller
 * @name DataTypeDeleteCtrl
 * @function
 * @description
 * The controller for deleting data type
 */
function DataTypeDeleteCtrl($scope, $stateParams, DataType) {
    $scope.delete = function(item) {
        console.log(item)
        DataType.delete({
            id: item.id
        }, function() {
            console.log("data type record with id: " + item.id + " deleted")
        })
    };
}