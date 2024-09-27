/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50743
 Source Host           : localhost:3306
 Source Schema         : ipv6app

 Target Server Type    : MySQL
 Target Server Version : 50743
 File Encoding         : 65001

 Date: 22/09/2024 07:12:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gf_hosts
-- ----------------------------
DROP TABLE IF EXISTS `gf_hosts`;
CREATE TABLE `gf_hosts`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `host_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主机唯一标识',
  `ipv6_address` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主机IPv6地址',
  `free_port` int(11) NOT NULL COMMENT '主机空闲端口',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `host_id`(`host_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
