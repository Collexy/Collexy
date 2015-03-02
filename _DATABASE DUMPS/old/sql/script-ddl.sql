CREATE DATABASE collexy
  WITH ENCODING='UTF8'
       CONNECTION LIMIT=-1;

CREATE EXTENSION ltree;

CREATE TABLE public."user"
(
   id bigserial NOT NULL, 
   username character varying NOT NULL, 
   first_name character varying, 
   last_name character varying, 
   password character(60) NOT NULL, 
   CONSTRAINT user_pkey PRIMARY KEY (id)
) 
WITH (
  OIDS = FALSE
)
;

CREATE TABLE node
(
  id bigserial NOT NULL,
  path ltree NOT NULL,
  name character varying(255),
  node_type smallint NOT NULL,
  created_by bigint NOT NULL,
  created_date timestamp without time zone NOT NULL DEFAULT now(),
  CONSTRAINT node_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

CREATE TABLE data_type
(
  id serial NOT NULL,
  node_id bigint NOT NULL,
  html text NOT NULL,
  CONSTRAINT data_type_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

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
  CONSTRAINT content_type_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

CREATE TABLE template
(
  id serial NOT NULL,
  node_id bigint NOT NULL,
  name character varying,
  html text,
  parent_template_node_id bigint,
  CONSTRAINT template_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE template
  OWNER TO postgres;

CREATE TABLE content
(
  id bigserial NOT NULL,
  node_id bigint NOT NULL,
  content_type_node_id bigint NOT NULL,
  meta jsonb,
  CONSTRAINT document_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

