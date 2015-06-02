angular.module("myApp").service('Content', Content);
angular.module("myApp").service('ContentChildren', ContentChildren);
angular.module("myApp").service('ContentParents', ContentParents);
angular.module("myApp").service('ContentContextMenu', ContentContextMenu);
angular.module("myApp").service('ContentType', ContentType);
angular.module("myApp").service('ContentTypeChildren', ContentTypeChildren);

// angular.module("myApp").service('MediaContextMenu', MediaContextMenu);

function Content($resource) {
    return $resource('/api/content/:id', {}, {
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
            method: 'DELETE'
        }
    });
}

function ContentChildren($resource) {
    return $resource('/api/content/:id/children', {}, {
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

function ContentParents($resource) {
    return $resource('/api/content/:id/parents', {}, {
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

function ContentContextMenu($resource) {
    return $resource('/api/content/:id/contextmenu', {}, {
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

function ContentType($resource) {
    return $resource('/api/content-type/:id', {
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

function ContentTypeChildren($resource) {
    return $resource('/api/content-type/:id/children', {}, {
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

// function MediaContextMenu($resource) {
//     return $resource('/api/media/:id/contextmenu', {}, {
//         //query: { method: 'GET', params: { id: 'id' }, isArray: true },
//         update: {
//             method: 'PUT',
//             isArray: false
//         },
//         create: {
//             method: 'POST',
//             isArray: false
//         }
//         // delete: { method: 'delete', isArray: false }
//     });
// }