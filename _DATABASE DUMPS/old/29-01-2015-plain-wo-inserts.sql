--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-01-29 19:36:57

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- TOC entry 196 (class 3079 OID 11855)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2355 (class 0 OID 0)
-- Dependencies: 196
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 197 (class 3079 OID 41274)
-- Name: citext; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- TOC entry 2356 (class 0 OID 0)
-- Dependencies: 197
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- TOC entry 198 (class 3079 OID 16394)
-- Name: ltree; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;


--
-- TOC entry 2357 (class 0 OID 0)
-- Dependencies: 198
-- Name: EXTENSION ltree; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';


SET search_path = public, pg_catalog;

--
-- TOC entry 281 (class 1255 OID 16655)
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
-- TOC entry 290 (class 1255 OID 33087)
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
-- TOC entry 287 (class 1255 OID 16671)
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
-- TOC entry 179 (class 1259 OID 16616)
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
-- TOC entry 178 (class 1259 OID 16614)
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
-- TOC entry 2358 (class 0 OID 0)
-- Dependencies: 178
-- Name: content_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_id_seq OWNED BY content.id;


--
-- TOC entry 177 (class 1259 OID 16605)
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
-- TOC entry 176 (class 1259 OID 16603)
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
-- TOC entry 2359 (class 0 OID 0)
-- Dependencies: 176
-- Name: content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_type_id_seq OWNED BY content_type.id;


--
-- TOC entry 175 (class 1259 OID 16594)
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
-- TOC entry 174 (class 1259 OID 16592)
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
-- TOC entry 2360 (class 0 OID 0)
-- Dependencies: 174
-- Name: data_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_type_id_seq OWNED BY data_type.id;


--
-- TOC entry 188 (class 1259 OID 57658)
-- Name: domain; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE domain (
    id integer NOT NULL,
    node_id bigint,
    name character varying
);


ALTER TABLE domain OWNER TO postgres;

--
-- TOC entry 189 (class 1259 OID 57661)
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
-- TOC entry 2361 (class 0 OID 0)
-- Dependencies: 189
-- Name: domain_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE domain_id_seq OWNED BY domain.id;


--
-- TOC entry 193 (class 1259 OID 57775)
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
-- TOC entry 194 (class 1259 OID 57784)
-- Name: member_group; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE member_group (
    id integer NOT NULL,
    name character varying,
    description text
);


ALTER TABLE member_group OWNER TO postgres;

--
-- TOC entry 195 (class 1259 OID 57787)
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
-- TOC entry 2362 (class 0 OID 0)
-- Dependencies: 195
-- Name: member_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_group_id_seq OWNED BY member_group.id;


--
-- TOC entry 192 (class 1259 OID 57773)
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
-- TOC entry 2363 (class 0 OID 0)
-- Dependencies: 192
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_id_seq OWNED BY member.id;


--
-- TOC entry 190 (class 1259 OID 57676)
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
-- TOC entry 191 (class 1259 OID 57688)
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
-- TOC entry 2364 (class 0 OID 0)
-- Dependencies: 191
-- Name: member_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_type_id_seq OWNED BY member_type.id;


--
-- TOC entry 181 (class 1259 OID 16627)
-- Name: node; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE node (
    id bigint NOT NULL,
    path ltree,
    name character varying(255),
    node_type smallint NOT NULL,
    created_by bigint NOT NULL,
    created_date timestamp without time zone DEFAULT now() NOT NULL,
    parent_id bigint
);


ALTER TABLE node OWNER TO postgres;

--
-- TOC entry 180 (class 1259 OID 16625)
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
-- TOC entry 2365 (class 0 OID 0)
-- Dependencies: 180
-- Name: node_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE node_id_seq OWNED BY node.id;


--
-- TOC entry 184 (class 1259 OID 49478)
-- Name: permission; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE permission (
    id integer NOT NULL,
    name character varying NOT NULL,
    description character varying
);


ALTER TABLE permission OWNER TO postgres;

--
-- TOC entry 185 (class 1259 OID 49481)
-- Name: permission_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE permission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE permission_id_seq OWNER TO postgres;

--
-- TOC entry 2366 (class 0 OID 0)
-- Dependencies: 185
-- Name: permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE permission_id_seq OWNED BY permission.id;


--
-- TOC entry 187 (class 1259 OID 49492)
-- Name: role; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE role (
    id integer NOT NULL,
    name character varying NOT NULL,
    permission_ids integer[]
);


ALTER TABLE role OWNER TO postgres;

--
-- TOC entry 186 (class 1259 OID 49490)
-- Name: role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE role_id_seq OWNER TO postgres;

--
-- TOC entry 2367 (class 0 OID 0)
-- Dependencies: 186
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE role_id_seq OWNED BY role.id;


--
-- TOC entry 183 (class 1259 OID 16639)
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
-- TOC entry 182 (class 1259 OID 16637)
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
-- TOC entry 2368 (class 0 OID 0)
-- Dependencies: 182
-- Name: template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE template_id_seq OWNED BY template.id;


--
-- TOC entry 173 (class 1259 OID 16571)
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
    role_ids integer[]
);


ALTER TABLE "user" OWNER TO postgres;

--
-- TOC entry 172 (class 1259 OID 16569)
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
-- TOC entry 2369 (class 0 OID 0)
-- Dependencies: 172
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- TOC entry 2186 (class 2604 OID 16619)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content ALTER COLUMN id SET DEFAULT nextval('content_id_seq'::regclass);


--
-- TOC entry 2185 (class 2604 OID 16608)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content_type ALTER COLUMN id SET DEFAULT nextval('content_type_id_seq'::regclass);


--
-- TOC entry 2184 (class 2604 OID 16597)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_type ALTER COLUMN id SET DEFAULT nextval('data_type_id_seq'::regclass);


--
-- TOC entry 2193 (class 2604 OID 57663)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY domain ALTER COLUMN id SET DEFAULT nextval('domain_id_seq'::regclass);


--
-- TOC entry 2195 (class 2604 OID 57778)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member ALTER COLUMN id SET DEFAULT nextval('member_id_seq'::regclass);


--
-- TOC entry 2198 (class 2604 OID 57789)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_group ALTER COLUMN id SET DEFAULT nextval('member_group_id_seq'::regclass);


--
-- TOC entry 2194 (class 2604 OID 57690)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_type ALTER COLUMN id SET DEFAULT nextval('member_type_id_seq'::regclass);


--
-- TOC entry 2187 (class 2604 OID 16630)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY node ALTER COLUMN id SET DEFAULT nextval('node_id_seq'::regclass);


--
-- TOC entry 2191 (class 2604 OID 49483)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY permission ALTER COLUMN id SET DEFAULT nextval('permission_id_seq'::regclass);


--
-- TOC entry 2192 (class 2604 OID 49495)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY role ALTER COLUMN id SET DEFAULT nextval('role_id_seq'::regclass);


--
-- TOC entry 2189 (class 2604 OID 16642)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY template ALTER COLUMN id SET DEFAULT nextval('template_id_seq'::regclass);


--
-- TOC entry 2182 (class 2604 OID 16574)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- TOC entry 2331 (class 0 OID 16616)
-- Dependencies: 179
-- Data for Name: content; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY content (id, node_id, content_type_node_id, meta, public_access) FROM stdin;
13	33	5	{"prop2": "prop2", "prop3": "prop3", "page_title": "About me", "page_content": "About page description", "template_node_id": 22}	\N
2	10	5	{"prop2": "prop2", "prop3": "sample page prop 3", "page_title": "Sample page title", "page_content": "Sample page content goes here", "template_node_id": 22}	\N
12	32	5	{"prop2": "p2", "prop3": "p3", "page_title": "Yet another test page title override", "page_content": "Yet another test page description goes here", "template_node_id": 8}	\N
14	23	16	\N	\N
7	13	5	{"prop2": "11", "prop3": "Another page prop 31", "page_title": "Another page title1", "page_content": "Another page content goes here1", "template_node_id": 25}	\N
19	40	5	{"prop2": "p2", "prop3": "p3", "page_title": "test page 2 child page title override", "page_content": "test page 2 child desc", "template_node_id": 25}	\N
4	19	15	{"alt": "Postgresql image alt text", "url": "/media/2014/10/postgresql.png", "path": "media\\\\postgresql.png", "caption": "This is the caption of the postgresql image", "description": "Postgresql image description"}	\N
18	39	5	{"prop2": "tp2p2", "prop3": "tp2p3", "page_title": "test page 2 title override", "page_content": "test page 2 content", "template_node_id": 22}	\N
5	11	5	{"prop3": "sample child page level 1 page prop 3", "page_title": "Child page level 1 title", "page_content": "Sample page - child page level 1 content goes here", "template_node_id": 22}	\N
15	36	16	{"path": "media\\\\Another Image Folder2"}	\N
8	24	15	{"alt": "Goku SSJ3 image alt text", "url": "/media/Sample picture folder/Goku_SSJ3.jpg", "path": "media\\\\Sample picture folder\\\\Goku_SSJ3.jpg", "caption": "This is the caption of the Goku SSJ3 image", "description": "Goku SSJ3 image description"}	\N
20	41	15	{"alt": "gopher alt", "url": "/media/Another Image Folder1/gopher.jpg", "path": "media\\\\Another Image Folder2\\\\gopher.jpg", "title": "gopher title", "caption": "gopher caption", "description": "gopher description"}	\N
11	31	5	{"prop3": "msp3", "page_title": "My Sub Page Override", "page_content": "mysubpage desc", "template_node_id": 8}	\N
29	54	15	{"alt": "catduck.jpg", "path": "media\\\\Subfolder depth test\\\\Level 1\\\\catduck.jpg", "title": "catduck.jpg", "caption": "catduck.jpg", "description": "catduck1.jpg"}	\N
3	18	15	{"alt": "Gopher image alt text1", "url": "/media/2014/10/gopher.jpg", "path": "media\\\\gopher.jpg", "caption": "This is the caption of the gopher image1", "description": "Gopher image description1", "temporary property": "lol"}	\N
9	29	5	{"prop2": "prop2", "prop3": "prop3", "page_title": "test page title override1", "page_content": "This is just a test page", "template_node_id": 8}	\N
6	12	5	{"prop3": "sample child page level 2 page prop 3", "page_title": "Child page level 2 title", "page_content": "Sample page - child page level 2 content goes here1", "template_node_id": 22}	{"groups": [1], "members": [1]}
16	37	16	{"path": "media\\\\Subfolder depth test"}	\N
17	38	16	{"path": "media\\\\2014"}	\N
21	45	16	{"path": "media\\\\2014\\\\12"}	\N
1	9	4	{"prop2": "Home page prop 2", "domains": ["localhost:8080", "localhost:8080/test"], "facebook": "facebook.com/home", "copyright": "&copy; 2014 codeish.com", "site_name": "Collexy cms test site", "page_title": "Home page title", "site_tagline": "Test site tagline", "template_node_id": 7}	\N
10	30	5	{"prop2": "", "prop3": "mypageprop3", "page_title": "Login page test", "page_content": "This is a login page for members", "template_node_id": 25}	\N
25	50	15	{"alt": "tiny.jpg", "path": "media\\\\Subfolder depth test\\\\Level 1\\\\Level 2\\\\tiny.jpg", "title": "tiny.jpg", "caption": "tiny.jpg", "description": "tiny1.jpg"}	\N
28	53	15	{"alt": "AngularLogo alt", "path": "media\\\\Subfolder depth test\\\\AnguarLogo.png", "title": "AngularLogo.png", "caption": "AngularLogo caption", "description": "AngularLogo desc"}	\N
27	52	15	{"alt": "taco-hamster.jpg", "path": "media\\\\Subfolder depth test\\\\Level 1\\\\Level 2\\\\taco-hamster.jpg", "title": "taco-hamster.jpg", "caption": "taco-hamster.jpg", "description": "taco-hamster.jpg"}	\N
24	49	15	{"alt": "blomkals-hamster.jpg", "path": "media\\\\Subfolder depth test\\\\Level 1\\\\Level 2\\\\blomkals-hamster.jpg", "title": "blomkals-hamster.jpg", "caption": "blomkals-hamster.jpg", "description": "blomkals-hamster.jpg"}	\N
31	56	15	{"alt": "ducks.jpg", "path": "media\\\\2014\\\\12\\\\ducks.jpg", "title": "ducks.jpg", "caption": "ducks.jpg", "description": "ducks.jpg3"}	\N
32	57	15	{"alt": "sleeping-kitten.jpg", "path": "media\\\\2014\\\\12\\\\sleeping-kitten.jpg", "title": "sleeping-kitten.jpg", "caption": "sleeping-kitten.jpg", "description": "sleeping-kitten.jpg"}	\N
30	55	15	{"alt": "cat-prays.jpg", "path": "media\\\\2014\\\\12\\\\cat-prays.jpg", "title": "cat-prays.jpg", "caption": "cat-prays.jpg", "description": "cat-prays.jpg"}	\N
22	46	16	{"path": "media\\\\Subfolder depth test\\\\Level 1"}	\N
23	47	16	{"path": "media\\\\Subfolder depth test\\\\Level 1\\\\Level 2"}	\N
26	51	15	{"alt": "dog.jpg", "path": "media\\\\Subfolder depth test\\\\Level 1\\\\Level 2\\\\dog.jpg", "title": "dog.jpg", "caption": "dog.jpg", "description": "dog.jpg"}	\N
\.


--
-- TOC entry 2370 (class 0 OID 0)
-- Dependencies: 178
-- Name: content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_id_seq', 32, true);


--
-- TOC entry 2329 (class 0 OID 16605)
-- Dependencies: 177
-- Data for Name: content_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) FROM stdin;
2	4	ctHome	Home Some description	fa fa-folder-o	fa fa-folder-o	3	{"template_node_id": 7, "allowed_templates_node_id": [7], "allowed_content_types_node_id": [5]}	[{"name":"Content","properties":[{"name":"site_name","order":2,"data_type_node_id":2,"help_text":"help text","description":"Site name goes here."},{"name":"site_tagline","order":3,"data_type_node_id":2,"help_text":"help text","description":"Site tagline goes here."},{"name":"copyright","order":4,"data_type_node_id":2,"help_text":"help text","description":"Copyright here."},{"name":"domains","order":5,"data_type_node_id":59,"help_text":"help text","description":"Domains goes here."}]},{"name":"Social","properties":[{"name":"facebook","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your facebook link here."},{"name":"twitter","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your twitter link here."},{"name":"linkedin","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your linkedin link here."}]}]
4	15	Image	Image content type description	fa fa-folder-o	fa fa-folder-o	\N	null	[{"name":"Image","properties":[{"name":"path","order":1,"data_type_node_id":2,"help_text":"help text","description":"URL goes here."},{"name":"title","order":2,"data_type_node_id":2,"help_text":"help text","description":"The title entered here can override the above one."},{"name":"caption","order":3,"data_type_node_id":14,"help_text":"help text","description":"Caption goes here."},{"name":"alt","order":4,"data_type_node_id":14,"help_text":"help text","description":"Alt goes here."},{"name":"description","order":5,"data_type_node_id":14,"help_text":"help text","description":"Description goes here."},{"name":"file_upload","order":1,"data_type_node_id":48,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties","properties":[{"name":"temporary property","order":1,"data_type_node_id":2,"help_text":"help text","description":"Temporary description goes here."}]}]
1	3	Master	Some description	fa fa-folder-o	fa fa-folder-o	\N	\N	[{"name": "Content", "properties": [{"name": "page_title", "order": 1, "data_type_node_id": 2, "help_text": "help text", "description": "The page title overrides the name the page has been given."}]}, {"name": "Properties", "properties": [{"name": "prop2", "order": 1, "data_type_node_id": 2, "help_text": "help text2", "description": "description2"}, {"name": "prop3", "order": 2, "data_type_node_id": 2, "help_text": "help text3", "description": "description3"}]}]
5	28	ctTestContentTypeAlias	ct-test desc	ct-test-icon	ct-test-thumbnail	3	{"template_node_id": 8, "allowed_templates_node_id": [8, 22, 25], "allowed_content_types_node_id": [5]}	[{"name":"Mytab1","properties":[{"name":"property name1","order":1,"data_type_node_id":2,"help_text":"prop help text1","description":"prop description1"},{"name":"property name2","order":2,"data_type_node_id":14,"help_text":"prop help text2","description":"prop description2"}]},{"name":"Mytab2","properties":[{"name":"property name3","order":1,"data_type_node_id":27,"help_text":"prop help text3","description":"prop description3"},{"name":"property name4","order":2,"data_type_node_id":14,"help_text":"prop help text4","description":"prop description4"}]}]
3	5	ctPage	Page content type desc	fa fa-folder-o	fa fa-folder-o	3	{"template_node_id": 8, "allowed_templates_node_id": [8, 22, 25], "allowed_content_types_node_id": [5]}	[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":14,"help_text":"Help text for page contentent","description":"Page content description"}]}]
7	35	mtTestMediaType alias	mtTest desc	mtTest-icon	mtTest-thumbnail	16	{"allowed_content_types_node_id": [15, 17, 16]}	[{"name":"mytab"}]
8	43	CT test alias	ct-test-desc	ct-test-icon	ct-test-thumb	3	{"template_node_id": "8", "allowed_templates_node_id": [22, 8, 25, 42], "allowed_content_types_node_id": [5, 28]}	[{"name":"mytab"}]
9	44	Test Content Type 2 alias	tc2-desc	tc2-icon	tc2-thumb	3	{"template_node_id": "8", "allowed_templates_node_id": [22, 8, 25, 26, 42], "allowed_content_types_node_id": [28, 43, 5]}	[{"name":"mytab"}]
6	16	mtFolder	Folder media type description1	mt-icon1	mt-thumbnail1	\N	{"allowed_content_types_node_id": [16, 15]}	[{"name":"Folder","properties":[{"name":"folder_browser","order":1,"data_type_node_id":34,"help_text":"prop help text","description":"prop description"},{"name":"path","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties"}]
\.


--
-- TOC entry 2371 (class 0 OID 0)
-- Dependencies: 176
-- Name: content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_id_seq', 9, true);


--
-- TOC entry 2327 (class 0 OID 16594)
-- Dependencies: 175
-- Data for Name: data_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY data_type (id, node_id, html, alias) FROM stdin;
6	59	<div>\n\t<input type="text"/> <button type="button">Add domain</button><br>\n\t<ul>\n\t\t<li ng-repeat="domain in data.meta[prop.name]">{{domain}}</li>\n\t</ul>\n\t<button type="button">Delete selected</button>\n</div>	defDomains
2	14	<textarea id="{{prop.name}}" ng-model="data.meta[prop.name]">	defTextarea
1	2	<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]">	defTextInput
3	27	<colorpicker>The color picker data type is not implemented yet!</colorpicker>	dtColorPicker
4	34	<folderbrowser>This is an awesome folder browser (unimplemented datatype)</folderbrowser>	dtFolderBrowser
5	48	<input type="file" file-input="test.files" multiple />\n<button ng-click="upload()" type="button">Upload</button>\n<li ng-repeat="file in test.files">{{file.name}}</li>\n<!--<input type="file" onchange="angular.element(this).scope().filesChanged(this)" multiple />\n<button ng-click="upload()">Upload</button>\n<li ng-repeat="file in files">{{file.name}}</li>-->	dtFileUpload
\.


--
-- TOC entry 2372 (class 0 OID 0)
-- Dependencies: 174
-- Name: data_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_type_id_seq', 6, true);


--
-- TOC entry 2340 (class 0 OID 57658)
-- Dependencies: 188
-- Data for Name: domain; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY domain (id, node_id, name) FROM stdin;
1	9	localhost:8080
\.


--
-- TOC entry 2373 (class 0 OID 0)
-- Dependencies: 189
-- Name: domain_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('domain_id_seq', 1, true);


--
-- TOC entry 2345 (class 0 OID 57775)
-- Dependencies: 193
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member (id, username, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_node_id, group_ids) FROM stdin;
1	default_member	$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa	default_member@mail.com	{"comments": "default user comments"}	2015-01-22 14:25:38.904	\N	2015-01-29 09:59:27.661	\N	1	QPXOHWKZJA654NLAE3PL3E4NYA64OKR67PVMIG4UNACJVZP32LMQ	61	{1}
\.


--
-- TOC entry 2346 (class 0 OID 57784)
-- Dependencies: 194
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member_group (id, name, description) FROM stdin;
1	authenticated_member	All logged in members
\.


--
-- TOC entry 2374 (class 0 OID 0)
-- Dependencies: 195
-- Name: member_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_group_id_seq', 1, true);


--
-- TOC entry 2375 (class 0 OID 0)
-- Dependencies: 192
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_id_seq', 1, true);


--
-- TOC entry 2342 (class 0 OID 57676)
-- Dependencies: 190
-- Data for Name: member_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member_type (id, node_id, alias, description, icon, parent_member_type_node_id, meta, tabs) FROM stdin;
1	61	mtMember	Default member type	fa fa-user fa-fw	1	\N	[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_node_id": 14}]}]
\.


--
-- TOC entry 2376 (class 0 OID 0)
-- Dependencies: 191
-- Name: member_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_type_id_seq', 1, true);


--
-- TOC entry 2333 (class 0 OID 16627)
-- Dependencies: 181
-- Data for Name: node; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY node (id, path, name, node_type, created_by, created_date, parent_id) FROM stdin;
14	1.14	Textarea	11	1	2014-10-27 02:40:41.179	1
48	1.48	File upload	11	1	2014-12-05 19:56:17.883	\N
10	1.9.10	Sample Page	1	1	2014-10-22 16:51:00.215	9
60	1.6.60	404	3	1	2015-01-20 13:46:33.668	6
20	1.20	Sidebar 1	3	1	2014-11-10 09:03:20.514	1
15	1.15	Image	7	1	2014-10-28 15:16:25.972	1
36	1.36	Another Image Folder2	2	1	2014-12-02 01:00:51.206	1
21	1.21	Sidebar 2	3	1	2014-11-10 23:56:55.038	1
1	1	root	5	1	2014-10-22 16:51:00.215	\N
2	1.2	Text input	11	1	2014-10-22 16:51:00.215	1
3	1.3	Master	4	1	2014-10-22 16:51:00.215	1
13	1.9.13	Another Page	1	1	2014-10-26 23:27:14.571	9
9	1.9	Home	1	1	2014-10-22 16:51:00.215	1
61	1.61	Member	12	1	2015-01-22 15:55:13.957	1
28	1.3.28	Test Content Type	4	1	2014-11-26 04:20:48.026	3
26	1.6.25.26	Child of test template	3	1	2014-11-26 01:39:42.816	25
46	1.37.46	Level 1	2	1	2014-12-05 17:02:13.875	37
41	1.36.41	gopher.jpg	2	1	2014-12-02 02:08:26.737	36
17	1.17	File	7	1	2014-10-28 15:18:13.4	1
18	1.18	gopher.jpg	2	1	2014-10-28 15:50:47.303	1
16	1.16	Folder	7	1	2014-10-28 15:18:13.4	1
35	1.35	Media Type Test	7	1	2014-12-01 22:09:43.783	1
23	1.23	Sample picture folder	2	1	2014-11-17 16:57:14.654	1
47	1.37.46.47	Level 2	2	1	2014-12-05 17:02:46.762	46
43	1.3.43	Content Type Test	4	1	2014-12-02 12:38:59.527	3
55	1.38.45.55	cat-prays.jpg	2	1	2014-12-06 13:07:08.943	45
54	1.37.46.54	catduck.jpg	2	1	2014-12-06 03:44:40.07	46
44	1.3.44	Test Content Type 2	4	1	2014-12-02 12:48:25.307	3
58	1.6.58	Unauthorized	3	1	2014-12-15 14:24:22.063	6
34	1.34	Folder Browser	11	1	2014-12-01 16:09:46.488	1
19	1.19	postgresql.png	2	1	2014-10-28 17:53:37.488	1
27	1.27	Color Picker	11	1	2014-11-26 02:20:17.638	1
33	1.9.33	About	1	1	2014-12-01 12:11:25.838	9
24	1.23.24	Goku_SSJ3.jpg	2	1	2014-11-17 16:58:57.285	23
11	1.9.10.11	Child Page Level 1	1	1	2014-10-26 23:19:44.735	10
4	1.3.4	Home	4	1	2014-10-22 16:51:00.215	3
31	1.9.30.31	MySubPage	1	1	2014-12-01 12:02:54.252	30
32	1.9.32	Yet another test page	1	1	2014-12-01 12:07:29.999	9
40	1.9.39.40	Test page 2 child	1	1	2014-12-02 01:45:49.78	39
39	1.9.39	Testpage 2	1	1	2014-12-02 01:43:33.233	9
5	1.3.5	Page	4	1	2014-10-22 16:51:00.215	3
51	1.37.46.47.51	dog.jpg	2	1	2014-12-05 21:06:49.532	47
52	1.37.46.47.52	taco-hamster.jpg	2	1	2014-12-05 21:22:45.227	47
22	1.6.22	Page with sidebars	3	1	2014-11-11 03:39:55.766	6
49	1.37.46.47.49	blomkals-hamster.jpg	2	1	2014-12-05 20:44:25.921	47
38	1.38	2014	2	1	2014-12-02 01:42:09.979	1
50	1.37.46.47.50	tiny.jpg	2	1	2014-12-05 21:05:42.816	47
57	1.38.45.57	sleeping-kitten.jpg	2	1	2014-12-06 14:28:52.117	45
37	1.37	Subfolder depth test3	2	1	2014-12-02 01:37:09.125	1
8	1.6.8	Page	3	1	2014-10-22 16:51:00.215	6
45	1.38.45	12	2	1	2014-12-05 16:18:29.762	38
6	1.6	Layout	3	1	2014-10-22 16:51:00.215	1
56	1.38.45.56	ducks.jpg	2	1	2014-12-06 13:10:14.637	45
53	1.37.53	AngularLogo.png	2	1	2014-12-06 03:36:14.425	37
42	1.6.42	Test template 2	3	1	2014-12-02 02:19:29.241	6
7	1.6.7	Home	3	1	2014-10-22 16:51:00.215	6
29	1.9.29	Test Page	1	1	2014-12-01 11:45:16.186	9
30	1.9.30	Login	1	1	2014-12-01 11:54:10.208	9
59	1.59	Domains	11	1	2015-01-19 21:22:06.945	\N
12	1.9.10.11.12	Child Page Level 2	1	1	2014-10-26 23:19:44.735	11
25	1.6.25	Login	3	1	2014-11-26 00:13:45.309	6
\.


--
-- TOC entry 2377 (class 0 OID 0)
-- Dependencies: 180
-- Name: node_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('node_id_seq', 61, true);


--
-- TOC entry 2336 (class 0 OID 49478)
-- Dependencies: 184
-- Data for Name: permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY permission (id, name, description) FROM stdin;
2	create_new_content	\N
4	delete_own_content	\N
6	edit_own_content	\N
8	administer_content_types	Warning: only give this permission to trusted roles!
1	use_administration_pages	Warning: only give this permission to trusted roles!
7	administer_content	Warning: only give this permission to trusted roles!
3	delete_any_content	Warning: only give this permission to trusted roles!
5	edit_any_content	Warning: only give this permission to trusted roles!
9	view_published_content	\N
10	view_unpublished_content	\N
11	administer_users	\N
12	administer_site_configuration	\N
13	administer_content_types	\N
14	administer_data_types	\N
\.


--
-- TOC entry 2378 (class 0 OID 0)
-- Dependencies: 185
-- Name: permission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('permission_id_seq', 14, true);


--
-- TOC entry 2339 (class 0 OID 49492)
-- Dependencies: 187
-- Data for Name: role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY role (id, name, permission_ids) FROM stdin;
2	authenticated_user	{9}
1	anonymous_user	{9}
3	administrator	{1,2,3,4,5,6,7,8,9,10,11,12,13,14}
\.


--
-- TOC entry 2379 (class 0 OID 0)
-- Dependencies: 186
-- Name: role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('role_id_seq', 3, true);


--
-- TOC entry 2335 (class 0 OID 16639)
-- Dependencies: 183
-- Data for Name: template; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY template (id, node_id, alias, is_partial, partial_template_node_ids, parent_template_node_id) FROM stdin;
10	58		f	{0}	6
6	22	Page with sidebars	f	{20,21}	6
1	6	Layout	f	{0}	0
11	60	tmpl404	f	{}	6
7	25	tplLogin	f	{20,21}	6
8	26	child of test template alias	f	{}	25
4	20	Sidebar 1	t	{}	0
5	21	Sidebar 2	t	{}	0
3	8	Page	f	{0}	6
9	42	tmpTestTemplate2	f	{0}	6
2	7	Home	f	{0}	6
\.


--
-- TOC entry 2380 (class 0 OID 0)
-- Dependencies: 182
-- Name: template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('template_id_seq', 12, true);


--
-- TOC entry 2325 (class 0 OID 16571)
-- Dependencies: 173
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, role_ids) FROM stdin;
1	soren	Soren	Tester	$2a$10$UNrly6WSmQnm495KAth6Auk4Z.11kjDBRFz8ZKjhqthytKFH/TjKq	soren@codeish.com	2014-10-21 16:51:00.215	\N	2015-01-13 02:05:18.422	\N	1	23NJGGPPTY2IYKMMTVPWOYNONQJHH5BGVWXCD5PQSKELZKYNQPBQ	{2}
2	admin	Admin	Demo	$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa	demo@codeish.com	2014-11-15 16:51:00.215	\N	2015-01-29 11:02:02.869	\N	1	Q5MRP2SZISGCSCHLCZUOU6DAPVWDKVAGQ4USEMZWCTGS6DGBHKOA	{3}
\.


--
-- TOC entry 2381 (class 0 OID 0)
-- Dependencies: 172
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 2, true);


--
-- TOC entry 2206 (class 2606 OID 16613)
-- Name: content_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY content_type
    ADD CONSTRAINT content_type_pkey PRIMARY KEY (id);


--
-- TOC entry 2204 (class 2606 OID 16602)
-- Name: data_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY data_type
    ADD CONSTRAINT data_type_pkey PRIMARY KEY (id);


--
-- TOC entry 2208 (class 2606 OID 16624)
-- Name: document_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY content
    ADD CONSTRAINT document_pkey PRIMARY KEY (id);


--
-- TOC entry 2211 (class 2606 OID 16636)
-- Name: node_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY node
    ADD CONSTRAINT node_pkey PRIMARY KEY (id);


--
-- TOC entry 2214 (class 2606 OID 16647)
-- Name: template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY template
    ADD CONSTRAINT template_pkey PRIMARY KEY (id);


--
-- TOC entry 2200 (class 2606 OID 49500)
-- Name: user_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- TOC entry 2202 (class 2606 OID 16579)
-- Name: user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 2209 (class 1259 OID 41272)
-- Name: idxgin; Type: INDEX; Schema: public; Owner: postgres; Tablespace: 
--

CREATE INDEX idxgin ON content USING gin (meta);


--
-- TOC entry 2212 (class 1259 OID 41273)
-- Name: template_partial_template_node_ids_idx; Type: INDEX; Schema: public; Owner: postgres; Tablespace: 
--

CREATE INDEX template_partial_template_node_ids_idx ON template USING gin (partial_template_node_ids);


--
-- TOC entry 2354 (class 0 OID 0)
-- Dependencies: 5
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2015-01-29 19:36:57

--
-- PostgreSQL database dump complete
--

