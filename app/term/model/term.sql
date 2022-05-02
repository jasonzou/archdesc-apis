CREATE TABLE `term`
(
	`id` INTEGER  NOT NULL AUTO_INCREMENT,
	`taxonomy_id` INTEGER  NOT NULL,
	`code` VARCHAR(1024),
	`parent_id` INTEGER,
	`lft` INTEGER,
	`rgt` INTEGER,
	`source_culture` VARCHAR(16)  NOT NULL,
	PRIMARY KEY (`id`),
	KEY `lft`(`lft`),
	INDEX `term_FI_2` (`taxonomy_id`),
	INDEX `term_FI_3` (`parent_id`)
)Engine=InnoDB;

#-----------------------------------------------------------------------------
#-- term_i18n
#-----------------------------------------------------------------------------



CREATE TABLE `term_i18n`
(
	`name` VARCHAR(1024),
	`id` INTEGER  NOT NULL AUTO_INCREMENT,
	`culture` VARCHAR(16)  NOT NULL,
	PRIMARY KEY (`id`)
)Engine=InnoDB;

