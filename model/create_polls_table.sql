CREATE TABLE `polls` (
    `title` VARCHAR (64) NOT NULL,
    `option1` VARCHAR(32) NOT NULL,
    `option2` VARCHAR(32) NOT NULL,
    `option3` VARCHAR(32) NOT NULL,
    `option4` VARCHAR(32) NOT NULL,
    `option5` VARCHAR(32) NOT NULL,
    `option1_vote` SMALLINT(2) DEFAULT 0,
    `option2_vote` SMALLINT(2) DEFAULT 0,
    `option3_vote` SMALLINT(2) DEFAULT 0,
    `option4_vote` SMALLINT(2) DEFAULT 0,
    `option5_vote` SMALLINT(2) DEFAULT 0,
    `total_votes` SMALLINT(2) DEFAULT 0,
    `created_by` SMALLINT(4) NOT NULL,
    `poll_status` BIT DEFAULT 1,
    `user_voted` JSON,
    `poll_id` SMALLINT(6) AUTO_INCREMENT PRIMARY KEY UNIQUE 
);