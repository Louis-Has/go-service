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

 Date: 26/08/2023 22:04:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品ID',
    `created_at`   datetime       DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   datetime       DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`   datetime       DEFAULT NULL COMMENT '删除时间',
    `product_name` varchar(100)    NOT NULL COMMENT '商品名称',
    `subtitle`     varchar(200)   DEFAULT '' COMMENT '商品副标题',
    `cate_id`      smallint       DEFAULT '0' COMMENT '类别ID',
    `price`        decimal(20, 2) DEFAULT '0.00' COMMENT '价格,单位-元,保留两位小数',
    `stock`        int            DEFAULT '0' COMMENT '库存数量',
    `status`       int            DEFAULT '0' COMMENT '商品状态.0-待上架 1-在售 2-下架 3-过期',
    PRIMARY KEY (`id`),
    KEY `product_name` (`product_name`),
    KEY `cate_id` (`cate_id`),
    KEY `price` (`price`),
    KEY `stock` (`stock`),
    KEY `status` (`status`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 8
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='商品表';

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO `product` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_name`, `subtitle`, `cate_id`, `price`,
                       `stock`, `status`)
VALUES (1, '2023-08-25 11:02:09', '2023-08-25 11:02:09', NULL, '趣多多', '', 1, 12.00, 800, 0);
INSERT INTO `product` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_name`, `subtitle`, `cate_id`, `price`,
                       `stock`, `status`)
VALUES (2, '2023-08-25 11:05:59', '2023-08-25 11:05:59', NULL, '苏打饼干', '', 1, 8.00, 200, 1);
INSERT INTO `product` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_name`, `subtitle`, `cate_id`, `price`,
                       `stock`, `status`)
VALUES (3, '2023-08-25 17:47:37', '2023-08-25 18:02:03', NULL, '洗衣机', '', 3, 600.00, 1, 0);
INSERT INTO `product` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_name`, `subtitle`, `cate_id`, `price`,
                       `stock`, `status`)
VALUES (4, '2023-08-25 17:48:29', '2023-08-25 18:02:07', NULL, '香蕉', '', 2, 6.00, 24, 1);
INSERT INTO `product` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_name`, `subtitle`, `cate_id`, `price`,
                       `stock`, `status`)
VALUES (5, '2023-08-25 17:48:56', '2023-08-25 18:01:50', NULL, '冰箱', '', 3, 2400.00, 2, 0);
INSERT INTO `product` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_name`, `subtitle`, `cate_id`, `price`,
                       `stock`, `status`)
VALUES (6, '2023-08-25 17:49:14', '2023-08-25 18:01:41', NULL, '电风扇', '', 3, 80.00, 10, 0);
INSERT INTO `product` (`id`, `created_at`, `updated_at`, `deleted_at`, `product_name`, `subtitle`, `cate_id`, `price`,
                       `stock`, `status`)
VALUES (7, '2023-08-25 17:49:52', '2023-08-25 18:01:33', NULL, '夹克', '', 4, 200.00, 12, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
