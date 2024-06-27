CREATE TABLE `t_withdraw_record` (
  `n_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `n_orderid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '订单ID',
  `n_playerid` bigint unsigned DEFAULT NULL COMMENT '玩家ID',
  `n_amount` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '提现金额，小数点2位',
  `n_state` smallint DEFAULT NULL COMMENT '订单状态: 0=审核中, 1=审核成功,2=提现成功,-1=审核失败,-2=提现失败,-3=提现提交失败',
  `n_des` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '订单失败描述',
  `n_createtime` datetime DEFAULT NULL COMMENT '订单创建时间',
  `n_paytime` datetime DEFAULT NULL COMMENT '订单放款时间',
  `n_checktime` datetime DEFAULT NULL COMMENT '订单审核时间',
  `n_coin` bigint DEFAULT NULL COMMENT '扣除金币数',
  `n_channel` int DEFAULT NULL COMMENT '用户渠道号',
  `n_ex_orderid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '外部订单号',
  `n_bankinfo` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '银行卡信息: 姓名和卡号, json',
  `n_ex_resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '外部支付平台回复内容',
  `n_back_coin` int DEFAULT '0' COMMENT '提现失败，返回金币是否成功: 1=成功, -1=失败,0=未返回。',
  `n_operator` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '操作人',
  `n_fee` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '提现费用',
  `n_from` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '来源',
  `n_partnerid` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '商户ID',
  `n_platform` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '充值平台',
  `n_after` bigint DEFAULT NULL COMMENT '提现之后的金额',
  `n_kind` int DEFAULT NULL COMMENT '提现银行卡类型: UPI, MPS',
  `n_show_err` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户端显示错误',
  `n_card_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '银行卡信息',
  `n_vpay_createtime` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT '' COMMENT '回调createtime',
  `n_retry` int DEFAULT '0' COMMENT '重试次数',
  `n_retry_state` int DEFAULT '0' COMMENT '重试提交状态: 0=未检测, 1=自动提交,2=手动审核',
  `n_query_state` int DEFAULT '0' COMMENT '重试提交状态: 0=未检测, 1=自动提交,2=手动审核',
  `n_query_resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '查询回复',
  `n_last_order_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '最后的订单ID',
  `n_red` tinyint DEFAULT '0' COMMENT '红点状态',
  `n_wd_flag` tinyint DEFAULT '0' COMMENT '提现标识 0普通 1代理',
  `n_user_flag` tinyint DEFAULT '0' COMMENT '0普通 1优质',
  PRIMARY KEY (`n_id`) USING BTREE,
  UNIQUE KEY `t_recharge_record_un` (`n_orderid`) USING BTREE,
  KEY `t_recharge_record_n_playerid_IDX` (`n_playerid`) USING BTREE,
  KEY `t_withdraw_record_n_state_IDX` (`n_state`) USING BTREE,
  KEY `vpay_createtime_IDX` (`n_vpay_createtime`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=778 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC COMMENT='提现记录表';