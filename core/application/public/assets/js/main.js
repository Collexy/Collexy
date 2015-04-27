// function addslashes( str ) {
//     return (str + '').replace(/[\\"']/g, '\\$&').replace(/\u0000/g, '\\0');
// }

function escapeRegExp(string) {
    return string.replace(/([.*+?^=!:${}()|\[\]\/\\])/g, "\\$1");
}

function replaceAll(string, find, replace) {
    return string.replace(new RegExp(escapeRegExp(find), 'g'), replace);
}

Array.prototype.unique = function() {
    var a = this.concat();
    for(var i=0; i<a.length; ++i) {
        for(var j=i+1; j<a.length; ++j) {
            if(a[i] === a[j])
                a.splice(j--, 1);
        }
    }

    return a;
};

Array.prototype.contains = function(obj) {
    var i = this.length;
    while (i--) {
        if (this[i] === obj) {
            return true;
        }
    }
    return false;
};

Array.prototype.containsId = function(id) {
    var i = this.length;
    while (i--) {
        if (this[i].id === id) {
            return true;
        }
    }
    return false;
};

// Array.prototype.containsTest = function(target, compareOnProp) {
//     var i = this.length;
//     while (i--) {
//         if (this[i][compareOnProp] === target) {
//             return true;
//         }
//     }
//     return false;
// };

// self executing function here
$(document).ready(function(){
	// $("*").on("mouseover", function (e){ 
	// 	alert($(e.currentTarget).attr("id"))
	// 	$(".scroller").perfectScrollbar("update"); 
	// });
	// setTimeout(function() {

	//   $('.scroller').perfectScrollbar('update');
	// }, 10000);
	// var toggleSessionDataState = localStorage.getItem('toggleSessionDataState');
	// if(toggleSessionDataState== null){
	// 	toggleSessionDataState = "collapse";
	// }


	// if(toggleSessionDataState== "collapse"){
	// 	document.getElementById('stuff').style.height="36px";
	// } else{
	// 	document.getElementById('stuff').style.height="auto";
	// 	//document.getElementById('stuff').setAttribute("style","height:auto;");
	// }
	
	// var toggleSessionDataAhref = document.getElementById('toggle-session-data')
	// toggleSessionDataAhref.onclick = toggleSessionDataStateClickHandler;

	// function toggleSessionDataStateClickHandler(){
	// 	if(toggleSessionDataState== "expanded"){
	// 		document.getElementById('stuff').style.height="36px";
	// 		toggleSessionDataState = "collapse"
	// 		localStorage.setItem('toggleSessionDataState',toggleSessionDataState);
	// 	} else{
	// 		document.getElementById('stuff').style.height="auto";
	// 		//document.getElementById('stuff').setAttribute("style","height:auto;");
	// 		toggleSessionDataState = "expanded"
	// 		localStorage.setItem('toggleSessionDataState',toggleSessionDataState);
	// 	}
	// }

 //   var nodes = [
	// 	{
	// 		"id": 18,
	// 		"path": "1.18",
	// 		"created_by": 1,
	// 		"name": "gopher.jpg",
	// 		"node_type": 2,
	// 		"created_date": "2014-10-28T15:50:47.303Z",
	// 		"parent_id": 1,
	// 		"children": []
	// 	},
	// 	{
	// 		"id": 19,
	// 		"path": "1.19",
	// 		"created_by": 1,
	// 		"name": "postgresql.png",
	// 		"node_type": 2,
	// 		"created_date": "2014-10-28T17:53:37.488Z",
	// 		"parent_id": 1,
	// 		"children": []
	// 	},
	// 	{
	// 		"id": 23,
	// 		"path": "1.23",
	// 		"created_by": 1,
	// 		"name": "Sample picture folder",
	// 		"node_type": 2,
	// 		"created_date": "2014-11-17T16:57:14.654Z",
	// 		"parent_id": 1,
	// 		"children": []
	// 	},
	// 	{
	// 		"id": 24,
	// 		"path": "1.23.24",
	// 		"created_by": 1,
	// 		"name": "Goku_SSJ3.jpg",
	// 		"node_type": 2,
	// 		"created_date": "2014-11-17T16:58:57.285Z",
	// 		"parent_id": 23,
	// 		"children": []
	// 	}
	// ];

	// function treeify(nodes) {
	//     var indexed_nodes = {}, tree_roots = [];
	//     for (var i = 0; i < nodes.length; i += 1) {
	//         indexed_nodes[nodes[i].id] = nodes[i];
	//     }
	//     for (var i = 0; i < nodes.length; i += 1) {
	//         var parent_id = nodes[i].parent_id;
	//         if (parent_id === 1) {
	//             tree_roots.push(nodes[i]);
	//         } else {
	//             indexed_nodes[parent_id].children.push(nodes[i]);
	//         }
	//     }
	//     return tree_roots;
	// }
	// alert("remember to remove this from main.js:: this is only temporary")
	// alert(JSON.stringify(treeify(nodes), undefined, "\t"));
});