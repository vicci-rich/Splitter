package config

const (
	GlobalSection    = "global"
	ProfilingSection = "profiling"
	MetricSection    = "metric"
	ListenSection    = "listen"
	DeploySection    = "deploy"
	LocationSection  = "location"
	AuthSection      = "auth"
	APISection       = "api"
	AccountSection   = "account"
	DatabaseSection  = "database"
	ProducerSection  = "producer"
	ConsumerSection  = "consumer"
)

type Global struct {
	MaxProcess int
	LocalMode  bool
}

type Profiling struct {
	Enable bool
	Host   string
	Port   string
}

type Metric struct {
	Enable bool
	Host   string
	Port   string
	Path   string
}

type Listen struct {
	Host string
	Port string
}

type Location struct {
	Enable        bool   `ini:"enable"`
	DatabaseFiles string `ini:"database_files"`
}

type Deploy struct {
	Mode string
}

type Auth struct {
	Enable     bool
	AdminKey   string
	EncryptKey string
}

type Database struct {
	Type         string
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
	SQLLogFile   string `ini:"sql_log_file"`
	Debug        bool
}

type KafkaProducer struct {
	BrokerList       string
	BufferSize       int
	ProducerNum      int
	FlushMessages    int
	FlushFrequency   int
	FlushMaxMessages int
	Timeout          int
	ReturnErrors     bool
}

type KafkaConsumer struct {
	BrokerList   string
	BufferSize   int
	ClientID     string `ini:"client_id"`
	GroupID      string `ini:"group_id"`
	ReturnErrors bool
}
