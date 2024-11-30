CREATE DATABASE IF NOT EXISTS blind_box;
CHARACTER SET utf8mb4;

CREATE TABLE `order_tab` (
  `order_id` bigint NOT NULL,
  `seller` varchar(256) NOT NULL,
  `nft_addr` varchar(256) NOT NULL,
  `token_id` bigint NOT NULL,
  `price` decimal(27,0) NOT NULL,
  `amount` bigint NOT NULL,
  `state` tinyint NOT NULL DEFAULT '1' COMMENT '1-active 2-inactive',
  `create_timestamp` bigint NOT NULL COMMENT 'row created block timestamp',
  `update_timestamp` bigint NOT NULL DEFAULT '0' COMMENT 'row updated block timestamp',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'db row create time',
  `update_time` datetime DEFAULT NULL COMMENT 'db row update time',
  PRIMARY KEY (`order_id`),
  KEY `idx_seller` (`seller`),
  KEY `idx_nft_addr_token_id` (`nft_addr`,`token_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

