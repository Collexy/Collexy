angular.module("myApp").controller("MimeTypeTreeCtrl", MimeTypeTreeCtrl);
angular.module("myApp").controller("MimeTypeEditCtrl", MimeTypeEditCtrl);
angular.module("myApp").controller("MimeTypeDeleteCtrl", MimeTypeDeleteCtrl);
/**
 * @ngdoc controller
 * @name MimeTypeTreeCtrl
 * @function
 * @description
 * The controller for the data type tree
 */
function MimeTypeTreeCtrl($scope, MIMEType) {
    $scope.ContextMenuServiceName = "MimeTypeContextMenu"
    $scope.tree = MIMEType.query();
    console.log($scope.tree)
}
/**
 * @ngdoc controller
 * @name MimeTypeEditCtrl
 * @function
 * @description
 * The controller for editing a data type
 */
function MimeTypeEditCtrl($scope, $stateParams, MIMEType, $compile, MediaType) {
   
    $scope.currentTab = 'mime-type';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        $scope.entity = MIMEType.get({
            id: $stateParams.id
        }, function(entity) {
            $scope.allMediaTypes = MediaType.query();
            // MIMEType.query(function(allMIMETypes){
            //     MediaType.query(function(allMediaTypes){
            //         $scope.allMediaTypes = allMediaTypes;

            //         var unavailableMediaTypeIds = [];
            //         for (var i = 0; i < allMIMETypes.length; i++) {
            //             if(typeof allMIMETypes[i].media_type_id != 'undefined'){
            //                 unavailableMediaTypeIds.push(allMIMETypes[i].media_type_id)
            //             }
            //         }

                    
            //         var availableMediaTypes = [];
            //         for (var i = 0; i < $scope.allMediaTypes.length; i++) {
            //             if(unavailableMediaTypeIds.indexOf($scope.allMediaTypes[i].id) == -1){
            //                 availableMediaTypes.push($scope.allMediaTypes[i])
            //             }
            //             if($scope.allMediaTypes[i].id == entity.media_type_id){
            //                 availableMediaTypes.push($scope.allMediaTypes[i])
            //             }
            //         }
            //         $scope.availableMediaTypes = availableMediaTypes;

            //         console.log($scope.allMediaTypes)
            //         console.log($scope.availableMediaTypes)
            //     });
            // })
        });
    } else {
        $scope.entity = {}
    }

    

    $scope.tree = MIMEType.query();
    
    $scope.submit = function() {
        console.log("submit")

        function success(response) {
            console.log("success", response)
            //$location.path("/admin/users");
        }

        function failure(response) {
            console.log("failure", response);
        }
        if ($stateParams.id) {
            console.log("update");
            MIMEType.update({
                id: $stateParams.id
            }, $scope.entity, success, failure);
            console.log($scope.entity)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            console.log($scope.entity)
            MIMEType.create($scope.entity, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
}
/**
 * @ngdoc controller
 * @name MimeTypeDeleteCtrl
 * @function
 * @description
 * The controller for deleting data type
 */
function MimeTypeDeleteCtrl($scope, $stateParams, MIMEType) {
    $scope.delete = function(item) {
        console.log(item)
        MIMEType.delete({
            id: item.id
        }, function() {
            console.log("data type record with id: " + item.id + " deleted")
        })
    };
}