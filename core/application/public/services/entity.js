var entityServices = angular.module('entityServices', ['ngResource']);

entityServices.factory('Content', ['$resource',
    function($resource) {
        return $resource('/api/content/:id', {}, {
            //query: { method: 'GET', params: { id: 'id' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {id: 'new'}, isArray: false },
            delete: { method: 'DELETE'}
        });
    }]);

entityServices.factory('ContentChildren', ['$resource',
    function($resource) {
        return $resource('/api/content/:id/children', {}, {
            //query: { method: 'GET', params: { id: 'id' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);

entityServices.factory('ContentContextMenu', ['$resource',
    function($resource) {
        return $resource('/api/content/:id/contextmenu', {}, {
            //query: { method: 'GET', params: { id: 'id' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);


entityServices.factory('ContentType', ['$resource',
    function($resource) {
        return $resource('/api/content-type/:id', {id: '@id'}, {
            getExtended: { method: 'GET', isArray: false },
            query: { method: 'GET', isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {id: 'new'}, isArray: false }

            // delete: { method: 'delete', isArray: false }
        });
    }]);

entityServices.factory('ContentTypeChildren', ['$resource',
    function($resource) {
        return $resource('/api/content-type/:id/children', {}, {
            //query: { method: 'GET', params: { id: 'id' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);

entityServices.factory('DataType', ['$resource',
    function($resource) {
        return $resource('/api/data-type/:id', {}, {
            query: { method: 'GET', params: {id: ''}, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {id: 'new'}, isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);

entityServices.factory('Template', ['$resource',
    function($resource) {
        return $resource('/api/template/:id', {}, {
            query: { method: 'GET', params: {id: ''}, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', params: {id: 'new'}, isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);

entityServices.factory('TemplateChildren', ['$resource',
    function($resource) {
        return $resource('/api/template/:id/children', {}, {
            //query: { method: 'GET', params: { id: 'id' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
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

// entityServices.factory('MemberType', ['$resource',
//     function($resource) {
//         return $resource('/api/member-type/:id', {id: '@id'}, {
//             getExtended: { method: 'GET', isArray: false },
//             query: { method: 'GET', isArray: true },
//             update: { method: 'PUT', isArray: false },
//             create: { method: 'POST', params: {id: 'new'}, isArray: false }

//             // delete: { method: 'delete', isArray: false }
//         });
//     }]);