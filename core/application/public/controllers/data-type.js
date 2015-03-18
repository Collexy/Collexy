var dataTypeControllers = angular.module('dataTypeControllers', []);

dataTypeControllers.controller('DataTypeTreeCtrl', ['$scope', '$stateParams', 'NodeChildren','Node', function ($scope, $stateParams, NodeChildren, Node) {

  $scope.delete = function(data) {
    data.nodes = [];
  };
  $scope.expand_collapse = function(data) {
    if(!data.show){
      if(data.nodes == undefined){
        data.nodes = [];
      }
      if(data.nodes.length == 0){
        // REST API call to fetch the current node's imdataTypete children
        data.nodes = NodeChildren.query({ nodeId: data.id}, function(node){
          //console.log(node)
        });
        //console.log(data.nodes)
        // data.nodes.push({
        //     "name": "Node-lol-works",
        //     "show": true,
        //     "nodes": []
        // })
      }
    }
    data.show = !data.show;
  }          
  $scope.add = function(data) {
    var post = data.nodes.length + 1;
    var newName = data.name + '-' + post;
                          data.nodes.push({name: newName, show: true, nodes: []});
  };
  //var dataTypeNodes = DataTypeNode.query(function(node){
  var dataTypeNodes = Node.query({'node-type': '11', 'levels': '1'},{},function(node){
          //console.log(node)
        });

  $scope.tree = dataTypeNodes;

  //$scope.tree = [{name: "Node", show:true, nodes: []}];
//   $scope.tree = [{
//         "id": 1,
//         "name": "Node",
//             "show": true,
//             "nodes": [{
//             "name": "Node-1",
//             "id": 2,
//                 "show": false,
//                 "nodes": [{
//                 "name": "Node-1-1",
//                     "show": false,
//                     "nodes": []
//             }, {
//                 "id": 3,
//                 "name": "Node-1-2",
//                     "show": false,
//                     "nodes": []
//             }, {
//                 "id": 4,
//                 "name": "Node-1-3",
//                     "show": false,
//                     "nodes": []
//             }]
//         }, {
//             "id": 5,
//             "name": "Node-2",
//                 "show": false,
//                 "nodes": [{
//                 "id": 6,
//                 "name": "Node-2-1",
//                     "show": false,
//                     "nodes": [{
//                     "id": 7,
//                     "name": "Node-2-1-1",
//                         "show": false,
//                         "nodes": []
//                 }, {
//                     "id": 8,
//                     "name": "Node-2-1-2-ggg",
//                         "show": false,
//                         "nodes": []
//                 }]
//             }]
//         }, {
//             "id": 9,
//             "name": "Node-3",
//                 "show": false,
//                 "nodes": []
//         }, {
//             "id": 10,
//             "name": "Node-4",
//                 "show": false,
//                 "nodes": [{
//                 "id": 11,  
//                 "name": "Node-4-1",
//                     "show": false,
//                     "nodes": [{
//                       "id": 12,
//                       "name": "Node-4-1-1",
//                           "show": false,
//                           "nodes": []
//                       }, {
//                           "id": 13,
//                           "name": "Node-4-1-2",
//                               "show": false,
//                               "nodes": []
//                       }]
//                 }]
//         }
//       ]
// }];

  $scope.menuOptions = [
      {
        "name": "Create",
        "target": "adminDataType.create",
        "attr": "href",
        "children": [
          {
            "name": "TextPage",
            "target": "adminDataType.create",
            "attr": "ui-sref"
          },
          {
            "name": "Product",
            "target": "adminDataType.create",
            "attr": "ui-sref"
          }
        ]
      },
      {
        "name": "Delete",
        "target": "adminDataType.delete",
        "attr": "ui-sref"
      }
  ];

  $scope.contextMenu = function(node_type){
    alert(node_type);
  }

  var offset = {
        left: 40,
        top: -80
  }

  var $oLay = angular.element(document.getElementById('overlay'))

  $scope.showOptions = function (item,$event) {       
      var overlayDisplay;

      if ($scope.currentItem === item) {
          $scope.currentItem = null;
           overlayDisplay='none'
      }else{
          $scope.currentItem = item;
          overlayDisplay='block'
      }
    
      var overLayCSS = {
          left: $event.clientX + offset.left + 'px',
          top: $event.clientY + offset.top + 'px',
          display: overlayDisplay
      }

       $oLay.css(overLayCSS)
  }

}]);

dataTypeControllers.controller('DataTypeTreeCtrlEdit', ['$scope', '$stateParams', 'DataType', 'Node', function ($scope, $stateParams, DataType) {

  $scope.editorOptions = {
        lineWrapping : true,
        lineNumbers: true,
        readOnly: 'nocursor',
        mode: 'htmlmixed',
    };
    
  $scope.currentTab = 'data-type';

  $scope.stateParams = $stateParams;
  if ($stateParams.nodeId) {

    $scope.node = DataType.get({ nodeId: $stateParams.nodeId}, function(node){
      
    });
    //User.get({ userId: $stateParams.userId} , function(phone) {
  }

  // $scope.allTemplates = TemplateNode.query(function(node){
  //   });

  // $scope.allDataTypes = DataTypeNode.query(function(node){
  //   });

  console.log($scope.node)
   
  $scope.toggleTab = function (item,$event) {
    $scope.currentTab = item;
  }

  $scope.checkAll = function() {
    //$scope.node.ct.meta.allowed_templates_node_id = $scope.allTemplates.map(function(item) { return item.id; });
  };
  $scope.uncheckAll = function() {
    //$scope.node.ct.meta.allowed_templates_node_id = [];
  };

  $scope.readFile = function() {
    var file = $scope.node.dt.Path;
    for(var i = 0; i < $scope.node.tmpl.length; i++){
      if($scope.node.tmpl[i].Path == $scope.node.dt.Path){
        $scope.node.dt.Html = $scope.node.tmpl[i].Html;
        $scope.node.dt.Name = $scope.node.tmpl[i].Name;
      }
    }
    //$scope.node.dt.Html = 
    // Create a new FileReader Object
    //my_parser('http://localhost:8080/public/views/settings/data-type/tmpl/text-input.html');
  };

  $scope.submit = function() {
    console.log("submit")

    function success(response) {
      console.log("success", response)
      //$location.path("/admin/users");
    }

    function failure(response) {
      console.log("failure", response);

      // _.each(response.data, function(errors, key) {
      //   if (errors.length > 0) {
      //     _.each(errors, function(e) {
      //       $scope.form[key].$dirty = true;
      //       $scope.form[key].$setValidity(e, false);
      //     });
      //   }
      // });
    }

    if ($stateParams.nodeId) {
      console.log("update");
      DataType.update({nodeId: $stateParams.nodeId}, $scope.node, success, failure);
      console.log($scope.node)
      //User.update($scope.user, success, failure);
    } else {
      console.log("create");
      DataType.create($scope.node, success, failure);
      //User.create($scope.user, success, failure);
    }

  }

}]);