--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-02-20 12:05:18

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- TOC entry 204 (class 3079 OID 11855)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2387 (class 0 OID 0)
-- Dependencies: 204
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 205 (class 3079 OID 41274)
-- Name: citext; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- TOC entry 2388 (class 0 OID 0)
-- Dependencies: 205
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- TOC entry 206 (class 3079 OID 16394)
-- Name: ltree; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;


--
-- TOC entry 2389 (class 0 OID 0)
-- Dependencies: 206
-- Name: EXTENSION ltree; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';


SET search_path = public, pg_catalog;

--
-- TOC entry 289 (class 1255 OID 16655)
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
-- TOC entry 298 (class 1255 OID 33087)
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
-- TOC entry 295 (class 1255 OID 16671)
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
-- TOC entry 199 (class 1259 OID 66129)
-- Name: adm_route; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE adm_route (
    id integer NOT NULL,
    name character varying,
    alias character varying,
    path ltree,
    parent_id integer,
    type smallint,
    icon character varying,
    url character varying,
    components jsonb,
    redirect_to character varying,
    data jsonb,
    ref character varying
);


ALTER TABLE adm_route OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 66127)
-- Name: ang_routes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE ang_routes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE ang_routes_id_seq OWNER TO postgres;

--
-- TOC entry 2390 (class 0 OID 0)
-- Dependencies: 198
-- Name: ang_routes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE ang_routes_id_seq OWNED BY adm_route.id;


--
-- TOC entry 181 (class 1259 OID 16616)
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
-- TOC entry 180 (class 1259 OID 16614)
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
-- TOC entry 2391 (class 0 OID 0)
-- Dependencies: 180
-- Name: content_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_id_seq OWNED BY content.id;


--
-- TOC entry 179 (class 1259 OID 16605)
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
-- TOC entry 178 (class 1259 OID 16603)
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
-- TOC entry 2392 (class 0 OID 0)
-- Dependencies: 178
-- Name: content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_type_id_seq OWNED BY content_type.id;


--
-- TOC entry 177 (class 1259 OID 16594)
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
-- TOC entry 176 (class 1259 OID 16592)
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
-- TOC entry 2393 (class 0 OID 0)
-- Dependencies: 176
-- Name: data_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_type_id_seq OWNED BY data_type.id;


--
-- TOC entry 186 (class 1259 OID 57658)
-- Name: domain; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE domain (
    id integer NOT NULL,
    node_id bigint,
    name character varying
);


ALTER TABLE domain OWNER TO postgres;

--
-- TOC entry 187 (class 1259 OID 57661)
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
-- TOC entry 2394 (class 0 OID 0)
-- Dependencies: 187
-- Name: domain_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE domain_id_seq OWNED BY domain.id;


--
-- TOC entry 191 (class 1259 OID 57775)
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
-- TOC entry 192 (class 1259 OID 57784)
-- Name: member_group; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE member_group (
    id integer NOT NULL,
    name character varying,
    description text
);


ALTER TABLE member_group OWNER TO postgres;

--
-- TOC entry 193 (class 1259 OID 57787)
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
-- TOC entry 2395 (class 0 OID 0)
-- Dependencies: 193
-- Name: member_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_group_id_seq OWNED BY member_group.id;


--
-- TOC entry 190 (class 1259 OID 57773)
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
-- TOC entry 2396 (class 0 OID 0)
-- Dependencies: 190
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_id_seq OWNED BY member.id;


--
-- TOC entry 188 (class 1259 OID 57676)
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
-- TOC entry 189 (class 1259 OID 57688)
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
-- TOC entry 2397 (class 0 OID 0)
-- Dependencies: 189
-- Name: member_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_type_id_seq OWNED BY member_type.id;


--
-- TOC entry 203 (class 1259 OID 74363)
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
    user_ids integer[],
    user_group_ids integer[]
);


ALTER TABLE menu_link OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 74361)
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
-- TOC entry 2398 (class 0 OID 0)
-- Dependencies: 202
-- Name: menu_link_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE menu_link_id_seq OWNED BY menu_link.id;


--
-- TOC entry 183 (class 1259 OID 16627)
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
-- TOC entry 182 (class 1259 OID 16625)
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
-- TOC entry 2399 (class 0 OID 0)
-- Dependencies: 182
-- Name: node_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE node_id_seq OWNED BY node.id;


--
-- TOC entry 196 (class 1259 OID 66016)
-- Name: permission; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE permission (
    id integer NOT NULL,
    name character varying
);


ALTER TABLE permission OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 66019)
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
-- TOC entry 2400 (class 0 OID 0)
-- Dependencies: 197
-- Name: permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE permission_id_seq OWNED BY permission.id;


--
-- TOC entry 200 (class 1259 OID 74340)
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
-- TOC entry 2401 (class 0 OID 0)
-- Dependencies: 200
-- Name: COLUMN route.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN route.path IS '
';


--
-- TOC entry 201 (class 1259 OID 74343)
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
-- TOC entry 2402 (class 0 OID 0)
-- Dependencies: 201
-- Name: route_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE route_id_seq OWNED BY route.id;


--
-- TOC entry 185 (class 1259 OID 16639)
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
-- TOC entry 184 (class 1259 OID 16637)
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
-- TOC entry 2403 (class 0 OID 0)
-- Dependencies: 184
-- Name: template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE template_id_seq OWNED BY template.id;


--
-- TOC entry 175 (class 1259 OID 16571)
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
    user_group_ids integer[]
);


ALTER TABLE "user" OWNER TO postgres;

--
-- TOC entry 194 (class 1259 OID 65997)
-- Name: user_group; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE user_group (
    id integer NOT NULL,
    name character varying,
    alias character varying,
    node_permissions jsonb,
    default_node_permissions integer[],
    permission_ids integer[],
    angular_route_ids integer[]
);


ALTER TABLE user_group OWNER TO postgres;

--
-- TOC entry 195 (class 1259 OID 66007)
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
-- TOC entry 2404 (class 0 OID 0)
-- Dependencies: 195
-- Name: user_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_group_id_seq OWNED BY user_group.id;


--
-- TOC entry 174 (class 1259 OID 16569)
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
-- TOC entry 2405 (class 0 OID 0)
-- Dependencies: 174
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- TOC entry 2222 (class 2604 OID 66132)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY adm_route ALTER COLUMN id SET DEFAULT nextval('ang_routes_id_seq'::regclass);


--
-- TOC entry 2209 (class 2604 OID 16619)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content ALTER COLUMN id SET DEFAULT nextval('content_id_seq'::regclass);


--
-- TOC entry 2208 (class 2604 OID 16608)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content_type ALTER COLUMN id SET DEFAULT nextval('content_type_id_seq'::regclass);


--
-- TOC entry 2207 (class 2604 OID 16597)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_type ALTER COLUMN id SET DEFAULT nextval('data_type_id_seq'::regclass);


--
-- TOC entry 2214 (class 2604 OID 57663)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY domain ALTER COLUMN id SET DEFAULT nextval('domain_id_seq'::regclass);


--
-- TOC entry 2216 (class 2604 OID 57778)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member ALTER COLUMN id SET DEFAULT nextval('member_id_seq'::regclass);


--
-- TOC entry 2219 (class 2604 OID 57789)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_group ALTER COLUMN id SET DEFAULT nextval('member_group_id_seq'::regclass);


--
-- TOC entry 2215 (class 2604 OID 57690)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_type ALTER COLUMN id SET DEFAULT nextval('member_type_id_seq'::regclass);


--
-- TOC entry 2224 (class 2604 OID 74366)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY menu_link ALTER COLUMN id SET DEFAULT nextval('menu_link_id_seq'::regclass);


--
-- TOC entry 2210 (class 2604 OID 16630)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY node ALTER COLUMN id SET DEFAULT nextval('node_id_seq'::regclass);


--
-- TOC entry 2221 (class 2604 OID 66021)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY permission ALTER COLUMN id SET DEFAULT nextval('permission_id_seq'::regclass);


--
-- TOC entry 2223 (class 2604 OID 74345)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY route ALTER COLUMN id SET DEFAULT nextval('route_id_seq'::regclass);


--
-- TOC entry 2212 (class 2604 OID 16642)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY template ALTER COLUMN id SET DEFAULT nextval('template_id_seq'::regclass);


--
-- TOC entry 2205 (class 2604 OID 16574)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- TOC entry 2220 (class 2604 OID 66009)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_group ALTER COLUMN id SET DEFAULT nextval('user_group_id_seq'::regclass);


--
-- TOC entry 2375 (class 0 OID 66129)
-- Dependencies: 199
-- Data for Name: adm_route; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY adm_route (id, name, alias, path, parent_id, type, icon, url, components, redirect_to, data, ref) FROM stdin;
2	Content	content	2	\N	1	fa fa-newspaper-o fa-fw	/admin/content	[{"single": "public/views/content/index.html"}]	\N	\N	\N
24	Edit	edit	6.14.24	14	3	\N	/edit/:nodeId	[{"single": "public/views/settings/template/edit.html"}]	\N	\N	\N
23	Edit	edit	6.9.23	9	3	\N	/edit/:nodeId	[{"single": "public/views/settings/content-type/edit.html"}]	\N	\N	\N
25	Edit	edit	6.12.25	12	3	\N	/edit/:name	[{"single": "public/views/settings/script/edit.html"}]	\N	\N	\N
4	Users	users	4	\N	1	fa fa-user fa-fw	/admin/users	[{"single": "public/views/users/index.html"}]	\N	\N	\N
5	Members	members	5	\N	1	fa fa-users fa-fw	/admin/members	[{"single": "public/views/members/index.html"}]	\N	\N	\N
6	Settings	settings	6	\N	1	fa fa-gear fa-fw	/admin/settings	[{"single": "public/views/settings/index.html"}]	\N	\N	\N
7	Login	login	7	\N	1	fa fa-sign-in fa-fw	/admin/login	[{"single": "public/views/admin/admin-login.html"}]	\N	\N	\N
8	Logout	logout	8	\N	1	fa fa-power-off fa-fw	/admin/logout	\N	\N	\N	\N
9	Content Type	contentType	6.9	6	2	fa fa-newspaper-o fa-fw	/content-type	[{"single": "public/views/settings/content-type/index.html"}]	\N	\N	\N
26	Edit	edit	6.13.26	13	3	\N	/edit/:name	[{"single": "public/views/settings/stylesheet/edit.html"}]	\N	\N	\N
12	Scripts	script	6.12	6	2	fa fa-file-code-o fa-fw	/script	[{"single": "public/views/settings/script/index.html"}]	\N	\N	\N
10	Media Types	mediaType	6.10	6	2	fa fa-files-o fa-fw	/media-type	[{"single": "public/views/settings/media-type/index.html"}]	\N	\N	\N
11	Data Types	dataType	6.11	6	2	fa fa-check-square-o fa-fw	/data-type	[{"single": "public/views/settings/data-type/index.html"}]	\N	\N	\N
13	Stylesheets	stylesheet	6.13	6	2	fa fa-desktop fa-fw	/stylesheet	[{"single": "public/views/settings/stylesheet/index.html"}]	\N	\N	\N
14	Templates	template	6.14	6	2	fa fa-eye fa-fw	/template	[{"single": "public/views/settings/template/index.html"}]	\N	\N	\N
15	Member Types	types	5.15	5	2	\N	/member-type	[{"single": "public/views/members/member-type/index.html"}]	\N	\N	\N
27	New	new	2.27	2	3	\N	/new?node_type&content_type_node_id&parent_id	[{"single": "public/views/content/new.html"}]	\N	\N	\N
16	Member Groups	group	5.16	5	2	\N	/member-group	[{"single": "public/views/members/member-group/index.html"}]	\N	\N	\N
17	Edit	edit	2.17	2	3	\N	/edit/:nodeId	[{"single": "public/views/content/edit.html"}]	\N	\N	\N
18	Edit	edit	3.18	3	3	\N	/edit/:nodeId	[{"single": "public/views/media/edit.html"}]	\N	\N	\N
19	Edit	edit	5.19	5	3	\N	/edit/:id	[{"single": "public/views/members/edit.html"}]	\N	\N	\N
21	Edit	edit	6.10.21	10	3	\N	/edit/:nodeId	[{"single": "public/views/settings/media-type/edit.html"}]	\N	\N	\N
22	Edit	edit	6.11.22	11	3	\N	/edit/:nodeId	[{"single": "public/views/settings/data-type/edit.html"}]	\N	\N	\N
28	New	new	3.28	3	3	\N	/new?node_type&content_type_node_id&parent_id	[{"single": "public/views/media/new.html"}]	\N	\N	\N
20	Edit	edit	5.15.20	15	3	\N	/edit/:nodeId	[{"single": "public/views/members/member-type/edit.html"}]	\N	\N	\N
30	New	new	6.10.30	10	3	\N	/new?type&parent	[{"single": "public/views/settings/media-type/new.html"}]	\N	\N	\N
29	New	new	6.9.29	9	3	\N	/new?type&parent	[{"single": "public/views/settings/content-type/new.html"}]	\N	\N	\N
31	New	new	6.11.31	11	3	\N	/new	[{"single": "public/views/settings/data-type/new.html"}]	\N	\N	\N
32	New	new	6.14.32	14	3	\N	/new?parent	[{"single": "public/views/settings/template/new.html"}]	\N	\N	\N
33	New	new	6.12.33	12	3	\N	/new?type&parent	[{"single": "public/views/settings/script/new.html"}]	\N	\N	\N
34	New	new	6.13.34	13	3	\N	/new?type&parent	[{"single": "public/views/settings/stylesheet/new.html"}]	\N	\N	\N
3	Media	media	3	\N	1	fa fa-file-image-o fa-fw	/admin/media	[{"single": "public/views/media/index.html"}]	\N	\N	\N
1	Index	index	1	\N	1	fa fa-home fa-fw	/admin	[{"single": "public/views/admin/dashboard.html"}]	\N	\N	\N
\.


--
-- TOC entry 2406 (class 0 OID 0)
-- Dependencies: 198
-- Name: ang_routes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('ang_routes_id_seq', 34, true);


--
-- TOC entry 2357 (class 0 OID 16616)
-- Dependencies: 181
-- Data for Name: content; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY content (id, node_id, content_type_node_id, meta, public_access) FROM stdin;
13	33	5	{"prop2": "prop2", "prop3": "prop3", "page_title": "About me", "page_content": "About page description", "template_node_id": 22}	\N
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
2	10	5	{"prop2": "prop2a", "prop3": "sample page prop 3", "page_title": "Sample page title", "page_content": "Sample page content goes here", "template_node_id": 22}	\N
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
-- TOC entry 2407 (class 0 OID 0)
-- Dependencies: 180
-- Name: content_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_id_seq', 32, true);


--
-- TOC entry 2355 (class 0 OID 16605)
-- Dependencies: 179
-- Data for Name: content_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY content_type (id, node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) FROM stdin;
4	15	Image	Image content type description	fa fa-folder-o	fa fa-folder-o	\N	null	[{"name":"Image","properties":[{"name":"path","order":1,"data_type_node_id":2,"help_text":"help text","description":"URL goes here."},{"name":"title","order":2,"data_type_node_id":2,"help_text":"help text","description":"The title entered here can override the above one."},{"name":"caption","order":3,"data_type_node_id":14,"help_text":"help text","description":"Caption goes here."},{"name":"alt","order":4,"data_type_node_id":14,"help_text":"help text","description":"Alt goes here."},{"name":"description","order":5,"data_type_node_id":14,"help_text":"help text","description":"Description goes here."},{"name":"file_upload","order":1,"data_type_node_id":48,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties","properties":[{"name":"temporary property","order":1,"data_type_node_id":2,"help_text":"help text","description":"Temporary description goes here."}]}]
1	3	Master	Some description	fa fa-folder-o	fa fa-folder-o	\N	\N	[{"name": "Content", "properties": [{"name": "page_title", "order": 1, "data_type_node_id": 2, "help_text": "help text", "description": "The page title overrides the name the page has been given."}]}, {"name": "Properties", "properties": [{"name": "prop2", "order": 1, "data_type_node_id": 2, "help_text": "help text2", "description": "description2"}, {"name": "prop3", "order": 2, "data_type_node_id": 2, "help_text": "help text3", "description": "description3"}]}]
5	28	ctTestContentTypeAlias	ct-test desc	ct-test-icon	ct-test-thumbnail	3	{"template_node_id": 8, "allowed_templates_node_id": [8, 22, 25], "allowed_content_types_node_id": [5]}	[{"name":"Mytab1","properties":[{"name":"property name1","order":1,"data_type_node_id":2,"help_text":"prop help text1","description":"prop description1"},{"name":"property name2","order":2,"data_type_node_id":14,"help_text":"prop help text2","description":"prop description2"}]},{"name":"Mytab2","properties":[{"name":"property name3","order":1,"data_type_node_id":27,"help_text":"prop help text3","description":"prop description3"},{"name":"property name4","order":2,"data_type_node_id":14,"help_text":"prop help text4","description":"prop description4"}]}]
3	5	ctPage	Page content type desc	fa fa-folder-o	fa fa-folder-o	3	{"template_node_id": 8, "allowed_templates_node_id": [8, 22, 25], "allowed_content_types_node_id": [5]}	[{"name":"Content","properties":[{"name":"page_content","order":2,"data_type_node_id":14,"help_text":"Help text for page contentent","description":"Page content description"}]}]
7	35	mtTestMediaType alias	mtTest desc	mtTest-icon	mtTest-thumbnail	16	{"allowed_content_types_node_id": [15, 17, 16]}	[{"name":"mytab"}]
8	43	CT test alias	ct-test-desc	ct-test-icon	ct-test-thumb	3	{"template_node_id": "8", "allowed_templates_node_id": [22, 8, 25, 42], "allowed_content_types_node_id": [5, 28]}	[{"name":"mytab"}]
9	44	Test Content Type 2 alias	tc2-desc	tc2-icon	tc2-thumb	3	{"template_node_id": "8", "allowed_templates_node_id": [22, 8, 25, 26, 42], "allowed_content_types_node_id": [28, 43, 5]}	[{"name":"mytab"}]
6	16	mtFolder	Folder media type description1	mt-icon1	mt-thumbnail1	\N	{"allowed_content_types_node_id": [16, 15]}	[{"name":"Folder","properties":[{"name":"folder_browser","order":1,"data_type_node_id":34,"help_text":"prop help text","description":"prop description"},{"name":"path","order":1,"data_type_node_id":2,"help_text":"prop help text","description":"prop description"}]},{"name":"Properties"}]
2	4	ctHome	Home Some description	fa fa-folder-o	fa fa-folder-o	3	{"template_node_id": 7, "allowed_templates_node_id": [7], "allowed_content_types_node_id": [5, 28]}	[{"name":"Content","properties":[{"name":"site_name","order":2,"data_type_node_id":2,"help_text":"help text","description":"Site name goes here."},{"name":"site_tagline","order":3,"data_type_node_id":2,"help_text":"help text","description":"Site tagline goes here."},{"name":"copyright","order":4,"data_type_node_id":2,"help_text":"help text","description":"Copyright here."},{"name":"domains","order":5,"data_type_node_id":59,"help_text":"help text","description":"Domains goes here."}]},{"name":"Social","properties":[{"name":"facebook","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your facebook link here."},{"name":"twitter","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your twitter link here."},{"name":"linkedin","order":1,"data_type_node_id":2,"help_text":"help text","description":"Enter your linkedin link here."}]}]
\.


--
-- TOC entry 2408 (class 0 OID 0)
-- Dependencies: 178
-- Name: content_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('content_type_id_seq', 9, true);


--
-- TOC entry 2353 (class 0 OID 16594)
-- Dependencies: 177
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
-- TOC entry 2409 (class 0 OID 0)
-- Dependencies: 176
-- Name: data_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_type_id_seq', 6, true);


--
-- TOC entry 2362 (class 0 OID 57658)
-- Dependencies: 186
-- Data for Name: domain; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY domain (id, node_id, name) FROM stdin;
1	9	localhost:8080
\.


--
-- TOC entry 2410 (class 0 OID 0)
-- Dependencies: 187
-- Name: domain_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('domain_id_seq', 1, true);


--
-- TOC entry 2367 (class 0 OID 57775)
-- Dependencies: 191
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member (id, username, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_node_id, group_ids) FROM stdin;
1	default_member	$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa	default_member@mail.com	{"comments": "default user comments"}	2015-01-22 14:25:38.904	\N	2015-02-19 23:46:00.495	\N	1	GIWES3RHMY5RKC7OZPOQTF5FQFWX32D5VLV3CAKT4HGKP5LZIENA	61	{1}
\.


--
-- TOC entry 2368 (class 0 OID 57784)
-- Dependencies: 192
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member_group (id, name, description) FROM stdin;
1	authenticated_member	All logged in members
\.


--
-- TOC entry 2411 (class 0 OID 0)
-- Dependencies: 193
-- Name: member_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_group_id_seq', 1, true);


--
-- TOC entry 2412 (class 0 OID 0)
-- Dependencies: 190
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_id_seq', 1, true);


--
-- TOC entry 2364 (class 0 OID 57676)
-- Dependencies: 188
-- Data for Name: member_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY member_type (id, node_id, alias, description, icon, parent_member_type_node_id, meta, tabs) FROM stdin;
1	61	mtMember	Default member type	fa fa-user fa-fw	1	\N	[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_node_id": 14}]}]
\.


--
-- TOC entry 2413 (class 0 OID 0)
-- Dependencies: 189
-- Name: member_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_type_id_seq', 1, true);


--
-- TOC entry 2379 (class 0 OID 74363)
-- Dependencies: 203
-- Data for Name: menu_link; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY menu_link (id, path, name, parent_id, route_id, icon, atts, type, menu, user_ids, user_group_ids) FROM stdin;
1	1	Index	\N	1	fa fa-home fa-fw	\N	1	main	\N	{1,2,3}
2	2	Content	\N	2	fa fa-newspaper-o fa-fw	\N	1	main	\N	{1,2,3}
3	3	Media	\N	3	fa fa-file-image-o fa-fw	\N	1	main	\N	{1,2,3}
4	4	Users	\N	4	fa fa-user fa-fw	\N	1	main	\N	{1,2,3}
5	5	Members	\N	5	fa fa-users fa-fw	\N	1	main	\N	{1,2,3}
8	6.8	Media Types	6	12	fa fa-files-o fa-fw	\N	1	main	\N	{1,2,3}
9	6.9	Data Types	6	13	fa fa-check-square-o fa-fw	\N	1	main	\N	{1,2,3}
10	6.10	Templates	6	14	fa fa-eye fa-fw	\N	1	main	\N	{1,2,3}
11	6.11	Scripts	6	15	fa fa-file-code-o fa-fw	\N	1	main	\N	{1,2,3}
12	6.12	Stylesheets	6	16	fa fa-desktop fa-fw	\N	1	main	\N	{1,2,3}
13	5.13	Member Types	5	31	fa fa-smile-o fa-fw	\N	1	main	\N	{1,2,3}
6	6	Settings	\N	6	fa fa-gear fa-fw	\N	1	main	\N	{1,2,3}
7	6.7	Content Types	6	11	fa fa-newspaper-o fa-fw	\N	1	main	\N	{1,2,3}
\.


--
-- TOC entry 2414 (class 0 OID 0)
-- Dependencies: 202
-- Name: menu_link_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('menu_link_id_seq', 13, true);


--
-- TOC entry 2359 (class 0 OID 16627)
-- Dependencies: 183
-- Data for Name: node; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY node (id, path, name, node_type, created_by, created_date, parent_id, user_permissions, user_group_permissions) FROM stdin;
14	1.14	Textarea	11	1	2014-10-27 02:40:41.179	1	\N	\N
48	1.48	File upload	11	1	2014-12-05 19:56:17.883	\N	\N	\N
60	1.6.60	404	3	1	2015-01-20 13:46:33.668	6	\N	\N
20	1.20	Sidebar 1	3	1	2014-11-10 09:03:20.514	1	\N	\N
15	1.15	Image	7	1	2014-10-28 15:16:25.972	1	\N	\N
36	1.36	Another Image Folder2	2	1	2014-12-02 01:00:51.206	1	\N	\N
21	1.21	Sidebar 2	3	1	2014-11-10 23:56:55.038	1	\N	\N
1	1	root	5	1	2014-10-22 16:51:00.215	\N	\N	\N
2	1.2	Text input	11	1	2014-10-22 16:51:00.215	1	\N	\N
3	1.3	Master	4	1	2014-10-22 16:51:00.215	1	\N	\N
13	1.9.13	Another Page	1	1	2014-10-26 23:27:14.571	9	\N	\N
9	1.9	Home	1	1	2014-10-22 16:51:00.215	1	\N	\N
61	1.61	Member	12	1	2015-01-22 15:55:13.957	1	\N	\N
28	1.3.28	Test Content Type	4	1	2014-11-26 04:20:48.026	3	\N	\N
26	1.6.25.26	Child of test template	3	1	2014-11-26 01:39:42.816	25	\N	\N
46	1.37.46	Level 1	2	1	2014-12-05 17:02:13.875	37	\N	\N
41	1.36.41	gopher.jpg	2	1	2014-12-02 02:08:26.737	36	\N	\N
17	1.17	File	7	1	2014-10-28 15:18:13.4	1	\N	\N
18	1.18	gopher.jpg	2	1	2014-10-28 15:50:47.303	1	\N	\N
16	1.16	Folder	7	1	2014-10-28 15:18:13.4	1	\N	\N
35	1.35	Media Type Test	7	1	2014-12-01 22:09:43.783	1	\N	\N
23	1.23	Sample picture folder	2	1	2014-11-17 16:57:14.654	1	\N	\N
47	1.37.46.47	Level 2	2	1	2014-12-05 17:02:46.762	46	\N	\N
43	1.3.43	Content Type Test	4	1	2014-12-02 12:38:59.527	3	\N	\N
55	1.38.45.55	cat-prays.jpg	2	1	2014-12-06 13:07:08.943	45	\N	\N
54	1.37.46.54	catduck.jpg	2	1	2014-12-06 03:44:40.07	46	\N	\N
44	1.3.44	Test Content Type 2	4	1	2014-12-02 12:48:25.307	3	\N	\N
58	1.6.58	Unauthorized	3	1	2014-12-15 14:24:22.063	6	\N	\N
34	1.34	Folder Browser	11	1	2014-12-01 16:09:46.488	1	\N	\N
19	1.19	postgresql.png	2	1	2014-10-28 17:53:37.488	1	\N	\N
27	1.27	Color Picker	11	1	2014-11-26 02:20:17.638	1	\N	\N
33	1.9.33	About	1	1	2014-12-01 12:11:25.838	9	\N	\N
24	1.23.24	Goku_SSJ3.jpg	2	1	2014-11-17 16:58:57.285	23	\N	\N
10	1.9.10	Sample Page	1	1	2014-10-22 16:51:00.215	9	[{"id": 2, "permissions": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]}]	\N
11	1.9.10.11	Child Page Level 1	1	1	2014-10-26 23:19:44.735	10	\N	[{"id": 1, "permissions": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]}]
4	1.3.4	Home	4	1	2014-10-22 16:51:00.215	3	\N	\N
31	1.9.30.31	MySubPage	1	1	2014-12-01 12:02:54.252	30	\N	\N
32	1.9.32	Yet another test page	1	1	2014-12-01 12:07:29.999	9	\N	\N
40	1.9.39.40	Test page 2 child	1	1	2014-12-02 01:45:49.78	39	\N	\N
39	1.9.39	Testpage 2	1	1	2014-12-02 01:43:33.233	9	\N	\N
5	1.3.5	Page	4	1	2014-10-22 16:51:00.215	3	\N	\N
51	1.37.46.47.51	dog.jpg	2	1	2014-12-05 21:06:49.532	47	\N	\N
52	1.37.46.47.52	taco-hamster.jpg	2	1	2014-12-05 21:22:45.227	47	\N	\N
22	1.6.22	Page with sidebars	3	1	2014-11-11 03:39:55.766	6	\N	\N
49	1.37.46.47.49	blomkals-hamster.jpg	2	1	2014-12-05 20:44:25.921	47	\N	\N
38	1.38	2014	2	1	2014-12-02 01:42:09.979	1	\N	\N
50	1.37.46.47.50	tiny.jpg	2	1	2014-12-05 21:05:42.816	47	\N	\N
57	1.38.45.57	sleeping-kitten.jpg	2	1	2014-12-06 14:28:52.117	45	\N	\N
37	1.37	Subfolder depth test3	2	1	2014-12-02 01:37:09.125	1	\N	\N
8	1.6.8	Page	3	1	2014-10-22 16:51:00.215	6	\N	\N
45	1.38.45	12	2	1	2014-12-05 16:18:29.762	38	\N	\N
6	1.6	Layout	3	1	2014-10-22 16:51:00.215	1	\N	\N
56	1.38.45.56	ducks.jpg	2	1	2014-12-06 13:10:14.637	45	\N	\N
53	1.37.53	AngularLogo.png	2	1	2014-12-06 03:36:14.425	37	\N	\N
42	1.6.42	Test template 2	3	1	2014-12-02 02:19:29.241	6	\N	\N
7	1.6.7	Home	3	1	2014-10-22 16:51:00.215	6	\N	\N
29	1.9.29	Test Page	1	1	2014-12-01 11:45:16.186	9	\N	\N
30	1.9.30	Login	1	1	2014-12-01 11:54:10.208	9	\N	\N
59	1.59	Domains	11	1	2015-01-19 21:22:06.945	\N	\N	\N
12	1.9.10.11.12	Child Page Level 2	1	1	2014-10-26 23:19:44.735	11	\N	\N
25	1.6.25	Login	3	1	2014-11-26 00:13:45.309	6	\N	\N
\.


--
-- TOC entry 2415 (class 0 OID 0)
-- Dependencies: 182
-- Name: node_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('node_id_seq', 61, true);


--
-- TOC entry 2372 (class 0 OID 66016)
-- Dependencies: 196
-- Data for Name: permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY permission (id, name) FROM stdin;
1	Create
2	Delete
3	Update
4	Move
5	Copy
7	Public access
8	Permissions
9	Send to publish
10	Sort
11	Publish
6	Change content type
12	Browse Node
\.


--
-- TOC entry 2416 (class 0 OID 0)
-- Dependencies: 197
-- Name: permission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('permission_id_seq', 12, true);


--
-- TOC entry 2376 (class 0 OID 74340)
-- Dependencies: 200
-- Data for Name: route; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY route (id, path, name, parent_id, url, components, is_abstract) FROM stdin;
6	6	settings	\N	/admin/settings	[{"single": "public/views/settings/index.html"}]	t
1	1	index	\N	/admin	[{"single": "public/views/admin/dashboard.html"}]	f
2	2	content	\N	/admin/content	[{"single": "public/views/content/index.html"}]	f
3	3	media	\N	/admin/media	[{"single": "public/views/media/index.html"}]	f
4	4	users	\N	/admin/users	[{"single": "public/views/users/index.html"}]	f
5	5	members	\N	/admin/members	[{"single": "public/views/members/index.html"}]	f
7	2.7	new	2	/new?node_type&content_type_node_id&parent_id	[{"single": "public/views/content/new.html"}]	f
8	2.8	edit	2	/edit/:nodeId	[{"single": "public/views/content/edit.html"}]	f
9	3.9	new	3	/new?node_type&content_type_node_id&parent_id	[{"single": "public/views/media/new.html"}]	f
10	3.10	edit	3	/edit/:nodeId	[{"single": "public/views/media/edit.html"}]	f
11	6.11	contentTypes	6	/content-type	[{"single": "public/views/settings/content-type/index.html"}]	f
12	6.12	mediaTypes	6	/media-type	[{"single": "public/views/settings/media-type/index.html"}]	f
13	6.13	dataTypes	6	/data-type	[{"single": "public/views/settings/data-type/index.html"}]	f
14	6.14	templates	6	/template	[{"single": "public/views/settings/template/index.html"}]	f
15	6.15	scripts	6	/script	[{"single": "public/views/settings/script/index.html"}]	f
16	6.16	stylesheets	6	/stylesheet	[{"single": "public/views/settings/stylesheet/index.html"}]	f
17	6.11.17	new	11	/new?type&parent	[{"single": "public/views/settings/content-type/new.html"}]	f
18	6.12.18	new	12	/new?type&parent	[{"single": "public/views/settings/media-type/new.html"}]	f
19	6.13.19	new	13	/new	[{"single": "public/views/settings/data-type/new.html"}]	f
20	6.14.20	new	14	/new?parent	[{"single": "public/views/settings/template/new.html"}]	f
21	6.15.21	new	15	/new?type&parent	[{"single": "public/views/settings/script/new.html"}]	f
22	6.16.22	new	16	/new?type&parent	[{"single": "public/views/settings/stylesheet/new.html"}]	f
23	6.11.23	edit	11	/edit/:nodeId	[{"single": "public/views/settings/content-type/edit.html"}]	f
24	6.12.24	edit	12	/edit/:nodeId	[{"single": "public/views/settings/media-type/edit.html"}]	f
25	6.13.25	edit	13	/edit/:nodeId	[{"single": "public/views/settings/data-type/edit.html"}]	f
26	6.14.26	edit	14	/edit/:nodeId	[{"single": "public/views/settings/template/edit.html"}]	f
27	6.15.27	edit	15	/edit/:name	[{"single": "public/views/settings/script/edit.html"}]	f
28	6.16.28	edit	16	/edit/:name	[{"single": "public/views/settings/stylesheet/edit.html"}]	f
29	5.29	edit	5	/edit/:id	[{"single": "public/views/members/edit.html"}]	f
30	5.30	new	5	/new	[{"single": "public/views/members/new.html"}]	f
31	5.31	memberTypes	5	/member-type	[{"single": "public/views/members/member-type/index.html"}]	f
32	5.31.32	edit	31	/edit/:nodeId	[{"single": "public/views/members/member-type/edit.html"}]	f
33	5.31.33	new	31	/new?type&parent	[{"single": "public/views/members/member-type/new.html"}]	f
\.


--
-- TOC entry 2417 (class 0 OID 0)
-- Dependencies: 201
-- Name: route_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('route_id_seq', 33, true);


--
-- TOC entry 2361 (class 0 OID 16639)
-- Dependencies: 185
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
-- TOC entry 2418 (class 0 OID 0)
-- Dependencies: 184
-- Name: template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('template_id_seq', 12, true);


--
-- TOC entry 2351 (class 0 OID 16571)
-- Dependencies: 175
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids) FROM stdin;
1	soren	Soren	Tester	$2a$10$UNrly6WSmQnm495KAth6Auk4Z.11kjDBRFz8ZKjhqthytKFH/TjKq	soren@codeish.com	2014-10-21 16:51:00.215	\N	2015-02-20 00:12:14.354	\N	1	YNHWOAMEFEOYDIQM66TBRQ7I45LR7FJQFT7FPDULDJXTWEFE2U2Q	{1}
2	admin	Admin	Demo	$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa	demo@codeish.com	2014-11-15 16:51:00.215	\N	2015-02-20 03:20:04.413	\N	1	CPNMASEFS223TVDTWZ6H7UZJ4XRA7J2WEGBI7TA4NNSHTIDJ3VBQ	{1}
\.


--
-- TOC entry 2370 (class 0 OID 65997)
-- Dependencies: 194
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY user_group (id, name, alias, node_permissions, default_node_permissions, permission_ids, angular_route_ids) FROM stdin;
2	Editor	editor	[{"id": 10, "permissions": [1, 9, 12]}, {"id": 11, "permissions": [1, 9]}]	{1,2,3,4,5,7,9,10,12}	{1,2,3,4,5,7,9,10,12}	\N
3	Writer	writer	[{"id": 10, "permissions": [1, 9]}, {"id": 11, "permissions": [1, 9, 12]}]	{1,3,9,12}	{1,3,9,12}	\N
1	Administrator	administrator	\N	{1,2,3,4,5,6,7,8,9,10,11,12}	{1,2,3,4,5,6,7,8,9,10,11,12}	{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34}
\.


--
-- TOC entry 2419 (class 0 OID 0)
-- Dependencies: 195
-- Name: user_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_group_id_seq', 3, true);


--
-- TOC entry 2420 (class 0 OID 0)
-- Dependencies: 174
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 2, true);


--
-- TOC entry 2232 (class 2606 OID 16613)
-- Name: content_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY content_type
    ADD CONSTRAINT content_type_pkey PRIMARY KEY (id);


--
-- TOC entry 2230 (class 2606 OID 16602)
-- Name: data_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY data_type
    ADD CONSTRAINT data_type_pkey PRIMARY KEY (id);


--
-- TOC entry 2234 (class 2606 OID 16624)
-- Name: document_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY content
    ADD CONSTRAINT document_pkey PRIMARY KEY (id);


--
-- TOC entry 2237 (class 2606 OID 16636)
-- Name: node_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY node
    ADD CONSTRAINT node_pkey PRIMARY KEY (id);


--
-- TOC entry 2240 (class 2606 OID 16647)
-- Name: template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY template
    ADD CONSTRAINT template_pkey PRIMARY KEY (id);


--
-- TOC entry 2226 (class 2606 OID 49500)
-- Name: user_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- TOC entry 2228 (class 2606 OID 16579)
-- Name: user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 2235 (class 1259 OID 41272)
-- Name: idxgin; Type: INDEX; Schema: public; Owner: postgres; Tablespace: 
--

CREATE INDEX idxgin ON content USING gin (meta);


--
-- TOC entry 2238 (class 1259 OID 41273)
-- Name: template_partial_template_node_ids_idx; Type: INDEX; Schema: public; Owner: postgres; Tablespace: 
--

CREATE INDEX template_partial_template_node_ids_idx ON template USING gin (partial_template_node_ids);


--
-- TOC entry 2386 (class 0 OID 0)
-- Dependencies: 5
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2015-02-20 12:05:18

--
-- PostgreSQL database dump complete
--

