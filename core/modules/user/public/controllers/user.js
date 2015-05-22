angular.module("myApp").controller("UserCtrl", UserCtrl);
angular.module("myApp").controller("UserListCtrl", UserListCtrl);
angular.module("myApp").controller("UserEditCtrl", UserEditCtrl);
angular.module("myApp").controller("UserProfileCtrl", UserProfileCtrl);
angular.module("myApp").controller("UserLoginCtrl", UserLoginCtrl);
angular.module("myApp").controller('UserDeleteCtrl', UserDeleteCtrl);
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function UserCtrl($scope, $state) {
    $scope.state = $state;
}
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function UserListCtrl($scope, $stateParams, User) {
    $scope.ContextMenuServiceName = "UserContextMenu"
    $scope.users = User.query();
}
/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function UserEditCtrl($scope, $stateParams, User, UserGroup, sessionService, $interpolate, ngDialog) {
    User.get({
        id: $stateParams.id
    }, function() {}).$promise.then(function(data) {
        $scope.data = data;
        $scope.currentTab = data.username;
        UserGroup.query().$promise.then(function(userGroups) {
            $scope.allUserGroups = userGroups;
            var availableUserGroups = [];
            var selectedUserGroups = [];
            for (var i = 0; i < data.user_group_ids.length; i++) {
                for (var j = 0; j < userGroups.length; j++) {
                    //console.log("[i] = " + data.user_group_ids[i] + ", [j]: " +userGroups[j].id)
                    if (data.user_group_ids[i] != userGroups[j].id) {
                        availableUserGroups.push(userGroups[j])
                    } else {
                        selectedUserGroups.push(userGroups[j])
                    }
                }
            }
            availableUserGroups.unique();
            selectedUserGroups.unique();
            $scope.availableUserGroups = availableUserGroups;
            $scope.selectedUserGroups = selectedUserGroups;
        }, function() {
            //ERR
        })
    }, function() {
        // ERR
    });
    $scope.moveItem = function(item, from, to) {
        alert("moveitem")
        var idx = from.indexOf(item);
        if (idx != -1) {
            from.splice(idx, 1);
            to.push(item);
        }
        var user_group_ids = [];
        for (var i = 0; i < $scope.selectedUserGroups.length; i++) {
            user_group_ids.push($scope.selectedUserGroups[i].id);
        }
        $scope.data.user_group_ids = user_group_ids;
    };
    $scope.moveAll = function(from, to) {
        angular.forEach(from, function(item) {
            to.push(item);
        });
        from.length = 0;
    };
}
var userControllers = angular.module('userControllers', []);
// userControllers.controller('UserListCtrl', ['$scope', 'User', function ($scope, User) {
//  $scope.users = User.query();
// }]);
// userControllers.controller('UserEditCtrl', ['$scope', '$stateParams', '$location', 'User', 'UserGroup', function($scope, $stateParams, $location, User, UserGroup) {
//   //$scope.user_groups = UserGroup.query();
//  //console.log($stateParams);
//   if ($stateParams.userId) {
//     $scope.user = User.get({ userId: $stateParams.userId}, function(user){
//     });
//     //User.get({ userId: $stateParams.userId} , function(phone) {
//   } else {
//     $scope.user = new User();
//   }
//   //   $scope.isSelected =
//   // function isSelected(listOfItems, item) {
//   //   //console.log(listOfItems);
//   //   if(listOfItems != undefined){
//   //     for(var i = 0; i< listOfItems.length; i++){
//   //       if(listOfItems[i]._id == item)
//   //         return true;
//   //     }
//   //   }
//   //   return false;
//   // };
//     $scope.isSelected =
//   function isSelected(userUserGroup, user_group) {
//     //console.log(listOfItems);
//     if(userUserGroup != undefined){
//         if(userUserGroup == user_group){
//           return true;
//         }
//     }
//     return false;
//   };
//   $scope.submit = function() {
//     console.log("submit")
//     function success(response) {
//       console.log("success", response)
//       $location.path("/admin/users");
//     }
//     function failure(response) {
//       console.log("failure", response)
//       _.each(response.data, function(errors, key) {
//         if (errors.length > 0) {
//           _.each(errors, function(e) {
//             $scope.form[key].$dirty = true;
//             $scope.form[key].$setValidity(e, false);
//           });
//         }
//       });
//     }
//     if ($stateParams.userId) {
//      User.update({userId: $stateParams.userId}, $scope.user, success, failure);
//       //User.update($scope.user, success, failure);
//     } else {
//      User.create($scope.user, success, failure);
//       //User.create($scope.user, success, failure);
//     }
//   }
// }]);
function UserProfileCtrl($scope, $stateParams, $location, $window, User) {
    // $scope.$watch('username', function () {
    //   console.log("kkkk" + $scope.username); 
    // });
    var user = undefined;
    if ($window.sessionStorage.token) {
        var encodedProfile = $window.sessionStorage.token.split('.')[1];
        user = JSON.parse(url_base64_decode(encodedProfile));
        $scope.user = user;
        console.log("kkkk " + angular.fromJson($scope.user).username);
        var user = angular.fromJson($scope.user);
        this.username = user.username;
        this.password = user.password;
        this.email = user.email;
        // angular.element("#adminbar").show();
    } else user = undefined;
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
            User.update({
                userId: user._id
            }, $scope.user, success, failure);
            //User.update($scope.user, success, failure);
        }
    }
}
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

function UserLoginCtrl($scope, $http, $window, $location, $state, authenticationService, sessionService, $cookies, $timeout) {
    // $scope.user = {username: 'admin', password: '44444'};
    console.log("lol1")
    $scope.userSession = sessionService.userSession;
    $scope.mySessionService = sessionService;
    $scope.$watch("mySessionService.getUser()", function(newValue, oldValue) {
        $scope.userSession = newValue;
        $state.go('content');
    }, true);
    console.log("lol2")
    $scope.message = '';
    $scope.submit = function(user) {
        $scope.user = user;
        alert("User: " + $scope.user.username + "logged in successfully!");
        //alert("username: " + $scope.user.username + ", password: " + $scope.user.password);
        // var req = {
        //   method: 'POST',
        //   url: '/api/user/login',
        //   data: {$scope.user}
        // }
        $http.post('/api/public/user/login', $scope.user).success(function(data, status, headers, config) {
            $timeout(function() {
                /* do stuff with $cookies here. */
                var usr = authenticationService.get({
                    sid: $cookies.get('sessionauth')
                }, function(usr) {
                    if (usr) {
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
        }).error(function(data, status, headers, config) {
            // Erase the token if the user fails to log in
            delete $window.sessionStorage.token;
            // Handle login errors here
            $scope.message = 'Error: Invalid user or password';
            alert($scope.message);
        });
    };
}

function MainCtrl($scope) {
    $scope.users = ['user1', 'user2', 'user3'];
}

function UserCtrl($scope) {
    $scope.users = ['user1', 'user2', 'user3'];
}

/**
 * @ngdoc controller
 * @name MemberDeleteCtrl
 * @function
 * @description
 * The controller for deleting a member
 */
function UserDeleteCtrl($scope, $stateParams, User) {
    $scope.delete = function(item) {
        console.log(item)
        User.delete({
            id: item.id
        }, function() {
            console.log("user record with id: " + item.id + " deleted")
        })
    };
}