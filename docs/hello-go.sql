-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- 主机： mysql:3306
-- 生成日期： 2024-01-05 06:35:26
-- 服务器版本： 8.0.25
-- PHP 版本： 8.2.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `hello-go`
--

--
-- 表的结构 `tbl_user`
--

CREATE TABLE IF NOT EXISTS `tbl_user` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `nickname` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户属性表';

-- --------------------------------------------------------

--
-- 表的结构 `tbl_user_api_auth`
--

CREATE TABLE IF NOT EXISTS `tbl_user_api_auth` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL COMMENT '用户id',
  `api_key` char(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'key',
  `api_secret` char(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'secret',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='认证表：api认证';

-- --------------------------------------------------------

--
-- 表的结构 `tbl_user_local_auth`
--

CREATE TABLE IF NOT EXISTS `tbl_user_local_auth` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL COMMENT '用户id',
  `username` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` char(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `created_at` int NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='认证表：本地认证LocalAuth';

-- --------------------------------------------------------

--
-- 表的结构 `tbl_user_oauth`
--

CREATE TABLE IF NOT EXISTS `tbl_user_oauth` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL COMMENT '用户id',
  `oauth_name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL,
  `oauth_id` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `oauth_access_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'access token',
  `oauth_expires` int NOT NULL COMMENT 'token有效期',
  `created_at` int NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='认证表：OAuth认证';
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
