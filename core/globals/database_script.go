package globals

var DbCreateScriptDML string = `--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-06-29 12:25:35

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- TOC entry 202 (class 3079 OID 11855)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2368 (class 0 OID 0)
-- Dependencies: 202
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 204 (class 3079 OID 97715)
-- Name: citext; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- TOC entry 2369 (class 0 OID 0)
-- Dependencies: 204
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- TOC entry 203 (class 3079 OID 97799)
-- Name: ltree; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;


--
-- TOC entry 2370 (class 0 OID 0)
-- Dependencies: 203
-- Name: EXTENSION ltree; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';


SET search_path = public, pg_catalog;

--
-- TOC entry 324 (class 1255 OID 97974)
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
-- TOC entry 325 (class 1255 OID 97975)
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


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 195 (class 1259 OID 98958)
-- Name: content; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE content (
    id bigint NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now(),
    content_type_id bigint,
    template_id bigint,
    meta jsonb,
    public_access_members jsonb,
    public_access_member_groups jsonb,
    user_permissions jsonb,
    user_group_permissions jsonb,
    is_abstract boolean DEFAULT false NOT NULL
);


ALTER TABLE content OWNER TO postgres;


--
-- TOC entry 194 (class 1259 OID 98956)
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
-- TOC entry 2372 (class 0 OID 0)
-- Dependencies: 194
-- Name: content_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_id_seq OWNED BY content.id;


--
-- TOC entry 193 (class 1259 OID 98944)
-- Name: content_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE content_type (
    id bigint NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    alias character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now(),
    description text,
    icon character varying,
    thumbnail character varying,
    meta jsonb,
    tabs jsonb,
    allow_at_root boolean DEFAULT false NOT NULL,
    is_container boolean DEFAULT false NOT NULL,
    is_abstract boolean DEFAULT false NOT NULL,
    allowed_content_type_ids bigint[],
    composite_content_type_ids bigint[],
    template_id bigint,
    allowed_template_ids bigint[]
);


ALTER TABLE content_type OWNER TO postgres;

--
-- TOC entry 192 (class 1259 OID 98942)
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
-- TOC entry 2374 (class 0 OID 0)
-- Dependencies: 192
-- Name: content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_type_id_seq OWNED BY content_type.id;


--
-- TOC entry 180 (class 1259 OID 98158)
-- Name: data_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE data_type (
    id bigint NOT NULL,
    name character varying,
    alias character varying,
    created_by bigint DEFAULT 0 NOT NULL,
    created_date timestamp without time zone DEFAULT now(),
    html text,
    editor_alias character varying,
    meta jsonb
);


ALTER TABLE data_type OWNER TO postgres;

--
-- TOC entry 181 (class 1259 OID 98162)
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
-- TOC entry 2375 (class 0 OID 0)
-- Dependencies: 181
-- Name: data_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_type_id_seq OWNED BY data_type.id;


--
-- TOC entry 185 (class 1259 OID 98724)
-- Name: media; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE media (
    id bigint NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now(),
    media_type_id bigint,
    meta jsonb,
    user_permissions jsonb,
    user_group_permissions jsonb
);


ALTER TABLE media OWNER TO postgres;

--
-- TOC entry 184 (class 1259 OID 98722)
-- Name: media_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE media_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media_id_seq OWNER TO postgres;

--
-- TOC entry 2376 (class 0 OID 0)
-- Dependencies: 184
-- Name: media_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE media_id_seq OWNED BY media.id;


--
-- TOC entry 187 (class 1259 OID 98772)
-- Name: media_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE media_type (
    id bigint NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    alias character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now(),
    description text,
    icon character varying,
    thumbnail character varying,
    meta jsonb,
    tabs jsonb,
    allow_at_root boolean DEFAULT false NOT NULL,
    is_container boolean DEFAULT false NOT NULL,
    is_abstract boolean DEFAULT false NOT NULL,
    allowed_media_type_ids bigint[],
    composite_media_type_ids bigint[]
);


ALTER TABLE media_type OWNER TO postgres;

--
-- TOC entry 186 (class 1259 OID 98770)
-- Name: media_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE media_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE media_type_id_seq OWNER TO postgres;

--
-- TOC entry 2377 (class 0 OID 0)
-- Dependencies: 186
-- Name: media_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE media_type_id_seq OWNED BY media_type.id;


--
-- TOC entry 172 (class 1259 OID 98009)
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
    member_type_id bigint,
    member_group_ids integer[]
);


ALTER TABLE member OWNER TO postgres;

--
-- TOC entry 179 (class 1259 OID 98131)
-- Name: member_group; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE member_group (
    id bigint NOT NULL,
    name character varying,
    alias character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now()
);


ALTER TABLE member_group OWNER TO postgres;

--
-- TOC entry 178 (class 1259 OID 98129)
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
-- TOC entry 2378 (class 0 OID 0)
-- Dependencies: 178
-- Name: member_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_group_id_seq OWNED BY member_group.id;


--
-- TOC entry 173 (class 1259 OID 98017)
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
-- TOC entry 2379 (class 0 OID 0)
-- Dependencies: 173
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_id_seq OWNED BY member.id;


--
-- TOC entry 197 (class 1259 OID 98969)
-- Name: member_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE member_type (
    id integer NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    alias character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now(),
    description text,
    icon character varying,
    thumbnail character varying,
    meta jsonb,
    tabs jsonb,
    is_abstract boolean DEFAULT false NOT NULL,
    composite_member_type_ids bigint[]
);


ALTER TABLE member_type OWNER TO postgres;

--
-- TOC entry 196 (class 1259 OID 98967)
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
-- TOC entry 2380 (class 0 OID 0)
-- Dependencies: 196
-- Name: member_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_type_id_seq OWNED BY member_type.id;


--
-- TOC entry 200 (class 1259 OID 99012)
-- Name: mime_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE mime_type (
    id bigint NOT NULL,
    name character varying,
    media_type_id bigint
);


ALTER TABLE mime_type OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 99015)
-- Name: mime_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE mime_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE mime_type_id_seq OWNER TO postgres;

--
-- TOC entry 2381 (class 0 OID 0)
-- Dependencies: 201
-- Name: mime_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE mime_type_id_seq OWNED BY mime_type.id;


--
-- TOC entry 199 (class 1259 OID 99001)
-- Name: permission; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE permission (
    id bigint NOT NULL,
    name character varying
);


ALTER TABLE permission OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 98999)
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
-- TOC entry 2382 (class 0 OID 0)
-- Dependencies: 198
-- Name: permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE permission_id_seq OWNED BY permission.id;


--
-- TOC entry 183 (class 1259 OID 98211)
-- Name: template; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE template (
    id bigint NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    alias character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now(),
    is_partial boolean DEFAULT false
);


ALTER TABLE template OWNER TO postgres;

--
-- TOC entry 182 (class 1259 OID 98209)
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
-- TOC entry 2383 (class 0 OID 0)
-- Dependencies: 182
-- Name: template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE template_id_seq OWNED BY template.id;


--
-- TOC entry 174 (class 1259 OID 98067)
-- Name: user; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE "user" (
    id bigint NOT NULL,
    username character varying NOT NULL,
    first_name character varying,
    last_name character varying,
    password character(60) NOT NULL,
    email character varying,
    created_date timestamp without time zone DEFAULT now(),
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
-- TOC entry 175 (class 1259 OID 98074)
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
-- TOC entry 176 (class 1259 OID 98080)
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
-- TOC entry 2384 (class 0 OID 0)
-- Dependencies: 176
-- Name: user_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_group_id_seq OWNED BY user_group.id;


--
-- TOC entry 177 (class 1259 OID 98082)
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
-- TOC entry 2385 (class 0 OID 0)
-- Dependencies: 177
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- TOC entry 2238 (class 2604 OID 98961)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content ALTER COLUMN id SET DEFAULT nextval('content_id_seq'::regclass);


--
-- TOC entry 2233 (class 2604 OID 98947)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content_type ALTER COLUMN id SET DEFAULT nextval('content_type_id_seq'::regclass);


--
-- TOC entry 2212 (class 2604 OID 98164)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_type ALTER COLUMN id SET DEFAULT nextval('data_type_id_seq'::regclass);


--
-- TOC entry 2218 (class 2604 OID 98727)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY media ALTER COLUMN id SET DEFAULT nextval('media_id_seq'::regclass);


--
-- TOC entry 2220 (class 2604 OID 98775)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY media_type ALTER COLUMN id SET DEFAULT nextval('media_type_id_seq'::regclass);


--
-- TOC entry 2205 (class 2604 OID 98088)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member ALTER COLUMN id SET DEFAULT nextval('member_id_seq'::regclass);


--
-- TOC entry 2210 (class 2604 OID 98134)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_group ALTER COLUMN id SET DEFAULT nextval('member_group_id_seq'::regclass);


--
-- TOC entry 2241 (class 2604 OID 98972)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_type ALTER COLUMN id SET DEFAULT nextval('member_type_id_seq'::regclass);


--
-- TOC entry 2245 (class 2604 OID 99017)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY mime_type ALTER COLUMN id SET DEFAULT nextval('mime_type_id_seq'::regclass);


--
-- TOC entry 2244 (class 2604 OID 99004)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY permission ALTER COLUMN id SET DEFAULT nextval('permission_id_seq'::regclass);


--
-- TOC entry 2215 (class 2604 OID 98214)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY template ALTER COLUMN id SET DEFAULT nextval('template_id_seq'::regclass);


--
-- TOC entry 2207 (class 2604 OID 98094)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- TOC entry 2209 (class 2604 OID 98095)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_group ALTER COLUMN id SET DEFAULT nextval('user_group_id_seq'::regclass);


--
-- TOC entry 2251 (class 2606 OID 99009)
-- Name: permission_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY permission
    ADD CONSTRAINT permission_name_key UNIQUE (name);


--
-- TOC entry 2247 (class 2606 OID 98109)
-- Name: user_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- TOC entry 2249 (class 2606 OID 98111)
-- Name: user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 2367 (class 0 OID 0)
-- Dependencies: 6
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2015-06-29 12:25:36

--
-- PostgreSQL database dump complete
--
`

var DbCreateScriptDDL string = `--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-06-29 12:26:22

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

SET search_path = public, pg_catalog;


--
-- TOC entry 2363 (class 0 OID 98067)
-- Dependencies: 174
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO "user" (id, username, first_name, last_name, password, email, created_date, updated_date, login_date, accessed_date, status, sid, user_group_ids, permissions) VALUES (1, '%s', 'Admin', 'Demo', '%s', '%s', '2014-11-15 16:51:00.215', NULL, '2015-06-26 18:45:51.652', NULL, 1, 'Q6BZMGO6GTBLVLFW2Z2PHZSVLISDGGH6GVKNLS2V2ZS44LXT4ZZA', '{1}', NULL);


--
-- TOC entry 2364 (class 0 OID 98074)
-- Dependencies: 175
-- Data for Name: user_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO user_group (id, name, alias, permissions) VALUES (2, 'Editor', 'editor', '{}');
INSERT INTO user_group (id, name, alias, permissions) VALUES (3, 'Writer', 'writer', '{}');
INSERT INTO user_group (id, name, alias, permissions) VALUES (1, 'Administrator', 'administrator', '{all,content_create,content_delete,content_update,content_section,content_browse,media_all,media_create,media_delete,media_update,media_section,media_browse,user_all,user_create,user_delete,user_update,user_section,user_browse,user_type_all,user_type_create,user_type_delete,user_type_update,user_type_section,user_type_browse,user_group_all,user_group_create,user_group_delete,user_group_update,user_group_section,user_group_browse,permission_all,permission_create,permission_delete,permission_update,permission_section,permission_browse,asset_all,asset_create,asset_delete,asset_update,asset_browse,asset_section,member_all,member_create,member_delete,member_update,member_section,member_browse,member_type_all,member_type_create,member_type_delete,member_type_update,member_type_section,member_type_browse,member_group_all,member_group_create,member_group_delete,member_group_update,member_group_section,member_group_browse,template_all,template_create,template_delete,template_update,template_section,template_browse,settings_section,settings_all,script_all,script_create,script_delete,script_update,script_section,script_browse,stylesheet_all,stylesheet_create,stylesheet_delete,stylesheet_update,stylesheet_section,stylesheet_browse,content_type_all,content_type_create,content_type_delete,content_type_update,content_type_section,content_type_browse,media_type_all,media_type_create,media_type_delete,media_type_update,media_type_section,media_type_browse,data_type_all,data_type_create,data_type_delete,data_type_update,data_type_section,data_type_browse,node_create,node_delete,node_update,node_move,node_copy,node_public_access,node_permissions,node_send_to_publish,node_publish,node_browse,node_change_content_type,node_sort,content_all,mime_type_all,mime_type_create,mime_type_delete,mime_type_update,mime_type_section,mime_type_browse}');


--
-- TOC entry 2408 (class 0 OID 0)
-- Dependencies: 176
-- Name: user_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_group_id_seq', 3, true);


--
-- TOC entry 2409 (class 0 OID 0)
-- Dependencies: 177
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('user_id_seq', 1, true);


--
-- TOC entry 2369 (class 0 OID 98158)
-- Dependencies: 180
-- Data for Name: data_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (2, 'Numeric Input', 'numeric_input', 0, '2015-03-26 23:47:44.854', '<input type="number" id="{{prop.name}}" ng-model="data.meta[prop.name]">', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (3, 'Textarea', 'textarea', 0, '2015-03-26 23:47:44.854', '<textarea id="{{prop.name}}" ng-model="data.meta[prop.name]">', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (18, 'True/False', 'true_false', 0, '2015-03-26 23:47:44.854', '<div><label><input type="checkbox" 
        ng-model="data.meta[prop.name]"
        ng-true-value="true"
        ng-false-value="false"
       ></label> {{prop.name}}
</div>', '', NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (5, 'Dropdown', 'dropdown', 0, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (6, 'Dropdown Multiple', 'dropdown_multiple', 0, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (8, 'Label', 'label', 0, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (9, 'Color Picker', 'color_picker', 0, '2015-03-26 23:47:44.854', '<colorpicker>The color picker data type is not implemented yet!</colorpicker>', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (10, 'Date Picker', 'date_picker', 0, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (11, 'Date Picker With Time', 'date_picker_time', 0, '2015-03-26 23:47:44.854', '<div class="well">
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
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (13, 'Media Picker', 'media_picker', 0, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (19, 'Richtext Editor', 'richtext_editor', 0, '2015-03-26 23:47:44.854', '<textarea ck-editor id="{{prop.name}}" name="{{prop.name}}" ng-model="data.meta[prop.name]" rows="10" cols="80"></textarea>', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (12, 'Content Picker', 'content_picker', 0, '2015-03-26 23:47:44.854', '<div ng-controller="Collexy.DataTypePropertyEditor.ContentPicker">
<div ng-repeat="cn in contentNodes"><label><input type="checkbox" checklist-model="data.meta[prop.name]" checklist-value="cn.id"></label> {{cn.name}}</div>
<br>
<button type="button" ng-click="checkAll()">check all</button>
<button type="button" ng-click="uncheckAll()">uncheck all</button>
</div>', 'Collexy.DataTypeEditor.ContentPicker', '{"content_type_id": 7}');
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (4, 'Radio Button List', 'radio_button_list', 0, '2015-03-26 23:47:44.854', '<div ng-controller="Collexy.DataTypePropertyEditor.RadioButtonList.Controller">
    <ul>
        <li ng-repeat="option in dataType.meta.options">
            <label>
                <input type="radio" name="radio-button-list-{{dataType.alias}}" ng-value="option.value" ng-model="data.meta[prop.name]"/>
                {{option.label}}
            </label>
        </li>
    </ul>
</div>', 'Collexy.DataTypeEditor.RadioButtonList', '{"options": [{"label": "Value 1", "value": "val1"}, {"label": "Value 2", "value": "val2"}]}');
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (7, 'Checkbox List', 'checkbox_list', 0, '2015-03-26 23:47:44.854', '', NULL, NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (1, 'Text Input', 'text_input', 0, '2015-03-26 23:47:44.854', '<input type="text" id="{{prop.name}}" ng-model="data.meta[prop.name]"/>', '', 'null');
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (17, 'Domains', 'domains', 0, '2015-03-26 23:47:44.854', '<div ng-controller="Collexy.DataTypeEditor.Domains.Controller">
    <input type="text" ng-model="domainToAdd"/> <button type="button" ng-click="addDomain()">Add domain</button><br>
    <ul>
        <li ng-repeat="domain in data.meta[prop.name]">
            {{domain}} <button type="button" ng-click="removeDomain(domain)">x</button>
        </li>
    </ul>
    
</div>', '', NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (14, 'Folder Browser', 'folder_browser', 0, '2015-03-26 23:47:44.854', '<style>
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
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (16, 'Upload Multiple', 'upload_multiple', 0, '2015-03-26 23:47:44.854', '<div ng-controller="Collexy.DataTypeEditor.FileUpload.Controller">
    <pre>{{originalData.meta}}</pre>
    <div ng-show="data.meta.attached_file">
    <input type="text" ng-readonly="true" ng-model="data.meta.attached_file">
    </div>
    <div ng-if="!originalData.meta.attached_file">
        <img style="max-width: 100%;" src="{{location_url}}/{{data.meta.attached_file.name}}"/>
    </div>
    <div ng-if="originalData.meta.attached_file">
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
</div>', '', NULL);
INSERT INTO data_type (id, name, alias, created_by, created_date, html, editor_alias, meta) VALUES (15, 'Upload', 'upload', 0, '2015-03-26 23:47:44.854', '<!--<div ng-controller="Collexy.DataTypeEditor.FileUpload.Controller" collexy-file-upload>
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
</div>-->
<div ngf-drop ngf-select ng-model="$parent.$parent.$parent.files" class="drop-box" 
     ngf-drag-over-class="dragover" ngf-multiple="true" ngf-allow-dir="true" accept="image/*,application/pdf">Drop file(s) here or click to upload</div>

<div ngf-no-file-drop>File Drag/Drop is not supported for this browser</div>
Files:
<ul>
    <li ng-repeat="f in files" style="font:smaller">{{f.name}} <img ng-show="f != null" ngf-src="f" class="thumb"></li>
</ul>
Upload Log:
<pre>{{log}}</pre>
<style>
    .button {
        -moz-appearance: button;
        /* Firefox */
        -webkit-appearance: button;
        /* Safari and Chrome */
        padding: 10px;
        margin: 10px;
        width: 70px;
    }
    .drop-box {
        background: #F8F8F8;
        border: 5px dashed #DDD;
        /*width: 200px;
        height: 65px;*/
        text-align: center;
        padding: 25px;
        margin: 10px 0;
    }
    .dragover {
        border: 5px dashed blue;
    }
</style>', '', NULL);


--
-- TOC entry 2399 (class 0 OID 0)
-- Dependencies: 181
-- Name: data_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('data_type_id_seq', 19, true);


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
-- TOC entry 2361 (class 0 OID 98009)
-- Dependencies: 172
-- Data for Name: member; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member (id, username, password, email, meta, created_date, updated_date, login_date, accessed_date, status, sid, member_type_id, member_group_ids) VALUES (1, 'default_member', '$2a$10$f9qZyhrTnjirqK53kY3jRu93AgSXUryUZwwFhOFxhh1R9t7LgHRGa', 'default_member@mail.com', '{"comments": "default user comments"}', '2015-01-22 14:25:38.904', NULL, '2015-06-13 03:51:35.925', NULL, 1, 'J4CNIJQH2JY5XOYR5MQFKYBDMM7572YF7GDYJA5JZ6GHZ7XGGE6A', 1, '{1}');


--
-- TOC entry 2368 (class 0 OID 98131)
-- Dependencies: 179
-- Data for Name: member_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_group (id, name, alias, created_by, created_date) VALUES (1, 'Authenticated Member', 'authenticated_member', 1, '2015-03-26 17:09:34.18');
INSERT INTO member_group (id, name, alias, created_by, created_date) VALUES (2, 'Member Group 2', 'member_group_2', 1, '2015-06-05 09:48:29.034');


--
-- TOC entry 2402 (class 0 OID 0)
-- Dependencies: 178
-- Name: member_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_group_id_seq', 2, true);


--
-- TOC entry 2403 (class 0 OID 0)
-- Dependencies: 173
-- Name: member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_id_seq', 1, true);


--
-- TOC entry 2386 (class 0 OID 98969)
-- Dependencies: 197
-- Data for Name: member_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO member_type (id, path, parent_id, name, alias, created_by, created_date, description, icon, thumbnail, meta, tabs, is_abstract, composite_member_type_ids) VALUES (1, '1', NULL, 'Member', 'member', 1, '2015-03-26 19:56:03.85', 'This is the default member type for Collexy members.', 'fa fa-user fa-fw', 'fa fa-user fa-fw', NULL, '[{"name": "Membership", "properties": [{"name": "comments", "order": 1, "help_text": "Help text for membership comments", "description": "Membership comments description", "data_type_id": 3}]}]', false, NULL);


--
-- TOC entry 2404 (class 0 OID 0)
-- Dependencies: 196
-- Name: member_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('member_type_id_seq', 1, true);


--
-- TOC entry 2389 (class 0 OID 99012)
-- Dependencies: 200
-- Data for Name: mime_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO mime_type (id, name, media_type_id) VALUES (2, 'image/png', 2);
INSERT INTO mime_type (id, name, media_type_id) VALUES (3, 'image/svg+xml', 2);
INSERT INTO mime_type (id, name, media_type_id) VALUES (1, 'image/jpeg', 2);


--
-- TOC entry 2405 (class 0 OID 0)
-- Dependencies: 201
-- Name: mime_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('mime_type_id_seq', 3, true);


--
-- TOC entry 2388 (class 0 OID 99001)
-- Dependencies: 199
-- Data for Name: permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO permission (id, name) VALUES (1, 'all');
INSERT INTO permission (id, name) VALUES (2, 'content_all');
INSERT INTO permission (id, name) VALUES (3, 'content_create');
INSERT INTO permission (id, name) VALUES (4, 'content_delete');
INSERT INTO permission (id, name) VALUES (5, 'content_update');
INSERT INTO permission (id, name) VALUES (6, 'content_section');
INSERT INTO permission (id, name) VALUES (7, 'content_browse');
INSERT INTO permission (id, name) VALUES (8, 'media_all');
INSERT INTO permission (id, name) VALUES (9, 'media_create');
INSERT INTO permission (id, name) VALUES (10, 'media_delete');
INSERT INTO permission (id, name) VALUES (11, 'media_update');
INSERT INTO permission (id, name) VALUES (12, 'media_section');
INSERT INTO permission (id, name) VALUES (13, 'media_browse');
INSERT INTO permission (id, name) VALUES (14, 'user_all');
INSERT INTO permission (id, name) VALUES (15, 'user_create');
INSERT INTO permission (id, name) VALUES (16, 'user_delete');
INSERT INTO permission (id, name) VALUES (17, 'user_update');
INSERT INTO permission (id, name) VALUES (18, 'user_section');
INSERT INTO permission (id, name) VALUES (19, 'user_browse');
INSERT INTO permission (id, name) VALUES (20, 'user_type_all');
INSERT INTO permission (id, name) VALUES (21, 'user_type_create');
INSERT INTO permission (id, name) VALUES (22, 'user_type_delete');
INSERT INTO permission (id, name) VALUES (23, 'user_type_update');
INSERT INTO permission (id, name) VALUES (24, 'user_type_section');
INSERT INTO permission (id, name) VALUES (25, 'user_type_browse');
INSERT INTO permission (id, name) VALUES (26, 'user_group_all');
INSERT INTO permission (id, name) VALUES (27, 'user_group_create');
INSERT INTO permission (id, name) VALUES (28, 'user_group_delete');
INSERT INTO permission (id, name) VALUES (29, 'user_group_update');
INSERT INTO permission (id, name) VALUES (30, 'user_group_section');
INSERT INTO permission (id, name) VALUES (31, 'user_group_browse');
INSERT INTO permission (id, name) VALUES (32, 'permission_all');
INSERT INTO permission (id, name) VALUES (33, 'permission_create');
INSERT INTO permission (id, name) VALUES (34, 'permission_delete');
INSERT INTO permission (id, name) VALUES (35, 'permission_update');
INSERT INTO permission (id, name) VALUES (36, 'permission_section');
INSERT INTO permission (id, name) VALUES (37, 'permission_browse');
INSERT INTO permission (id, name) VALUES (38, 'asset_all');
INSERT INTO permission (id, name) VALUES (39, 'asset_create');
INSERT INTO permission (id, name) VALUES (40, 'asset_delete');
INSERT INTO permission (id, name) VALUES (41, 'asset_update');
INSERT INTO permission (id, name) VALUES (42, 'asset_browse');
INSERT INTO permission (id, name) VALUES (43, 'asset_section');
INSERT INTO permission (id, name) VALUES (44, 'member_all');
INSERT INTO permission (id, name) VALUES (45, 'member_create');
INSERT INTO permission (id, name) VALUES (46, 'member_delete');
INSERT INTO permission (id, name) VALUES (47, 'member_update');
INSERT INTO permission (id, name) VALUES (48, 'member_section');
INSERT INTO permission (id, name) VALUES (49, 'member_browse');
INSERT INTO permission (id, name) VALUES (50, 'member_type_all');
INSERT INTO permission (id, name) VALUES (51, 'member_type_create');
INSERT INTO permission (id, name) VALUES (52, 'member_type_delete');
INSERT INTO permission (id, name) VALUES (53, 'member_type_update');
INSERT INTO permission (id, name) VALUES (54, 'member_type_section');
INSERT INTO permission (id, name) VALUES (55, 'member_type_browse');
INSERT INTO permission (id, name) VALUES (56, 'member_group_all');
INSERT INTO permission (id, name) VALUES (57, 'member_group_create');
INSERT INTO permission (id, name) VALUES (58, 'member_group_delete');
INSERT INTO permission (id, name) VALUES (59, 'member_group_update');
INSERT INTO permission (id, name) VALUES (60, 'member_group_section');
INSERT INTO permission (id, name) VALUES (61, 'member_group_browse');
INSERT INTO permission (id, name) VALUES (62, 'template_all');
INSERT INTO permission (id, name) VALUES (63, 'template_create');
INSERT INTO permission (id, name) VALUES (64, 'template_delete');
INSERT INTO permission (id, name) VALUES (65, 'template_update');
INSERT INTO permission (id, name) VALUES (66, 'template_section');
INSERT INTO permission (id, name) VALUES (67, 'template_browse');
INSERT INTO permission (id, name) VALUES (68, 'settings_section');
INSERT INTO permission (id, name) VALUES (69, 'settings_all');
INSERT INTO permission (id, name) VALUES (70, 'script_all');
INSERT INTO permission (id, name) VALUES (71, 'script_create');
INSERT INTO permission (id, name) VALUES (72, 'script_delete');
INSERT INTO permission (id, name) VALUES (73, 'script_update');
INSERT INTO permission (id, name) VALUES (74, 'script_section');
INSERT INTO permission (id, name) VALUES (75, 'script_browse');
INSERT INTO permission (id, name) VALUES (76, 'stylesheet_all');
INSERT INTO permission (id, name) VALUES (77, 'stylesheet_create');
INSERT INTO permission (id, name) VALUES (78, 'stylesheet_delete');
INSERT INTO permission (id, name) VALUES (79, 'stylesheet_update');
INSERT INTO permission (id, name) VALUES (80, 'stylesheet_section');
INSERT INTO permission (id, name) VALUES (81, 'stylesheet_browse');
INSERT INTO permission (id, name) VALUES (82, 'content_type_all');
INSERT INTO permission (id, name) VALUES (83, 'content_type_create');
INSERT INTO permission (id, name) VALUES (84, 'content_type_delete');
INSERT INTO permission (id, name) VALUES (85, 'content_type_update');
INSERT INTO permission (id, name) VALUES (86, 'content_type_section');
INSERT INTO permission (id, name) VALUES (87, 'content_type_browse');
INSERT INTO permission (id, name) VALUES (88, 'media_type_all');
INSERT INTO permission (id, name) VALUES (89, 'media_type_create');
INSERT INTO permission (id, name) VALUES (90, 'media_type_delete');
INSERT INTO permission (id, name) VALUES (91, 'media_type_update');
INSERT INTO permission (id, name) VALUES (92, 'media_type_section');
INSERT INTO permission (id, name) VALUES (93, 'media_type_browse');
INSERT INTO permission (id, name) VALUES (94, 'data_type_all');
INSERT INTO permission (id, name) VALUES (95, 'data_type_create');
INSERT INTO permission (id, name) VALUES (96, 'data_type_delete');
INSERT INTO permission (id, name) VALUES (97, 'data_type_update');
INSERT INTO permission (id, name) VALUES (98, 'data_type_section');
INSERT INTO permission (id, name) VALUES (99, 'data_type_browse');
INSERT INTO permission (id, name) VALUES (100, 'node_create');
INSERT INTO permission (id, name) VALUES (101, 'node_delete');
INSERT INTO permission (id, name) VALUES (102, 'node_update');
INSERT INTO permission (id, name) VALUES (103, 'node_move');
INSERT INTO permission (id, name) VALUES (104, 'node_copy');
INSERT INTO permission (id, name) VALUES (105, 'node_public_access');
INSERT INTO permission (id, name) VALUES (106, 'node_permissions');
INSERT INTO permission (id, name) VALUES (107, 'node_send_to_publish');
INSERT INTO permission (id, name) VALUES (108, 'node_publish');
INSERT INTO permission (id, name) VALUES (109, 'node_browse');
INSERT INTO permission (id, name) VALUES (110, 'node_change_content_type');
INSERT INTO permission (id, name) VALUES (111, 'node_sort');
INSERT INTO permission (id, name) VALUES (112, 'mime_type_all');
INSERT INTO permission (id, name) VALUES (113, 'mime_type_create');
INSERT INTO permission (id, name) VALUES (114, 'mime_type_delete');
INSERT INTO permission (id, name) VALUES (115, 'mime_type_update');
INSERT INTO permission (id, name) VALUES (116, 'mime_type_section');
INSERT INTO permission (id, name) VALUES (117, 'mime_type_browse');


--
-- TOC entry 2406 (class 0 OID 0)
-- Dependencies: 198
-- Name: permission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('permission_id_seq', 117, true);


-- Completed on 2015-06-29 12:26:22

--
-- PostgreSQL database dump complete
--
`
