angular.module("myApp").controller("TemplateTreeCtrl", TemplateTreeCtrl);
angular.module("myApp").controller("TemplateEditCtrl", TemplateEditCtrl);
angular.module("myApp").controller("TemplateDeleteCtrl", TemplateDeleteCtrl);
/**
 * @ngdoc controller
 * @name TemplateTreeCtrl
 * @function
 * @description
 * The controller for the template tree
 */
function TemplateTreeCtrl($scope, $stateParams, Template) {
    $scope.ContextMenuServiceName = "TemplateContextMenu"
    $scope.EntityChildrenServiceName = "TemplateChildren"
    Template.query({
        'levels': '1'
    }, {}, function(tree) {
        $scope.tree = tree;
    });
}
/**
 * @ngdoc controller
 * @name TemplateEditCtrl
 * @function
 * @description
 * The controller for editing a temlate
 */
function TemplateEditCtrl($scope, $stateParams, Template) {
    $scope.editorOptions = {
        lineWrapping: true,
        lineNumbers: true,
        readOnly: 'nocursor',
        mode: 'htmlmixed',
    };
    $scope.currentTab = 'template';
    $scope.stateParams = $stateParams;
    if ($stateParams.id) {
        $scope.node = Template.get({
            id: $stateParams.id
        }, function(node) {});
        //User.get({ userId: $stateParams.userId} , function(phone) {
    } else if ($stateParams.parent_id) {
        $scope.node = {
            "parent_template_id": parseInt($stateParams.parent_id)
        };
    } else {
        $scope.node = {}
    }

    $scope.allTemplates = Template.query({}, {}, function(node) {});

    console.log($scope.node)
    
    $scope.readFile = function() {
        var file = $scope.path;
        for (var i = 0; i < $scope.node.tmpl.length; i++) {
            if ($scope.node.tmpl[i].Path == $scope.node.dt.Path) {
                $scope.node.dt.Html = $scope.node.tmpl[i].Html;
                $scope.node.dt.Name = $scope.node.tmpl[i].Name;
            }
        }
        //$scope.node.dt.Html = 
        // Create a new FileReader Object
        //my_parser('http://localhost:8080/public/views/settings/data-type/tmpl/text-input.html');
    };
    
    $scope.isSelected = function isSelected(listOfItems, item) {
        if (listOfItems != undefined) {
            //console.log(listOfItems);
            for (var i = 0; i < listOfItems.length; i++) {
                //console.log(item)
                if (listOfItems[i] == item) {
                    //alert(item)
                    return true;
                }
            }
        }
        return false;
        // if (resArr.indexOf(item.toString()) > -1) {
        //     return true;
        //   } else {
        //     return false;
        //   }
        //   console.log(listOfItems);
    }
    $scope.submit = function() {
        console.log("submit")

        function success(response) {
            console.log("success", response)
            //$location.path("/admin/users");
        }

        function failure(response) {
            console.log("failure", response);
            // _.each(response.data, function(errors, key) {
            //   if (errors.length > 0) {
            //     _.each(errors, function(e) {
            //       $scope.form[key].$dirty = true;
            //       $scope.form[key].$setValidity(e, false);
            //     });
            //   }
            // });
        }
        if ($stateParams.id) {
            console.log("update");
            Template.update({
                id: $stateParams.id
            }, $scope.node, success, failure);
            console.log($scope.node)
            //User.update($scope.user, success, failure);
        } else {
            console.log("create");
            console.log($scope.node);
            Template.create($scope.node, success, failure);
            //User.create($scope.user, success, failure);
        }
    }
}
/**
 * @ngdoc controller
 * @name TemplateDeleteCtrl
 * @function
 * @description
 * The controller for deleting a template
 */
function TemplateDeleteCtrl($scope, $stateParams, Template) {
    $scope.delete = function(item) {
        console.log(item)
        Template.delete({
            id: item.id
        }, function() {
            console.log("template file and record with id: " + item.id + " deleted")
        })
    };
}