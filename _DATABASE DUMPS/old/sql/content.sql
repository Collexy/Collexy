-- Table: content

DROP TABLE content;

CREATE TABLE content
(
  id bigserial NOT NULL,
  node_id bigint NOT NULL,
  content_type_node_id bigint NOT NULL,
  meta jsonb,
  properties jsonb,
  CONSTRAINT document_pkey PRIMARY KEY (id, node_id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE content
  OWNER TO postgres;
