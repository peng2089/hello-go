SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";
SET NAMES utf8mb4;


-- 数据库 hello-go start --
CREATE DATABASE IF NOT EXISTS `hello-go` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `hello-go`;
CREATE TABLE IF NOT EXISTS `tbl_user` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
-- 数据库 hello-go end --