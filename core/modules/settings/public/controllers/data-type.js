angular.module("myApp").controller("DataTypeTreeCtrl", DataTypeTreeCtrl);
angular.module("myApp").controller("DataTypeTreeCtrlEdit", DataTypeTreeCtrlEdit);
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function DataTypeTreeCtrl($scope, DataType) {
    $scope.tree = DataType.query();
    $scope.clickToOpen = function(item) {
        ngDialog.open({
            template: item.url,
            scope: $scope
        });
    };
    $scope.deleteNode = function(item) {
        //alert("deleteNode")
        DataType.delete({
            id: item.entity.node.id
        }, function() {
            console.log("data type and node record deleted with id: " + item.id)
        })
    };
}
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function DataTypeTreeCtrlEdit($scope, $stateParams, DataType) {
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
    $scope.toggleTab = function(item, $event) {
        $scope.currentTab = item;
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
//var dataTypeControllers = angular.module('dataTypeControllers', []);
// dataTypeControllers.controller('DataTypeTreeCtrl', ['$scope', '$stateParams', 'NodeChildren','Node', 'DataType', 'sessionService', 'ContextMenu', '$interpolate', 'ngDialog', function ($scope, $stateParams, NodeChildren, Node, DataType, sessionService, ContextMenu, $interpolate, ngDialog) {
//   $scope.rootNode = {
//     "id": 1,
//     "allowedPermissions": ["node_create"],
//     "path": "1",
//     "name": "root",
//     "node_type": 5,
//     "created_by": 1,
//     "entity": {}
//   }
//   $scope.clickToOpen = function (item) {
//         ngDialog.open({ 
//           template: item.url,
//           scope: $scope 
//         });
//     };
//   $scope.deleteNode = function(item) {
//     //alert("deleteNode")
//     DataType.delete({nodeId: item.entity.node.id}, function(){
//       console.log("data type and node record deleted with nodeId: " + item.entity.node.id)
//     })
//   };
//   $scope.delete = function(data) {
//     data.nodes = [];
//   };
//   $scope.expand_collapse = function(data) {
//     if(!data.show){
//       if(data.nodes == undefined){
//         data.nodes = [];
//       }
//       if(data.nodes.length == 0){
//         // REST API call to fetch the current node's imdataTypete children
//         data.nodes = NodeChildren.query({ nodeId: data.id}, function(node){
//           //console.log(node)
//         });
//         //console.log(data.nodes)
//         // data.nodes.push({
//         //     "name": "Node-lol-works",
//         //     "show": true,
//         //     "nodes": []
//         // })
//       }
//     }
//     data.show = !data.show;
//   }          
//   $scope.add = function(data) {
//     var post = data.nodes.length + 1;
//     var newName = data.name + '-' + post;
//                           data.nodes.push({name: newName, show: true, nodes: []});
//   };
//   //var dataTypeNodes = DataTypeNode.query(function(node){
//   var dataTypeNodes = Node.query({'node-type': '11', 'levels': '1'},{},function(node){
//           //console.log(node)
//         });
//   $scope.tree = dataTypeNodes;
//   $scope.menuOptions = [
//       {
//         "name": "Create",
//         "target": "adminDataType.create",
//         "attr": "href",
//         "children": [
//           {
//             "name": "TextPage",
//             "target": "adminDataType.create",
//             "attr": "ui-sref"
//           },
//           {
//             "name": "Product",
//             "target": "adminDataType.create",
//             "attr": "ui-sref"
//           }
//         ]
//       },
//       {
//         "name": "Delete",
//         "target": "adminDataType.delete",
//         "attr": "ui-sref"
//       }
//   ];
//   $scope.contextMenu = function(node_type){
//     alert(node_type);
//   }
//   var offset = {
//         // left: 40,
//         // top: -80
//         left: 0,
//         top: -76
//   }
//   var $oLay = angular.element(document.getElementById('overlay'))
//   $scope.showOptions = function (item,$event) {
//       console.log("showoptions")
//       var overlayDisplay;
//       // if ($scope.currentItem === item){
//       if ($oLay.css("display") == "block") {
//           $scope.currentItem = null;
//            overlayDisplay='none'
//       }else{
//           $scope.currentItem = item;
//           overlayDisplay='block'
//       }
//       if(angular.element(document.getElementById('adminsubmenucontainer')).hasClass('expanded1')){
//         offset = {
//           // left: 40,
//           // top: -80
//           left: 0,
//           top: -121
//         }
//       }
//       var overLayCSS = {
//           // left: $event.clientX + offset.left + 'px',
//           // top: $event.clientY + offset.top + 'px',
//           left: $event.clientX + offset.left + 'px',
//           top: $event.clientY + offset.top + 'px',
//           display: overlayDisplay
//       }
//        $oLay.css(overLayCSS)
//   }
//   $scope.getEntityInfo = function(currentItem){
//     if(currentItem==undefined){
//       currentItem = $scope.rootNode;
//       $scope.getMenu(11);
//     } else {
//       currentItem['entity'] = DataType.get({ nodeId: currentItem.id}, function(data){
//         var tempArray = getUserNodePermissions(currentItem, sessionService.getUser());
//         var tempArray2 = [];
//         if(typeof tempArray[0] == 'object'){
//           for(var i = 0; i < tempArray.length; i++){
//             tempArray2.push(tempArray[i].id)
//           }
//           currentItem['allowedPermissions'] = tempArray2;
//         } else {
//           currentItem['allowedPermissions'] = tempArray;
//         }
//         $scope.getMenu(currentItem.node_type);
//       });
//     }
//   }
//   $scope.interpolate = function (value) {
//         return $interpolate(value)($scope);
//     };
//   $scope.getMenu = function (node_type){
//     //alert(currentItem.entity.node.node_type)
//     // First we get all pre-registered Context Menu items for the given nodeType
//     $scope.contextMenu = ContextMenu.query({},{nodeType:node_type}, function(menu){
//       //alert("lol1")
//     })
//     //alert($scope.contextMenu)
//   }
// }]);
// dataTypeControllers.controller('DataTypeTreeCtrlEdit', ['$scope', '$stateParams', 'DataType', 'Node', function ($scope, $stateParams, DataType) {
// }]);