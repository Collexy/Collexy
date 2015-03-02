-- Table: node

-- Login to psql and run the following
-- What is the result?
-- SELECT MAX(id) FROM node;

-- Then run...
-- This should be higher than the last result.
-- SELECT nextval('node_id_seq');

-- If it's not higher... run this set the sequence last to your highest pid it. 
-- (wise to run a quick pg_dump first...)
-- SELECT setval('node_id_seq', (SELECT MAX(id) FROM your_table));


-- DROP TABLE node;
 	
TRUNCATE TABLE node RESTART IDENTITY;

CREATE TABLE node
(
  id bigint NOT NULL DEFAULT nextval(('public.node_id_seq'::text)::regclass),
  path ltree NOT NULL,
  created_by bigint NOT NULL,
  name character varying(255),
  node_type integer NOT NULL,
  created_date timestamp without time zone NOT NULL DEFAULT now(),
  CONSTRAINT node_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE node
  OWNER TO postgres;

INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1', 1, 'root', 5);

ALTER SEQUENCE node_id_seq RESTART WITH 1;
UPDATE node SET id = DEFAULT;

INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1.2', 1, 'Text input', 11);


INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1.3.4', 1, 'Page', 4);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.5', 1, 'Layout', 3);


INSERT INTO node(
            path, created_by, name, node_type)
    VALUES ('1.5.6', 1, 'Page', 3);


INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.7', 1, 'Home', 1);

INSERT INTO node(
        path, created_by, name, node_type)
VALUES ('1.7.8', 1, 'Sample Page', 1);

