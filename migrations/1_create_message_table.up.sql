CREATE TABLE IF NOT EXISTS `messages`
(
    `id`         integer PRIMARY KEY AUTO_INCREMENT,
    `type`       varchar(20)        NOT NULL,
    `message`       varchar(50)     NOT NULL
);