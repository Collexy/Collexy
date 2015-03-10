--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-03-03 12:28:25

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

--
-- TOC entry 2339 (class 0 OID 84013)
-- Dependencies: 172
-- Data for Name: content; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY content (id, node_id, content_type_node_id, meta, public_access) FROM stdin;
1	42	36	{"prop2": "Home page prop 2", "domains": ["localhost:8080", "localhost:8080/test"], "facebook": "facebook.com/home", "copyright": "&copy; 2014 codeish.com", "site_name": "Collexy cms test site", "page_title": "Home page title", "site_tagline": "Test site tagline", "template_node_id": 22}	\N
2	43	39	{"prop2": "prop2a", "prop3": "sample page prop 3", "page_title": "Sample page title", "page_content": "Sample page content goes here", "template_node_id": 25}	\N
3	44	39	{"prop3": "sample child page level 1 page prop 3", "page_title": "Child page level 1 title", "page_content": "Sample page - child page level 1 content goes here", "template_node_id": 25}	\N
4	45	39	{"prop3": "sample child page level 2 page prop 3", "page_title": "Child page level 2 title", "page_content": "Sample page - child page level 2 content goes here1", "template_node_id": 25}	{"groups": [1], "members": [1]}
6	47	37	{"prop2": "prop2a", "prop3": "Hello world prop 3", "page_title": "Hello World", "page_content": "Welcome to Collexy. This is your first post. Edit or delete it, then start blogging", "template_node_id": 23}	\N
7	48	41	{"alt": "Gopher image alt text1", "url": "/media/2014/10/gopher.jpg", "path": "media\\\\gopher.jpg", "caption": "This is the caption of the gopher image1", "description": "Gopher image description1", "temporary property": "lol"}	\N
8	49	40	{"path": "media\\\\2014"}	\N
9	50	40	{"path": "media\\\\2014\\\\12"}	\N
10	51	41	{"alt": "sleeping-kitten.jpg", "path": "media\\\\2014\\\\12\\\\sleeping-kitten.jpg", "title": "sleeping-kitten.jpg", "caption": "sleeping-kitten.jpg", "description": "sleeping-kitten.jpg"}	\N
11	52	41	{"alt": "cat-prays.jpg", "path": "media\\\\2014\\\\12\\\\cat-prays.jpg", "title": "cat-prays.jpg", "caption": "cat-prays.jpg", "description": "cat-prays.jpg"}	\N
5	46	38	{"prop2": "prop2a", "prop3": "Posts overview prop 3", "page_title": "Posts", "page_content": "Sample page content goes here", "template_node_id": 24}	\N
\.


--
-- TOC entry 2370 (class 0 OID 0)
-- Dependencies: 173
-- Name: content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_id_seq', 1, false);


--
-- TOC entry 2341 (class 0 OID 84021)
-- Dependencies: 174
-- Data for Name: content_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) FROM stdin;
1	35	Collexy.Master	Some description	fa fa-folder-o	fa fa-folder-o	\N	\N	[{"name": "Content", "properties": [{"name": "page_title", "order": 1, "data_type_node_id": 2, "help_text": "help text", "description": "The page title overrides the name the page has been given."}]}, {"name": "Properties", "properties": [{"name": "prop2", "order": 1, "data_type_node_id": 2, "help_text": "help text2", "description": "description2"}, {"name": "prop3", "order": 2, "data_type_node_id": 2, "help_text": "help text3", "description": "description3"}]}]
2	36	Collexy.Home	Home Some description	fa fa-folder-o	fa fa-folder-o	35	{"template_node_id": 22, "allowed_templates_node_id": [22], "allowed_content_types_node_id": [37, 38, 39]}	[{"name":"Content","properties":[{"name":"site_name","order":2,"data_type_node_id":2,"help_text":"help text","description":"Site name goes here."},{"name":"site_tagline","order":3,"data_type_node_id":2,"help_text":"help text","description":"Site tagline goes here."},{"name":"copyright","order":4,"data_type_node_id":2,"help_text":"help text","description":"Copyright here."},{"name":"domains","order":5,"data_type_node_id":19,"help_text":"help text","description":"Domains goes here."}]},{"name":"Social","properties":[{"name":"facebook","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your facebook link here."},{"name":"twitter","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your twitter link here."},{"name":"linkedin","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your linkedin link here."}]}]
3	37	Collexy.Post	Post content type desc	fa fa-folder-o	fa fa-folder-o	35	{"template_node_id": 23, "allowed_templates_node_id": [23], "allowed_content_types_node_id": [37]}	[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":4,"help_text":"Help text for page contentent","description":"Page content description"}]}]
4	38	Collexy.PostOverview	Post overview content type desc	fa fa-folder-o	fa fa-folder-o	35	{"template_node_id": 24, "allowed_templates_node_id": [24], "allowed_content_types_node_id": [38]}	[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":4,"help_text":"Help text for page contentent","description":"Page content description"}]}]
5	39	Collexy.Page	Page content type desc	fa fa-folder-o	fa fa-folder-o	35	{"template_node_id": 25, "allowed_templates_node_id": [25], "allowed_content_types_node_id": [39]}	[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":4,"help_text":"Help text for page contentent","description":"Page content description"}]}]
6	40	Collexy.Folder	Folder media type description1	mt-icon1	mt-thumbnail1	\N	{"allowed_content_types_node_id": [40, 41]}	[{"name":"Folder","properties":[{"name":"folder_browser","order":1,"data_type_node_id":15,"help_text":"prop help text","description":"prop description"},{"name":"path","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties"}]
7	41	Collexy.Image	Image content type description	fa fa-folder-o	fa fa-folder-o	\N	null	[{"name":"Image","properties":[{"name":"path","order":1,"data_type_node_id":2,"help_text":"help text","description":"URL goes here."},{"name":"title","order":2,"data_type_node_id":2,"help_text":"help text","description":"The title entered here can override the above one."},{"name":"caption","order":3,"data_type_node_id":4,"help_text":"help text","description":"Caption goes here."},{"name":"alt","order":4,"data_type_node_id":4,"help_text":"help text","description":"Alt goes here."},{"name":"description","order":5,"data_type_node_id":4,"help_text":"help text","description":"Description goes here."},{"name":"file_upload","order":1,"data_type_node_id":16,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties","properties":[{"name":"temporary property","order":1,"data_type_node_id":2,"help_text":"help text","description":"Temporary description goes here."}]}]
\.


--
-- TOC entry 2371 (class 0 OID 0)
-- Dependencies: 175
-- Name: content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_id_seq', 1, false);


--
-- TOC entry 2343 (class 0 OID 84029)
-- Dependencies: 176
-- Data for Name: data_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY data_type (id, node_id, html, alias) FROM stdin;
1	2	<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">	Collexy.TextField
2	3	<input type="number" id="{{prop.name}}" ng-model="data.meta[prop.name]">	Collexy.NumberField
3	4	<textarea id="{{prop.name}}" ng-model="data.meta[prop.name]">	Collexy.Textarea
4	5		Collexy.Radiobox
5	6		Collexy.RadioboxList
6	7		Collexy.Dropdown
7	8		Collexy.DropdownMultiple
8	9		Collexy.Checkbox
9	10		Collexy.CheckboxList
10	11		Collexy.Label
11	12	<colorpicker>The color picker data type is not implemented yet!</colorpicker>	Collexy.ColorPicker
12	13		Collexy.DatePicker
13	14		Collexy.DatePickerTime
14	15	<folderbrowser>This is an awesome folder browser (unimplemented datatype)</folderbrowser>	Collexy.FolderBrowser
15	16	<input type="file" file-input="test.files" multiple />\n<button ng-click="upload()" type="button">Upload</button>\n<li ng-repeat="file in test.files">{{file.name}}</li>\n<!--<input type="file" onchange="angular.element(this).scope().filesChanged(this)" multiple />\n<button ng-click="upload()">Upload</button>\n<li ng-repeat="file in files">{{file.name}}</li>-->	Collexy.Upload
16	17		Collexy.RichtextEditor
17	18		Collexy.TrueFalse
18	19	<div>\n\t<input type="text"/> <button type="button">Add domain</button><br>\n\t<ul>\n\t\t<li ng-repeat="domain in data.meta[prop.name]">{{domain}}</li>\n\t</ul>\n\t<button type="button">Delete selected</button>\n</div>	Collexy.Domains
\.


--
-- TOC entry 2372 (class 0 OID 0)
-- Dependencies: 177
-- Name: data_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_type_id_seq', 1, false);


--
-- TOC entry 2345 (class 0 OID 84037)
-- Dependencies: 178
-- Data for Name: domain; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY domain (id, node_id, name) FROM stdin;
\.


--
-- TOC entry 2373 (class 0 OID 0)
-- Dependencies: 179
-- Name: domain_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('domain_id_seq', 1, false);


--
-- TOC entry 2347 (class 0 OID 84045)
-- Dependencies: 180
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member (id, username, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_node_id, group_ids) FROM stdin;
1	default_member	$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa	default_member@mail.com	{"comments": "default user comments"}	2015-01-22 14:25:38.904	\N	2015-02-19 23:46:00.495	\N	1	GIWES3RHMY5RKC7OZPOQTF5FQFWX32D5VLV3CAKT4HGKP5LZIENA	20	{1}
\.


--
-- TOC entry 2348 (class 0 OID 84053)
-- Dependencies: 181
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member_group (id, name, description) FROM stdin;
1	authenticated_member	All logged in members
\.


--
-- TOC entry 2374 (class 0 OID 0)
-- Dependencies: 182
-- Name: member_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_group_id_seq', 1, false);


--
-- TOC entry 2375 (class 0 OID 0)
-- Dependencies: 183
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_id_seq', 1, false);


--
-- TOC entry 2351 (class 0 OID 84063)
-- Dependencies: 184
-- Data for Name: member_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member_type (id, node_id, alias, description, icon, parent_member_type_node_id, meta, tabs) FROM stdin;
1	20	Collexy.Member	Default member type	fa fa-user fa-fw	1	\N	[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_node_id": 4}]}]
\.


--
-- TOC entry 2376 (class 0 OID 0)
-- Dependencies: 185
-- Name: member_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_type_id_seq', 1, false);


--
-- TOC entry 2353 (class 0 OID 84071)
-- Dependencies: 186
-- Data for Name: menu_link; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, permissions) FROM stdin;
1	1	Content	\N	1	fa fa-newspaper-o fa-fw	\N	1	main	{content_section}
2	2	Media	\N	2	fa fa-file-image-o fa-fw	\N	1	main	{media_section}
3	3	Users	\N	3	fa fa-user fa-fw	\N	1	main	{users_section}
4	4	Members	\N	4	fa fa-users fa-fw	\N	1	main	{members_section}
5	5	Settings	\N	5	fa fa-gear fa-fw	\N	1	main	{settings_section}
6	5.6	Content Types	5	10	fa fa-newspaper-o fa-fw	\N	1	main	{content_types_section}
7	5.7	Media Types	5	11	fa fa-files-o fa-fw	\N	1	main	{media_types_section}
8	5.8	Data Types	5	12	fa fa-check-square-o fa-fw	\N	1	main	{data_types_section}
9	5.9	Templates	5	13	fa fa-eye fa-fw	\N	1	main	{templates_section}
10	6.10	Scripts	5	14	fa fa-file-code-o fa-fw	\N	1	main	{scripts_section}
11	6.11	Stylesheets	5	15	fa fa-desktop fa-fw	\N	1	main	{stylesheets_section}
12	5.12	Member Types	4	30	fa fa-smile-o fa-fw	\N	1	main	{member_types_section}
13	5.13	Member Groups	4	33	fa fa-smile-o fa-fw	\N	1	main	{member_groups_section}
\.


--
-- TOC entry 2377 (class 0 OID 0)
-- Dependencies: 187
-- Name: menu_link_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('menu_link_id_seq', 1, false);


--
-- TOC entry 2355 (class 0 OID 84079)
-- Dependencies: 188
-- Data for Name: node; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) FROM stdin;
1	1	root	5	1	2014-10-22 16:51:00.215	\N	\N	\N
2	1.2	Text input	11	1	2014-10-22 16:51:00.215	1	\N	\N
3	1.3	Numeric input	11	1	2014-10-22 16:51:00.215	1	\N	\N
4	1.4	Textarea	11	1	2014-10-22 16:51:00.215	1	\N	\N
5	1.5	Radiobox	11	1	2014-10-22 16:51:00.215	1	\N	\N
6	1.6	Radiobox list	11	1	2014-10-22 16:51:00.215	1	\N	\N
7	1.7	Dropdown	11	1	2014-10-22 16:51:00.215	1	\N	\N
8	1.8	Dropdown multiple	11	1	2014-10-22 16:51:00.215	1	\N	\N
9	1.9	Checkbox	11	1	2014-10-22 16:51:00.215	1	\N	\N
10	1.10	Checkbox list	11	1	2014-10-22 16:51:00.215	1	\N	\N
11	1.11	Label	11	1	2014-10-22 16:51:00.215	1	\N	\N
12	1.12	Color picker	11	1	2014-10-22 16:51:00.215	1	\N	\N
13	1.13	Date picker	11	1	2014-10-22 16:51:00.215	1	\N	\N
14	1.14	Date picker with time	11	1	2014-10-22 16:51:00.215	1	\N	\N
15	1.15	Folder browser	11	1	2014-10-22 16:51:00.215	1	\N	\N
16	1.16	Upload	11	1	2014-10-22 16:51:00.215	1	\N	\N
17	1.17	Richtext editor	11	1	2014-10-22 16:51:00.215	1	\N	\N
18	1.18	True/false	11	1	2014-10-22 16:51:00.215	1	\N	\N
19	1.19	Domains	11	1	2014-10-22 16:51:00.215	1	\N	\N
20	1.20	Member	12	1	2014-10-22 16:51:00.215	1	\N	\N
21	1.21	Master	3	1	2014-10-22 16:51:00.215	1	\N	\N
22	1.21.22	Home	3	1	2014-10-22 16:51:00.215	21	\N	\N
23	1.21.23	Post	3	1	2014-10-22 16:51:00.215	21	\N	\N
24	1.21.24	Post Overview	3	1	2014-10-22 16:51:00.215	21	\N	\N
25	1.21.25	Page	3	1	2014-10-22 16:51:00.215	21	\N	\N
26	1.21.26	Login	3	1	2014-10-22 16:51:00.215	21	\N	\N
27	1.21.27	Register	3	1	2014-10-22 16:51:00.215	21	\N	\N
28	1.21.28	404	3	1	2014-10-22 16:51:00.215	21	\N	\N
29	1.21.29	Unauthorized	3	1	2014-10-22 16:51:00.215	21	\N	\N
30	1.30	Top Navigation	3	1	2014-10-22 16:51:00.215	1	\N	\N
31	1.31	Post Overview Widget	3	1	2014-10-22 16:51:00.215	1	\N	\N
32	1.32	Featured Pages Widget	3	1	2014-10-22 16:51:00.215	1	\N	\N
33	1.33	Recent Posts Widget	3	1	2014-10-22 16:51:00.215	1	\N	\N
34	1.34	Social	3	1	2014-10-22 16:51:00.215	1	\N	\N
35	1.35	Master	4	1	2014-10-22 16:51:00.215	1	\N	\N
36	1.35.36	Home	4	1	2014-10-22 16:51:00.215	35	\N	\N
37	1.35.37	Post	4	1	2014-10-22 16:51:00.215	35	\N	\N
38	1.35.38	Post Overview	4	1	2014-10-22 16:51:00.215	35	\N	\N
39	1.35.39	Page	4	1	2014-10-22 16:51:00.215	35	\N	\N
40	1.40	Folder	7	1	2014-10-22 16:51:00.215	1	\N	\N
41	1.41	Image	7	1	2014-10-22 16:51:00.215	1	\N	\N
42	1.42	Home	1	1	2014-10-22 16:51:00.215	1	\N	\N
43	1.42.43	Sample Page	1	1	2014-10-22 16:51:00.215	42	[{"id": 2, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]	\N
44	1.42.43.44	Child Page Level 1	1	1	2014-10-26 23:19:44.735	43	\N	[{"id": 1, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]
45	1.42.43.44.45	Child Page Level 2	1	1	2014-10-26 23:19:44.735	44	\N	\N
48	1.48	gopher.jpg	2	1	2014-10-28 15:50:47.303	1	\N	\N
49	1.49	2014	2	1	2014-12-02 01:42:09.979	1	\N	\N
50	1.49.50	12	2	1	2014-12-05 16:18:29.762	49	\N	\N
51	1.49.50.51	cat-prays.jpg	2	1	2014-12-06 13:07:08.943	50	\N	\N
52	1.49.50.52	sleeping-kitten.jpg	2	1	2014-12-06 14:28:52.117	50	\N	\N
47	1.42.46.47	Hello World	1	1	2014-10-22 16:51:00.215	46	\N	\N
46	1.42.46	Posts	1	1	2014-10-22 16:51:00.215	42	\N	\N
\.


--
-- TOC entry 2378 (class 0 OID 0)
-- Dependencies: 189
-- Name: node_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('node_id_seq', 1, false);


--
-- TOC entry 2357 (class 0 OID 84088)
-- Dependencies: 190
-- Data for Name: permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY permission (name) FROM stdin;
node_create
node_delete
node_update
node_move
node_copy
node_public_access
node_permissions
node_send_to_publish
node_publish
node_browse
node_change_content_type
admin
content_all
content_create
content_delete
content_update
content_section
content_browse
media_all
media_create
media_delete
media_update
media_section
media_browse
users_all
users_create
users_delete
users_update
users_section
users_browse
user_types_all
user_types_create
user_types_delete
user_types_update
user_types_section
user_types_browse
user_groups_all
user_groups_create
user_groups_delete
user_groups_update
user_groups_section
user_groups_browse
members_all
members_create
members_delete
members_update
members_section
members_browse
member_types_all
member_types_create
member_types_delete
member_types_update
member_types_section
member_types_browse
member_groups_all
member_groups_create
member_groups_delete
member_groups_update
member_groups_section
member_groups_browse
templates_all
templates_create
templates_delete
templates_update
templates_section
templates_browse
scripts_all
scripts_create
scripts_delete
scripts_update
scripts_section
scripts_browse
stylesheets_all
stylesheets_create
stylesheets_delete
stylesheets_update
stylesheets_section
stylesheets_browse
settings_section
settings_all
node_sort
content_types_all
content_types_create
content_types_delete
content_types_update
content_types_section
content_types_browse
media_types_all
media_types_create
media_types_delete
media_types_update
media_types_section
media_types_browse
data_types_all
data_types_create
data_types_delete
data_types_update
data_types_section
data_types_browse
\.


--
-- TOC entry 2358 (class 0 OID 84094)
-- Dependencies: 191
-- Data for Name: route; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY route (id, path, name, parent_id, url, components, is_abstract) FROM stdin;
1	content	content	\N	/admin/content	[{"single": "public/views/content/index.html"}]	f
2	media	media	\N	/admin/media	[{"single": "public/views/media/index.html"}]	f
3	users	users	\N	/admin/users	[{"single": "public/views/users/index.html"}]	f
4	members	members	\N	/admin/members	[{"single": "public/views/members/index.html"}]	f
5	settings	settings	\N	/admin/settings	[{"single": "public/views/settings/index.html"}]	t
6	content.new	new	1	/new?node_type&content_type_node_id&parent_id	[{"single": "public/views/content/new.html"}]	f
7	content.edit	edit	1	/edit/:nodeId	[{"single": "public/views/content/edit.html"}]	f
8	media.new	new	2	/new?node_type&content_type_node_id&parent_id	[{"single": "public/views/media/new.html"}]	f
9	media.edit	edit	2	/edit/:nodeId	[{"single": "public/views/media/edit.html"}]	f
10	settings.contentTypes	contentTypes	5	/content-type	[{"single": "public/views/settings/content-type/index.html"}]	f
11	settings.mediaTypes	mediaTypes	5	/media-type	[{"single": "public/views/settings/media-type/index.html"}]	f
12	settings.dataTypes	dataTypes	5	/data-type	[{"single": "public/views/settings/data-type/index.html"}]	f
13	settings.templates	templates	5	/template	[{"single": "public/views/settings/template/index.html"}]	f
14	settings.scripts	scripts	5	/script	[{"single": "public/views/settings/script/index.html"}]	f
15	settings.stylesheets	stylesheets	5	/stylesheet	[{"single": "public/views/settings/stylesheet/index.html"}]	f
16	settings.contentTypes.new	new	10	/new?type&parent	[{"single": "public/views/settings/content-type/new.html"}]	f
17	settings.mediaTypes.new	new	11	/new?type&parent	[{"single": "public/views/settings/media-type/new.html"}]	f
18	settings.dataTypes.new	new	12	/new	[{"single": "public/views/settings/data-type/new.html"}]	f
19	settings.templates.new	new	13	/new?parent	[{"single": "public/views/settings/template/new.html"}]	f
20	settings.scripts.new	new	14	/new?type&parent	[{"single": "public/views/settings/script/new.html"}]	f
21	settings.stylesheets.new	new	15	/new?type&parent	[{"single": "public/views/settings/stylesheet/new.html"}]	f
22	settings.contentTypes.edit	edit	10	/edit/:nodeId	[{"single": "public/views/settings/content-type/edit.html"}]	f
23	settings.mediaTypes.edit	edit	11	/edit/:nodeId	[{"single": "public/views/settings/media-type/edit.html"}]	f
24	settings.dataTypes.edit	edit	12	/edit/:nodeId	[{"single": "public/views/settings/data-type/edit.html"}]	f
25	settings.templates.edit	edit	13	/edit/:nodeId	[{"single": "public/views/settings/template/edit.html"}]	f
26	settings.scripts.edit	edit	14	/edit/:name	[{"single": "public/views/settings/script/edit.html"}]	f
27	settings.stylesheets.edit	edit	15	/edit/:name	[{"single": "public/views/settings/stylesheet/edit.html"}]	f
28	members.edit	edit	4	/edit/:id	[{"single": "public/views/members/edit.html"}]	f
29	members.new	new	4	/new	[{"single": "public/views/members/new.html"}]	f
30	members.memberTypes	memberTypes	4	/member-type	[{"single": "public/views/members/member-type/index.html"}]	f
31	members.memberTypes.edit	edit	30	/edit/:nodeId	[{"single": "public/views/members/member-type/edit.html"}]	f
32	members.memberTypes.new	new	30	/new?type&parent	[{"single": "public/views/members/member-type/new.html"}]	f
33	members.memberGroups	memberTypes	4	/member-group	[{"single": "public/views/members/member-group/index.html"}]	f
34	members.memberGroups.edit	edit	33	/edit/:id	[{"single": "public/views/members/member-group/edit.html"}]	f
35	members.memberGroups.new	new	33	/new?type&parent	[{"single": "public/views/members/member-group/new.html"}]	f
\.


--
-- TOC entry 2379 (class 0 OID 0)
-- Dependencies: 192
-- Name: route_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('route_id_seq', 1, false);


--
-- TOC entry 2360 (class 0 OID 84102)
-- Dependencies: 193
-- Data for Name: template; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) FROM stdin;
1	21	Collexy.Master	f	{30,34}	\N
2	22	Collexy.Home	f	{32,33}	21
3	23	Collexy.Post	f	{32,33}	21
4	24	Collexy.PostOverview	f	{32}	21
5	25	Collexy.Page	f	{32,33}	21
6	26	Collexy.Login	f	\N	21
7	27	Collexy.Register	f	\N	21
8	28	Collexy.404	f	\N	21
9	29	Collexy.Unauthorized	f	\N	21
10	30	Collexy.TopNavigation	t	\N	\N
11	31	Collexy.PostOverviewWidget	t	\N	\N
12	32	Collexy.FeaturedPagesWidget	t	\N	\N
13	33	Collexy.RecentPostsWidget	t	\N	\N
14	34	Collexy.Social	t	\N	\N
\.


--
-- TOC entry 2380 (class 0 OID 0)
-- Dependencies: 194
-- Name: template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('template_id_seq', 1, false);


--
-- TOC entry 2362 (class 0 OID 84111)
-- Dependencies: 195
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids, permissions) FROM stdin;
1	admin	Admin	Demo	$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa	demo@codeish.com	2014-11-15 16:51:00.215	\N	2015-03-02 13:07:57.994	\N	1	IF2LJ42JDYHDEH55IXUTCTYMWEFO6PD2VJGYTWVYT353US52LFXQ	{1}	\N
\.


--
-- TOC entry 2363 (class 0 OID 84118)
-- Dependencies: 196
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY user_group (id, name, alias, permissions) FROM stdin;
2	Editor	editor	\N
3	Writer	writer	\N
1	Administrator	administrator	{node_create,node_delete,node_update,node_move,node_copy,node_public_access,node_permissions,node_send_to_publish,node_publish,node_browse,node_change_content_type,admin,content_all,content_create,content_delete,content_update,content_section,content_browse,media_all,media_create,media_delete,media_update,media_section,media_browse,users_all,users_create,users_delete,users_update,users_section,users_browse,user_types_all,user_types_create,user_types_delete,user_types_update,user_types_section,user_types_browse,user_groups_all,user_groups_create,user_groups_delete,user_groups_update,user_groups_section,user_groups_browse,members_all,members_create,members_delete,members_update,members_section,members_browse,member_types_all,member_types_create,member_types_delete,member_types_update,member_types_section,member_types_browse,member_groups_all,member_groups_create,member_groups_delete,member_groups_update,member_groups_section,member_groups_browse,templates_all,templates_create,templates_delete,templates_update,templates_section,templates_browse,scripts_all,scripts_create,scripts_delete,scripts_update,scripts_section,scripts_browse,stylesheets_all,stylesheets_create,stylesheets_delete,stylesheets_update,stylesheets_section,stylesheets_browse,settings_section,settings_all,node_sort,content_types_all,content_types_create,content_types_delete,content_types_update,content_types_section,content_types_browse,media_types_all,media_types_create,media_types_delete,media_types_update,media_types_section,media_types_browse,data_types_all,data_types_create,data_types_delete,data_types_update,data_types_section,data_types_browse}
\.


--
-- TOC entry 2381 (class 0 OID 0)
-- Dependencies: 197
-- Name: user_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_group_id_seq', 1, false);


--
-- TOC entry 2382 (class 0 OID 0)
-- Dependencies: 198
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 1, false);


-- Completed on 2015-03-03 12:28:25

--
-- PostgreSQL database dump complete
--

