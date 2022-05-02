CREATE TABLE `taxonomy`
(
	`id` INTEGER  NOT NULL AUTO_INCREMENT,
	`usage` VARCHAR(1024),
	`parent_id` INTEGER,
	`source_culture` VARCHAR(16)  NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `taxonomy_FI_2` (`parent_id`)
)Engine=InnoDB;

#-----------------------------------------------------------------------------
#-- taxonomy_i18n
#-----------------------------------------------------------------------------



CREATE TABLE `taxonomy_i18n`
(
	`name` VARCHAR(1024),
	`note` TEXT,
	`id` INTEGER  NOT NULL AUTO_INCREMENT,
	`culture` VARCHAR(16)  NOT NULL,
	PRIMARY KEY (`id`)
)Engine=InnoDB;

