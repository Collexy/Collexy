var adminControllers = angular.module('adminControllers', []);

adminControllers.controller('AdminContentCtrl', ['$scope', '$interpolate', 'sessionService', '$state', function ($scope, $interpolate, sessionService, $state) {
	//$scope.users = User.query();
    $scope.state = $state;
    $scope.userSession = sessionService.userSession;
    $scope.mySessionService = sessionService;
    console.log("adminmenustrl")
    console.log(sessionService.userSession)

    $scope.$watch("mySessionService.getUser()", function(newValue, oldValue) {
        $scope.userSession = newValue;
        console.log($scope.userSession)
    },true);

	$scope.isDefined = function(obj, prop) {
        if(obj != null ||obj != undefined){
            if(prop in obj){
                //alert("obj: " + obj.name + ", prop: " + prop + ", :::: true")
                return true;
            } else {
                //alert("obj: " + obj.name + ", prop: " + prop + ", :::: false")
                return false;
            }
        } else {
            //alert("obj: " + obj.name + ", prop: " + prop + ", :::: false")
            return false;
        }
	    
	}

	$scope.userHasPermission = function(permissionsString){
        //alert(permissionsString)
        permissions = permissionsString.split(",");
        
        var permFound = false;
        var hasPermissions = false;

        var user = $scope.userSession;

        // First check if a the currently logged in user has specific permissions per user-level
        if($scope.isDefined(user,"permissions")){
            if(user.permissions.length > 0){

            }
        } else if($scope.isDefined(user,"user_groups")){ // If first check fails, check permissions for each group if any
            if(user.user_groups.length > 0){
                i_loop:
                for(var i=0; i < permissions.length; i++){
                    permFound = false;
                    j_loop:
                    for (var j=0; j < user.user_groups.length; j++){
                        
                        if(permFound){
                            break j_loop;
                        }
                        k_loop:
                        for (var k=0; k < user.user_groups[j].permissions.length; k++){
                            if(permFound){
                                break k_loop;
                            }

                            if(permissions[i] == user.user_groups[j].permissions[k]){
                                permFound = true;
                            }
                        }
                    }  
                }
            }
        }

        hasPermissions = permFound;
        //console.log(hasPermissions)
        //alert(hasPermissions)
        return hasPermissions;
    }

    // $scope.getContentByContentTypeId = function(id){}
}]);

adminControllers.controller('AdminMenuCtrl', ['$scope', '$state', 'AngularRoute', 'MenuLink', function ($scope, $state, AngularRoute, MenuLink) {
	
    //$scope.sections = AngularRoute.query({'type': '1'},{}, function(section){})
    $scope.sections = MenuLink.query({},{name: 'main'}, function(section){})
    $scope.currentSectionId = 0;
    $scope.toggleSubMenu = function(id){

    	//alert($scope.sections.length)
    	var subMenuItems = [];
    	for(var i = 0; i < $scope.sections.length; i++){
    		if('parent_id' in $scope.sections[i]){
    			if($scope.sections[i].parent_id == id){
	    			subMenuItems.push($scope.sections[i]);
	    		}
    		}
    		
    	}
    	$scope.subMenuItems = subMenuItems;

    	console.log($scope.subMenuItems)

    	if($scope.currentSectionId == 0){
    		$scope.currentSectionId = id;
    	}



    	if(angular.element('#adminsubmenucontainer').hasClass("collapse1")){
    		if($scope.subMenuItems.length>0){
    			if($scope.currentSectionId == id){
	    			angular.element('#adminsubmenucontainer').removeClass("collapse1");
	    			angular.element('#adminsubmenucontainer').addClass("expanded1");
                    angular.forEach(angular.element(".nosubmenu-margin-top"), function(value, key){
                         var a = angular.element(value);
                         a.removeClass('nosubmenu-margin-top');
                         a.addClass('submenu-margin-top');
                    });
    			} else {
	    			
	    			angular.element('#adminsubmenucontainer').removeClass("collapse1");
	    			angular.element('#adminsubmenucontainer').addClass("expanded1");
                    angular.forEach(angular.element(".nosubmenu-margin-top"), function(value, key){
                         var a = angular.element(value);
                         a.removeClass('nosubmenu-margin-top');
                         a.addClass('submenu-margin-top');
                    });
	    			$scope.currentSectionId = id;
	    		}
    		}
    	} else {
    		if($scope.currentSectionId == id){
    			angular.element('#adminsubmenucontainer').removeClass("expanded1");
    			angular.element('#adminsubmenucontainer').addClass("collapse1");
                angular.forEach(angular.element(".submenu-margin-top"), function(value, key){
                         var a = angular.element(value);
                         a.removeClass('submenu-margin-top');
                         a.addClass('nosubmenu-margin-top');
                    });
    			$scope.currentSectionId = 0;
    		} else {
    			var hasSubs = false;
    			for(var i = 0; i < subMenuItems.length; i++){
    				if(id==subMenuItems[i].parent_id){
    					hasSubs = true;
    					break;
    				}
    			}
    			if(!hasSubs){
    				angular.element('#adminsubmenucontainer').removeClass("expanded1");
    				angular.element('#adminsubmenucontainer').addClass("collapse1");
                    angular.forEach(angular.element(".submenu-margin-top"), function(value, key){
                         var a = angular.element(value);
                         a.removeClass('submenu-margin-top');
                         a.addClass('nosubmenu-margin-top');
                    });
    			}
    			
    			$scope.currentSectionId = id;
    		}
    		
    	}
    	
    }
}]);

adminControllers.controller('AdminTopUserInfoCtrl', ['$scope', 'sessionService', '$state', function ($scope, sessionService, $state) {
	
	$scope.userSession = sessionService.userSession;
	$scope.mySessionService = sessionService;
	console.log("adminmenustrl")
	console.log(sessionService.userSession)

	$scope.$watch("mySessionService.getUser()", function(newValue, oldValue) {
		$scope.userSession = newValue;
    },true);
}]);