-- SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
--     content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
--     ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
--     ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabs, ct.meta as ct_meta,
--     ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.master_content_type_node_id as ctm_master_content_type_node_id, ctm.name as ctm_name,
--     ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta as ctm_meta,
--     data_type.id as data_type_id, data_type.node_id as data_type.node_id, data_type.html as data_type_html
--     FROM node
--     JOIN content
--     ON content.node_id = node.id
--     JOIN content_type as ct
--     ON ct.node_id = content.content_type_node_id
--     JOIN content_type as ctm
--     ON ctm.node_id = ct.master_content_type_node_id
--     JOIN data_type
--     ON data_type.node_id = content_type.tabs->'data_type'
--     WHERE node.id = 10

-- select * from content_type where tabs::json->>'data_type' = '1'

--select tabs from content_type where id=1;
--select * from json_to_recordset('[{"name": "tab1", "properties": [{"name": "prop1", "order": 1, "data_type": 1, "help_text": "help text", "description": "description"}]}]') as x(name text, properties json)
--select * from json_to_record('{"name": "tab1", "properties": [{"name": "prop1", "order": 1, "data_type": 1, "help_text": "help text", "description": "description"}]}') as x(name text, properties json) where x.properties->'data_type' = 1;
--select * from data_type

-- SELECT tabs, tabs->>'name' as namelol
-- FROM 
-- content_type 
-- JOIN data_type ON (data_type.id = content_type.tabs->>'id')

-- select row_to_json(row)
-- from (
--     select u.*, urd AS user_role
--     from users u
--     inner join (
--         select ur.*, d
--         from user_roles ur
--         inner join role_duties d on d.id = ur.duty_id
--     ) urd(id,name,description,duty_id,duty) on urd.id = u.user_role_id
-- ) row;

-- select row_to_json(row)
-- from (
-- SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
--     content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
--     ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
--     ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabs, ct.meta as ct_meta,
--     ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.master_content_type_node_id as ctm_master_content_type_node_id, ctm.name as ctm_name,
--     ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta as ctm_meta
--     FROM node
--     JOIN content
--     ON content.node_id = node.id
--     JOIN content_type as ct
--     ON ct.node_id = content.content_type_node_id
--     JOIN content_type as ctm
--     ON ctm.node_id = ct.master_content_type_node_id
--     WHERE node.id = 10
-- )row 
-- JOIN data_type
-- ON
-- row.row_to_json::jsonb @> '{"node_id":10}'::jsonb

--select * from json_to_recordset('[{"a":1,"b":"foo"},{"a":"2","c":"bar"}]') as x(a int, b text);
--select * from json_to_recordset('[{"name":"tab1", "properties": [{"name": "prop1", "order": 1, "data_type": 1},{"name": "prop2", "order": 2, "data_type": 1}]}]') as x(name text, properties json)
-- select properties
-- from(
-- 	select * 
-- 	from (
-- 		select * 
-- 		from json_to_recordset('[{"name":"tab1", "properties": [{"name": "prop1", "order": 1, "data_type": 1},{"name": "prop2", "order": 2, "data_type": 1}]}]') 
-- 		as x(name text, properties json)
-- 	)lol
-- )oki

-- CREATE OR REPLACE FUNCTION public.json_merge(data json, merge_data json)
-- RETURNS json
-- LANGUAGE sql
-- AS $$
-- SELECT ('{'||string_agg(to_json(key)||':'||value, ',')||'}')::json
-- FROM (
-- WITH to_merge AS (
-- SELECT * FROM json_each(merge_data)
-- )
-- SELECT *
-- FROM json_each(data)
-- WHERE key NOT IN (SELECT key FROM to_merge)
-- UNION ALL
-- SELECT * FROM to_merge
-- ) t;
-- $$;
-- 
-- SELECT json_merge('[{"name":"tab1", "properties": [{"name": "prop1", "order": 1, "data_type": 1},{"name": "prop2", "order": 2, "data_type": 1}]}]', '[{"name":"tab1", "properties": [{"name": "prop1", "order": 1, "data_type": 1},{"name": "prop2", "order": 2, "data_type": 2}]}]');

-- select row_to_json(row)
-- from (
-- SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
--     content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
--     ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
--     ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabs, ct.meta as ct_meta,
--     ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.master_content_type_node_id as ctm_master_content_type_node_id, ctm.name as ctm_name,
--     ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta as ctm_meta
--     FROM node
--     JOIN content
--     ON content.node_id = node.id
--     JOIN content_type as ct
--     ON ct.node_id = content.content_type_node_id
--     JOIN content_type as ctm
--     ON ctm.node_id = ct.master_content_type_node_id
--     --WHERE node.id = 10
-- )row
-- where row.row_to_json::jsonb @> '{"node_id":10}'::jsonb
-- 

--drop function populate();
-- CREATE OR REPLACE FUNCTION populate() RETURNS json AS $$
-- DECLARE
--     -- declarations
--     ret RECORD;
--     js  JSON;
-- BEGIN
--     SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
--     content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
--     ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
--     ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabs, ct.meta as ct_meta,
--     ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.master_content_type_node_id as ctm_master_content_type_node_id, ctm.name as ctm_name,
--     ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta as ctm_meta
--     FROM node
--     JOIN content
--     ON content.node_id = node.id
--     JOIN content_type as ct
--     ON ct.node_id = content.content_type_node_id
--     JOIN content_type as ctm
--     ON ctm.node_id = ct.master_content_type_node_id
--     WHERE node.id = 10
--     INTO ret;
--     SELECT tabs from ret INTO js;
--     RETURN js;
-- END;
-- $$ LANGUAGE plpgsql;
-- 
-- select populate();

-- select row_to_json(row)
-- from (
-- SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
--     content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
--     ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
--     ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabs, ct.meta as ct_meta,
--     ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.master_content_type_node_id as ctm_master_content_type_node_id, ctm.name as ctm_name,
--     ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta as ctm_meta
--     FROM node
--     JOIN content
--     ON content.node_id = node.id
--     JOIN content_type as ct
--     ON ct.node_id = content.content_type_node_id
--     JOIN content_type as ctm
--     ON ctm.node_id = ct.master_content_type_node_id
--     WHERE node.id = 10
-- )row

-- WITH t(data) AS (
--    VALUES ('{"objects":[{"src":"foo.png"}
--                        ,{"src":"bar.png"}]
--            ,"background":"background.png"}'::json)
--    ) 
-- SELECT *
-- FROM   t, json_array_elements(t.data#>'{objects}') AS o
-- WHERE  o->>'src' = 'foo.png';


-- select row_to_json(row)
-- from (
-- SELECT
--     content_type.id, content_type.tabs
--     FROM content_type
--     WHERE id=1
-- )row


--SELECT * from( 

--select * from json_populate_recordset(null::myrowtype, '[{"a":1,"b":2},{"a":3,"b":4}]')

    --SELECT * FROM
      --  json_populate_recordset(null::myrowtype,'[{"name": "tab1", "properties": [{"name": "prop1", "order": 1, "data_type": 1, "help_text": "help text", "description": "description"}]}, {"name": "tab2", "properties": [{"name": "prop2", "order": 1, "data_type": 1, "help_text": "help text2", "description": "description2"}, {"name": "prop3", "order": 2, "data_type": 1, "help_text": "help text3", "description": "description3"}]}]')
-- select row_to_json(row)
-- from (
-- select oki.*,data_type.node_id as data_type_node_id, data_type.html as data_type_html
-- from(
--     select * 
--     from json_to_recordset(
--         (select properties from(
--             (select * 
--                 from json_to_recordset(
--                     (SELECT content_type.tabs 
--                     FROM content_type 
--                     WHERE id=1)
--                 ) as x(name text, properties json)
--                 where name='tab2'
--             )
--         )ghgh)
--     )as x(name text, "order" int, data_type int, help_text text, description text)
-- )oki
-- JOIN data_type
-- ON data_type.id = oki.data_type
-- )row


-- select row_to_json(row)
-- from (
-- select oki.*,data_type.node_id as data_type_node_id, data_type.html as data_type_html
-- from(
--     select * 
--     from json_to_recordset(
--         (select properties from(
--             (select * 
--                 from json_to_recordset(
--                     (
-- 		    SELECT coalesce(ct_tabs,ctm_tabs) as tabs
-- 			FROM 
-- 			(	
-- 			    SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
-- 			    content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
-- 			    ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
-- 			    ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabs, ct.meta as ct_meta,
-- 			    ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.master_content_type_node_id as ctm_master_content_type_node_id, ctm.name as ctm_name,
-- 			    ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta as ctm_meta
-- 			    FROM node
-- 			    JOIN content
-- 			    ON content.node_id = node.id
-- 			    JOIN content_type as ct
-- 			    ON ct.node_id = content.content_type_node_id
-- 			    JOIN content_type as ctm
-- 			    ON ctm.node_id = ct.master_content_type_node_id
-- 			    WHERE node.id = 10
-- 			) dsfds
--                     )
--                 ) as x(name text, properties json)
--                 where name='tab2'
--             )
--         )ghgh)
--     )as x(name text, "order" int, data_type int, help_text text, description text)
-- )oki
-- JOIN data_type
-- ON data_type.id = oki.data_type
-- )row


-- select json_agg(row)
-- from (
-- select oki.*,data_type.html as data_type_html
-- from(
--     select * 
--     from json_to_recordset(
--         (select properties from(
--             (select * 
--                 from json_to_recordset(
--                     (
-- 		    SELECT tabs
-- 			FROM 
-- 			(	
-- 			    SELECT 
-- 			    content_type.id, content_type.tabs
-- 			    FROM content_type
-- 			    WHERE id=1
-- 			) dsfds
--                     )
--                 ) as x(name text, properties json)
--                 where name='tab2'
--             )
--         )ghgh)
--     )as x(name text, "order" int, data_type int, help_text text, description text)
-- )oki
-- JOIN data_type
-- ON data_type.id = oki.data_type
-- )row
-- 

-- SELECT content_type.id, content_type.tabs as original, ss as extended 
-- FROM content_type,
-- LATERAL (
-- 	select json_agg(row)
-- 	from (
-- 		select oki.*,data_type.html as data_type_html
-- 		from(
-- 		    select * 
-- 		    from json_to_recordset(
-- 			(select properties from(
-- 			    (select * 
-- 				from json_to_recordset(
-- 				    (
-- 					SELECT json_agg(ggg)
-- 					from(
-- 						SELECT tabs
-- 						FROM 
-- 						(	
-- 						    SELECT 
-- 						    content_type.id, content_type.tabs
-- 						    FROM content_type as ct
-- 						    WHERE ct.id=content_type.id
-- 						) dsfds
-- 					)ggg
-- 				    )
-- 				) as x(name text, properties json)
-- 				--where name='tab2'
-- 			    )
-- 			)ghgh)
-- 		    )as x(name text, "order" int, data_type int, help_text text, description text)
-- 		)oki
-- 		JOIN data_type
-- 		ON data_type.id = oki.data_type
-- 	)row
-- ) ss;



-- select * 
-- from json_to_recordset(
--      (
-- 	--SELECT json_agg(ggg)
-- 	--from(
-- 		SELECT tabs
-- 		FROM 
-- 		(	
-- 		    SELECT 
-- 		    *
-- 		    FROM content_type
-- 		    WHERE content_type.id=1
-- 		) dsfds
-- 	--)ggg
--      )
-- ) as x(name text, properties json)


-- SELECT content_type.id, content_type.tabs as original, ss.json_agg as extended 
-- FROM content_type,
-- LATERAL (
-- 	select json_agg(row)
-- 	from (
-- 		select oki.*,data_type.html as data_type_html
-- 		from(
-- 		    select * 
-- 		    from json_to_recordset(
-- 			(select properties from(
-- 			    (select * 
-- 				from json_to_recordset(
-- 				    (
-- 					SELECT json_agg(ggg)
-- 					from(
-- 						SELECT tabs
-- 						FROM 
-- 						(	
-- 						    SELECT 
-- 						    *
-- 						    FROM content_type as ct
-- 						    WHERE ct.id=content_type.id
-- 						) dsfds
-- 					)ggg
-- 				    )
-- 				) as x(name text, properties json)
-- 				--where name='tab2'
-- 			    )
-- 			)ghgh)
-- 		    )as x(name text, "order" int, data_type int, help_text text, description text)
-- 		)oki
-- 		JOIN data_type
-- 		ON data_type.id = oki.data_type
-- 	)row
-- ) ss;


-- select name,properties
-- from json_to_recordset(
-- (
-- 	select * 
-- 	from json_to_recordset(
-- 		(
-- 			SELECT json_agg(ggg)
-- 			from(
-- 				SELECT tabs
-- 				FROM 
-- 				(	
-- 				    SELECT 
-- 				    *
-- 				    FROM content_type
-- 				    WHERE content_type.id=1
-- 				) dsfds
-- 			)ggg
-- 		)
-- 	) as x(tabs json)
-- )
-- ) as y(name text, properties json)



-- --select json_agg(row2) from((
-- SELECT content_type.id, content_type.tabs as original, gf.json_agg as new_tabs
-- FROM content_type,
-- LATERAL (
-- 	select json_agg(row1) from((
-- 	select y.name, ss.extended_properties
-- 	from json_to_recordset(
-- 		(
-- 			select * 
-- 			from json_to_recordset(
-- 				(
-- 					SELECT json_agg(ggg)
-- 					from(
-- 						SELECT tabs
-- 						FROM 
-- 						(	
-- 						    SELECT 
-- 						    *
-- 						    FROM content_type as ct
-- 						    WHERE ct.id=content_type.id
-- 						) dsfds
-- 					)ggg
-- 				)
-- 			) as x(tabs json)
-- 		)
-- 	) as y(name text, properties json),
-- 	LATERAL (
-- 		select json_agg(row) as extended_properties
-- 		from(
-- 			select name, "order", data_type, data_type.html as data_type_html, help_text, description
-- 			from json_to_recordset(properties) 
-- 			as k(name text, "order" int, data_type int, help_text text, description text)
-- 			JOIN data_type
-- 			ON data_type.id = k.data_type
-- 			)row
-- 	) ss
-- 	))row1
-- ) gf
-- --))row2



-- SELECT *
-- FROM 
-- (
-- 	SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
-- 	content.id as content_id, content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta::json as content_meta,
-- 	ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
-- 	ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabss, ct.meta::json as ct_meta,
-- 	ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.name as ctm_name,
-- 	ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta::json as ctm_meta
-- 	FROM node
-- 	JOIN content
-- 	ON content.node_id = node.id
-- 	JOIN content_type as ct
-- 	ON ct.node_id = content.content_type_node_id
-- 	JOIN content_type as ctm
-- 	ON ctm.node_id = ct.master_content_type_node_id
-- 	WHERE node.id=9
-- )noden,
-- (
-- --select json_agg(row2) from((
-- SELECT content_type.id as content_type_id, content_type.master_content_type_node_id as master_content_type_node_id, gf.json_agg as ct_tabs
-- FROM content_type,
-- LATERAL (
-- 	select json_agg(row1) from((
-- 	select y.name, ss.extended_properties
-- 	from json_to_recordset(
-- 		(
-- 			select * 
-- 			from json_to_recordset(
-- 				(
-- 					SELECT json_agg(ggg)
-- 					from(
-- 						SELECT tabs
-- 						FROM 
-- 						(	
-- 						    SELECT *
-- 						    FROM content_type as ct
-- 						    WHERE ct.id=content_type.id
-- 						) dsfds
-- 
-- 					)ggg
-- 				)
-- 			) as x(tabs json)
-- 		)
-- 	) as y(name text, properties json),
-- 	LATERAL (
-- 		select json_agg(row) as extended_properties
-- 		from(
-- 			select name, "order", data_type, data_type.html as data_type_html, help_text, description
-- 			from json_to_recordset(properties) 
-- 			as k(name text, "order" int, data_type int, help_text text, description text)
-- 			JOIN data_type
-- 			ON data_type.id = k.data_type
-- 			)row
-- 	) ss
-- 	UNION all((
-- 	select p.name, ss2.extended_properties
-- 	from json_to_recordset(
-- 		(
-- 			select * 
-- 			from json_to_recordset(
-- 				(
-- 					SELECT json_agg(ggg)
-- 					from(
-- 						SELECT tabs
-- 						FROM 
-- 						(	
-- 						    SELECT *
-- 						    FROM content_type as ctm
-- 						    WHERE ctm.node_id=content_type.master_content_type_node_id
-- 						) dsfds
-- 
-- 					)ggg
-- 				)
-- 			) as x(tabs json)
-- 		)
-- 	) as p(name text, properties json),
-- 	LATERAL (
-- 		select json_agg(row) as extended_properties
-- 		from(
-- 			select name, "order", data_type, data_type.html as data_type_html, help_text, description
-- 			from json_to_recordset(properties) 
-- 			as k(name text, "order" int, data_type int, help_text text, description text)
-- 			JOIN data_type
-- 			ON data_type.id = k.data_type
-- 			)row
-- 	) ss2
-- 	))
-- 	))row1
-- ) gf
-- ) rgr
-- JOIN content_type as lollo
-- ON lollo.node_id = rgr.master_content_type_node_id
-- where rgr.content_type_id = ct_id;
-- --))row2



SELECT *
FROM 
(
	SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
	content.id as content_id, content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
	ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
	ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.meta::json as ct_meta,
	ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.name as ctm_name,
	ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.meta::json as ctm_meta
	FROM node
	JOIN content
	ON content.node_id = node.id
	JOIN content_type as ct
	ON ct.node_id = content.content_type_node_id
	JOIN content_type as ctm
	ON ctm.node_id = ct.master_content_type_node_id
	WHERE node.id=10
)noden,
(
SELECT content_type.id as content_type_id, content_type.master_content_type_node_id as master_content_type_node_id, gf.json_agg as ct_tabs, gf2.json_agg as ctm_tabs
FROM content_type,
LATERAL (
	select json_agg(row1) from((
	select y.name, ss.properties
	from json_to_recordset(
		(
			select * 
			from json_to_recordset(
				(
					SELECT json_agg(ggg)
					from(
						SELECT tabs
						FROM 
						(	
						    SELECT *
						    FROM content_type as ct
						    WHERE ct.id=content_type.id
						) dsfds

					)ggg
				)
			) as x(tabs json)
		)
	) as y(name text, properties json),
	LATERAL (
		--select json_agg(row) as extended_properties
		select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id',row.data_type, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
		from(
			select name, "order", data_type, data_type.html as data_type_html, help_text, description
			from json_to_recordset(properties) 
			as k(name text, "order" int, data_type int, help_text text, description text)
			JOIN data_type
			ON data_type.id = k.data_type
			)row
	) ss
	))row1
) gf,
LATERAL (
	select json_agg(row2) from((
	select p.name, ss2.properties
	from json_to_recordset(
		(
			select * 
			from json_to_recordset(
				(
					SELECT json_agg(ggg)
					from(
						SELECT tabs
						FROM 
						(	
						    SELECT *
						    FROM content_type as ctm
						    WHERE ctm.node_id=content_type.master_content_type_node_id
						) dsfds

					)ggg
				)
			) as x(tabs json)
		)
	) as p(name text, properties json),
	LATERAL (
		select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id',row.data_type, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
		--select json_agg(row) as extended_properties
		from(
			select name, "order", data_type, data_type.html as data_type_html, help_text, description
			from json_to_recordset(properties) 
			as k(name text, "order" int, data_type int, help_text text, description text)
			JOIN data_type
			ON data_type.id = k.data_type
			)row
	) ss2
	))row2
	)gf2
) rgr

where rgr.content_type_id = ct_id;


--update content set meta='{"page_title":"Home page title", "site_name": "Collexy cms test site", "site_tagline": "Test site tagline", "copyright": "&copy; 2014 codeish.com", "prop2": "Home page prop 2", "facebook": "facebook.com/home"}' where node_id=9;

--update content_type set tabs='[{"name": "Content", "properties": [{"name": "site_name", "order": 2, "data_type": 1, "help_text": "help text", "description": "Site name goes here."}, {"name": "site_tagline", "order": 3, "data_type": 1, "help_text": "help text", "description": "Site tagline goes here."}, {"name": "copyright", "order": 4, "data_type": 1, "help_text": "help text", "description": "Copyright here."}]},{"name": "Social", "properties": [{"name": "facebook", "order": 1, "data_type": 1, "help_text": "help text", "description": "Enter your facebook link here."},{"name": "twitter", "order": 1, "data_type": 1, "help_text": "help text", "description": "Enter your twitter link here."},{"name": "linkedin", "order": 1, "data_type": 1, "help_text": "help text", "description": "Enter your linkedin link here."}]}]' where id=2;
select * from content_type;
--select* from node
--update node set path='1.9.13' where id=13
--alter table content alter column meta type jsonb using meta::jsonb;
--update content set meta = '{"page_title": "Home page title", "prop3": "content prop 3", "facebook": "facebook.com/profile"}' where id = 1;
--update data_type set html = '<input type="text" id="{{prop.name}}" ng-model="node.content.meta[prop.name]">';
--update content_type set tabs = '[{"name": "Content", "properties": [{"name": "page_content", "order": 2, "data_type": 2, "help_text": "help text", "description": "Page description goes here."}]}]' where id=3;
--select * from content_type;