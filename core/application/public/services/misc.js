// var underscoreServices = angular.module('underscoreServices', ['ngResource']); 
// underscoreServices.factory('_', ['$window', function ($window) { return $window._; }]);
angular.module("myApp").service('Route', ["$resource", Route]);
angular.module("myApp").service('Section', ["$resource", Section]);
angular.module("myApp").service('DataTypeEditor', ["$resource", DataTypeEditor]);
angular.module("myApp").service('FileService', ["$resource", FileService]);

function Route($resource) {
    //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
    return $resource('/api/route', {}, {
        //query: { method: 'GET', isArray: true },
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

function Section($resource) {
    //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
    return $resource('/api/section/:name', {
        name: '@name'
    }, {
        query: {
            method: 'GET',
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

function DataTypeEditor($resource) {
    //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
    return $resource('/api/data-type-editor/:alias', {
        name: '@alias'
    }, {
        // query: {
        //     method: 'GET',
        //     isArray: true
        // },
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


function FileService($resource) {
    //return $resource('/api/node/node-type', {nodeTypeId: "@nodeTypeId", levels: "@levels"}, {
    return $resource('/api/file', {}, {
        //query: { method: 'GET', isArray: true },
        update: {
            method: 'PUT',
            isArray: false
        },
        create: {
            method: 'POST',
            isArray: false
        },
        delete: { method: 'DELETE' }
    });
}