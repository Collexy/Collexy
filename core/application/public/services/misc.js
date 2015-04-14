var contextMenuServices = angular.module('contextMenuServices', ['ngResource']);

// contextMenuServices.factory('ContextMenuService', function() {
//     return {
//       element: null,
//       menuElement: null
//     };
//   })

contextMenuServices.factory('ContextMenu', ['$resource',
    function($resource) {
        //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
        return $resource('/api/public/contextmenutest/:nodeType', {nodeType: '@nodeType'}, {
            query: { method: 'GET', isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);


var underscoreServices = angular.module('underscoreServices', ['ngResource']); 

underscoreServices.factory('_', ['$window', function ($window) { return $window._; }]);




var angularRouteService = angular.module('angularRouteService', ['ngResource']);

angularRouteService.factory('AngularRoute', ['$resource',
    function($resource) {
        //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
        return $resource('/api/angular-route', {}, {
            //query: { method: 'GET', isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);


angularRouteService.factory('Route', ['$resource',
    function($resource) {
        //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
        return $resource('/api/route', {}, {
            //query: { method: 'GET', isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);

angularRouteService.factory('Section', ['$resource',
    function($resource) {
        //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
        return $resource('/api/section/:name', {name: '@name'}, {
            query: { method: 'GET', isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);