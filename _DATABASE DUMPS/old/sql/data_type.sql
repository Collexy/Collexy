-- Table: data_type

-- DROP TABLE data_type;

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
ALTER TABLE data_type
  OWNER TO postgres;

INSERT INTO data_type(
            node_id, html)
    VALUES (2, '<input type="text" id="{{prop.name}}" ng-model="node.content.meta[prop.name]">');


INSERT INTO data_type(
            node_id, html)
    VALUES (14, '<textarea id="{{prop.name}}" ng-model="node.content.meta[prop.name]">');

SELECT * FROM node
JOIN data_type ON node.id = data_type.node_id
WHERE node.id = 2;