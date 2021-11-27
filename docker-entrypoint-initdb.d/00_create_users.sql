DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `email` VARCHAR(255),
  PRIMARY KEY (`id`)
);

INSERT INTO users (id, name, email) VALUES (1, "Yamada", "yamada@example.com");
INSERT INTO users (id, name, email) VALUES (2, "Tanaka", "tanaka@example.com");