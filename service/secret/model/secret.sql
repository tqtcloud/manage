CREATE TABLE `secret` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `vendor` varchar(255)  NOT NULL DEFAULT '' COMMENT '云厂商:腾讯/阿里/华为',
    `accessKeyId` varchar(255)  NOT NULL DEFAULT '' COMMENT '云厂商 AK accessKeyId',
    `accessKeySecret` varchar(255)  NOT NULL DEFAULT '' COMMENT '云厂商 SK accessKeySecret',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_mobile_unique` (`accessKeyId`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
