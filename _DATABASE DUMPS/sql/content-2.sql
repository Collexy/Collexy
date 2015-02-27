-- select * from content JOIN node ON content.node_id = node.id;


-- update content set meta='{"facebook":"http://facebook.com/profileid", "google_plus":"http://plus.google.com/profileid"}' where id=1;


-- select * from content


-- SELECT * FROM node
-- JOIN content
-- ON content.node_id = node.id
-- JOIN content_type as ct
-- ON ct.node_id = content.content_type_node_id
-- JOIN content_type as ctm
-- ON ctm.node_id = ct.master_content_type_node_id

-- select * from content_type

SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
    content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
    ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
    ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabs, ct.meta as ct_meta,
    ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.master_content_type_node_id as ctm_master_content_type_node_id, ctm.name as ctm_name,
    ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta as ctm_meta
    FROM node
    JOIN content
    ON content.node_id = node.id
    JOIN content_type as ct
    ON ct.node_id = content.content_type_node_id
    JOIN content_type as ctm
    ON ctm.node_id = ct.master_content_type_node_id
    WHERE node.id = 8
-- 
-- update content_type set master_content_type_node_id = null where id = 1;
-- select * from content_type