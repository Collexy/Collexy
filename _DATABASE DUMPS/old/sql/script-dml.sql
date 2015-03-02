-- node
-- guids:
-- 1) document
-- 2) media
-- 3) template
-- 4) content_type
-- 5) root
-- 6) content_item
-- 7) media type
-- 8) recycle bin
-- 9) stylesheet
-- 10) script
-- 11) data_type

-- =?----- ROOT
INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1', 1, 'root', 5);

-- data_type
INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1.2', 1, 'Text input', 11);

-- content_type
INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1.3', 1, 'Master', 4);

INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1.3.4', 1, 'Home', 4);
    
INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1.3.5', 1, 'Page', 4);

-- template
INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.6', 1, 'Layout', 3);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.6.7', 1, 'Home', 3);

INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1.6.8', 1, 'Page', 3);

-- document
INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.9', 1, 'Home', 1);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.9.10', 1, 'Sample Page', 1);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.9.10.11', 1, 'Child Page Level 1', 1);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.9.10.11.12', 1, 'Child Page Level 2', 1);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.9.12', 1, 'Another Page', 1);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.14', 1, 'Textarea', 11);

-- content type 
INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.15', 1, 'Image', 4);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.16', 1, 'Folder', 4);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.17', 1, 'File', 4);

-- media
INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.18', 1, 'gopher.jpg', 2);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.19', 1, 'postgresql.png', 2);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.20', 1, 'Sidebar 1', 3);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.21', 1, 'Sidebar 2', 3);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.6.22', 1, 'Page with sidebars', 3);

-- =?----- DATA TYPE
INSERT INTO data_type(
            node_id, html)
    VALUES (2, '<input type="text" id="%v" class="%v" placeholder="%v" value="%v" ng-model="%v">');

INSERT INTO template (node_id, name, html)
  VALUES (6, 'Home', '<!DOCTYPE html><html><head><title>Some title</title></head><body><h1>Layout template html</h1><p>This should eventually just be some placeholder containers and template tags</p></body></html>');

INSERT INTO template (node_id, name, html, parent_template_node_id)
  VALUES (7, 'Home', '<h2>Home layout template</h2>', 6);

INSERT INTO template (node_id, name, html, parent_template_node_id)
  VALUES (8, 'Page', '<h2>Page layout template</h2>', 6);

INSERT INTO template (node_id, name)
  VALUES (20, 'Sidebar 1');

INSERT INTO template (node_id, name)
  VALUES (21, 'Sidebar 2');

INSERT INTO template (node_id,parent_template_node_id,name, partial_templates)
  VALUES (22, 6, 'Page with sidebars', '[{"id": 4, "name": "Sidebar 1", "node_id": 20}, {"id": 5, "name": "Sidebar 2", "node_id": 21}]');




INSERT INTO content_type(
            node_id, name, description, icon, thumbnail,
            tabs)
    VALUES (3, 'Master', 'Some description', 'fa fa-folder-o', 'fa fa-folder-o',
            '[{"name": "tab1","properties": [{"name": "prop1","order": 1,"help_text": "help text","description": "description","data_type": 1}]},{"name": "tab2","properties": [{"name": "prop2","order": 1,"help_text": "help text2","description": "description2","data_type": 1},{"name": "prop3","order": 2,"help_text": "help text3","description": "description3","data_type": 1}]}]');

INSERT INTO content_type(
            node_id, name, description, icon, thumbnail, master_content_type_node_id,
            tabs, meta)
    VALUES (4, 'Home', 'Home Some description', 'fa fa-folder-o', 'fa fa-folder-o', 3,
            ARRAY['{"name": "social","properties": [{"name": "facebook","order": 1,"help_text": "fb social help text","description": "fb social description","data_type": 1}, {"name": "google_plus","order": 2,"help_text": "g+ social help text","description": "g+ social description","data_type": 1}]}'::jsonb]
  , '{"allowed_content_types_node_id": [4], "template_node_id": 7,"allowed_templates_node_id": [7]}'::jsonb);

INSERT INTO content_type(
            node_id, name, description, icon, thumbnail, master_content_type_node_id, meta)
    VALUES (5, 'Page', 'Page content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 3, '{"allowed_content_types_node_id": [5], "template_node_id": 8,"allowed_templates_node_id": [8]}'::jsonb);

INSERT INTO content_type(
            node_id, name, description, icon, thumbnail, tabs)
    VALUES (15, 'Image', 'Image content type description', 'fa fa-folder-o', 'fa fa-folder-o',
            '[{"name": "Image", "properties": [{"name": "url", "order": 1, "data_type": 1, "help_text": "help text", "description": "URL goes here."},{"name": "title", "order": 2, "data_type": 1, "help_text": "help text", "description": "The title entered here can override the above one."}, {"name": "caption", "order": 3, "data_type": 2, "help_text": "help text", "description": "Caption goes here."}, {"name": "alt", "order": 4, "data_type": 1, "help_text": "help text", "description": "Alt goes here."}, {"name": "description", "order": 5, "data_type": 2, "help_text": "help text", "description": "Description goes here."}]},{"name": "Properties", "properties": [{"name": "temporary property", "order": 1, "data_type": 1, "help_text": "help text", "description": "Temporary description goes here."}]}]');

INSERT INTO content (node_id, content_type_node_id)
  VALUES (9, 4);

INSERT INTO content (node_id, content_type_node_id)
  VALUES (10, 5);

INSERT INTO content (node_id, content_type_node_id)
  VALUES (15, 5);

INSERT INTO content (node_id, content_type_node_id, meta)
  VALUES (18, 15, '{"url": "/media/2014/10/gopher.jpg", "caption": "This is the caption of the gopher image", "alt": "Gopher image alt text", "description": "Gopher image description"}');

INSERT INTO content (node_id, content_type_node_id)
  VALUES (19, 15);

-- ALTER SEQUENCE content_id_seq RESTART WITH 1;
-- UPDATE content SET id = DEFAULT;