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

 Date: 26/08/2023 00:49:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_mes
-- ----------------------------
DROP TABLE IF EXISTS `user_mes`;
CREATE TABLE `user_mes`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `created_at`    datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`    datetime DEFAULT NULL COMMENT '删除时间',
    `user_name`     varchar(255)    NOT NULL COMMENT '用户名',
    `signed_person` tinyint  DEFAULT '0' COMMENT '标记用户，0表示默认；1表示正式用户；2表示VIP',
    PRIMARY KEY (`id`),
    KEY `deleted_at` (`deleted_at`),
    KEY `user_name` (`user_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 7
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user_mes
-- ----------------------------
BEGIN;
INSERT INTO `user_mes` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_name`, `signed_person`)
VALUES (1, '2023-04-01 15:31:00', '2023-04-01 15:31:00', NULL, '唐家', 0);
INSERT INTO `user_mes` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_name`, `signed_person`)
VALUES (2, '2023-04-01 15:35:02', '2023-04-01 15:35:02', NULL, '朱朱', 1);
INSERT INTO `user_mes` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_name`, `signed_person`)
VALUES (3, '2023-04-01 15:36:30', '2023-04-01 15:36:30', NULL, '朱朱23', 1);
INSERT INTO `user_mes` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_name`, `signed_person`)
VALUES (4, '2023-04-01 16:01:42', '2023-04-01 16:01:42', NULL, '小鱼', 1);
INSERT INTO `user_mes` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_name`, `signed_person`)
VALUES (5, '2023-08-25 16:45:14', '2023-08-25 16:48:42', NULL, '天天', 0);
INSERT INTO `user_mes` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_name`, `signed_person`)
VALUES (6, '2023-08-25 16:45:14', '2023-08-25 16:48:42', NULL, '天天2', 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
