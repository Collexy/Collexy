--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-06-02 08:54:10

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

--
-- TOC entry 2373 (class 0 OID 98842)
-- Dependencies: 198
-- Data for Name: content; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (1, '1', NULL, 'Home', 'home', 1, '2015-03-27 21:22:51.805', 2, '{"title": "Home", "domains": ["localhost:8080", "localhost:8080/test"], "copyright": "&copy; 2014 codeish.com", "site_name": "Collexy test site", "about_text": "<p>This is <strong>TXT</strong>, yet another free responsive site template designed by <a href=\"http://n33.co\">AJ</a> for <a href=\"http://html5up.net\">HTML5 UP</a>. It is released under the <a href=\"http://html5up.net/license/\">Creative Commons Attribution</a> license so feel free to use it for whatever you are working on (personal or commercial), just be sure to give us credit for the design. That is basically it :)</p>", "about_title": "About title here", "banner_link": "http://somelink.test", "hide_banner": false, "hide_in_nav": false, "is_featured": false, "template_id": 2, "site_tagline": "Test site tagline", "banner_header": "Banner header goes here", "facebook_link": "facebook.com/home", "banner_link_text": "Click Here!", "banner_subheader": "Banner subheader goes here", "banner_background_image": "/media/Sample Images/TXT/banner.jpg"}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (5, '1.5', 1, 'Get Involved', 'get_involved', 1, '2015-03-27 21:51:57.503', 5, '{"image": "/media/Sample Images/TXT/pic04.jpg", "title": "Get Involved", "content": "Get Involved content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (10, '1.6.10', 6, 'Amazing Post', 'amazing_post', 1, '2015-03-27 22:05:14.042', 3, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Amazing Post", "content": "<p>What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post.</p>", "sub_header": "Amazing subheader here", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (8, '1.6.8', 6, 'Txt Starter Kit For Collexy Released', 'collexy_starter_kit', 1, '2015-03-27 21:59:24.379', 3, '{"title": "TXT Starter Kit For Collexy Released", "content": "The collexy TXT starter kit is just awesome!", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (9, '1.6.9', 6, 'You Need To Read This', 'read_this', 1, '2015-03-27 22:03:09.422', 3, '{"title": "You Need To Read This", "content": "See - you really needed to read this post!", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (6, '1.6', 1, 'Posts', 'posts', 1, '2015-03-27 21:54:10.787', 4, '{"title": "Posts", "hide_in_nav": false, "is_featured": true, "template_id": 5}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (2, '1.2', 1, 'Welcome', 'welcome', 1, '2015-03-27 21:31:55.462', 5, '{"image": "/media/Sample Images/TXT/pic01.jpg", "title": "Welcome", "content": "Welcome content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3, "test_radio_button_list": ["val2"]}', NULL, NULL, '{"2": {"permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}}', NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (3, '1.3', 1, 'Getting Started', 'getting_started', 1, '2015-03-27 21:46:13.265', 5, '{"image": "/media/Sample Images/TXT/pic02.jpg", "title": "Getting Started", "content": "Getting Started content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3}', NULL, NULL, NULL, '{"1": {"permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}, "2": {"permissions": ["node_update"]}}', false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (11, '1.6.11', 6, 'Categories', 'categories', 1, '2015-03-27 22:17:32.659', 6, '{"title": "Categories", "content": "Categories", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (4, '1.4', 1, 'Documentation', 'documentation', 1, '2015-03-27 21:50:23.197', 5, '{"image": "/media/Sample Images/TXT/pic03.jpg", "title": "Documentation", "content": "<p>Documentation content goes here1</p>\n", "hide_in_nav": false, "is_featured": true, "template_id": 3}', '{"1": true}', '{"1": true}', NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (13, '1.13', 1, '404', '404', 1, '2015-03-27 22:20:10.169', 5, '{"title": "404", "content": "404 content goes here", "hide_in_nav": true, "is_featured": false, "template_id": 9}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (14, '1.14', 1, 'Login', 'login', 1, '2015-03-27 22:21:19.482', 5, '{"title": "Login", "content": "Login content goes here", "hide_in_nav": true, "is_featured": false, "template_id": 7}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (12, '1.6.11.12', 11, 'Category 1', 'category_1', 1, '2015-03-27 22:18:45.865', 7, '{"title": "Category 1", "content": "Category 1 content", "hide_in_nav": false, "is_featured": true, "template_id": 6}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (7, '1.6.7', 6, 'Hello World', 'hello_world', 1, '2015-03-27 21:55:03.797', 3, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Hello World", "content": "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas vel tellus venenatis, iaculis eros eu, pellentesque felis. Mauris eleifend venenatis maximus. Fusce condimentum nulla augue, sed elementum nisl dictum ut. Sed ex arcu, efficitur eu finibus ac, convallis ut eros. Ut faucibus elit erat, ac venenatis velit cursus quis. Phasellus sapien elit, ullamcorper ac placerat at, consectetur eget ex. Integer augue sem, tempor nec hendrerit et, ullamcorper ut arcu.</p>\n\n<p>Pellentesque auctor et arcu at tristique. Suspendisse ipsum sapien, vulputate quis cursus eu, rhoncus sed nisi. Nulla euismod mauris vitae tellus iaculis convallis. Sed sodales, risus id sollicitudin aliquet, purus justo convallis dui, sit amet imperdiet elit mauris accumsan velit. Suspendisse dapibus sit amet quam in porta. Nam eleifend sodales dolor eget tempor. Sed pharetra aliquam dui, ultricies scelerisque orci luctus at. Proin eleifend neque quis dolor facilisis sollicitudin. Integer vel ligula nec metus sagittis lacinia at quis arcu. Sed in sem ut mauris laoreet euismod. Integer eu tincidunt lectus, nec varius libero. Proin nec interdum ex. Quisque non lacinia lectus, luctus molestie mi. Fusce lacus est, rhoncus sed nunc at, fermentum luctus ipsum.</p>\n\n<h3>Nunc pulvinar metus a erat fermentum bibendum</h3>\n\n<p>Phasellus mattis tempor dolor vitae feugiat. Sed aliquet massa nisi, in imperdiet mauris auctor in. Nam consectetur ut erat at suscipit. Integer faucibus eleifend rhoncus. Praesent vel bibendum elit, ut molestie metus. Maecenas efficitur, magna vel scelerisque pretium, magna elit vehicula massa, dignissim posuere felis enim a lectus. Donec eget semper urna. Praesent vel nisi id lacus tincidunt pretium vitae eu sapien. Duis varius nisi velit, nec maximus arcu blandit sit amet. Proin dapibus dui et elit dapibus, sit amet rhoncus nisl lobortis. Nunc pretium, lorem eu dignissim mollis, ex nisi mollis lectus, eu blandit arcu nisl vel elit. Mauris risus ipsum, elementum quis eleifend ut, venenatis sit amet orci. Donec ac orci aliquam, vulputate odio eget, pulvinar elit. Cras molestie urna eget justo hendrerit aliquam.</p>\n", "categories": [12], "sub_header": "Subheader for Hello World", "hide_in_nav": false, "is_featured": true, "template_id": 4, "date_published": "2015-16-03 20:55:38"}', NULL, NULL, NULL, NULL, false);


--
-- TOC entry 2369 (class 0 OID 98785)
-- Dependencies: 194
-- Data for Name: content_backup; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (1, '1', NULL, 'Home', 'home', 1, '2015-03-27 21:22:51.805', 2, '{"title": "Home", "domains": ["localhost:8080", "localhost:8080/test"], "copyright": "&copy; 2014 codeish.com", "site_name": "Collexy test site", "about_text": "<p>This is <strong>TXT</strong>, yet another free responsive site template designed by <a href=\"http://n33.co\">AJ</a> for <a href=\"http://html5up.net\">HTML5 UP</a>. It is released under the <a href=\"http://html5up.net/license/\">Creative Commons Attribution</a> license so feel free to use it for whatever you are working on (personal or commercial), just be sure to give us credit for the design. That is basically it :)</p>", "about_title": "About title here", "banner_link": "http://somelink.test", "hide_banner": false, "hide_in_nav": false, "is_featured": false, "template_id": 2, "site_tagline": "Test site tagline", "banner_header": "Banner header goes here", "facebook_link": "facebook.com/home", "banner_link_text": "Click Here!", "banner_subheader": "Banner subheader goes here", "banner_background_image": "/media/Sample Images/TXT/banner.jpg"}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (5, '1.5', 1, 'Get Involved', 'get_involved', 1, '2015-03-27 21:51:57.503', 5, '{"image": "/media/Sample Images/TXT/pic04.jpg", "title": "Get Involved", "content": "Get Involved content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (10, '1.6.10', 6, 'Amazing Post', 'amazing_post', 1, '2015-03-27 22:05:14.042', 3, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Amazing Post", "content": "<p>What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post.</p>", "sub_header": "Amazing subheader here", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (11, '11', NULL, 'Sample Images', 'sample_images', 1, '2015-03-27 22:08:29.415', 6, '{"path": "media\\Sample Images"}', NULL, NULL, NULL, 2, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (12, '11.12', 11, 'TXT', 'txt', 1, '2015-03-27 22:09:40.207', 6, '{"path": "media\\Sample Images\\TXT"}', NULL, NULL, NULL, 2, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (19, '1.6.19', 6, 'Categories', 'categories', 1, '2015-03-27 22:17:32.659', 8, '{"title": "Categories", "content": "Categories", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (20, '1.6.19.20', 19, 'Category 1', 'category_1', 1, '2015-03-27 22:18:45.865', 9, '{"title": "Category 1", "content": "Category 1 content", "hide_in_nav": false, "is_featured": true, "template_id": 6}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (21, '1.21', 1, '404', '404', 1, '2015-03-27 22:20:10.169', 5, '{"title": "404", "content": "404 content goes here", "hide_in_nav": true, "is_featured": false, "template_id": 9}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (22, '1.22', 1, 'Login', 'login', 1, '2015-03-27 22:21:19.482', 5, '{"title": "Login", "content": "Login content goes here", "hide_in_nav": true, "is_featured": false, "template_id": 7}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (8, '1.6.8', 6, 'Txt Starter Kit For Collexy Released', 'collexy_starter_kit', 1, '2015-03-27 21:59:24.379', 3, '{"title": "TXT Starter Kit For Collexy Released", "content": "The collexy TXT starter kit is just awesome!", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (9, '1.6.9', 6, 'You Need To Read This', 'read_this', 1, '2015-03-27 22:03:09.422', 3, '{"title": "You Need To Read This", "content": "See - you really needed to read this post!", "hide_in_nav": false, "is_featured": true, "template_id": 4}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (6, '1.6', 1, 'Posts', 'posts', 1, '2015-03-27 21:54:10.787', 4, '{"title": "Posts", "hide_in_nav": false, "is_featured": true, "template_id": 5}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (15, '11.12.15', 12, 'pic03.jpg', 'pic3', 1, '2015-03-27 22:13:10.64', 7, '{"alt": "pic03.jpg", "path": "media\\Sample Images\\TXT\\pic03.jpg", "title": "pic03.jpg", "caption": "pic03.jpg", "description": "pic03.jpg", "attached_file": {"name": "pic03.jpg", "size": 8984, "type": "image/jpeg", "lastModified": 1427893165426, "lastModifiedDate": "2015-04-01T12:59:25.426Z", "webkitRelativePath": ""}}', NULL, NULL, NULL, 2, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (4, '1.4', 1, 'Documentation', 'documentation', 1, '2015-03-27 21:50:23.197', 5, '{"image": "/media/Sample Images/TXT/pic03.jpg", "title": "Documentation", "content": "<p>Documentation content goes here1</p>\n", "hide_in_nav": false, "is_featured": true, "template_id": 3}', '{"groups": [1], "members": [1]}', NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (16, '11.12.16', 12, 'pic04.jpg', 'pic4', 1, '2015-03-27 22:13:35.245', 7, '{"alt": "pic04.jpg", "path": "media\\Sample Images\\TXT\\pic04.jpg", "title": "pic04.jpg", "caption": "pic04.jpg", "description": "pic04.jpg", "attached_file": {"name": "pic04.jpg", "size": 23592, "type": "image/jpeg", "lastModified": 1427893165426, "lastModifiedDate": "2015-04-01T12:59:25.426Z", "webkitRelativePath": ""}}', NULL, NULL, NULL, 2, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (14, '11.12.14', 12, 'pic02.jpg', 'pic2', 1, '2015-03-27 22:12:24.478', 7, '{"alt": "pic02.jpg", "path": "media\\Sample Images\\TXT\\pic02.jpg", "title": "pic02.jpg", "caption": "pic02.jpg", "description": "pic02.jpg", "attached_file": {"name": "pic02.jpg", "size": 19811, "type": "image/jpeg", "lastModified": 1427893165425, "lastModifiedDate": "2015-04-01T12:59:25.425Z", "webkitRelativePath": ""}}', NULL, NULL, NULL, 2, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (2, '1.2', 1, 'Welcome', 'welcome', 1, '2015-03-27 21:31:55.462', 5, '{"image": "/media/Sample Images/TXT/pic01.jpg", "title": "Welcome", "content": "Welcome content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3, "test_radio_button_list": ["val2"]}', NULL, '{"2": {"permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}}', NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (3, '1.3', 1, 'Getting Started', 'getting_started', 1, '2015-03-27 21:46:13.265', 5, '{"image": "/media/Sample Images/TXT/pic02.jpg", "title": "Getting Started", "content": "Getting Started content goes here", "hide_in_nav": false, "is_featured": true, "template_id": 3}', NULL, NULL, '{"1": {"permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}, "2": {"permissions": ["node_update"]}}', 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (7, '1.6.7', 6, 'Hello World', 'hello_world', 1, '2015-03-27 21:55:03.797', 3, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Hello World", "content": "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas vel tellus venenatis, iaculis eros eu, pellentesque felis. Mauris eleifend venenatis maximus. Fusce condimentum nulla augue, sed elementum nisl dictum ut. Sed ex arcu, efficitur eu finibus ac, convallis ut eros. Ut faucibus elit erat, ac venenatis velit cursus quis. Phasellus sapien elit, ullamcorper ac placerat at, consectetur eget ex. Integer augue sem, tempor nec hendrerit et, ullamcorper ut arcu.</p>\n\n<p>Pellentesque auctor et arcu at tristique. Suspendisse ipsum sapien, vulputate quis cursus eu, rhoncus sed nisi. Nulla euismod mauris vitae tellus iaculis convallis. Sed sodales, risus id sollicitudin aliquet, purus justo convallis dui, sit amet imperdiet elit mauris accumsan velit. Suspendisse dapibus sit amet quam in porta. Nam eleifend sodales dolor eget tempor. Sed pharetra aliquam dui, ultricies scelerisque orci luctus at. Proin eleifend neque quis dolor facilisis sollicitudin. Integer vel ligula nec metus sagittis lacinia at quis arcu. Sed in sem ut mauris laoreet euismod. Integer eu tincidunt lectus, nec varius libero. Proin nec interdum ex. Quisque non lacinia lectus, luctus molestie mi. Fusce lacus est, rhoncus sed nunc at, fermentum luctus ipsum.</p>\n\n<h3>Nunc pulvinar metus a erat fermentum bibendum</h3>\n\n<p>Phasellus mattis tempor dolor vitae feugiat. Sed aliquet massa nisi, in imperdiet mauris auctor in. Nam consectetur ut erat at suscipit. Integer faucibus eleifend rhoncus. Praesent vel bibendum elit, ut molestie metus. Maecenas efficitur, magna vel scelerisque pretium, magna elit vehicula massa, dignissim posuere felis enim a lectus. Donec eget semper urna. Praesent vel nisi id lacus tincidunt pretium vitae eu sapien. Duis varius nisi velit, nec maximus arcu blandit sit amet. Proin dapibus dui et elit dapibus, sit amet rhoncus nisl lobortis. Nunc pretium, lorem eu dignissim mollis, ex nisi mollis lectus, eu blandit arcu nisl vel elit. Mauris risus ipsum, elementum quis eleifend ut, venenatis sit amet orci. Donec ac orci aliquam, vulputate odio eget, pulvinar elit. Cras molestie urna eget justo hendrerit aliquam.</p>\n", "categories": [20], "sub_header": "Subheader for Hello World", "hide_in_nav": false, "is_featured": true, "template_id": 4, "date_published": "2015-16-03 20:55:38"}', NULL, NULL, NULL, 1, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (13, '11.12.13', 12, 'pic01.jpg', 'pic1', 1, '2015-03-27 22:10:35.745', 7, '{"alt": "pic01.jpg", "path": "media\\Sample Images\\TXT\\pic01.jpg", "title": "pic01.jpg", "caption": "pic01.jpg", "description": "pic01.jpg", "attached_file": {"name": "pic01.jpg", "size": 22026, "type": "image/jpeg", "lastModified": 1427893165424, "lastModifiedDate": "2015-04-01T12:59:25.424Z", "webkitRelativePath": ""}}', NULL, NULL, NULL, 2, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (17, '11.12.17', 12, 'pic05.jpg', 'pic5', 1, '2015-03-27 22:14:05.966', 7, '{"alt": "pic05.jpg", "path": "media\\Sample Images\\TXT\\pic05.jpg", "title": "pic05.jpg", "caption": "pic05.jpg", "description": "pic05.jpg", "attached_file": {"name": "pic05.jpg", "size": 74874, "type": "image/jpeg", "lastModified": 1427893165427, "lastModifiedDate": "2015-04-01T12:59:25.427Z", "webkitRelativePath": ""}}', NULL, NULL, NULL, 2, false);
INSERT INTO content_backup (id, path, parent_id, name, alias, created_by, created_date, content_type_id, meta, public_access, user_permissions, user_group_permissions, type_id, is_abstract) VALUES (18, '11.12.18', 12, 'banner.jpg', 'banner', 1, '2015-03-27 22:14:35.241', 7, '{"alt": "banner.jpg", "path": "media\\Sample Images\\TXT\\banner.jpg", "title": "banner.jpg", "caption": "banner.jpg", "description": "banner.jpg", "attached_file": {"name": "banner.jpg", "size": 269179, "type": "image/jpeg", "lastModified": 1427893165424, "lastModifiedDate": "2015-04-01T12:59:25.424Z", "webkitRelativePath": ""}}', NULL, NULL, NULL, 2, false);


--
-- TOC entry 2378 (class 0 OID 0)
-- Dependencies: 193
-- Name: content_backup_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_backup_id_seq', 1, false);


--
-- TOC entry 2379 (class 0 OID 0)
-- Dependencies: 197
-- Name: content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_id_seq', 14, true);


--
-- TOC entry 2362 (class 0 OID 98227)
-- Dependencies: 187
-- Data for Name: content_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids, template_id, allowed_template_ids) VALUES (7, '1.7', 1, 'Category', 'category', 1, '2015-03-27 18:02:14.279', 'Category content type description', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', NULL, '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for category contentent", "description": "Category content description", "data_type_id": 19}]}]', false, false, false, '{7}', '{8}', 6, '{6}');
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids, template_id, allowed_template_ids) VALUES (4, '1.4', 1, 'Post Overview', 'post_overview', 1, '2015-03-27 17:53:03.252', 'Post Overview content type description', 'fa fa-newspaper-o fa-fw', 'fa fa-newspaper-o fa-fw', NULL, '[]', false, false, false, '{3,6}', '{8}', 5, '{5}');
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids, template_id, allowed_template_ids) VALUES (6, '1.6', 1, 'Categories', 'categories', 1, '2015-03-27 17:59:30.925', 'Categories content type description', 'fa fa-folder-open-o fa-fw', 'fa fa-folder-open-o fa-fw', NULL, '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for category contentent", "description": "Category content description", "data_type_id": 19}]}]', false, false, false, '{7}', NULL, NULL, NULL);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids, template_id, allowed_template_ids) VALUES (1, '1', NULL, 'Master', 'master', 1, '2015-03-27 17:46:05.405', 'Master content type description', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', NULL, '[{"name": "Content", "properties": [{"name": "title", "order": 1, "help_text": "help text", "description": "The page title overrides the name the page has been given.", "data_type_id": 1}]}, {"name": "Properties", "properties": [{"name": "hide_in_nav", "order": 1, "help_text": "help text2", "description": "description2", "data_type_id": 18}]}]', false, false, true, NULL, NULL, NULL, NULL);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids, template_id, allowed_template_ids) VALUES (3, '1.3', 1, 'Post', 'post', 1, '2015-03-27 17:51:17.53', 'Post content type description', 'fa fa-file-text-o fa-fw', 'fa fa-file-text-o fa-fw', NULL, '[{"name": "Content", "properties": [{"name": "is_featured", "order": 2, "help_text": "help text2", "description": "description2", "data_type_id": 18}, {"name": "image", "order": 3, "help_text": "Help text for image", "description": "Image url", "data_type_id": 1}, {"name": "sub_header", "order": 4, "help_text": "Help text for subheader", "description": "Subheader description", "data_type_id": 1}, {"name": "content", "order": 5, "help_text": "Help text for post content", "description": "Post content description", "data_type_id": 19}, {"name": "categories", "order": 6, "help_text": "help text2", "description": "description2", "data_type_id": 12}, {"name": "date_published", "order": 7, "help_text": "help date picker with time", "description": "date picker w time", "data_type_id": 11}]}]', false, false, false, '{3}', '{8}', 4, '{4}');
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids, template_id, allowed_template_ids) VALUES (5, '1.5', 1, 'Page', 'page', 1, '2015-03-27 17:54:15.03', 'Page content type description', 'fa fa-file-o fa-fw', 'fa fa-file-o fa-fw', NULL, '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for page contentent", "description": "Page content description", "data_type_id": 19}, {"name": "test_radio_button_list", "order": 3, "help_text": "Help text for test radio button", "description": "Page test radio button desc", "data_type_id": 4}]}]', false, false, false, '{5}', '{8}', 3, '{3,7,8,9,10}');
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids, template_id, allowed_template_ids) VALUES (8, '8', NULL, 'SEO', 'seo', 1, '2015-04-20 14:03:59.172', 'Search Engine Optimization content type', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', NULL, '[{"name": "SEO", "properties": [{"name": "meta_title", "order": 1, "help_text": "Help text for meta title", "description": "Meta title description", "data_type_id": 1}, {"name": "meta_description", "order": 2, "help_text": "Help text for meta description", "description": "Mets description description", "data_type_id": 3}]}]', false, false, true, NULL, '{8}', NULL, NULL);
INSERT INTO content_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids, template_id, allowed_template_ids) VALUES (2, '1.2', 1, 'Home', 'home', 1, '2015-03-27 17:47:50.897', 'Home content type description', 'fa fa-home fa-fw', 'fa fa-home fa-fw', NULL, '[{"name": "Content", "properties": [{"name": "site_name", "order": 2, "help_text": "help text", "description": "Site name goes here.", "data_type_id": 1}, {"name": "site_tagline", "order": 3, "help_text": "help text", "description": "Site tagline goes here.", "data_type_id": 1}, {"name": "copyright", "order": 4, "help_text": "help text", "description": "Copyright here.", "data_type_id": 1}, {"name": "domains", "order": 5, "help_text": "help text", "description": "Domains goes here.", "data_type_id": 17}]}, {"name": "Social", "properties": [{"name": "facebook_link", "order": 1, "help_text": "help text", "description": "Enter your facebook link here.", "data_type_id": 1}, {"name": "twitter_link", "order": 2, "help_text": "help text", "description": "Enter your twitter link here.", "data_type_id": 1}, {"name": "linkedin_link", "order": 3, "help_text": "help text", "description": "Enter your linkedin link here.", "data_type_id": 1}, {"name": "google_link", "order": 4, "help_text": "help text", "description": "Enter your Google+ profile link here.", "data_type_id": 1}, {"name": "rss_link", "order": 5, "help_text": "help text", "description": "Enter your RSS feed link here.", "data_type_id": 1}]}, {"name": "Banner", "properties": [{"name": "hide_banner", "order": 1, "help_text": "help text2", "description": "description2", "data_type_id": 18}, {"name": "banner_header", "order": 2, "help_text": "help text", "description": "Banner header.", "data_type_id": 1}, {"name": "banner_subheader", "order": 3, "help_text": "help text", "description": "Banner subheader.", "data_type_id": 1}, {"name": "banner_link_text", "order": 4, "help_text": "help text", "description": "Banner link text.", "data_type_id": 1}, {"name": "banner_link", "order": 5, "help_text": "help text", "description": "Banner link should ideally use a content picker data type.", "data_type_id": 1}, {"name": "banner_background_image", "order": 6, "help_text": "help text", "description": "This should ideally use the upload data type.", "data_type_id": 1}]}, {"name": "About", "properties": [{"name": "about_title", "order": 1, "help_text": "help text", "description": "About title.", "data_type_id": 1}, {"name": "about_text", "order": 2, "help_text": "help text", "description": "About text.", "data_type_id": 19}]}]', true, false, false, '{3,4,5}', '{8}', 2, '{2}');


--
-- TOC entry 2371 (class 0 OID 98797)
-- Dependencies: 196
-- Data for Name: content_type_backup; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (9, '1.9', 1, 'Category', 'category', 1, '2015-03-27 18:02:14.279', 'Category content type description', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', '{"template_id": 6, "allowed_template_ids": [6]}', '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for category contentent", "description": "Category content description", "data_type_id": 19}]}]', 1, false, false, false, '{9}', '{10}');
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (4, '1.4', 1, 'Post Overview', 'post_overview', 1, '2015-03-27 17:53:03.252', 'Post Overview content type description', 'fa fa-newspaper-o fa-fw', 'fa fa-newspaper-o fa-fw', '{"template_id": 5, "allowed_template_ids": [5]}', '[]', 1, false, false, false, '{3,8}', '{10}');
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (8, '1.8', 1, 'Categories', 'categories', 1, '2015-03-27 17:59:30.925', 'Categories content type description', 'fa fa-folder-open-o fa-fw', 'fa fa-folder-open-o fa-fw', NULL, '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for category contentent", "description": "Category content description", "data_type_id": 19}]}]', 1, false, false, false, '{9}', NULL);
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (1, '1', NULL, 'Master', 'master', 1, '2015-03-27 17:46:05.405', 'Master content type description', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', NULL, '[{"name": "Content", "properties": [{"name": "title", "order": 1, "help_text": "help text", "description": "The page title overrides the name the page has been given.", "data_type_id": 1}]}, {"name": "Properties", "properties": [{"name": "hide_in_nav", "order": 1, "help_text": "help text2", "description": "description2", "data_type_id": 18}]}]', 1, false, false, true, NULL, NULL);
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (3, '1.3', 1, 'Post', 'post', 1, '2015-03-27 17:51:17.53', 'Post content type description', 'fa fa-file-text-o fa-fw', 'fa fa-file-text-o fa-fw', '{"template_id": 4, "allowed_template_ids": [4]}', '[{"name": "Content", "properties": [{"name": "is_featured", "order": 2, "help_text": "help text2", "description": "description2", "data_type_id": 18}, {"name": "image", "order": 3, "help_text": "Help text for image", "description": "Image url", "data_type_id": 1}, {"name": "sub_header", "order": 4, "help_text": "Help text for subheader", "description": "Subheader description", "data_type_id": 1}, {"name": "content", "order": 5, "help_text": "Help text for post content", "description": "Post content description", "data_type_id": 19}, {"name": "categories", "order": 6, "help_text": "help text2", "description": "description2", "data_type_id": 12}, {"name": "date_published", "order": 7, "help_text": "help date picker with time", "description": "date picker w time", "data_type_id": 11}]}]', 1, false, false, false, '{3}', '{10}');
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (5, '1.5', 1, 'Page', 'page', 1, '2015-03-27 17:54:15.03', 'Page content type description', 'fa fa-file-o fa-fw', 'fa fa-file-o fa-fw', '{"template_id": 3, "allowed_template_ids": [3, 7, 8, 9, 10]}', '[{"name": "Content", "properties": [{"name": "content", "order": 2, "help_text": "Help text for page contentent", "description": "Page content description", "data_type_id": 19}, {"name": "test_radio_button_list", "order": 3, "help_text": "Help text for test radio button", "description": "Page test radio button desc", "data_type_id": 4}]}]', 1, false, false, false, '{5}', '{10}');
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (2, '1.2', 1, 'Home', 'home', 1, '2015-03-27 17:47:50.897', 'Home content type description', 'fa fa-home fa-fw', 'fa fa-home fa-fw', '{"template_id": 2, "allowed_template_ids": [2]}', '[{"name": "Content", "properties": [{"name": "site_name", "order": 2, "help_text": "help text", "description": "Site name goes here.", "data_type_id": 1}, {"name": "site_tagline", "order": 3, "help_text": "help text", "description": "Site tagline goes here.", "data_type_id": 1}, {"name": "copyright", "order": 4, "help_text": "help text", "description": "Copyright here.", "data_type_id": 1}, {"name": "domains", "order": 5, "help_text": "help text", "description": "Domains goes here.", "data_type_id": 17}]}, {"name": "Social", "properties": [{"name": "facebook_link", "order": 1, "help_text": "help text", "description": "Enter your facebook link here.", "data_type_id": 1}, {"name": "twitter_link", "order": 2, "help_text": "help text", "description": "Enter your twitter link here.", "data_type_id": 1}, {"name": "linkedin_link", "order": 3, "help_text": "help text", "description": "Enter your linkedin link here.", "data_type_id": 1}, {"name": "google_link", "order": 4, "help_text": "help text", "description": "Enter your Google+ profile link here.", "data_type_id": 1}, {"name": "rss_link", "order": 5, "help_text": "help text", "description": "Enter your RSS feed link here.", "data_type_id": 1}]}, {"name": "Banner", "properties": [{"name": "hide_banner", "order": 1, "help_text": "help text2", "description": "description2", "data_type_id": 18}, {"name": "banner_header", "order": 2, "help_text": "help text", "description": "Banner header.", "data_type_id": 1}, {"name": "banner_subheader", "order": 3, "help_text": "help text", "description": "Banner subheader.", "data_type_id": 1}, {"name": "banner_link_text", "order": 4, "help_text": "help text", "description": "Banner link text.", "data_type_id": 1}, {"name": "banner_link", "order": 5, "help_text": "help text", "description": "Banner link should ideally use a content picker data type.", "data_type_id": 1}, {"name": "banner_background_image", "order": 6, "help_text": "help text", "description": "This should ideally use the upload data type.", "data_type_id": 1}]}, {"name": "About", "properties": [{"name": "about_title", "order": 1, "help_text": "help text", "description": "About title.", "data_type_id": 1}, {"name": "about_text", "order": 2, "help_text": "help text", "description": "About text.", "data_type_id": 19}]}]', 1, true, false, false, '{3,4,5}', '{10}');
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (6, '6', NULL, 'Folder', 'folder', 1, '2015-03-27 17:55:47.388', 'Folder media type description', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', NULL, '[{"name": "Folder", "properties": [{"name": "folder_browser", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 14}, {"name": "path", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 1}]}, {"name": "Properties"}]', 2, true, false, false, '{6,7}', NULL);
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (10, '10', NULL, 'SEO', 'seo', 1, '2015-04-20 14:03:59.172', 'Search Engine Optimization content type', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', NULL, '[{"name": "SEO", "properties": [{"name": "meta_title", "order": 1, "help_text": "Help text for meta title", "description": "Meta title description", "data_type_id": 1}, {"name": "meta_description", "order": 2, "help_text": "Help text for meta description", "description": "Mets description description", "data_type_id": 3}]}]', 1, false, false, true, NULL, '{10}');
INSERT INTO content_type_backup (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, type_id, allow_at_root, is_container, is_abstract, allowed_content_type_ids, composite_content_type_ids) VALUES (7, '7', NULL, 'Image', 'image', 1, '2015-03-27 17:57:48.335', 'Image media type description', 'fa fa-image fa-fw', 'fa fa-image fa-fw', NULL, '[{"name": "Image", "properties": [{"name": "path", "order": 1, "help_text": "help text", "description": "URL goes here.", "data_type_id": 1}, {"name": "title", "order": 2, "help_text": "help text", "description": "The title entered here can override the above one.", "data_type_id": 1}, {"name": "caption", "order": 3, "help_text": "help text", "description": "Caption goes here.", "data_type_id": 3}, {"name": "alt", "order": 4, "help_text": "help text", "description": "Alt goes here.", "data_type_id": 3}, {"name": "description", "order": 5, "help_text": "help text", "description": "Description goes here.", "data_type_id": 3}, {"name": "file_upload", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 16}]}, {"name": "Properties", "properties": [{"name": "temporary property", "order": 1, "help_text": "help text", "description": "Temporary description goes here.", "data_type_id": 1}]}]', 2, true, false, false, NULL, NULL);


--
-- TOC entry 2380 (class 0 OID 0)
-- Dependencies: 195
-- Name: content_type_backup_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_backup_id_seq', 1, false);


--
-- TOC entry 2381 (class 0 OID 0)
-- Dependencies: 188
-- Name: content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_id_seq', 8, true);


--
-- TOC entry 2358 (class 0 OID 98158)
-- Dependencies: 183
-- Data for Name: data_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (1, '1', NULL, 'Text Input', 'text_input', 1, '2015-03-26 23:47:44.854', '<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (2, '2', NULL, 'Numeric Input', 'numeric_input', 1, '2015-03-26 23:47:44.854', '<input type="number" id="{{prop.name}}" ng-model="data.meta[prop.name]">', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (3, '3', NULL, 'Textarea', 'textarea', 1, '2015-03-26 23:47:44.854', '<textarea id="{{prop.name}}" ng-model="data.meta[prop.name]">', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (5, '5', NULL, 'Dropdown', 'dropdown', 1, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (6, '6', NULL, 'Dropdown Multiple', 'dropdown_multiple', 1, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (8, '8', NULL, 'Label', 'label', 1, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (9, '9', NULL, 'Color Picker', 'color_picker', 1, '2015-03-26 23:47:44.854', '<colorpicker>The color picker data type is not implemented yet!</colorpicker>', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (10, '10', NULL, 'Date Picker', 'date_picker', 1, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (11, '11', NULL, 'Date Picker With Time', 'date_picker_time', 1, '2015-03-26 23:47:44.854', '<div class="well">
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
</script>', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (13, '13', NULL, 'Media Picker', 'media_picker', 1, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (17, '17', NULL, 'Domains', 'domains', 1, '2015-03-26 23:47:44.854', '<div>
    <input type="text"/> <button type="button">Add domain</button><br>
    <ul>
        <li ng-repeat="domain in data.meta[prop.name]">{{domain}}</li>
    </ul>
    <button type="button">Delete selected</button>
</div>', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (19, '19', NULL, 'Richtext Editor', 'richtext_editor', 1, '2015-03-26 23:47:44.854', '<textarea ck-editor id="{{prop.name}}" name="{{prop.name}}" ng-model="data.meta[prop.name]" rows="10" cols="80"></textarea>', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (18, '18', NULL, 'True/False', 'true_false', 1, '2015-03-26 23:47:44.854', '<div><label><input type="checkbox" type="checkbox"
        ng-model="data.meta[prop.name]"
        ng-true-value="true"
        ng-false-value="false"
       ></label> {{prop.name}}
</div>', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (12, '12', NULL, 'Content Picker', 'content_picker', 1, '2015-03-26 23:47:44.854', '<div ng-controller="Collexy.DataTypePropertyEditor.ContentPicker">
<div ng-repeat="cn in contentNodes"><label><input type="checkbox" checklist-model="data.meta[prop.name]" checklist-value="cn.id"></label> {{cn.name}}</div>
<br>
<button type="button" ng-click="checkAll()">check all</button>
<button type="button" ng-click="uncheckAll()">uncheck all</button>
</div>', 'Collexy.DataTypeEditor.ContentPicker', '{"content_type_id": 7}');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (4, '4', NULL, 'Radio Button List', 'radio_button_list', 1, '2015-03-26 23:47:44.854', '<div ng-controller="Collexy.DataTypePropertyEditor.RadioButtonList.Controller">
    <ul>
    	<li ng-repeat="option in dataType.meta.options">
    		<label>
    			<input type="radio" name="radio-button-list-{{dataType.alias}}" ng-value="option.value" ng-model="data.meta[prop.name]"/>
    			{{option.label}}
    		</label>
    	</li>
    </ul>
</div>', 'Collexy.DataTypeEditor.RadioButtonList', '{"options": [{"label": "Value 1", "value": "val1"}, {"label": "Value 2", "value": "val2"}]}');
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (7, '7', NULL, 'Checkbox List', 'checkbox_list', 1, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (15, '15', NULL, 'Upload', 'upload', 1, '2015-03-26 23:47:44.854', '<div ng-controller="Collexy.DataTypeEditor.FileUpload.Controller" collexy-file-upload>
    <div ng-show="persistedFiles.length > 0">
        <ul>
            <li ng-repeat="file in persistedFiles">
                {{file}}
            </li>
        </ul>
    </div>
    <input type="file" file-input="files" multiple />
	<button ng-click="upload()" type="button">Upload</button>
    <ul ng-show="files.length > 0">
        <li ng-repeat="file in files">{{file.name}}</li>
    </ul>
</div>', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (16, '16', NULL, 'Upload Multiple', 'upload_multiple', 1, '2015-03-26 23:47:44.854', '<div ng-controller="Collexy.DataTypeEditor.FileUpload.Controller">
    <pre>{{originalData.meta}}</pre>
    <div ng-show="data.meta.attached_file">
	<input type="text" ng-readonly="true" ng-model="data.meta.attached_file">
    </div>
    <div ng-show="originalData.meta.attached_file">
	<img style="max-width: 100%;" src="{{location_url}}/{{originalData.meta.attached_file.name}}"/>
    </div>
    <input type="checkbox" ng-model="clearFiles" id="clearFiles" name="clearFiles"/>
    <label for="clearFiles">Remove file</label>
    <div ng-hide="clearFiles">
        <hr>
        <input type="file" file-input="files"/>
        <ul ng-show="files.length > 0">
            <li ng-repeat="file in files">{{file.name}}</li>
        </ul>
    </div>

    <!--<pre>{{originalData.meta}}</pre>
    <div ng-show="data.meta[prop.name].persisted_files.length > 0">
	<input type="text" ng-readonly="true" ng-model="data.meta[prop.name].persisted_files[$index]" ng-repeat="file in data.meta[prop.name].persisted_files">
    </div>
    <div ng-show="originalData.meta[prop.name].persisted_files.length > 0">
        <ul>
            <li ng-repeat="file in originalData.meta[prop.name].persisted_files">
                <img style="max-width: 100%;" src="{{location_url}}/{{file.name}}"/>
            </li>
        </ul>
    </div>
    <input type="checkbox" ng-model="clearFiles" id="clearFiles" name="clearFiles"/>
    <label for="clearFiles">Remove file(s)</label>
    <div ng-hide="clearFiles">
        <hr>
        <input type="file" file-input="files" prop-name="{{prop.name}}" multiple />
        <ul ng-show="files.length > 0">
            <li ng-repeat="file in files">{{file.name}}</li>
        </ul>
    </div>-->
</div>', NULL, NULL);
INSERT INTO data_type (id, path, parent_id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (14, '14', NULL, 'Folder Browser', 'folder_browser', 1, '2015-03-26 23:47:44.854', '<style>
    .col-ulist-3 {
	columns: 4;
	-webkit-columns: 4;
	-moz-columns: 4;
	padding-left: 0;
    }
    .col-ulist-3 img {
	max-width: 100%;
    } 
    .collexy-folder-browser li a{ 
	display: block; 
	background-color: whitesmoke;
	text-align: center;
	max-width: 100%;
	position: relative;
    }
    .folder-browser-img-placeholder{
	display: inline-block;
	padding: 1em;
    }
    .collexy-folder-browser .folder-browser-img-placeholder i { font-size: 3em; }
    .collexy-folder-browser-img-overlay {
	position: absolute;
	top: 0;
	width: 100%;
	height: 100%;
	background-color: rgba(255,255,255,0.8);
	opacity: 0;
	/**-webkit-transition: all 0.5s ease;
	-moz-transition: all 0.5s ease;
	-o-transition: all 0.5s ease;
	transition: all 0.5s ease;*/
	z-index: 10;
	font-size: 0.8em;
    }
    .collexy-folder-browser-img-overlay:hover {
	opacity:1;
    }
    .collexy-folder-browser img {
	z-index: 1;
    }
</style>
<div ng-controller="Collexy.DataTypePropertyEditor.FolderBrowser" class="collexy-folder-browser">
    <div ng-show="folder.children.length > 0">
        <ul class="col-ulist-3">
            <li ng-repeat="child in folder.children">
		<a ui-sref="media.edit({id:child.id})">
		    <span ng-if="child.meta.attached_file.type == undefined || child.meta.attached_file.type.indexOf(''image'') < 0" class="folder-browser-img-placeholder">
		        <i ng-class="child.content_type.icon"></i><br>
		        {{child.name}}
		    </span>
		    <span class="collexy-folder-browser-img-overlay" ng-if="child.meta.attached_file.type != undefined && child.meta.attached_file.type.indexOf(''image'') > -1">
			Name: {{child.meta.attached_file.name}}<br>
			Type: {{child.meta.attached_file.type}}<br>
			Size: {{child.meta.attached_file.size}} bytes
		    </span>
		    <img ng-if="child.meta.attached_file.type != undefined && child.meta.attached_file.type.indexOf(''image'') > -1" src="{{location_url}}/{{data.name}}/{{child.meta.attached_file.name}}"/>
		</a>
                <!--<img src="{{location_url}}/{{child.name}}"/>-->
            </li>
        </ul>
    </div>
</div>', NULL, NULL);


--
-- TOC entry 2382 (class 0 OID 0)
-- Dependencies: 184
-- Name: data_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_type_id_seq', 19, true);


--
-- TOC entry 2365 (class 0 OID 98724)
-- Dependencies: 190
-- Data for Name: media; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (1, '1', NULL, 'Sample Images', 1, '2015-03-27 22:08:29.415', 1, '{"path": "media\\Sample Images"}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (2, '1.2', 1, 'TXT', 1, '2015-03-27 22:09:40.207', 1, '{"path": "media\\Sample Images\\TXT"}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (3, '1.2.3', 2, 'pic01.jpg', 1, '2015-03-27 22:10:35.745', 2, '{"alt": "pic01.jpg", "path": "media\\Sample Images\\TXT\\pic01.jpg", "title": "pic01.jpg", "caption": "pic01.jpg", "description": "pic01.jpg", "attached_file": {"name": "pic01.jpg", "size": 22026, "type": "image/jpeg", "lastModified": 1427893165424, "lastModifiedDate": "2015-04-01T12:59:25.424Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (4, '1.2.4', 2, 'pic02.jpg', 1, '2015-03-27 22:12:24.478', 2, '{"alt": "pic02.jpg", "path": "media\\Sample Images\\TXT\\pic02.jpg", "title": "pic02.jpg", "caption": "pic02.jpg", "description": "pic02.jpg", "attached_file": {"name": "pic02.jpg", "size": 19811, "type": "image/jpeg", "lastModified": 1427893165425, "lastModifiedDate": "2015-04-01T12:59:25.425Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (5, '1.2.5', 2, 'pic03.jpg', 1, '2015-03-27 22:13:10.64', 2, '{"alt": "pic03.jpg", "path": "media\\Sample Images\\TXT\\pic03.jpg", "title": "pic03.jpg", "caption": "pic03.jpg", "description": "pic03.jpg", "attached_file": {"name": "pic03.jpg", "size": 8984, "type": "image/jpeg", "lastModified": 1427893165426, "lastModifiedDate": "2015-04-01T12:59:25.426Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (6, '1.2.6', 2, 'pic04.jpg', 1, '2015-03-27 22:13:35.245', 2, '{"alt": "pic04.jpg", "path": "media\\Sample Images\\TXT\\pic04.jpg", "title": "pic04.jpg", "caption": "pic04.jpg", "description": "pic04.jpg", "attached_file": {"name": "pic04.jpg", "size": 23592, "type": "image/jpeg", "lastModified": 1427893165426, "lastModifiedDate": "2015-04-01T12:59:25.426Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (7, '1.2.7', 2, 'pic05.jpg', 1, '2015-03-27 22:14:05.966', 2, '{"alt": "pic05.jpg", "path": "media\\Sample Images\\TXT\\pic05.jpg", "title": "pic05.jpg", "caption": "pic05.jpg", "description": "pic05.jpg", "attached_file": {"name": "pic05.jpg", "size": 74874, "type": "image/jpeg", "lastModified": 1427893165427, "lastModifiedDate": "2015-04-01T12:59:25.427Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (8, '1.2.8', 2, 'banner.jpg', 1, '2015-03-27 22:14:35.241', 2, '{"alt": "banner.jpg", "path": "media\\Sample Images\\TXT\\banner.jpg", "title": "banner.jpg", "caption": "banner.jpg", "description": "banner.jpg", "attached_file": {"name": "banner.jpg", "size": 269179, "type": "image/jpeg", "lastModified": 1427893165424, "lastModifiedDate": "2015-04-01T12:59:25.424Z", "webkitRelativePath": ""}}', NULL, NULL);


--
-- TOC entry 2383 (class 0 OID 0)
-- Dependencies: 189
-- Name: media_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('media_id_seq', 8, true);


--
-- TOC entry 2367 (class 0 OID 98772)
-- Dependencies: 192
-- Data for Name: media_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO media_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_media_type_ids, composite_media_type_ids) VALUES (2, '2', NULL, 'Image', 'image', 1, '2015-03-27 17:57:48.335', 'Image media type description', 'fa fa-image fa-fw', 'fa fa-image fa-fw', NULL, '[{"name": "Image", "properties": [{"name": "path", "order": 1, "help_text": "help text", "description": "URL goes here.", "data_type_id": 1}, {"name": "title", "order": 2, "help_text": "help text", "description": "The title entered here can override the above one.", "data_type_id": 1}, {"name": "caption", "order": 3, "help_text": "help text", "description": "Caption goes here.", "data_type_id": 3}, {"name": "alt", "order": 4, "help_text": "help text", "description": "Alt goes here.", "data_type_id": 3}, {"name": "description", "order": 5, "help_text": "help text", "description": "Description goes here.", "data_type_id": 3}, {"name": "file_upload", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 16}]}, {"name": "Properties", "properties": [{"name": "temporary property", "order": 1, "help_text": "help text", "description": "Temporary description goes here.", "data_type_id": 1}]}]', true, false, false, NULL, NULL);
INSERT INTO media_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_media_type_ids, composite_media_type_ids) VALUES (1, '1', NULL, 'Folder', 'folder', 1, '2015-03-27 17:55:47.388', 'Folder media type description', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', NULL, '[{"name": "Folder", "properties": [{"name": "folder_browser", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 14}, {"name": "path", "order": 1, "help_text": "prop help text", "description": "prop description", "data_type_id": 1}]}, {"name": "Properties"}]', true, false, false, '{1,2}', NULL);


--
-- TOC entry 2384 (class 0 OID 0)
-- Dependencies: 191
-- Name: media_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('media_type_id_seq', 2, true);


--
-- TOC entry 2347 (class 0 OID 98009)
-- Dependencies: 172
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member (id, username, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_id, member_group_ids) VALUES (1, 'default_member', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'default_member@mail.com', '{"comments": "default user comments"}', '2015-01-22 14:25:38.904', NULL, '2015-06-01 16:40:41.34', NULL, 1, 'FTZNP34EBRMBYBPZ7GOIUKRBFVLGL553KFKHKFQUWHIKBKYLLZ4Q', 1, '{1}');


--
-- TOC entry 2355 (class 0 OID 98131)
-- Dependencies: 180
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_group (id, path, parent_id, name, alias, created_by, created_date) VALUES (1, '1', NULL, 'Authenticated Member', 'authenticated_member', 1, '2015-03-26 17:09:34.18');


--
-- TOC entry 2385 (class 0 OID 0)
-- Dependencies: 179
-- Name: member_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_group_id_seq', 1, true);


--
-- TOC entry 2386 (class 0 OID 0)
-- Dependencies: 173
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_id_seq', 1, true);


--
-- TOC entry 2356 (class 0 OID 98139)
-- Dependencies: 181
-- Data for Name: member_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, meta, tabs) VALUES (1, '1', NULL, 'Member', 'member', 1, '2015-03-26 19:56:03.85', 'This is the default member type for Collexy members.', 'fa fa-user fa-fw', NULL, '[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_id": 3}]}]');


--
-- TOC entry 2387 (class 0 OID 0)
-- Dependencies: 182
-- Name: member_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_type_id_seq', 1, true);


--
-- TOC entry 2349 (class 0 OID 98044)
-- Dependencies: 174
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
-- TOC entry 2361 (class 0 OID 98211)
-- Dependencies: 186
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
-- TOC entry 2388 (class 0 OID 0)
-- Dependencies: 185
-- Name: template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('template_id_seq', 18, true);


--
-- TOC entry 2350 (class 0 OID 98067)
-- Dependencies: 175
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids, permissions) VALUES (1, 'admin', 'Admin', 'Demo', '$2a$10$CWn3i3CKMJzhRGJ3B9TIeO.ePxgzajTFoB2cH5fpXkiZ7Az9jrmue', 'soren@codeish.com', '2014-11-15 16:51:00.215', NULL, '2015-06-01 16:44:24.794', NULL, 1, 'QJDUM6PO4TU2B65MQLFDTKELEKMKVIOHRAATJEZO2DY4ST7IWMHQ', '{1}', NULL);
INSERT INTO "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids, permissions) VALUES (2, 'test_user', 'Chuck', 'Norris', '$2a$10$CWn3i3CKMJzhRGJ3B9TIeO.ePxgzajTFoB2cH5fpXkiZ7Az9jrmue', 'chuck@norris.com', NULL, NULL, '2015-05-19 09:52:30.916', NULL, 1, '6CNDLLZ224QEIOSFSW7JP4VNXPT5FQKELBJ3NJQCCHMW7WQD3EVA', '{2}', NULL);


--
-- TOC entry 2351 (class 0 OID 98074)
-- Dependencies: 176
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO user_group (id, name, alias, permissions) VALUES (2, 'Editor', 'editor', '{}');
INSERT INTO user_group (id, name, alias, permissions) VALUES (3, 'Writer', 'writer', '{}');
INSERT INTO user_group (id, name, alias, permissions) VALUES (1, 'Administrator', 'administrator', '{node_create,node_delete,node_update,node_move,node_copy,node_public_access,node_permissions,node_send_to_publish,node_publish,node_browse,node_change_content_type,admin,content_all,content_create,content_delete,content_update,content_section,content_browse,media_all,media_create,media_delete,media_update,media_section,media_browse,user_all,user_create,user_delete,user_update,user_section,user_browse,user_type_all,user_type_create,user_type_delete,user_type_update,user_type_section,user_type_browse,user_group_all,user_group_create,user_group_delete,user_group_update,user_group_section,user_group_browse,member_all,member_create,member_delete,member_update,member_section,member_browse,member_type_all,member_type_create,member_type_delete,member_type_update,member_type_section,member_type_browse,member_group_all,member_group_create,member_group_delete,member_group_update,member_group_section,member_group_browse,template_all,template_create,template_delete,template_update,template_section,template_browse,script_all,script_create,script_delete,script_update,script_section,script_browse,stylesheet_all,stylesheet_create,stylesheet_delete,stylesheet_update,stylesheet_section,stylesheet_browse,settings_section,settings_all,node_sort,content_type_all,content_type_create,content_type_delete,content_type_update,content_type_section,content_type_browse,media_type_all,media_type_create,media_type_delete,media_type_update,media_type_section,media_type_browse,data_type_all,data_type_create,data_type_delete,data_type_update,data_type_section,data_type_browse}');


--
-- TOC entry 2389 (class 0 OID 0)
-- Dependencies: 177
-- Name: user_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_group_id_seq', 3, true);


--
-- TOC entry 2390 (class 0 OID 0)
-- Dependencies: 178
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 2, true);


-- Completed on 2015-06-02 08:54:10

--
-- PostgreSQL database dump complete
--

