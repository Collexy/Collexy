//Registering a controller after app bootstrap
// $controllerProviderRef.register('Collexy.DataTypeEditor.ContentPicker', function($scope, ContentType)
// {
// 	alert("lol")
// 	$scope.contentTypes = ContentType.query();
// });
// alert("lol")
angular.module("myApp").controller("Collexy.DataTypeEditor.RadioButtonList.Controller", CollexyDataTypeEditorRadioButtonListController);
angular.module("myApp").controller("Collexy.DataTypePropertyEditor.RadioButtonList.Controller", CollexyDataTypePropertyEditorRadioButtonListController);

function CollexyDataTypeEditorRadioButtonListController($scope, ContentType) {
    $scope.add = function() {
        if (typeof $scope.entity.meta.options !== 'undefined') {
            var optionsArr = $scope.entity.meta.options;
            optionsArr.push({
                "value": "value here",
                "label": "label here"
            })
            $scope.entity.meta.options = optionsArr;
        }
    }
    $scope.remove = function(option) {
        $scope.entity.meta.options = $scope.entity.meta.options.filter(function(el) {
            return el !== option;
        });
    }
    // // alert("lol")
    // $scope.contentTypes = ContentType.query();
    // $scope.convertToInt = function(id){
    //     return parseInt(id, 10);
    // };
}

function CollexyDataTypePropertyEditorRadioButtonListController($scope, Content) {
    // //console.log($scope.data.meta[$scope.prop.name][0])
    // //console.log($scope.tabs)
    var dataType = null;
    for (var i = 0; i < $scope.tabs.length; i++) {
        for (var j = 0; j < $scope.tabs[i].properties.length; j++) {
            if ($scope.tabs[i].properties[j].name == $scope.prop.name) {
                dataType = $scope.tabs[i].properties[j].data_type;
            }
        }
    }
    $scope.dataType = dataType;
    // Content.query({
    //        'type-id': '1',
    //        'content-type': dataType.meta.content_type_id
    //        //'content-type': $scope.data.meta[$scope.prop.name].data_type.meta.content_type_id
    //    }, {}, function(contentNodes) {
    //        //var parentControllerScope = angular.element.controller().parent().scope();
    //        //console.log(parentControllerScope)
    //        $scope.contentNodes = contentNodes;
    //    });
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