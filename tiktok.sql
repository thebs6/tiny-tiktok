CREATE DATABASE IF NOT EXISTS `tiktok` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `tiktok`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `uid`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`        varchar(128)        NOT NULL DEFAULT '' COMMENT '用户昵称',
    `password`    varchar(128)        NOT NULL DEFAULT 1  COMMENT '密码',
    `avatar`      varchar(128)        NOT NULL DEFAULT '' COMMENT '头像',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user`
VALUES (1, 'Jerry', 'password', '', '2022-04-01 10:00:00', '2022-04-01 10:00:00'),
       (2, 'Tom', 'password', '', '2022-04-01 10:00:00', '2022-04-01 10:00:00');