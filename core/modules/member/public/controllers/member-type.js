angular.module("myApp").controller("MemberTypeListCtrl", MemberTypeListCtrl);
angular.module("myApp").controller("MemberTypeEditCtrl", MemberTypeEditCtrl);

/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function MemberTypeListCtrl($scope, MemberType){
    $scope.tree = MemberType.query();
}

/**
 * @ngdoc controller
 * @name ContentTreeCtrl
 * @function
 * @description
 * The controller for deleting content
 */
function MemberTypeEditCtrl($scope, $stateParams, MemberType, DataType){
    $scope.currentTab = 'member-type';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        MemberType.getExtended({extended: true}, {id: $stateParams.id}, function(){}).$promise.then(function(entity){
            $scope.entity = entity;
        }, function(){
            console.log("Database error: Error fetching MemberType")
        });
    }

    DataType.query().$promise.then(function(allDataTypes){
        $scope.allDataTypes = allDataTypes;
    }, function(){
        console.log("Database error: Error fetching Data Types")
    });

    $scope.submit = function() {
        console.log("submit")

        function success(response) {
            console.log("success", response)
            //$location.path("/admin/users");
        }

        function failure(response) {
            console.log("failure", response);
        }

        if ($stateParams.id) {
            console.log("update");
            MemberType.update({id: $stateParams.id}, $scope.entity, success, failure);
            console.log($scope.entity)
        } else {
            console.log("create");
            MemberType.create($scope.entity, success, failure);
        }

    }
    $scope.aliasOrName = function(alias, name){
        if(alias != null && alias != ""){
            return alias;
        }
        return name;
    }

    $scope.toggleTab = function (item,$event) {
        $scope.currentTab = item;
    }

    $scope.addTab = function(){
        if('tabs' in $scope.entity){
            
        }
        else{
            $scope.node["tabs"] = [];
        }
        tab = {"name": "mytab", "properties" : []}
        $scope.entity.tabs.push(tab);
    }
    $scope.addProp = function(tab){
        if('tabs' in $scope.entity){
            var tabs = $scope.entity.tabs;
            if(tabs.length > 0){
                for(var i = 0; i < tabs.length; i++){
                    if(tabs[i].name == tab){
                        if('properties' in tabs[i]){

                        }
                        else{
                            tabs[i].properties = [];
                        }
                        tabs[i].properties.push({"name":"property name","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"});
                    }
                }
                $scope.entity.tabs = tabs;
            }
            
        }
    }
}


var memberTypeControllers = angular.module('memberTypeControllers', []);