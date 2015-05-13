

// angular.module("myApp").controller("Collexy.DataTypeEditor.ContentPicker", CollexyDataTypeEditorContentPicker);
angular.module("myApp").controller("Collexy.DataTypePropertyEditor.FolderBrowser", CollexyDataTypePropertyEditorFolderBrowser);

// function CollexyDataTypeEditorContentPicker($scope, ContentType){
// 	// alert("lol")
// 	$scope.contentTypes = ContentType.query();

// 	$scope.convertToInt = function(id){
// 	    return parseInt(id, 10);
// 	};
// }

function CollexyDataTypePropertyEditorFolderBrowser($scope, Content, ContentChildren, $stateParams){
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
		ContentChildren.query({
        'id': $stateParams.id
	        //'content-type': $scope.data.meta[$scope.prop.name].data_type.meta.content_type_id
	    }, {}, function(children) {
	        //var parentControllerScope = angular.element.controller().parent().scope();
	        //console.log(parentControllerScope)
	        $scope.folder["children"] = children;
	        if(typeof $scope.folder["children"] != 'undefined'){
	        	for(var i = 0; i < $scope.folder["children"].length; i++){
	        		
	        		if(typeof $scope.folder["children"][i].content_type.tabs != 'undefined'){
				    	for(var j = 0; j < $scope.folder["children"][i].content_type.tabs.length; j++){

				    		if(typeof $scope.folder["children"][i].content_type.tabs[j].properties != 'undefined'){
								for(var k = 0; k < $scope.folder["children"][i].content_type.tabs[j].properties.length; k++){

									if($scope.folder["children"][i].content_type.tabs[j].properties[k].data_type_id == 16){
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