--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-03-02 00:34:30

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

--
-- TOC entry 2339 (class 0 OID 82782)
-- Dependencies: 172
-- Data for Name: content; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content VALUES (13, 33, 5, '{"prop2": "prop2", "prop3": "prop3", "page_title": "About me", "page_content": "About page description", "template_node_id": 22}', NULL);
INSERT INTO content VALUES (12, 32, 5, '{"prop2": "p2", "prop3": "p3", "page_title": "Yet another test page title override", "page_content": "Yet another test page description goes here", "template_node_id": 8}', NULL);
INSERT INTO content VALUES (14, 23, 16, NULL, NULL);
INSERT INTO content VALUES (7, 13, 5, '{"prop2": "11", "prop3": "Another page prop 31", "page_title": "Another page title1", "page_content": "Another page content goes here1", "template_node_id": 25}', NULL);
INSERT INTO content VALUES (19, 40, 5, '{"prop2": "p2", "prop3": "p3", "page_title": "test page 2 child page title override", "page_content": "test page 2 child desc", "template_node_id": 25}', NULL);
INSERT INTO content VALUES (4, 19, 15, '{"alt": "Postgresql image alt text", "url": "/media/2014/10/postgresql.png", "path": "media\\postgresql.png", "caption": "This is the caption of the postgresql image", "description": "Postgresql image description"}', NULL);
INSERT INTO content VALUES (5, 11, 5, '{"prop3": "sample child page level 1 page prop 3", "page_title": "Child page level 1 title", "page_content": "Sample page - child page level 1 content goes here", "template_node_id": 22}', NULL);
INSERT INTO content VALUES (15, 36, 16, '{"path": "media\\Another Image Folder2"}', NULL);
INSERT INTO content VALUES (8, 24, 15, '{"alt": "Goku SSJ3 image alt text", "url": "/media/Sample picture folder/Goku_SSJ3.jpg", "path": "media\\Sample picture folder\\Goku_SSJ3.jpg", "caption": "This is the caption of the Goku SSJ3 image", "description": "Goku SSJ3 image description"}', NULL);
INSERT INTO content VALUES (20, 41, 15, '{"alt": "gopher alt", "url": "/media/Another Image Folder1/gopher.jpg", "path": "media\\Another Image Folder2\\gopher.jpg", "title": "gopher title", "caption": "gopher caption", "description": "gopher description"}', NULL);
INSERT INTO content VALUES (11, 31, 5, '{"prop3": "msp3", "page_title": "My Sub Page Override", "page_content": "mysubpage desc", "template_node_id": 8}', NULL);
INSERT INTO content VALUES (29, 54, 15, '{"alt": "catduck.jpg", "path": "media\\Subfolder depth test\\Level 1\\catduck.jpg", "title": "catduck.jpg", "caption": "catduck.jpg", "description": "catduck1.jpg"}', NULL);
INSERT INTO content VALUES (3, 18, 15, '{"alt": "Gopher image alt text1", "url": "/media/2014/10/gopher.jpg", "path": "media\\gopher.jpg", "caption": "This is the caption of the gopher image1", "description": "Gopher image description1", "temporary property": "lol"}', NULL);
INSERT INTO content VALUES (9, 29, 5, '{"prop2": "prop2", "prop3": "prop3", "page_title": "test page title override1", "page_content": "This is just a test page", "template_node_id": 8}', NULL);
INSERT INTO content VALUES (6, 12, 5, '{"prop3": "sample child page level 2 page prop 3", "page_title": "Child page level 2 title", "page_content": "Sample page - child page level 2 content goes here1", "template_node_id": 22}', '{"groups": [1], "members": [1]}');
INSERT INTO content VALUES (2, 10, 5, '{"prop2": "prop2a", "prop3": "sample page prop 3", "page_title": "Sample page title", "page_content": "Sample page content goes here", "template_node_id": 22}', NULL);
INSERT INTO content VALUES (16, 37, 16, '{"path": "media\\Subfolder depth test"}', NULL);
INSERT INTO content VALUES (17, 38, 16, '{"path": "media\\2014"}', NULL);
INSERT INTO content VALUES (21, 45, 16, '{"path": "media\\2014\\12"}', NULL);
INSERT INTO content VALUES (18, 39, 5, '{"prop2": "tp2p2", "prop3": "tp2p3", "page_title": "test page 2 title override1", "page_content": "test page 2 content", "template_node_id": 22}', NULL);
INSERT INTO content VALUES (1, 9, 4, '{"prop2": "Home page prop 2", "domains": ["localhost:8080", "localhost:8080/test"], "facebook": "facebook.com/home", "copyright": "&copy; 2014 codeish.com", "site_name": "Collexy cms test site", "page_title": "Home page title", "site_tagline": "Test site tagline", "template_node_id": 7}', NULL);
INSERT INTO content VALUES (10, 30, 5, '{"prop2": "", "prop3": "mypageprop3", "page_title": "Login page test", "page_content": "This is a login page for members", "template_node_id": 25}', NULL);
INSERT INTO content VALUES (25, 50, 15, '{"alt": "tiny.jpg", "path": "media\\Subfolder depth test\\Level 1\\Level 2\\tiny.jpg", "title": "tiny.jpg", "caption": "tiny.jpg", "description": "tiny1.jpg"}', NULL);
INSERT INTO content VALUES (28, 53, 15, '{"alt": "AngularLogo alt", "path": "media\\Subfolder depth test\\AnguarLogo.png", "title": "AngularLogo.png", "caption": "AngularLogo caption", "description": "AngularLogo desc"}', NULL);
INSERT INTO content VALUES (27, 52, 15, '{"alt": "taco-hamster.jpg", "path": "media\\Subfolder depth test\\Level 1\\Level 2\\taco-hamster.jpg", "title": "taco-hamster.jpg", "caption": "taco-hamster.jpg", "description": "taco-hamster.jpg"}', NULL);
INSERT INTO content VALUES (24, 49, 15, '{"alt": "blomkals-hamster.jpg", "path": "media\\Subfolder depth test\\Level 1\\Level 2\\blomkals-hamster.jpg", "title": "blomkals-hamster.jpg", "caption": "blomkals-hamster.jpg", "description": "blomkals-hamster.jpg"}', NULL);
INSERT INTO content VALUES (31, 56, 15, '{"alt": "ducks.jpg", "path": "media\\2014\\12\\ducks.jpg", "title": "ducks.jpg", "caption": "ducks.jpg", "description": "ducks.jpg3"}', NULL);
INSERT INTO content VALUES (32, 57, 15, '{"alt": "sleeping-kitten.jpg", "path": "media\\2014\\12\\sleeping-kitten.jpg", "title": "sleeping-kitten.jpg", "caption": "sleeping-kitten.jpg", "description": "sleeping-kitten.jpg"}', NULL);
INSERT INTO content VALUES (30, 55, 15, '{"alt": "cat-prays.jpg", "path": "media\\2014\\12\\cat-prays.jpg", "title": "cat-prays.jpg", "caption": "cat-prays.jpg", "description": "cat-prays.jpg"}', NULL);
INSERT INTO content VALUES (22, 46, 16, '{"path": "media\\Subfolder depth test\\Level 1"}', NULL);
INSERT INTO content VALUES (23, 47, 16, '{"path": "media\\Subfolder depth test\\Level 1\\Level 2"}', NULL);
INSERT INTO content VALUES (26, 51, 15, '{"alt": "dog.jpg", "path": "media\\Subfolder depth test\\Level 1\\Level 2\\dog.jpg", "title": "dog.jpg", "caption": "dog.jpg", "description": "dog.jpg"}', NULL);


--
-- TOC entry 2370 (class 0 OID 0)
-- Dependencies: 173
-- Name: content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_id_seq', 33, true);


--
-- TOC entry 2341 (class 0 OID 82790)
-- Dependencies: 174
-- Data for Name: content_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content_type VALUES (4, 15, 'Image', 'Image content type description', 'fa fa-folder-o', 'fa fa-folder-o', NULL, 'null', '[{"name":"Image","properties":[{"name":"path","order":1,"data_type_node_id":2,"help_text":"help text","description":"URL goes here."},{"name":"title","order":2,"data_type_node_id":2,"help_text":"help text","description":"The title entered here can override the above one."},{"name":"caption","order":3,"data_type_node_id":14,"help_text":"help text","description":"Caption goes here."},{"name":"alt","order":4,"data_type_node_id":14,"help_text":"help text","description":"Alt goes here."},{"name":"description","order":5,"data_type_node_id":14,"help_text":"help text","description":"Description goes here."},{"name":"file_upload","order":1,"data_type_node_id":48,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties","properties":[{"name":"temporary property","order":1,"data_type_node_id":2,"help_text":"help text","description":"Temporary description goes here."}]}]');
INSERT INTO content_type VALUES (1, 3, 'Master', 'Some description', 'fa fa-folder-o', 'fa fa-folder-o', NULL, NULL, '[{"name": "Content", "properties": [{"name": "page_title", "order": 1, "data_type_node_id": 2, "help_text": "help text", "description": "The page title overrides the name the page has been given."}]}, {"name": "Properties", "properties": [{"name": "prop2", "order": 1, "data_type_node_id": 2, "help_text": "help text2", "description": "description2"}, {"name": "prop3", "order": 2, "data_type_node_id": 2, "help_text": "help text3", "description": "description3"}]}]');
INSERT INTO content_type VALUES (5, 28, 'ctTestContentTypeAlias', 'ct-test desc', 'ct-test-icon', 'ct-test-thumbnail', 3, '{"template_node_id": 8, "allowed_templates_node_id": [8, 22, 25], "allowed_content_types_node_id": [5]}', '[{"name":"Mytab1","properties":[{"name":"property name1","order":1,"data_type_node_id":2,"help_text":"prop help text1","description":"prop description1"},{"name":"property name2","order":2,"data_type_node_id":14,"help_text":"prop help text2","description":"prop description2"}]},{"name":"Mytab2","properties":[{"name":"property name3","order":1,"data_type_node_id":27,"help_text":"prop help text3","description":"prop description3"},{"name":"property name4","order":2,"data_type_node_id":14,"help_text":"prop help text4","description":"prop description4"}]}]');
INSERT INTO content_type VALUES (3, 5, 'ctPage', 'Page content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 3, '{"template_node_id": 8, "allowed_templates_node_id": [8, 22, 25], "allowed_content_types_node_id": [5]}', '[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":14,"help_text":"Help text for page contentent","description":"Page content description"}]}]');
INSERT INTO content_type VALUES (7, 35, 'mtTestMediaType alias', 'mtTest desc', 'mtTest-icon', 'mtTest-thumbnail', 16, '{"allowed_content_types_node_id": [15, 17, 16]}', '[{"name":"mytab"}]');
INSERT INTO content_type VALUES (8, 43, 'CT test alias', 'ct-test-desc', 'ct-test-icon', 'ct-test-thumb', 3, '{"template_node_id": "8", "allowed_templates_node_id": [22, 8, 25, 42], "allowed_content_types_node_id": [5, 28]}', '[{"name":"mytab"}]');
INSERT INTO content_type VALUES (9, 44, 'Test Content Type 2 alias', 'tc2-desc', 'tc2-icon', 'tc2-thumb', 3, '{"template_node_id": "8", "allowed_templates_node_id": [22, 8, 25, 26, 42], "allowed_content_types_node_id": [28, 43, 5]}', '[{"name":"mytab"}]');
INSERT INTO content_type VALUES (6, 16, 'mtFolder', 'Folder media type description1', 'mt-icon1', 'mt-thumbnail1', NULL, '{"allowed_content_types_node_id": [16, 15]}', '[{"name":"Folder","properties":[{"name":"folder_browser","order":1,"data_type_node_id":34,"help_text":"prop help text","description":"prop description"},{"name":"path","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties"}]');
INSERT INTO content_type VALUES (2, 4, 'ctHome', 'Home Some description', 'fa fa-folder-o', 'fa fa-folder-o', 3, '{"template_node_id": 7, "allowed_templates_node_id": [7], "allowed_content_types_node_id": [5, 28]}', '[{"name":"Content","properties":[{"name":"site_name","order":2,"data_type_node_id":2,"help_text":"help text","description":"Site name goes here."},{"name":"site_tagline","order":3,"data_type_node_id":2,"help_text":"help text","description":"Site tagline goes here."},{"name":"copyright","order":4,"data_type_node_id":2,"help_text":"help text","description":"Copyright here."},{"name":"domains","order":5,"data_type_node_id":59,"help_text":"help text","description":"Domains goes here."}]},{"name":"Social","properties":[{"name":"facebook","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your facebook link here."},{"name":"twitter","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your twitter link here."},{"name":"linkedin","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your linkedin link here."}]}]');


--
-- TOC entry 2371 (class 0 OID 0)
-- Dependencies: 175
-- Name: content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_id_seq', 9, true);


--
-- TOC entry 2343 (class 0 OID 82798)
-- Dependencies: 176
-- Data for Name: data_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO data_type VALUES (6, 59, '<div>
	<input type="text"/> <button type="button">Add domain</button><br>
	<ul>
		<li ng-repeat="domain in data.meta[prop.name]">{{domain}}</li>
	</ul>
	<button type="button">Delete selected</button>
</div>', 'defDomains');
INSERT INTO data_type VALUES (2, 14, '<textarea id="{{prop.name}}" ng-model="data.meta[prop.name]">', 'defTextarea');
INSERT INTO data_type VALUES (1, 2, '<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">', 'defTextInput');
INSERT INTO data_type VALUES (3, 27, '<colorpicker>The color picker data type is not implemented yet!</colorpicker>', 'dtColorPicker');
INSERT INTO data_type VALUES (4, 34, '<folderbrowser>This is an awesome folder browser (unimplemented datatype)</folderbrowser>', 'dtFolderBrowser');
INSERT INTO data_type VALUES (5, 48, '<input type="file" file-input="test.files" multiple />
<button ng-click="upload()" type="button">Upload</button>
<li ng-repeat="file in test.files">{{file.name}}</li>
<!--<input type="file" onchange="angular.element(this).scope().filesChanged(this)" multiple />
<button ng-click="upload()">Upload</button>
<li ng-repeat="file in files">{{file.name}}</li>-->', 'dtFileUpload');


--
-- TOC entry 2372 (class 0 OID 0)
-- Dependencies: 177
-- Name: data_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_type_id_seq', 6, true);


--
-- TOC entry 2345 (class 0 OID 82806)
-- Dependencies: 178
-- Data for Name: domain; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO domain VALUES (1, 9, 'localhost:8080');


--
-- TOC entry 2373 (class 0 OID 0)
-- Dependencies: 179
-- Name: domain_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('domain_id_seq', 1, true);


--
-- TOC entry 2347 (class 0 OID 82814)
-- Dependencies: 180
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member VALUES (1, 'default_member', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'default_member@mail.com', '{"comments": "default user comments"}', '2015-01-22 14:25:38.904', NULL, '2015-02-19 23:46:00.495', NULL, 1, 'GIWES3RHMY5RKC7OZPOQTF5FQFWX32D5VLV3CAKT4HGKP5LZIENA', 61, '{1}');


--
-- TOC entry 2348 (class 0 OID 82822)
-- Dependencies: 181
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_group VALUES (1, 'authenticated_member', 'All logged in members');


--
-- TOC entry 2374 (class 0 OID 0)
-- Dependencies: 182
-- Name: member_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_group_id_seq', 1, true);


--
-- TOC entry 2375 (class 0 OID 0)
-- Dependencies: 183
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_id_seq', 1, true);


--
-- TOC entry 2351 (class 0 OID 82832)
-- Dependencies: 184
-- Data for Name: member_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_type VALUES (1, 61, 'mtMember', 'Default member type', 'fa fa-user fa-fw', 1, NULL, '[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_node_id": 14}]}]');


--
-- TOC entry 2376 (class 0 OID 0)
-- Dependencies: 185
-- Name: member_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_type_id_seq', 1, true);


--
-- TOC entry 2353 (class 0 OID 82840)
-- Dependencies: 186
-- Data for Name: menu_link; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO menu_link VALUES (3, '3', 'Media', NULL, 3, 'fa fa-file-image-o fa-fw', NULL, 1, 'main', '{media_section}');
INSERT INTO menu_link VALUES (4, '4', 'Users', NULL, 4, 'fa fa-user fa-fw', NULL, 1, 'main', '{users_section}');
INSERT INTO menu_link VALUES (5, '5', 'Members', NULL, 5, 'fa fa-users fa-fw', NULL, 1, 'main', '{members_section}');
INSERT INTO menu_link VALUES (6, '6', 'Settings', NULL, 6, 'fa fa-gear fa-fw', NULL, 1, 'main', '{settings_section}');
INSERT INTO menu_link VALUES (7, '6.7', 'Content Types', 6, 11, 'fa fa-newspaper-o fa-fw', NULL, 1, 'main', '{content_types_section}');
INSERT INTO menu_link VALUES (8, '6.8', 'Media Types', 6, 12, 'fa fa-files-o fa-fw', NULL, 1, 'main', '{media_types_section}');
INSERT INTO menu_link VALUES (9, '6.9', 'Data Types', 6, 13, 'fa fa-check-square-o fa-fw', NULL, 1, 'main', '{data_types_section}');
INSERT INTO menu_link VALUES (10, '6.10', 'Templates', 6, 14, 'fa fa-eye fa-fw', NULL, 1, 'main', '{templates_section}');
INSERT INTO menu_link VALUES (11, '6.11', 'Scripts', 6, 15, 'fa fa-file-code-o fa-fw', NULL, 1, 'main', '{scripts_section}');
INSERT INTO menu_link VALUES (12, '6.12', 'Stylesheets', 6, 16, 'fa fa-desktop fa-fw', NULL, 1, 'main', '{stylesheets_section}');
INSERT INTO menu_link VALUES (13, '5.13', 'Member Types', 5, 31, 'fa fa-smile-o fa-fw', NULL, 1, 'main', '{member_types_section}');
INSERT INTO menu_link VALUES (2, '2', 'Content', NULL, 2, 'fa fa-newspaper-o fa-fw', NULL, 1, 'main', '{content_section}');


--
-- TOC entry 2377 (class 0 OID 0)
-- Dependencies: 187
-- Name: menu_link_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('menu_link_id_seq', 13, true);


--
-- TOC entry 2355 (class 0 OID 82848)
-- Dependencies: 188
-- Data for Name: node; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO node VALUES (14, '1.14', 'Textarea', 11, 1, '2014-10-27 02:40:41.179', 1, NULL, NULL);
INSERT INTO node VALUES (48, '1.48', 'File upload', 11, 1, '2014-12-05 19:56:17.883', NULL, NULL, NULL);
INSERT INTO node VALUES (10, '1.9.10', 'Sample Page', 1, 1, '2014-10-22 16:51:00.215', 9, '[{"id": 2, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]', NULL);
INSERT INTO node VALUES (60, '1.6.60', '404', 3, 1, '2015-01-20 13:46:33.668', 6, NULL, NULL);
INSERT INTO node VALUES (20, '1.20', 'Sidebar 1', 3, 1, '2014-11-10 09:03:20.514', 1, NULL, NULL);
INSERT INTO node VALUES (15, '1.15', 'Image', 7, 1, '2014-10-28 15:16:25.972', 1, NULL, NULL);
INSERT INTO node VALUES (36, '1.36', 'Another Image Folder2', 2, 1, '2014-12-02 01:00:51.206', 1, NULL, NULL);
INSERT INTO node VALUES (21, '1.21', 'Sidebar 2', 3, 1, '2014-11-10 23:56:55.038', 1, NULL, NULL);
INSERT INTO node VALUES (1, '1', 'root', 5, 1, '2014-10-22 16:51:00.215', NULL, NULL, NULL);
INSERT INTO node VALUES (2, '1.2', 'Text input', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (3, '1.3', 'Master', 4, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (13, '1.9.13', 'Another Page', 1, 1, '2014-10-26 23:27:14.571', 9, NULL, NULL);
INSERT INTO node VALUES (9, '1.9', 'Home', 1, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (61, '1.61', 'Member', 12, 1, '2015-01-22 15:55:13.957', 1, NULL, NULL);
INSERT INTO node VALUES (28, '1.3.28', 'Test Content Type', 4, 1, '2014-11-26 04:20:48.026', 3, NULL, NULL);
INSERT INTO node VALUES (26, '1.6.25.26', 'Child of test template', 3, 1, '2014-11-26 01:39:42.816', 25, NULL, NULL);
INSERT INTO node VALUES (46, '1.37.46', 'Level 1', 2, 1, '2014-12-05 17:02:13.875', 37, NULL, NULL);
INSERT INTO node VALUES (41, '1.36.41', 'gopher.jpg', 2, 1, '2014-12-02 02:08:26.737', 36, NULL, NULL);
INSERT INTO node VALUES (17, '1.17', 'File', 7, 1, '2014-10-28 15:18:13.4', 1, NULL, NULL);
INSERT INTO node VALUES (18, '1.18', 'gopher.jpg', 2, 1, '2014-10-28 15:50:47.303', 1, NULL, NULL);
INSERT INTO node VALUES (16, '1.16', 'Folder', 7, 1, '2014-10-28 15:18:13.4', 1, NULL, NULL);
INSERT INTO node VALUES (35, '1.35', 'Media Type Test', 7, 1, '2014-12-01 22:09:43.783', 1, NULL, NULL);
INSERT INTO node VALUES (23, '1.23', 'Sample picture folder', 2, 1, '2014-11-17 16:57:14.654', 1, NULL, NULL);
INSERT INTO node VALUES (47, '1.37.46.47', 'Level 2', 2, 1, '2014-12-05 17:02:46.762', 46, NULL, NULL);
INSERT INTO node VALUES (43, '1.3.43', 'Content Type Test', 4, 1, '2014-12-02 12:38:59.527', 3, NULL, NULL);
INSERT INTO node VALUES (55, '1.38.45.55', 'cat-prays.jpg', 2, 1, '2014-12-06 13:07:08.943', 45, NULL, NULL);
INSERT INTO node VALUES (54, '1.37.46.54', 'catduck.jpg', 2, 1, '2014-12-06 03:44:40.07', 46, NULL, NULL);
INSERT INTO node VALUES (44, '1.3.44', 'Test Content Type 2', 4, 1, '2014-12-02 12:48:25.307', 3, NULL, NULL);
INSERT INTO node VALUES (58, '1.6.58', 'Unauthorized', 3, 1, '2014-12-15 14:24:22.063', 6, NULL, NULL);
INSERT INTO node VALUES (34, '1.34', 'Folder Browser', 11, 1, '2014-12-01 16:09:46.488', 1, NULL, NULL);
INSERT INTO node VALUES (19, '1.19', 'postgresql.png', 2, 1, '2014-10-28 17:53:37.488', 1, NULL, NULL);
INSERT INTO node VALUES (27, '1.27', 'Color Picker', 11, 1, '2014-11-26 02:20:17.638', 1, NULL, NULL);
INSERT INTO node VALUES (33, '1.9.33', 'About', 1, 1, '2014-12-01 12:11:25.838', 9, NULL, NULL);
INSERT INTO node VALUES (24, '1.23.24', 'Goku_SSJ3.jpg', 2, 1, '2014-11-17 16:58:57.285', 23, NULL, NULL);
INSERT INTO node VALUES (11, '1.9.10.11', 'Child Page Level 1', 1, 1, '2014-10-26 23:19:44.735', 10, NULL, '[{"id": 1, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]');
INSERT INTO node VALUES (39, '1.9.39', 'Testpage 2', 1, 1, '2014-12-02 01:43:33.233', 9, NULL, NULL);
INSERT INTO node VALUES (4, '1.3.4', 'Home', 4, 1, '2014-10-22 16:51:00.215', 3, NULL, NULL);
INSERT INTO node VALUES (31, '1.9.30.31', 'MySubPage', 1, 1, '2014-12-01 12:02:54.252', 30, NULL, NULL);
INSERT INTO node VALUES (32, '1.9.32', 'Yet another test page', 1, 1, '2014-12-01 12:07:29.999', 9, NULL, NULL);
INSERT INTO node VALUES (40, '1.9.39.40', 'Test page 2 child', 1, 1, '2014-12-02 01:45:49.78', 39, NULL, NULL);
INSERT INTO node VALUES (5, '1.3.5', 'Page', 4, 1, '2014-10-22 16:51:00.215', 3, NULL, NULL);
INSERT INTO node VALUES (51, '1.37.46.47.51', 'dog.jpg', 2, 1, '2014-12-05 21:06:49.532', 47, NULL, NULL);
INSERT INTO node VALUES (52, '1.37.46.47.52', 'taco-hamster.jpg', 2, 1, '2014-12-05 21:22:45.227', 47, NULL, NULL);
INSERT INTO node VALUES (22, '1.6.22', 'Page with sidebars', 3, 1, '2014-11-11 03:39:55.766', 6, NULL, NULL);
INSERT INTO node VALUES (49, '1.37.46.47.49', 'blomkals-hamster.jpg', 2, 1, '2014-12-05 20:44:25.921', 47, NULL, NULL);
INSERT INTO node VALUES (38, '1.38', '2014', 2, 1, '2014-12-02 01:42:09.979', 1, NULL, NULL);
INSERT INTO node VALUES (50, '1.37.46.47.50', 'tiny.jpg', 2, 1, '2014-12-05 21:05:42.816', 47, NULL, NULL);
INSERT INTO node VALUES (57, '1.38.45.57', 'sleeping-kitten.jpg', 2, 1, '2014-12-06 14:28:52.117', 45, NULL, NULL);
INSERT INTO node VALUES (37, '1.37', 'Subfolder depth test3', 2, 1, '2014-12-02 01:37:09.125', 1, NULL, NULL);
INSERT INTO node VALUES (8, '1.6.8', 'Page', 3, 1, '2014-10-22 16:51:00.215', 6, NULL, NULL);
INSERT INTO node VALUES (45, '1.38.45', '12', 2, 1, '2014-12-05 16:18:29.762', 38, NULL, NULL);
INSERT INTO node VALUES (6, '1.6', 'Layout', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (56, '1.38.45.56', 'ducks.jpg', 2, 1, '2014-12-06 13:10:14.637', 45, NULL, NULL);
INSERT INTO node VALUES (53, '1.37.53', 'AngularLogo.png', 2, 1, '2014-12-06 03:36:14.425', 37, NULL, NULL);
INSERT INTO node VALUES (42, '1.6.42', 'Test template 2', 3, 1, '2014-12-02 02:19:29.241', 6, NULL, NULL);
INSERT INTO node VALUES (7, '1.6.7', 'Home', 3, 1, '2014-10-22 16:51:00.215', 6, NULL, NULL);
INSERT INTO node VALUES (29, '1.9.29', 'Test Page', 1, 1, '2014-12-01 11:45:16.186', 9, NULL, NULL);
INSERT INTO node VALUES (30, '1.9.30', 'Login', 1, 1, '2014-12-01 11:54:10.208', 9, NULL, NULL);
INSERT INTO node VALUES (59, '1.59', 'Domains', 11, 1, '2015-01-19 21:22:06.945', NULL, NULL, NULL);
INSERT INTO node VALUES (12, '1.9.10.11.12', 'Child Page Level 2', 1, 1, '2014-10-26 23:19:44.735', 11, NULL, NULL);
INSERT INTO node VALUES (25, '1.6.25', 'Login', 3, 1, '2014-11-26 00:13:45.309', 6, NULL, NULL);


--
-- TOC entry 2378 (class 0 OID 0)
-- Dependencies: 189
-- Name: node_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('node_id_seq', 62, true);


--
-- TOC entry 2357 (class 0 OID 82857)
-- Dependencies: 190
-- Data for Name: permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO permission VALUES ('node_create');
INSERT INTO permission VALUES ('node_delete');
INSERT INTO permission VALUES ('node_update');
INSERT INTO permission VALUES ('node_move');
INSERT INTO permission VALUES ('node_copy');
INSERT INTO permission VALUES ('node_public_access');
INSERT INTO permission VALUES ('node_permissions');
INSERT INTO permission VALUES ('node_send_to_publish');
INSERT INTO permission VALUES ('node_publish');
INSERT INTO permission VALUES ('node_browse');
INSERT INTO permission VALUES ('node_change_content_type');
INSERT INTO permission VALUES ('admin');
INSERT INTO permission VALUES ('content_all');
INSERT INTO permission VALUES ('content_create');
INSERT INTO permission VALUES ('content_delete');
INSERT INTO permission VALUES ('content_update');
INSERT INTO permission VALUES ('content_section');
INSERT INTO permission VALUES ('content_browse');
INSERT INTO permission VALUES ('media_all');
INSERT INTO permission VALUES ('media_create');
INSERT INTO permission VALUES ('media_delete');
INSERT INTO permission VALUES ('media_update');
INSERT INTO permission VALUES ('media_section');
INSERT INTO permission VALUES ('media_browse');
INSERT INTO permission VALUES ('users_all');
INSERT INTO permission VALUES ('users_create');
INSERT INTO permission VALUES ('users_delete');
INSERT INTO permission VALUES ('users_update');
INSERT INTO permission VALUES ('users_section');
INSERT INTO permission VALUES ('users_browse');
INSERT INTO permission VALUES ('user_types_all');
INSERT INTO permission VALUES ('user_types_create');
INSERT INTO permission VALUES ('user_types_delete');
INSERT INTO permission VALUES ('user_types_update');
INSERT INTO permission VALUES ('user_types_section');
INSERT INTO permission VALUES ('user_types_browse');
INSERT INTO permission VALUES ('user_groups_all');
INSERT INTO permission VALUES ('user_groups_create');
INSERT INTO permission VALUES ('user_groups_delete');
INSERT INTO permission VALUES ('user_groups_update');
INSERT INTO permission VALUES ('user_groups_section');
INSERT INTO permission VALUES ('user_groups_browse');
INSERT INTO permission VALUES ('members_all');
INSERT INTO permission VALUES ('members_create');
INSERT INTO permission VALUES ('members_delete');
INSERT INTO permission VALUES ('members_update');
INSERT INTO permission VALUES ('members_section');
INSERT INTO permission VALUES ('members_browse');
INSERT INTO permission VALUES ('member_types_all');
INSERT INTO permission VALUES ('member_types_create');
INSERT INTO permission VALUES ('member_types_delete');
INSERT INTO permission VALUES ('member_types_update');
INSERT INTO permission VALUES ('member_types_section');
INSERT INTO permission VALUES ('member_types_browse');
INSERT INTO permission VALUES ('member_groups_all');
INSERT INTO permission VALUES ('member_groups_create');
INSERT INTO permission VALUES ('member_groups_delete');
INSERT INTO permission VALUES ('member_groups_update');
INSERT INTO permission VALUES ('member_groups_section');
INSERT INTO permission VALUES ('member_groups_browse');
INSERT INTO permission VALUES ('templates_all');
INSERT INTO permission VALUES ('templates_create');
INSERT INTO permission VALUES ('templates_delete');
INSERT INTO permission VALUES ('templates_update');
INSERT INTO permission VALUES ('templates_section');
INSERT INTO permission VALUES ('templates_browse');
INSERT INTO permission VALUES ('scripts_all');
INSERT INTO permission VALUES ('scripts_create');
INSERT INTO permission VALUES ('scripts_delete');
INSERT INTO permission VALUES ('scripts_update');
INSERT INTO permission VALUES ('scripts_section');
INSERT INTO permission VALUES ('scripts_browse');
INSERT INTO permission VALUES ('stylesheets_all');
INSERT INTO permission VALUES ('stylesheets_create');
INSERT INTO permission VALUES ('stylesheets_delete');
INSERT INTO permission VALUES ('stylesheets_update');
INSERT INTO permission VALUES ('stylesheets_section');
INSERT INTO permission VALUES ('stylesheets_browse');
INSERT INTO permission VALUES ('settings_section');
INSERT INTO permission VALUES ('settings_all');
INSERT INTO permission VALUES ('node_sort');
INSERT INTO permission VALUES ('content_types_all');
INSERT INTO permission VALUES ('content_types_create');
INSERT INTO permission VALUES ('content_types_delete');
INSERT INTO permission VALUES ('content_types_update');
INSERT INTO permission VALUES ('content_types_section');
INSERT INTO permission VALUES ('content_types_browse');
INSERT INTO permission VALUES ('media_types_all');
INSERT INTO permission VALUES ('media_types_create');
INSERT INTO permission VALUES ('media_types_delete');
INSERT INTO permission VALUES ('media_types_update');
INSERT INTO permission VALUES ('media_types_section');
INSERT INTO permission VALUES ('media_types_browse');
INSERT INTO permission VALUES ('data_types_all');
INSERT INTO permission VALUES ('data_types_create');
INSERT INTO permission VALUES ('data_types_delete');
INSERT INTO permission VALUES ('data_types_update');
INSERT INTO permission VALUES ('data_types_section');
INSERT INTO permission VALUES ('data_types_browse');


--
-- TOC entry 2358 (class 0 OID 82863)
-- Dependencies: 191
-- Data for Name: route; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO route VALUES (7, 'content.new', 'new', 2, '/new?node_type&content_type_node_id&parent_id', '[{"single": "public/views/content/new.html"}]', false);
INSERT INTO route VALUES (8, 'content.edit', 'edit', 2, '/edit/:nodeId', '[{"single": "public/views/content/edit.html"}]', false);
INSERT INTO route VALUES (9, 'media.new', 'new', 3, '/new?node_type&content_type_node_id&parent_id', '[{"single": "public/views/media/new.html"}]', false);
INSERT INTO route VALUES (10, 'media.edit', 'edit', 3, '/edit/:nodeId', '[{"single": "public/views/media/edit.html"}]', false);
INSERT INTO route VALUES (2, 'content', 'content', NULL, '/admin/content', '[{"single": "public/views/content/index.html"}]', false);
INSERT INTO route VALUES (3, 'media', 'media', NULL, '/admin/media', '[{"single": "public/views/media/index.html"}]', false);
INSERT INTO route VALUES (4, 'users', 'users', NULL, '/admin/users', '[{"single": "public/views/users/index.html"}]', false);
INSERT INTO route VALUES (5, 'members', 'members', NULL, '/admin/members', '[{"single": "public/views/members/index.html"}]', false);
INSERT INTO route VALUES (6, 'settings', 'settings', NULL, '/admin/settings', '[{"single": "public/views/settings/index.html"}]', true);
INSERT INTO route VALUES (11, 'settings.contentTypes', 'contentTypes', 6, '/content-type', '[{"single": "public/views/settings/content-type/index.html"}]', false);
INSERT INTO route VALUES (12, 'settings.mediaTypes', 'mediaTypes', 6, '/media-type', '[{"single": "public/views/settings/media-type/index.html"}]', false);
INSERT INTO route VALUES (13, 'settings.dataTypes', 'dataTypes', 6, '/data-type', '[{"single": "public/views/settings/data-type/index.html"}]', false);
INSERT INTO route VALUES (14, 'settings.templates', 'templates', 6, '/template', '[{"single": "public/views/settings/template/index.html"}]', false);
INSERT INTO route VALUES (15, 'settings.scripts', 'scripts', 6, '/script', '[{"single": "public/views/settings/script/index.html"}]', false);
INSERT INTO route VALUES (16, 'settings.stylesheets', 'stylesheets', 6, '/stylesheet', '[{"single": "public/views/settings/stylesheet/index.html"}]', false);
INSERT INTO route VALUES (17, 'settings.contentTypes.new', 'new', 11, '/new?type&parent', '[{"single": "public/views/settings/content-type/new.html"}]', false);
INSERT INTO route VALUES (18, 'settings.mediaTypes.new', 'new', 12, '/new?type&parent', '[{"single": "public/views/settings/media-type/new.html"}]', false);
INSERT INTO route VALUES (19, 'settings.dataTypes.new', 'new', 13, '/new', '[{"single": "public/views/settings/data-type/new.html"}]', false);
INSERT INTO route VALUES (20, 'settings.templates.new', 'new', 14, '/new?parent', '[{"single": "public/views/settings/template/new.html"}]', false);
INSERT INTO route VALUES (21, 'settings.scripts.new', 'new', 15, '/new?type&parent', '[{"single": "public/views/settings/script/new.html"}]', false);
INSERT INTO route VALUES (22, 'settings.stylesheets.new', 'new', 16, '/new?type&parent', '[{"single": "public/views/settings/stylesheet/new.html"}]', false);
INSERT INTO route VALUES (23, 'settings.contentTypes.edit', 'edit', 11, '/edit/:nodeId', '[{"single": "public/views/settings/content-type/edit.html"}]', false);
INSERT INTO route VALUES (24, 'settings.mediaTypes.edit', 'edit', 12, '/edit/:nodeId', '[{"single": "public/views/settings/media-type/edit.html"}]', false);
INSERT INTO route VALUES (25, 'settings.dataTypes.edit', 'edit', 13, '/edit/:nodeId', '[{"single": "public/views/settings/data-type/edit.html"}]', false);
INSERT INTO route VALUES (26, 'settings.templates.edit', 'edit', 14, '/edit/:nodeId', '[{"single": "public/views/settings/template/edit.html"}]', false);
INSERT INTO route VALUES (27, 'settings.scripts.edit', 'edit', 15, '/edit/:name', '[{"single": "public/views/settings/script/edit.html"}]', false);
INSERT INTO route VALUES (28, 'settings.stylesheets.edit', 'edit', 16, '/edit/:name', '[{"single": "public/views/settings/stylesheet/edit.html"}]', false);
INSERT INTO route VALUES (29, 'members.edit', 'edit', 5, '/edit/:id', '[{"single": "public/views/members/edit.html"}]', false);
INSERT INTO route VALUES (30, 'members.new', 'new', 5, '/new', '[{"single": "public/views/members/new.html"}]', false);
INSERT INTO route VALUES (31, 'members.memberTypes', 'memberTypes', 5, '/member-type', '[{"single": "public/views/members/member-type/index.html"}]', false);
INSERT INTO route VALUES (32, 'members.memberTypes.edit', 'edit', 31, '/edit/:nodeId', '[{"single": "public/views/members/member-type/edit.html"}]', false);
INSERT INTO route VALUES (33, 'members.memberTypes.new', 'new', 31, '/new?type&parent', '[{"single": "public/views/members/member-type/new.html"}]', false);


--
-- TOC entry 2379 (class 0 OID 0)
-- Dependencies: 192
-- Name: route_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('route_id_seq', 33, true);


--
-- TOC entry 2360 (class 0 OID 82871)
-- Dependencies: 193
-- Data for Name: template; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO template VALUES (10, 58, '', false, '{0}', 6);
INSERT INTO template VALUES (6, 22, 'Page with sidebars', false, '{20,21}', 6);
INSERT INTO template VALUES (1, 6, 'Layout', false, '{0}', 0);
INSERT INTO template VALUES (11, 60, 'tmpl404', false, '{}', 6);
INSERT INTO template VALUES (7, 25, 'tplLogin', false, '{20,21}', 6);
INSERT INTO template VALUES (8, 26, 'child of test template alias', false, '{}', 25);
INSERT INTO template VALUES (4, 20, 'Sidebar 1', true, '{}', 0);
INSERT INTO template VALUES (5, 21, 'Sidebar 2', true, '{}', 0);
INSERT INTO template VALUES (3, 8, 'Page', false, '{0}', 6);
INSERT INTO template VALUES (9, 42, 'tmpTestTemplate2', false, '{0}', 6);
INSERT INTO template VALUES (2, 7, 'Home', false, '{0}', 6);


--
-- TOC entry 2380 (class 0 OID 0)
-- Dependencies: 194
-- Name: template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('template_id_seq', 12, true);


--
-- TOC entry 2362 (class 0 OID 82880)
-- Dependencies: 195
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO "user" VALUES (1, 'soren', 'Soren', 'Tester', '$2a$10$UNrly6WSmQnm495KAth6Auk4Z.11kjDBRFz8ZKjhqthytKFH/TjKq', 'soren@codeish.com', '2014-10-21 16:51:00.215', NULL, '2015-02-20 00:12:14.354', NULL, 1, 'YNHWOAMEFEOYDIQM66TBRQ7I45LR7FJQFT7FPDULDJXTWEFE2U2Q', '{1}', NULL);
INSERT INTO "user" VALUES (2, 'admin', 'Admin', 'Demo', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'demo@codeish.com', '2014-11-15 16:51:00.215', NULL, '2015-02-27 15:12:12.285', NULL, 1, 'ZMLZCCH7WXTLDCMOXAHV3PMRB3NPR5A33PQUSFLWS2QA5CGH5YPQ', '{1}', NULL);


--
-- TOC entry 2363 (class 0 OID 82887)
-- Dependencies: 196
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO user_group VALUES (2, 'Editor', 'editor', NULL);
INSERT INTO user_group VALUES (3, 'Writer', 'writer', NULL);
INSERT INTO user_group VALUES (1, 'Administrator', 'administrator', '{node_create,node_delete,node_update,node_move,node_copy,node_public_access,node_permissions,node_send_to_publish,node_publish,node_browse,node_change_content_type,admin,content_all,content_create,content_delete,content_update,content_section,content_browse,media_all,media_create,media_delete,media_update,media_section,media_browse,users_all,users_create,users_delete,users_update,users_section,users_browse,user_types_all,user_types_create,user_types_delete,user_types_update,user_types_section,user_types_browse,user_groups_all,user_groups_create,user_groups_delete,user_groups_update,user_groups_section,user_groups_browse,members_all,members_create,members_delete,members_update,members_section,members_browse,member_types_all,member_types_create,member_types_delete,member_types_update,member_types_section,member_types_browse,member_groups_all,member_groups_create,member_groups_delete,member_groups_update,member_groups_section,member_groups_browse,templates_all,templates_create,templates_delete,templates_update,templates_section,templates_browse,scripts_all,scripts_create,scripts_delete,scripts_update,scripts_section,scripts_browse,stylesheets_all,stylesheets_create,stylesheets_delete,stylesheets_update,stylesheets_section,stylesheets_browse,settings_section,settings_all,node_sort,content_types_all,content_types_create,content_types_delete,content_types_update,content_types_section,content_types_browse,media_types_all,media_types_create,media_types_delete,media_types_update,media_types_section,media_types_browse,data_types_all,data_types_create,data_types_delete,data_types_update,data_types_section,data_types_browse}');


--
-- TOC entry 2381 (class 0 OID 0)
-- Dependencies: 197
-- Name: user_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_group_id_seq', 3, true);


--
-- TOC entry 2382 (class 0 OID 0)
-- Dependencies: 198
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 2, true);


-- Completed on 2015-03-02 00:34:31

--
-- PostgreSQL database dump complete
--

