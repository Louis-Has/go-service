/*
 Navicat Premium Data Transfer

 Source Server         : Test-mysql
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : testDB

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 24/08/2023 03:21:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_receive_mes
-- ----------------------------
DROP TABLE IF EXISTS `order_receive_mes`;
CREATE TABLE `order_receive_mes`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT comment '收货信息表id',
    `created_at`        datetime     DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at`        datetime     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`        datetime     DEFAULT NULL comment '删除时间',
    `user_id`           bigint(20)      NOT NULL COMMENT '用户id',
    `receiver_name`     varchar(20)  DEFAULT '' COMMENT '收货姓名',
    `receiver_phone`    varchar(20)  DEFAULT '' COMMENT '收货固定电话',
    `receiver_province` varchar(20)  DEFAULT '' COMMENT '省份',
    `receiver_city`     varchar(20)  DEFAULT '' COMMENT '城市',
    `receiver_district` varchar(20)  DEFAULT '' COMMENT '区/县',
    `receiver_address`  varchar(200) DEFAULT '' COMMENT '详细地址',
    PRIMARY KEY (`id`),
    KEY `deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='收货信息表';

SET FOREIGN_KEY_CHECKS = 1;
