angular.module("myApp").service('DataType', ["$resource", DataType]);
angular.module("myApp").service('Template', ["$resource", Template]);
angular.module("myApp").service('TemplateChildren', ["$resource", TemplateChildren]);
angular.module("myApp").service('Directory', ["$resource", Directory]);
angular.module("myApp").service('DirectoryDelete', ["$resource", DirectoryDelete]);
angular.module("myApp").service('ContentTypeContextMenu', ["$resource", ContentTypeContextMenu]);
angular.module("myApp").service('MediaTypeContextMenu', ["$resource", MediaTypeContextMenu]);
angular.module("myApp").service('DataTypeContextMenu', ["$resource", DataTypeContextMenu]);
angular.module("myApp").service('TemplateContextMenu', ["$resource", TemplateContextMenu]);
angular.module("myApp").service('DirectoryContextMenu', ["$resource", DirectoryContextMenu]);


function DataType($resource) {
    return $resource('/api/data-type/:id', {}, {
        // query: {
        //     method: 'GET',
        //     params: {
        //         id: ''
        //     },
        //     isArray: true
        // },
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

function DirectoryDelete($resource) {
    return $resource('/api/directory/:path', {}, {
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

function ContentTypeContextMenu($resource) {
    return $resource('/api/content-type/:id/contextmenu', {}, {
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

function MediaTypeContextMenu($resource) {
    return $resource('/api/media-type/:id/contextmenu', {}, {
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

function DataTypeContextMenu($resource) {
    return $resource('/api/data-type/:id/contextmenu', {}, {
        //query: { method: 'GET', params: { id: 'id' }, isArray: true },
        update: {
            method: 'PUT',
            isArray: false
        },
        create: {
            method: 'POST',
            isArray: false
        },
        //delete: { method: 'delete', isArray: false }
    });
}

function TemplateContextMenu($resource) {
    return $resource('/api/template/:id/contextmenu', {}, {
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

function DirectoryContextMenu($resource) {
    return $resource('/api/directory/:rootdir/:name/:is_dir/contextmenu', {}, {
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

