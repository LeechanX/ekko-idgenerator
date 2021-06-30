
CREATE TABLE `global_worker_id` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `mac_address` BIGINT(20),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_mac` (`mac_address`) USING BTREE
);

CREATE TABLE `time_report` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `mac_address` BIGINT(20),
  `ts` BIGINT(20),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_mac` (`mac_address`) USING BTREE
);
