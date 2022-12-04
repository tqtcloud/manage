CREATE TABLE IF NOT EXISTS `task` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `taskname` varchar(255)  NOT NULL DEFAULT '' COMMENT '同步任务名称',
    `vendor` varchar(255)  NOT NULL DEFAULT '' COMMENT '云厂商:腾讯/阿里/华为',
    `tasktype` varchar(255)  NOT NULL DEFAULT '' COMMENT '任务同步类型，主机/rds/slb',
    `secret_id` int(64)  NOT NULL COMMENT '用于操作资源的ak,sk Id',
    `secret_desc` text NOT NULL COMMENT '凭证描述',
    `region` varchar(64) NOT NULL COMMENT '操作区域Region',
    `taskuser` varchar(64) NOT NULL COMMENT '发起同步的用户',
    `status` varchar(254) NOT NULL COMMENT '任务当前状态',
    `message` text NOT NULL COMMENT '任务失败相关信息',
    `start_at` bigint(20) NOT NULL COMMENT '任务开始时间',
    `end_at` bigint(20) NOT NULL COMMENT '任务结束时间',
    `total_succeed` int(11) NOT NULL COMMENT '总共操作成功的资源数量',
    `total_failed` int(11) NOT NULL COMMENT '总共操作失败的资源数量',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_task_name` (`taskname`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='资源操作任务管理';
