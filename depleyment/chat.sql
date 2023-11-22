DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uuid` VARCHAR(255) NOT NULL COMMENT 'uuid',
    `username` VARCHAR(150) NOT NULL COMMENT 'username',
    `nickname` VARCHAR(150) DEFAULT NULL COMMENT 'nickname',
    `email` VARCHAR(255) DEFAULT NULL COMMENT 'EMAIL',
    `password` VARCHAR(150) NOT NULL COMMENT 'password',
    `avatar` VARCHAR(250) DEFAULT NULL COMMENT 'avatar',
    `create_time` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'CREATE TIME',
    `update_time` TIMESTAMP(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'UPDATE TIME',
    `delete_time` DATETIME(3) DEFAULT NULL COMMENT 'DELETE TIME',
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `idx_uuid` (`uuid`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '用户表';

DROP TABLE IF EXISTS `user_friends`;

CREATE TABLE IF NOT EXISTS `user_friends` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` INT DEFAULT NULL COMMENT 'USER ID',
    `friend_id` INT DEFAULT NULL COMMENT 'FRIEND ID',
    `create_time` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'CREATE TIME',
    `update_time` TIMESTAMP(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'UPDATE TIME',
    `delete_time` DATETIME(3) DEFAULT NULL COMMENT 'DELETE TIME',
    PRIMARY KEY (`id`),
    KEY `idx_user_friends_user_id` (`user_id`),
    KEY `idx_user_friends_friends_id` (`friend_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '用户好友关系表';

DROP TABLE IF EXISTS `message`;

CREATE TABLE IF NOT EXISTS `message` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'id',
    `from_user_id` INT DEFAULT NULL COMMENT '发送方id',
    `to_user_id` INT DEFAULT NULL COMMENT '目标用户id',
    `content` VARCHAR(2500) DEFAULT NULL COMMENT '消息内容',
    `url` VARCHAR(350) DEFAULT NULL COMMENT '内容地址',
    `pic` TEXT COMMENT '缩略图',
    `message_type` SMALLINT DEFAULT NULL COMMENT '消息类型：1.单聊，2.群聊',
    `content_type` SMALLINT DEFAULT NULL COMMENT '消息内容类型：1.text 2.file 3.image 4.audio 5.video',
    `delete_time` DATETIME(3) DEFAULT NULL COMMENT 'DELETE TIME',
    `create_time` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'create time',
    `update_time` TIMESTAMP(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'UPDATE TIME',   
    PRIMARY KEY (`id`),
    KEY `idx_message_to_user_id` (`to_user_id`),
    KEY `idx_message_from_user_id` (`from_user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '消息表';

DROP TABLE IF EXISTS `group`;

CREATE TABLE IF NOT EXISTS `group` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'id',
    `admin_id` INT DEFAULT NULL COMMENT 'ADMIN ID',
    `name` VARCHAR(150) DEFAULT NULL COMMENT 'group name',
    `notice` VARCHAR(300) DEFAULT NULL COMMENT 'group notice',
    `uuid` VARCHAR(150) DEFAULT NULL COMMENT 'uuid',
    `create_time` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'CREATE TIME',
    `update_time` TIMESTAMP(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'UPDATE TIME',
    `delete_time` DATETIME(3) DEFAULT NULL COMMENT 'delete time',
    PRIMARY KEY (`id`),
    KEY `idx_groups_admin_id` (`admin_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '群组表';

CREATE TABLE IF NOT EXISTS `group_member` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id` INT DEFAULT NULL COMMENT 'USER ID',
    `group_id` INT DEFAULT NULL COMMENT 'GROUP ID',
    `nickname` VARCHAR(150) DEFAULT NULL COMMENT 'USER NICKNAME',
    `create_time` TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3) COMMENT 'CREATE TIME',
    `update_time` TIMESTAMP(3) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'UPDATE TIME',
    `delete_time` BIGINT UNSIGNED DEFAULT NULL COMMENT 'DELETE TIME',
    `mute` SMALLINT DEFAULT NULL COMMENT '是否是禁言',
    PRIMARY KEY (`id`),
    KEY `idx_group_member_user_id` (`user_id`),
    KEY `idx_group_member_group_id` (`group_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '用户群组表';
