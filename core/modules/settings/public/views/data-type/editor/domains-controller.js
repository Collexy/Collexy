

	angular.module("myApp").controller("Collexy.DataTypeEditor.Domains.Controller", CollexyDataTypeEditorDomains);
	
	function CollexyDataTypeEditorDomains($scope){
		if(typeof $scope.data.meta == 'undefined'){
			$scope.data["meta"] = {}
		}
		if(typeof $scope.data.meta["domains"] == 'undefined'){
			$scope.data.meta["domains"] = [];
		}
		
		$scope.addDomain = function(){
			$scope.data.meta["domains"].push($scope.domainToAdd);
			$scope.domainToAdd = "";
		}
		$scope.removeDomain = function(domain){
			var pos = $scope.data.meta["domains"].indexOf(domain);
			if(pos > -1){
				$scope.data.meta.domains.splice(pos,1);
			}
		}
	}