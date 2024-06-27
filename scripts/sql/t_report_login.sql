CREATE TABLE `t_report_login` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `appId` int DEFAULT NULL COMMENT '产品ID',
  `channelId` int DEFAULT NULL COMMENT '渠道ID',
  `playerId` bigint NOT NULL COMMENT '玩家ID',
  `nickName` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '昵称',
  `type` int DEFAULT NULL COMMENT '类型:0上线,1下线',
  `ip` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '登录IP',
  `loginDevice` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '登陆设备',
  `deviceCode` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '设备IMEI(唯一识别码)',
  `clientVersion` varchar(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '客户端版本号',
  `osInfomation` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '客户端系统信息',
  `beginTime` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '登录时间',
  `duration` bigint NOT NULL COMMENT '在线时间:秒',
  `endTime` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '离线时间',
  `location` varchar(80) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '地区',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `t_report_login_un` (`playerId`,`beginTime`) USING BTREE,
  KEY `t_report_login_playerId_IDX` (`playerId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13776 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='玩家登陆/离线';