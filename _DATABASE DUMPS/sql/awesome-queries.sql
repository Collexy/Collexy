-- SELECT node.name, test.*
-- FROM node,
-- LATERAL (
-- 	SELECT user_group.*, elem->'permissions' AS user_group_node_permissions
-- 	FROM user_group, jsonb_array_elements(user_group.node_permissions) elem
-- 	WHERE (elem->>'id' = node.id::text AND user_group.id = ANY('{2,3}') AND elem->'permissions' @> '12') 
-- 	OR (elem->'permissions' = NULL AND 12 = ANY(user_group.default_node_permissions))
-- ) as test


SELECT DISTINCT ON (my_node.id) my_node.*
FROM user_group AS my_user_group,
LATERAL
(
	SELECT node.*, elem->'permissions' AS user_group_node_permissions
	FROM node
	LEFT OUTER JOIN
	jsonb_array_elements(my_user_group.node_permissions) elem
	ON elem->>'id' = node.id::text
	--ORDER BY node.id
)my_node
WHERE (my_user_group.id = ANY('{2,3}')) 
AND (user_group_node_permissions @> '12' OR (user_group_node_permissions IS NULL AND 12 = ANY(my_user_group.default_node_permissions)));
--LIMIT 20;




-- SELECT DISTINCT ON (my_node.id) my_node.*
-- FROM user_group AS my_user_group,
-- LATERAL
-- (
-- 	SELECT node.*, elem->'permissions' AS user_group_node_permissions
-- 	FROM node
-- 	LEFT OUTER JOIN
-- 	jsonb_array_elements(my_user_group.node_permissions) elem
-- 	ON elem->>'id' = node.id::text
-- 	ORDER BY node.id
-- )my_node
-- WHERE (my_user_group.id = ANY('{2,3}')) 
-- AND (user_group_node_permissions @> '12' OR (user_group_node_permissions IS NULL AND 12 = ANY(my_user_group.default_node_permissions)));

--SELECT * from user_group;

--update user_group set default_node_permissions = '{1,2,3,4,5,7,9,10,12}' WHERE id=2;
--update user_group set node_permissions = '[{"id": 10, "permissions": [1, 9, 12]}, {"id": 11, "permissions": [1, 9]}]' WHERE id=2;
--SELECT * from user_group;

-- SELECT my_node.*, my_user_group.default_node_permissions FROM
-- user_group AS my_user_group,
-- LATERAL
-- (
-- 	SELECT node.*, elem->'permissions' AS group_node_permissions
-- 	FROM node
-- 	LEFT OUTER JOIN
-- 	jsonb_array_elements(my_user_group.node_permissions) elem
-- 	ON elem->>'id' = node.id::text
-- 	ORDER BY node.id
-- )my_node
-- WHERE my_user_group.id = 3 AND (group_node_permissions @> '12' OR (group_node_permissions IS NULL AND 12 = ANY(my_user_group.default_node_permissions)));

-- SELECT my_node.*, my_user_group.default_node_permissions FROM
-- user_group AS my_user_group,
-- LATERAL
-- (
-- 	SELECT node.*, replace(replace(elem->>'permissions', '[', '{'),']','}')::integer[] AS group_node_permissions
-- 	FROM node
-- 	LEFT OUTER JOIN
-- 	jsonb_array_elements(my_user_group.node_permissions) elem
-- 	ON elem->>'id' = node.id::text
-- 	ORDER BY node.id
-- )my_node 
-- WHERE my_user_group.id = 3 AND (12 = ANY(group_node_permissions) OR (group_node_permissions IS NULL AND 12 = ANY(my_user_group.default_node_permissions)));

