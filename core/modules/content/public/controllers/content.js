angular.module("myApp").controller("ContentTreeCtrl", ContentTreeCtrl);
angular.module("myApp").controller("ContentTreeCtrlEdit", ContentTreeCtrlEdit);


/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function ContentTreeCtrl($scope, $stateParams, ContentChildren, Node, Content, ContentType, sessionService, ContextMenu, $interpolate, ngDialog, ContentContextMenu) {
  var allowedContentTypeNodes = [];
  var allowedContentTypes = [];

  ContentType.query({'type-id': '1'},{},function(){
  }).$promise.then(function(data){
    allowedContentTypes = data;
  }, function(){

  })
  // .$promise.then(function(data){
  //   console.log("success")
  //   console.log(allowedContentTypeNodes[0])
  //   for(var i = 0; i < allowedContentTypeNodes[0].length; i++){
  //       var ct = ContentType.get({id: allowedContentTypeNodes[0][i].id}, function(){});
  //       allowedContentTypes.push(ct);
        
  //   }
  // }, function(error) {
  //     // error handler
  // });



  
  $scope.rootNode = {
    "id": 0,
    "allowedPermissions": ["node_create"],
    "path": "1",
    "name": "root",
    "node_type": 5,
    "created_by": 1,
    // "entity": {
      "allowedContentTypes": allowedContentTypes
    // }
  }

  $scope.clickToOpen = function (item) {
        ngDialog.open({ 
          template: item.url,
          scope: $scope 
        });
    };

  $scope.deleteNode = function(item) {
    //alert("deleteNode")
    Content.delete({id: item.entity.node.id}, function(){
      console.log("content and node record deleted with id: " + item.entity.node.id)
    })
    
  };

  $scope.interpolate = function (value) {
        return $interpolate(value)($scope);
    };

  $scope.user = sessionService.getUser();
  $scope.delete = function(data) {
    data.nodes = [];
  };
  $scope.expand_collapse = function(data) {
    if(!data.show){
      if(data.nodes == undefined){
        data.nodes = [];
      }
      if(data.nodes.length == 0){
        // REST API call to fetch the current node's immediate children
        data.nodes = ContentChildren.query({ id: data.id}, function(node){
          //console.log(node)
        });

      }
    }
    data.show = !data.show;
  }          
  $scope.add = function(data) {
    var post = data.nodes.length + 1;
    var newName = data.name + '-' + post;
                          data.nodes.push({name: newName, show: true, nodes: []});
  };
  // var contentNodes = Node.query({},{'nodeTypeId': 1, 'levels': '1'},function(node){
  var contentNodes = Content.query({'type-id': '1', 'levels': '1'},{},function(node){
          //console.log(node)
        });

  $scope.tree = contentNodes;

  // $scope.menuOptions = [
  //     {
  //       "name": "Create",
  //       "target": "adminContent.create",
  //       "attr": "href",
  //       "children": [
  //         {
  //           "name": "TextPage",
  //           "target": "adminContent.create",
  //           "attr": "ui-sref"
  //         },
  //         {
  //           "name": "Product",
  //           "target": "adminContent.create",
  //           "attr": "ui-sref"
  //         }
  //       ]
  //     },
  //     {
  //       "name": "Delete",
  //       "target": "adminContent.delete",
  //       "attr": "ui-sref"
  //     }
  // ];

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
    }

    // if(currentItem==undefined){
    //   currentItem = $scope.rootNode;
    // } else {
      ContentContextMenu.query({id: currentItem.id},function(){

      }).$promise.then(function(data){
        $scope.contextMenu = data;
      });
    // }
    
    // console.log("getEntityInfo")
    // //console.log(currentItem);
    // if(currentItem==undefined){
    //   currentItem = $scope.rootNode;
    //   // data = currentItem
    //   // console.log($scope.currentItem)
    //   // $scope.currentItem = currentItem;
    //   // console.log($scope.currentItem)
    //   // console.log(currentItem)
    //   $scope.getMenu(1);
    // }
    // if(currentItem.node_type != 5){
    //   allowedContentTypes = [];

    //   currentItem['entity'] = Content.get({ id: currentItem.id}, function(data){
    //     var allowedContentTypes = [];
    //     //console.log(data.content_type.meta)
    //     for(var i = 0; i < data.content_type.meta.allowed_content_type_ids.length; i++){
    //         var ct = ContentType.get({id: data.content_type.meta.allowed_content_type_ids[i]}, function(){});
    //         allowedContentTypes.push(ct);
            
    //     }
    //     data['allowedContentTypes'] = allowedContentTypes;
    //     //alert(sessionService.getUser())

    //     var tempArray = getUserNodePermissions(currentItem, sessionService.getUser());
    //     var tempArray2 = [];
    //     if(typeof tempArray[0] == 'object'){
    //       for(var i = 0; i < tempArray.length; i++){
    //         tempArray2.push(tempArray[i].id)
    //       }
    //       currentItem['allowedPermissions'] = tempArray2;
    //     } else {
    //       currentItem['allowedPermissions'] = tempArray;
    //     }

    //     // currentItem['allowedPermissions'] = getUserNodePermissions(currentItem, sessionService.getUser());

    //     $scope.getMenu(currentItem.node_type);
    //   });
    // }
    
    
  }

  $scope.getMenu = function (node_type){
    //alert(currentItem.entity.node.node_type)
    // First we get all pre-registered Context Menu items for the given nodeType
    $scope.contextMenu = ContextMenu.query({},{nodeType:node_type}, function(menu){
      //alert("lol1")
    })
    //alert($scope.contextMenu)
  }

}
// }]);



function ContentTreeCtrlEdit($scope, $stateParams, Content, Template, ContentType, Node, $interpolate) {
  //$scope._ = _;

  var tabs = [];

  $scope.stateParams = $stateParams;
    if ($stateParams.id) {

      $scope.data = Content.get({ id: $stateParams.id}, function(data){
        if(data.content_type.tabs != null){
          tabs = data.content_type.tabs;
        }
        if(data.content_type.parent_content_types != null){
          for(var i = 0; i < data.content_type.parent_content_types.length; i++){
            if(data.content_type.parent_content_types[i].tabs != null){
              tabs = tabs.concat(data.content_type.parent_content_types[i].tabs).unique();
            }
          }
        }
        if(data.content_type.composite_content_types != null){
          for(var i = 0; i < data.content_type.composite_content_types.length; i++){
            if(data.content_type.composite_content_types[i].tabs != null){
              tabs = tabs.concat(data.content_type.composite_content_types[i].tabs).unique();
            }
          }
        }
        console.log(tabs);
        $scope.tabs = tabs;
        $scope.currentTab = tabs[0].name;
      
    });
    //User.get({ userId: $stateParams.userId} , function(phone) {
  } else{
    if($scope.stateParams.content_type_id){
      var ct = ContentType.getExtended({extended: true},{id: $scope.stateParams.content_type_id}, function(c){
        if(c.tabs != null){
          tabs = c.tabs;
        }
        if(c.parent_content_types != null){
          for(var i = 0; i < c.parent_content_types.length; i++){
            if(c.parent_content_types[i].tabs != null){
              tabs = tabs.concat(c.parent_content_types[i].tabs).unique();
            }
          }
        }
        console.log(tabs);
        $scope.tabs = tabs;
        $scope.currentTab = tabs[0].name;
      });
      $scope.data = { content_type: ct }

    }
    if($scope.stateParams.parent_id){
      if(typeof $scope.data.node !== 'undefined'){
        $scope.data.node["parent_id"] = parseInt($scope.stateParams.parent_id);
      }
      else {
        $scope.data["node"] = { parent_id: parseInt($scope.stateParams.parent_id)}
      }
      
    }
    if($scope.stateParams.content_type_id){
      if(typeof $scope.data !== 'undefined'){
        $scope.data["content_type_id"] = parseInt($scope.stateParams.content_type_id);
      }
      
    }
    if($scope.stateParams.node_type){
      if(typeof $scope.data.node !== 'undefined'){
          $scope.data.node["node_type"] = parseInt($scope.stateParams.node_type);
      }
      else {
        $scope.data["node"] = { node_type: parseInt($scope.stateParams.node_type)}
      }
      
      
    }
  }

  // if(typeof $scope.data !== 'undefined'){
  //    if(typeof $scope.data.content_type !== 'undefined'){
  //      if(typeof $scope.data.content_type.meta !== 'undefined'){
  //        if(typeof $scope.data.content_type.meta.allowed_templates_node_id !== 'undefined'
  //         && typeof($scope.data.meta !== 'undefined')){
  //         if(typeof($scope.data.meta.template_node_id == 'undefined')){
  //           $scope.data.meta.template_node_id = parseInt($scope.data.content_type.meta.template_node_id);
  //         }
  //        }
  //      }
  //    }
  // }
  

  $scope.allTemplates = Template.query({},{},function(node){
    });

  $scope.aliasOrNodeName = function(alias, node_name){
    if(alias != null && alias != ""){
      return alias;
    }
    return node_name;
  }

  
  $scope.filteredTemplates = function () {
    return $scope.allTemplates.filter(function (template) {
      return $scope.data.content_type.meta.allowed_template_ids.indexOf(template.id) !== -1;
    });
  };

  //console.log($scope.node)
  
   
  $scope.toggleTab = function (item,$event) {
    $scope.currentTab = item;
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

    if ($stateParams.id) {
      console.log("update");
      Content.update({id: $stateParams.id}, $scope.data, success, failure);
      console.log($scope.data)
      //User.update($scope.user, success, failure);
    } else {
      console.log("create");
      Content.create($scope.data, success, failure);
      //User.create($scope.user, success, failure);
    }

  }

  $scope.interpolate = function (value) {
        //alert(value)
        return $interpolate(value)($scope);
    };

    //$scope.contentNodes = $scope.contentNodes = Node.query({'node-type': '1', 'content-type':'64'},{},function(node){});

  
}


function getUserNodePermissions(currentItem, user){
    console.log(currentItem);
    var allowedArray = [];

    if(currentItem.entity.node.user_permissions != null){
      //for(var i = 0; i < user.permissions.length; i++){
        for(var j = 0; j< currentItem.entity.node.user_permissions.length; j++){
          if(user.id==currentItem.entity.node.user_permissions[j].id){
            // for(var k=0; k<currentItem.entity.node.user_permissions[j].permissions.length; k++){
            //   if(currentItem.entity.node.user_permissions[j].permissions[k] == user.permissions[i].id){
            //     allowedArray.push(user.permissions[i])
            //   }
            // }
            allowedArray = currentItem.entity.node.user_permissions[j].permissions;
            return allowedArray;
          }
        }
      //}
    } else if(user.user_groups != null){
      if(currentItem.entity.node.user_group_permissions != null){
        for(var h=0; h<user.user_groups.length; h++){
          if(user.user_groups[h].permissions != null){
            for(var i = 0; i < user.user_groups[h].permissions.length; i++){
              for(var j = 0; j < currentItem.entity.node.user_group_permissions.length; j++){
                if(user.user_groups[h].id==currentItem.entity.node.user_group_permissions[j].id){
                  for(var k=0; k<currentItem.entity.node.user_group_permissions[j].permissions.length; k++){
                    if(currentItem.entity.node.user_permissions[j].permissions[k] == user.user_groups[h].permissions[i].id){
                      allowedArray.push(user.user_groups[h].permissions[i])
                    }
                  }
                }
              }
            }
          }
        }
      } else {
        //if(user.user_groups != null){
          for(var i=0; i<user.user_groups.length; i++){
            if(user.user_groups[i].permissions != null){
              allowedArray = allowedArray.concat(user.user_groups[i].permissions).unique();
            }
          }
        //}
      }
    } 

    return allowedArray;
    
  }