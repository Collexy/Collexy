var mediaTypeControllers = angular.module('mediaTypeControllers', []);

mediaTypeControllers.controller('MediaTypeTreeCtrl', ['$scope', '$stateParams', 'NodeChildren','Node', function ($scope, $stateParams, NodeChildren, Node) {
  
  $scope.delete = function(data) {
    data.nodes = [];
  };
  $scope.expand_collapse = function(data) {
    if(!data.show){
      if(data.nodes == undefined){
        data.nodes = [];
      }
      if(data.nodes.length == 0){
        // REST API call to fetch the current node's imcontentTypete children
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
  //var contentTypeNodes = ContentTypeNode.query(function(node){
  var contentTypeNodes = Node.query({'node-type': '7', 'levels': '1'},{},function(node){
          //console.log(node)
        });
 
  $scope.tree = contentTypeNodes;

  $scope.menuOptions = [
      {
        "name": "Create",
        "target": "adminContentType.create",
        "attr": "href",
        "children": [
          {
            "name": "TextPage",
            "target": "adminContentType.create",
            "attr": "ui-sref"
          },
          {
            "name": "Product",
            "target": "adminContentType.create",
            "attr": "ui-sref"
          }
        ]
      },
      {
        "name": "Delete",
        "target": "adminContentType.delete",
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

mediaTypeControllers.controller('MediaTypeTreeCtrlEdit', ['$scope', '$stateParams', 'ContentType', 'Node', 'DataType', function ($scope, $stateParams, ContentType, Node, DataType) {
  $scope.currentTab = 'content-type';
  $scope.stateParams = $stateParams;
  if ($stateParams.nodeId) {

    $scope.node = ContentType.get({ nodeId: $stateParams.nodeId}, function(node){
      
    });
    //User.get({ userId: $stateParams.userId} , function(phone) {
  } else if ($stateParams.parent) {
    $scope.node = { "parent_content_type_node_id" : parseInt($stateParams.parent)}
  } else {
    $scope.node = {}
  }

  if($scope.stateParams.type){
      if(typeof $scope.node.node !== 'undefined'){
          $scope.node.node["node_type"] = parseInt($scope.stateParams.type);
      }
      else {
        $scope.node["node"] = { node_type: parseInt($scope.stateParams.type)}
      }
      
      
    }

  $scope.allTemplates = Node.query({'node-type': '3'},{},function(node){
    });

  $scope.allContentTypes = Node.query({'node-type': '4'},{},function(node){
    });

  $scope.allMediaTypes = Node.query({'node-type': '7'},{},function(node){
    });

  $scope.allDataTypes = DataType.query({},{},function(node){
    });

  console.log($scope.allDataTypes)
   
  $scope.toggleTab = function (item,$event) {
    $scope.currentTab = item;
  }

  $scope.checkAll = function() {
    $scope.node.ct.meta.allowed_templates_node_id = $scope.allTemplates.map(function(item) { return item.id; });
  };
  $scope.uncheckAll = function() {
    $scope.node.ct.meta.allowed_templates_node_id = [];
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
      ContentType.update({nodeId: $stateParams.nodeId}, $scope.node, success, failure);
      console.log($scope.node)
      //User.update($scope.user, success, failure);
    } else {
      console.log("create");
      ContentType.create($scope.node, success, failure);
      //User.create($scope.user, success, failure);
    }

  }
  $scope.aliasOrNodeName = function(alias, node_name){
    if(alias != null && alias != ""){
      return alias;
    }
    return node_name;
  }

  $scope.addTab = function(){
    if('tabs' in $scope.node){
      
    }
    else{
      $scope.node["tabs"] = [];
    }
    tab = {"name": "mytab", "properties" : []}
    $scope.node.tabs.push(tab);
  }
  $scope.addProp = function(tab){
    if('tabs' in $scope.node){
      var tabs = $scope.node.tabs;
      if(tabs.length > 0){
        for(var i = 0; i < tabs.length; i++){
          if(tabs[i].name == tab){
            if('properties' in tabs[i]){

            }
            else{
              tabs[i].properties = [];
            }
            tabs[i].properties.push({"name":"property name","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"});
          }
        }
        $scope.node.tabs = tabs;
      }
      
    }
  }
}]);