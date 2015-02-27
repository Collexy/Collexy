var nodeServices = angular.module('nodeServices', ['ngResource']);

nodeServices.factory('NodeChildren', ['$resource',
    function($resource) {
        return $resource('/api/node/:nodeId/children', {}, {
            query: { method: 'GET', params: { nodeId: 'nodeId' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);

nodeServices.factory('Node', ['$resource',
    function($resource) {
        //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
        return $resource('/api/node', {}, {
            //query: { method: 'GET', isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);