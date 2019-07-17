package eos

import (
	"github.com/jdcloud-bds/bds/common/cuckoofilter"
	"github.com/jdcloud-bds/bds/common/metric"
	model "github.com/jdcloud-bds/bds/service/model/eos"
)

const (
	MetricReceiveMessages                       = "receive_messages"
	MetricParseDataError                        = "parse_data_error"
	MetricVaildationSuccess                     = "validation_success"
	MetricVaildationError                       = "validation_error"
	MetricDatabaseRollback                      = "database_rollback"
	MetricDatabaseCommit                        = "database_commit"
	MetricCronWorkerJob                         = "cron_worker_job"
	MetricCronWorkerJobUpdateMetaData           = "cron_worker_job_update_meta_data"
	MetricCronWorkerJobRefreshContractAddresses = "cron_worker_job_refresh_contract_addresses"
	MetricCronWorkerJobRefreshPoolName          = "cron_worker_job_refresh_pool_name"
	MetricRPCCall                               = "rpc_call"
	MetricRevertBlock                           = "revert_block"
	TimeYmdHmssFormatISO                        = "2006-01-02T15:04:05.000"
)

var (
	stats                 = metric.NewMap("eos")
	contractAddressFilter = cuckoofilter.New()
)

type EOSBlockData struct {
	Block        *model.Block
	Transactions []*model.Transaction
	Actions      []*model.Action
}
