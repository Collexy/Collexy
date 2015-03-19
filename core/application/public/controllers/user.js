var userControllers = angular.module('userControllers', []);

userControllers.controller('UserListCtrl', ['$scope', 'User', function ($scope, User) {
	$scope.users = User.query();
}]);

userControllers.controller('UserEditCtrl', ['$scope', '$stateParams', '$location', 'User', 'UserGroup', function($scope, $stateParams, $location, User, UserGroup) {
  //$scope.user_groups = UserGroup.query();
	//console.log($stateParams);
  if ($stateParams.userId) {

    $scope.user = User.get({ userId: $stateParams.userId}, function(user){

    });
    //User.get({ userId: $stateParams.userId} , function(phone) {
  } else {

    $scope.user = new User();
  }

  //   $scope.isSelected =
  // function isSelected(listOfItems, item) {
  //   //console.log(listOfItems);
  //   if(listOfItems != undefined){
  //     for(var i = 0; i< listOfItems.length; i++){
  //       if(listOfItems[i]._id == item)
  //         return true;
  //     }
    
  //   }
  //   return false;
  // };

    $scope.isSelected =
  function isSelected(userUserGroup, user_group) {
    //console.log(listOfItems);
    if(userUserGroup != undefined){
      
        if(userUserGroup == user_group){
          
          return true;
        }
      
    
    }
    return false;
  };

  $scope.submit = function() {
    console.log("submit")

    function success(response) {
      console.log("success", response)
      $location.path("/admin/users");
    }

    function failure(response) {
      console.log("failure", response)

      _.each(response.data, function(errors, key) {
        if (errors.length > 0) {
          _.each(errors, function(e) {
            $scope.form[key].$dirty = true;
            $scope.form[key].$setValidity(e, false);
          });
        }
      });
    }

    if ($stateParams.userId) {
    	User.update({userId: $stateParams.userId}, $scope.user, success, failure);
      //User.update($scope.user, success, failure);
    } else {
    	User.create($scope.user, success, failure);
      //User.create($scope.user, success, failure);
    }

  }
}]);

userControllers.controller('UserProfileCtrl', ['$scope', '$stateParams', '$location', '$window', 'User', function($scope, $stateParams, $location, $window, User) {
   // $scope.$watch('username', function () {
   //   console.log("kkkk" + $scope.username); 
   // });
  var user = undefined;
  if($window.sessionStorage.token){
    var encodedProfile = $window.sessionStorage.token.split('.')[1];
    user = JSON.parse(url_base64_decode(encodedProfile));
    $scope.user = user;
    console.log("kkkk " + angular.fromJson($scope.user).username); 
    var user = angular.fromJson($scope.user);
    this.username = user.username;
    this.password = user.password;
    this.email = user.email;
    // angular.element("#adminbar").show();
      }
  else
    user = undefined;
  
  
  
  
  //$scope.user = User.get({ userId: $stateParams.userId}, function(user){

  $scope.submit = function() {
    console.log("submit")

    function success(response) {
      console.log("success", response)
      $location.path("/admin/users");
    }

    function failure(response) {
      console.log("failure", response)

      _.each(response.data, function(errors, key) {
        if (errors.length > 0) {
          _.each(errors, function(e) {
            $scope.form[key].$dirty = true;
            $scope.form[key].$setValidity(e, false);
          });
        }
      });
    }

    if (user._id) {
      User.update({userId: user._id}, $scope.user, success, failure);
      //User.update($scope.user, success, failure);
    }

  }
}]);

//this is used to parse the profile
function url_base64_decode(str) {
  var output = str.replace('-', '+').replace('_', '/');
  switch (output.length % 4) {
    case 0:
      break;
    case 2:
      output += '==';
     break;
    case 3:
      output += '=';
      break;
    default:
      throw 'Illegal base64url string!';
  }
  return window.atob(output); //polifyll https://github.com/davidchambers/Base64.js
}

userControllers.controller('UserLoginCtrl', ['$scope', '$http','$window', '$location', '$state', 'authenticationService', 'sessionService','$cookies', '$timeout', function ($scope, $http, $window, $location, $state, authenticationService, sessionService, $cookies, $timeout) {
  // $scope.user = {username: 'admin', password: '44444'};
  console.log("lol1")
  $scope.userSession = sessionService.userSession;
    $scope.mySessionService = sessionService;
  $scope.$watch("mySessionService.getUser()", function(newValue, oldValue) {
        $scope.userSession = newValue;
        $state.go('content');
    },true);
  console.log("lol2")
  
  $scope.message = '';
  $scope.submit = function (user) {
    $scope.user = user;
    alert("User: " + $scope.user.username + "logged in successfully!");
    //alert("username: " + $scope.user.username + ", password: " + $scope.user.password);

    // var req = {
    //   method: 'POST',
    //   url: '/api/user/login',
    //   data: {$scope.user}
    // }

    $http
      .post('/api/public/user/login', $scope.user)
        .success(function (data, status, headers, config) {
          $timeout(function() { 
          /* do stuff with $cookies here. */ 
            var usr = authenticationService.get({sid:$cookies.sessionauth}, function(usr){
              if(usr){
                window.location.reload(true);
              }
            });
          }, 100);
          
          
          // if(data){
          //   alert(data)
          //   console.log(data)
          //   sessionService.setUser(data)
          //   $state.go('adminDashboard');
          // }
          // else{
          //   $scope.message = 'Error: Invalid user or password. For the right login credentials please see the cover of the project report!';
          //   alert($scope.message);
          // }
        })
        .error(function (data, status, headers, config) {
          // Erase the token if the user fails to log in
          delete $window.sessionStorage.token;

          // Handle login errors here
          $scope.message = 'Error: Invalid user or password';
          alert($scope.message);
        });
  };
}]);



function MainCtrl($scope) {
	$scope.users = [
	'user1',
	'user2',
	'user3'
	];
}

function UserCtrl($scope) {
	$scope.users = [
	'user1',
	'user2',
	'user3'
	];
}
