CREATE TABLE characters (
  id bigint NOT NULL,
  user_id bigint NOT NULL,
  image CHARACTER VARYING(256),
  name CHARACTER VARYING(128),
  PRIMARY KEY (id, user_id)
);
CREATE TABLE users (
  id bigint NOT NULL,
  quote CHARACTER VARYING(1024) NOT NULL DEFAULT '',
  date timestamp NOT NULL DEFAULT '1970-01-01 00:00:00-00',
  favorite bigint,
  claim_count int NOT NULL DEFAULT 0,
  PRIMARY KEY (id)
);
ALTER TABLE characters
ADD CONSTRAINT FK_users_TO_characters FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE users
ADD CONSTRAINT FK_characters_TO_users FOREIGN KEY (favorite) REFERENCES characters (id);