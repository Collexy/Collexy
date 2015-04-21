angular.module("myApp").service('DataType', ["$resource", DataType]);
angular.module("myApp").service('Template', ["$resource", Template]);
angular.module("myApp").service('TemplateChildren', ["$resource", TemplateChildren]);
angular.module("myApp").service('Directory', ["$resource", Directory]);

function DataType($resource) {
    return $resource('/api/data-type/:id', {}, {
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
            params: {
                id: 'new'
            },
            isArray: false
        }
        // delete: { method: 'delete', isArray: false }
    });
}

function Template($resource) {
    return $resource('/api/template/:id', {}, {
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
            params: {
                id: 'new'
            },
            isArray: false
        }
        // delete: { method: 'delete', isArray: false }
    });
}

function TemplateChildren($resource) {
    return $resource('/api/template/:id/children', {}, {
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

function Directory($resource) {
    return $resource('/api/directory/:rootdir/:name', {}, {
        query: {
            method: 'GET',
            isArray: false
        },
        update: {
            method: 'PUT',
            isArray: false
        },
        create: {
            method: 'POST',
            params: {
                name: 'new'
            },
            isArray: false
        }
        // delete: { method: 'delete', isArray: false }
    });
}