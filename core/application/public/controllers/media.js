function escapeRegExp(string) {
    return string.replace(/([.*+?^=!:${}()|\[\]\/\\])/g, "\\$1");
}

function replaceAll(string, find, replace) {
  return string.replace(new RegExp(escapeRegExp(find), 'g'), replace);
}

var mediaControllers = angular.module('mediaControllers', []);

mediaControllers.controller('MediaTreeCtrl', ['$scope', '$stateParams', 'NodeChildren','Node', 'Content', 'ContentType', function ($scope, $stateParams, NodeChildren, Node, Content, ContentType) {
  var contentTypes = Node.query({'node-type': '7'},{},function(node){});

  $scope.rootNode = {
    "id": 1,
    "path": "1",
    "name": "root",
    "node_type": 5,
    "created_by": 1,
    "entity": {
      "allowedContentTypes": contentTypes
    }
  }

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
        data.nodes = NodeChildren.query({ nodeId: data.id}, function(node){
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
  var contentNodes = Node.query({'node-type': '2', 'levels': '1'},{},function(node){
          //console.log(node)
        });

  $scope.tree = contentNodes;

  var offset = {
        // left: 40,
        // top: -80
        left: -115,
        top: -120
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
      var contentTypes = Node.query({'node-type': '7'},{},function(node){});
      
      currentItem = {
        "id": 1,
        "path": "1",
        "name": "root",
        "node_type": 5,
        "created_by": 1,
        "entity": {
          "allowedContentTypes": contentTypes
        }
      }
    }
      
    console.log(currentItem);
    if(currentItem.node_type != 5){
      currentItem['entity'] = Content.get({ nodeId: currentItem.id}, function(data){
        var allowedContentTypes = [];
        //console.log(data.content_type.meta)
        for(var i = 0; i < data.content_type.meta.allowed_content_types_node_id.length; i++){
            var ct = ContentType.get({nodeId: data.content_type.meta.allowed_content_types_node_id[i]}, function(){});
            allowedContentTypes.push(ct);
            
        }
        data['allowedContentTypes'] = allowedContentTypes;
      });
    }
    
    
    
  }


}]);

mediaControllers.controller('MediaTreeCtrlEdit', ['$scope', '$http', '$stateParams', 'Content', 'Template', 'ContentType', function ($scope, $http, $stateParams, Content, Template, ContentType) {
  //$scope._ = _;

  var tabs = [];

  $scope.stateParams = $stateParams;
    if ($stateParams.nodeId && $stateParams.nodeId != 'new') {

      $scope.data = Content.get({ nodeId: $stateParams.nodeId}, function(data){
        data.node.old_name = data.node.name;
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
        console.log(tabs);
        $scope.tabs = tabs;
        $scope.currentTab = tabs[0].name;
      
    });
    //User.get({ userId: $stateParams.userId} , function(phone) {
  } else{
    $scope.data = {}
    if($scope.stateParams.content_type_node_id){
      var ct = ContentType.getExtended({extended: true},{nodeId: $scope.stateParams.content_type_node_id}, function(c){
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
      $scope.data["content_type"] = ct;

    }
    if($scope.stateParams.parent_id){
      if(typeof $scope.data.node !== 'undefined'){
        $scope.data.node["parent_id"] = parseInt($scope.stateParams.parent_id);
      }
      else {
        $scope.data["node"] = { parent_id: parseInt($scope.stateParams.parent_id)}
      }
      
    }
    if($scope.stateParams.content_type_node_id){
      if(typeof $scope.data !== 'undefined'){
        $scope.data["content_type_node_id"] = parseInt($scope.stateParams.content_type_node_id);
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
      return $scope.data.content_type.meta.allowed_templates_node_id.indexOf(template.node_id) !== -1;
    });
  };

  //console.log($scope.node)
  
   
  $scope.toggleTab = function (item,$event) {
    $scope.currentTab = item;
  }

  // $scope.filesChanged = function(elm){
  //   $scope.files=elm.files
  //   $scope.$apply();
  //   console.log("mediaControllers scope: ")
  // console.log($scope.files);
  // }

  $scope.test = {files:undefined}

  $scope.upload=function(escapedPath){
    console.log($scope.test.files)
    var fd = new FormData() // put these 3 lines in a service
    angular.forEach($scope.test.files, function(file){
      fd.append('file', file)
    })
    $http.post('/api/directory/upload-file-test?path='+escapedPath, fd,
    {
      transformRequest: angular.identity, // returns first argument it is passed
      headers:{'Content-Type': undefined} //multipart/form-data
    })
    .success(function(d){
      console.log(d)
      console.log("works?")
    })
  }

  $scope.submit = function() {
    console.log("submit")

    function success(response) {
      console.log("success", response)
      var escapedPath = replaceAll(response.meta.path, '\\', '%5C');

      //console.log($scope.test.files);
      if($scope.test.files.length > 0){
        $scope.upload(escapedPath);
      }
      //$scope.upload(escapedPath);
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
      Content.update({nodeId: $stateParams.nodeId}, $scope.data, success, failure);
      console.log($scope.data)
      //User.update($scope.user, success, failure);
    } else {
      console.log("create");
      Content.create($scope.data, success, failure);
      //User.create($scope.user, success, failure);
    }

  }
  
}]);

