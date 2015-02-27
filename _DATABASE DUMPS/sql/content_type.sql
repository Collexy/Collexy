-- Table: content_type

-- DROP TABLE content_type;

CREATE TABLE content_type
(
  id serial NOT NULL,
  node_id bigint NOT NULL,
  name character varying(255) NOT NULL,
  description character varying(255),
  icon character varying(255),
  thumbnail character varying(255),
  master_content_type integer,
  tabs jsonb[],
  meta jsonb,
  CONSTRAINT document_type_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE content_type
  OWNER TO postgres;

INSERT INTO content_type(
            node_id, name, description, icon, thumbnail, 
            tabs, meta)
    VALUES (3, 'Master', 'Some description', 'fa fa-folder-o', 'fa fa-folder-o', 
            ARRAY['{"name": "tab1","properties": [{"name": "prop1","order": 1,"help_text": "help text","description": "description","data_type": 1}]}'::jsonb,
            '{"name": "tab2","properties": [{"name": "prop2","order": 1,"help_text": "help text2","description": "description2","data_type": 1},{"name": "prop3","order": 2,"help_text": "help text3","description": "description3","data_type": 1}]}'::jsonb]
  , '{"allowed_types": [2,3]}');


INSERT INTO content_type(
            node_id, name, description, icon, thumbnail, master_content_type, meta)
    VALUES (4, 'Page', 'Page content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 1, '{"template_node_id":6}'::jsonb);

INSERT INTO content_type(
            node_id, name, description, icon, thumbnail, 
            tabs, meta)
    VALUES (9, 'Home', 'Home Some description', 'fa fa-folder-o', 'fa fa-folder-o', 
            ARRAY['{"name": "social","properties": [{"name": "facebook","order": 1,"help_text": "fb social help text","description": "fb social description","data_type": 1}, {"name": "google_plus","order": 2,"help_text": "g+ social help text","description": "g+ social description","data_type": 1}]}'::jsonb]
  , '{"allowed_types": [2,3]}');