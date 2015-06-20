angular.module("myApp").controller("MediaTypeTreeCtrl", MediaTypeTreeCtrl);
angular.module("myApp").controller("MediaTypeEditCtrl", MediaTypeEditCtrl);
angular.module("myApp").controller("MediaTypeDeleteCtrl", MediaTypeDeleteCtrl);
/**
 * @ngdoc controller
 * @name MediaTypeTreeCtrl
 * @function
 * @description
 * The controller for the media type tree
 */
function MediaTypeTreeCtrl($scope, $stateParams, MediaTypeChildren, MediaType, sessionService, $interpolate, ngDialog) {
    $scope.ContextMenuServiceName = "MediaTypeContextMenu"
    $scope.EntityChildrenServiceName = "MediaTypeChildren"
    MediaType.query({
        'levels': '1'
    }, {}, function(tree) {
        $scope.tree = tree;
    });
}
/**
 * @ngdoc controller
 * @name MediaTypeEditCtrl
 * @function
 * @description
 * The controller for editing a media type
 */
function MediaTypeEditCtrl($scope, $stateParams, MediaType, DataType, MIMEType, $interval) {
    console.log($scope.node)
    $scope.currentTab = 'media-type';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        $scope.node = MediaType.get({
            extended: true
        }, {
            id: $stateParams.id
        }, function(node) {
            console.log(node);
            MIMEType.query({}, function(){}).$promise.then(function(mime_types){
                console.log(mime_types)
                //$scope.mime_types = mime_types;
                $scope.allMIMETypes = mime_types;
                var availableMIMETypes = [];
                var selectedMIMETypes = [];
                if (typeof node.id != 'undefined') {
                    for (var i = 0; i < mime_types.length; i++) {
                        //var mime_type_name = mime_types[i].name
                        //alert(mime_type_name)
                        if (typeof mime_types[i].media_type_id == 'undefined') {
                            
                                availableMIMETypes.push(mime_types[i])
                            
                            
                        } else {
                            
                            if(mime_types[i].media_type_id == node.id){
                                
                                selectedMIMETypes.push(mime_types[i])
                            }
                            
                        }
                    }
                }
                availableMIMETypes.unique();
                selectedMIMETypes.unique();
                $scope.availableMIMETypes = availableMIMETypes;
                $scope.selectedMIMETypes = selectedMIMETypes;
                if (selectedMIMETypes.length == 0) {
                    $scope.availableMIMETypes = mime_types;
                }
            })
        });

        
    } else {

        if ($stateParams.parent_id) {
            $scope.node = {
                "parent_id": parseInt($stateParams.parent_id),
                "created_by": $scope.userSession.id
            }
        } else {
            $scope.entity = {
                "created_by": $scope.userSession.id
            }
        }

        MIMEType.query({}, function() {}).$promise.then(function(mime_types) {
            $scope.allMIMETypes = mime_types;
            var availableMIMETypes = [];
            var selectedMIMETypes = [];
            availableMIMETypes.unique();
            selectedMIMETypes.unique();
            $scope.availableMIMETypes = availableMIMETypes;
            $scope.selectedMIMETypes = selectedMIMETypes;
            if (selectedMIMETypes.length == 0) {
                $scope.availableMIMETypes = mime_types;
            }
        }, function() {
            // error
        })
    }
    if ($scope.stateParams.type_id) {
        if (typeof $scope.node !== 'undefined') {
            $scope.node["type_id"] = parseInt($scope.stateParams.type_id);
        } else {
            $scope["node"] = {
                type_id: parseInt($scope.stateParams.type_id)
            }
        }
    }
    MediaType.query({
        'type-id': '2'
    }, {}, function(allMediaTypes) {
        $scope.allMediaTypes = allMediaTypes
        var availableCompositeMediaTypes = []
        for (var i = 0; i < allMediaTypes.length; i++) {
            if (typeof $scope.node.parent_media_types !== 'undefined' && $scope.node.parent_media_types.length > 0) {
                if ($scope.node.parent_media_types.containsId(allMediaTypes[i].id)) {} else {
                    availableCompositeMediaTypes.push(allMediaTypes[i])
                }
            } else {
                if(allMediaTypes[i].id != $stateParams.id){
                    availableCompositeMediaTypes.push(allMediaTypes[i])
                }
                
            }
        }
        $scope.availableCompositeMediaTypes = availableCompositeMediaTypes;
        console.log(availableCompositeMediaTypes)
    });
    $scope.allDataTypes = DataType.query({}, {}, function(node) {});
    console.log($scope.allDataTypes)

    $scope.moveMIMEType = function(item, from, to) {
        //alert("moveMember")
        var idx = from.indexOf(item);
        if (idx != -1) {
            from.splice(idx, 1);
            to.push(item);
        }
        var mime_types = [];
        for (var i = 0; i < $scope.selectedMIMETypes.length; i++) {
            mime_types.push($scope.selectedMIMETypes[i]);
        }
        $scope.media_type_mime_types = mime_types;

    };


    $scope.submit = function() {
        console.log("submit")

        function success(response) {
            console.log("success", response)
            //$location.path("/admin/users");
            var initialMimeTypesForMediaType = [];
            for (var i = 0; i < $scope.allMIMETypes.length; i++) {
                if($scope.allMIMETypes[i].media_type_id == $scope.node.id){
                    initialMimeTypesForMediaType.push($scope.allMIMETypes[i])
                }
            };

            var removed = []
            for (var i = 0; i < initialMimeTypesForMediaType.length; i++) {
                var found = false;
                for (var j = 0; j < $scope.media_type_mime_types.length; j++) {
                    if(initialMimeTypesForMediaType[i].id == $scope.media_type_mime_types[j].id){
                        found = true;
                        break;
                    }
                };
                if(!found){
                    var obj = initialMimeTypesForMediaType[i];
                    delete obj.media_type_id
                    removed.push(obj)
                }
            };

            var added = []
            for (var i = 0; i < $scope.media_type_mime_types.length; i++) {
                var found = false;
                for (var j = 0; j < initialMimeTypesForMediaType.length; j++) {
                    if($scope.media_type_mime_types[i].id == initialMimeTypesForMediaType[j].id){
                        found = true;
                        break;
                    }
                };
                if(!found){
                    var obj = $scope.media_type_mime_types[i];
                    obj["media_type_id"] = $scope.node.id;
                    added.push(obj)
                }
            };


            var allMIMETypesCombined = initialMimeTypesForMediaType.concat(removed)
            var allMIMETypesCombined = allMIMETypesCombined.concat(added)
            // allMIMETypes.unique();
            // MAKE PUT REQUEST TO MIMEType service
            var i = 0;

            $interval(function() {
                MIMEType.update({
                    id: allMIMETypesCombined[i].id
                }, allMIMETypesCombined[i], function(){}, function(){});
                i++;
            }, 1000, allMIMETypesCombined.length)
            
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
            MediaType.update({
                id: $stateParams.id
            }, $scope.node, success, failure);
            console.log($scope.node)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            MediaType.create($scope.node, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
    $scope.addTab = function() {
        if ('tabs' in $scope.node) {} else {
            $scope.node["tabs"] = [];
        }
        tab = {
            "name": "mytab",
            "properties": []
        }
        $scope.node.tabs.push(tab);
    }
    $scope.addProp = function(tab) {
        if ('tabs' in $scope.node) {
            var tabs = $scope.node.tabs;
            if (tabs.length > 0) {
                for (var i = 0; i < tabs.length; i++) {
                    if (tabs[i].name == tab) {
                        if ('properties' in tabs[i]) {} else {
                            tabs[i].properties = [];
                        }
                        tabs[i].properties.push({
                            "name": "property name",
                            "order": 1,
                            "data_type_id": 2,
                            "help_text": "prop help text",
                            "description": "prop description"
                        });
                    }
                }
                $scope.node.tabs = tabs;
            }
        }
    }
    
    // $scope.$watch("test", function(newValue, oldValue) {
    //             $scope.test = newValue;
    //             console.log(newValue)
    //         }, true);
}
/**
 * @ngdoc controller
 * @name MediaTypeDeleteCtrl
 * @function
 * @description
 * The controller for deleting media type
 */
function MediaTypeDeleteCtrl($scope, $stateParams, MediaType) {
    $scope.delete = function(item) {
        console.log(item)
        MediaType.delete({
            id: item.id
        }, function() {
            console.log("media type record with id: " + item.id + " deleted")
        })
    };
}