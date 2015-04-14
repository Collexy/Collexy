--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-04-14 19:24:29

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

--
-- TOC entry 2337 (class 0 OID 98246)
-- Dependencies: 193
-- Data for Name: content; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (1, '1', NULL, 'Home', 'home', 1, '2015-03-27 21:22:51.805', 2, '{"title": "Home", "domains": ["localhost:8080", "localhost:8080/test"], "copyright": "&copy; 2014 codeish.com", "site_name": "Collexy test site", "about_text": "<p>This is <strong>TXT</strong>, yet another free responsive site template designed by <a href=\"http://n33.co\">AJ</a> for <a href=\"http://html5up.net\">HTML5 UP</a>. It is released under the <a href=\"http://html5up.net/license/\">Creative Commons Attribution</a> license so feel free to use it for whatever you are working on (personal or commercial), just be sure to give us credit for the design. That is basically it :)</p>", "about_title": "About title here", "banner_link": "http://somelink.test", "hide_banner": false, "hide_in_nav": false, "is_featured": false, "template_id": 2, "site_tagline": "Test site tagline", "banner_header": "Banner header goes here", "facebook_link": "facebook.com/home", "banner_link_text": "Click Here!", "banner_subheader": "Banner subheader goes here", "banner_background_image": "/media/Sample Images/TXT/banner.jpg"}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (2, '1.2', 1, 'Welcome', 'welcome', 1, '2015-03-27 21:31:55.462', 5, '{"image": "/media/Sample Images/TXT/pic01.jpg", "title": "Welcome", "content": "Welcome content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3}', NULL, '[{"id": 2, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]', NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (3, '1.3', 1, 'Getting Started', 'getting_started', 1, '2015-03-27 21:46:13.265', 5, '{"image": "/media/Sample Images/TXT/pic02.jpg", "title": "Getting Started", "content": "Getting Started content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3}', NULL, NULL, '[{"id": 1, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]', 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (4, '1.4', 1, 'Documentation', 'documentation', 1, '2015-03-27 21:50:23.197', 5, '{"image": "/media/Sample Images/TXT/pic03.jpg", "title": "Documentation", "content": "Documentation content goes here1", "hide_in_nav": false, "is_featured": true, "template_id": 3}', '{"groups": [1], "members": [1]}', NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (5, '1.5', 1, 'Get Involved', 'get_involved', 1, '2015-03-27 21:51:57.503', 5, '{"image": "/media/Sample Images/TXT/pic04.jpg", "title": "Get Involved", "content": "Get Involved content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (6, '1.6', 1, 'Posts', 'posts', 1, '2015-03-27 21:54:10.787', 4, '{"title": "Posts", "hide_in_nav": false, "is_featured": true, "template_node_id": 5}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (10, '1.6.10', 6, 'Amazing Post', 'amazing_post', 1, '2015-03-27 22:05:14.042', 3, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Amazing Post", "content": "<p>What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post.</p>", "sub_header": "Amazing subheader here", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (11, '1.11', 1, 'Sample Images', 'sample_images', 1, '2015-03-27 22:08:29.415', 6, '{"path": "media\\Sample Images"}', NULL, NULL, NULL, 2);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (14, '1.11.12.14', 12, 'pic02.jpg', 'pic2', 1, '2015-03-27 22:12:24.478', 7, '{"alt": "pic02.jpg", "path": "media\\Sample Images\\TXT\\pic02.jpg", "title": "pic02.jpg", "caption": "pic02.jpg", "description": "pic02.jpg"}', NULL, NULL, NULL, 2);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (12, '1.11.12', 11, 'TXT', 'txt', 1, '2015-03-27 22:09:40.207', 6, '{"path": "media\\Sample Images\\TXT"}', NULL, NULL, NULL, 2);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (13, '1.11.12.13', 12, 'pic01.jpg', 'pic1', 1, '2015-03-27 22:10:35.745', 7, '{"alt": "pic01.jpg", "path": "media\\Sample Images\\TXT\\pic01.jpg", "title": "pic01.jpg", "caption": "pic01.jpg", "description": "pic01.jpg"}', NULL, NULL, NULL, 2);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (15, '1.11.12.15', 12, 'pic03.jpg', 'pic3', 1, '2015-03-27 22:13:10.64', 7, '{"alt": "pic03.jpg", "path": "media\\Sample Images\\TXT\\pic03.jpg", "title": "pic03.jpg", "caption": "pic03.jpg", "description": "pic03.jpg"}', NULL, NULL, NULL, 2);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (16, '1.11.12.16', 12, 'pic04.jpg', 'pic4', 1, '2015-03-27 22:13:35.245', 7, '{"alt": "pic04.jpg", "path": "media\\Sample Images\\TXT\\pic04.jpg", "title": "pic04.jpg", "caption": "pic04.jpg", "description": "pic04.jpg"}', NULL, NULL, NULL, 2);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (17, '1.11.12.17', 12, 'pic05.jpg', 'pic5', 1, '2015-03-27 22:14:05.966', 7, '{"alt": "pic05.jpg", "path": "media\\Sample Images\\TXT\\pic05.jpg", "title": "pic05.jpg", "caption": "pic05.jpg", "description": "pic05.jpg"}', NULL, NULL, NULL, 2);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (18, '1.11.12.18', 12, 'banner.jpg', 'banner', 1, '2015-03-27 22:14:35.241', 7, '{"alt": "banner.jpg", "path": "media\\Sample Images\\TXT\\banner.jpg", "title": "banner.jpg", "caption": "banner.jpg", "description": "banner.jpg"}', NULL, NULL, NULL, 2);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (19, '1.6.19', 6, 'Categories', 'categories', 1, '2015-03-27 22:17:32.659', 8, '{"title": "Categories", "content": "Categories", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (20, '1.6.19.20', 19, 'Category 1', 'category_1', 1, '2015-03-27 22:18:45.865', 9, '{"title": "Category 1", "content": "Category 1 content", "hide_in_nav": false, "is_featured": true, "template_id": 6}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (21, '1.21', 1, '404', '404', 1, '2015-03-27 22:20:10.169', 5, '{"title": "404", "content": "404 content goes here", "hide_in_nav": true, "is_featured": false, "template_id": 9}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (22, '1.22', 1, 'Login', 'login', 1, '2015-03-27 22:21:19.482', 5, '{"title": "Login", "content": "Login content goes here", "hide_in_nav": true, "is_featured": false, "template_id": 7}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (8, '1.6.8', 6, 'Txt Starter Kit For Collexy Released', 'collexy_starter_kit', 1, '2015-03-27 21:59:24.379', 3, '{"title": "TXT Starter Kit For Collexy Released", "content": "The collexy TXT starter kit is just awesome!", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (9, '1.6.9', 6, 'You Need To Read This', 'read_this', 1, '2015-03-27 22:03:09.422', 3, '{"title": "You Need To Read This", "content": "See - you really needed to read this post!", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, 1);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id) VALUES (7, '1.6.7', 6, 'Hello World', 'hello_world', 1, '2015-03-27 21:55:03.797', 3, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Hello World", "content": "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas vel tellus venenatis, iaculis eros eu, pellentesque felis. Mauris eleifend venenatis maximus. Fusce condimentum nulla augue, sed elementum nisl dictum ut. Sed ex arcu, efficitur eu finibus ac, convallis ut eros. Ut faucibus elit erat, ac venenatis velit cursus quis. Phasellus sapien elit, ullamcorper ac placerat at, consectetur eget ex. Integer augue sem, tempor nec hendrerit et, ullamcorper ut arcu.</p>\n\n<p>Pellentesque auctor et arcu at tristique. Suspendisse ipsum sapien, vulputate quis cursus eu, rhoncus sed nisi. Nulla euismod mauris vitae tellus iaculis convallis. Sed sodales, risus id sollicitudin aliquet, purus justo convallis dui, sit amet imperdiet elit mauris accumsan velit. Suspendisse dapibus sit amet quam in porta. Nam eleifend sodales dolor eget tempor. Sed pharetra aliquam dui, ultricies scelerisque orci luctus at. Proin eleifend neque quis dolor facilisis sollicitudin. Integer vel ligula nec metus sagittis lacinia at quis arcu. Sed in sem ut mauris laoreet euismod. Integer eu tincidunt lectus, nec varius libero. Proin nec interdum ex. Quisque non lacinia lectus, luctus molestie mi. Fusce lacus est, rhoncus sed nunc at, fermentum luctus ipsum.</p>\n\n<h3>Nunc pulvinar metus a erat fermentum bibendum</h3>\n\n<p>Phasellus mattis tempor dolor vitae feugiat. Sed aliquet massa nisi, in imperdiet mauris auctor in. Nam consectetur ut erat at suscipit. Integer faucibus eleifend rhoncus. Praesent vel bibendum elit, ut molestie metus. Maecenas efficitur, magna vel scelerisque pretium, magna elit vehicula massa, dignissim posuere felis enim a lectus. Donec eget semper urna. Praesent vel nisi id lacus tincidunt pretium vitae eu sapien. Duis varius nisi velit, nec maximus arcu blandit sit amet. Proin dapibus dui et elit dapibus, sit amet rhoncus nisl lobortis. Nunc pretium, lorem eu dignissim mollis, ex nisi mollis lectus, eu blandit arcu nisl vel elit. Mauris risus ipsum, elementum quis eleifend ut, venenatis sit amet orci. Donec ac orci aliquam, vulputate odio eget, pulvinar elit. Cras molestie urna eget justo hendrerit aliquam.</p>\n", "categories": [20], "sub_header": "Subheader for Hello World", "hide_in_nav": false, "is_featured": true, "template_id": 4, "date_published": "2015-16-03 20:55:38"}', NULL, NULL, NULL, 1);


--
-- TOC entry 2343 (class 0 OID 0)
-- Dependencies: 194
-- Name: content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_id_seq', 22, true);


--
-- TOC entry 2335 (class 0 OID 98227)
-- Dependencies: 191
-- Data for Name: content_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (1, '1', NULL, 'Master', 'master', 1, '2015-03-27 17:46:05.405', 'Master content type description', '', '', NULL, '[{"name": "Content", "properties": [{"name": "title", "order": 1, "help_text": "help text", "description": "The page title overrides the name the page has been given.", "data_type_id": 1}]}, {"name": "Properties", "properties": [{"name": "hide_in_nav", "order": 1, "help_text": "help text2", "description": "description2", "data_type_id": 18}]}]', 1);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (2, '1.2', 1, 'Home', 'home', 1, '2015-03-27 17:47:50.897', 'Home content type description', 'fa fa-home fa-fw', 'fa fa-home fa-fw', '{"template_id": 2, "allowed_template_ids": [2], "allowed_content_type_ids": [3, 4, 5]}', '[{"name": "Content", "properties": [{"name": "site_name", "order": 2, "help_text": "help text", "description": "Site name goes here.", "data_type_id": 1}, {"name": "site_tagline", "order": 3, "help_text": "help text", "description": "Site tagline goes here.", "data_type_id": 1}, {"name": "copyright", "order": 4, "help_text": "help text", "description": "Copyright here.", "data_type_id": 1}, {"name": "domains", "order": 5, "help_text": "help text", "description": "Domains goes here.", "data_type_id": 17}]}, {"name": "Social", "properties": [{"name": "facebook_link", "order": 1, "help_text": "help text", "description": "Enter your facebook link here.", "data_type_id": 1}, {"name": "twitter_link", "order": 2, "help_text": "help text", "description": "Enter your twitter link here.", "data_type_id": 1}, {"name": "linkedin_link", "order": 3, "help_text": "help text", "description": "Enter your linkedin link here.", "data_type_id": 1}, {"name": "google_link", "order": 4, "help_text": "help text", "description": "Enter your Google+ profile link here.", "data_type_id": 1}, {"name": "rss_link", "order": 5, "help_text": "help text", "description": "Enter your RSS feed link here.", "data_type_id": 1}]}, {"name": "Banner", "properties": [{"name": "hide_banner", "order": 1, "help_text": "help text2", "description": "description2", "data_type_id": 18}, {"name": "banner_header", "order": 2, "help_text": "help text", "description": "Banner header.", "data_type_id": 1}, {"name": "banner_subheader", "order": 3, "help_text": "help text", "description": "Banner subheader.", "data_type_id": 1}, {"name": "banner_link_text", "order": 4, "help_text": "help text", "description": "Banner link text.", "data_type_id": 1}, {"name": "banner_link", "order": 5, "help_text": "help text", "description": "Banner link should ideally use a content picker data type.", "data_type_id": 1}, {"name": "banner_background_image", "order": 6, "help_text": "help text", "description": "This should ideally use the upload data type.", "data_type_id": 1}]}, {"name": "About", "properties": [{"name": "about_title", "order": 1, "help_text": "help text", "description": "About title.", "data_type_id": 1}, {"name": "about_text", "order": 2, "help_text": "help text", "description": "About text.", "data_type_id": 19}]}]', 1);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (3, '1.3', 1, 'Post', 'post', 1, '2015-03-27 17:51:17.53', 'Post content type description', 'fa fa-file-text-o fa-fw', 'fa fa-file-text-o fa-fw', '{"template_id": 4, "allowed_template_ids": [4], "allowed_content_type_ids": [3]}', '[{"name": "Content", "properties": [{"name": "is_featured", "order": 2, "help_text": "help text2", "description": "description2", "data_type_id": 18}, {"name": "image", "order": 3, "help_text": "Help text for image", "description": "Image url", "data_type_id": 1}, {"name": "sub_header", "order": 4, "help_text": "Help text for subheader", "description": "Subheader description", "data_type_id": 1}, {"name": "content", "order": 5, "help_text": "Help text for post content", "description": "Post content description", "data_type_id": 19}, {"name": "categories", "order": 6, "help_text": "help text2", "description": "description2", "data_type_id": 12}, {"name": "date_published", "order": 7, "help_text": "help date picker with time", "description": "date picker w time", "data_type_id": 11}]}]', 1);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (4, '1.4', 1, 'Post Overview', 'post_overview', 1, '2015-03-27 17:53:03.252', 'Post Overview content type description', 'fa fa-newspaper-o fa-fw', 'fa fa-newspaper-o fa-fw', '{"template_id": 5, "allowed_templates_ids": [5], "allowed_content_type_ids": [3, 8]}', '[]', 1);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (5, '1.5', 1, 'Page', 'page', 1, '2015-03-27 17:54:15.03', 'Page content type description', 'fa fa-file-o fa-fw', 'fa fa-file-o fa-fw', '{"template_id": 3, "allowed_template_ids": [3, 7, 8, 9, 10], "allowed_content_type_ids": [5]}', '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for page contentent", "description": "Page content description", "data_type_id": 19}]}]', 1);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (6, '1.6', NULL, 'Folder', 'folder', 1, '2015-03-27 17:55:47.388', 'Folder media type description', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', '{"allowed_content_type_ids": [6, 7]}', '[{"name": "Folder", "properties": [{"name": "folder_browser", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 14}, {"name": "path", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 1}]}, {"name": "Properties"}]', 2);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (7, '1.7', NULL, 'Image', 'image', 1, '2015-03-27 17:57:48.335', 'Image media type description', 'fa fa-image fa-fw', 'fa fa-image fa-fw', NULL, '[{"name": "Image", "properties": [{"name": "path", "order": 1, "help_text": "help text", "description": "URL goes here.", "data_type_id": 1}, {"name": "title", "order": 2, "help_text": "help text", "description": "The title entered here can override the above one.", "data_type_id": 1}, {"name": "caption", "order": 3, "help_text": "help text", "description": "Caption goes here.", "data_type_id": 3}, {"name": "alt", "order": 4, "help_text": "help text", "description": "Alt goes here.", "data_type_id": 3}, {"name": "description", "order": 5, "help_text": "help text", "description": "Description goes here.", "data_type_id": 3}, {"name": "file_upload", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 15}]}, {"name": "Properties", "properties": [{"name": "temporary property", "order": 1, "help_text": "help text", "description": "Temporary description goes here.", "data_type_id": 1}]}]', 2);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (8, '1.8', 1, 'Categories', 'categories', 1, '2015-03-27 17:59:30.925', 'Categories content type description', 'fa fa-folder-open-o fa-fw', 'fa fa-folder-open-o fa-fw', '{"allowed_content_type_ids": [9]}', '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for category contentent", "description": "Category content description", "data_type_id": 19}]}]', 1);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id) VALUES (9, '1.9', 1, 'Category', 'category', 1, '2015-03-27 18:02:14.279', 'Category content type description', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', '{"template_id": 6, "allowed_template_ids": [6], "allowed_content_types_ids": [9]}', '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for category contentent", "description": "Category content description", "data_type_id": 19}]}]', 1);


--
-- TOC entry 2344 (class 0 OID 0)
-- Dependencies: 192
-- Name: content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_id_seq', 9, true);


--
-- TOC entry 2331 (class 0 OID 98158)
-- Dependencies: 187
-- Data for Name: data_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (1, '1', NULL, 'Text Input', 'text_input', 1, '2015-03-26 23:47:44.854', '<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (2, '2', NULL, 'Numeric Input', 'numeric_input', 1, '2015-03-26 23:47:44.854', '<input type="number" id="{{prop.name}}" ng-model="data.meta[prop.name]">');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (3, '3', NULL, 'Textarea', 'textarea', 1, '2015-03-26 23:47:44.854', '<textarea id="{{prop.name}}" ng-model="data.meta[prop.name]">');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (4, '4', NULL, 'Radiobox', 'radiobox', 1, '2015-03-26 23:47:44.854', '');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (5, '5', NULL, 'Dropdown', 'dropdown', 1, '2015-03-26 23:47:44.854', '');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (6, '6', NULL, 'Dropdown Multiple', 'dropdown_multiple', 1, '2015-03-26 23:47:44.854', '');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (7, '7', NULL, 'Checkbox List', 'checkbox_list', 1, '2015-03-26 23:47:44.854', '');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (8, '8', NULL, 'Label', 'label', 1, '2015-03-26 23:47:44.854', '');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (9, '9', NULL, 'Color Picker', 'color_picker', 1, '2015-03-26 23:47:44.854', '<colorpicker>The color picker data type is not implemented yet!</colorpicker>');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (10, '10', NULL, 'Date Picker', 'date_picker', 1, '2015-03-26 23:47:44.854', '');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (11, '11', NULL, 'Date Picker With Time', 'date_picker_time', 1, '2015-03-26 23:47:44.854', '<div class="well">
  <div id="datetimepicker1" class="input-append date">
    <input data-format="dd-MM-yyyy hh:mm:ss" type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]"></input>
    <span class="add-on">
      <i class="fa fa-calendar" data-time-icon="icon-time" data-date-icon="icon-calendar">
      </i>
    </span>
  </div>
</div>

<script type="text/javascript">
  $(function() {
    $("#datetimepicker1").datetimepicker({
      language: "en"
    });
  });
</script>');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (12, '12', NULL, 'Content Picker', 'content_picker', 1, '2015-03-26 23:47:44.854', '<!--<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">-->

<div ng-repeat="cn in contentNodes"><label><input type="checkbox" checklist-model="data.meta[prop.name]" checklist-value="cn.id"></label> {{cn.name}}</div>
<br>
<button type="button" ng-click="checkAll()">check all</button>
<button type="button" ng-click="uncheckAll()">uncheck all</button>');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (13, '13', NULL, 'Media Picker', 'media_picker', 1, '2015-03-26 23:47:44.854', '');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (14, '14', NULL, 'Folder Browser', 'folder_browser', 1, '2015-03-26 23:47:44.854', '<folderbrowser>This is an awesome folder browser (unimplemented datatype)</folderbrowser>');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (15, '15', NULL, 'Upload', 'upload', 1, '2015-03-26 23:47:44.854', '');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (16, '16', NULL, 'Upload Multiple', 'upload_multiple', 1, '2015-03-26 23:47:44.854', '<input type="file" file-input="test.files" multiple />
<button ng-click="upload()" type="button">Upload</button>
<li ng-repeat="file in test.files">{{file.name}}</li>


<!--<input type="file" onchange="angular.element(this).scope().filesChanged(this)" multiple />
<button ng-click="upload()">Upload</button>
<li ng-repeat="file in files">{{file.name}}</li>-->');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (17, '17', NULL, 'Domains', 'domains', 1, '2015-03-26 23:47:44.854', '<div>
    <input type="text"/> <button type="button">Add domain</button><br>
    <ul>
        <li ng-repeat="domain in data.meta[prop.name]">{{domain}}</li>
    </ul>
    <button type="button">Delete selected</button>
</div>');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (18, '18', NULL, 'True/False', 'true_false', 1, '2015-03-26 23:47:44.854', '<div><label><input type="checkbox" type="checkbox"
       ng-model="data.meta[prop.name]"
       [name="{{prop.name}}"]
       [ng-true-value="true"]
       [ng-false-value=""]
       [ng-change=""]></label> {{prop.name}}
</div>');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html) VALUES (19, '19', NULL, 'Richtext Editor', 'richtext_editor', 1, '2015-03-26 23:47:44.854', '<textarea ck-editor id="{{prop.name}}" name="{{prop.name}}" ng-model="data.meta[prop.name]" rows="10" cols="80"></textarea>');


--
-- TOC entry 2345 (class 0 OID 0)
-- Dependencies: 188
-- Name: data_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_type_id_seq', 19, true);


--
-- TOC entry 2316 (class 0 OID 98009)
-- Dependencies: 172
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member (id, username, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_id, member_group_ids) VALUES (1, 'default_member', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'default_member@mail.com', '{"comments": "default user comments"}', '2015-01-22 14:25:38.904', NULL, '2015-04-09 16:21:43.78', NULL, 1, 'ECQQPP7OVVWUQUILXCQQJ65ARRYQEGRIQFK3JKRULFC4UY4DIRXQ', 1, '{68}');


--
-- TOC entry 2328 (class 0 OID 98131)
-- Dependencies: 184
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_group (id, path, parent_id, name, alias, created_by, created_date) VALUES (1, '1', NULL, 'Authenticated Member', 'authenticated_member', 1, '2015-03-26 17:09:34.18');


--
-- TOC entry 2346 (class 0 OID 0)
-- Dependencies: 183
-- Name: member_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_group_id_seq', 1, true);


--
-- TOC entry 2347 (class 0 OID 0)
-- Dependencies: 173
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_id_seq', 1, true);


--
-- TOC entry 2329 (class 0 OID 98139)
-- Dependencies: 185
-- Data for Name: member_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, meta, tabs) VALUES (1, '1', NULL, 'Member', 'member', 1, '2015-03-26 19:56:03.85', 'This is the default member type for Collexy members.', 'fa fa-user fa-fw', NULL, '[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_id": 3}]}]');


--
-- TOC entry 2348 (class 0 OID 0)
-- Dependencies: 186
-- Name: member_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_type_id_seq', 1, true);


--
-- TOC entry 2318 (class 0 OID 98027)
-- Dependencies: 174
-- Data for Name: menu_link; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (1, '1', 'Content', NULL, 1, 'fa fa-newspaper-o fa-fw', NULL, 1, 'main', '{content_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (2, '2', 'Media', NULL, 2, 'fa fa-file-image-o fa-fw', NULL, 1, 'main', '{media_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (3, '3', 'Users', NULL, 3, 'fa fa-user fa-fw', NULL, 1, 'main', '{users_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (4, '4', 'Members', NULL, 4, 'fa fa-users fa-fw', NULL, 1, 'main', '{members_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (5, '5', 'Settings', NULL, 5, 'fa fa-gear fa-fw', NULL, 1, 'main', '{settings_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (6, '5.6', 'Content Types', 5, 10, 'fa fa-newspaper-o fa-fw', NULL, 1, 'main', '{content_types_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (7, '5.7', 'Media Types', 5, 11, 'fa fa-files-o fa-fw', NULL, 1, 'main', '{media_types_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (8, '5.8', 'Data Types', 5, 12, 'fa fa-check-square-o fa-fw', NULL, 1, 'main', '{data_types_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (9, '5.9', 'Templates', 5, 13, 'fa fa-eye fa-fw', NULL, 1, 'main', '{templates_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (10, '6.10', 'Scripts', 5, 14, 'fa fa-file-code-o fa-fw', NULL, 1, 'main', '{scripts_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (11, '6.11', 'Stylesheets', 5, 15, 'fa fa-desktop fa-fw', NULL, 1, 'main', '{stylesheets_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (12, '5.12', 'Member Types', 4, 30, 'fa fa-smile-o fa-fw', NULL, 1, 'main', '{member_types_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (13, '5.13', 'Member Groups', 4, 33, 'fa fa-smile-o fa-fw', NULL, 1, 'main', '{member_groups_section}');
INSERT INTO menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) VALUES (14, '3.14', 'User Groups', 3, 38, 'fa fa-smile-o fa-fw', NULL, 1, 'main', '{user_groups_section}');


--
-- TOC entry 2349 (class 0 OID 0)
-- Dependencies: 175
-- Name: menu_link_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('menu_link_id_seq', 14, true);


--
-- TOC entry 2320 (class 0 OID 98044)
-- Dependencies: 176
-- Data for Name: permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO permission (name) VALUES ('node_create');
INSERT INTO permission (name) VALUES ('node_delete');
INSERT INTO permission (name) VALUES ('node_update');
INSERT INTO permission (name) VALUES ('node_move');
INSERT INTO permission (name) VALUES ('node_copy');
INSERT INTO permission (name) VALUES ('node_public_access');
INSERT INTO permission (name) VALUES ('node_permissions');
INSERT INTO permission (name) VALUES ('node_send_to_publish');
INSERT INTO permission (name) VALUES ('node_publish');
INSERT INTO permission (name) VALUES ('node_browse');
INSERT INTO permission (name) VALUES ('node_change_content_type');
INSERT INTO permission (name) VALUES ('admin');
INSERT INTO permission (name) VALUES ('content_all');
INSERT INTO permission (name) VALUES ('content_create');
INSERT INTO permission (name) VALUES ('content_delete');
INSERT INTO permission (name) VALUES ('content_update');
INSERT INTO permission (name) VALUES ('content_section');
INSERT INTO permission (name) VALUES ('content_browse');
INSERT INTO permission (name) VALUES ('media_all');
INSERT INTO permission (name) VALUES ('media_create');
INSERT INTO permission (name) VALUES ('media_delete');
INSERT INTO permission (name) VALUES ('media_update');
INSERT INTO permission (name) VALUES ('media_section');
INSERT INTO permission (name) VALUES ('media_browse');
INSERT INTO permission (name) VALUES ('users_all');
INSERT INTO permission (name) VALUES ('users_create');
INSERT INTO permission (name) VALUES ('users_delete');
INSERT INTO permission (name) VALUES ('users_update');
INSERT INTO permission (name) VALUES ('users_section');
INSERT INTO permission (name) VALUES ('users_browse');
INSERT INTO permission (name) VALUES ('user_types_all');
INSERT INTO permission (name) VALUES ('user_types_create');
INSERT INTO permission (name) VALUES ('user_types_delete');
INSERT INTO permission (name) VALUES ('user_types_update');
INSERT INTO permission (name) VALUES ('user_types_section');
INSERT INTO permission (name) VALUES ('user_types_browse');
INSERT INTO permission (name) VALUES ('user_groups_all');
INSERT INTO permission (name) VALUES ('user_groups_create');
INSERT INTO permission (name) VALUES ('user_groups_delete');
INSERT INTO permission (name) VALUES ('user_groups_update');
INSERT INTO permission (name) VALUES ('user_groups_section');
INSERT INTO permission (name) VALUES ('user_groups_browse');
INSERT INTO permission (name) VALUES ('members_all');
INSERT INTO permission (name) VALUES ('members_create');
INSERT INTO permission (name) VALUES ('members_delete');
INSERT INTO permission (name) VALUES ('members_update');
INSERT INTO permission (name) VALUES ('members_section');
INSERT INTO permission (name) VALUES ('members_browse');
INSERT INTO permission (name) VALUES ('member_types_all');
INSERT INTO permission (name) VALUES ('member_types_create');
INSERT INTO permission (name) VALUES ('member_types_delete');
INSERT INTO permission (name) VALUES ('member_types_update');
INSERT INTO permission (name) VALUES ('member_types_section');
INSERT INTO permission (name) VALUES ('member_types_browse');
INSERT INTO permission (name) VALUES ('member_groups_all');
INSERT INTO permission (name) VALUES ('member_groups_create');
INSERT INTO permission (name) VALUES ('member_groups_delete');
INSERT INTO permission (name) VALUES ('member_groups_update');
INSERT INTO permission (name) VALUES ('member_groups_section');
INSERT INTO permission (name) VALUES ('member_groups_browse');
INSERT INTO permission (name) VALUES ('templates_all');
INSERT INTO permission (name) VALUES ('templates_create');
INSERT INTO permission (name) VALUES ('templates_delete');
INSERT INTO permission (name) VALUES ('templates_update');
INSERT INTO permission (name) VALUES ('templates_section');
INSERT INTO permission (name) VALUES ('templates_browse');
INSERT INTO permission (name) VALUES ('scripts_all');
INSERT INTO permission (name) VALUES ('scripts_create');
INSERT INTO permission (name) VALUES ('scripts_delete');
INSERT INTO permission (name) VALUES ('scripts_update');
INSERT INTO permission (name) VALUES ('scripts_section');
INSERT INTO permission (name) VALUES ('scripts_browse');
INSERT INTO permission (name) VALUES ('stylesheets_all');
INSERT INTO permission (name) VALUES ('stylesheets_create');
INSERT INTO permission (name) VALUES ('stylesheets_delete');
INSERT INTO permission (name) VALUES ('stylesheets_update');
INSERT INTO permission (name) VALUES ('stylesheets_section');
INSERT INTO permission (name) VALUES ('stylesheets_browse');
INSERT INTO permission (name) VALUES ('settings_section');
INSERT INTO permission (name) VALUES ('settings_all');
INSERT INTO permission (name) VALUES ('node_sort');
INSERT INTO permission (name) VALUES ('content_types_all');
INSERT INTO permission (name) VALUES ('content_types_create');
INSERT INTO permission (name) VALUES ('content_types_delete');
INSERT INTO permission (name) VALUES ('content_types_update');
INSERT INTO permission (name) VALUES ('content_types_section');
INSERT INTO permission (name) VALUES ('content_types_browse');
INSERT INTO permission (name) VALUES ('media_types_all');
INSERT INTO permission (name) VALUES ('media_types_create');
INSERT INTO permission (name) VALUES ('media_types_delete');
INSERT INTO permission (name) VALUES ('media_types_update');
INSERT INTO permission (name) VALUES ('media_types_section');
INSERT INTO permission (name) VALUES ('media_types_browse');
INSERT INTO permission (name) VALUES ('data_types_all');
INSERT INTO permission (name) VALUES ('data_types_create');
INSERT INTO permission (name) VALUES ('data_types_delete');
INSERT INTO permission (name) VALUES ('data_types_update');
INSERT INTO permission (name) VALUES ('data_types_section');
INSERT INTO permission (name) VALUES ('data_types_browse');


--
-- TOC entry 2321 (class 0 OID 98050)
-- Dependencies: 177
-- Data for Name: route; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (1, 'content', 'content', NULL, '/admin/content', '[{"single": "public/views/content/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (2, 'media', 'media', NULL, '/admin/media', '[{"single": "public/views/media/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (4, 'members', 'members', NULL, '/admin/members', '[{"single": "public/views/members/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (5, 'settings', 'settings', NULL, '/admin/settings', '[{"single": "public/views/settings/index.html"}]', true);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (10, 'settings.contentTypes', 'contentTypes', 5, '/content-type', '[{"single": "public/views/settings/content-type/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (11, 'settings.mediaTypes', 'mediaTypes', 5, '/media-type', '[{"single": "public/views/settings/media-type/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (12, 'settings.dataTypes', 'dataTypes', 5, '/data-type', '[{"single": "public/views/settings/data-type/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (13, 'settings.templates', 'templates', 5, '/template', '[{"single": "public/views/settings/template/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (14, 'settings.scripts', 'scripts', 5, '/script', '[{"single": "public/views/settings/script/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (15, 'settings.stylesheets', 'stylesheets', 5, '/stylesheet', '[{"single": "public/views/settings/stylesheet/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (16, 'settings.contentTypes.new', 'new', 10, '/new?type&parent', '[{"single": "public/views/settings/content-type/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (17, 'settings.mediaTypes.new', 'new', 11, '/new?type&parent', '[{"single": "public/views/settings/media-type/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (18, 'settings.dataTypes.new', 'new', 12, '/new', '[{"single": "public/views/settings/data-type/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (19, 'settings.templates.new', 'new', 13, '/new?parent', '[{"single": "public/views/settings/template/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (20, 'settings.scripts.new', 'new', 14, '/new?type&parent', '[{"single": "public/views/settings/script/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (21, 'settings.stylesheets.new', 'new', 15, '/new?type&parent', '[{"single": "public/views/settings/stylesheet/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (26, 'settings.scripts.edit', 'edit', 14, '/edit/:name', '[{"single": "public/views/settings/script/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (27, 'settings.stylesheets.edit', 'edit', 15, '/edit/:name', '[{"single": "public/views/settings/stylesheet/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (28, 'members.edit', 'edit', 4, '/edit/:id', '[{"single": "public/views/members/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (29, 'members.new', 'new', 4, '/new', '[{"single": "public/views/members/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (30, 'members.memberTypes', 'memberTypes', 4, '/member-type', '[{"single": "public/views/members/member-type/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (32, 'members.memberTypes.new', 'new', 30, '/new?type&parent', '[{"single": "public/views/members/member-type/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (33, 'members.memberGroups', 'memberGroups', 4, '/member-group', '[{"single": "public/views/members/member-group/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (35, 'members.memberGroups.new', 'new', 33, '/new', '[{"single": "public/views/members/member-group/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (24, 'settings.dataTypes.edit', 'edit', 12, '/edit/:id', '[{"single": "public/views/settings/data-type/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (34, 'members.memberGroups.edit', 'edit', 33, '/edit/:id', '[{"single": "public/views/members/member-group/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (31, 'members.memberTypes.edit', 'edit', 30, '/edit/:id', '[{"single": "public/views/members/member-type/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (22, 'settings.contentTypes.edit', 'edit', 10, '/edit/:id', '[{"single": "public/views/settings/content-type/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (23, 'settings.mediaTypes.edit', 'edit', 11, '/edit/:id', '[{"single": "public/views/settings/media-type/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (25, 'settings.templates.edit', 'edit', 13, '/edit/:id', '[{"single": "public/views/settings/template/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (7, 'content.edit', 'edit', 1, '/edit/:id', '[{"single": "public/views/content/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (9, 'media.edit', 'edit', 2, '/edit/:id', '[{"single": "public/views/media/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (6, 'content.new', 'new', 1, '/new?type_id&content_type_id&parent_id', '[{"single": "public/views/content/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (8, 'media.new', 'new', 2, '/new?type_id&content_type_id&parent_id', '[{"single": "public/views/media/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (38, 'users.userGroups', 'userGroups', 3, '/user-group', '[{"single": "core/modules/user/public/views/user-group/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (39, 'users.userGroups.edit', 'edit', 38, '/edit/:id', '[{"single": "core/modules/user/public/views/user-group/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (40, 'users.userGroups.new', 'new', 38, '/new', '[{"single": "core/modules/user/public/views/user-group/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (36, 'users.new', 'new', 3, '/new', '[{"single": "core/modules/user/public/views/user/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (37, 'users.edit', 'edit', 3, '/edit/:id', '[{"single": "core/modules/user/public/views/user/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (3, 'users', 'users', NULL, '/admin/users', '[{"single": "core/modules/user/public/views/user/index.html"}]', false);


--
-- TOC entry 2350 (class 0 OID 0)
-- Dependencies: 178
-- Name: route_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('route_id_seq', 40, true);


--
-- TOC entry 2334 (class 0 OID 98211)
-- Dependencies: 190
-- Data for Name: template; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (1, '1', NULL, 'Layout', 'layout', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (2, '1.2', 1, 'Home', 'home', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (3, '1.3', 1, 'Page', 'page', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (4, '1.4', 1, 'Post', 'post', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (5, '1.5', 1, 'Post Overview', 'post_overview', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (6, '1.6', 1, 'Category', 'category', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (7, '1.7', 1, 'Login', 'login', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (8, '1.8', 1, 'Register', 'register', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (9, '1.9', 1, '404', '404', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (10, '1.10', 1, 'Unauthorized', 'unauthorized', 1, '2015-03-27 03:46:27.018', false);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (11, '11', NULL, 'Top Navigation', 'top_navigation', 1, '2015-03-27 03:52:39.752', true);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (12, '12', NULL, 'Featured Pages Widget', 'featured_pages_widget', 1, '2015-03-27 03:52:39.752', true);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (13, '13', NULL, 'Recent Posts Widget', 'recent_posts_widget', 1, '2015-03-27 03:52:39.752', true);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (14, '14', NULL, 'Post Overview Widget', 'post_overview_widget', 1, '2015-03-27 03:52:39.752', true);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (15, '15', NULL, 'Category List Widget', 'category_list_widget', 1, '2015-03-27 03:52:39.752', true);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (16, '16', NULL, 'Social Widget', 'social_widget', 1, '2015-03-27 03:52:39.752', true);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (17, '17', NULL, 'About Widget', 'about_widget', 1, '2015-03-27 03:52:39.752', true);
INSERT INTO template (id, path, parent_id, name, alias, created_by, created_date, is_partial) VALUES (18, '18', NULL, 'Login Widget', 'login_widget', 1, '2015-03-27 03:52:39.752', true);


--
-- TOC entry 2351 (class 0 OID 0)
-- Dependencies: 189
-- Name: template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('template_id_seq', 18, true);


--
-- TOC entry 2323 (class 0 OID 98067)
-- Dependencies: 179
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids, permissions) VALUES (1, 'admin', 'Admin', 'Demo', '$2a$10$CWn3i3CKMJzhRGJ3B9TIeO.ePxgzajTFoB2cH5fpXkiZ7Az9jrmue', 'soren@codeish.com', '2014-11-15 16:51:00.215', NULL, '2015-04-14 19:09:38.582', NULL, 1, 'L6I2UYLEYZDENGOY6KKH774IONP4TDQYXVWJRYH23LC2YZZK3OJQ', '{1}', NULL);


--
-- TOC entry 2324 (class 0 OID 98074)
-- Dependencies: 180
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO user_group (id, name, alias, permissions) VALUES (1, 'Administrator', 'administrator', '{node_create,node_delete,node_update,node_move,node_copy,node_public_access,node_permissions,node_send_to_publish,node_publish,node_browse,node_change_content_type,admin,content_all,content_create,content_delete,content_update,content_section,content_browse,media_all,media_create,media_delete,media_update,media_section,media_browse,users_all,users_create,users_delete,users_update,users_section,users_browse,user_types_all,user_types_create,user_types_delete,user_types_update,user_types_section,user_types_browse,user_groups_all,user_groups_create,user_groups_delete,user_groups_update,user_groups_section,user_groups_browse,members_all,members_create,members_delete,members_update,members_section,members_browse,member_types_all,member_types_create,member_types_delete,member_types_update,member_types_section,member_types_browse,member_groups_all,member_groups_create,member_groups_delete,member_groups_update,member_groups_section,member_groups_browse,templates_all,templates_create,templates_delete,templates_update,templates_section,templates_browse,scripts_all,scripts_create,scripts_delete,scripts_update,scripts_section,scripts_browse,stylesheets_all,stylesheets_create,stylesheets_delete,stylesheets_update,stylesheets_section,stylesheets_browse,settings_section,settings_all,node_sort,content_types_all,content_types_create,content_types_delete,content_types_update,content_types_section,content_types_browse,media_types_all,media_types_create,media_types_delete,media_types_update,media_types_section,media_types_browse,data_types_all,data_types_create,data_types_delete,data_types_update,data_types_section,data_types_browse}');
INSERT INTO user_group (id, name, alias, permissions) VALUES (2, 'Editor', 'editor', '{}');
INSERT INTO user_group (id, name, alias, permissions) VALUES (3, 'Writer', 'writer', '{}');


--
-- TOC entry 2352 (class 0 OID 0)
-- Dependencies: 181
-- Name: user_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_group_id_seq', 3, true);


--
-- TOC entry 2353 (class 0 OID 0)
-- Dependencies: 182
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 1, true);


-- Completed on 2015-04-14 19:24:30

--
-- PostgreSQL database dump complete
--

