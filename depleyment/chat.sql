CREATE TABEL GOCHAT;

DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uuid` varchar(255) NOT NULL COMMENT 'uuid',
    `username` varchar(150) NOT NULL COMMENT 'username',
    `nickname` varchar(150) DEFAULT NULL COMMENT 'nickname',
    `email` varchar(255) DEFAULT NULL COMMENT 'EMAIL',
    `password` varchar(150) NOT NULL COMMENT 'password',
    `avator` varchar(250) DEFAULT NULL COMMENT 'avator',
    `create_time` DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'CREATE TIME',
    `update_time` DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'UPDATE TIME',
    `delete_time` datetime(3) DEFAULT NULL COMMENT 'DELETE TIME',
    PRIMARY KEY(`id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `idx_uuid` (`uuid`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '用户表';

DROP TABLE IF EXISTS `user_friends`;

CREATE TABLE IF NOT EXISTS `user_friends`(
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` INT DEFAULT NULL COMMENT 'USER ID',
    `friend_id` INT DEFAULT NULL COMMENT 'FRIEND ID',
    `create_time` DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'CREATE TIME',
    `update_time` DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'UPDATE TIME',
    `delete_time` datetime(3) DEFAULT NULL COMMENT 'DELETE TIME',
    PRIMARY KEY(`id`),
    KEY `idx_user_friends_user_id`(`user_id`),
    KEY `idex_user_friends_friends_id`(`friend_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '用户好友关系表';

CREATE TABLE IF EXISTS `message` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'id',
    `from_user_id` INT DEFAULT NULL COMMENT '发送方id',
    `to_user_id` INT DEFAULT NULL COMMENT '目标用户id',
    `content` varchar(2500) DEFAULT NULL COMMENT '消息内容',
    `url` varchar(350) DEFAULT NULL COMMENT '内容地址',
    `pic` text COMMENT '缩略图',
    `message_type` SMALLINT DEFAULT NULL COMMENT '消息类型：1.单聊，2.群聊',
    `conent_type` SMALLINT default null comment '消息内容类型：1.text 2.file 3.image 4.audio 5.vidio',
    `delete_time` datetime(3) DEFAULT NULL COMMENT 'DELETE TIME',
    `create_time` DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'create time',
    `update_time` DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'UPDATE TIME',   
    PRIMARY KEY(`id`),
    KEY `idx_message_to_user_id` (`to_user_id`),
    KEY `idex_message_from_user_id`(`from_user_id`)
) engine =InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '消息表'




