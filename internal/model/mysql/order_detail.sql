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
-- Table structure for order_detail
-- ----------------------------
DROP TABLE IF EXISTS `order_detail`;
CREATE TABLE `order_detail`
(
    `id`                 bigint unsigned NOT NULL AUTO_INCREMENT comment '订单明细id',
    `created_at`         datetime DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `updated_at`         datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`         datetime DEFAULT NULL comment '删除时间',
    `order_id`           bigint          NOT NULL comment '订单id',
    `user_id`            varchar(255)    NOT NULL comment '用户id',
    `product_id`         bigint          NOT NULL comment '商品id',
    `current_unit_price` decimal(20, 2)  NOT NULL comment '生成订单时的商品单价，单位是元,保留两位小数',
    `product_quantity`   bigint          NOT NULL comment '商品数量',
    `total_price`        decimal(20, 2)  NOT NULL COMMENT '商品总价,单位是元,保留两位小数',
    PRIMARY KEY (`id`),
    KEY `deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='订单明细表';

SET FOREIGN_KEY_CHECKS = 1;
