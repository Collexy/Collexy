/**
 * Checklist-model
 * AngularJS directive for list of checkboxes
 */

 /* DB STUFF
 -- user_permissions for id:2 welcome
-- "[{"id": 2, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]"

-- user group permissions for id=3, getting started
-- "[{"id": 1, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]"
  */

 /* USE LIKE THIS
    <td ng-repeat="ug in userGroupsWithPermissions">
        <strong>{{ug.name}}</strong>
        <ul class="stripped-list">
            <!-- <li ng-repeat="p in allPermissions"><label><input type="checkbox" checklist-model="data.user_group_permissions[ug.id].permissions" checklist-value="p.name"> {{p.name}}</label></li> -->

            <!-- TEMPORARY SOLUTION ONLY!!!!!!! until changes in user_group_permission in db to "id": { "permissions" : []} structure-->
            <li ng-repeat="p in allPermissions">
                <span ng-repeat="temp in data.user_group_permissions" ng-if="temp.id==ug.id">
                    <label><input type="checkbox" checklist-model="data.user_group_permissions[$index].permissions" checklist-value="p.name"> {{p.name}}</label>
                </span>
                
            </li>
        </ul>
    </td>
    <td ng-repeat="ug in userGroupsWithoutPermissions">
        <strong>{{ug.name}}</strong>
        <ul class="stripped-list">

            <!-- TEMPORARY SOLUTION ONLY!!!!!!! until changes in user_group_permission in db to "id": { "permissions" : []} structure-->
            <li ng-repeat="p in allPermissions">
                <span>
                    <label><input type="checkbox" checklist-model="data.user_group_permissions" checklist-value="{'id': ug.id, 'permissions': [p.name]}"> {{p.name}}</label>
                </span>
                
            </li>
        </ul>
    </td>
 */
angular.module('checklist-model', []).directive('checklistModel', ['$parse', '$compile',
    function($parse, $compile) {
        // contains
        function contains(arr, item) {
            if(typeof item === 'object'){
                //alert("lol")
                if (angular.isArray(arr)) {
                    for (var i = 0; i < arr.length; i++) {
                        // if (angular.equals(arr[i], item)) {
                        //     // return true;
                        // } else 
                        if(typeof arr[i].id !== 'undefined' && typeof item.id !== 'undefined'){
                            if (arr[i].id == item.id) {
                                //console.log(arr[i].id + " i: " + i)
                                for (var k in item) {
                                    if(Array.isArray(item[k])){
                                        if(typeof arr[i][k] !== 'undefined'){
                                            if(Array.isArray(arr[i][k])){
                                                //if()
                                                for(var j=0; j<arr[i][k].length; j++){
                                                    // console.log("arr[i][k][j]")
                                                    // console.log(arr[i][k][j])
                                                    // console.log("item[k][0]")
                                                    // console.log(item[k][0])
                                                    if(arr[i][k][j] == item[k][0]){
                                                        return true;
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                //return true;
                            }
                        }
                    }
                }
                return false;
            } else {
                if (angular.isArray(arr)) {
                    for (var i = 0; i < arr.length; i++) {
                        if (angular.equals(arr[i], item)) {
                            return true;
                        }
                    }
                }
                return false;
            }
            
            // if (angular.isArray(arr)) {
            //     for (var i = 0; i < arr.length; i++) {
            //         if (angular.equals(arr[i], item)) {
            //             return true;
            //         } else if(typeof arr[i].id !== 'undefined' && typeof item.id !== 'undefined'){
            //             if (arr[i].id == item.id) {
            //                 return true;
            //             }
            //         }
            //     }
            // }
            // return false;
        }
        // add
        function add(arr, item) {
            arr = angular.isArray(arr) ? arr : [];
            for (var i = 0; i < arr.length; i++) {
                if (angular.equals(arr[i], item)) {
                    return arr;
                }
            }
            arr.push(item);
            return arr;
        }
        // remove
        function remove(arr, item) {
            console.log(arr)
            console.log(item)
            if (angular.isArray(arr)) {
                for (var i = 0; i < arr.length; i++) {
                    if (angular.equals(arr[i], item)) {
                        arr.splice(i, 1);
                        break;
                    }
                }
            }
            return arr;
        }
        function removeObject(arr, item) {
            var found = false;
            var arrayIndex = -1;
            if (angular.isArray(arr)) {
                for (var i = 0; i < arr.length; i++) {
                    if (arr[i].id == item.id) {
                        //arr.splice(i, 1);
                        found = true;
                        arrayIndex = i;
                        break;
                    }
                }

                if(found){

                    for(var k in item){

                        if(Array.isArray(item[k])){

                            for(var i = 0; i < item[k].length; i++){
                                for(var j= 0; j < arr[arrayIndex][k].length; j++){

                                    if(arr[arrayIndex][k][j] == item[k][i]){
                                        // delete arr[arrayIndex][k][j];
                                        arr[arrayIndex][k].splice(j,1);
                                        return arr;
                                    }
                                }
                            }
                        } else {
                            // if(arr[arrayIndex][k] == item[k]){
                            //     delete arr[arrayIndex][k];
                            //     return arr;
                            // }
                        }
                    }
                }
            }
            return arr;
        }
        function mergeTest(arr, item){
            arr = angular.isArray(arr) ? arr : [];
            var itemFound = false;
            var arrayIndex = -1;
            for (var i = 0; i < arr.length; i++) {
                if(arr[i].id != item.id){
                    //
                } else {
                    itemFound = true;
                    arrayIndex = i;
                    break;
                }
            }
            if(itemFound){
                //alert("do deep nested merge")
                var itemCopy = {}
                for (var k in item) {

                    if(Array.isArray(item[k])){
                        //alert("lolda")
                        // if(arr[k] === Array){
                            itemCopy[k] = []
                            //console.log(item[k])
                            for(var j=0; j<item[k].length; j++){

                                itemCopy[k].push(item[k][j]);
                            }
                            for(var l=0; l<arr[arrayIndex][k].length; l++){
                                itemCopy[k].push(arr[arrayIndex][k][l]);
                            }
                            
                        // } else {
                        //     itemCopy[k] = item[k];
                        // }
                        
                    } else {
                        //alert("lol")
                        itemCopy[k] = item[k];
                    }
                }
                arr[arrayIndex] = itemCopy;
                return arr;

            } else {
                arr.push(item);
                return arr;
            }
            
        }
        // http://stackoverflow.com/a/19228302/1458162
        function postLinkFn(scope, elem, attrs) {
            // compile with `ng-model` pointing to `checked`
            $compile(elem)(scope);
            // getter / setter for original model
            var getter = $parse(attrs.checklistModel);
            var setter = getter.assign;
            // value added to list
            var value = $parse(attrs.checklistValue)(scope.$parent);
            // watch UI checked change
            scope.$watch('checked', function(newValue, oldValue) {
                if (newValue === oldValue) {
                    return;
                }
                var current = getter(scope.$parent);
                if (newValue === true) {
                    if(typeof value === 'object'){
                        setter(scope.$parent, mergeTest(current, value));
                    } else {
                        setter(scope.$parent, add(current, value));
                    }
                    
                } else {
                    if(typeof value === 'object'){
                        setter(scope.$parent, removeObject(current, value));
                    } else {
                        setter(scope.$parent, remove(current, value));
                    }
                }
            });
            // watch original model change
            scope.$parent.$watch(attrs.checklistModel, function(newArr, oldArr) {
                scope.checked = contains(newArr, value);
            }, true);
        }
        return {
            restrict: 'A',
            priority: 1000,
            terminal: true,
            scope: true,
            compile: function(tElement, tAttrs) {
                if (tElement[0].tagName !== 'INPUT' || !tElement.attr('type', 'checkbox')) {
                    throw 'checklist-model should be applied to `input[type="checkbox"]`.';
                }
                if (!tAttrs.checklistValue) {
                    throw 'You should provide `checklist-value`.';
                }
                // exclude recursion
                tElement.removeAttr('checklist-model');
                // local scope var storing individual checkbox model
                tElement.attr('ng-model', 'checked');
                return postLinkFn;
            }
        };
    }
]);