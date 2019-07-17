package pusher

import "time"

type SourceDatabaseHash struct {
	AID         int64     `xorm:"aid bigint autoincr pk"`
	TName       string    `xorm:"table_name varchar(255) notnull"`
	RangeStart  int64     `xorm:"range_start bigint notnull"`
	RangeEnd    int64     `xorm:"range_end bigint notnull"`
	Hash        string    `xorm:"hash varchar(255) notnull"`
	CreatedTime time.Time `xorm:"created_time created notnull"`
	UpdatedTime time.Time `xorm:"updated_time updated notnull"`
}

func (t SourceDatabaseHash) TableName() string {
	return "source_database_hash"
}
