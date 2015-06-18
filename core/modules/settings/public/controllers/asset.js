angular.module("myApp").controller("AssetTreeCtrl", AssetTreeCtrl);
angular.module("myApp").controller("AssetEditCtrl", AssetEditCtrl);
angular.module("myApp").controller("AssetDeleteCtrl", AssetDeleteCtrl);

function AssetTreeCtrl($scope, $stateParams, $state, Directory, DirectoryContextMenu) {
    //$scope.rootdir = $state.current.data.rootdir;
    $scope.rootdir = $state.current.name.split(".")[1] + "s";
    
    //alert($scope.rootdir);
    var directoryNodes;
    Directory.query({
        rootdir: $scope.rootdir
    }, function() {}).$promise.then(function(data) {
        directoryNodes = data;
        $scope.tree = directoryNodes;
        $scope.rootNode = data;
        console.log($scope.rootNode)
    }, function(err) {
        // ERROR
    });
    // $scope.clickToOpen = function(url) {
    //     ngDialog.open({
    //         template: url,
    //         scope: $scope
    //     });
    // };
    
    var offset = {
        // left: 40,
        // top: -80
        left: 0,
        top: -76
    }
    var $oLay = angular.element(document.getElementById('overlay'))
    $scope.showOptions = function(item, $event) {
        

        console.log("showoptions")
        var overlayDisplay;
        // if ($scope.currentItem === item){
        if ($oLay.css("display") == "block") {
            $scope.currentItem = null;
            overlayDisplay = 'none'
        } else {
            $scope.currentItem = item;
            overlayDisplay = 'block'
        }
        if (angular.element(document.getElementById('adminsubmenucontainer')).hasClass('expanded1')) {
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
        $scope.getEntityInfo(item)
        $oLay.css(overLayCSS)
        $event.preventDefault();
        $event.stopPropagation();
    }

    $scope.getEntityInfo = function(currentItem){
        console.log(currentItem)
        if (currentItem == undefined) {
            currentItem = {
                path: "root",
                info: { name: "root", is_dir: true}
            }
            //scope.currentItem = currentItem;
        }

        if(typeof currentItem.info.is_dir == 'undefined'){
            currentItem.info["is_dir"] = false;
        }
        
        DirectoryContextMenu.query({
            rootdir: $scope.rootdir,
            // name: addslashes(currentItem.path)
            name: currentItem.path,
            is_dir: currentItem.info.is_dir
        }, {}, function() {}).$promise.then(function(data) {
            console.log($scope.rootdir)
            //console.log(data)
            //var parentScope = element.parent().parent().parent().parent().parent().parent().parent().scope();
            //scope.$parent.$parent.$parent.$parent.$parent.contextMenu = data;
            //var s = angular.element(document.getElementsByClassName('outer-list-container')[0]).scope()
            $scope.contextMenu = data;
            //console.log(scope.$parent.$parent.$parent.$parent.$parent)
            //scope.currentItem = currentItem;
        });
    }
}

function AssetEditCtrl($scope, $stateParams, Directory, $state, Upload) {
    $scope.$watch('files', function () {
      console.log($scope.files[0])
        //$scope.upload($scope.files);
    });

    $scope.rootdir = $state.current.name.split(".")[1] + "s";
    
    if ($stateParams.name) {
        $scope.data = Directory.get({
            rootdir: $scope.rootdir,
            name: $stateParams.name
        }, function(data) {
            $scope.data.old_path = data.path;
            console.log(data)
            if (data.info.is_dir) {
                $scope.type = 'folder'
                $scope.currentTab = $scope.type
            } else {
                $scope.type = 'file'
                $scope.currentTab = $scope.type;

                var fileName = data.info.name;
                var fileNameArr = fileName.split(".")
                var fileExtension = fileNameArr[fileNameArr.length-1]
                var isImage = false;
                var codeMirrorMode = "";
                // should go into its own separate .json file
                var types = {
                        "img": ["jpg", "png", "gif", "svg"],
                        "code": {
                            "js" : "javascript",
                            "css" : "css"
                        }
                    };


                if(types["img"].indexOf(fileExtension) > -1){
                    isImage = true;
                }

                if(!isImage){
                    // file is not an image
                    for (var k in types["code"]){
                        if(k==fileExtension){
                            //alert(types["code"][k])
                            codeMirrorMode = types["code"][k];
                            $scope.editorOptions = {
                                lineWrapping: true,
                                lineNumbers: true,
                                //readOnly: 'nocursor',
                                mode: codeMirrorMode, // eg. 'css', 'javascript',
                                indentUnit: 4,
                                tabMode: 'spaces',
                            };
                            break;
                        }
                    }
                }
                $scope.isImage = isImage;
                $scope.fileExtension = fileExtension;

            }

        });
        //User.get({ userId: $stateParams.userId} , function(phone) {
    } else { 


        $scope.data = {
            "info": {}
        }

        if ($stateParams.type == 'folder') {
            $scope.data.info.is_dir = true;
            $scope.type = 'folder';
            //$scope.currentTab = $scope.type;
        } else if ($stateParams.type == 'file') {
            $scope.data.info.is_dir = false;
            $scope.type = 'file';
            $scope.editorOptions = {
                                lineWrapping: true,
                                lineNumbers: true,
                                //readOnly: 'nocursor',
                                mode: 'htmlmixed', // eg. 'css', 'javascript',
                                indentUnit: 4,
                                tabMode: 'spaces',
                            };
        }
        
        if ($stateParams.parent) {
            $scope.data.parent = $stateParams.parent;
            if (!$scope.data.info.name) {
                $scope.data.path = $scope.data.parent + "\\" + $scope.data.info.name;
            }
        }

        $scope.currentTab = $scope.type;

        // $scope.$watch("data", function(newValue, oldValue) {
        //     $scope.data = newValue;
        //     alert("data changed")
        // }, true);
        
        // $scope.$watch("files", function(newValue, oldValue) {
        //     //$scope.data.name = newValue[0].info.name;
        //     if(typeof newValue != 'undefined'){
        //         if(typeof newValue[0] != 'undefined'){
        //             if(typeof newValue[0].name != 'undefined'){
        //                 console.log(newValue[0].name)
        //                 $scope.data["info"]["name"] = newValue[0].name;
        //             }
        //         }
        //     }
            
        // }, true);
    }
    
    
    
    //alert( $scope.type);
    
    // $scope.toggleTab = function(item, $event) {
    //     $scope.currentTab = item;
    // }
    $scope.updateName = function(name) {
        $scope.data.path = $scope.data.parent + "\\" + name;
    }
    $scope.updateParentPath = function(name) {
        $scope.data.parent = name;
        $scope.data.path = $scope.data.parent + "\\" + $scope.data.info.name;
    }
    $scope.updateNewPath = function(name) {
        var path = $scope.data.path;
        var pathEnding = path.lastIndexOf('\\');
        //var currentFileFolderName = path.substring(pathEnding + 1);
        var currentFileFolderName = $scope.data.info.name;
        var pathBeginning = path.substring(0, pathEnding + 1);
        $scope.data.path = pathBeginning + currentFileFolderName;
    }

    $scope.pathToUrl = function (text) {
        text = text.replace(/\\/g, '/');
        return text;
    }
    $scope.submit = function() {
        //$scope.$broadcast("formSubmitSuccess"); 
        //$scope.files = ["lol123"]
        // alert($scope.files[0]) 
        // var files = $scope.files
        // for (var i = 0; i < files.length; i++) {
        //     console.log(files[i])
        // }
        $scope.upload($scope.files);


        // console.log("submit")

        // function success(response) {
        //     console.log("success", response)
        //     //$location.path("/admin/users");
        // }

        // function failure(response) {
        //     console.log("failure", response);

        // }
        // //console.log($stateParams)
        // if ($stateParams.type) {
        //     console.log("create");
        //     console.log($scope.data)
        //     Directory.create({
        //         rootdir: $scope.rootdir
        //     }, $scope.data, success, failure);
        //     //User.create($scope.user, success, failure);
        // } else {
        //     console.log("update");
        //     Directory.update({
        //         rootdir: $scope.rootdir,
        //         name: $stateParams.name
        //     }, $scope.data, success, failure);
        //     console.log($scope.data)
        //     //User.update($scope.user, success, failure);
        // }
    }

    // $scope.files = [];
    // $scope.$watch("files", function(newValue, oldValue) {
    //     $scope.files = newValue;
    //     console.log($scope.files)
    //     // var lol = [];
    //     // for(var i = 0; i < $scope.files.length; i++){
    //     //  lol.push($scope.files[i].name)
    //     // }
    //     // $scope.data.meta["file_upload"].persisted_files = lol
    // }, true);

    

    $scope.upload = function (files) {
        alert("lol")

        console.log($scope.files)
        if (files && files.length && !$stateParams.type) {
            for (var i = 0; i < files.length; i++) {
                var file = files[i];
                // var pathIndexEnd = $scope.data.path.lastIndexOf("\\")
                // var path = $scope.data.path.substring(0,pathIndexEnd)
                // alert(path)
                var escapedPath = replaceAll($scope.data.path, '\\', '%5C');
                console.log('file is ' + JSON.stringify(file));
                var uploadUrl = '/api/directory/upload-file-test?path=' + escapedPath;
                Upload.upload({
                    url: uploadUrl,
                    fields: {'username': $scope.username},
                    file: file
                }).progress(function (evt) {
                    var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
                    console.log('progress: ' + progressPercentage + '% ' + evt.config.file.name);
                }).success(function (data, status, headers, config) {
                    console.log('file ' + config.file.name + 'uploaded. Response: ' + data);
                });
            }
        } else {
            function success(response) {
                console.log("success", response)

                if (files && files.length && $stateParams.type) {
                    for (var i = 0; i < files.length; i++) {
                        var file = files[i];
                        // var pathIndexEnd = $scope.data.path.lastIndexOf("\\")
                        // var path = $scope.data.path.substring(0,pathIndexEnd)
                        // alert(path)
                        var escapedPath = replaceAll($scope.data.path, '\\', '%5C');
                        console.log('file is ' + JSON.stringify(file));
                        var uploadUrl = '/api/directory/upload-file-test?path=' + escapedPath;
                        Upload.upload({
                            url: uploadUrl,
                            fields: {'username': $scope.username},
                            file: file
                        }).progress(function (evt) {
                            var progressPercentage = parseInt(100.0 * evt.loaded / evt.total);
                            console.log('progress: ' + progressPercentage + '% ' + evt.config.file.name);
                        }).success(function (data, status, headers, config) {
                            console.log('file ' + config.file.name + 'uploaded. Response: ' + data);
                        });
                    }
                }
                //$location.path("/admin/users");
            }

            function failure(response) {
                console.log("failure", response);

            }
            //console.log($stateParams)
            if ($stateParams.type) {
                console.log("create");
                console.log($scope.data)
                Directory.create({},
                    $scope.data, success, failure);
                //User.create($scope.user, success, failure);
            } else {
                console.log("update");
                Directory.update({},
                    $scope.data, success, failure);
                console.log($scope.data)
                //User.update($scope.user, success, failure);
            }
        }
    };
    //$scope.persistedFiles = [pathToUrl("media\\Sample Images\\TXT\\pic04.jpg")];
    //console.log($scope.data.meta["file_upload"]["persisted_files"])
    

    
}
/**
 * @ngdoc controller
 * @name DirectoryDeleteCtrl
 * @function
 * @description
 * The controller for deleting files and file directories
 */
function AssetDeleteCtrl($scope, $stateParams, DirectoryDelete) {
    // console.log($scope.currentItem)
    $scope.delete = function(item) {
        
        // $scope.currentItem = item;
        // console.log(item)
        // var path = item.path
        // var path1 = path.split("\\")
        // var path2="";
        // var startAt = -1;
        // for (var i = 0; i < path1.length; i++) {
        //     if(startAt != -1){
        //         path2 = path2 + "\\" + path1[i]
        //     }
        //     if(path1[i]=='assets'){
        //         path2 = "assets";
        //         startAt = i;
        //     }
            
        // };
        // var url = pathToUrl(path2)
        // alert(url)

        DirectoryDelete.delete({
            "path" : item.path
        }, function() {
            console.log("directory file/folder with name: " + item.info.name + " deleted")
        })
    };
}