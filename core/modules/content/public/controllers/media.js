angular.module("myApp").controller("MediaTreeCtrl", MediaTreeCtrl);
angular.module("myApp").controller("MediaEditCtrl", MediaEditCtrl);
angular.module("myApp").controller("MediaDeleteCtrl", MediaDeleteCtrl);
/**
 * @ngdoc controller
 * @name MediaTreeCtrl
 * @function
 * @description
 * The controller for the media tree
 */
function MediaTreeCtrl($scope, Content) {
    $scope.ContextMenuServiceName = "MediaContextMenu"
    $scope.EntityChildrenServiceName = "ContentChildren"
    Content.query({
        'type-id': '2',
        'levels': '1'
    }, {}, function(tree) {
        $scope.tree = tree;
    });
}
/**
 * @ngdoc controller
 * @name MediaTreeCtrl
 * @function
 * @description
 * The controller for editing media
 */
function MediaEditCtrl($scope, $http, $stateParams, Content, Template, ContentType, ContentParents) {
    // Tabs
    var tabs = [];
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        // Edit
        Content.get({
            id: $stateParams.id
        }, function(data) {
            if (data.content_type.tabs != null) {
                tabs = data.content_type.tabs;
            }
            if (data.content_type.parent_content_types != null) {
                for (var i = 0; i < data.content_type.parent_content_types.length; i++) {
                    if (data.content_type.parent_content_types[i].tabs != null) {
                        tabs = tabs.concat(data.content_type.parent_content_types[i].tabs).unique();
                    }
                }
            }
            if (data.content_type.composite_content_types != null) {
                for (var i = 0; i < data.content_type.composite_content_types.length; i++) {
                    if (data.content_type.composite_content_types[i].tabs != null) {
                        tabs = tabs.concat(data.content_type.composite_content_types[i].tabs).unique();
                    }
                }
            }
            console.log(tabs);
            $scope.tabs = tabs;
            $scope.currentTab = tabs[0].name;
            $scope.data = data;

            console.log("lol")
            console.log(data.path)
            ContentParents.query(
                {
                    "id": data.parent_id
                }, function(){}).$promise.then(function(contentParents){
                        var location = "media\\";
                        for(var i = 0; i < contentParents.length; i++){
                            location = location + contentParents[i].name;
                            if(i != contentParents.length-1){
                                location = location + "\\"
                            }
                        }
                        $scope.location = location;
                        $scope.location_url = pathToUrl(location)
                        console.log(location)
                    }, 
                    function(){
                        //error
                        var location = "media";
                        $scope.location = location;
                        $scope.location_url = pathToUrl(location)
                        console.log(location)
                    }
                )

            $scope.originalData = angular.copy(data);
            $scope.latestData = angular.copy(data);
        });
    } else {
        // New
        if ($scope.stateParams.content_type_id) {
            var ct = ContentType.getExtended({
                extended: true
            }, {
                id: $scope.stateParams.content_type_id
            }, function(c) {
                if (c.tabs != null) {
                    tabs = c.tabs;
                }
                if (c.parent_content_types != null) {
                    for (var i = 0; i < c.parent_content_types.length; i++) {
                        if (c.parent_content_types[i].tabs != null) {
                            tabs = tabs.concat(c.parent_content_types[i].tabs).unique();
                        }
                    }
                }
                if (c.composite_content_types != null) {
                    for (var i = 0; i < c.composite_content_types.length; i++) {
                        if (c.composite_content_types[i].tabs != null) {
                            tabs = tabs.concat(c.composite_content_types[i].tabs).unique();
                        }
                    }
                }
                console.log(tabs);
                $scope.tabs = tabs;
                $scope.currentTab = tabs[0].name;
            });
            $scope.data = {
                content_type: ct
            }
        }
        if ($scope.stateParams.parent_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["parent_id"] = parseInt($scope.stateParams.parent_id);
            } else {
                $scope.data = {
                    parent_id: parseInt($scope.stateParams.parent_id)
                }
            }
        }
        if ($scope.stateParams.content_type_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["content_type_id"] = parseInt($scope.stateParams.content_type_id);
            } else {
                $scope.data = {
                    content_type_id: parseInt($scope.stateParams.content_type_id)
                }
            }
        }
        if ($scope.stateParams.type_id) {
            if (typeof $scope.data !== 'undefined') {
                $scope.data["type_id"] = parseInt($scope.stateParams.type_id);
            } else {
                $scope.data = {
                    type_id: parseInt($scope.stateParams.type_id)
                }
            }
        }
        $scope.originalData = angular.copy(data);
        $scope.latestData = angular.copy(data);
    }
    
    
    // $scope.filesChanged = function(elm){
    //   $scope.files=elm.files
    //   $scope.$apply();
    //   console.log("mediaControllers scope: ")
    // console.log($scope.files);
    // }
    // $scope.files = [];
    // $scope.persistedFiles = [pathToUrl("media\\Sample Images\\TXT\\pic04.jpg")];

    $scope.test = {
        files: undefined
    }

    // $scope.submit = function() {
    //     $scope.$emit("formSubmit"); 
    // }

    $scope.submit = function() {
        console.log("submit")

        function success(response) {
            console.log("success", response)
            var escapedPath = replaceAll($scope.location, '\\', '%5C');
            
            $scope.$broadcast("formSubmitSuccess");  
            // if ($scope.files.length > 0) {
            //     $scope.upload(escapedPath);
            // }
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
            Content.update({
                id: $stateParams.id
            }, $scope.data, success, failure);
            console.log($scope.data)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            Content.create($scope.data, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
}
/**
 * @ngdoc controller
 * @name MediaDeleteCtrl
 * @function
 * @description
 * The controller for deleting media
 */
function MediaDeleteCtrl($scope, $stateParams, Content) {
    $scope.delete = function(item) {
        console.log(item)
        Content.delete({
            id: item.id
        }, function() {
            console.log("content record with id: " + item.id + " deleted")
        })
    };
}