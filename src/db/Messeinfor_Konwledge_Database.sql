CREATE DATABASE messeinfor_kownledge;
CREATE SCHEMA dev;
CREATE TABLE dev.title
(
  id VARCHAR(60) NOT NULL,
  title VARCHAR(200),
  expand BOOLEAN DEFAULT FALSE  NOT NULL,
  sort SERIAL NOT NULL,
  level INT NOT NULL,
  parent_id VARCHAR(60)
);
CREATE UNIQUE INDEX title_id_uindex ON dev.title (id);
CREATE UNIQUE INDEX title_sort_uindex ON dev.title (sort);