angular.module("myApp").service('User', ["$resource", User]);
angular.module("myApp").service('UserGroup', ["$resource", UserGroup]);
angular.module("myApp").service('UserContextMenu', ["$resource", UserContextMenu]);
angular.module("myApp").service('UserGroupContextMenu', ["$resource", UserGroupContextMenu]);
angular.module("myApp").service('Permission', ["$resource", Permission]);

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
            // params: {
            //     id: 'new'
            // },
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

function Permission($resource) {
    return $resource('/api/permission/:name', {}, {
        query: {
            method: 'GET',
            params: {
                name: ''
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

function UserContextMenu($resource) {
    return $resource('/api/user/:id/contextmenu', {}, {
        //query: { method: 'GET', params: { id: 'id' }, isArray: true },
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

function UserGroupContextMenu($resource) {
    return $resource('/api/user-group/:id/contextmenu', {}, {
        //query: { method: 'GET', params: { id: 'id' }, isArray: true },
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