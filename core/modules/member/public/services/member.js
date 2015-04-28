angular.module("myApp").service('Member', ["$resource", Member]);
angular.module("myApp").service('MemberGroup', ["$resource", MemberGroup]);
angular.module("myApp").service('MemberType', ["$resource", MemberType]);
angular.module("myApp").service('MemberContextMenu', ["$resource", MemberContextMenu]);


function Member($resource) {
    return $resource('/api/member/:id', {}, {
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

function MemberGroup($resource) {
    return $resource('/api/member-group/:id', {}, {
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

function MemberType($resource) {
    return $resource('/api/member-type/:id', {
        id: '@id'
    }, {
        getExtended: {
            method: 'GET',
            isArray: false
        },
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

function MemberContextMenu($resource) {
    return $resource('/api/member/:id/contextmenu', {}, {
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