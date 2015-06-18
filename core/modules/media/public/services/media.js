angular.module("myApp").service('Media', Media);
angular.module("myApp").service('MediaChildren', MediaChildren);
angular.module("myApp").service('MediaParents', MediaParents);
angular.module("myApp").service('MediaContextMenu', MediaContextMenu);
angular.module("myApp").service('MediaType', MediaType);
angular.module("myApp").service('MediaTypeChildren', MediaTypeChildren);

angular.module("myApp").service('MediaContextMenu', MediaContextMenu);

function Media($resource) {
    return $resource('/api/media/:id', {}, {
        //query: { method: 'GET', params: { id: 'id' }, isArray: true },
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
        },
        delete: {
            method: 'DELETE',
            params: {
                path: '@path'
            }
        }
    });
}

function MediaChildren($resource) {
    return $resource('/api/media/:id/children', {}, {
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

function MediaParents($resource) {
    return $resource('/api/media/:id/parents', {}, {
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

function MediaContextMenu($resource) {
    return $resource('/api/media/:id/contextmenu', {}, {
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

function MediaType($resource) {
    return $resource('/api/media-type/:id', {
        id: '@id'
    }, {
        getExtended: {
            method: 'GET',
            isArray: false
        },
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
            params: {
                id: 'new'
            },
            isArray: false
        }
        // delete: { method: 'delete', isArray: false }
    });
}

function MediaTypeChildren($resource) {
    return $resource('/api/media-type/:id/children', {}, {
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

function MediaContextMenu($resource) {
    return $resource('/api/media/:id/contextmenu', {}, {
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