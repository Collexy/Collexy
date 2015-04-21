angular.module("myApp").controller("ContentTypeTreeCtrl", ContentTypeTreeCtrl);

/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function ContentTypeTreeCtrl($scope, $stateParams, ContentTypeChildren, ContentType, sessionService, ContextMenu, $interpolate, ngDialog) {
  $scope.rootNode = {
    "id": 1,
    "allowedPermissions": ["node_create"],
    "path": "1",
    "name": "root",
    "node_type": 5,
    "created_by": 1,
    "entity": {}
  }

  $scope.clickToOpen = function (item) {
        ngDialog.open({ 
          template: item.url,
          scope: $scope 
        });
    };

  $scope.deleteNode = function(item) {
    //alert("deleteNode")
    ContentType.delete({id: item.entity.node.id}, function(){
      console.log("content type and node record deleted with id: " + item.entity.node.id)
    })
    
  };

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
        data.nodes = ContentTypeChildren.query({ id: data.id}, function(node){
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
  var contentTypeNodes = ContentType.query({'type-id': '1', 'levels': '1'},{},function(node){
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
        // left: 40,
        // top: -80
        left: 0,
        top: -76
  }

  var $oLay = angular.element(document.getElementById('overlay'))

  $scope.showOptions = function (item,$event) {
      console.log("showoptions")
      var overlayDisplay;
      // if ($scope.currentItem === item){
      if ($oLay.css("display") == "block") {
          $scope.currentItem = null;
           overlayDisplay='none'
      }else{
          $scope.currentItem = item;
          overlayDisplay='block'
      }

      if(angular.element(document.getElementById('adminsubmenucontainer')).hasClass('expanded1')){
        offset = {
          // left: 40,
          // top: -80
          left: 0,
          top: -121
        }
      }
    
      var overLayCSS = {
          // left: $event.clientX + offset.left + 'px',
          // top: $event.clientY + offset.top + 'px',
          left: $event.clientX + offset.left + 'px',
          top: $event.clientY + offset.top + 'px',
          display: overlayDisplay
      }

       $oLay.css(overLayCSS)
  }

  $scope.getEntityInfo = function(currentItem){
    if(currentItem==undefined){
      
      currentItem = $scope.rootNode;
      $scope.getMenu(4);
    } else {
      currentItem['entity'] = ContentType.get({ id: currentItem.id}, function(data){
        var tempArray = getUserNodePermissions(currentItem, sessionService.getUser());
        var tempArray2 = [];
        if(typeof tempArray[0] == 'object'){
          for(var i = 0; i < tempArray.length; i++){
            tempArray2.push(tempArray[i].id)
          }
          currentItem['allowedPermissions'] = tempArray2;
        } else {
          currentItem['allowedPermissions'] = tempArray;
        }

        $scope.getMenu(currentItem.node_type);
      });
    }


  }

  $scope.interpolate = function (value) {
        return $interpolate(value)($scope);
    };

  $scope.getMenu = function (node_type){
    //alert(currentItem.entity.node.node_type)
    // First we get all pre-registered Context Menu items for the given nodeType
    $scope.contextMenu = ContextMenu.query({},{nodeType:node_type}, function(menu){
      //alert("lol1")
    })
    //alert($scope.contextMenu)
  }
}


var contentTypeControllers = angular.module('contentTypeControllers', []);

contentTypeControllers.controller('ContentTypeEditCtrl', ['$scope', '$stateParams', 'ContentType', 'Node', 'DataType', 'Template', function ($scope, $stateParams, ContentType, Node, DataType, Template) {
  $scope.currentTab = 'content-type';
  $scope.stateParams = $stateParams;
  if ($stateParams.id) {

    $scope.node = ContentType.get({extended: true},{ id: $stateParams.id}, function(node){
      console.log(node)
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

  $scope.allTemplates = Template.query({},{},function(node){
    });

  $scope.allContentTypes = ContentType.query({'type-id': '1'},{},function(allContentTypes){
    var availableCompositeContentTypes = []
    for(var i = 0; i < allContentTypes.length; i++){
      if($scope.node.parent_content_types.containsId(allContentTypes[i].id)){

      } else{
        availableCompositeContentTypes.push(allContentTypes[i])
      }
    }
    $scope.availableCompositeContentTypes = availableCompositeContentTypes;
    console.log(availableCompositeContentTypes)
    // for(var i = 0; i < allContentTypes.length; j++){
    //   if(allContentTypes[i].id == $scope.node.id){

    //   } else {
    //     for(var j= 0; i < $scope.node.parent_content_types.length; j++){
    //       if(allContentTypes[i].id == $scope.node.parent_content_types[j]){

    //       }
    //     }
    //   }
    // }
  });

  // $scope.allMediaTypes = Node.query({'node-type': '7'},{},function(node){
  //   });
  
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

    if ($stateParams.id) {
      console.log("update");
      ContentType.update({id: $stateParams.id}, $scope.node, success, failure);
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