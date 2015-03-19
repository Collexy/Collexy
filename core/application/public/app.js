'use strict';


var $stateProviderRef = null;
var $urlRouterProviderRef = null;

// Declare app level module which depends on components
angular.module('myApp', [
  'ui.router',
  'adminControllers',
  'authInterceptorService',
  'authenticationService',
  'sessionService',
  'userControllers',
  'memberControllers',
  'userServices',
  'memberServices',
  //'contentControllers',
  'mediaControllers',
  'contentTypeControllers',
  'mediaTypeControllers',
  'directoryControllers',
  //'stylesheetControllers',
  'dataTypeControllers',
  'templateControllers',
  'nodeServices',
  'entityServices',
  'contextMenuServices',
  'ui.utils',
  'checklist-model',
  'memberTypeControllers',
  'angularRouteService',
  'ngDialog',
  'ui.codemirror'
  // 'ng-context-menu'
  // 'underscoreServices'
])
.config(function($stateProvider,$urlRouterProvider,$locationProvider, $httpProvider, $provide) {
	

	$urlRouterProviderRef = $urlRouterProvider;
	$stateProvider
			.state('login', {
			url: '/admin/login',
			templateUrl: 'public/views/admin/admin-login.html'
		});
	// $provide.decorator('$state', function($delegate, $stateParams) {
 //        $delegate.forceReload = function() {
 //            return $delegate.go($delegate.current, $stateParams, {
 //                reload: true,
 //                inherit: false,
 //                notify: true
 //            });
 //        };
 //        return $delegate;
 //    });
	// $stateProvider
	// 	.state('adminLogin', {
	// 		url: '/admin/login',
	// 		templateUrl: 'public/views/admin/admin-login.html',
	// 		data: {
	// 			access: { requiredAuthentication: false }
	// 		} 
	// 	})
	// 	.state('adminLogout', {
	// 		url: '/admin/logout',
	// 		//templateUrl: 'public/views/admin/admin-logout.html',
	// 		data: {
	// 			access: { requiredAuthentication: false }
	// 		} 
	// 	})
	// 	// .state('adminIndex', {
	// 	// 	url: '/admin',
	// 	// 	abstract: true,
	// 	// 	templateUrl: 'public/views/admin/index.html',
	// 	// 	data: {
	// 	// 		access: { requiredAuthentication: true }
	// 	// 	}     
	// 	// })
	// 	.state('adminDashboard', {
	// 		url: '/admin',
	// 		templateUrl: 'public/views/admin/dashboard.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['use_administration_pages']
	// 		}     
	// 	})
	// 	.state('adminContent', {
	// 		url: '/admin/content',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/content/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['use_administration_pages']
	// 		}     
	// 	})
	// 	.state('adminContent.edit', {
	// 		url: '/edit/:nodeId',
	// 		templateUrl: 'public/views/content/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminContent.new', {
	// 		url: '/new?node_type&content_type_node_id&parent_id',
	// 		templateUrl: 'public/views/content/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminMedia', {
	// 		url: '/admin/media',
	// 		templateUrl: 'public/views/media/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['use_administration_pages']
	// 		}     
	// 	})
	// 	.state('adminMedia.edit', {
	// 		url: '/edit/:nodeId',
	// 		templateUrl: 'public/views/media/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminMedia.new', {
	// 		url: '/new?node_type&content_type_node_id&parent_id',
	// 		templateUrl: 'public/views/media/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminUsers', {
	// 		url: '/admin/users',
	// 		templateUrl: 'public/views/users/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['use_administration_pages', 'administer_users']
	// 		}     
	// 	})
	// 	.state('adminMembers', {
	// 		url: '/admin/members',
	// 		templateUrl: 'public/views/members/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['use_administration_pages', 'administer_users']
	// 		}     
	// 	})
	// 	.state('adminMembers.edit', {
	// 		url: '/edit/:id',
	// 		templateUrl: 'public/views/members/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminMembers.new', {
	// 		url: '/new?node_type&member_type_node_id&parent_id',
	// 		templateUrl: 'public/views/members/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminMembers.types', {
	// 		url: '/member-type',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/members/member-type/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['administer_content_types']
	// 		}     
	// 	})
	// 	.state('adminMembers.types.edit', {
	// 		url: '/edit/:nodeId',
	// 		templateUrl: 'public/views/members/member-type/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminMembers.types.new', {
	// 		url: '/new?node_type&member_type_node_id&parent_id',
	// 		templateUrl: 'public/views/members/member-type/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminMembers.roles', {
	// 		url: '/member-role',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/members/member-role/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['administer_content_types']
	// 		}     
	// 	})
	// 	.state('adminMembers.permissions', {
	// 		url: '/member-permission',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/members/member-permission/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['administer_content_types']
	// 		}     
	// 	})
	// 	.state('adminSettings', {
	// 		url: '/admin/settings',
	// 		templateUrl: 'public/views/settings/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['use_administration_pages', 'administer_site_configuration']
	// 		}     
	// 	})
	// 	.state('adminSettings.contentType', {
	// 		url: '/content-type',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/settings/content-type/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['administer_content_types']
	// 		}     
	// 	})
	// 	.state('adminSettings.contentType.edit', {
	// 		url: '/edit/:nodeId',
	// 		templateUrl: 'public/views/settings/content-type/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.contentType.new', {
	// 		url: '/new?type&parent',
	// 		templateUrl: 'public/views/settings/content-type/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.mediaType', {
	// 		url: '/media-type',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/settings/media-type/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.mediaType.edit', {
	// 		url: '/edit/:nodeId',
	// 		templateUrl: 'public/views/settings/media-type/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.mediaType.new', {
	// 		url: '/new?type&parent',
	// 		templateUrl: 'public/views/settings/media-type/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.dataType', {
	// 		url: '/data-type',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/settings/data-type/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			//permissions: ['administer_data_types']
	// 		}     
	// 	})
	// 	.state('adminSettings.dataType.edit', {
	// 		url: '/edit/:nodeId',
	// 		templateUrl: 'public/views/settings/data-type/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.dataType.new', {
	// 		url: '/new',
	// 		templateUrl: 'public/views/settings/data-type/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.template', {
	// 		url: '/template',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/settings/template/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.template.edit', {
	// 		url: '/edit/:nodeId',
	// 		templateUrl: 'public/views/settings/template/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.template.new', {
	// 		url: '/new?parent',
	// 		templateUrl: 'public/views/settings/template/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.script', {
	// 		url: '/script',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/settings/script/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			rootdir: 'script' 
	// 		}     
	// 	})
	// 	.state('adminSettings.script.edit', {
	// 		url: '/edit/:name',
	// 		templateUrl: 'public/views/settings/script/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.script.new', {
	// 		url: '/new?type&parent',
	// 		templateUrl: 'public/views/settings/script/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.stylesheet', {
	// 		url: '/stylesheet',
	// 		// abstract: true,
	// 		templateUrl: 'public/views/settings/stylesheet/index.html',
	// 		data: {
	// 			access: { requiredAuthentication: true },
	// 			rootdir: 'stylesheet'
	// 		}     
	// 	})
	// 	.state('adminSettings.stylesheet.edit', {
	// 		url: '/edit/:name',
	// 		templateUrl: 'public/views/settings/stylesheet/edit.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	})
	// 	.state('adminSettings.stylesheet.new', {
	// 		url: '/new?type&parent',
	// 		templateUrl: 'public/views/settings/stylesheet/new.html',
	// 		data: {
	// 			access: { requiredAuthentication: true }
	// 		}     
	// 	});
	$locationProvider.html5Mode(true);
	$stateProviderRef = $stateProvider;
	//$httpProvider.interceptors.push('authInterceptorService');
})

.constant('_', window._)

.run(['$rootScope', '$state', '$stateParams', 'authenticationService', '$location', '$window', '$q', '$cookies', 'sessionService', 'AngularRoute', 'Route', '$timeout', function ($rootScope, $state,$stateParams, authenticationService, $location, $window, $q, $cookies, sessionService, AngularRoute, Route, $timeout) {
	//console.log(JSON.parse(localStorage["lastStateParams"]));
	var routes = Route.query({}, function(routes){
		//console.log(routes)
		angular.forEach(routes, function (value, key){
			var templateUrl = "";
			var abstract = false;

			if('components' in value){
				if(value.components.length > 0){
					templateUrl = value.components[0].single
				}
			}
			
			if(value.is_abstract){
				//alert(value.is_abstract)
				var state = {
					"url": value.url,
					//"parent" : parent,
					"abstract": true,
					//"views":{}
					"templateUrl": templateUrl
				};
				$stateProviderRef.state(value.path, state);
			} else {
				var state = {
					"url": value.url,
					//"parent" : parent,
					//"abstract": value.abstract,
					//"views":{}
					"templateUrl": templateUrl
				};
				$stateProviderRef.state(value.path, state);
			}

			// var lol = window.location.href.split("/admin");
			// //alert("/admin" +lol[1])
			// if(value.url == "/admin" + lol[1]){

			// 	$state.go(value.path)
			// }
			//$rootScope.$broadcast('$stateChangeStart');
		});
		$state.go(JSON.parse(localStorage["lastState"]).name,JSON.parse(localStorage["lastStateParams"]))

		// console.log($rootScope.$state.current.name)
		// $rootScope.$state.forceReload();
		//$state.transitionTo($state.current, $stateParams, { reload: true, inherit: false, notify: true });
	});

			// if($cookies.sessionauth != null){

			// 		var user = authenticationService.get({sid:$cookies.sessionauth}, function(user){

			// 			var allowUserAccess = true;

						
			// 			if(!allowUserAccess){
			// 				console.log("test")
			// 				$state.go('login', toParams, {notify: false}).then(function() {	
			// 			    	$rootScope.$broadcast('$stateChangeSuccess', toState, toParams, fromState, fromParams);
			// 				});
			// 				event.preventDefault();
			// 			} else {
			// 				sessionService.setUser(user)
			// 			}
			// 		})

			// } else {
			// 	$state.go('login', toParams, {notify: false}).then(function() {	
			// 	    	$rootScope.$broadcast('$stateChangeSuccess', toState, toParams, fromState, fromParams);
			// 		});
			// 	event.preventDefault();
			// }

	// var angularRoutes = AngularRoute.query({}, function(routes){
	// 	console.log(routes)
	// 	angular.forEach(routes, function (value, key){ 
	// 		//console.log(value.components[0].single)
	// 		var templateUrl = "";
	// 		var parent = "";

	// 		if('parent' in value){
	// 			parent = value.parent;
	// 		}

	// 		if('components' in value){
	// 			if(value.components.length > 0){
	// 				templateUrl = value.components[0].single
	// 			}
	// 		}
	// 		if(templateUrl != ""){
	// 			var state = {
	// 				"url": value.url,
	// 				//"parent" : parent,
	// 				//"abstract": value.abstract,
	// 				//"views":{}
	// 				"templateUrl": templateUrl
	// 			};
	// 		} else {
	// 			var state = {
	// 				"url": value.url,
	// 				//"parent" : parent,
	// 				//"abstract": value.abstract,
	// 				//"views":{}
	// 				//"templateUrl": templateUrl
	// 			};
	// 		}
          

	// 		// here we configure the views
	// 		// angular.forEach(value.components, function (value1,key1) 
	// 		// {
	// 		// 	console.log(value1)
	// 		//   state.views[key1] = {
	// 		//     templateUrl : value1.single,
	// 		//   };
	// 		// });
	// 		if('parent' in value){
	// 			$stateProviderRef.state(value.parent + "." + value.alias, state);
	// 		} else {
	// 			$stateProviderRef.state(value.alias, state);
	// 		}
          
 //      });
	// })




	// angular.forEach(angularRoutes, function (value, key) 
 //      { 
 //          var state = {
 //            "url": value.url,
 //            //"parent" : value.parent,
 //            //"abstract": value.abstract,
 //            "views":{}
 //            //"templateUrl": value[0].single
 //          };

 //          // here we configure the views
 //          angular.forEach(value.components, function (value1,key1) 
 //          {
 //          	alert(key1 + ": : " + value1)
 //            state.views[key1] = {
 //              templateUrl : value1,
 //            };
 //          });

 //          $stateProviderRef.state(value.name, state);
 //      });
	$rootScope.$on("$viewContentLoading",function(event, viewConfig){
		// HACK ISH with the timeout - DOM manipulation should be done in a directive!
		// $timeout(function() {
			
	        
	 //      },0);
		
	});

	$rootScope.$on("$stateChangeSuccess",function (event, toState, toParams, fromState, fromParams) {
		$rootScope.state = toState;
		localStorage.setItem("lastState", JSON.stringify(toState));
		localStorage.setItem("lastStateParams", JSON.stringify(toParams))
		
	});
	$rootScope.$on("$stateChangeStart",function (event, toState, toParams, fromState, fromParams) {
		//alert("stateChangeStart")
		// if (toState != null && toState.data.access != null && toState.data.access.requiredAuthentication) {
			if($cookies.sessionauth != null){

				
		// 		if("permissions" in toState.data){
					var user = authenticationService.get({sid:$cookies.sessionauth}, function(user){
		// 				console.log(user)
		// 				var userPermissions = [];
		// 				var i = 0;
		// 				//var hasPermission = false;
						var allowUserAccess = true;
		// 				if(user != null && "id" in user){
		// 					console.log("user is defined")
							
		// 					if(typeof user.user_groups != 'undefined'){
		// 						console.log("user_groups in user [x]")
		// 						while(i < user.user_groups.length){
		// 							if("permissions" in user.user_groups[i]){
		// 								console.log("permissions in user_group [x]")
		// 								for(var k = 0; k< user.user_groups[i].permissions.length; k++){
		// 									userPermissions.push(user.user_groups[i].permissions[k].name)
		// 								}
		// 								i++;
		// 							}
		// 						}
		// 					}
		// 				}

		// 				var userPermissionsUnique = userPermissions.unique();
		// 				console.log(userPermissionsUnique)
		// 				for(var l=0; l < toState.data.permissions.length; l++){
		// 					//console.log(toState.data.permissions[l])
		// 					if(userPermissionsUnique.indexOf(toState.data.permissions[l]) == -1){
		// 						//console.log("FALSE")
		// 						allowUserAccess = false;
		// 					}
		// 				}
		// 				//console.log(allowUserAccess)
						
						if(!allowUserAccess){
							console.log("test")
							$state.go('login', toParams, {notify: false}).then(function() {	
						    	$rootScope.$broadcast('$stateChangeSuccess', toState, toParams, fromState, fromParams);
							});
							event.preventDefault();
						} else {
							sessionService.setUser(user)
						}
					})
		// 		} else {
		// 			$state.go('adminLogin', toParams, {notify: false}).then(function() {	
		// 		    	$rootScope.$broadcast('$stateChangeSuccess', toState, toParams, fromState, fromParams);
		// 			});
		// 			event.preventDefault();
		// 		}

			} else {
				$state.go('login', toParams, {notify: false}).then(function() {	
				    	$rootScope.$broadcast('$stateChangeSuccess', toState, toParams, fromState, fromParams);
					});
				event.preventDefault();
			}
		// }
    });
}])

.filter('unsafe', function($sce) {
    return function(val) {
        return $sce.trustAsHtml(val);
    };
})

.filter('capitalize', function() {
    return function(input, all) {
      return (!!input) ? input.replace(/([^\W_]+[^\s-]*) */g, function(txt){return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();}) : '';
    }
  })

.filter('pathToUrl', function(){
  return function(text){
  	text = text.replace(/\\/g, '/');
  	return text;
  }
})

// .filter('unique', function() {

//   return function (arr, field) {
//     var o = {}, i, l = arr.length, r = [];
//     for(i=0; i<l;i+=1) {
//       o[arr[i][field]] = arr[i];
//     }
//     for(i in o) {
//       r.push(o[i]);
//     }
//     return r;
//   };
// })

.directive('wrapInput', [function () {
   return {
      replace: true,
      transclude: true,
      //template: '<div>{{prop.data_type.Html}}</div>'
      template: '<div class="input-wrapper" ng-transclude></div>'
   };
}])

.directive('compile',function($compile, $timeout){
    return{
        restrict:'A',
        link: function(scope,elem,attrs){
            $timeout(function(){                
                $compile(elem.contents())(scope);    
            });
        }        
    };
})

.directive('ngContextMenu', function($parse, $compile, $document) {
	var offset = {
        left: 40,
        top: -20
    }
    return function(scope, element, attrs) {


    	//console.log(scope.menuOptions);
    	var template = "<ul>\
    				<li ng-repeat='option in currentItem.nodes'>\
    					<a>{{ option.label }}\
    						<ul ng-if='option.nodes' style='padding: 1em 0; list-style-type: none;'>\
    							<li ng-repeat='child in option.nodes'>\
									<a>{{ child.label }}</a>\
								</li>\
							</ul>\
						</a>\
					</li>\
				</ul>";
    	//var lol = '<ul><li ng-repeat="option in currentItem.nodes"><a>{{ option.label }} <ul ng-if="option.nodes" style="padding: 1em 0; list-style-type: none;"><li ng-repeat="child in option.nodes"><a>{{ child.label }}</a></li></ul></a></li></ul>';
    	var $oLay = angular.element(document.getElementById('overlay'))
    	
        var fn = $parse(attrs.ngContextMenu);

        // scope.showOptions = function (item,$event) {       
        //     var overlayDisplay;
        //     if (scope.currentItem === item) {
        //         scope.currentItem = null;
        //          overlayDisplay='none'
        //     }else{
        //          scope.currentItem = item;
        //         overlayDisplay='block'
        //     }
          
        //     var overLayCSS = {
        //         left: $event.clientX + offset.left + 'px',
        //         top: $event.clientY + offset.top + 'px',
        //         display: overlayDisplay
        //     }

        //      $oLay.css(overLayCSS)
        // }

        element.bind('contextmenu', function(event) {
        	//alert(scope.currentItem);

        	$oLay = angular.element(document.getElementById('overlay'))
            scope.$apply(function() {
            	if(scope.getEntityInfo != undefined)
            		scope.getEntityInfo(scope.data);
                event.preventDefault();
                event.stopPropagation();
                //$oLay.html('<p>showing options for: {{currentItem.label}}</p>').show();
                
                fn(scope, {$event:event});
                // $oLay.html(template).show();
                // $compile($oLay.contents())(scope);
            //     if(scope.currentItem!= null)
	           //      if('nodes' in scope.currentItem)
	        			// console.log(scope.currentItem.nodes)
            });
        });
        function handleClickEvent(event){
        	//scope.currentItem = null;
        	$oLay.css({display:'none'})
        }
        $document.bind('click', handleClickEvent);
    };
})

// Would IsolateScope be better here? 
// Rather than relying on a controller function? 
// Would it make more like a self contained component?
.directive('fileInput', ['$parse', function($parse){
	return {
		restrict: 'A',
		link: function(scope, elm, attrs){
  			if(typeof(scope.test) == undefined){
		      scope.test = { "files": []}
		    }
		    if(typeof(scope.test.files) !== undefined){
		      scope.test["files"] =[]
		    }
			elm.bind('change', function(){

				$parse(attrs.fileInput)
				.assign(scope,elm[0].files)
				scope.$apply()
			})
		}
	}
}])

.directive('showonhoverparent',
   function() {
      return {
         link : function(scope, element, attrs) {
            element.parent().parent().bind('mouseenter', function() {
                element.show();
            });
            element.parent().parent().bind('mouseleave', function() {
                 element.hide();
            });
       }
   };
})

.directive('ckEditor', function() {
  return {
    require: '?ngModel',
    link: function(scope, elm, attr, ngModel) {
      var ck = CKEDITOR.replace(elm[0]);

      if (!ngModel) return;

      ck.on('instanceReady', function() {
        ck.setData(ngModel.$viewValue);
      });

      function updateModel() {
          scope.$apply(function() {
              ngModel.$setViewValue(ck.getData());
          });
      }

      ck.on('change', updateModel);
      ck.on('key', updateModel);
      ck.on('dataReady', updateModel);

      ngModel.$render = function(value) {
        ck.setData(ngModel.$viewValue);
      };
    }
  };
})

.directive('topNavMargin', function($timeout) {
  return {
    link: function(scope, element, attrs) {
    	$timeout(function(){
	        if(angular.element('#adminsubmenucontainer').hasClass("collapse1")){
				if(element.hasClass("submenu-margin-top")){
					element.removeClass('submenu-margin-top');
	                element.addClass('nosubmenu-margin-top');
				}
			} else {
				if(element.hasClass("nosubmenu-margin-top")){
					element.removeClass('nosubmenu-margin-top');
	                 element.addClass('submenu-margin-top');
				}
				
			}
	    });
      
    }
  };
});

// .directive('ngContextMenu', [
// 	'$parse',
//     '$document',
//     'ContextMenuService',
//     function($parse, $document, ContextMenuService) {

//       return {
//         restrict: 'A',
//         scope: {
//           'callback': '&contextMenu',
//           'disabled': '&contextMenuDisabled'
//         },
//         link: function($scope, $element, $attrs) {
//         	alert($scope.menuOptions);
// 	        var data = $parse($attrs.ngContextMenu)($scope);
// 	        console.log(data);

// 	        var opened = false;
//           	function open(event, menuElement) {
// 	            menuElement.addClass('open');

// 	            var doc = $document[0].documentElement;
// 	            var docLeft = (window.pageXOffset || doc.scrollLeft) -
// 	                          (doc.clientLeft || 0),
// 	                docTop = (window.pageYOffset || doc.scrollTop) -
// 	                         (doc.clientTop || 0),
// 	                elementWidth = menuElement[0].scrollWidth,
// 	                elementHeight = menuElement[0].scrollHeight;
// 	            var docWidth = doc.clientWidth + docLeft,
// 	              docHeight = doc.clientHeight + docTop,
// 	              totalWidth = elementWidth + event.pageX,
// 	              totalHeight = elementHeight + event.pageY,
// 	              left = Math.max(event.pageX - docLeft, 0),
// 	              top = Math.max(event.pageY - docTop, 0);

// 	            if (totalWidth > docWidth) {
// 	              left = left - (totalWidth - docWidth);
// 	            }

// 	            if (totalHeight > docHeight) {
// 	              top = top - (totalHeight - docHeight);
// 	            }

// 	            menuElement.css('top', top + 'px');
// 	            menuElement.css('left', left + 'px');
// 	            opened = true;
//           	}

// 			function close(menuElement) {
// 				menuElement.removeClass('open');
// 				opened = false;
// 			}


// 	        $element.bind('contextmenu', function(event) {
// 				if (ContextMenuService.menuElement !== null) {
// 					close(ContextMenuService.menuElement);
// 				}
// 				ContextMenuService.menuElement = angular.element(
// 					document.getElementById($attrs.target)
// 				);
// 				ContextMenuService.element = event.target;
// 				//console.log('set', ContextMenuService.element);

// 				event.preventDefault();
// 				event.stopPropagation();
// 				$scope.$apply(function() {
// 					$scope.callback({ $event: event });
// 				});
// 				$scope.$apply(function() {
// 					open(event, ContextMenuService.menuElement);
// 				});
	                
// 	        });

// 	        function handleClickEvent(event) {
// 	            if (opened &&
// 	              (event.button !== 2 ||
// 	               event.target !== ContextMenuService.element)) {
// 	              $scope.$apply(function() {
// 	                close(ContextMenuService.menuElement);
// 	              });
// 	            }
//           	}


//           	// Firefox treats a right-click as a click and a contextmenu event
//           	// while other browsers just treat it as a contextmenu event
// 			$document.bind('click', handleClickEvent);
// 			$document.bind('contextmenu', handleClickEvent);

// 			$scope.$on('$destroy', function() {
// 				//console.log('destroy');
// 				$document.unbind('click', handleClickEvent);
// 				$document.unbind('contextmenu', handleClickEvent);
// 			});
// 	    }
//     };
// }]);

// .directive('ngRightClick', function($parse) {
//     return function(scope, element, attrs) {
//         var fn = $parse(attrs.ngRightClick);
//         element.bind('contextmenu', function(event) {
//             scope.$apply(function() {
//                 event.preventDefault();
//                 fn(scope, {$event:event});
//             });
//         });
//     };
// });

// .directive('contextMenu', [
//     '$document',
//     'ContextMenuService',
//     function($document, ContextMenuService) {
//       return {
//         restrict: 'A',
//         scope: {
//           'callback': '&contextMenu',
//           'disabled': '&contextMenuDisabled'
//         },
//         link: function($scope, $element, $attrs) {
//           var opened = false;

//           function open(event, menuElement) {
//             menuElement.addClass('open');

//             var doc = $document[0].documentElement;
//             var docLeft = (window.pageXOffset || doc.scrollLeft) -
//                           (doc.clientLeft || 0),
//                 docTop = (window.pageYOffset || doc.scrollTop) -
//                          (doc.clientTop || 0),
//                 elementWidth = menuElement[0].scrollWidth,
//                 elementHeight = menuElement[0].scrollHeight;
//             var docWidth = doc.clientWidth + docLeft,
//               docHeight = doc.clientHeight + docTop,
//               totalWidth = elementWidth + event.pageX,
//               totalHeight = elementHeight + event.pageY,
//               left = Math.max(event.pageX - docLeft, 0),
//               top = Math.max(event.pageY - docTop, 0);

//             if (totalWidth > docWidth) {
//               left = left - (totalWidth - docWidth);
//             }

//             if (totalHeight > docHeight) {
//               top = top - (totalHeight - docHeight);
//             }

//             menuElement.css('top', top + 'px');
//             menuElement.css('left', left + 'px');
//             opened = true;
//           }

//           function close(menuElement) {
//             menuElement.removeClass('open');
//             opened = false;
//           }

//           $element.bind('contextmenu', function(event) {
//             if (!$scope.disabled()) {
//               if (ContextMenuService.menuElement !== null) {
//                 close(ContextMenuService.menuElement);
//               }
//               ContextMenuService.menuElement = angular.element(
//                 document.getElementById($attrs.target)
//               );
//               ContextMenuService.element = event.target;
//               //console.log('set', ContextMenuService.element);

//               event.preventDefault();
//               event.stopPropagation();
//               $scope.$apply(function() {
//                 $scope.callback({ $event: event });
//               });
//               $scope.$apply(function() {
//                 open(event, ContextMenuService.menuElement);
//               });
//             }
//           });

//           function handleKeyUpEvent(event) {
//             //console.log('keyup');
//             if (!$scope.disabled() && opened && event.keyCode === 27) {
//               $scope.$apply(function() {
//                 close(ContextMenuService.menuElement);
//               });
//             }
//           }

//           function handleClickEvent(event) {
//             if (!$scope.disabled() &&
//               opened &&
//               (event.button !== 2 ||
//                event.target !== ContextMenuService.element)) {
//               $scope.$apply(function() {
//                 close(ContextMenuService.menuElement);
//               });
//             }
//           }

//           $document.bind('keyup', handleKeyUpEvent);
//           // Firefox treats a right-click as a click and a contextmenu event
//           // while other browsers just treat it as a contextmenu event
//           $document.bind('click', handleClickEvent);
//           $document.bind('contextmenu', handleClickEvent);

//           $scope.$on('$destroy', function() {
//             //console.log('destroy');
//             $document.unbind('keyup', handleKeyUpEvent);
//             $document.unbind('click', handleClickEvent);
//             $document.unbind('contextmenu', handleClickEvent);
//           });
//         }
//       };
//     }
//   ]);
