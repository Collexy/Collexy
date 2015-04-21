angular.module("myApp").service('User', ["$resource", User]);
angular.module("myApp").service('UserGroup', ["$resource", UserGroup]);

function User($resource) {
    return $resource('/api/user/:id', {}, {
        query: {
            method: 'GET',
            params: {
                id: ''
            },
            isArray: true
        },
        update: {
            method: 'PUT',
            isArray: false
        },
        create: {
            method: 'POST',
            isArray: false
        }
        // delete: { method: 'delete', isArray: false }
    });
}

function UserGroup($resource) {
    return $resource('/api/user-group/:id', {}, {
        query: {
            method: 'GET',
            params: {
                nodeId: ''
            },
            isArray: true
        },
        update: {
            method: 'PUT',
            isArray: false
        },
        create: {
            method: 'POST',
            isArray: false
        }
        // delete: { method: 'delete', isArray: false }
    });
}
// var userServices = angular.module('userServices', ['ngResource']);
// userServices.factory('User', ['$resource',
//     function($resource) {
//         return $resource('/api/user/:id', {}, {
//             query: { method: 'GET', params: { id: '' }, isArray: true },
//             update: { method: 'PUT', isArray: false },
//             create: { method: 'POST', isArray: false }
//             // delete: { method: 'delete', isArray: false }
//         });
//     }]);
// userServices.factory('UserGroup', ['$resource',
//     function($resource) {
//         return $resource('/api/user-group/:id', {}, {
//             query: { method: 'GET', params: { nodeId: '' }, isArray: true },
//             update: { method: 'PUT', isArray: false },
//             create: { method: 'POST', isArray: false }
//             // delete: { method: 'delete', isArray: false }
//         });
//     }]);