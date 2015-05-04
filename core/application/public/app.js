'use strict';
var $stateProviderRef = null;
var $urlRouterProviderRef = null;
var $controllerProviderRef = null;
// Declare app level module which depends on components
angular.module('myApp', ['ui.router', 'ngCookies', 'ngResource', 'ui.utils', 'checklist-model', 'ngDialog', 'ui.codemirror', 'perfect_scrollbar']).config(function($stateProvider, $urlRouterProvider, $locationProvider, $httpProvider, $provide, $controllerProvider) {
    $controllerProviderRef = $controllerProvider
    $urlRouterProviderRef = $urlRouterProvider;
    $stateProvider.state('login', {
        url: '/admin/login',
        templateUrl: 'public/views/admin/admin-login.html'
    });
    $locationProvider.html5Mode(true);
    $stateProviderRef = $stateProvider;
    //$httpProvider.interceptors.push('authInterceptorService');
}).constant('_', window._).run(['$rootScope', '$state', '$stateParams', 'authenticationService', '$location', '$window', '$q', '$cookies', 'sessionService', 'Route', '$timeout',
    function($rootScope, $state, $stateParams, authenticationService, $location, $window, $q, $cookies, sessionService, Route, $timeout) {
        //console.log(JSON.parse(localStorage["lastStateParams"]));
        var routes = Route.query({}, function(routes) {
            console.log(routes)
            angular.forEach(routes, function(value, key) {
                var templateUrl = "";
                var abstract = false;
                if (value.is_abstract) {
                    //alert(value.is_abstract)
                    var state = {
                        "url": value.url,
                        //"parent" : parent,
                        "abstract": true,
                        //"views":{}
                        "templateUrl": value.template_url
                        //"reloadOnSearch": true
                    };
                    $stateProviderRef.state(value.state, state);
                } else {
                    var state = {
                        "url": value.url,
                        //"parent" : parent,
                        //"abstract": value.abstract,
                        //"views":{}
                        "templateUrl": value.template_url
                        //"reloadOnSearch": true
                    };
                    $stateProviderRef.state(value.state, state);
                }
            });
            $state.go(JSON.parse(localStorage["lastState"]).name, JSON.parse(localStorage["lastStateParams"]))
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
        $rootScope.$on("$viewContentLoading", function(event, viewConfig) {
            // HACK ISH with the timeout - DOM manipulation should be done in a directive!
            // $timeout(function() {
            //      },0);
        });
        $rootScope.$on("$stateChangeSuccess", function(event, toState, toParams, fromState, fromParams) {
            $rootScope.state = toState;
            localStorage.setItem("lastState", JSON.stringify(toState));
            localStorage.setItem("lastStateParams", JSON.stringify(toParams))
        });
        $rootScope.$on("$stateChangeStart", function(event, toState, toParams, fromState, fromParams) {
            //alert("stateChangeStart")
            // if (toState != null && toState.data.access != null && toState.data.access.requiredAuthentication) {
            if ($cookies.sessionauth != null) {
                // 		if("permissions" in toState.data){
                var user = authenticationService.get({
                    sid: $cookies.sessionauth
                }, function(user) {
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
                    if (!allowUserAccess) {
                        console.log("test")
                        $state.go('login', toParams, {
                            notify: false
                        }).then(function() {
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
                $state.go('login', toParams, {
                    notify: false
                }).then(function() {
                    $rootScope.$broadcast('$stateChangeSuccess', toState, toParams, fromState, fromParams);
                });
                event.preventDefault();
            }
            // }
        });
    }
]).filter('unsafe', function($sce) {
    return function(val) {
        return $sce.trustAsHtml(val);
    };
}).filter('capitalize', function() {
    return function(input, all) {
        return ( !! input) ? input.replace(/([^\W_]+[^\s-]*) */g, function(txt) {
            return txt.charAt(0).toUpperCase() + txt.substr(1).toLowerCase();
        }) : '';
    }
}).filter('pathToUrl', function() {
    return function(text) {
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
.directive('wrapInput', [

    function() {
        return {
            replace: true,
            transclude: true,
            //template: '<div>{{prop.data_type.Html}}</div>'
            template: '<div class="input-wrapper" ng-transclude></div>'
        };
    }
]).directive('compile', function($compile, $timeout) {
    return {
        restrict: 'A',
        link: function(scope, elem, attrs) {
            $timeout(function() {
                $compile(elem.contents())(scope);
            });
        }
    };
}).directive('ngContextMenu', function($parse, $compile, $document, $injector) {
    var offset = {
        left: 40,
        top: -20
    }
    return {
        scope: false,
        link: function(scope, element, attrs) {
            var parentControllerScope = "";
            if(element.hasClass('collexy-controller')){
                parentControllerScope = element.scope();
            } else {
                parentControllerScope = element.closest('.collexy-controller').scope()
            }

            parentControllerScope.$watch("contextMenu", function(newValue, oldValue) {
                parentControllerScope.contextMenu = newValue;
                //alert($scope.contextMenu)
            }, true);
            //console.log(scope.$parent)
            var fn = $parse(attrs.ngContextMenu);
            var $oLay = angular.element(document.getElementById('overlay'))
            var offset = {
                left: 0,
                top: -76
            }
            scope.showOptions = function(item, $event) {
                console.log("showoptions")
                var overlayDisplay;
                if ($oLay.css("display") == "block") {
                    parentControllerScope.currentItem = null;
                    overlayDisplay = 'none'
                } else {
                    parentControllerScope.currentItem = item;
                    overlayDisplay = 'block'
                }
                if (angular.element(document.getElementById('adminsubmenucontainer')).hasClass('expanded1')) {
                    offset = {
                        left: 0,
                        top: -121
                    }
                }
                var overLayCSS = {
                    left: $event.clientX + offset.left + 'px',
                    top: $event.clientY + offset.top + 'px',
                    display: overlayDisplay
                }
                $oLay.css(overLayCSS)
            }
            scope.getEntityInfo = function(currentItem) {
                console.log("getEntityInfo")
                //alert(parentControllerScope.ContextMenuServiceName)
                var serviceName = parentControllerScope.ContextMenuServiceName;
                //console.log(currentItem)
                if (currentItem == undefined) {
                    currentItem = {
                        id: 0
                    }
                    //scope.currentItem = currentItem;
                }
                // console.log(currentItem)
                var myService = $injector.get(serviceName);
                myService.query({
                    id: currentItem.id
                }, function() {}).$promise.then(function(data) {
                    //console.log(data)
                    //var parentScope = element.parent().parent().parent().parent().parent().parent().parent().scope();
                    //scope.$parent.$parent.$parent.$parent.$parent.contextMenu = data;
                    //var s = angular.element(document.getElementsByClassName('outer-list-container')[0]).scope()
                    parentControllerScope.contextMenu = data;
                    //console.log(scope.$parent.$parent.$parent.$parent.$parent)
                    //scope.currentItem = currentItem;
                });
            }
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
                //scope.showOptions(item,)
                //$oLay = angular.element(document.getElementById('overlay'))
                scope.$apply(function() {
                    if (scope.getEntityInfo != undefined) scope.getEntityInfo(scope.data);
                    event.preventDefault();
                    event.stopPropagation();
                    //$oLay.html('<p>showing options for: {{currentItem.label}}</p>').show();
                    fn(scope, {
                        $event: event
                    });
                    //element.append($compile(overlayTemplate)(scope))
                    // $oLay.html(template).show();
                    // $compile($oLay.contents())(scope);
                    //     if(scope.currentItem!= null)
                    //      if('nodes' in scope.currentItem)
                    // console.log(scope.currentItem.nodes)
                });
            });

            function handleClickEvent(event) {
                //scope.currentItem = null;
                $oLay.css({
                    display: 'none'
                })
            }
            $document.bind('click', handleClickEvent);
        }
    }
}).directive('ngClickDialog', function($parse, ngDialog) {
    return function($scope, element, attr) {
        var item = $parse(attr.ngClickDialog)($scope)
        element.bind('click', function(event) {
            ngDialog.open({
                template: item.url,
                scope: $scope
            });
        })
    }
}).directive('ngExpandCollapse', function($parse, $injector) {
    return function($scope, element, attr) {
        var parentControllerScope = element.closest('.collexy-controller').scope()
        var item = $parse(attr.ngExpandCollapse)($scope)
        element.bind('click', function(event) {
            //alert("lol")
            $scope.expand_collapse(item)
        })
        $scope.expand_collapse = function(data) {
            var serviceName = parentControllerScope.EntityChildrenServiceName;
            if (!data.show) {
                if (data.nodes == undefined) {
                    data.nodes = [];
                }
                if (data.nodes.length == 0) {
                    // REST API call to fetch the current node's immediate children
                    var myService = $injector.get(serviceName);
                    data.nodes = myService.query({
                        id: data.id
                    }, function(node) {
                        //console.log(node)
                    });
                }
            }
            data.show = !data.show;
        }
    }
}).directive('collexyToggleTab', function($parse) {
    return function($scope, element, attrs) {
        var fn = $parse(attrs.collexyToggleTab);
        element.bind('click', function(event) {
            fn($scope, {
                $event: event
            });
        })
        $scope.toggleTab = function(item, $event) {
            var parentControllerScope = element.closest('.collexy-controller').scope()
            parentControllerScope.currentTab = item;
        }
    }
})
// Would IsolateScope be better here? 
// Rather than relying on a controller function? 
// Would it make more like a self contained component?
.directive('fileInput', ['$parse',
    function($parse) {
        return {
            restrict: 'A',
            link: function(scope, elm, attrs) {
                if (typeof(scope.test) == undefined) {
                    scope.test = {
                        "files": []
                    }
                }
                if (typeof(scope.test.files) !== undefined) {
                    scope.test["files"] = []
                }
                elm.bind('change', function() {
                    $parse(attrs.fileInput).assign(scope, elm[0].files)
                    scope.$apply()
                })
            }
        }
    }
]).directive('showonhoverparent', function() {
    return {
        link: function(scope, element, attrs) {
            element.parent().parent().bind('mouseenter', function() {
                element.show();
            });
            element.parent().parent().bind('mouseleave', function() {
                element.hide();
            });
        }
    };
}).directive('ckEditor', function() {
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
}).directive('topNavMargin', function($timeout) {
    return {
        link: function(scope, element, attrs) {
            $timeout(function() {
                if (angular.element('#adminsubmenucontainer').hasClass("collapse1")) {
                    if (element.hasClass("submenu-margin-top")) {
                        element.removeClass('submenu-margin-top');
                        element.addClass('nosubmenu-margin-top');
                    }
                } else {
                    if (element.hasClass("nosubmenu-margin-top")) {
                        element.removeClass('nosubmenu-margin-top');
                        element.addClass('submenu-margin-top');
                    }
                }
            });
        }
    };
}).directive('perfectScroll', function() {
    return {
        restrict: 'A',
        template: '<div ng-transclude></div>',
        transclude: true,
        scope: {},
        link: function(scope, element) {
            element.perfectScrollbar();
            element.bind("mouseover", function(e) {
                e.stopPropagation();
                e.preventDefault();
                element.perfectScrollbar('update');
            });
        }
    }
}).directive('psMouseOver', function() {
    return {
        link: function(scope, element) {
            element.bind("mouseover", function(e) {
                e.stopPropagation();
                e.preventDefault();
                element.perfectScrollbar('update');
            });
        }
    }
}).directive('ngRightClick', function($parse) {
    return function(scope, element, attrs) {
        var fn = $parse(attrs.ngRightClick);
        element.bind('contextmenu', function(event) {
            scope.$apply(function() {
                event.preventDefault();
                fn(scope, {$event:event});
            });
        });
    };
})
.directive('script', function() {
    return {
      restrict: 'E',
      scope: false,
      link: function(scope, elem, attr) {
        if (attr.type === 'text/javascript-lazy') {
          var code = elem.text();
          var f = new Function(code);
          f();
        }
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