package ltc

import (
	"fmt"
	"time"
)

const (
	TablePrefix = "ltc"
)

type Meta struct {
	ID          int64     `xorm:"id bigint autoincr pk"`
	Name        string    `xorm:"name varchar(255) notnull unique(IDX_ltc_meta_name)"`
	LastID      int64     `xorm:"last_id bigint notnull"`
	Count       int64     `xorm:"count bigint notnull"`
	CreatedTime time.Time `xorm:"created_time created notnull"`
	UpdatedTime time.Time `xorm:"updated_time updated notnull"`
}

func (t Meta) TableName() string {
	return tableName("meta")
}

func tableName(s string) string {
	if len(TablePrefix) == 0 {
		return s
	}
	return fmt.Sprintf("%s_%s", TablePrefix, s)
}
