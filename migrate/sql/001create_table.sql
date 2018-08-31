CREATE DATABASE parserms;
USE parserms;

CREATE TABLE `orderOLX` (
  `id` INT NOT NULL AUTO_INCREMENT,
	`user_id` INT NOT NULL,
	`url` VARCHAR(1024) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
	`page_limit` INT DEFAULT 0,
	`mail` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
	`expiration_time` INT NOT NULL,
	`frequency` INT NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE advertisements (
	`id` INT NOT NULL AUTO_INCREMENT,
	`order_id` INT NOT NULL,
	`title` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
	`url` VARCHAR(1024) NOT NULL DEFAULT '' COLLATE 'utf8_unicode_ci',
	`created_at`  BIGINT NOT NULL,
	PRIMARY KEY (`id`),
	FOREIGN KEY (`order_id`) REFERENCES `orderOLX`(`id`)
)
COLLATE='utf8_unicode_ci'
ENGINE=InnoDB