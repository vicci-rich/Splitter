package bch

import (
	"fmt"
)

type Node struct {
	ID                int64  `xorm:"id bigint autoincr pk"`
	IP                string `xorm:"ip varchar(40) notnull unique(address)"`
	Port              string `xorm:"port varchar(5) notnull unique(address)"`
	IPVersion         int    `xorm:"ip_version tinyint notnull"`
	UserAgent         string `xorm:"user_agent varchar(256) notnull"`
	ProtocolVersion   int64  `xorm:"protocol_version bigint notnull"`
	Services          string `xorm:"services varchar(256) notnull"`
	LastSeenTimestamp int64  `xorm:"last_seen_timestamp bigint notnull"`
	LastBlockHeight   int64  `xorm:"last_block_height bigint notnull"`
	CountryName       string `xorm:"country_name varchar(256) notnull"`
	RegionName        string `xorm:"region_name varchar(256) notnull"`
	CityName          string `xorm:"city_name varchar(256) notnull"`
	Owner             string `xorm:"owner varchar(256) notnull"`
	ISP               string `xorm:"isp varchar(256) notnull"`
	Latitude          string `xorm:"latitude varchar(256) notnull"`
	Longitude         string `xorm:"longitude varchar(256) notnull"`
	Timezone          string `xorm:"timezone varchar(256) notnull"`
	UTCOffset         string `xorm:"utc_offset varchar(256) notnull"`
	IDDCode           string `xorm:"idd_code varchar(256) notnull"`
	CountryCode       string `xorm:"country_code varchar(256) notnull"`
	ContinentCode     string `xorm:"continent_code varchar(256) notnull"`
}

func (t Node) TableName() string {
	return tableName("node")
}

func (t Node) String() string {
	return fmt.Sprintf("%s:%s", t.IP, t.Port)
}
