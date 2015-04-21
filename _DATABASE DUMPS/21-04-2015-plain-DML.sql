--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4beta3
-- Dumped by pg_dump version 9.4beta3
-- Started on 2015-04-21 10:40:40

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- TOC entry 191 (class 3079 OID 11855)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2311 (class 0 OID 0)
-- Dependencies: 191
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


--
-- TOC entry 193 (class 3079 OID 97715)
-- Name: citext; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- TOC entry 2312 (class 0 OID 0)
-- Dependencies: 193
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- TOC entry 192 (class 3079 OID 97799)
-- Name: ltree; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;


--
-- TOC entry 2313 (class 0 OID 0)
-- Dependencies: 192
-- Name: EXTENSION ltree; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';


SET search_path = public, pg_catalog;

--
-- TOC entry 313 (class 1255 OID 97974)
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
-- TOC entry 314 (class 1255 OID 97975)
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
-- TOC entry 315 (class 1255 OID 97976)
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
-- TOC entry 189 (class 1259 OID 98246)
-- Name: content; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE content (
    id bigint NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    alias character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now(),
    content_type_id bigint,
    meta jsonb,
    public_access jsonb,
    user_permissions jsonb,
    user_group_permissions jsonb,
    type_id smallint,
    is_abstract boolean DEFAULT false NOT NULL
);


ALTER TABLE content OWNER TO postgres;

--
-- TOC entry 190 (class 1259 OID 98253)
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
-- TOC entry 2314 (class 0 OID 0)
-- Dependencies: 190
-- Name: content_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_id_seq OWNED BY content.id;


--
-- TOC entry 187 (class 1259 OID 98227)
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
    type_id smallint,
    allow_at_root boolean DEFAULT false NOT NULL,
    is_container boolean DEFAULT false NOT NULL,
    is_abstract boolean DEFAULT false NOT NULL,
    allowed_content_type_ids bigint[],
    composite_content_type_ids bigint[]
);


ALTER TABLE content_type OWNER TO postgres;

--
-- TOC entry 188 (class 1259 OID 98230)
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
-- TOC entry 2315 (class 0 OID 0)
-- Dependencies: 188
-- Name: content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE content_type_id_seq OWNED BY content_type.id;


--
-- TOC entry 183 (class 1259 OID 98158)
-- Name: data_type; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE data_type (
    id bigint NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    alias character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now(),
    html text
);


ALTER TABLE data_type OWNER TO postgres;

--
-- TOC entry 184 (class 1259 OID 98162)
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
-- TOC entry 2316 (class 0 OID 0)
-- Dependencies: 184
-- Name: data_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE data_type_id_seq OWNED BY data_type.id;


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
-- TOC entry 180 (class 1259 OID 98131)
-- Name: member_group; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE member_group (
    id bigint NOT NULL,
    path ltree,
    parent_id bigint,
    name character varying,
    alias character varying,
    created_by bigint,
    created_date timestamp without time zone DEFAULT now()
);


ALTER TABLE member_group OWNER TO postgres;

--
-- TOC entry 179 (class 1259 OID 98129)
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
-- TOC entry 2317 (class 0 OID 0)
-- Dependencies: 179
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
-- TOC entry 2318 (class 0 OID 0)
-- Dependencies: 173
-- Name: member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_id_seq OWNED BY member.id;


--
-- TOC entry 181 (class 1259 OID 98139)
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
    meta jsonb,
    tabs jsonb
);


ALTER TABLE member_type OWNER TO postgres;

--
-- TOC entry 182 (class 1259 OID 98142)
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
-- TOC entry 2319 (class 0 OID 0)
-- Dependencies: 182
-- Name: member_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE member_type_id_seq OWNED BY member_type.id;


--
-- TOC entry 174 (class 1259 OID 98044)
-- Name: permission; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE permission (
    name character varying NOT NULL
);


ALTER TABLE permission OWNER TO postgres;

--
-- TOC entry 186 (class 1259 OID 98211)
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
-- TOC entry 185 (class 1259 OID 98209)
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
-- TOC entry 2320 (class 0 OID 0)
-- Dependencies: 185
-- Name: template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE template_id_seq OWNED BY template.id;


--
-- TOC entry 175 (class 1259 OID 98067)
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
-- TOC entry 176 (class 1259 OID 98074)
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
-- TOC entry 177 (class 1259 OID 98080)
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
-- TOC entry 2321 (class 0 OID 0)
-- Dependencies: 177
-- Name: user_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_group_id_seq OWNED BY user_group.id;


--
-- TOC entry 178 (class 1259 OID 98082)
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
-- TOC entry 2322 (class 0 OID 0)
-- Dependencies: 178
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE user_id_seq OWNED BY "user".id;


--
-- TOC entry 2186 (class 2604 OID 98255)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content ALTER COLUMN id SET DEFAULT nextval('content_id_seq'::regclass);


--
-- TOC entry 2181 (class 2604 OID 98232)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY content_type ALTER COLUMN id SET DEFAULT nextval('content_type_id_seq'::regclass);


--
-- TOC entry 2176 (class 2604 OID 98164)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY data_type ALTER COLUMN id SET DEFAULT nextval('data_type_id_seq'::regclass);


--
-- TOC entry 2168 (class 2604 OID 98088)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member ALTER COLUMN id SET DEFAULT nextval('member_id_seq'::regclass);


--
-- TOC entry 2172 (class 2604 OID 98134)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_group ALTER COLUMN id SET DEFAULT nextval('member_group_id_seq'::regclass);


--
-- TOC entry 2174 (class 2604 OID 98144)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY member_type ALTER COLUMN id SET DEFAULT nextval('member_type_id_seq'::regclass);


--
-- TOC entry 2178 (class 2604 OID 98214)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY template ALTER COLUMN id SET DEFAULT nextval('template_id_seq'::regclass);


--
-- TOC entry 2170 (class 2604 OID 98094)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);


--
-- TOC entry 2171 (class 2604 OID 98095)
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY user_group ALTER COLUMN id SET DEFAULT nextval('user_group_id_seq'::regclass);


--
-- TOC entry 2190 (class 2606 OID 98105)
-- Name: permission_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY permission
    ADD CONSTRAINT permission_name_key UNIQUE (name);


--
-- TOC entry 2192 (class 2606 OID 98109)
-- Name: user_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- TOC entry 2194 (class 2606 OID 98111)
-- Name: user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres; Tablespace: 
--

ALTER TABLE ONLY "user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 2310 (class 0 OID 0)
-- Dependencies: 6
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2015-04-21 10:40:41

--
-- PostgreSQL database dump complete
--

