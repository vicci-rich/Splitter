## Splitter Module Reference

### Directory structure
```$xslt
package_name
├── handler_rpc.go # RPC request wrapper
├── job.go         # Worker job implement
├── meta.go        # Global structure define
├── method.go      # Block data parse implement
├── schema.json    # Kafka data JSON schema
├── splitter.go    # Package entry
└── worker.go      # Background worker implement
```
> package_name: symbol name of blockchain 

### Implement
> example: bitcoin

#### meta.go
###### global structure define
```
type BTCBlockData struct {}
```
###### structure member order
1. prioritization
2. relationship group

#### splitter.go
###### Structure define
```
type SplitterConfig struct {}
type BTCSplitter struct {}
```

###### Implement method
```
func NewSplitter(cfg *Config) (*BTCSplitter, error)
func (p *BTCSplitter) Start()
func (p *BTCSplitter) Stop()
func (p *BTCSplitter) CheckBlock(data *BTCBlockData)
func (p *BTCSplitter) CheckMissedBlock()
func (p *BTCSplitter) Save(data *BTCBlockData)
func (p *BTCSPlitter) RevertBlock(height int64, tx *service.Transaction) error
```

#### method.go
###### Implement method
```
func ParseBlock(data string) (*BTCBlockData, error)
```

#### worker.go
###### Implement method for cron worker
```
func NewCronWorker(splitter *BTCSplitter) *CronWorker
func (w *CronWorker) Prepare() error
func (w *CronWorker) Start() error
func (w *CronWorker) Stop()
```

#### job.go
###### Implement method for update meta data job
```
func newUpdateMetaDataJob(splitter *BTCSplitter) *updateMetaDataJob
func (j *updateMetaDataJob) Run()
func (j *updateMetaDataJob) Name() string
```

#### handler_rpc.go
###### Implement method
```
func newRPCHandler(c *jsonrpc.Client) (*rpcHandler, error)
func (h *rpcHandler) SendBlock(height int64) error
func (h *rpcHandler) SendBatchBlock(startHeight, endHeight int64) error
```

#### schema.json
###### JSON schema validator
- https://www.jsonschemavalidator.net

###### JSON schema reference
- https://ajv.js.org/keywords.html
- https://json-schema.org/understanding-json-schema/index.html

###### JSON schema generator
- https://www.jsonschema.net

###### JSON schema fake data generator
- http://json-schema-faker.js.org
