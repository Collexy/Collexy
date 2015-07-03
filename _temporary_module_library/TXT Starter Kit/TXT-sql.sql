--
-- TOC entry 2384 (class 0 OID 98958)
-- Dependencies: 195
-- Data for Name: content; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (1, '1', NULL, 'Home', 1, '2015-03-27 21:22:51.805', 2, 2, '{"title": "Home", "domains": ["localhost:8080", "localhost:8080/test"], "copyright": "&copy; 2014 codeish.com", "site_name": "%s", "about_text": "<p>This is <strong>TXT</strong>, yet another free responsive site template designed by <a href=\"http://n33.co\">AJ</a> for <a href=\"http://html5up.net\">HTML5 UP</a>. It is released under the <a href=\"http://html5up.net/license/\">Creative Commons Attribution</a> license so feel free to use it for whatever you are working on (personal or commercial), just be sure to give us credit for the design. That is basically it :)</p>", "about_title": "About title here", "banner_link": "http://somelink.test", "hide_banner": false, "hide_in_nav": false, "is_featured": false, "site_tagline": "Test site tagline", "banner_header": "Banner header goes here", "facebook_link": "facebook.com/home", "banner_link_text": "Click Here!", "banner_subheader": "Banner subheader goes here", "banner_background_image": "/media/Sample Images/TXT/banner.jpg"}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (5, '1.5', 1, 'Get Involved', 1, '2015-03-27 21:51:57.503', 5, 3, '{"image": "/media/Sample Images/TXT/pic04.jpg", "title": "Get Involved", "content": "Get Involved content goes here", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (8, '1.6.8', 6, 'Txt Starter Kit For Collexy Released', 1, '2015-03-27 21:59:24.379', 3, 4, '{"title": "TXT Starter Kit For Collexy Released", "content": "The collexy TXT starter kit is just awesome!", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (9, '1.6.9', 6, 'You Need To Read This', 1, '2015-03-27 22:03:09.422', 3, 4, '{"title": "You Need To Read This", "content": "See - you really needed to read this post!", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (6, '1.6', 1, 'Posts', 1, '2015-03-27 21:54:10.787', 4, 5, '{"title": "Posts", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (2, '1.2', 1, 'Welcome', 1, '2015-03-27 21:31:55.462', 5, 3, '{"image": "/media/Sample Images/TXT/pic01.jpg", "title": "Welcome", "content": "Welcome content goes here", "hide_in_nav": false, "is_featured": true, "test_radio_button_list": ["val2"]}', NULL, NULL, '{"2": {"permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}}', NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (3, '1.3', 1, 'Getting Started', 1, '2015-03-27 21:46:13.265', 5, 3, '{"image": "/media/Sample Images/TXT/pic02.jpg", "title": "Getting Started", "content": "Getting Started content goes here", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, '{"1": {"permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}, "2": {"permissions": ["node_update"]}}', false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (11, '1.6.11', 6, 'Categories', 1, '2015-03-27 22:17:32.659', 6, NULL, '{"title": "Categories", "content": "Categories", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (4, '1.4', 1, 'Documentation', 1, '2015-03-27 21:50:23.197', 5, 3, '{"image": "/media/Sample Images/TXT/pic03.jpg", "title": "Documentation", "content": "<p>Documentation content goes here1</p>\n", "hide_in_nav": false, "is_featured": true}', '{"1": true}', '{"1": true}', NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (13, '1.13', 1, '404', 1, '2015-03-27 22:20:10.169', 5, 9, '{"title": "404", "content": "404 content goes here", "hide_in_nav": true, "is_featured": false}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (14, '1.14', 1, 'Login', 1, '2015-03-27 22:21:19.482', 5, 7, '{"title": "Login", "content": "Login content goes here", "hide_in_nav": true, "is_featured": false}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (12, '1.6.11.12', 11, 'Category 1', 1, '2015-03-27 22:18:45.865', 7, 6, '{"title": "Category 1", "content": "Category 1 content", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (7, '1.6.7', 6, 'Hello World', 1, '2015-03-27 21:55:03.797', 3, 4, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Hello World", "content": "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas vel tellus venenatis, iaculis eros eu, pellentesque felis. Mauris eleifend venenatis maximus. Fusce condimentum nulla augue, sed elementum nisl dictum ut. Sed ex arcu, efficitur eu finibus ac, convallis ut eros. Ut faucibus elit erat, ac venenatis velit cursus quis. Phasellus sapien elit, ullamcorper ac placerat at, consectetur eget ex. Integer augue sem, tempor nec hendrerit et, ullamcorper ut arcu.</p>\n\n<p>Pellentesque auctor et arcu at tristique. Suspendisse ipsum sapien, vulputate quis cursus eu, rhoncus sed nisi. Nulla euismod mauris vitae tellus iaculis convallis. Sed sodales, risus id sollicitudin aliquet, purus justo convallis dui, sit amet imperdiet elit mauris accumsan velit. Suspendisse dapibus sit amet quam in porta. Nam eleifend sodales dolor eget tempor. Sed pharetra aliquam dui, ultricies scelerisque orci luctus at. Proin eleifend neque quis dolor facilisis sollicitudin. Integer vel ligula nec metus sagittis lacinia at quis arcu. Sed in sem ut mauris laoreet euismod. Integer eu tincidunt lectus, nec varius libero. Proin nec interdum ex. Quisque non lacinia lectus, luctus molestie mi. Fusce lacus est, rhoncus sed nunc at, fermentum luctus ipsum.</p>\n\n<h3>Nunc pulvinar metus a erat fermentum bibendum</h3>\n\n<p>Phasellus mattis tempor dolor vitae feugiat. Sed aliquet massa nisi, in imperdiet mauris auctor in. Nam consectetur ut erat at suscipit. Integer faucibus eleifend rhoncus. Praesent vel bibendum elit, ut molestie metus. Maecenas efficitur, magna vel scelerisque pretium, magna elit vehicula massa, dignissim posuere felis enim a lectus. Donec eget semper urna. Praesent vel nisi id lacus tincidunt pretium vitae eu sapien. Duis varius nisi velit, nec maximus arcu blandit sit amet. Proin dapibus dui et elit dapibus, sit amet rhoncus nisl lobortis. Nunc pretium, lorem eu dignissim mollis, ex nisi mollis lectus, eu blandit arcu nisl vel elit. Mauris risus ipsum, elementum quis eleifend ut, venenatis sit amet orci. Donec ac orci aliquam, vulputate odio eget, pulvinar elit. Cras molestie urna eget justo hendrerit aliquam.</p>\n", "categories": [12], "sub_header": "Subheader for Hello World", "hide_in_nav": false, "is_featured": true, "date_published": "2015-16-03 20:55:38"}', NULL, NULL, NULL, NULL, false);
INSERT INTO content (id, path, parent_id, name, created_by, created_date, content_type_id, template_id, meta, public_access_members, public_access_member_groups, user_permissions, user_group_permissions, is_abstract) VALUES (10, '1.6.10', 6, 'Amazing Post', 1, '2015-03-27 22:05:14.042', 3, 4, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Amazing Post", "content": "<p>What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post.</p>\n", "sub_header": "Amazing subheader here!", "hide_in_nav": false, "is_featured": true}', NULL, NULL, NULL, NULL, false);


--
-- TOC entry 2396 (class 0 OID 0)
-- Dependencies: 194
-- Name: content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_id_seq', 14, true);













--
-- TOC entry 2382 (class 0 OID 98944)
-- Dependencies: 193
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
-- TOC entry 2398 (class 0 OID 0)
-- Dependencies: 192
-- Name: content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_id_seq', 8, true);










--
-- TOC entry 2374 (class 0 OID 98724)
-- Dependencies: 185
-- Data for Name: media; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (1, '1', NULL, 'Sample Images', 1, '2015-03-27 22:08:29.415', 1, '{"path": "media\\Sample Images"}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (3, '1.2.3', 2, 'pic01.jpg', 1, '2015-03-27 22:10:35.745', 2, '{"alt": "pic01.jpg", "path": "media\\Sample Images\\TXT\\pic01.jpg", "title": "pic01.jpg", "caption": "pic01.jpg", "description": "pic01.jpg", "attached_file": {"name": "pic01.jpg", "size": 22026, "type": "image/jpeg", "lastModified": 1427893165424, "lastModifiedDate": "2015-04-01T12:59:25.424Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (4, '1.2.4', 2, 'pic02.jpg', 1, '2015-03-27 22:12:24.478', 2, '{"alt": "pic02.jpg", "path": "media\\Sample Images\\TXT\\pic02.jpg", "title": "pic02.jpg", "caption": "pic02.jpg", "description": "pic02.jpg", "attached_file": {"name": "pic02.jpg", "size": 19811, "type": "image/jpeg", "lastModified": 1427893165425, "lastModifiedDate": "2015-04-01T12:59:25.425Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (5, '1.2.5', 2, 'pic03.jpg', 1, '2015-03-27 22:13:10.64', 2, '{"alt": "pic03.jpg", "path": "media\\Sample Images\\TXT\\pic03.jpg", "title": "pic03.jpg", "caption": "pic03.jpg", "description": "pic03.jpg", "attached_file": {"name": "pic03.jpg", "size": 8984, "type": "image/jpeg", "lastModified": 1427893165426, "lastModifiedDate": "2015-04-01T12:59:25.426Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (2, '1.2', 1, 'TXT', 1, '2015-03-27 22:09:40.207', 1, '{"path": "media\\Sample Images\\TXT"}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (7, '1.2.7', 2, 'pic05.jpg', 1, '2015-03-27 22:14:05.966', 2, '{"alt": "pic05.jpg", "path": "media\\Sample Images\\TXT\\pic05.jpg", "title": "pic05.jpg", "caption": "pic05.jpg", "description": "pic05.jpg", "attached_file": {"name": "pic05.jpg", "size": 74874, "type": "image/jpeg", "lastModified": 1427893165427, "lastModifiedDate": "2015-04-01T12:59:25.427Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (8, '1.2.8', 2, 'banner.jpg', 1, '2015-03-27 22:14:35.241', 2, '{"alt": "banner.jpg", "path": "media\\Sample Images\\TXT\\banner.jpg", "title": "banner.jpg", "caption": "banner.jpg", "description": "banner.jpg", "attached_file": {"name": "banner.jpg", "size": 269179, "type": "image/jpeg", "lastModified": 1427893165424, "lastModifiedDate": "2015-04-01T12:59:25.424Z", "webkitRelativePath": ""}}', NULL, NULL);
INSERT INTO media (id, path, parent_id, name, created_by, created_date, media_type_id, meta, user_permissions, user_group_permissions) VALUES (6, '1.2.6', 2, 'pic04.jpg', 1, '2015-03-27 22:13:35.245', 2, '{"alt": "pic04.jpg", "path": "media\\Sample Images\\TXT\\pic04.jpg", "title": "pic04.jpg", "caption": "pic04.jpg", "description": "pic04.jpg", "attached_file": {"name": "pic04.jpg", "size": 23592, "type": "image/jpeg", "lastModified": 1427893165426, "lastModifiedDate": "2015-04-01T12:59:25.426Z", "webkitRelativePath": ""}}', NULL, NULL);


--
-- TOC entry 2400 (class 0 OID 0)
-- Dependencies: 184
-- Name: media_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('media_id_seq', 8, true);
















--
-- TOC entry 2376 (class 0 OID 98772)
-- Dependencies: 187
-- Data for Name: media_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO media_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_media_type_ids, composite_media_type_ids) VALUES (1, '1', NULL, 'Folder', 'folder', 0, '2015-03-27 17:55:47.388', 'Folder media type description1', 'fa fa-folder-o fa-fw', 'fa fa-folder-o fa-fw', NULL, '[{"name": "Folder", "properties": [{"name": "folder_browser", "order": 1, "data_type": {"id": 14, "html": "<style>\n    .col-ulist-3 {\n\tcolumns: 4;\n\t-webkit-columns: 4;\n\t-moz-columns: 4;\n\tpadding-left: 0;\n    }\n    .col-ulist-3 img {\n\tmax-width: 100%;\n    } \n    .collexy-folder-browser li a{ \n\tdisplay: block; \n\tbackground-color: whitesmoke;\n\ttext-align: center;\n\tmax-width: 100%;\n\tposition: relative;\n    }\n    .folder-browser-img-placeholder{\n\tdisplay: inline-block;\n\tpadding: 1em;\n    }\n    .collexy-folder-browser .folder-browser-img-placeholder i { font-size: 3em; }\n    .collexy-folder-browser-img-overlay {\n\tposition: absolute;\n\ttop: 0;\n\twidth: 100%;\n\theight: 100%;\n\tbackground-color: rgba(255,255,255,0.8);\n\topacity: 0;\n\t/**-webkit-transition: all 0.5s ease;\n\t-moz-transition: all 0.5s ease;\n\t-o-transition: all 0.5s ease;\n\ttransition: all 0.5s ease;*/\n\tz-index: 10;\n\tfont-size: 0.8em;\n    }\n    .collexy-folder-browser-img-overlay:hover {\n\topacity:1;\n    }\n    .collexy-folder-browser img {\n\tz-index: 1;\n    }\n</style>\n<div ng-controller=\"Collexy.DataTypePropertyEditor.FolderBrowser\" class=\"collexy-folder-browser\">\n    <div ng-show=\"folder.children.length > 0\">\n        <ul class=\"col-ulist-3\">\n            <li ng-repeat=\"child in folder.children\">\n\t\t<a ui-sref=\"media.edit({id:child.id})\">\n\t\t    <span ng-if=\"child.meta.attached_file.type == undefined || child.meta.attached_file.type.indexOf(''image'') < 0\" class=\"folder-browser-img-placeholder\">\n\t\t        <i ng-class=\"child.content_type.icon\"></i><br>\n\t\t        {{child.name}}\n\t\t    </span>\n\t\t    <span class=\"collexy-folder-browser-img-overlay\" ng-if=\"child.meta.attached_file.type != undefined && child.meta.attached_file.type.indexOf(''image'') > -1\">\n\t\t\tName: {{child.meta.attached_file.name}}<br>\n\t\t\tType: {{child.meta.attached_file.type}}<br>\n\t\t\tSize: {{child.meta.attached_file.size}} bytes\n\t\t    </span>\n\t\t    <img ng-if=\"child.meta.attached_file.type != undefined && child.meta.attached_file.type.indexOf(''image'') > -1\" src=\"{{location_url}}/{{data.name}}/{{child.meta.attached_file.name}}\"/>\n\t\t</a>\n                <!--<img src=\"{{location_url}}/{{child.name}}\"/>-->\n            </li>\n        </ul>\n    </div>\n</div>", "name": "Folder Browser", "alias": "folder_browser", "created_by": 1}, "help_text": "prop help text", "description": "prop description", "data_type_id": 14}, {"name": "upload_multiple", "order": 1, "data_type": {"id": 15, "html": "<!--<div ng-controller=\"Collexy.DataTypeEditor.FileUpload.Controller\" collexy-file-upload>\n    <div ng-show=\"persistedFiles.length > 0\">\n        <ul>\n            <li ng-repeat=\"file in persistedFiles\">\n                {{file}}\n            </li>\n        </ul>\n    </div>\n    <input type=\"file\" file-input=\"files\" multiple />\n\t<button ng-click=\"upload()\" type=\"button\">Upload</button>\n    <ul ng-show=\"files.length > 0\">\n        <li ng-repeat=\"file in files\">{{file.name}}</li>\n    </ul>\n</div>-->\n<div ngf-drop ngf-select ng-model=\"$parent.$parent.$parent.files\" class=\"drop-box\" \n     ngf-drag-over-class=\"dragover\" ngf-multiple=\"true\" ngf-allow-dir=\"true\" accept=\"image/*,application/pdf\">Drop file(s) here or click to upload</div>\n\n<div ngf-no-file-drop>File Drag/Drop is not supported for this browser</div>\nFiles:\n<ul>\n    <li ng-repeat=\"f in files\" style=\"font:smaller\">{{f.name}} <img ng-show=\"f != null\" ngf-src=\"f\" class=\"thumb\"></li>\n</ul>\nUpload Log:\n<pre>{{log}}</pre>\n<style>\n    .button {\n        -moz-appearance: button;\n        /* Firefox */\n        -webkit-appearance: button;\n        /* Safari and Chrome */\n        padding: 10px;\n        margin: 10px;\n        width: 70px;\n    }\n    .drop-box {\n        background: #F8F8F8;\n        border: 5px dashed #DDD;\n        /*width: 200px;\n        height: 65px;*/\n        text-align: center;\n        padding: 25px;\n        margin: 10px 0;\n    }\n    .dragover {\n        border: 5px dashed blue;\n    }\n</style>", "name": "Upload", "alias": "upload", "created_by": 1}, "help_text": "prop help text", "description": "prop description", "data_type_id": 15}]}, {"name": "Properties"}]', true, false, false, '{1,2}', '{0}');
INSERT INTO media_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, allow_at_root, is_container, is_abstract, allowed_media_type_ids, composite_media_type_ids) VALUES (2, '2', NULL, 'Image', 'image', 1, '2015-03-27 17:57:48.335', 'Image media type description', 'fa fa-image fa-fw', 'fa fa-image fa-fw', NULL, '[{"name": "Image", "properties": [{"name": "title", "order": 2, "data_type": {"id": 1, "html": "<input type=\"text\" id=\"{{prop.name}}\" ng-model=\"data.meta[prop.name]\"/>", "name": "Text Input", "alias": "text_input", "created_by": 1}, "help_text": "help text", "description": "The title entered here can override the above one.", "data_type_id": 1}, {"name": "caption", "order": 3, "data_type": {"id": 3, "html": "<textarea id=\"{{prop.name}}\" ng-model=\"data.meta[prop.name]\">", "name": "Textarea", "alias": "textarea", "created_by": 1}, "help_text": "help text", "description": "Caption goes here.", "data_type_id": 3}, {"name": "alt", "order": 4, "data_type": {"id": 3, "html": "<textarea id=\"{{prop.name}}\" ng-model=\"data.meta[prop.name]\">", "name": "Textarea", "alias": "textarea", "created_by": 1}, "help_text": "help text", "description": "Alt goes here.", "data_type_id": 3}, {"name": "description", "order": 5, "data_type": {"id": 3, "html": "<textarea id=\"{{prop.name}}\" ng-model=\"data.meta[prop.name]\">", "name": "Textarea", "alias": "textarea", "created_by": 1}, "help_text": "help text", "description": "Description goes here.", "data_type_id": 3}, {"name": "file_upload", "order": 1, "data_type": {"id": 16, "html": "<div ng-controller=\"Collexy.DataTypeEditor.FileUpload.Controller\">\n    <pre>{{originalData.meta}}</pre>\n    <div ng-show=\"data.meta.attached_file\">\n\t<input type=\"text\" ng-readonly=\"true\" ng-model=\"data.meta.attached_file\">\n    </div>\n    <div ng-if=\"!originalData.meta.attached_file\">\n        <img style=\"max-width: 100%;\" src=\"{{location_url}}/{{data.meta.attached_file.name}}\"/>\n    </div>\n    <div ng-if=\"originalData.meta.attached_file\">\n\t\t<img style=\"max-width: 100%;\" src=\"{{location_url}}/{{originalData.meta.attached_file.name}}\"/>\n    </div>\n    <input type=\"checkbox\" ng-model=\"clearFiles\" id=\"clearFiles\" name=\"clearFiles\"/>\n    <label for=\"clearFiles\">Remove file</label>\n    <div ng-hide=\"clearFiles\">\n        <hr>\n        <input type=\"file\" file-input=\"files\"/>\n        <ul ng-show=\"files.length > 0\">\n            <li ng-repeat=\"file in files\">{{file.name}}</li>\n        </ul>\n    </div>\n\n    <!--<pre>{{originalData.meta}}</pre>\n    <div ng-show=\"data.meta[prop.name].persisted_files.length > 0\">\n\t<input type=\"text\" ng-readonly=\"true\" ng-model=\"data.meta[prop.name].persisted_files[$index]\" ng-repeat=\"file in data.meta[prop.name].persisted_files\">\n    </div>\n    <div ng-show=\"originalData.meta[prop.name].persisted_files.length > 0\">\n        <ul>\n            <li ng-repeat=\"file in originalData.meta[prop.name].persisted_files\">\n                <img style=\"max-width: 100%;\" src=\"{{location_url}}/{{file.name}}\"/>\n            </li>\n        </ul>\n    </div>\n    <input type=\"checkbox\" ng-model=\"clearFiles\" id=\"clearFiles\" name=\"clearFiles\"/>\n    <label for=\"clearFiles\">Remove file(s)</label>\n    <div ng-hide=\"clearFiles\">\n        <hr>\n        <input type=\"file\" file-input=\"files\" prop-name=\"{{prop.name}}\" multiple />\n        <ul ng-show=\"files.length > 0\">\n            <li ng-repeat=\"file in files\">{{file.name}}</li>\n        </ul>\n    </div>-->\n</div>", "name": "Upload Multiple", "alias": "upload_multiple", "created_by": 1}, "help_text": "prop help text", "description": "prop description", "data_type_id": 16}]}, {"name": "Properties", "properties": [{"name": "temporary property", "order": 1, "data_type": {"id": 1, "html": "<input type=\"text\" id=\"{{prop.name}}\" ng-model=\"data.meta[prop.name]\"/>", "name": "Text Input", "alias": "text_input", "created_by": 1}, "help_text": "help text", "description": "Temporary description goes here.", "data_type_id": 1}]}]', true, false, false, '{0}', '{0}');


--
-- TOC entry 2401 (class 0 OID 0)
-- Dependencies: 186
-- Name: media_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('media_type_id_seq', 2, true);









--
-- TOC entry 2372 (class 0 OID 98211)
-- Dependencies: 183
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
-- TOC entry 2407 (class 0 OID 0)
-- Dependencies: 182
-- Name: template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('template_id_seq', 19, true);






















