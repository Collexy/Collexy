angular.module("myApp").controller("Collexy.DataTypeEditor.FileUpload.Controller", CollexyDataTypeEditorFileUploadController);
// angular.module("myApp").controller("Collexy.DataTypePropertyEditor.ContentPicker", CollexyDataTypePropertyEditorContentPicker);

function CollexyDataTypeEditorFileUploadController($scope, ContentType, $http, FileService) {
    // $scope.files = [];
    //$scope.persistedFiles = [pathToUrl("media\\Sample Images\\TXT\\pic04.jpg")];
    //console.log($scope.data.meta["file_upload"]["persisted_files"])
    // $scope.$watch("files", function(newValue, oldValue) {
    //     $scope.files = newValue;
    //     //console.log($scope.files)
    //     // var lol = [];
    //     // for(var i = 0; i < $scope.files.length; i++){
    //     // 	lol.push($scope.files[i].name)
    //     // }
    //     // $scope.data.meta["file_upload"].persisted_files = lol
    // }, true);

	// $scope.$watch("data", function(newValue, oldValue) {
	// 	$scope.data = newValue;
	// 	console.log($scope.data)
 //    }, true);

    $scope.$watch("clearFiles", function(newValue, oldValue) {
    	if (newValue == true) {
            clearFiles(true);
        } else {
        	clearFiles(false);
        }
        $scope.clearFiles = newValue;
        console.log($scope.clearFiles)
    }, true);

    function clearFiles(isTrue) {
    	//alert(isTrue)
    	if(isTrue){
    		//clear the current files
        	$scope.files = [];
        	// $scope.data.meta["attached_file"] = [];
            delete $scope.data.meta["attached_file"];
    	} else {
            if(typeof $scope.originalData != 'undefined'){
                if(typeof $scope.originalData["meta"] != 'undefined'){
                    if(typeof $scope.originalData["meta"]["attached_file"] != 'undefined'){
                        $scope.data["meta"]["attached_file"] = $scope.originalData["meta"]["attached_file"];
                    }
                }
            }
            // if(typeof $scope.data.meta["file_upload"] == 'undefined'){
            //     $scope.data.meta["file_upload"] = {};
            // }
            // if(typeof $scope.originalData.meta["file_upload"] != 'undefined'){
            //     $scope.data.meta["file_upload"]["persisted_files"] = $scope.originalData.meta["file_upload"].persisted_files;
            // }
    		
    	}
        
        // $scope.data.meta["file_upload"].persisted_files = [];
    }

    // $scope.$on("formSubmitSuccess", function (event, args) {
    // 	// alert("formSubmitSuccess event")
    // 	var escapedPath = replaceAll($scope.location, '\\', '%5C');
    // 	if(typeof $scope.files != 'undefined'){
    // 		if($scope.files.length > 0){
    // 			$scope.upload(escapedPath);

    // 			if(typeof $scope.originalData.meta["file_upload"].persisted_files != undefined){
		  //   		if($scope.originalData.meta["file_upload"].persisted_files.length > 0){
		  //   			$scope.deleteFiles($scope.location, $scope.originalData.meta["file_upload"].persisted_files);
		  //   		}
		  //   	}
    // 		} else {
    // 			if(typeof $scope.originalData.meta["file_upload"].persisted_files != undefined){
		  //   		if($scope.originalData.meta["file_upload"].persisted_files.length > 0){
		  //   			if($scope.clearFiles){
		  //   				$scope.deleteFiles($scope.location, $scope.originalData.meta["file_upload"].persisted_files);
		  //   			}
		  //   		}
		  //   	}
    // 		}
    // 	} else {
    // 		if(typeof $scope.originalData.meta["file_upload"].persisted_files != undefined){
	   //  		if($scope.originalData.meta["file_upload"].persisted_files.length > 0){
	   //  			if($scope.clearFiles){
	   //  				$scope.deleteFiles($scope.location, $scope.originalData.meta["file_upload"].persisted_files);
	   //  			}
	   //  		}
	   //  	}
    // 	}
    	

    // 	// if(typeof $scope.data.meta != 'undefined'){
    // 	// 	if(typeof $scope.data.meta["file_upload"] != 'undefined'){
    // 	// 		if(typeof $scope.data.meta["file_upload"].persisted_files != 'undefined'){
    // 	// 			if($scope.data.meta["file_upload"].persisted_files.length > 0){
    // 	// 				for(var i = 0; i < $scope.data.meta["file_upload"].persisted_files.length; i++){
    // 	// 					var isSameAsOrig = false;
    // 	// 					if($scope.data.meta["file_upload"].persisted_files[i] == $scope.originalData.meta["file_upload"].persisted_files[i]){
    // 	// 						var isSameAsOrig = false;
    // 	// 					}
    // 	// 					if(!isSameAsOrig){

    // 	// 					}
    // 	// 				}
    // 	// 			}
    // 	// 		}
    // 	// 	}
    // 	// }
    	
    	
    //     // $scope.$apply(function () {
        	
    //     // })
    // });

	$scope.deleteFile = function(location, fileToDelete){
		alert("deleteFile() fired");
		console.log(fileToDelete)
			FileService.delete({
				path: location,
				fileName: fileToDelete.name
			}, function(){
				console.log("Location: " + location + " with filename: " + fileToDelete.name + " has successfully been deleted");
			})
		// for(var i = 0; filesArray.length; i++){
		// 	FileService.delete({
		// 		path: escapedPath,
		// 		fileName: filesArray[i]
		// 	}, function(){
		// 		console.log("File: " + escapedPath + " with filename: " + filesArray[i] + " has successfully been deleted");
		// 	})
		// }
	}

    $scope.upload = function(escapedPath) {
        console.log("$scope.upload")
        if(typeof escapedPath == 'undefined'){
            escapedPath = replaceAll($scope.location, '\\', '%5C');
        } 
        var fd = new FormData() // put these 3 lines in a service
        angular.forEach($scope.files, function(file) {
            console.log("angular.forEach")
            console.log(file)
            fd.append('file', file)
        })
        $http.post('/api/directory/upload-file-test?path=' + escapedPath, fd, {
            transformRequest: angular.identity, // returns first argument it is passed
            headers: {
                'Content-Type': undefined
            } //multipart/form-data
        }).success(function(d) {
            console.log("File(s) successfully uploaded")
        })
    }

  //   $scope.readAsDataURL = function(file){
  //   	var reader = new FileReader()
  //   	reader.onload = function(){
  //   		alert(this.result)
  //   		return this.result;
		// };
  //   	reader.readAsDataURL(file)
    	

  //   }
    
 	// $scope.$on("filesSelected", function (event, args) {
  //       $scope.$apply(function () {
  //       	console.log("lol")
  //       	console.log(args.files)
  //       	$scope.files = args.files;
  //       	var files = [];
  //           for(var i = 0; i< args.files.length; i++){
  //           	//$scope.files.push({ alias: $scope.data.alias, file: args.files[i] });
  //           	files.push(args.files[i].name)
  //           }
  //           $scope.data.meta["file_upload"]["persisted_files"] = files;
  //           $scope.latestData = $scope.data;
  //       })
  //   });


    // $scope.$on("filesSelected", function (event, args) {
    //     $scope.$apply(function () {
    //     	console.log("lol")
    //     	console.log(args.files)
    //     	//set the files collection
    //         fileManager.setFiles($scope.data.alias, args.files);
    //         var files = [];
    //         for(var i = 0; i< args.files.length; i++){
    //         	//$scope.files.push({ alias: $scope.data.alias, file: args.files[i] });
    //         	files.push(args.files[i].name)
    //         }
    //         $scope.data.meta["persisted_files"] = files;
    //     })
    // });
}

// function CollexyDataTypePropertyEditorContentPicker($scope, Content) {
//     //console.log($scope.data.meta[$scope.prop.name][0])
//     //console.log($scope.tabs)
//     var dataType = null;
//     for (var i = 0; i < $scope.tabs.length; i++) {
//         for (var j = 0; j < $scope.tabs[i].properties.length; j++) {
//             if ($scope.tabs[i].properties[j].name == $scope.prop.name) {
//                 dataType = $scope.tabs[i].properties[j].data_type;
//             }
//         }
//     }
//     Content.query({
//         'type-id': '1',
//         'content-type': dataType.meta.content_type_id
//         //'content-type': $scope.data.meta[$scope.prop.name].data_type.meta.content_type_id
//     }, {}, function(contentNodes) {
//         //var parentControllerScope = angular.element.controller().parent().scope();
//         //console.log(parentControllerScope)
//         $scope.contentNodes = contentNodes;
//     });
// }