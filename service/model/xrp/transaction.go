package xrp

type Transaction struct {
	ID              int64  `xorm:"id bigint autoincr pk"`
	Account         string `xorm:"account varchar(68) notnull index"`
	TransactionType string `xorm:"transaction_type char(30) notnull index"`
	//The type of transaction. Valid types include: Payment, OfferCreate, OfferCancel, TrustSet,
	// AccountSet, SetRegularKey, SignerListSet, EscrowCreate, EscrowFinish, EscrowCancel,
	// PaymentChannelCreate, PaymentChannelFund, PaymentChannelClaim, and DepositPreauth.
	Fee                int64  `xorm:"fee bigint notnull"` //in drops
	Sequence           int64  `xorm:"sequence bigint notnull"`
	AccountTxnID       string `xorm:"account_txn_id varchar(128) null"`
	Flags              int64  `xorm:"flags bigint null"`
	LastLedgerSequence int64  `xorm:"last_ledger_sequence bigint null"`
	Memos              string `xorm:"memos varchar(1024) null"`
	Signers            string `xorm:"signers varchar(2048) null"`
	SourceTag          int64  `xorm:"source_tag bigint null"`
	SigningPubKey      string `xorm:"signing_pub_key varchar(132) null"`
	TxnSignature       string `xorm:"txn_signature varchar(380) null"`

	Hash              string `xorm:"hash varchar(128) notnull index"`
	LedgerIndex       int64  `xorm:"ledger_index bigint index"`
	AffectedNodesLen  int    `xorm:"affected_nodes_len bigint null"`
	TransactionResult string `xorm:"transaction_result varchar(30) null"`
	TransactionIndex  int    `xorm:"transaction_index bigint null"`
	Validated         int    `xorm:"validated tinyint null"`
	Date              int64  `xorm:"date bigint null"`

	//additional
	CloseTime int64 `xorm:"close_time bigint notnull"`

	//Payment
	Amount         int64  `xorm:"amount bigint notnull"` //if xrp, then in drops, else be -1 and ref: fk-Amount
	Destination    string `xorm:"destination varchar(68) null"`
	DestinationTag int64  `xorm:"destination_tag bigint null"`
	InvoiceID      string `xorm:"invoice_id varchar(128) null"`
	//Pathes fk
	PathesLen  int   `xorm:"pathes_len bigint null"`
	SendMax    int64 `xorm:"send_max bigint notnull"`    //if xrp, then in drops, else be -1 and ref: fk-Amount
	DeliverMin int64 `xorm:"deliver_min bigint notnull"` //if xrp, then in drops, else be -1 and ref: fk-Amount

	//OfferCreate  OfferCancel
	Expiration    int64 `xorm:"expiration bigint null"`
	OfferSequence int   `xorm:"offer_sequence bigint null"`
	TakerGets     int64 `xorm:"taker_gets bigint null"` //if xrp, then in drops, else be -1 and ref: fk-Amount
	TakerPays     int64 `xorm:"taker_pays bigint null"` //if xrp, then in drops, else be -1 and ref: fk-Amount

	//TrustSet
	LimitAmount int64 `xorm:"limit_amount bigint null"` //if xrp, then in drops, else be -1 and ref: fk-Amount
	QualityIn   int64 `xorm:"quality_in bigint null"`
	QualityOut  int64 `xorm:"quality_out bigint null"`

	//AccountSet
	ClearFlag    int    `xorm:"clear_flag bigint null"`
	Domain       string `xorm:"domain varchar(512) null"`
	EmailHash    string `xorm:"email_hash char(64) null"`
	MessageKey   string `xorm:"message_key varchar(68) null"`
	SetFlag      int    `xorm:"set_flag bigint null"`
	TransferRate int    `xorm:"transfer_rate bigint null"`
	TickSize     int    `xorm:"tick_size bigint null"`
	//WalletLocator  WalletSize: not used

	//SetRegularKey
	RegularKey string `xorm:"regular_key varchar(68) null"`

	//SignerListSet
	SignerQuorum int `xorm:"signer_quorum bigint null"`

	//EscrowCreate
	CancelAfter int64  `xorm:"cancel_after bigint null"`
	FinishAfter int64  `xorm:"finish_after bigint null"`
	Condition   string `xorm:"condition varchar(512) null"`

	//EscrowFinish
	Owner       string `xorm:"owner varchar(68) null"`
	Fulfillment string `xorm:"fulfillment varchar(512) null"`
	//EscrowCancel
	//PaymentChannelCreate
	SettleDelay int64  `xorm:"settle_delay bigint null"`
	PublicKey   string `xorm:"public_key char(66) null"`
	//PaymentChannelFund
	Channel string `xorm:"channel char(64) null"`
	//PaymentChannelClaim
	Balance int64 `xorm:"balance bigint null"`
	//DepositPreauth
	Authorize   string `xorm:"authorize varchar(68) null"`
	UnAuthorize string `xorm:"un_authorize varchar(68) null"`
}

func (t Transaction) TableName() string {
	return tableName("transaction")
}
