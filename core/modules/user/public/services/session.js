var sessionService = angular.module('sessionService', []);

sessionService.service('sessionService', function($rootScope) {
    var userSession

    this.setUser = function(u){
    	userSession = u
    }

    this.getUser = function(){
    	return userSession;
    }
});
