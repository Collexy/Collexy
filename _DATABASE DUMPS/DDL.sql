INSERT INTO node VALUES (1, '1', 'root', 5, 1, '2014-10-22 16:51:00.215', NULL, NULL, NULL);
INSERT INTO node VALUES (2, '1.2', 'Text input', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (3, '1.3', 'Numeric input', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (4, '1.4', 'Textarea', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (5, '1.5', 'Radiobox', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (6, '1.6', 'Radiobox list', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (7, '1.7', 'Dropdown', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (8, '1.8', 'Dropdown multiple', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (9, '1.9', 'Checkbox', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (10, '1.10', 'Checkbox list', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (11, '1.11', 'Label', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (12, '1.12', 'Color picker', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (13, '1.13', 'Date picker', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (14, '1.14', 'Date picker with time', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (15, '1.15', 'Folder browser', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (16, '1.16', 'Upload', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (17, '1.17', 'Richtext editor', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (18, '1.18', 'True/false', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (19, '1.19', 'Domains', 11, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);

INSERT INTO node VALUES (20, '1.20', 'Member', 12, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);

INSERT INTO node VALUES (21, '1.21', 'Master', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (22, '1.21.22', 'Home', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node VALUES (23, '1.21.23', 'Post', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node VALUES (24, '1.21.24', 'Post Overview', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node VALUES (25, '1.21.25', 'Page', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node VALUES (26, '1.21.26', 'Login', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node VALUES (27, '1.21.27', 'Register', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node VALUES (28, '1.21.28', '404', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node VALUES (29, '1.21.29', 'Unauthorized', 3, 1, '2014-10-22 16:51:00.215', 21, NULL, NULL);
INSERT INTO node VALUES (30, '1.30', 'Top Navigation', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (31, '1.31', 'Post Overview Widget', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (32, '1.32', 'Featured Pages Widget', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (33, '1.33', 'Recent Posts Widget', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (34, '1.34', 'Social', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);

INSERT INTO template VALUES (1, 21, 'Collexy.Master', false, '{30,34}', NULL);
INSERT INTO template VALUES (2, 22, 'Collexy.Home', false, '{32,33}', 21);
INSERT INTO template VALUES (3, 23, 'Collexy.Post', false, '{32,33}', 21);
INSERT INTO template VALUES (4, 24, 'Collexy.PostOverview', false, '{32}', 21);
INSERT INTO template VALUES (5, 25, 'Collexy.Page', false, '{32,33}', 21);
INSERT INTO template VALUES (6, 26, 'Collexy.Login', false, NULL, 21);
INSERT INTO template VALUES (7, 27, 'Collexy.Register', false, NULL, 21);
INSERT INTO template VALUES (8, 28, 'Collexy.404', false, NULL, 21);
INSERT INTO template VALUES (9, 29, 'Collexy.Unauthorized', false, NULL, 21);
INSERT INTO template VALUES (10, 30, 'Collexy.TopNavigation', true, NULL, NULL);
INSERT INTO template VALUES (11, 31, 'Collexy.PostOverviewWidget', true, NULL, NULL);
INSERT INTO template VALUES (12, 32, 'Collexy.FeaturedPagesWidget', true, NULL, NULL);
INSERT INTO template VALUES (13, 33, 'Collexy.RecentPostsWidget', true, NULL, NULL);
INSERT INTO template VALUES (14, 34, 'Collexy.Social', true, NULL, NULL);


INSERT INTO data_type VALUES (1, 2, '<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">', 'Collexy.TextField');
INSERT INTO data_type VALUES (2, 3, '<input type="number" id="{{prop.name}}" ng-model="data.meta[prop.name]">', 'Collexy.NumberField');
INSERT INTO data_type VALUES (3, 4, '<textarea id="{{prop.name}}" ng-model="data.meta[prop.name]">', 'Collexy.Textarea');
INSERT INTO data_type VALUES (4, 5, '', 'Collexy.Radiobox');
INSERT INTO data_type VALUES (5, 6, '', 'Collexy.RadioboxList');
INSERT INTO data_type VALUES (6, 7, '', 'Collexy.Dropdown');
INSERT INTO data_type VALUES (7, 8, '', 'Collexy.DropdownMultiple');
INSERT INTO data_type VALUES (8, 9, '', 'Collexy.Checkbox');
INSERT INTO data_type VALUES (9, 10, '', 'Collexy.CheckboxList');
INSERT INTO data_type VALUES (10, 11, '', 'Collexy.Label');
INSERT INTO data_type VALUES (11, 12, '<colorpicker>The color picker data type is not implemented yet!</colorpicker>', 'Collexy.ColorPicker');
INSERT INTO data_type VALUES (12, 13, '', 'Collexy.DatePicker');
INSERT INTO data_type VALUES (13, 14, '', 'Collexy.DatePickerTime');
INSERT INTO data_type VALUES (14, 15, '<folderbrowser>This is an awesome folder browser (unimplemented datatype)</folderbrowser>', 'Collexy.FolderBrowser');
INSERT INTO data_type VALUES (15, 16, '<input type="file" file-input="test.files" multiple />
<button ng-click="upload()" type="button">Upload</button>
<li ng-repeat="file in test.files">{{file.name}}</li>
<!--<input type="file" onchange="angular.element(this).scope().filesChanged(this)" multiple />
<button ng-click="upload()">Upload</button>
<li ng-repeat="file in files">{{file.name}}</li>-->', 'Collexy.Upload');
INSERT INTO data_type VALUES (16, 17, '', 'Collexy.RichtextEditor');
INSERT INTO data_type VALUES (17, 18, '', 'Collexy.TrueFalse');
INSERT INTO data_type VALUES (18, 19, '<div>
	<input type="text"/> <button type="button">Add domain</button><br>
	<ul>
		<li ng-repeat="domain in data.meta[prop.name]">{{domain}}</li>
	</ul>
	<button type="button">Delete selected</button>
</div>', 'Collexy.Domains');

--INSERT INTO member_type VALUES (1, 20, 'Umbraco.Member', 'Default member type', 'fa fa-user fa-fw', NULL, NULL, '[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_node_id": 4}]}]');

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


INSERT INTO route VALUES (1, 'content', 'content', NULL, '/admin/content', '[{"single": "public/views/content/index.html"}]', false);
INSERT INTO route VALUES (2, 'media', 'media', NULL, '/admin/media', '[{"single": "public/views/media/index.html"}]', false);
INSERT INTO route VALUES (3, 'users', 'users', NULL, '/admin/users', '[{"single": "public/views/users/index.html"}]', false);
INSERT INTO route VALUES (4, 'members', 'members', NULL, '/admin/members', '[{"single": "public/views/members/index.html"}]', false);
INSERT INTO route VALUES (5, 'settings', 'settings', NULL, '/admin/settings', '[{"single": "public/views/settings/index.html"}]', true);
INSERT INTO route VALUES (6, 'content.new', 'new', 1, '/new?node_type&content_type_node_id&parent_id', '[{"single": "public/views/content/new.html"}]', false);
INSERT INTO route VALUES (7, 'content.edit', 'edit', 1, '/edit/:nodeId', '[{"single": "public/views/content/edit.html"}]', false);
INSERT INTO route VALUES (8, 'media.new', 'new', 2, '/new?node_type&content_type_node_id&parent_id', '[{"single": "public/views/media/new.html"}]', false);
INSERT INTO route VALUES (9, 'media.edit', 'edit', 2, '/edit/:nodeId', '[{"single": "public/views/media/edit.html"}]', false);
INSERT INTO route VALUES (10, 'settings.contentTypes', 'contentTypes', 5, '/content-type', '[{"single": "public/views/settings/content-type/index.html"}]', false);
INSERT INTO route VALUES (11, 'settings.mediaTypes', 'mediaTypes', 5, '/media-type', '[{"single": "public/views/settings/media-type/index.html"}]', false);
INSERT INTO route VALUES (12, 'settings.dataTypes', 'dataTypes', 5, '/data-type', '[{"single": "public/views/settings/data-type/index.html"}]', false);
INSERT INTO route VALUES (13, 'settings.templates', 'templates', 5, '/template', '[{"single": "public/views/settings/template/index.html"}]', false);
INSERT INTO route VALUES (14, 'settings.scripts', 'scripts', 5, '/script', '[{"single": "public/views/settings/script/index.html"}]', false);
INSERT INTO route VALUES (15, 'settings.stylesheets', 'stylesheets', 5, '/stylesheet', '[{"single": "public/views/settings/stylesheet/index.html"}]', false);
INSERT INTO route VALUES (16, 'settings.contentTypes.new', 'new', 10, '/new?type&parent', '[{"single": "public/views/settings/content-type/new.html"}]', false);
INSERT INTO route VALUES (17, 'settings.mediaTypes.new', 'new', 11, '/new?type&parent', '[{"single": "public/views/settings/media-type/new.html"}]', false);
INSERT INTO route VALUES (18, 'settings.dataTypes.new', 'new', 12, '/new', '[{"single": "public/views/settings/data-type/new.html"}]', false);
INSERT INTO route VALUES (19, 'settings.templates.new', 'new', 13, '/new?parent', '[{"single": "public/views/settings/template/new.html"}]', false);
INSERT INTO route VALUES (20, 'settings.scripts.new', 'new', 14, '/new?type&parent', '[{"single": "public/views/settings/script/new.html"}]', false);
INSERT INTO route VALUES (21, 'settings.stylesheets.new', 'new', 15, '/new?type&parent', '[{"single": "public/views/settings/stylesheet/new.html"}]', false);
INSERT INTO route VALUES (22, 'settings.contentTypes.edit', 'edit', 10, '/edit/:nodeId', '[{"single": "public/views/settings/content-type/edit.html"}]', false);
INSERT INTO route VALUES (23, 'settings.mediaTypes.edit', 'edit', 11, '/edit/:nodeId', '[{"single": "public/views/settings/media-type/edit.html"}]', false);
INSERT INTO route VALUES (24, 'settings.dataTypes.edit', 'edit', 12, '/edit/:nodeId', '[{"single": "public/views/settings/data-type/edit.html"}]', false);
INSERT INTO route VALUES (25, 'settings.templates.edit', 'edit', 13, '/edit/:nodeId', '[{"single": "public/views/settings/template/edit.html"}]', false);
INSERT INTO route VALUES (26, 'settings.scripts.edit', 'edit', 14, '/edit/:name', '[{"single": "public/views/settings/script/edit.html"}]', false);
INSERT INTO route VALUES (27, 'settings.stylesheets.edit', 'edit', 15, '/edit/:name', '[{"single": "public/views/settings/stylesheet/edit.html"}]', false);
INSERT INTO route VALUES (28, 'members.edit', 'edit', 4, '/edit/:id', '[{"single": "public/views/members/edit.html"}]', false);
INSERT INTO route VALUES (29, 'members.new', 'new', 4, '/new', '[{"single": "public/views/members/new.html"}]', false);
INSERT INTO route VALUES (30, 'members.memberTypes', 'memberTypes', 4, '/member-type', '[{"single": "public/views/members/member-type/index.html"}]', false);
INSERT INTO route VALUES (31, 'members.memberTypes.edit', 'edit', 30, '/edit/:nodeId', '[{"single": "public/views/members/member-type/edit.html"}]', false);
INSERT INTO route VALUES (32, 'members.memberTypes.new', 'new', 30, '/new?type&parent', '[{"single": "public/views/members/member-type/new.html"}]', false);
INSERT INTO route VALUES (33, 'members.memberGroups', 'memberTypes', 4, '/member-group', '[{"single": "public/views/members/member-group/index.html"}]', false);
INSERT INTO route VALUES (34, 'members.memberGroups.edit', 'edit', 33, '/edit/:id', '[{"single": "public/views/members/member-group/edit.html"}]', false);
INSERT INTO route VALUES (35, 'members.memberGroups.new', 'new', 33, '/new?type&parent', '[{"single": "public/views/members/member-group/new.html"}]', false);

INSERT INTO menu_link VALUES (1, '1', 'Content', NULL, 1, 'fa fa-newspaper-o fa-fw', NULL, 1, 'main', '{content_section}');
INSERT INTO menu_link VALUES (2, '2', 'Media', NULL, 2, 'fa fa-file-image-o fa-fw', NULL, 1, 'main', '{media_section}');
INSERT INTO menu_link VALUES (3, '3', 'Users', NULL, 3, 'fa fa-user fa-fw', NULL, 1, 'main', '{users_section}');
INSERT INTO menu_link VALUES (4, '4', 'Members', NULL, 4, 'fa fa-users fa-fw', NULL, 1, 'main', '{members_section}');
INSERT INTO menu_link VALUES (5, '5', 'Settings', NULL, 5, 'fa fa-gear fa-fw', NULL, 1, 'main', '{settings_section}');
INSERT INTO menu_link VALUES (6, '5.6', 'Content Types', 5, 10, 'fa fa-newspaper-o fa-fw', NULL, 1, 'main', '{content_types_section}');
INSERT INTO menu_link VALUES (7, '5.7', 'Media Types', 5, 11, 'fa fa-files-o fa-fw', NULL, 1, 'main', '{media_types_section}');
INSERT INTO menu_link VALUES (8, '5.8', 'Data Types', 5, 12, 'fa fa-check-square-o fa-fw', NULL, 1, 'main', '{data_types_section}');
INSERT INTO menu_link VALUES (9, '5.9', 'Templates', 5, 13, 'fa fa-eye fa-fw', NULL, 1, 'main', '{templates_section}');
INSERT INTO menu_link VALUES (10, '6.10', 'Scripts', 5, 14, 'fa fa-file-code-o fa-fw', NULL, 1, 'main', '{scripts_section}');
INSERT INTO menu_link VALUES (11, '6.11', 'Stylesheets', 5, 15, 'fa fa-desktop fa-fw', NULL, 1, 'main', '{stylesheets_section}');
INSERT INTO menu_link VALUES (12, '5.12', 'Member Types', 4, 30, 'fa fa-smile-o fa-fw', NULL, 1, 'main', '{member_types_section}');
INSERT INTO menu_link VALUES (13, '5.13', 'Member Groups', 4, 33, 'fa fa-smile-o fa-fw', NULL, 1, 'main', '{member_groups_section}');


INSERT INTO user_group VALUES (2, 'Editor', 'editor', NULL);
INSERT INTO user_group VALUES (3, 'Writer', 'writer', NULL);
INSERT INTO user_group VALUES (1, 'Administrator', 'administrator', '{node_create,node_delete,node_update,node_move,node_copy,node_public_access,node_permissions,node_send_to_publish,node_publish,node_browse,node_change_content_type,admin,content_all,content_create,content_delete,content_update,content_section,content_browse,media_all,media_create,media_delete,media_update,media_section,media_browse,users_all,users_create,users_delete,users_update,users_section,users_browse,user_types_all,user_types_create,user_types_delete,user_types_update,user_types_section,user_types_browse,user_groups_all,user_groups_create,user_groups_delete,user_groups_update,user_groups_section,user_groups_browse,members_all,members_create,members_delete,members_update,members_section,members_browse,member_types_all,member_types_create,member_types_delete,member_types_update,member_types_section,member_types_browse,member_groups_all,member_groups_create,member_groups_delete,member_groups_update,member_groups_section,member_groups_browse,templates_all,templates_create,templates_delete,templates_update,templates_section,templates_browse,scripts_all,scripts_create,scripts_delete,scripts_update,scripts_section,scripts_browse,stylesheets_all,stylesheets_create,stylesheets_delete,stylesheets_update,stylesheets_section,stylesheets_browse,settings_section,settings_all,node_sort,content_types_all,content_types_create,content_types_delete,content_types_update,content_types_section,content_types_browse,media_types_all,media_types_create,media_types_delete,media_types_update,media_types_section,media_types_browse,data_types_all,data_types_create,data_types_delete,data_types_update,data_types_section,data_types_browse}');

INSERT INTO "user" VALUES (1, 'admin', 'Admin', 'Demo', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'demo@codeish.com', '2014-11-15 16:51:00.215', NULL, '2015-02-27 15:12:12.285', NULL, 1, 'ZMLZCCH7WXTLDCMOXAHV3PMRB3NPR5A33PQUSFLWS2QA5CGH5YPQ', '{1}', NULL);

INSERT INTO member_group VALUES (1, 'authenticated_member', 'All logged in members');

INSERT INTO member VALUES (1, 'default_member', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'default_member@mail.com', '{"comments": "default user comments"}', '2015-01-22 14:25:38.904', NULL, '2015-02-19 23:46:00.495', NULL, 1, 'GIWES3RHMY5RKC7OZPOQTF5FQFWX32D5VLV3CAKT4HGKP5LZIENA', 20, '{1}');
INSERT INTO member_type VALUES (1, 20, 'Collexy.Member', 'Default member type', 'fa fa-user fa-fw', 1, NULL, '[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_node_id": 4}]}]');

INSERT INTO node VALUES (35, '1.35', 'Master', 4, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (36, '1.35.36', 'Home', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node VALUES (37, '1.35.37', 'Post', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node VALUES (38, '1.35.38', 'Post Overview', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node VALUES (39, '1.35.39', 'Page', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);

INSERT INTO node VALUES (40, '1.40', 'Folder', 7, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (41, '1.41', 'Image', 7, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);


INSERT INTO content_type VALUES (1, 35, 'Collexy.Master', 'Some description', 'fa fa-folder-o', 'fa fa-folder-o', NULL, NULL, '[{"name": "Content", "properties": [{"name": "page_title", "order": 1, "data_type_node_id": 2, "help_text": "help text", "description": "The page title overrides the name the page has been given."}]}, {"name": "Properties", "properties": [{"name": "prop2", "order": 1, "data_type_node_id": 2, "help_text": "help text2", "description": "description2"}, {"name": "prop3", "order": 2, "data_type_node_id": 2, "help_text": "help text3", "description": "description3"}]}]');
INSERT INTO content_type VALUES (2, 36, 'Collexy.Home', 'Home Some description', 'fa fa-folder-o', 'fa fa-folder-o', 35, '{"template_node_id": 22, "allowed_templates_node_id": [22], "allowed_content_types_node_id": [37,38,39]}', '[{"name":"Content","properties":[{"name":"site_name","order":2,"data_type_node_id":2,"help_text":"help text","description":"Site name goes here."},{"name":"site_tagline","order":3,"data_type_node_id":2,"help_text":"help text","description":"Site tagline goes here."},{"name":"copyright","order":4,"data_type_node_id":2,"help_text":"help text","description":"Copyright here."},{"name":"domains","order":5,"data_type_node_id":19,"help_text":"help text","description":"Domains goes here."}]},{"name":"Social","properties":[{"name":"facebook","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your facebook link here."},{"name":"twitter","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your twitter link here."},{"name":"linkedin","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your linkedin link here."}]}]');
INSERT INTO content_type VALUES (3, 37, 'Collexy.Post', 'Post content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 35, '{"template_node_id": 23, "allowed_templates_node_id": [23], "allowed_content_types_node_id": [37]}', '[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":4,"help_text":"Help text for page contentent","description":"Page content description"}]}]');
INSERT INTO content_type VALUES (4, 38, 'Collexy.PostOverview', 'Post overview content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 35, '{"template_node_id": 24, "allowed_templates_node_id": [24], "allowed_content_types_node_id": [38]}', '[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":4,"help_text":"Help text for page contentent","description":"Page content description"}]}]');
INSERT INTO content_type VALUES (5, 39, 'Collexy.Page', 'Page content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 35, '{"template_node_id": 25, "allowed_templates_node_id": [25], "allowed_content_types_node_id": [39]}', '[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":4,"help_text":"Help text for page contentent","description":"Page content description"}]}]');


INSERT INTO content_type VALUES (6, 40, 'Collexy.Folder', 'Folder media type description1', 'mt-icon1', 'mt-thumbnail1', NULL, '{"allowed_content_types_node_id": [40,41]}', '[{"name":"Folder","properties":[{"name":"folder_browser","order":1,"data_type_node_id":15,"help_text":"prop help text","description":"prop description"},{"name":"path","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties"}]');
INSERT INTO content_type VALUES (7, 41, 'Collexy.Image', 'Image content type description', 'fa fa-folder-o', 'fa fa-folder-o', NULL, 'null', '[{"name":"Image","properties":[{"name":"path","order":1,"data_type_node_id":2,"help_text":"help text","description":"URL goes here."},{"name":"title","order":2,"data_type_node_id":2,"help_text":"help text","description":"The title entered here can override the above one."},{"name":"caption","order":3,"data_type_node_id":4,"help_text":"help text","description":"Caption goes here."},{"name":"alt","order":4,"data_type_node_id":4,"help_text":"help text","description":"Alt goes here."},{"name":"description","order":5,"data_type_node_id":4,"help_text":"help text","description":"Description goes here."},{"name":"file_upload","order":1,"data_type_node_id":16,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties","properties":[{"name":"temporary property","order":1,"data_type_node_id":2,"help_text":"help text","description":"Temporary description goes here."}]}]');

INSERT INTO node VALUES (42, '1.42', 'Home', 1, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (43, '1.42.43', 'Sample Page', 1, 1, '2014-10-22 16:51:00.215', 42, '[{"id": 2, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]', NULL);
INSERT INTO node VALUES (44, '1.42.43.44', 'Child Page Level 1', 1, 1, '2014-10-26 23:19:44.735', 43, NULL, '[{"id": 1, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]');
INSERT INTO node VALUES (45, '1.42.43.44.45', 'Child Page Level 2', 1, 1, '2014-10-26 23:19:44.735', 44, NULL, NULL);

INSERT INTO content VALUES (1, 42, 36, '{"prop2": "Home page prop 2", "domains": ["localhost:8080", "localhost:8080/test"], "facebook": "facebook.com/home", "copyright": "&copy; 2014 codeish.com", "site_name": "Collexy cms test site", "page_title": "Home page title", "site_tagline": "Test site tagline", "template_node_id": 22}', NULL);

INSERT INTO content VALUES (2, 43, 39, '{"prop2": "prop2a", "prop3": "sample page prop 3", "page_title": "Sample page title", "page_content": "Sample page content goes here", "template_node_id": 25}', NULL);
INSERT INTO content VALUES (3, 44, 39, '{"prop3": "sample child page level 1 page prop 3", "page_title": "Child page level 1 title", "page_content": "Sample page - child page level 1 content goes here", "template_node_id": 25}', NULL);
INSERT INTO content VALUES (4, 45, 39, '{"prop3": "sample child page level 2 page prop 3", "page_title": "Child page level 2 title", "page_content": "Sample page - child page level 2 content goes here1", "template_node_id": 25}', '{"groups": [1], "members": [1]}');

INSERT INTO node VALUES (46, '1.42.46', 'Posts', 1, 1, '2014-10-22 16:51:00.215', 42, NULL, NULL);
INSERT INTO node VALUES (47, '1.42.46.47', 'Hello World', 1, 1, '2014-10-22 16:51:00.215', 46, NULL, NULL);

INSERT INTO content VALUES (5, 46, 38, '{"prop2": "prop2a", "prop3": "Posts overview prop 3", "page_title": "Posts", "page_content": "Sample page content goes here", "template_node_id": 24}', NULL);
INSERT INTO content VALUES (6, 47, 37, '{"prop2": "prop2a", "prop3": "Hello world prop 3", "page_title": "Hello World", "page_content": "Welcome to Collexy. This is your first post. Edit or delete it, then start blogging", "template_node_id": 23}', NULL);

INSERT INTO node VALUES (48, '1.48', 'gopher.jpg', 2, 1, '2014-10-28 15:50:47.303', 1, NULL, NULL);
INSERT INTO node VALUES (49, '1.49', '2014', 2, 1, '2014-12-02 01:42:09.979', 1, NULL, NULL);
INSERT INTO node VALUES (50, '1.49.50', '12', 2, 1, '2014-12-05 16:18:29.762', 49, NULL, NULL);
INSERT INTO node VALUES (51, '1.49.50.51', 'cat-prays.jpg', 2, 1, '2014-12-06 13:07:08.943', 50, NULL, NULL);
INSERT INTO node VALUES (52, '1.49.50.52', 'sleeping-kitten.jpg', 2, 1, '2014-12-06 14:28:52.117', 50, NULL, NULL);

INSERT INTO content VALUES (7, 48, 41, '{"alt": "Gopher image alt text1", "url": "/media/2014/10/gopher.jpg", "path": "media\\gopher.jpg", "caption": "This is the caption of the gopher image1", "description": "Gopher image description1", "temporary property": "lol"}', NULL);
INSERT INTO content VALUES (8, 49, 40, '{"path": "media\\2014"}', NULL);
INSERT INTO content VALUES (9, 50, 40, '{"path": "media\\2014\\12"}', NULL);
INSERT INTO content VALUES (10, 51, 41, '{"alt": "sleeping-kitten.jpg", "path": "media\\2014\\12\\sleeping-kitten.jpg", "title": "sleeping-kitten.jpg", "caption": "sleeping-kitten.jpg", "description": "sleeping-kitten.jpg"}', NULL);
INSERT INTO content VALUES (11, 52, 41, '{"alt": "cat-prays.jpg", "path": "media\\2014\\12\\cat-prays.jpg", "title": "cat-prays.jpg", "caption": "cat-prays.jpg", "description": "cat-prays.jpg"}', NULL);
