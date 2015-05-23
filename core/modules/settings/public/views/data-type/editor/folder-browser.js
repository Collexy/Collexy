

// angular.module("myApp").controller("Collexy.DataTypeEditor.MediaPicker", CollexyDataTypeEditorMediaPicker);
angular.module("myApp").controller("Collexy.DataTypePropertyEditor.FolderBrowser", CollexyDataTypePropertyEditorFolderBrowser);

// function CollexyDataTypeEditorMediaPicker($scope, MediaType){
// 	// alert("lol")
// 	$scope.mediaTypes = MediaType.query();

// 	$scope.convertToInt = function(id){
// 	    return parseInt(id, 10);
// 	};
// }

function CollexyDataTypePropertyEditorFolderBrowser($scope, Media, MediaChildren, $stateParams){
	//console.log($scope.data.meta[$scope.prop.name][0])
	//console.log($scope.tabs)
	// var dataType = null;
	// for(var i = 0; i < $scope.tabs.length; i++){
	// 	for(var j = 0; j < $scope.tabs[i].properties.length; j++){
	// 		if($scope.tabs[i].properties[j].name == $scope.prop.name){
	// 			dataType = $scope.tabs[i].properties[j].data_type;
	// 		}
	// 	}
	// }
	$scope.folder = {}

	if($stateParams.id){
		MediaChildren.query({
        'id': $stateParams.id
	        //'media-type': $scope.data.meta[$scope.prop.name].data_type.meta.media_type_id
	    }, {}, function(children) {
	        //var parentControllerScope = angular.element.controller().parent().scope();
	        //console.log(parentControllerScope)
	        $scope.folder["children"] = children;
	        if(typeof $scope.folder["children"] != 'undefined'){
	        	for(var i = 0; i < $scope.folder["children"].length; i++){
	        		
	        		if(typeof $scope.folder["children"][i].media_type.tabs != 'undefined'){
				    	for(var j = 0; j < $scope.folder["children"][i].media_type.tabs.length; j++){

				    		if(typeof $scope.folder["children"][i].media_type.tabs[j].properties != 'undefined'){
								for(var k = 0; k < $scope.folder["children"][i].media_type.tabs[j].properties.length; k++){

									if($scope.folder["children"][i].media_type.tabs[j].properties[k].data_type_id == 16){
										//dataType = $scope.tabs[j].properties[k].data_type;

										$scope.folder["children"][i]["has_upload"] = true;
									}
								}
							}
						}
					}
			    }
	        }
	        
	    });


	}
	
}