CREATE TABLE IF NOT EXISTS `prueba-tecnica-reby`.`users` (
  `id_user` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(32) NOT NULL,
  `email` VARCHAR(64) NOT NULL,
  `balance` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id_user`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `prueba-tecnica-reby`.`vehicles` (
  `id_vehicle` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(64) NOT NULL,
  `state` ENUM('free', 'in_use') NOT NULL DEFAULT 'FREE',
  PRIMARY KEY (`id_vehicle`))
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `prueba-tecnica-reby`.`rides` (
  `id_ride` INT NOT NULL AUTO_INCREMENT,
  `created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `finished` DATETIME NOT NULL,
  `id_vehicle` INT NOT NULL,
  `id_user` INT NOT NULL,
  PRIMARY KEY (`id_ride`),
  INDEX `fk_Rides_Vehicles_idx` (`id_vehicle` ASC) VISIBLE,
  INDEX `fk_Rides_Users1_idx` (`id_user` ASC) VISIBLE,
  CONSTRAINT `fk_Rides_Vehicles`
    FOREIGN KEY (`id_vehicle`)
    REFERENCES `prueba-tecnica-reby`.`vehicles` (`id_vehicle`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Rides_Users1`
    FOREIGN KEY (`id_user`)
    REFERENCES `prueba-tecnica-reby`.`users` (`id_user`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;