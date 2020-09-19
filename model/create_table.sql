CREATE TABLE `users`(
    `email` varchar(32) NOT NULL,
    `password` varchar(64) NOT NULL,
    `user_id` SMALLINT(4) AUTO_INCREMENT,
    PRIMARY KEY(user_id)
);