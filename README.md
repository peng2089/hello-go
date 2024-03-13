# README

## 运行
```bash
$ docker compose build // 构建镜像
$ docker compose up -d // 启动

// 注册
$ curl -v -d "username=admin&password=123456" http://localhost:8080/auth/register
// 登录
$ curl -v -d "username=admin&password=123456" http://localhost:8080/auth/login
// 获取用户信息
$ curl -H "Authorization: Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteS1hdXRoLXNlcnZlciIsInVpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiJ9.e9aqP0U3Fk4c8qbBiDDQe8TwkowzMgf6xsdlZQMvmAVTPvF4WB19nj784yVpxG5oo1O_Oq9EOM6S6VV3pyyKklipPGEcBFriRtlWV5qI-NotdxwJKXoZ7CqfEyArCA0I0NjtyAwos4lSMr9dv45mEhkK5fs62VRVjTzY1e5MPVBMTa0m8mFu8KalkjRCFtaKduRZvA5EakmH--dqjLVtibLI2xARmvgOKgzY5Bk3Sv-QAbJAdMu8_ftqRtfyxin9Vy-g8EKAAut6wF0dKDjnP7ljmFQocshMi5h6mJ7JNsB46T6jr1wxIKNiZ3TG-6d00zyaJuTlMme9_jn9fszpBA" http://localhost:8080/user/info
```

## 数据库
```sql
DROP TABLE IF EXISTS `tbl_user`;
CREATE TABLE `tbl_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

INSERT INTO `tbl_user` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e');
```

## Note
- [Gin](https://gin-gonic.com/zh-cn/docs/quickstart/)
- [Gorm指南](https://gorm.io/zh_CN/docs/index.html)
- [go-redis](https://redis.io/docs/connect/clients/go/)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)