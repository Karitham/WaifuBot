CREATE TABLE characters (
  id BIGINT NOT NULL,
  user_id BIGINT NOT NULL,
  image CHARACTER VARYING(256),
  name CHARACTER VARYING(128),
  PRIMARY KEY (id, user_id)
);
CREATE TABLE users (
  id BIGINT NOT NULL,
  quote CHARACTER VARYING(1024) DEFAULT '' NOT NULL,
  date timestamp NOT NULL DEFAULT '1970-01-01 00:00:00-00',
  favorite BIGINT,
  PRIMARY KEY (id)
);
ALTER TABLE characters
ADD CONSTRAINT FK_users_TO_characters FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE users
ADD CONSTRAINT FK_characters_TO_users FOREIGN KEY (favorite) REFERENCES characters (id) ON UPDATE CASCADE ON DELETE CASCADE;