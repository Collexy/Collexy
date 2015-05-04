//Registering a controller after app bootstrap
	// $controllerProviderRef.register('Collexy.DataTypeEditor.ContentPicker', function($scope, ContentType)
	// {
	// 	alert("lol")
	// 	$scope.contentTypes = ContentType.query();
	// });
// alert("lol")

	angular.module("myApp").controller("Collexy.DataTypeEditor.ContentPicker", CollexyDataTypeEditorContentPicker);
	angular.module("myApp").controller("Collexy.DataTypePropertyEditor.ContentPicker", CollexyDataTypePropertyEditorContentPicker);

	function CollexyDataTypeEditorContentPicker($scope, ContentType){
		// alert("lol")
		$scope.contentTypes = ContentType.query();

		$scope.convertToInt = function(id){
		    return parseInt(id, 10);
		};
	}

	function CollexyDataTypePropertyEditorContentPicker($scope, Content){
		//console.log($scope.data.meta[$scope.prop.name][0])
		//console.log($scope.tabs)
		var dataType = null;
		for(var i = 0; i < $scope.tabs.length; i++){
			for(var j = 0; j < $scope.tabs[i].properties.length; j++){
				if($scope.tabs[i].properties[j].name == $scope.prop.name){
					dataType = $scope.tabs[i].properties[j].data_type;
				}
			}
		}
		Content.query({
	        'type-id': '1',
	        'content-type': dataType.meta.content_type_id
	        //'content-type': $scope.data.meta[$scope.prop.name].data_type.meta.content_type_id
	    }, {}, function(contentNodes) {
	        //var parentControllerScope = angular.element.controller().parent().scope();
	        //console.log(parentControllerScope)
	        $scope.contentNodes = contentNodes;
	    });
	}


	

// 	// var queueLen = angular.module('myApp')._invokeQueue.length;
// 	// // Register the controls/directives/services we just loaded
// 	// var queue = angular.module('myApp')._invokeQueue;
// 	// for(var i=queueLen;i<queue.length;i++) {
// 	//     var call = queue[i];
// 	//     // call is in the form [providerName, providerFunc, providerArguments]
// 	//     var provider = $controllerProviderRef;
// 	//     if(provider) {
// 	//     	alert(provider)
// 	//         // e.g. $controllerProvider.register("Ctrl", function() { ... })
// 	//         provider[call[1]].apply(provider, call[2]);
// 	//     }
// 	// }
// $('body').injector().invoke(function($compile, $rootScope) {
//     $compile($('#ctrl'))($rootScope);
//     $rootScope.$apply();
// });