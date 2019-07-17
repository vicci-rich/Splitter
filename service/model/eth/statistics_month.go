package eth

/*
DROP TABLE IF EXISTS `eth_month`;
CREATE TABLE `eth_month` (
  `id`                   INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `time`                 VARCHAR(20)      NOT NULL COMMENT '日期',
  `active_address_count` INT(10) UNSIGNED NOT NULL COMMENT '月活',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_time` (`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type StatisticsMonth struct {
	ID                      int64 `xorm:"id bigint autoincr pk"`
	Timestamp               int64 `xorm:"timestamp int notnull default '0' unique(IDX_timestamp)"`
	ActiveAddressCount      int   `xorm:"active_address_count int notnull default '0'"`
	TotalActiveAddressCount int   `xorm:"total_active_address_count int notnull default '0'"`
}

func (t StatisticsMonth) TableName() string {
	return tableName("statistics_month")
}
