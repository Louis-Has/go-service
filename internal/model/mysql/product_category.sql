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

 Date: 26/08/2023 22:05:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for product_category
-- ----------------------------
DROP TABLE IF EXISTS `product_category`;
CREATE TABLE `product_category`
(
    `id`         smallint unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    `parent_id`  smallint DEFAULT '0' COMMENT '父类别id当id=0时说明是根节点,一级类别',
    `name`       varchar(50)       NOT NULL COMMENT '类别名称',
    PRIMARY KEY (`id`),
    KEY `name` (`name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='商品类别表';

-- ----------------------------
-- Records of product_category
-- ----------------------------
BEGIN;
INSERT INTO `product_category` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `name`)
VALUES (1, '2023-08-25 11:12:58', '2023-08-25 11:12:58', NULL, 0, '干燥食品');
INSERT INTO `product_category` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `name`)
VALUES (2, '2023-08-25 11:13:38', '2023-08-25 11:13:38', NULL, 0, '水果');
INSERT INTO `product_category` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `name`)
VALUES (3, '2023-08-25 17:48:01', '2023-08-25 17:48:01', NULL, 0, '家电');
INSERT INTO `product_category` (`id`, `created_at`, `updated_at`, `deleted_at`, `parent_id`, `name`)
VALUES (4, '2023-08-25 17:49:22', '2023-08-25 17:49:39', NULL, 0, '衣物');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
