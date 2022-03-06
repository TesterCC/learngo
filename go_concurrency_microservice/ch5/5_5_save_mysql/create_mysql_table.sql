CREATE DATABASE IF NOT EXISTS `testdb` DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

drop table if exists `tuser`;
CREATE TABLE `tuser` (
    `id` INT(10) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NULL DEFAULT NULL,
    `habits` VARCHAR(128) NULL DEFAULT NULL,
    `created_time` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
);