CREATE DATABASE IF NOT EXISTS `CanEat`;

use CanEat;
CREATE TABLE IF NOT EXISTS `foods`
(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `name` VARCHAR(256) NOT NULL,
    `eatable_percent` INT UNSIGNED DEFAULT 100 NOT NULL,
    `energy_kcal` INT UNSIGNED NOT NULL,
    `water_g` FLOAT NOT NULL,
    `protein_g` FLOAT NOT NULL,
    `fat_g` FLOAT NOT NULL,
    `dietary_fiber_g` FLOAT NOT NULL,
    `carbohydrate_g` FLOAT NOT NULL,
    `vitamin_A_ug` FLOAT NOT NULL,
    `vitamin_B1_mg` FLOAT NOT NULL,
    `vitamin_B2_mg` FLOAT NOT NULL,
    `vitamin_B3_mg` FLOAT NOT NULL COMMENT "niacin",
    `vitamin_C_mg` FLOAT NOT NULL,
    `vitamin_E_mg` FLOAT NOT NULL,
    `Na_mg` FLOAT NOT NULL,
    `Ca_mg` FLOAT NOT NULL,
    `Fe_mg` FLOAT NOT NULL,
    `cholesterin_mg` FLOAT NOT NULL,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
