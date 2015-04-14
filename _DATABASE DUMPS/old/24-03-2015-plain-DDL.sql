--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-03-24 12:22:06

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;

--
-- TOC entry 2340 (class 0 OID 95095)
-- Dependencies: 172
-- Data for Name: content; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (2, 43, 39, '{"image": "/media/Sample Images/TXT/pic01.jpg", "title": "Welcome", "content": "Welcome content goes here", "hide_in_nav": false, "is_featured": true, "template_node_id": 25}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (3, 44, 39, '{"image": "/media/Sample Images/TXT/pic02.jpg", "title": "Getting Started", "content": "Getting Started content goes here", "hide_in_nav": false, "is_featured": true, "template_node_id": 25}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (4, 45, 39, '{"image": "/media/Sample Images/TXT/pic03.jpg", "title": "Documentation", "content": "Documentation content goes here1", "hide_in_nav": false, "is_featured": true, "template_node_id": 25}', '{"groups": [1], "members": [1]}');
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (5, 46, 39, '{"image": "/media/Sample Images/TXT/pic04.jpg", "title": "Get Involved", "content": "Get Involved content goes here", "hide_in_nav": false, "is_featured": true, "template_node_id": 25}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (6, 47, 38, '{"title": "Posts", "hide_in_nav": false, "is_featured": true, "template_node_id": 24}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (8, 49, 37, '{"title": "TXT Starter Kit For Collexy Released", "content": "The collexy TXT starter kit is just awesome!", "hide_in_nav": false, "is_featured": true, "template_node_id": 23}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (9, 50, 37, '{"title": "You Need To Read This", "content": "See - you really needed to read this post!", "hide_in_nav": false, "is_featured": true, "template_node_id": 23}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (11, 52, 40, '{"path": "media\\Sample Images"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (12, 53, 40, '{"path": "media\\Sample Images\\TXT"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (14, 55, 41, '{"alt": "pic02.jpg", "path": "media\\Sample Images\\TXT\\pic02.jpg", "title": "pic02.jpg", "caption": "pic02.jpg", "description": "pic02.jpg"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (15, 56, 41, '{"alt": "pic03.jpg", "path": "media\\Sample Images\\TXT\\pic03.jpg", "title": "pic03.jpg", "caption": "pic03.jpg", "description": "pic03.jpg"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (16, 57, 41, '{"alt": "pic04.jpg", "path": "media\\Sample Images\\TXT\\pic04.jpg", "title": "pic04.jpg", "caption": "pic04.jpg", "description": "pic04.jpg"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (17, 58, 41, '{"alt": "pic05.jpg", "path": "media\\Sample Images\\TXT\\pic05.jpg", "title": "pic05.jpg", "caption": "pic05.jpg", "description": "pic05.jpg"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (18, 59, 41, '{"alt": "banner.jpg", "path": "media\\Sample Images\\TXT\\banner.jpg", "title": "banner.jpg", "caption": "banner.jpg", "description": "banner.jpg"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (19, 63, 62, '{"title": "Categories", "content": "Categories", "hide_in_nav": false, "is_featured": true}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (20, 65, 64, '{"title": "Category 1", "content": "Category 1 content", "hide_in_nav": false, "is_featured": true, "template_node_id": 60}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (10, 51, 37, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Amazing Post", "content": "<p>What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post. What an amazing post.</p>", "sub_header": "Amazing subheader here", "hide_in_nav": false, "is_featured": true, "template_node_id": 23}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (1, 42, 36, '{"title": "Home", "domains": ["localhost:8080", "localhost:8080/test"], "copyright": "&copy; 2014 codeish.com", "site_name": "Collexy Test Site", "about_text": "<p>This is <strong>TXT</strong>, yet another free responsive site template designed by <a href=\"http://n33.co\">AJ</a> for <a href=\"http://html5up.net\">HTML5 UP</a>. It is released under the <a href=\"http://html5up.net/license/\">Creative Commons Attribution</a> license so feel free to use it for whatever you are working on (personal or commercial), just be sure to give us credit for the design. That is basically it :)</p>", "about_title": "About title here", "banner_link": "http://somelink.test", "hide_banner": false, "hide_in_nav": false, "is_featured": false, "site_tagline": "Test site tagline", "banner_header": "Banner header goes here", "facebook_link": "facebook.com/home", "banner_link_text": "Click Here!", "banner_subheader": "Banner subheader goes here", "template_node_id": 22, "banner_background_image": "/media/Sample Images/TXT/banner.jpg"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (13, 54, 41, '{"alt": "pic01.jpg", "path": "media\\Sample Images\\TXT\\pic01.jpg", "title": "pic01.jpg", "caption": "pic01.jpg", "description": "pic01.jpg"}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (21, 66, 39, '{"title": "404", "content": "404 content goes here", "hide_in_nav": true, "is_featured": false, "template_node_id": 28}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (22, 67, 39, '{"title": "Login", "content": "Login content goes here", "hide_in_nav": true, "is_featured": false, "template_node_id": 26}', NULL);
INSERT INTO content (id, node_id, content_type_node_id, meta, public_access) VALUES (7, 48, 37, '{"image": "/media/Sample Images/TXT/pic05.jpg", "title": "Hello World", "content": "<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas vel tellus venenatis, iaculis eros eu, pellentesque felis. Mauris eleifend venenatis maximus. Fusce condimentum nulla augue, sed elementum nisl dictum ut. Sed ex arcu, efficitur eu finibus ac, convallis ut eros. Ut faucibus elit erat, ac venenatis velit cursus quis. Phasellus sapien elit, ullamcorper ac placerat at, consectetur eget ex. Integer augue sem, tempor nec hendrerit et, ullamcorper ut arcu.</p>\n\n<p>Pellentesque auctor et arcu at tristique. Suspendisse ipsum sapien, vulputate quis cursus eu, rhoncus sed nisi. Nulla euismod mauris vitae tellus iaculis convallis. Sed sodales, risus id sollicitudin aliquet, purus justo convallis dui, sit amet imperdiet elit mauris accumsan velit. Suspendisse dapibus sit amet quam in porta. Nam eleifend sodales dolor eget tempor. Sed pharetra aliquam dui, ultricies scelerisque orci luctus at. Proin eleifend neque quis dolor facilisis sollicitudin. Integer vel ligula nec metus sagittis lacinia at quis arcu. Sed in sem ut mauris laoreet euismod. Integer eu tincidunt lectus, nec varius libero. Proin nec interdum ex. Quisque non lacinia lectus, luctus molestie mi. Fusce lacus est, rhoncus sed nunc at, fermentum luctus ipsum.</p>\n\n<h3>Nunc pulvinar metus a erat fermentum bibendum</h3>\n\n<p>Phasellus mattis tempor dolor vitae feugiat. Sed aliquet massa nisi, in imperdiet mauris auctor in. Nam consectetur ut erat at suscipit. Integer faucibus eleifend rhoncus. Praesent vel bibendum elit, ut molestie metus. Maecenas efficitur, magna vel scelerisque pretium, magna elit vehicula massa, dignissim posuere felis enim a lectus. Donec eget semper urna. Praesent vel nisi id lacus tincidunt pretium vitae eu sapien. Duis varius nisi velit, nec maximus arcu blandit sit amet. Proin dapibus dui et elit dapibus, sit amet rhoncus nisl lobortis. Nunc pretium, lorem eu dignissim mollis, ex nisi mollis lectus, eu blandit arcu nisl vel elit. Mauris risus ipsum, elementum quis eleifend ut, venenatis sit amet orci. Donec ac orci aliquam, vulputate odio eget, pulvinar elit. Cras molestie urna eget justo hendrerit aliquam.</p>\n", "categories": [65], "sub_header": "Subheader for Hello World", "hide_in_nav": false, "is_featured": true, "date_published": "2015-16-03 20:55:38", "template_node_id": 23}', NULL);


--
-- TOC entry 2371 (class 0 OID 0)
-- Dependencies: 173
-- Name: content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_id_seq', 22, true);


--
-- TOC entry 2342 (class 0 OID 95103)
-- Dependencies: 174
-- Data for Name: content_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (1, 35, 'Collexy.Master', 'Some description', 'fa fa-folder-o', 'fa fa-folder-o', NULL, NULL, '[{"name":"Content","properties":[{"name":"title","order":1,"data_type_node_id":2,"help_text":"help text","description":"The page title overrides the name the page has been given."}]},{"name":"Properties","properties":[{"name":"hide_in_nav","order":1,"data_type_node_id":18,"help_text":"help text2","description":"description2"}]}]');
INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (8, 62, 'Categories', 'Cat desc', 'fa fa-folder-open-o fa-fw', 'fa fa-tag', 35, '{"allowed_content_types_node_id": [64]}', '[{"name":"Content","properties":[{"name":"content","order":2,"data_type_node_id":4,"help_text":"Help text for category contentent","description":"Category content description"}]}]');
INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (9, 64, 'Category', 'Cat desc', 'fa fa-folder-o fa-fw', 'fa fa-tag', 35, '{"template_node_id": 60, "allowed_templates_node_id": [60], "allowed_content_types_node_id": [64]}', '[{"name":"Content","properties":[{"name":"content","order":2,"data_type_node_id":4,"help_text":"Help text for category contentent","description":"Category content description"}]}]');
INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (3, 37, 'Collexy.Post', 'Post content type desc', 'fa fa-file-text-o fa-fw', 'fa fa-folder-o', 35, '{"template_node_id": 23, "allowed_templates_node_id": [23], "allowed_content_types_node_id": [37]}', '[{"name":"Content","properties":[{"name":"is_featured","order":2,"data_type_node_id":18,"help_text":"help text2","description":"description2"},{"name":"image","order":3,"data_type_node_id":2,"help_text":"Help text for image","description":"Image url"},{"name":"sub_header","order":4,"data_type_node_id":2,"help_text":"Help text for subheader","description":"Subheader description"},{"name":"content","order":5,"data_type_node_id":17,"help_text":"Help text for post content","description":"Post content description"},{"name":"categories","order":6,"data_type_node_id":6,"help_text":"help text2","description":"description2"},{"name":"date_published","order":7,"data_type_node_id":14,"help_text":"help date picker with time","description":"date picker w time"}]}]');
INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (5, 39, 'Collexy.Page', 'Page content type desc', 'fa fa-file-o fa-fw', 'fa fa-folder-o', 35, '{"template_node_id": 25, "allowed_templates_node_id": [25, 28, 26, 27, 29], "allowed_content_types_node_id": [39]}', '[{"name":"Content","properties":[{"name":"content","order":2,"data_type_node_id":4,"help_text":"Help text for page contentent","description":"Page content description"}]}]');
INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (2, 36, 'Collexy.Home', 'Home Some description', 'fa fa-home fa-fw', 'fa fa-folder-o', 35, '{"template_node_id": 22, "allowed_templates_node_id": [22], "allowed_content_types_node_id": [37, 38, 39]}', '[{"name":"Content","properties":[{"name":"site_name","order":2,"data_type_node_id":2,"help_text":"help text","description":"Site name goes here."},{"name":"site_tagline","order":3,"data_type_node_id":2,"help_text":"help text","description":"Site tagline goes here."},{"name":"copyright","order":4,"data_type_node_id":2,"help_text":"help text","description":"Copyright here."},{"name":"domains","order":5,"data_type_node_id":19,"help_text":"help text","description":"Domains goes here."}]},{"name":"Social","properties":[{"name":"facebook_link","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your facebook link here."},{"name":"twitter_link","order":2,"data_type_node_id":2,"help_text":"help text","description":"Enter your twitter link here."},{"name":"linkedin_link","order":3,"data_type_node_id":2,"help_text":"help text","description":"Enter your linkedin link here."},{"name":"google_link","order":4,"data_type_node_id":2,"help_text":"help text","description":"Enter your Google+ profile link here."},{"name":"rss_link","order":5,"data_type_node_id":2,"help_text":"help text","description":"Enter your RSS feed link here."}]},{"name":"Banner","properties":[{"name": "hide_banner", "order": 1, "data_type_node_id": 18, "help_text": "help text2", "description": "description2"},{"name": "banner_header", "order": 2, "data_type_node_id": 2, "help_text": "help text", "description": "Banner header."},{"name": "banner_subheader", "order": 3, "data_type_node_id": 2, "help_text": "help text", "description": "Banner subheader."},{"name": "banner_link_text", "order": 4, "data_type_node_id": 2, "help_text": "help text", "description": "Banner link text."},{"name": "banner_link", "order": 5, "data_type_node_id": 2, "help_text": "help text", "description": "Banner link should ideally use a content picker data type."},{"name": "banner_background_image", "order": 6, "data_type_node_id": 2, "help_text": "help text", "description": "This should ideally use the upload data type."}]},{"name":"About","properties":[{"name": "about_title", "order": 1, "data_type_node_id": 2, "help_text": "help text", "description": "About title."},{"name": "about_text", "order": 2, "data_type_node_id": 4, "help_text": "help text", "description": "About text."}]}]');
INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (4, 38, 'Collexy.PostOverview', 'Post overview content type desc', 'fa fa-newspaper-o fa-fw', 'fa fa-folder-o', 35, '{"template_node_id": 24, "allowed_templates_node_id": [24], "allowed_content_types_node_id": [64, 37]}', '[]');
INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (6, 40, 'mtFolder', 'Folder media type description1', 'fa fa-folder-o fa-fw', 'mt-thumbnail1', NULL, '{"allowed_content_types_node_id": [40, 41]}', '[{"name":"Folder","properties":[{"name":"folder_browser","order":1,"data_type_node_id":34,"help_text":"prop help text","description":"prop description"},{"name":"path","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties"}]');
INSERT INTO content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES (7, 41, 'Collexy.Image', 'Image content type description', 'fa fa-image fa-fw', 'fa fa-folder-o', NULL, 'null', '[{"name":"Image","properties":[{"name":"path","order":1,"data_type_node_id":2,"help_text":"help text","description":"URL goes here."},{"name":"title","order":2,"data_type_node_id":2,"help_text":"help text","description":"The title entered here can override the above one."},{"name":"caption","order":3,"data_type_node_id":4,"help_text":"help text","description":"Caption goes here."},{"name":"alt","order":4,"data_type_node_id":4,"help_text":"help text","description":"Alt goes here."},{"name":"description","order":5,"data_type_node_id":4,"help_text":"help text","description":"Description goes here."},{"name":"file_upload","order":1,"data_type_node_id":16,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties","properties":[{"name":"temporary property","order":1,"data_type_node_id":2,"help_text":"help text","description":"Temporary description goes here."}]}]');


--
-- TOC entry 2372 (class 0 OID 0)
-- Dependencies: 175
-- Name: content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_id_seq', 9, true);


--
-- TOC entry 2344 (class 0 OID 95111)
-- Dependencies: 176
-- Data for Name: data_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO data_type (id, node_id, html, alias) VALUES (1, 2, '<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">', 'Collexy.TextField');
INSERT INTO data_type (id, node_id, html, alias) VALUES (2, 3, '<input type="number" id="{{prop.name}}" ng-model="data.meta[prop.name]">', 'Collexy.NumberField');
INSERT INTO data_type (id, node_id, html, alias) VALUES (3, 4, '<textarea id="{{prop.name}}" ng-model="data.meta[prop.name]">', 'Collexy.Textarea');
INSERT INTO data_type (id, node_id, html, alias) VALUES (4, 5, '', 'Collexy.Radiobox');
INSERT INTO data_type (id, node_id, html, alias) VALUES (6, 7, '', 'Collexy.Dropdown');
INSERT INTO data_type (id, node_id, html, alias) VALUES (7, 8, '', 'Collexy.DropdownMultiple');
INSERT INTO data_type (id, node_id, html, alias) VALUES (9, 10, '', 'Collexy.CheckboxList');
INSERT INTO data_type (id, node_id, html, alias) VALUES (10, 11, '', 'Collexy.Label');
INSERT INTO data_type (id, node_id, html, alias) VALUES (11, 12, '<colorpicker>The color picker data type is not implemented yet!</colorpicker>', 'Collexy.ColorPicker');
INSERT INTO data_type (id, node_id, html, alias) VALUES (12, 13, '', 'Collexy.DatePicker');
INSERT INTO data_type (id, node_id, html, alias) VALUES (14, 15, '<folderbrowser>This is an awesome folder browser (unimplemented datatype)</folderbrowser>', 'Collexy.FolderBrowser');
INSERT INTO data_type (id, node_id, html, alias) VALUES (18, 19, '<div>
    <input type="text"/> <button type="button">Add domain</button><br>
    <ul>
        <li ng-repeat="domain in data.meta[prop.name]">{{domain}}</li>
    </ul>
    <button type="button">Delete selected</button>
</div>', 'Collexy.Domains');
INSERT INTO data_type (id, node_id, html, alias) VALUES (8, 9, '', 'Collexy.MediaPicker');
INSERT INTO data_type (id, node_id, html, alias) VALUES (5, 6, '<!--<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">-->

<div ng-repeat="cn in contentNodes"><label><input type="checkbox" checklist-model="data.meta[prop.name]" checklist-value="cn.id"></label> {{cn.name}}</div>
<br>
<button type="button" ng-click="checkAll()">check all</button>
<button type="button" ng-click="uncheckAll()">uncheck all</button>', 'Collexy.ContentPicker');
INSERT INTO data_type (id, node_id, html, alias) VALUES (16, 17, '<textarea ck-editor id="{{prop.name}}" name="{{prop.name}}" ng-model="data.meta[prop.name]" rows="10" cols="80"></textarea>', 'Collexy.RichtextEditor');
INSERT INTO data_type (id, node_id, html, alias) VALUES (13, 14, '<div class="well">
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
    $(''#datetimepicker1'').datetimepicker({
      language: ''en''
    });
  });
</script>', 'Collexy.DatePickerTime');
INSERT INTO data_type (id, node_id, html, alias) VALUES (15, 16, '<input type="file" file-input="test.files" multiple />
<button ng-click="upload()" type="button">Upload</button>
<li ng-repeat="file in test.files">{{file.name}}</li>


<!--<input type="file" onchange="angular.element(this).scope().filesChanged(this)" multiple />
<button ng-click="upload()">Upload</button>
<li ng-repeat="file in files">{{file.name}}</li>-->', 'Collexy.Upload');
INSERT INTO data_type (id, node_id, html, alias) VALUES (17, 18, '<div><label><input type="checkbox" type="checkbox"
       ng-model="data.meta[prop.name]"
       [name="{{prop.name}}"]
       [ng-true-value="true"]
       [ng-false-value=""]
       [ng-change=""]></label> {{prop.name}}
</div>', 'Collexy.TrueFalse');


--
-- TOC entry 2373 (class 0 OID 0)
-- Dependencies: 177
-- Name: data_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_type_id_seq', 18, true);


--
-- TOC entry 2346 (class 0 OID 95119)
-- Dependencies: 178
-- Data for Name: domain; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 2374 (class 0 OID 0)
-- Dependencies: 179
-- Name: domain_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('domain_id_seq', 1, true);


--
-- TOC entry 2348 (class 0 OID 95127)
-- Dependencies: 180
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member (id, username, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_node_id, group_ids) VALUES (1, 'default_member', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'default_member@mail.com', '{"comments": "default user comments"}', '2015-01-22 14:25:38.904', NULL, '2015-03-17 11:52:55.345', NULL, 1, 'LVWNOSHTKQ6CMNF3SMXLBVWPH7N6TOWCYNRER2L64O2J23Y4K2MQ', 20, '{1}');


--
-- TOC entry 2349 (class 0 OID 95135)
-- Dependencies: 181
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_group (id, name, description) VALUES (1, 'authenticated_member', 'All logged in members');


--
-- TOC entry 2375 (class 0 OID 0)
-- Dependencies: 182
-- Name: member_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_group_id_seq', 1, true);


--
-- TOC entry 2376 (class 0 OID 0)
-- Dependencies: 183
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_id_seq', 1, true);


--
-- TOC entry 2352 (class 0 OID 95145)
-- Dependencies: 184
-- Data for Name: member_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_type (id, node_id, alias, description, icon, parent_member_type_node_id, meta, tabs) VALUES (1, 20, 'Collexy.Member', 'Default member type', 'fa fa-user fa-fw', 1, NULL, '[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_node_id": 4}]}]');


--
-- TOC entry 2377 (class 0 OID 0)
-- Dependencies: 185
-- Name: member_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_type_id_seq', 1, true);


--
-- TOC entry 2354 (class 0 OID 95153)
-- Dependencies: 186
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


--
-- TOC entry 2378 (class 0 OID 0)
-- Dependencies: 187
-- Name: menu_link_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('menu_link_id_seq', 13, true);


--
-- TOC entry 2356 (class 0 OID 95161)
-- Dependencies: 188
-- Data for Name: node; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (1, '1', 'root', 5, 1, '2014-10-22 16:51:00.215', NULL, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (2, '1.2', 'Text input', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (3, '1.3', 'Numeric input', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (4, '1.4', 'Textarea', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (5, '1.5', 'Radiobox', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (7, '1.7', 'Dropdown', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (8, '1.8', 'Dropdown multiple', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (10, '1.10', 'Checkbox list', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (11, '1.11', 'Label', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (12, '1.12', 'Color picker', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (13, '1.13', 'Date picker', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (15, '1.15', 'Folder browser', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (19, '1.19', 'Domains', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (20, '1.20', 'Member', 12, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (21, '1.21', 'Layout', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (22, '1.21.22', 'Home', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (23, '1.21.23', 'Post', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (24, '1.21.24', 'Post Overview', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (25, '1.21.25', 'Page', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (26, '1.21.26', 'Login', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (27, '1.21.27', 'Register', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (30, '1.30', 'Top Navigation', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (31, '1.31', 'Post Overview Widget', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (32, '1.32', 'Featured Pages Widget', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (33, '1.33', 'Recent Posts Widget', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (34, '1.34', 'Social', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (35, '1.35', 'Master', 4, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (36, '1.35.36', 'Home', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (37, '1.35.37', 'Post', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (38, '1.35.38', 'Post Overview', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (40, '1.40', 'Folder', 7, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (41, '1.41', 'Image', 7, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (43, '1.42.43', 'Welcome', 1, 1, '2014-10-22 16:51:00.215', 42, '[{"id": 2, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]', NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (44, '1.42.44', 'Getting Started', 1, 1, '2014-10-26 23:19:44.735', 42, NULL, '[{"id": 1, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]');
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (45, '1.42.45', 'Documentation', 1, 1, '2014-10-26 23:19:44.735', 42, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (46, '1.42.46', 'Get Involved', 1, 1, '2014-10-26 23:19:44.735', 42, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (47, '1.42.47', 'Posts', 1, 1, '2014-10-22 16:51:00.215', 42, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (49, '1.42.47.49', 'TXT Starter Kit For Collexy Released', 1, 1, '2014-10-22 16:51:00.215', 47, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (50, '1.42.47.50', 'You Need To Read This', 1, 1, '2014-10-22 16:51:00.215', 47, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (52, '1.52', 'Sample Images', 2, 1, '2014-12-02 01:42:09.979', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (53, '1.52.53', 'TXT', 2, 1, '2014-12-05 16:18:29.762', 52, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (55, '1.42.53.55', 'pic02.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (56, '1.42.53.56', 'pic03.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (57, '1.42.53.57', 'pic04.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (58, '1.42.53.58', 'pic05.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (59, '1.42.53.59', 'banner.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (42, '1.42', 'Home', 1, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (9, '1.9', 'Media Picker', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (17, '1.17', 'Richtext editor', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (61, '1.61', 'Category List Widget', 3, 1, '2015-03-10 00:44:02.866', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (65, '1.42.47.63.65', 'Category 1', 1, 1, '2015-03-10 01:28:32.023', 63, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (64, '1.35.64', 'Category', 4, 1, '2015-03-10 01:17:20.015', 35, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (62, '1.35.62', 'Categories', 4, 1, '2015-03-10 00:44:02.866', 35, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (60, '1.21.60', 'Category', 3, 1, '2015-03-10 00:44:02.866', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (63, '1.42.47.63', 'Categories', 1, 1, '2015-03-10 00:44:02.866', 47, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (51, '1.42.47.51', 'Amazing Post', 1, 1, '2015-03-12 16:51:00.215', 47, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (6, '1.6', 'Content Picker', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (39, '1.35.39', 'Page', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (66, '1.42.66', '404', 1, 1, '2015-03-12 18:42:54.439', 42, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (28, '1.21.28', '404', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (14, '1.14', 'Date picker with time', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (16, '1.16', 'Upload', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (54, '1.52.53.54', 'pic01.jpg', 2, 1, '2014-12-06 13:07:08.943', 53, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (18, '1.18', 'True/false', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (48, '1.42.47.48', 'Hello World', 1, 1, '2014-10-22 16:51:00.215', 47, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (29, '1.21.29', 'Unauthorized', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (67, '1.42.67', 'Login', 1, 1, '2015-03-12 20:49:09.637', 42, NULL, NULL);
INSERT INTO node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) VALUES (68, '1.68', 'Login Widget', 3, 1, '2015-03-13 10:53:45.924', 1, NULL, NULL);


--
-- TOC entry 2379 (class 0 OID 0)
-- Dependencies: 189
-- Name: node_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('node_id_seq', 68, true);


--
-- TOC entry 2358 (class 0 OID 95170)
-- Dependencies: 190
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
-- TOC entry 2359 (class 0 OID 95176)
-- Dependencies: 191
-- Data for Name: route; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (1, 'content', 'content', NULL, '/admin/content', '[{"single": "public/views/content/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (2, 'media', 'media', NULL, '/admin/media', '[{"single": "public/views/media/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (3, 'users', 'users', NULL, '/admin/users', '[{"single": "public/views/users/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (4, 'members', 'members', NULL, '/admin/members', '[{"single": "public/views/members/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (5, 'settings', 'settings', NULL, '/admin/settings', '[{"single": "public/views/settings/index.html"}]', true);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (6, 'content.new', 'new', 1, '/new?node_type&content_type_node_id&parent_id', '[{"single": "public/views/content/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (7, 'content.edit', 'edit', 1, '/edit/:nodeId', '[{"single": "public/views/content/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (8, 'media.new', 'new', 2, '/new?node_type&content_type_node_id&parent_id', '[{"single": "public/views/media/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (9, 'media.edit', 'edit', 2, '/edit/:nodeId', '[{"single": "public/views/media/edit.html"}]', false);
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
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (22, 'settings.contentTypes.edit', 'edit', 10, '/edit/:nodeId', '[{"single": "public/views/settings/content-type/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (23, 'settings.mediaTypes.edit', 'edit', 11, '/edit/:nodeId', '[{"single": "public/views/settings/media-type/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (24, 'settings.dataTypes.edit', 'edit', 12, '/edit/:nodeId', '[{"single": "public/views/settings/data-type/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (25, 'settings.templates.edit', 'edit', 13, '/edit/:nodeId', '[{"single": "public/views/settings/template/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (26, 'settings.scripts.edit', 'edit', 14, '/edit/:name', '[{"single": "public/views/settings/script/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (27, 'settings.stylesheets.edit', 'edit', 15, '/edit/:name', '[{"single": "public/views/settings/stylesheet/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (28, 'members.edit', 'edit', 4, '/edit/:id', '[{"single": "public/views/members/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (29, 'members.new', 'new', 4, '/new', '[{"single": "public/views/members/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (30, 'members.memberTypes', 'memberTypes', 4, '/member-type', '[{"single": "public/views/members/member-type/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (31, 'members.memberTypes.edit', 'edit', 30, '/edit/:nodeId', '[{"single": "public/views/members/member-type/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (32, 'members.memberTypes.new', 'new', 30, '/new?type&parent', '[{"single": "public/views/members/member-type/new.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (33, 'members.memberGroups', 'memberTypes', 4, '/member-group', '[{"single": "public/views/members/member-group/index.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (34, 'members.memberGroups.edit', 'edit', 33, '/edit/:id', '[{"single": "public/views/members/member-group/edit.html"}]', false);
INSERT INTO route (id, path, name, parent_id, url, components, is_abstract) VALUES (35, 'members.memberGroups.new', 'new', 33, '/new?type&parent', '[{"single": "public/views/members/member-group/new.html"}]', false);


--
-- TOC entry 2380 (class 0 OID 0)
-- Dependencies: 192
-- Name: route_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('route_id_seq', 35, true);


--
-- TOC entry 2361 (class 0 OID 95184)
-- Dependencies: 193
-- Data for Name: template; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (1, 21, 'Collexy.Layout', false, '{30,34}', NULL);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (2, 22, 'Collexy.Home', false, '{32,33}', 21);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (3, 23, 'Collexy.Post', false, '{32,33}', 21);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (4, 24, 'Collexy.PostOverview', false, '{32}', 21);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (5, 25, 'Collexy.Page', false, '{32,33}', 21);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (6, 26, 'Collexy.Login', false, NULL, 21);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (7, 27, 'Collexy.Register', false, NULL, 21);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (10, 30, 'Collexy.TopNavigation', true, NULL, NULL);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (11, 31, 'Collexy.PostOverviewWidget', true, NULL, NULL);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (12, 32, 'Collexy.FeaturedPagesWidget', true, NULL, NULL);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (13, 33, 'Collexy.RecentPostsWidget', true, NULL, NULL);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (14, 34, 'Collexy.Social', true, NULL, NULL);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (16, 61, 'Collexy.CategoryListWidget', true, NULL, NULL);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (15, 60, 'Collexy.Category', false, '{}', 21);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (9, 29, 'Collexy.Unauthorized', false, NULL, NULL);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (8, 28, 'Collexy.404', false, NULL, 21);
INSERT INTO template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) VALUES (17, 68, 'Collexy.LoginWidget', true, NULL, NULL);


--
-- TOC entry 2381 (class 0 OID 0)
-- Dependencies: 194
-- Name: template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('template_id_seq', 17, true);


--
-- TOC entry 2363 (class 0 OID 95193)
-- Dependencies: 195
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids, permissions) VALUES (1, 'admin', 'Admin', 'Demo', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'admin@codeish.com', '2014-11-15 16:51:00.215', NULL, '2015-03-24 02:29:39.894', NULL, 1, 'SBZJXRY35B23EYX3WSM2IA6GTUOG7OIRYUPBV2MQCIGHV4ROZRIA', '{1}', NULL);


--
-- TOC entry 2364 (class 0 OID 95200)
-- Dependencies: 196
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO user_group (id, name, alias, permissions) VALUES (1, 'Administrator', 'administrator', '{node_create,node_delete,node_update,node_move,node_copy,node_public_access,node_permissions,node_send_to_publish,node_publish,node_browse,node_change_content_type,admin,content_all,content_create,content_delete,content_update,content_section,content_browse,media_all,media_create,media_delete,media_update,media_section,media_browse,users_all,users_create,users_delete,users_update,users_section,users_browse,user_types_all,user_types_create,user_types_delete,user_types_update,user_types_section,user_types_browse,user_groups_all,user_groups_create,user_groups_delete,user_groups_update,user_groups_section,user_groups_browse,members_all,members_create,members_delete,members_update,members_section,members_browse,member_types_all,member_types_create,member_types_delete,member_types_update,member_types_section,member_types_browse,member_groups_all,member_groups_create,member_groups_delete,member_groups_update,member_groups_section,member_groups_browse,templates_all,templates_create,templates_delete,templates_update,templates_section,templates_browse,scripts_all,scripts_create,scripts_delete,scripts_update,scripts_section,scripts_browse,stylesheets_all,stylesheets_create,stylesheets_delete,stylesheets_update,stylesheets_section,stylesheets_browse,settings_section,settings_all,node_sort,content_types_all,content_types_create,content_types_delete,content_types_update,content_types_section,content_types_browse,media_types_all,media_types_create,media_types_delete,media_types_update,media_types_section,media_types_browse,data_types_all,data_types_create,data_types_delete,data_types_update,data_types_section,data_types_browse}');
INSERT INTO user_group (id, name, alias, permissions) VALUES (2, 'Editor', 'editor', NULL);
INSERT INTO user_group (id, name, alias, permissions) VALUES (3, 'Writer', 'writer', NULL);


--
-- TOC entry 2382 (class 0 OID 0)
-- Dependencies: 197
-- Name: user_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_group_id_seq', 3, true);


--
-- TOC entry 2383 (class 0 OID 0)
-- Dependencies: 198
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 1, true);


-- Completed on 2015-03-24 12:22:06

--
-- PostgreSQL database dump complete
--

