  CREATE TABLE IF NOT EXISTS `users_role` (
    `id` int(11) NOT NULL,
    `name` character varying(50) NOT NULL,
    `code` int(2) NOT NULL,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`id`)
  );

  INSERT INTO users_role (id, name, code) VALUES (1, "visitor", 0);
  INSERT INTO users_role (id. name, code) VALUES (2, "manager", 10);
  INSERT INTO users_role (id, name, code) VALUES (3, "technician", 20);

  CREATE TABLE IF NOT EXISTS `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` character varying(50) NOT NULL,
    `username` character varying(30) NOT NULL UNIQUE,
    `password` character varying(255) NOT NULL,
    `user_role_id` int(11) NOT NULL,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_role_id`) REFERENCES `users_role` (`id`)
  );

  CREATE TABLE IF NOT EXISTS `tasks` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` character varying(100) NOT NULL,
    `description` character varying(2500) NOT NULL,
    `user_id` int(11) NOT NULL,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `finished_at` TIMESTAMP,
    `deleted_at` TIMESTAMP,
    `created_at` TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
  );

  CREATE TABLE IF NOT EXISTS `migrations` (
    `name` character varying(100) NOT NULL PRIMARY KEY,
    `applied_at` TIMESTAMP NOT NULL
  );

  INSERT INTO migrations VALUES ('0001.up.sql', NOW());