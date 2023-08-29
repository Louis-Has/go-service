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
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`
(
    `id`                   bigint unsigned NOT NULL AUTO_INCREMENT comment '订单id',
    `created_at`           datetime       DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at`           datetime       DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`           datetime       DEFAULT NULL comment '删除时间',
    `user_id`              bigint          NOT NULL comment '用户id',
    `order_receive_mes_id` bigint         DEFAULT 0 COMMENT '收货信息表id',
    `payment`              decimal(20, 2) DEFAULT 0 COMMENT '实际付款金额,单位是元,保留两位小数',
    `payment_type`         tinyint(4)     DEFAULT 1 COMMENT '支付类型,1-在线支付',
    `postage`              int(10)        DEFAULT 0 COMMENT '运费,单位是元',
    `status`               smallint(6)    DEFAULT 10 COMMENT '订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭',
    `payment_at`           datetime       DEFAULT NULL comment '支付时间',
    `send_at`              datetime       DEFAULT NULL comment '发货时间',
    `completed_at`         datetime       DEFAULT NULL comment '订单完成时间',
    PRIMARY KEY (`id`),
    KEY `deleted_at` (`deleted_at`),
    key `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
