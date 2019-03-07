-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema internal
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema internal
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `internal` DEFAULT CHARACTER SET utf8 ;
-- -----------------------------------------------------
-- Schema symp_internal
-- -----------------------------------------------------
-- This schema was created for a stub table

-- -----------------------------------------------------
-- Schema symp_internal
--
-- This schema was created for a stub table
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `symp_internal` ;
USE `internal` ;

-- -----------------------------------------------------
-- Table `internal`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`users` (
  `id` INT NOT NULL,
  `name` VARCHAR(255) NULL,
  `real_name` VARCHAR(255) NULL,
  `email` VARCHAR(64) NULL,
  `created_at` DATETIME NULL,
  `pass_hash` VARCHAR(255) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`lecture_suggestion`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`lecture_suggestion` (
  `id` INT NOT NULL,
  `lecture_name` VARCHAR(45) NULL,
  `speaker_name` VARCHAR(45) NULL,
  `speaker_bio` VARCHAR(45) NULL,
  `lecture_desc` VARCHAR(45) NULL,
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
CREATE TABLE IF NOT EXISTS `internal`.`topic_suggestion` (
  `id` INT NOT NULL,
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
CREATE TABLE IF NOT EXISTS `internal`.`image_suggestion` (
  `id` INT NOT NULL,
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
COMMENT = 'TODO: make sure combination of voter and voted subject is unique (COMPOSITE KEY)';


-- -----------------------------------------------------
-- Table `internal`.`topic_suggestion_votes`
-- -----------------------------------------------------
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
COMMENT = 'TODO: make sure combination of voter and voted subject is unique (COMPOSITE KEY)';


-- -----------------------------------------------------
-- Table `internal`.`image_suggestion_votes`
-- -----------------------------------------------------
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
COMMENT = 'TODO: make sure combination of voter and voted subject is unique (COMPOSITE KEY)';


-- -----------------------------------------------------
-- Table `internal`.`symp_meta`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`symp_meta` (
  `key` VARCHAR(255) NOT NULL,
  `value` VARCHAR(255) NULL,
  `id` INT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'Contains server-wide configuration -- i.e. phase whole event is in, website configuration variables';


-- -----------------------------------------------------
-- Table `internal`.`website_pages`
-- -----------------------------------------------------
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
-- Table `internal`.`lectures`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`lectures` (
  `id` INT NOT NULL,
  `lecture_name` VARCHAR(255) NULL,
  `speaker_name` VARCHAR(32) NULL,
  `speaker_bio` VARCHAR(255) NULL,
  `lecture_desc` VARCHAR(255) NULL,
  `from_nonprague` TINYINT(1) NULL,
  `preferences` VARCHAR(255) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`user_privileges`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`user_privileges` (
  `id` INT NOT NULL,
  `name` VARCHAR(64) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`groups`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`groups` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `group_name` VARCHAR(32) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'TODO: decide on keeping this -- many-to-many grouping of people, for mass messaging, \ninstruction management and so on -- more sensible for the type of event we\'re dealing with\nthis should replace ROLES';


-- -----------------------------------------------------
-- Table `internal`.`privileges_by_group`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`privileges_by_group` (
  `group_id` INT NOT NULL,
  `priv_id` INT NOT NULL,
  PRIMARY KEY (`group_id`, `priv_id`),
  INDEX `fk_privileges_by_role_user_privileges1_idx` (`priv_id` ASC),
  CONSTRAINT `fk_privileges_by_role_user_privileges1`
    FOREIGN KEY (`priv_id`)
    REFERENCES `internal`.`user_privileges` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_privileges_by_group_1`
    FOREIGN KEY (`group_id`)
    REFERENCES `internal`.`groups` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`timeslots`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`timeslots` (
  `id` INT NOT NULL,
  `from` DATETIME NULL,
  `to` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'CURRENTLY OF CONSTANT VALUE -- stores times of day for lectures to be held at';


-- -----------------------------------------------------
-- Table `internal`.`rooms`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`rooms` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`lectures`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`lectures` (
  `id` INT NOT NULL,
  `lecture_name` VARCHAR(255) NULL,
  `speaker_name` VARCHAR(32) NULL,
  `speaker_bio` VARCHAR(255) NULL,
  `lecture_desc` VARCHAR(255) NULL,
  `from_nonprague` TINYINT(1) NULL,
  `preferences` VARCHAR(255) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`timeslots_x_rooms_to_lectures`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`timeslots_x_rooms_to_lectures` (
  `rooms_id` INT NOT NULL,
  `lecture_id` INT NOT NULL,
  `timeslots_id` INT NOT NULL,
  PRIMARY KEY (`rooms_id`, `lecture_id`, `timeslots_id`),
  INDEX `fk_timeslots_x_rooms_to_lectures_rooms1_idx` (`rooms_id` ASC),
  INDEX `fk_timeslots_x_rooms_to_lectures_lecture_suggestion_copy21_idx` (`lecture_id` ASC),
  INDEX `fk_timeslots_x_rooms_to_lectures_timeslots1_idx` (`timeslots_id` ASC),
  CONSTRAINT `fk_timeslots_x_rooms_to_lectures_rooms1`
    FOREIGN KEY (`rooms_id`)
    REFERENCES `internal`.`rooms` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_timeslots_x_rooms_to_lectures_lecture_suggestion_copy21`
    FOREIGN KEY (`lecture_id`)
    REFERENCES `internal`.`lectures` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_timeslots_x_rooms_to_lectures_timeslots1`
    FOREIGN KEY (`timeslots_id`)
    REFERENCES `internal`.`timeslots` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'TODO: make sure there is a composite key on this one ( unique combination of timeslot and room, but each can appear many times';


-- -----------------------------------------------------
-- Table `internal`.`messages`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`messages` (
  `message_id` INT NOT NULL AUTO_INCREMENT,
  `from` INT NOT NULL,
  `to` INT NOT NULL,
  `to_group` INT NOT NULL,
  `contents` VARCHAR(255) NULL,
  `title` VARCHAR(225) NULL,
  `urgency` INT NULL,
  `associated_deadline` DATETIME NULL,
  PRIMARY KEY (`message_id`),
  INDEX `fk_messages_users1_idx` (`from` ASC),
  INDEX `fk_messages_groups1_idx` (`to_group` ASC),
  INDEX `fk_messages_1_idx` (`to` ASC),
  CONSTRAINT `fk_messages_users1`
    FOREIGN KEY (`from`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_messages_groups1`
    FOREIGN KEY (`to_group`)
    REFERENCES `internal`.`groups` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_messages_1`
    FOREIGN KEY (`to`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `internal`.`groups_by_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `internal`.`groups_by_users` (
  `groups_id` INT NOT NULL,
  `users_id` INT NOT NULL,
  PRIMARY KEY (`groups_id`, `users_id`),
  INDEX `fk_groups_has_users_users1_idx` (`users_id` ASC),
  INDEX `fk_groups_has_users_groups1_idx` (`groups_id` ASC),
  CONSTRAINT `fk_groups_has_users_groups1`
    FOREIGN KEY (`groups_id`)
    REFERENCES `internal`.`groups` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_groups_has_users_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `internal`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

USE `symp_internal` ;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
