CREATE TABLE `users`
(
    `email` varchar
(32) NOT NULL UNIQUE,
    `password` varchar
(64) NOT NULL,
    `user_id` SMALLINT
(4) AUTO_INCREMENT UNIQUE,
    PRIMARY KEY
(user_id)
);