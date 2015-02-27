var userServices = angular.module('userServices', ['ngResource']);

userServices.factory('User', ['$resource',
    function($resource) {
        return $resource('/api/user/:userId', {}, {
            query: { method: 'GET', params: { userId: '' }, isArray: true },
            update: { method: 'PUT', isArray: false },
            create: { method: 'POST', isArray: false }
            // delete: { method: 'delete', isArray: false }
        });
    }]);