-- MySQL Script generated by MySQL Workbench
-- Tue Oct 31 16:36:53 2023
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema melodymeter
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `melodymeter` ;

-- -----------------------------------------------------
-- Schema melodymeter
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `melodymeter` DEFAULT CHARACTER SET utf8 ;
USE `melodymeter` ;

-- -----------------------------------------------------
-- Table `melodymeter`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `melodymeter`.`users` ;

CREATE TABLE IF NOT EXISTS `melodymeter`.`users` (
  `id` INT NOT NULL,
  `name` VARCHAR(245) NOT NULL,
  `email` VARCHAR(245) NOT NULL,
  `isadm` TINYINT(1) NOT NULL DEFAULT 0,
  `password` VARCHAR(245) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `melodymeter`.`albuns`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `melodymeter`.`albuns` ;

CREATE TABLE IF NOT EXISTS `melodymeter`.`albuns` (
  `id` INT NOT NULL,
  `name` VARCHAR(245) NOT NULL,
  `imagem` VARCHAR(245) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `melodymeter`.`songs`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `melodymeter`.`songs` ;

CREATE TABLE IF NOT EXISTS `melodymeter`.`songs` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  `description` VARCHAR(45) NULL,
  `author` VARCHAR(45) NULL,
  `year` VARCHAR(45) NULL,
  `duration` VARCHAR(45) NULL,
  `albuns_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_songs_albuns`
    FOREIGN KEY (`albuns_id`)
    REFERENCES `melodymeter`.`albuns` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `melodymeter`.`rates`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `melodymeter`.`rates` ;

CREATE TABLE IF NOT EXISTS `melodymeter`.`rates` (
  `id` INT NOT NULL,
  `rate` VARCHAR(2) NULL,
  `comment` VARCHAR(245) NULL,
  `songs_id` INT NOT NULL,
  `users_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_rates_songs1`
    FOREIGN KEY (`songs_id`)
    REFERENCES `melodymeter`.`songs` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_rates_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `melodymeter`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;