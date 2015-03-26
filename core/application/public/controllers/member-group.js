var memberGroupControllers = angular.module('memberGroupControllers', []);

memberGroupControllers.controller('MemberGroupListCtrl', ['$scope', '$stateParams', 'NodeChildren','Node', function ($scope, $stateParams, NodeChildren, Node) {

  $scope.delete = function(data) {
    data.nodes = [];
  };
  $scope.expand_collapse = function(data) {
    if(!data.show){
      if(data.nodes == undefined){
        data.nodes = [];
      }
      if(data.nodes.length == 0){
        // REST API call to fetch the current node's immemberGroupte children
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
  //var memberGroupNodes = MemberGroupNode.query(function(node){
  var memberGroupNodes = Node.query({'node-type': '13', 'levels': '1'},{},function(node){
          //console.log(node)
        });
 
  $scope.tree = memberGroupNodes;

  // $scope.menuOptions = [
  //     {
  //       "name": "Create",
  //       "target": "adminMemberGroup.create",
  //       "attr": "href",
  //       "children": [
  //         {
  //           "name": "TextPage",
  //           "target": "adminMemberGroup.create",
  //           "attr": "ui-sref"
  //         },
  //         {
  //           "name": "Product",
  //           "target": "adminMemberGroup.create",
  //           "attr": "ui-sref"
  //         }
  //       ]
  //     },
  //     {
  //       "name": "Delete",
  //       "target": "adminMemberGroup.delete",
  //       "attr": "ui-sref"
  //     }
  // ];

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



memberGroupControllers.controller('MemberGroupEditCtrl', ['$scope', '$stateParams', 'Node', function ($scope, $stateParams, Node) {
  $scope.currentTab = 'properties';
  $scope.stateParams = $stateParams;
  if ($stateParams.nodeId) {

    $scope.node = Node.get({id: $stateParams.nodeId}, function(node){
      
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
      Node.update({id: $stateParams.nodeId}, $scope.node, success, failure);
      console.log($scope.node)
      //User.update($scope.user, success, failure);
    } else {
      console.log("create");
      Node.create($scope.node, success, failure);
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