package globals

var DbCreateScriptDML string = `--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-03-02 00:27:45

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- TOC entry 199 (class 3079 OID 11855)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2346 (class 0 OID 0)
-- Dependencies: 199
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 201 (class 3079 OID 82512)
-- Name: citext; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- TOC entry 2347 (class 0 OID 0)
-- Dependencies: 201
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- TOC entry 200 (class 3079 OID 82596)
-- Name: ltree; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;


--
-- TOC entry 2348 (class 0 OID 0)
-- Dependencies: 200
-- Name: EXTENSION ltree; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';


SET search_path = public, pg_catalog;

--
-- TOC entry 321 (class 1255 OID 82771)
-- Name: json_merge(json, json); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION json_merge(data json, merge_data json) RETURNS json
    LANGUAGE sql
    AS $$
SELECT ('{'||string_agg(to_json(key)||':'||value, ',')||'}')::json
FROM (
WITH to_merge AS (
SELECT * FROM json_each(merge_data)
)
SELECT *
FROM json_each(data)
WHERE key NOT IN (SELECT key FROM to_merge)
UNION ALL
SELECT * FROM to_merge
) t;
$$;


ALTER FUNCTION public.json_merge(data json, merge_data json) OWNER TO postgres;

--
-- TOC entry 322 (class 1255 OID 82772)
-- Name: json_object_update_key(json, text, anyelement); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION json_object_update_key(json json, key_to_set text, value_to_set anyelement) RETURNS json
    LANGUAGE sql IMMUTABLE STRICT
    AS $$
SELECT CASE
  WHEN ("json" -> "key_to_set") IS NULL THEN "json"
  ELSE COALESCE(
    (SELECT ('{' || string_agg(to_json("key") || ':' || "value", ',') || '}')
       FROM (SELECT *
               FROM json_each("json")
              WHERE "key" <> "key_to_set"
              UNION ALL
             SELECT "key_to_set", to_json("value_to_set")) AS "fields"),
    '{}'
  )::json
END
$$;


ALTER FUNCTION public.json_object_update_key(json json, key_to_set text, value_to_set anyelement) OWNER TO postgres;

--
-- TOC entry 323 (class 1255 OID 82773)
-- Name: populate(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION populate() RETURNS record
    LANGUAGE plpgsql
    AS $$
DECLARE
    -- declarations
    ret RECORD;
BEGIN
    SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
    content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
    ct.id as ct_id, ct.node_id as ct_node_id, ct.master_content_type_node_id as ct_master_content_type_node_id, ct.name as ct_name,
    ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.tabs as ct_tabs, ct.meta as ct_meta,
    ctm.id as ctm_id, ctm.node_id as ctm_node_id, ctm.master_content_type_node_id as ctm_master_content_type_node_id, ctm.name as ctm_name,
    ctm.description as ctm_description, ctm.icon as ctm_icon, ctm.thumbnail as ctm_thumbnail, ctm.tabs as ctm_tabs, ctm.meta as ctm_meta
    FROM node
    JOIN content
    ON content.node_id = node.id
    JOIN content_type as ct
    ON ct.node_id = content.content_type_node_id
    JOIN content_type as ctm
    ON ctm.node_id = ct.master_content_type_node_id
    WHERE node.id = 10
    INTO ret;
    RETURN ret;
END;
$$;


ALTER FUNCTION public.populate() OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 172 (class 1259 OID 82782)
-- Name: content; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE content (
    id bigint NOT NULL,
    node_id bigint NOT NULL,
    content_type_node_id bigint NOT NULL,
    meta jsonb,
    public_access jsonb
);


ALTER TABLE content OWNER TO postgres;

--
-- TOC entry 173 (class 1259 OID 82788)
-- Name: content_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE content_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE content_id_seq OWNER TO postgres;

--
-- TOC entry 2349 (class 0 OID 0)
-- Dependencies: 173
-- Name: content_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_id_seq OWNED BY content.id;


--
-- TOC entry 174 (class 1259 OID 82790)
-- Name: content_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE content_type (
    id integer NOT NULL,
    node_id bigint NOT NULL,
    alias character varying(255) NOT NULL,
    description character varying(255),
    icon character varying(255),
    thumbnail character varying(255),
    parent_content_type_node_id integer,
    meta jsonb,
    tabs json
);


ALTER TABLE content_type OWNER TO postgres;

--
-- TOC entry 175 (class 1259 OID 82796)
-- Name: content_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE content_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE content_type_id_seq OWNER TO postgres;

--
-- TOC entry 2350 (class 0 OID 0)
-- Dependencies: 175
-- Name: content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_type_id_seq OWNED BY content_type.id;


--
-- TOC entry 176 (class 1259 OID 82798)
-- Name: data_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE data_type (
    id integer NOT NULL,
    node_id bigint NOT NULL,
    html text NOT NULL,
    alias character varying
);


ALTER TABLE data_type OWNER TO postgres;

--
-- TOC entry 177 (class 1259 OID 82804)
-- Name: data_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE data_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE data_type_id_seq OWNER TO postgres;

--
-- TOC entry 2351 (class 0 OID 0)
-- Dependencies: 177
-- Name: data_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_type_id_seq OWNED BY data_type.id;


--
-- TOC entry 178 (class 1259 OID 82806)
-- Name: domain; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE domain (
    id integer NOT NULL,
    node_id bigint,
    name character varying
);


ALTER TABLE domain OWNER TO postgres;

--
-- TOC entry 179 (class 1259 OID 82812)
-- Name: domain_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE domain_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE domain_id_seq OWNER TO postgres;

--
-- TOC entry 2352 (class 0 OID 0)
-- Dependencies: 179
-- Name: domain_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE domain_id_seq OWNED BY domain.id;


--
-- TOC entry 180 (class 1259 OID 82814)
-- Name: member; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE member (
    id bigint NOT NULL,
    username character varying NOT NULL,
    password character(60),
    email character varying,
    meta jsonb,
    created_date timestamp without time zone DEFAULT now() NOT NULL,
    updated_date timestamp without time zone,
    login_date timestamp without time zone,
    accessed_date timestamp without time zone,
    status smallint DEFAULT 1 NOT NULL,
    sid character varying,
    member_type_node_id bigint,
    group_ids integer[]
);


ALTER TABLE member OWNER TO postgres;

--
-- TOC entry 181 (class 1259 OID 82822)
-- Name: member_group; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE member_group (
    id integer NOT NULL,
    name character varying,
    description text
);


ALTER TABLE member_group OWNER TO postgres;

--
-- TOC entry 182 (class 1259 OID 82828)
-- Name: member_group_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE member_group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE member_group_id_seq OWNER TO postgres;

--
-- TOC entry 2353 (class 0 OID 0)
-- Dependencies: 182
-- Name: member_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_group_id_seq OWNED BY member_group.id;


--
-- TOC entry 183 (class 1259 OID 82830)
-- Name: member_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE member_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE member_id_seq OWNER TO postgres;

--
-- TOC entry 2354 (class 0 OID 0)
-- Dependencies: 183
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_id_seq OWNED BY member.id;


--
-- TOC entry 184 (class 1259 OID 82832)
-- Name: member_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE member_type (
    id bigint NOT NULL,
    node_id bigint,
    alias character varying,
    description character varying,
    icon character varying,
    parent_member_type_node_id bigint,
    meta jsonb,
    tabs jsonb
);


ALTER TABLE member_type OWNER TO postgres;

--
-- TOC entry 185 (class 1259 OID 82838)
-- Name: member_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE member_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE member_type_id_seq OWNER TO postgres;

--
-- TOC entry 2355 (class 0 OID 0)
-- Dependencies: 185
-- Name: member_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_type_id_seq OWNED BY member_type.id;


--
-- TOC entry 186 (class 1259 OID 82840)
-- Name: menu_link; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE menu_link (
    id bigint NOT NULL,
    path ltree,
    name character varying,
    parent_id bigint,
    route_id bigint,
    icon character varying,
    atts jsonb,
    type smallint,
    menu character varying,
    permissions text[]
);


ALTER TABLE menu_link OWNER TO postgres;

--
-- TOC entry 187 (class 1259 OID 82846)
-- Name: menu_link_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE menu_link_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE menu_link_id_seq OWNER TO postgres;

--
-- TOC entry 2356 (class 0 OID 0)
-- Dependencies: 187
-- Name: menu_link_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE menu_link_id_seq OWNED BY menu_link.id;


--
-- TOC entry 188 (class 1259 OID 82848)
-- Name: node; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE node (
    id bigint NOT NULL,
    path ltree,
    name character varying(255),
    node_type smallint NOT NULL,
    created_by bigint NOT NULL,
    created_date timestamp without time zone DEFAULT now() NOT NULL,
    parent_id bigint,
    user_permissions jsonb,
    user_group_permissions jsonb
);


ALTER TABLE node OWNER TO postgres;

--
-- TOC entry 189 (class 1259 OID 82855)
-- Name: node_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE node_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE node_id_seq OWNER TO postgres;

--
-- TOC entry 2357 (class 0 OID 0)
-- Dependencies: 189
-- Name: node_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE node_id_seq OWNED BY node.id;


--
-- TOC entry 190 (class 1259 OID 82857)
-- Name: permission; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE permission (
    name character varying NOT NULL
);


ALTER TABLE permission OWNER TO postgres;

--
-- TOC entry 191 (class 1259 OID 82863)
-- Name: route; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE route (
    id bigint NOT NULL,
    path ltree,
    name character varying,
    parent_id bigint,
    url character varying,
    components jsonb,
    is_abstract boolean
);


ALTER TABLE route OWNER TO postgres;

--
-- TOC entry 2358 (class 0 OID 0)
-- Dependencies: 191
-- Name: COLUMN route.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN route.path IS '
';


--
-- TOC entry 192 (class 1259 OID 82869)
-- Name: route_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE route_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE route_id_seq OWNER TO postgres;

--
-- TOC entry 2359 (class 0 OID 0)
-- Dependencies: 192
-- Name: route_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE route_id_seq OWNED BY route.id;


--
-- TOC entry 193 (class 1259 OID 82871)
-- Name: template; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE template (
    id integer NOT NULL,
    node_id bigint NOT NULL,
    alias character varying,
    is_partial boolean DEFAULT false NOT NULL,
    partial_template_node_ids bigint[],
    parent_template_node_id bigint
);


ALTER TABLE template OWNER TO postgres;

--
-- TOC entry 194 (class 1259 OID 82878)
-- Name: template_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE template_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE template_id_seq OWNER TO postgres;

--
-- TOC entry 2360 (class 0 OID 0)
-- Dependencies: 194
-- Name: template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE template_id_seq OWNED BY template.id;


--
-- TOC entry 195 (class 1259 OID 82880)
-- Name: user; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE "user" (
    id bigint NOT NULL,
    username character varying NOT NULL,
    first_name character varying,
    last_name character varying,
    password character(60) NOT NULL,
    email character varying,
    created_date timestamp without time zone,
    updated_date timestamp without time zone,
    login_date timestamp without time zone,
    accessed_date timestamp without time zone,
    status smallint DEFAULT 1 NOT NULL,
    sid character varying,
    user_group_ids integer[],
    permissions text[]
);


ALTER TABLE "user" OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 82887)
-- Name: user_group; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE user_group (
    id integer NOT NULL,
    name character varying,
    alias character varying,
    permissions text[]
);


ALTER TABLE user_group OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 82893)
-- Name: user_group_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE user_group_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_group_id_seq OWNER TO postgres;

--
-- TOC entry 2361 (class 0 OID 0)
-- Dependencies: 197
-- Name: user_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_group_id_seq OWNED BY user_group.id;


--
-- TOC entry 198 (class 1259 OID 82895)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE user_id_seq OWNER TO postgres;

--
-- TOC entry 2362 (class 0 OID 0)
-- Dependencies: 198
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- TOC entry 2194 (class 2604 OID 82898)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content ALTER COLUMN id SET DEFAULT nextval('content_id_seq'::regclass);


--
-- TOC entry 2195 (class 2604 OID 82899)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content_type ALTER COLUMN id SET DEFAULT nextval('content_type_id_seq'::regclass);


--
-- TOC entry 2196 (class 2604 OID 82900)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_type ALTER COLUMN id SET DEFAULT nextval('data_type_id_seq'::regclass);


--
-- TOC entry 2197 (class 2604 OID 82901)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY domain ALTER COLUMN id SET DEFAULT nextval('domain_id_seq'::regclass);


--
-- TOC entry 2200 (class 2604 OID 82902)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member ALTER COLUMN id SET DEFAULT nextval('member_id_seq'::regclass);


--
-- TOC entry 2201 (class 2604 OID 82903)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_group ALTER COLUMN id SET DEFAULT nextval('member_group_id_seq'::regclass);


--
-- TOC entry 2202 (class 2604 OID 82904)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_type ALTER COLUMN id SET DEFAULT nextval('member_type_id_seq'::regclass);


--
-- TOC entry 2203 (class 2604 OID 82905)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY menu_link ALTER COLUMN id SET DEFAULT nextval('menu_link_id_seq'::regclass);


--
-- TOC entry 2205 (class 2604 OID 82906)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY node ALTER COLUMN id SET DEFAULT nextval('node_id_seq'::regclass);


--
-- TOC entry 2206 (class 2604 OID 82907)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY route ALTER COLUMN id SET DEFAULT nextval('route_id_seq'::regclass);


--
-- TOC entry 2208 (class 2604 OID 82908)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY template ALTER COLUMN id SET DEFAULT nextval('template_id_seq'::regclass);


--
-- TOC entry 2210 (class 2604 OID 82909)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- TOC entry 2211 (class 2604 OID 82910)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_group ALTER COLUMN id SET DEFAULT nextval('user_group_id_seq'::regclass);


--
-- TOC entry 2216 (class 2606 OID 82912)
-- Name: content_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY content_type
    ADD CONSTRAINT content_type_pkey PRIMARY KEY (id);


--
-- TOC entry 2218 (class 2606 OID 82914)
-- Name: data_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY data_type
    ADD CONSTRAINT data_type_pkey PRIMARY KEY (id);


--
-- TOC entry 2213 (class 2606 OID 82916)
-- Name: document_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY content
    ADD CONSTRAINT document_pkey PRIMARY KEY (id);


--
-- TOC entry 2220 (class 2606 OID 82918)
-- Name: node_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY node
    ADD CONSTRAINT node_pkey PRIMARY KEY (id);


--
-- TOC entry 2222 (class 2606 OID 82920)
-- Name: permission_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY permission
    ADD CONSTRAINT permission_name_key UNIQUE (name);


--
-- TOC entry 2225 (class 2606 OID 82922)
-- Name: template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY template
    ADD CONSTRAINT template_pkey PRIMARY KEY (id);


--
-- TOC entry 2227 (class 2606 OID 82924)
-- Name: user_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- TOC entry 2229 (class 2606 OID 82926)
-- Name: user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 2214 (class 1259 OID 82927)
-- Name: idxgin; Type: INDEX; Schema: public; Owner: postgres; Tablespace: 
--

CREATE INDEX idxgin ON content USING gin (meta);


--
-- TOC entry 2223 (class 1259 OID 82928)
-- Name: template_partial_template_node_ids_idx; Type: INDEX; Schema: public; Owner: postgres; Tablespace: 
--

CREATE INDEX template_partial_template_node_ids_idx ON template USING gin (partial_template_node_ids);


--
-- TOC entry 2345 (class 0 OID 0)
-- Dependencies: 6
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2015-03-02 00:27:46

--
-- PostgreSQL database dump complete
--`

var DbCreateScriptDDL string = `--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-03-03 12:27:46

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

INSERT INTO content VALUES (1, 42, 36, '{"is_featured": false, "hide_in_nav": false, "domains": ["localhost:8080", "localhost:8080/test"], "facebook_link": "facebook.com/home", "copyright": "&copy; 2014 codeish.com", "site_name": "%s", "title": "Home page title", "site_tagline": "Test site tagline", "template_node_id": 22, "hide_banner": false, "banner_header": "Banner header goes here", "banner_subheader": "Banner subheader goes here", "banner_link": "http://somelink.test", "banner_link_text": "Click Here!", "banner_background_image": "/media/Sample Images/TXT/banner.jpg", "about_title": "About title here", "about_text": "About text goes here"}', NULL);
INSERT INTO content VALUES (2, 43, 39, '{"is_featured": true, "hide_in_nav": false, "title": "Welcome", "content": "Welcome content goes here", "template_node_id": 25, "image": "/media/Sample Images/TXT/pic01.jpg"}', NULL);
INSERT INTO content VALUES (3, 44, 39, '{"is_featured": true, "hide_in_nav": false, "title": "Getting Started", "content": "Getting Started content goes here", "template_node_id": 25, "image": "/media/Sample Images/TXT/pic02.jpg"}', NULL);
INSERT INTO content VALUES (4, 45, 39, '{"is_featured": true, "hide_in_nav": false, "title": "Documentation", "content": "Documentation content goes here1", "template_node_id": 25, "image": "/media/Sample Images/TXT/pic03.jpg"}', '{"groups": [1], "members": [1]}');
INSERT INTO content VALUES (5, 46, 39, '{"is_featured": true, "hide_in_nav": false, "title": "Get Involved", "content": "Get Involved content goes here", "template_node_id": 25, "image": "/media/Sample Images/TXT/pic04.jpg"}', NULL);
INSERT INTO content VALUES (6, 47, 38, '{"is_featured": true, "hide_in_nav": false, "title": "Posts", "template_node_id": 24}', NULL);
INSERT INTO content VALUES (7, 48, 37, '{"is_featured": true, "hide_in_nav": false, "title": "Hello World", "content": "Welcome to Collexy. This is your first post. Edit or delete it, then start blogging", "template_node_id": 23, "image": "/media/Sample Images/TXT/pic05.jpg", "sub_header": "Subheader for Hello World"}', NULL);
INSERT INTO content VALUES (8, 49, 37, '{"is_featured": true, "hide_in_nav": false, "title": "TXT Starter Kit For Collexy Released", "content": "The collexy TXT starter kit is just awesome!", "template_node_id": 23}', NULL);
INSERT INTO content VALUES (9, 50, 37, '{"is_featured": true, "hide_in_nav": false, "title": "You Need To Read This", "content": "See - you really needed to read this post!", "template_node_id": 23}', NULL);
INSERT INTO content VALUES (10, 51, 37, '{"is_featured": true, "hide_in_nav": false, "title": "Amazing Post", "content": "What an amazing post.", "template_node_id": 23}', NULL);

INSERT INTO content VALUES (11, 52, 40, '{"path": "media\\Sample Images"}', NULL);
INSERT INTO content VALUES (12, 53, 40, '{"path": "media\\Sample Images\\TXT"}', NULL);
INSERT INTO content VALUES (13, 54, 41, '{"alt": "pic01.jpg", "path": "media\\Sample Images\\TXT\\pic01.jpg", "title": "pic01.jpg", "caption": "pic01.jpg", "description": "pic01.jpg"}', NULL);
INSERT INTO content VALUES (14, 55, 41, '{"alt": "pic02.jpg", "path": "media\\Sample Images\\TXT\\pic02.jpg", "title": "pic02.jpg", "caption": "pic02.jpg", "description": "pic02.jpg"}', NULL);
INSERT INTO content VALUES (15, 56, 41, '{"alt": "pic03.jpg", "path": "media\\Sample Images\\TXT\\pic03.jpg", "title": "pic03.jpg", "caption": "pic03.jpg", "description": "pic03.jpg"}', NULL);
INSERT INTO content VALUES (16, 57, 41, '{"alt": "pic04.jpg", "path": "media\\Sample Images\\TXT\\pic04.jpg", "title": "pic04.jpg", "caption": "pic04.jpg", "description": "pic04.jpg"}', NULL);
INSERT INTO content VALUES (17, 58, 41, '{"alt": "pic05.jpg", "path": "media\\Sample Images\\TXT\\pic05.jpg", "title": "pic05.jpg", "caption": "pic05.jpg", "description": "pic05.jpg"}', NULL);
INSERT INTO content VALUES (18, 59, 41, '{"alt": "banner.jpg", "path": "media\\Sample Images\\TXT\\banner.jpg", "title": "banner.jpg", "caption": "banner.jpg", "description": "banner.jpg"}', NULL);

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

INSERT INTO content_type VALUES (1, 35, 'Collexy.Master', 'Some description', 'fa fa-folder-o', 'fa fa-folder-o', NULL, NULL, '[{"name":"Content","properties":[{"name":"title","order":1,"data_type_node_id":2,"help_text":"help text","description":"The page title overrides the name the page has been given."}]},{"name":"Properties","properties":[{"name":"hide_in_nav","order":1,"data_type_node_id":18,"help_text":"help text2","description":"description2"}]}]');

INSERT INTO content_type VALUES (2, 36, 'Collexy.Home', 'Home Some description', 'fa fa-folder-o', 'fa fa-folder-o', 35, '{"template_node_id": 22, "allowed_templates_node_id": [22], "allowed_content_types_node_id": [37, 38, 39]}', '[{"name":"Content","properties":[{"name":"site_name","order":2,"data_type_node_id":2,"help_text":"help text","description":"Site name goes here."},{"name":"site_tagline","order":3,"data_type_node_id":2,"help_text":"help text","description":"Site tagline goes here."},{"name":"copyright","order":4,"data_type_node_id":2,"help_text":"help text","description":"Copyright here."},{"name":"domains","order":5,"data_type_node_id":19,"help_text":"help text","description":"Domains goes here."}]},{"name":"Social","properties":[{"name":"facebook_link","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your facebook link here."},{"name":"twitter_link","order":2,"data_type_node_id":2,"help_text":"help text","description":"Enter your twitter link here."},{"name":"linkedin_link","order":3,"data_type_node_id":2,"help_text":"help text","description":"Enter your linkedin link here."},{"name":"google_link","order":4,"data_type_node_id":2,"help_text":"help text","description":"Enter your Google+ profile link here."},{"name":"rss_link","order":5,"data_type_node_id":2,"help_text":"help text","description":"Enter your RSS feed link here."}]},{"name":"Banner","properties":[{"name": "hide_banner", "order": 1, "data_type_node_id": 18, "help_text": "help text2", "description": "description2"},{"name": "banner_header", "order": 2, "data_type_node_id": 2, "help_text": "help text", "description": "Banner header."},{"name": "banner_subheader", "order": 3, "data_type_node_id": 2, "help_text": "help text", "description": "Banner subheader."},{"name": "banner_link_text", "order": 4, "data_type_node_id": 2, "help_text": "help text", "description": "Banner link text."},{"name": "banner_link", "order": 5, "data_type_node_id": 2, "help_text": "help text", "description": "Banner link should ideally use a content picker data type."},{"name": "banner_background_image", "order": 6, "data_type_node_id": 2, "help_text": "help text", "description": "This should ideally use the upload data type."}]},{"name":"About","properties":[{"name": "about_title", "order": 1, "data_type_node_id": 2, "help_text": "help text", "description": "About title."},{"name": "about_text", "order": 2, "data_type_node_id": 4, "help_text": "help text", "description": "About text."}]}]');

INSERT INTO content_type VALUES (3, 37, 'Collexy.Post', 'Post content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 35, '{"template_node_id": 23, "allowed_templates_node_id": [23], "allowed_content_types_node_id": [37]}', '[{"name":"Content","properties":[{"name":"is_featured","order":2,"data_type_node_id":18,"help_text":"help text2","description":"description2"},{"name":"image","order":3,"data_type_node_id":2,"help_text":"Help text for image","description":"Image url"},{"name":"sub_header","order":4,"data_type_node_id":2,"help_text":"Help text for subheader","description":"Subheader description"},{"name":"content","order":5,"data_type_node_id":4,"help_text":"Help text for post content","description":"Post content description"}]}]');

INSERT INTO content_type VALUES (4, 38, 'Collexy.PostOverview', 'Post overview content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 35, '{"template_node_id": 24, "allowed_templates_node_id": [24], "allowed_content_types_node_id": [38]}', NULL);
INSERT INTO content_type VALUES (5, 39, 'Collexy.Page', 'Page content type desc', 'fa fa-folder-o', 'fa fa-folder-o', 35, '{"template_node_id": 25, "allowed_templates_node_id": [25], "allowed_content_types_node_id": [39]}', '[{"name":"Content","properties":[{"name":"content","order":2,"data_type_node_id":4,"help_text":"Help text for page contentent","description":"Page content description"}]}]');


INSERT INTO content_type VALUES (6, 40, 'mtFolder', 'Folder media type description1', 'mt-icon1', 'mt-thumbnail1', NULL, '{"allowed_content_types_node_id": [16, 15]}', '[{"name":"Folder","properties":[{"name":"folder_browser","order":1,"data_type_node_id":34,"help_text":"prop help text","description":"prop description"},{"name":"path","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties"}]');

INSERT INTO content_type VALUES (7, 41, 'Collexy.Image', 'Image content type description', 'fa fa-folder-o', 'fa fa-folder-o', NULL, 'null', '[{"name":"Image","properties":[{"name":"path","order":1,"data_type_node_id":2,"help_text":"help text","description":"URL goes here."},{"name":"title","order":2,"data_type_node_id":2,"help_text":"help text","description":"The title entered here can override the above one."},{"name":"caption","order":3,"data_type_node_id":4,"help_text":"help text","description":"Caption goes here."},{"name":"alt","order":4,"data_type_node_id":4,"help_text":"help text","description":"Alt goes here."},{"name":"description","order":5,"data_type_node_id":4,"help_text":"help text","description":"Description goes here."},{"name":"file_upload","order":1,"data_type_node_id":16,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties","properties":[{"name":"temporary property","order":1,"data_type_node_id":2,"help_text":"help text","description":"Temporary description goes here."}]}]');


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

INSERT INTO member VALUES (1, 'default_member', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'default_member@mail.com', '{"comments": "default user comments"}', '2015-01-22 14:25:38.904', NULL, '2015-02-19 23:46:00.495', NULL, 1, 'GIWES3RHMY5RKC7OZPOQTF5FQFWX32D5VLV3CAKT4HGKP5LZIENA', 20, '{1}');


--
-- TOC entry 2348 (class 0 OID 84053)
-- Dependencies: 181
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_group VALUES (1, 'authenticated_member', 'All logged in members');


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

INSERT INTO member_type VALUES (1, 20, 'Collexy.Member', 'Default member type', 'fa fa-user fa-fw', 1, NULL, '[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_node_id": 4}]}]');


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

INSERT INTO node VALUES (21, '1.21', 'Layout', 3, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
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

INSERT INTO node VALUES (35, '1.35', 'Master', 4, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (36, '1.35.36', 'Home', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node VALUES (37, '1.35.37', 'Post', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node VALUES (38, '1.35.38', 'Post Overview', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);
INSERT INTO node VALUES (39, '1.35.39', 'Page', 4, 1, '2014-10-22 16:51:00.215', 35, NULL, NULL);

INSERT INTO node VALUES (40, '1.40', 'Folder', 7, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (41, '1.41', 'Image', 7, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);

INSERT INTO node VALUES (42, '1.42', 'Home', 1, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
INSERT INTO node VALUES (43, '1.42.43', 'Welcome', 1, 1, '2014-10-22 16:51:00.215', 42, '[{"id": 2, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]', NULL);
INSERT INTO node VALUES (44, '1.42.44', 'Getting Started', 1, 1, '2014-10-26 23:19:44.735', 42, NULL, '[{"id": 1, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]');
INSERT INTO node VALUES (45, '1.42.45', 'Documentation', 1, 1, '2014-10-26 23:19:44.735', 42, NULL, NULL);
INSERT INTO node VALUES (46, '1.42.46', 'Get Involved', 1, 1, '2014-10-26 23:19:44.735', 42, NULL, NULL); -- NEW
INSERT INTO node VALUES (47, '1.42.47', 'Posts', 1, 1, '2014-10-22 16:51:00.215', 42, NULL, NULL);
INSERT INTO node VALUES (48, '1.42.47.48', 'Hello World', 1, 1, '2014-10-22 16:51:00.215', 47, NULL, NULL);
INSERT INTO node VALUES (49, '1.42.47.49', 'TXT Starter Kit For Collexy Released', 1, 1, '2014-10-22 16:51:00.215', 47, NULL, NULL);
INSERT INTO node VALUES (50, '1.42.47.50', 'You Need To Read This', 1, 1, '2014-10-22 16:51:00.215', 47, NULL, NULL);
INSERT INTO node VALUES (51, '1.42.47.51', 'Amazing Post', 1, 1, '2014-10-22 16:51:00.215', 47, NULL, NULL);

INSERT INTO node VALUES (52, '1.52', 'Sample Images', 2, 1, '2014-12-02 01:42:09.979', 1, NULL, NULL);

INSERT INTO node VALUES (53, '1.52.53', 'TXT', 2, 1, '2014-12-05 16:18:29.762', 52, NULL, NULL);
INSERT INTO node VALUES (54, '1.52.53.54', 'pic01.jpg', 2, 1, '2014-12-06 13:07:08.943', 53, NULL, NULL);
INSERT INTO node VALUES (55, '1.42.53.55', 'pic02.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node VALUES (56, '1.42.53.56', 'pic03.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node VALUES (57, '1.42.53.57', 'pic04.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node VALUES (58, '1.42.53.58', 'pic05.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);
INSERT INTO node VALUES (59, '1.42.53.59', 'banner.jpg', 2, 1, '2014-12-06 14:28:52.117', 53, NULL, NULL);


-- INSERT INTO node VALUES (42, '1.42', 'Home', 1, 1, '2014-10-22 16:51:00.215', 1, NULL, NULL);
-- INSERT INTO node VALUES (43, '1.42.43', 'Sample Page', 1, 1, '2014-10-22 16:51:00.215', 42, '[{"id": 2, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]', NULL);
-- INSERT INTO node VALUES (44, '1.42.43.44', 'Child Page Level 1', 1, 1, '2014-10-26 23:19:44.735', 43, NULL, '[{"id": 1, "permissions": ["node_create", "node_delete", "node_update", "node_move", "node_copy", "node_public_access", "node_permissions", "node_send_to_publish", "node_sort", "node_publish", "node_browse", "node_change_content_type"]}]');
-- INSERT INTO node VALUES (45, '1.42.43.44.45', 'Child Page Level 2', 1, 1, '2014-10-26 23:19:44.735', 44, NULL, NULL);
-- INSERT INTO node VALUES (46, '1.42.46', 'Posts', 1, 1, '2014-10-22 16:51:00.215', 42, NULL, NULL);
-- INSERT INTO node VALUES (47, '1.42.46.47', 'Hello World', 1, 1, '2014-10-22 16:51:00.215', 46, NULL, NULL);
-- INSERT INTO node VALUES (48, '1.48', 'gopher.jpg', 2, 1, '2014-10-28 15:50:47.303', 1, NULL, NULL);
-- INSERT INTO node VALUES (49, '1.49', '2014', 2, 1, '2014-12-02 01:42:09.979', 1, NULL, NULL);
-- INSERT INTO node VALUES (50, '1.49.50', '12', 2, 1, '2014-12-05 16:18:29.762', 49, NULL, NULL);
-- INSERT INTO node VALUES (51, '1.49.50.51', 'cat-prays.jpg', 2, 1, '2014-12-06 13:07:08.943', 50, NULL, NULL);
-- INSERT INTO node VALUES (52, '1.49.50.52', 'sleeping-kitten.jpg', 2, 1, '2014-12-06 14:28:52.117', 50, NULL, NULL);



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
-- TOC entry 2358 (class 0 OID 84094)
-- Dependencies: 191
-- Data for Name: route; Type: TABLE DATA; Schema: public; Owner: postgres
--

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

INSERT INTO template VALUES (1, 21, 'Collexy.Layout', false, '{30,34}', NULL);
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

INSERT INTO "user" VALUES (1, '%s', 'Admin', 'Demo', '%s', '%s', '2014-11-15 16:51:00.215', NULL, '2015-03-02 13:07:57.994', NULL, 1, 'IF2LJ42JDYHDEH55IXUTCTYMWEFO6PD2VJGYTWVYT353US52LFXQ', '{1}', NULL);


--
-- TOC entry 2363 (class 0 OID 84118)
-- Dependencies: 196
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: postgres
--


INSERT INTO user_group VALUES (1, 'Administrator', 'administrator', '{node_create,node_delete,node_update,node_move,node_copy,node_public_access,node_permissions,node_send_to_publish,node_publish,node_browse,node_change_content_type,admin,content_all,content_create,content_delete,content_update,content_section,content_browse,media_all,media_create,media_delete,media_update,media_section,media_browse,users_all,users_create,users_delete,users_update,users_section,users_browse,user_types_all,user_types_create,user_types_delete,user_types_update,user_types_section,user_types_browse,user_groups_all,user_groups_create,user_groups_delete,user_groups_update,user_groups_section,user_groups_browse,members_all,members_create,members_delete,members_update,members_section,members_browse,member_types_all,member_types_create,member_types_delete,member_types_update,member_types_section,member_types_browse,member_groups_all,member_groups_create,member_groups_delete,member_groups_update,member_groups_section,member_groups_browse,templates_all,templates_create,templates_delete,templates_update,templates_section,templates_browse,scripts_all,scripts_create,scripts_delete,scripts_update,scripts_section,scripts_browse,stylesheets_all,stylesheets_create,stylesheets_delete,stylesheets_update,stylesheets_section,stylesheets_browse,settings_section,settings_all,node_sort,content_types_all,content_types_create,content_types_delete,content_types_update,content_types_section,content_types_browse,media_types_all,media_types_create,media_types_delete,media_types_update,media_types_section,media_types_browse,data_types_all,data_types_create,data_types_delete,data_types_update,data_types_section,data_types_browse}');
INSERT INTO user_group VALUES (2, 'Editor', 'editor', NULL);
INSERT INTO user_group VALUES (3, 'Writer', 'writer', NULL);

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


-- Completed on 2015-03-03 12:27:47

--
-- PostgreSQL database dump complete
--`