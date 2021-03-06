angular.module("myApp").service("authenticationService", authenticationService);

function authenticationService($resource) {
    return $resource('/api/auth/:sid', {}, {
        query: {
            method: 'GET',
            isArray: false
        }
    });
}
// var authInterceptorService = angular.module('authInterceptorService', ['ngResource', 'ngCookies']);
// authInterceptorService.factory('authInterceptorService', function ($cookies, $q, $window, $location, authenticationService) {
//     console.log("authInterceptorService");
//     console.log($cookies.sessionauth)
//     return {
//         request: function (config) {
//             config.headers = config.headers || {};
//             if ($window.sessionStorage.token) {
//                 config.headers.Authorization = 'Bearer ' + $window.sessionStorage.token;
//             }
//             return config;
//         },
//         requestError: function(rejection) {
//             return $q.reject(rejection);
//         },
//         /* Set Authentication.isAuthenticated to true if 200 received */
//         response: function (response) {
//             // if (response != null && response.status == 200 && $window.sessionStorage.token && !authenticationService.isAuthenticated) {
//             //     authenticationService.isAuthenticated = true;
//             // }
//             return response || $q.when(response);
//         },
//         /* Revoke client authentication if 401 is received */
//         responseError: function(rejection) {
//             if (rejection != null && rejection.status === 401 && ($window.sessionStorage.token || authenticationService.isAuthenticated)) {
//                 delete $window.sessionStorage.token;
//                 authenticationService.isAuthenticated = false;
//                 $location.path("/admin/login");
//             }
//             return $q.reject(rejection);
//         }
//  };
// });
// var authenticationService = angular.module('authenticationService', ['ngResource']);
// authenticationService.factory('authenticationService', ['$resource',
//     function($resource) {
//         return $resource('/api/auth/:sid', {}, {
//             query: { method: 'GET', isArray: false }
//         });
//     }]);