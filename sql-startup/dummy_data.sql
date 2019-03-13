-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema internal
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `internal` ;

-- -----------------------------------------------------
-- Schema internal
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `internal` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema symp_internal
-- -----------------------------------------------------
-- This schema was created for a stub table
USE `internal` ;

-- -----------------------------------------------------
-- Table `internal`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`users` ;

CREATE TABLE IF NOT EXISTS `internal`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NULL,
  `real_name` VARCHAR(255) NULL,
  `email` VARCHAR(64) NULL,
  `created_at` DATETIME NULL DEFAULT NOW(),
  `pass_hash` VARCHAR(255) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`lecture_suggestion`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`lecture_suggestion` ;

CREATE TABLE IF NOT EXISTS `internal`.`lecture_suggestion` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `lecture_name` VARCHAR(64) NULL,
  `speaker_name` VARCHAR(64) NULL,
  `from_nonprague` TINYINT(1) NULL,
  `speaker_bio` VARCHAR(255) NULL,
  `lecture_desc` VARCHAR(255) NULL,
  `preferences` VARCHAR(255) NULL,
  `author` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_lecture_suggestion_users1_idx` (`author` ASC),
  CONSTRAINT `fk_lecture_suggestion_users1`
    FOREIGN KEY (`author`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`topic_suggestion`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`topic_suggestion` ;

CREATE TABLE IF NOT EXISTS `internal`.`topic_suggestion` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `desc` VARCHAR(255) NULL,
  `topic_name` VARCHAR(45) NULL,
  `illustration_url` VARCHAR(128) NULL,
  `author` INT NOT NULL,
  PRIMARY KEY (`id`, `author`),
  INDEX `fk_topic_suggestion_users1_idx` (`author` ASC),
  CONSTRAINT `fk_topic_suggestion_users1`
    FOREIGN KEY (`author`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`image_suggestion`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`image_suggestion` ;

CREATE TABLE IF NOT EXISTS `internal`.`image_suggestion` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `filename_uri` VARCHAR(128) NULL,
  `desc` VARCHAR(255) NULL,
  `author` INT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_image_suggestion_users1_idx` (`author` ASC),
  CONSTRAINT `fk_image_suggestion_users1`
    FOREIGN KEY (`author`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`lecture_suggestion_votes`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`lecture_suggestion_votes` ;

CREATE TABLE IF NOT EXISTS `internal`.`lecture_suggestion_votes` (
  `voter_id` INT NOT NULL,
  `for_or_against` TINYINT(1) NULL,
  `comment` VARCHAR(255) NULL,
  `lecture_id` INT NULL,
  `lecture_suggestion_id` INT NOT NULL,
  PRIMARY KEY (`voter_id`, `lecture_suggestion_id`),
  INDEX `fk_lecture_suggestion_votes_lecture_suggestion1_idx` (`lecture_suggestion_id` ASC),
  CONSTRAINT `fk_lecture_suggestion`
    FOREIGN KEY (`lecture_suggestion_id`)
    REFERENCES `internal`.`lecture_suggestion` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_voter_id_lecture`
    FOREIGN KEY (`voter_id`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'TODO: make sure combination of voter and voted subject is un' /* comment truncated */ /*ique (COMPOSITE KEY)*/;


-- -----------------------------------------------------
-- Table `internal`.`topic_suggestion_votes`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`topic_suggestion_votes` ;

CREATE TABLE IF NOT EXISTS `internal`.`topic_suggestion_votes` (
  `voter_id` INT NOT NULL,
  `for_or_against` TINYINT(1) NULL,
  `comment` VARCHAR(255) NULL,
  `topic_id` INT NULL,
  PRIMARY KEY (`voter_id`),
  INDEX `topic_id_idx` (`topic_id` ASC),
  CONSTRAINT `fk_topic_id`
    FOREIGN KEY (`topic_id`)
    REFERENCES `internal`.`topic_suggestion` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_voter_id_topic`
    FOREIGN KEY (`voter_id`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'TODO: make sure combination of voter and voted subject is un' /* comment truncated */ /*ique (COMPOSITE KEY)*/;


-- -----------------------------------------------------
-- Table `internal`.`image_suggestion_votes`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`image_suggestion_votes` ;

CREATE TABLE IF NOT EXISTS `internal`.`image_suggestion_votes` (
  `voter_id` INT NOT NULL,
  `for_or_against` TINYINT(1) NULL,
  `comment` VARCHAR(255) NULL,
  `image_id` INT NULL,
  PRIMARY KEY (`voter_id`),
  INDEX `image_suggestion_vote_idx` (`image_id` ASC),
  CONSTRAINT `image_suggestion_vote`
    FOREIGN KEY (`image_id`)
    REFERENCES `internal`.`image_suggestion` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_voter_id_image`
    FOREIGN KEY (`voter_id`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'TODO: make sure combination of voter and voted subject is un' /* comment truncated */ /*ique (COMPOSITE KEY)*/;


-- -----------------------------------------------------
-- Table `internal`.`symp_meta`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`symp_meta` ;

CREATE TABLE IF NOT EXISTS `internal`.`symp_meta` (
  `key` VARCHAR(255) NOT NULL,
  `value` VARCHAR(255) NULL,
  `id` INT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'Contains server-wide configuration -- i.e. phase whole event' /* comment truncated */ /* is in, website configuration variables*/;


-- -----------------------------------------------------
-- Table `internal`.`website_pages`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`website_pages` ;

CREATE TABLE IF NOT EXISTS `internal`.`website_pages` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `title` VARCHAR(45) NULL,
  `nav_name` VARCHAR(45) NULL,
  `removable` TINYINT(1) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`website_fields`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`website_fields` ;

CREATE TABLE IF NOT EXISTS `internal`.`website_fields` (
  `id` INT NOT NULL,
  `key` VARCHAR(45) NULL,
  `value` VARCHAR(45) NULL,
  `page_id` INT NOT NULL,
  PRIMARY KEY (`id`, `page_id`),
  INDEX `fk_website_fields_website_pages1_idx` (`page_id` ASC),
  CONSTRAINT `fk_website_fields_website_pages1`
    FOREIGN KEY (`page_id`)
    REFERENCES `internal`.`website_pages` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`rooms`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`rooms` ;

CREATE TABLE IF NOT EXISTS `internal`.`rooms` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(32) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`times`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`times` ;

CREATE TABLE IF NOT EXISTS `internal`.`times` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `begins_at` INT NULL,
  `ends_at` INT NULL,
  `label` VARCHAR(64) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'CURRENTLY OF CONSTANT VALUE -- stores times of day for lectu' /* comment truncated */ /*res to be held at*/;


-- -----------------------------------------------------
-- Table `internal`.`user_permissions`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`user_permissions` ;

CREATE TABLE IF NOT EXISTS `internal`.`user_permissions` (
  `id` INT NOT NULL,
  `name` VARCHAR(64) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`collections`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`collections` ;

CREATE TABLE IF NOT EXISTS `internal`.`collections` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `collection_name` VARCHAR(32) NULL,
  `hierarchy_pos` INT NULL,
  INDEX `fk_collections_hierarchy_pos` (`hierarchy_pos` ASC),
  PRIMARY KEY (`id`))
  -- CONSTRAINT `immediate_superior`
  --   FOREIGN KEY (`immediate_superior`)
  --   REFERENCES `internal`.`collections` (`id`)
  --   ON DELETE NO ACTION
  --   ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`permissions_by_collections`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`permissions_by_collections` ;

CREATE TABLE IF NOT EXISTS `internal`.`permissions_by_collections` (
  `collections_id` INT NOT NULL,
  `perm_id` INT NOT NULL,
  PRIMARY KEY (`collections_id`, `perm_id`),
  INDEX `fk_permissions_by_role_user_permissions1_idx` (`perm_id` ASC),
  CONSTRAINT `fk_permissions_by_role_user_permissions1`
    FOREIGN KEY (`perm_id`)
    REFERENCES `internal`.`user_permissions` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_permissions_by_collections_1`
    FOREIGN KEY (`collections_id`)
    REFERENCES `internal`.`collections` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;




-- -----------------------------------------------------
-- Table `internal`.`rooms`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`rooms` ;

CREATE TABLE IF NOT EXISTS `internal`.`rooms` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`lectures`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`lectures` ;

CREATE TABLE IF NOT EXISTS `internal`.`lectures` (
  `id` INT NOT NULL,
  `lecture_name` VARCHAR(64) NULL,
  `speaker_name` VARCHAR(64) NULL,
  `speaker_bio` VARCHAR(255) NULL,
  `lecture_desc` VARCHAR(255) NULL,
  `from_nonprague` TINYINT(1) NULL,
  `preferences` VARCHAR(255) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `internal`.`days`
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS `internal`.`days` (
  `id` INT NOT NULL,
  `name` VARCHAR(255) NULL,
  `date` DATE UNIQUE NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `internal`.`days_x_timeslots_x_rooms_to_lectures`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`days_x_timeslots_x_rooms_to_lectures` ;

CREATE TABLE IF NOT EXISTS `internal`.`days_x_timeslots_x_rooms_to_lectures` (
  `day_id` INT NOT NULL,
  `timeslots_id` INT NOT NULL,
  `rooms_id` INT NOT NULL,
  `lecture_id` INT NOT NULL,
  PRIMARY KEY (`timeslots_id`, `rooms_id`, `day_id`),
  INDEX `fk_timeslots_x_rooms_to_lectures_timeslots1_idx` (`timeslots_id` ASC),
  INDEX `fk_timeslots_x_rooms_to_lectures_rooms1_idx` (`rooms_id` ASC),
  INDEX `fk_timeslots_x_rooms_to_lectures_lecture_suggestion_copy21_idx` (`lecture_id` ASC),
  CONSTRAINT `fk_timeslots_x_rooms_to_lectures_timeslots1`
    FOREIGN KEY (`timeslots_id`)
    REFERENCES `internal`.`times` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_timeslots_x_rooms_to_days`
    FOREIGN KEY (`day_id`)
    REFERENCES `internal`.`days` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_timeslots_x_rooms_to_lectures_rooms1`
    FOREIGN KEY (`rooms_id`)
    REFERENCES `internal`.`rooms` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_timeslots_x_rooms_to_lectures_lecture_suggestion_copy21`
    FOREIGN KEY (`lecture_id`)
    REFERENCES `internal`.`lectures` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'TODO: make sure there is a composite key on this one ( uniqu' /* comment truncated */ /*e combination of timeslot and room, but each can appear many times*/;


-- -----------------------------------------------------
-- Table `internal`.`messages`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`messages` ;

CREATE TABLE IF NOT EXISTS `internal`.`messages` (
  `message_id` INT NOT NULL AUTO_INCREMENT,
  `from` INT NULL,
  `to` INT NULL,
  `contents` VARCHAR(255) NULL,
  `title` VARCHAR(225) NULL,
  `urgency` INT NULL,
  `associated_deadline` DATETIME NULL,
  `to_collection` INT NOT NULL,
  PRIMARY KEY (`message_id`, `to_collection`),
  INDEX `fk_messages_users1_idx` (`to` ASC),
  INDEX `fk_messages_users_to_idx` (`from` ASC),
  INDEX `fk_messages_collections1_idx` (`to_collection` ASC),
  CONSTRAINT `fk_messages_users_from`
    FOREIGN KEY (`to`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_messages_users_to`
    FOREIGN KEY (`from`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_messages_collections1`
    FOREIGN KEY (`to_collection`)
    REFERENCES `internal`.`collections` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`collections_by_users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`collections_by_users` ;

CREATE TABLE IF NOT EXISTS `internal`.`collections_by_users` (
  `collections_id` INT NOT NULL,
  `users_id` INT NOT NULL,
  PRIMARY KEY (`collections_id`, `users_id`),
  INDEX `fk_collections_has_users_users1_idx` (`users_id` ASC),
  INDEX `fk_collections_has_users_collections1_idx` (`collections_id` ASC),
  CONSTRAINT `fk_collections_has_users_collections1`
    FOREIGN KEY (`collections_id`)
    REFERENCES `internal`.`collections` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_collections_has_users_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `internal`.`api_tokens`
-- store used api tokens
-- -----------------------------------------------------
DROP TABLE IF EXISTS `internal`.`api_tokens` ;

CREATE TABLE IF NOT EXISTS `internal`.`api_tokens` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `token` VARCHAR(255) UNIQUE NOT NULL,
  `valid` TINYINT(1) NULL,
  `expires` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

-- -----------------------------------------------------
-- Data for table `internal`.`users`
-- -----------------------------------------------------
START TRANSACTION;
USE `internal`;
INSERT INTO `internal`.`users` (`id`, `name`, `real_name`, `email`, `created_at`, `pass_hash`) VALUES (1, 'root', 'Růt Růtovičová', 'root@symp.cz', '2000-01-01 00:00:00', '$2y$10$lSoRURTzK1110tlK8rYvbeDv.i7pJbe0tk6bQ6U.9j2XzJD9/lXCi\n');
INSERT INTO `internal`.`users` (`id`, `name`, `real_name`, `email`, `created_at`, `pass_hash`) VALUES (2, 'user', 'Usir Userov', 'user@symp.cz', '2000-01-01 00:00:00', '$2y$10$4HyPyMawnaQ1K0fzGSWu6u2B2nH1scjgMcRcBRlHLNS8teSQVsw2i\n');
INSERT INTO `internal`.`users` (`id`, `name`, `real_name`, `email`, `created_at`, `pass_hash`) VALUES (3, 'editor', 'Edith Orová', 'edit@symp.cz', '2000-01-01 00:00:00', '$2y$10$lSoRURTzK1110tlK8rYvbeDv.i7pJbe0tk6bQ6U.9j2XzJD9/lXCi\n');
INSERT INTO `internal`.`users` (`id`, `name`, `real_name`, `email`, `created_at`, `pass_hash`) VALUES (4, 'coordinator', 'Dušan Koordinátor', 'coordinator@symp.cz', '2000-01-01 00:00:00', '$2y$10$4HyPyMawnaQ1K0fzGSWu6u2B2nH1scjgMcRcBRlHLNS8teSQVsw2i\n');
COMMIT;


-- -----------------------------------------------------
-- Data for table `internal`.`user_permissions`
-- -----------------------------------------------------
START TRANSACTION;
USE `internal`;
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (1, 'MAKE_USER');
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (2, 'READ_USERLIST');
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (3, 'READ_TIMETABLE');
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (4, 'ALTER_TIMETABLE');
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (5, 'SUGGEST_LECTURE');
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (6, 'ALTER_LECTURE_CONTENTS');
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (7, 'INSERT_LECTURE');
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (8, 'DELETE_LECTURE');
INSERT INTO `internal`.`user_permissions` (`id`, `name`) VALUES (9, 'ALTER_USERS');

COMMIT;


-- -----------------------------------------------------
-- Data for table `internal`.`collections`
-- -----------------------------------------------------
START TRANSACTION;
USE `internal`;
INSERT INTO `internal`.`collections` (`id`, `collection_name`, `hierarchy_pos`) VALUES (1, 'ADMINS', 0);
INSERT INTO `internal`.`collections` (`id`, `collection_name`, `hierarchy_pos`) VALUES (2, 'USERS', 3);
INSERT INTO `internal`.`collections` (`id`, `collection_name`, `hierarchy_pos`) VALUES (3, 'EDITORS', 2);
INSERT INTO `internal`.`collections` (`id`, `collection_name`, `hierarchy_pos`) VALUES (4, 'COORDINATORS', 1);
INSERT INTO `internal`.`collections` (`id`, `collection_name`, `hierarchy_pos`) VALUES (5, 'NOBODY', 4);

COMMIT;



-- -----------------------------------------------------
-- Data for table `internal`.`collections_by_users`
-- -----------------------------------------------------
START TRANSACTION;
USE `internal`;
INSERT INTO `internal`.`collections_by_users` (`collections_id`, `users_id`) VALUES (1, 1);
INSERT INTO `internal`.`collections_by_users` (`collections_id`, `users_id`) VALUES (2, 2);
INSERT INTO `internal`.`collections_by_users` (`collections_id`, `users_id`) VALUES (3, 3);
INSERT INTO `internal`.`collections_by_users` (`collections_id`, `users_id`) VALUES (4, 4);
INSERT INTO `internal`.`collections_by_users` (`collections_id`, `users_id`) VALUES (2, 4);


COMMIT;

-- -----------------------------------------------------
-- Data for table `internal`.`permissions_by_collections`
-- -----------------------------------------------------
START TRANSACTION;
USE `internal`;
-- ADMIN -- everything
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 1);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 2);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 3);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 4);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 5);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 6);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 7);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 8);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (1, 9);
-- USER -- basic things
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (2, 1);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (2, 2);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (2, 3);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (2, 5);
-- EDITOR -- alter lecture contents
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (3, 1);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (3, 2);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (3, 3);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (3, 5);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (3, 6);
-- COORDINATOR -- can do everything but touch user accounts, perhaps private information
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (4, 1);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (4, 2);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (4, 3);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (4, 4);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (4, 5);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (4, 6);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (4, 7);
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (4, 8);
-- NOBODY -- just to allow non-users to register TODO: consider null instead of adding a new label
INSERT INTO `internal`.`permissions_by_collections` (`collections_id`, `perm_id`) VALUES (5, 1);
COMMIT;

START TRANSACTION;
USE `internal`;
-- DAYS
INSERT INTO `internal`.`days` (`id`, `name`, `date`) VALUES (1, "Sobota",'2019-03-17');
-- INSERT INTO `internal`.`days` (`id`, `name`, `date`) VALUES (2, "Neděle",'2019-03-18');
COMMIT;

-- LECTURES
START TRANSACTION;
USE `internal`;
INSERT INTO `internal`.`lectures` (`id`, `lecture_name`, `speaker_name`, `speaker_bio`, `lecture_desc`, `from_nonprague`, `preferences` ) VALUES (1, 'Převážně o jsoucnech', 'Ing. Mgr. Evžen Korelát, CSc.','Věhlasný akademik a amatérský pletač ponožek.','Jak už jsem řekl, bude to převážně o jsoucnech',1,'Konzumuje pouze plyny.');
INSERT INTO `internal`.`lectures` (`id`, `lecture_name`, `speaker_name`, `speaker_bio`, `lecture_desc`, `from_nonprague`, `preferences` ) VALUES (2, 'Pojednání o synestetických korelátech', 'Ing. Mgr. Arnošt Narativ, PHd.','Věhlasný pletač ponožek a amatérský akademik.','Název přednášky mluví sám za sebe.',0,'Má alergii na kapaliny.');
COMMIT;

-- ROOMS
START TRANSACTION;
USE `internal`;
INSERT INTO `internal`.`rooms` (`id`, `name`) VALUES (1, "P12");
INSERT INTO `internal`.`rooms` (`id`, `name`) VALUES (2, "Sborovna");
INSERT INTO `internal`.`rooms` (`id`, `name`) VALUES (3, "USV");
COMMIT;

-- TIMESLOTS
START TRANSACTION;
USE `internal`;
INSERT INTO `internal`.`times` (`id`, `begins_at`, `ends_at`, `label`)  VALUES (1, 32,36,'8 do 9'); -- od 8 do 9
INSERT INTO `internal`.`times` (`id`, `begins_at`, `ends_at`, `label`)  VALUES (2, 36,40,'9 do 10'); -- od 9 do 10
INSERT INTO `internal`.`times` (`id`, `begins_at`, `ends_at`, `label`)  VALUES (3, 40,44,'10 do 11'); -- od 10 do 11
INSERT INTO `internal`.`times` (`id`, `begins_at`, `ends_at`, `label`)  VALUES (4, 44,48,'11 do 12'); -- od 11 do 12
COMMIT;

START TRANSACTION;
USE `internal`;

INSERT INTO `internal`.`days_x_timeslots_x_rooms_to_lectures` (`day_id`, `timeslots_id`, `rooms_id`, `lecture_id`) VALUES (1,1,1,1);
INSERT INTO `internal`.`days_x_timeslots_x_rooms_to_lectures` (`day_id`, `timeslots_id`, `rooms_id`, `lecture_id`) VALUES (1,2,1,2);
INSERT INTO `internal`.`days_x_timeslots_x_rooms_to_lectures` (`day_id`, `timeslots_id`, `rooms_id`, `lecture_id`) VALUES (1,3,2,2);
INSERT INTO `internal`.`days_x_timeslots_x_rooms_to_lectures` (`day_id`, `timeslots_id`, `rooms_id`, `lecture_id`) VALUES (1,4,2,1);

COMMIT;
