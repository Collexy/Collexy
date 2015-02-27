var entityServices = angular.module('entityServices', ['ngResource']);

entityServices.factory('Content', ['$resource',
    function($resource) {
        return $resource('/api/content/:nodeId', {}, {
            query: { method: 'GET', params: { nodeId: 'nodeId' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {nodeId: 'new'}, isArray: false },
            delete: { method: 'DELETE'}
        });
    }]);

entityServices.factory('ContentType', ['$resource',
    function($resource) {
        return $resource('/api/content-type/:nodeId', {nodeId: '@nodeId'}, {
            getExtended: { method: 'GET', isArray: false },
            query: { method: 'GET', isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {nodeId: 'new'}, isArray: false }

            // delete: { method: 'delete', isArray: false }
        });
    }]);

entityServices.factory('DataType', ['$resource',
    function($resource) {
        return $resource('/api/data-type/:nodeId', {}, {
            query: { method: 'GET', params: {nodeId: ''}, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {nodeId: 'new'}, isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);

entityServices.factory('Template', ['$resource',
    function($resource) {
        return $resource('/api/template/:nodeId', {}, {
            query: { method: 'GET', params: {nodeId: ''}, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {nodeId: 'new'}, isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);


entityServices.factory('Directory', ['$resource',
    function($resource) {
        return $resource('/api/directory/:rootdir/:name', {}, {
            query: { method: 'GET', isArray: false },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {name: 'new'}, isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);

entityServices.factory('MemberType', ['$resource',
    function($resource) {
        return $resource('/api/member-type/:nodeId', {nodeId: '@nodeId'}, {
            getExtended: { method: 'GET', isArray: false },
            query: { method: 'GET', isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {nodeId: 'new'}, isArray: false }

            // delete: { method: 'delete', isArray: false }
        });
    }]);