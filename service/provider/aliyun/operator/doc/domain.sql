CREATE TABLE IF NOT EXISTS `domain` (
    `instance_id` varchar(64)  NOT NULL COMMENT '实例Id',
    `domain_name` varchar(64)  NOT NULL COMMENT '域名名称',
    `belong` varchar(64)  NOT NULL COMMENT '所属业务线',
    `domain_audit_status` varchar(64)  NOT NULL COMMENT '域名实名认证状态 FAILED：实名认证失败 SUCCEED：实名认证成功 NONAUDIT：未实名认证 AUDITING：审核中',
    `registration_time` varchar(64)  NOT NULL COMMENT '注册时间',
    `expiration_date_status` varchar(64)  NOT NULL COMMENT '域名过期状态: 1：域名未过期 2：域名已过期',
    `expired_time` varchar(64)  NOT NULL COMMENT '域名到期日期',
    `expiration_curr_date_diff` varchar(64)  NOT NULL COMMENT '域名到期日和当前的时间的天数差值',
    `registrant_type` varchar(64)  NOT NULL COMMENT '域名注册类型: 1：个人 2：企业',
    `domain_status` varchar(64)  NOT NULL COMMENT '域名状态: 1：急需续费 2：急需赎回 3：正常',
    PRIMARY KEY (`domain_name`),
    KEY (`belong`),
    UNIQUE KEY  `idx_instance_id` (`instance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='域名信息表';