var memberServices = angular.module('memberServices', ['ngResource']);

memberServices.factory('Member', ['$resource',
    function($resource) {
        return $resource('/api/member/:id', {}, {
            query: { method: 'GET', params: { id: '' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);