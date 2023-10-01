# Golang-Api

This project consists of CRUD operations using Golang for backend, and SQL for database.
It performs various operations on table called books, which stores the info.

// Sql table create code

CREATE TABLE `books` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(100) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	`author` VARCHAR(30) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	`quantity` INT(11) NULL DEFAULT NULL,
	`unique_id` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `UNIQUE KEY` (`id`) USING BTREE,
	UNIQUE INDEX `title` (`title`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=57
;
