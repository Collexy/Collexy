angular.module("myApp").controller("DirectoryTreeCtrl", DirectoryTreeCtrl);
angular.module("myApp").controller("DirectoryTreeCtrlEdit", DirectoryTreeCtrlEdit);

function DirectoryTreeCtrl($scope, $stateParams, $state, Directory, ngDialog) {
    //$scope.rootdir = $state.current.data.rootdir;
    $scope.rootdir = $state.current.name.split(".")[1] + "s";
    //alert(rootdir);
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
    $scope.clickToOpen = function(url) {
        ngDialog.open({
            template: url,
            scope: $scope
        });
    };
    $scope.deleteNode = function(item) {
        //alert("deleteNode")
        Directory.delete({
            'rootdir': rootdir,
            name: item.info.name
        }, function() {
            console.log("content and node record deleted with nodeId: " + item.entity.node.id)
        })
    };
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
        $oLay.css(overLayCSS)
    }
}

function DirectoryTreeCtrlEdit($scope, $stateParams, Directory, $state) {
    //console.log($state.current)
    $scope.rootdir = $state.current.name.split(".")[1] + "s";
    if ($scope.rootdir == 'stylesheets') {
        $scope.editorOptions = {
            lineWrapping: true,
            lineNumbers: true,
            readOnly: 'nocursor',
            mode: 'css',
        };
    } else {
        $scope.editorOptions = {
            lineWrapping: true,
            lineNumbers: true,
            readOnly: 'nocursor',
            mode: 'javascript',
        };
    }
    //alert(rootdir);
    //$scope.rootdir = $state.current.data.rootdir;
    //alert(rootdir)
    //$scope.currentTab = $scope.rootdir
    //$scope.stateParams = $stateParams;
    if ($stateParams.name) {
        $scope.data = Directory.get({
            rootdir: $scope.rootdir,
            name: $stateParams.name
        }, function(node) {
            $scope.data.old_path = node.path;
            console.log(node)
            if (node.info.is_dir) {
                $scope.type = 'folder'
                $scope.currentTab = $scope.type;
            } else {
                //alert($scope.type)
                $scope.type = 'file'
                $scope.currentTab = $scope.type;
            }
        });
        //User.get({ userId: $stateParams.userId} , function(phone) {
    }
    if (!$scope.data) {
        $scope.data = {
            "info": {}
        }
    }
    // if(!$scope.data.type) {
    //   $scope.type = 'file';
    // }
    if ($stateParams.type == 'folder') {
        $scope.data.info.is_dir = true;
        $scope.type = 'folder';
        //$scope.currentTab = $scope.type;
    } else if ($stateParams.type == 'file') {
        $scope.data.info.is_dir = false;
        $scope.type = 'file';
        $scope.currentTab = $scope.type;
    }
    if ($stateParams.parent) {
        $scope.data.parent = $stateParams.parent;
        if (!$scope.data.info.name) {
            $scope.data.path = $scope.data.parent + "\\" + $scope.data.info.name;
        }
    }
    //alert( $scope.type);
    $scope.currentTab = $scope.type;
    $scope.toggleTab = function(item, $event) {
        $scope.currentTab = item;
    }
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
        //console.log($stateParams)
        if ($stateParams.type) {
            console.log("create");
            console.log($scope.data)
            Directory.create({
                rootdir: $scope.rootdir
            }, $scope.data, success, failure);
            //User.create($scope.user, success, failure);
        } else {
            console.log("update");
            Directory.update({
                rootdir: $scope.rootdir,
                name: $stateParams.name
            }, $scope.data, success, failure);
            console.log($scope.data)
            //User.update($scope.user, success, failure);
        }
    }
}