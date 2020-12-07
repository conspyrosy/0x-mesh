// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqltypes

import (
	"fmt"
	"io"
	"strconv"
)

type AcceptedOrderResult struct {
	// The order that was accepted, including metadata.
	Order *OrderWithMetadata `json:"order"`
	// Whether or not the order is new. Set to true if this is the first time this Mesh node has accepted the order
	// and false otherwise.
	IsNew bool `json:"isNew"`
}

type AddOrdersOpts struct {
	// Pinned orders will not be affected by any DDoS prevention or incentive
	// mechanisms and will stay in storage until they are no longer fillable
	// (assuming the following flags are all false). Defaults to true.
	Pinned *bool `json:"pinned"`
	// Indicates that the orders being added should be kept by the database after
	// cancellation.
	KeepCancelled *bool `json:"keepCancelled"`
	// Indicates that the orders being added should be kept by the database after
	// expiry.
	KeepExpired *bool `json:"keepExpired"`
	// Indicates that the orders being added should be kept by the database after
	// being fully filled.
	KeepFullyFilled *bool `json:"keepFullyFilled"`
	// Indicates that the orders being added should be kept by the database after
	// becoming unfunded.
	KeepUnfunded *bool `json:"keepUnfunded"`
}

// The results of the addOrders mutation. Includes which orders were accepted and which orders where rejected.
type AddOrdersResults struct {
	// The set of orders that were accepted. Accepted orders will be watched and order events will be emitted if
	// their status changes.
	Accepted []*AcceptedOrderResult `json:"accepted"`
	// The set of orders that were rejected, including the reason they were rejected. Rejected orders will not be
	// watched.
	Rejected []*RejectedOrderResult `json:"rejected"`
}

// An on-chain contract event.
type ContractEvent struct {
	// The hash of the block where the event was generated.
	BlockHash string `json:"blockHash"`
	// The hash of the transaction where the event was generated.
	TxHash string `json:"txHash"`
	// The index of the transaction where the event was generated.
	TxIndex int `json:"txIndex"`
	// The index of the event log.
	LogIndex int `json:"logIndex"`
	// True when this was an event that was removed due to a block-reorg. False otherwise.
	IsRemoved bool `json:"isRemoved"`
	// The address of the contract that generated the event.
	Address string `json:"address"`
	// The kind of event (e.g. "ERC20TransferEvent").
	Kind string `json:"kind"`
	// The parameters for the event. The parameters are different for each event kind, but will always
	// be a set of key-value pairs.
	Parameters interface{} `json:"parameters"`
}

// The block number and block hash for the latest block that has been processed by Mesh.
type LatestBlock struct {
	// The block number encoded as a numerical string.
	Number string `json:"number"`
	// The block hash encoded as a hexadecimal string.
	Hash string `json:"hash"`
}

// A signed 0x order according to the [protocol specification](https://github.com/0xProject/0x-protocol-specification/blob/master/v3/v3-specification.md#order-message-format).
type NewOrder struct {
	ChainID               string `json:"chainId"`
	ExchangeAddress       string `json:"exchangeAddress"`
	MakerAddress          string `json:"makerAddress"`
	MakerAssetData        string `json:"makerAssetData"`
	MakerAssetAmount      string `json:"makerAssetAmount"`
	MakerFeeAssetData     string `json:"makerFeeAssetData"`
	MakerFee              string `json:"makerFee"`
	TakerAddress          string `json:"takerAddress"`
	TakerAssetData        string `json:"takerAssetData"`
	TakerAssetAmount      string `json:"takerAssetAmount"`
	TakerFeeAssetData     string `json:"takerFeeAssetData"`
	TakerFee              string `json:"takerFee"`
	SenderAddress         string `json:"senderAddress"`
	FeeRecipientAddress   string `json:"feeRecipientAddress"`
	ExpirationTimeSeconds string `json:"expirationTimeSeconds"`
	Salt                  string `json:"salt"`
	Signature             string `json:"signature"`
}

// A signed 0x order according to the [protocol specification](https://github.com/0xProject/0x-protocol-specification/blob/master/v3/v3-specification.md#order-message-format.)
type Order struct {
	ChainID               string `json:"chainId"`
	ExchangeAddress       string `json:"exchangeAddress"`
	MakerAddress          string `json:"makerAddress"`
	MakerAssetData        string `json:"makerAssetData"`
	MakerAssetAmount      string `json:"makerAssetAmount"`
	MakerFeeAssetData     string `json:"makerFeeAssetData"`
	MakerFee              string `json:"makerFee"`
	TakerAddress          string `json:"takerAddress"`
	TakerAssetData        string `json:"takerAssetData"`
	TakerAssetAmount      string `json:"takerAssetAmount"`
	TakerFeeAssetData     string `json:"takerFeeAssetData"`
	TakerFee              string `json:"takerFee"`
	SenderAddress         string `json:"senderAddress"`
	FeeRecipientAddress   string `json:"feeRecipientAddress"`
	ExpirationTimeSeconds string `json:"expirationTimeSeconds"`
	Salt                  string `json:"salt"`
	Signature             string `json:"signature"`
}

type OrderEvent struct {
	// The order that was affected.
	Order *OrderWithMetadata `json:"order"`
	// A way of classifying the effect that the order event had on the order. You can
	// think of different end states as different "types" of order events.
	EndState OrderEndState `json:"endState"`
	// The timestamp for the order event, which can be used for bookkeeping purposes.
	// If the order event was generated as a direct result of on-chain events (e.g., FILLED,
	// UNFUNDED, CANCELED), then it is set to the latest block timestamp at which the order
	// was re-validated. Otherwise (e.g., for ADDED, STOPPED_WATCHING), the timestamp corresponds
	// when the event was generated on the server side.
	Timestamp string `json:"timestamp"`
	// Contains all the contract events that triggered the order to be re-validated.
	// All events that _may_ have affected the state of the order are included here.
	// It is guaranteed that at least one of the events included here will have affected
	// the order's state, but there may also be some false positives.
	ContractEvents []*ContractEvent `json:"contractEvents"`
}

// A filter on orders. Can be used in queries to only return orders that meet certain criteria.
type OrderFilter struct {
	Field OrderField `json:"field"`
	Kind  FilterKind `json:"kind"`
	// value must match the type of the filter field.
	Value interface{} `json:"value"`
}

// A sort ordering for orders. Can be used in queries to control the order in which results are returned.
type OrderSort struct {
	Field     OrderField    `json:"field"`
	Direction SortDirection `json:"direction"`
}

// A signed 0x order along with some additional metadata about the order which is not part of the 0x protocol specification.
type OrderWithMetadata struct {
	ChainID               string `json:"chainId"`
	ExchangeAddress       string `json:"exchangeAddress"`
	MakerAddress          string `json:"makerAddress"`
	MakerAssetData        string `json:"makerAssetData"`
	MakerAssetAmount      string `json:"makerAssetAmount"`
	MakerFeeAssetData     string `json:"makerFeeAssetData"`
	MakerFee              string `json:"makerFee"`
	TakerAddress          string `json:"takerAddress"`
	TakerAssetData        string `json:"takerAssetData"`
	TakerAssetAmount      string `json:"takerAssetAmount"`
	TakerFeeAssetData     string `json:"takerFeeAssetData"`
	TakerFee              string `json:"takerFee"`
	SenderAddress         string `json:"senderAddress"`
	FeeRecipientAddress   string `json:"feeRecipientAddress"`
	ExpirationTimeSeconds string `json:"expirationTimeSeconds"`
	Salt                  string `json:"salt"`
	Signature             string `json:"signature"`
	// The hash, which can be used to uniquely identify an order. Encoded as a hexadecimal string.
	Hash string `json:"hash"`
	// The remaining amount of the maker asset which has not yet been filled. Encoded as a numerical string.
	FillableTakerAssetAmount string `json:"fillableTakerAssetAmount"`
}

type RejectedOrderResult struct {
	// The hash of the order. May be null if the hash could not be computed.
	Hash *string `json:"hash"`
	// The order that was rejected.
	Order *Order `json:"order"`
	// A machine-readable code indicating why the order was rejected. This code is designed to
	// be used by programs and applications and will never change without breaking backwards-compatibility.
	Code RejectedOrderCode `json:"code"`
	// A human-readable message indicating why the order was rejected. This message may change
	// in future releases and is not covered by backwards-compatibility guarantees.
	Message string `json:"message"`
}

// Contains configuration options and various stats for Mesh.
type Stats struct {
	Version                           string       `json:"version"`
	PubSubTopicsV3                    []string     `json:"pubSubTopicsV3"`
	PubSubTopicsV4                    []string     `json:"pubSubTopicsV4"`
	Rendezvous                        string       `json:"rendezvous"`
	PeerID                            string       `json:"peerID"`
	EthereumChainID                   int          `json:"ethereumChainID"`
	LatestBlock                       *LatestBlock `json:"latestBlock"`
	NumPeers                          int          `json:"numPeers"`
	NumOrders                         int          `json:"numOrders"`
	NumOrdersIncludingRemoved         int          `json:"numOrdersIncludingRemoved"`
	StartOfCurrentUTCDay              string       `json:"startOfCurrentUTCDay"`
	EthRPCRequestsSentInCurrentUTCDay int          `json:"ethRPCRequestsSentInCurrentUTCDay"`
	EthRPCRateLimitExpiredRequests    int          `json:"ethRPCRateLimitExpiredRequests"`
	SecondaryRendezvous               []string     `json:"secondaryRendezvous"`
	// The max expiration time expressed as seconds since the Unix Epoch and encoded as a numerical string.
	// Any order with an expiration time greater than this maximum will be rejected by Mesh.
	MaxExpirationTime string `json:"maxExpirationTime"`
}

// The kind of comparison to be used in a filter.
type FilterKind string

const (
	FilterKindEqual          FilterKind = "EQUAL"
	FilterKindNotEqual       FilterKind = "NOT_EQUAL"
	FilterKindGreater        FilterKind = "GREATER"
	FilterKindGreaterOrEqual FilterKind = "GREATER_OR_EQUAL"
	FilterKindLess           FilterKind = "LESS"
	FilterKindLessOrEqual    FilterKind = "LESS_OR_EQUAL"
)

var AllFilterKind = []FilterKind{
	FilterKindEqual,
	FilterKindNotEqual,
	FilterKindGreater,
	FilterKindGreaterOrEqual,
	FilterKindLess,
	FilterKindLessOrEqual,
}

func (e FilterKind) IsValid() bool {
	switch e {
	case FilterKindEqual, FilterKindNotEqual, FilterKindGreater, FilterKindGreaterOrEqual, FilterKindLess, FilterKindLessOrEqual:
		return true
	}
	return false
}

func (e FilterKind) String() string {
	return string(e)
}

func (e *FilterKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterKind", str)
	}
	return nil
}

func (e FilterKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderEndState string

const (
	// The order was successfully validated and added to the Mesh node. The order is now being watched and any changes to
	// the fillability will result in subsequent order events.
	OrderEndStateAdded OrderEndState = "ADDED"
	// The order was filled for a partial amount. The order is still fillable up to the fillableTakerAssetAmount.
	OrderEndStateFilled OrderEndState = "FILLED"
	// The order was fully filled and its remaining fillableTakerAssetAmount is 0. The order is no longer fillable.
	OrderEndStateFullyFilled OrderEndState = "FULLY_FILLED"
	// The order was cancelled and is no longer fillable.
	OrderEndStateCancelled OrderEndState = "CANCELLED"
	// The order expired and is no longer fillable.
	OrderEndStateExpired OrderEndState = "EXPIRED"
	// The order was previously expired, but due to a block re-org it is no longer considered expired (should be rare).
	OrderEndStateUnexpired OrderEndState = "UNEXPIRED"
	// The order has become unfunded and is no longer fillable. This can happen if the maker makes a transfer or changes their allowance.
	OrderEndStateUnfunded OrderEndState = "UNFUNDED"
	// The fillability of the order has increased. This can happen if a previously processed fill event gets reverted due to a block re-org,
	// or if a maker makes a transfer or changes their allowance.
	OrderEndStateFillabilityIncreased OrderEndState = "FILLABILITY_INCREASED"
	// The order is potentially still valid but was removed for a different reason (e.g.
	// the database is full or the peer that sent the order was misbehaving). The order will no longer be watched
	// and no further events for this order will be emitted. In some cases, the order may be re-added in the
	// future.
	OrderEndStateStoppedWatching OrderEndState = "STOPPED_WATCHING"
)

var AllOrderEndState = []OrderEndState{
	OrderEndStateAdded,
	OrderEndStateFilled,
	OrderEndStateFullyFilled,
	OrderEndStateCancelled,
	OrderEndStateExpired,
	OrderEndStateUnexpired,
	OrderEndStateUnfunded,
	OrderEndStateFillabilityIncreased,
	OrderEndStateStoppedWatching,
}

func (e OrderEndState) IsValid() bool {
	switch e {
	case OrderEndStateAdded, OrderEndStateFilled, OrderEndStateFullyFilled, OrderEndStateCancelled, OrderEndStateExpired, OrderEndStateUnexpired, OrderEndStateUnfunded, OrderEndStateFillabilityIncreased, OrderEndStateStoppedWatching:
		return true
	}
	return false
}

func (e OrderEndState) String() string {
	return string(e)
}

func (e *OrderEndState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderEndState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderEndState", str)
	}
	return nil
}

func (e OrderEndState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// An enum containing all the order fields for which filters and/or sorting is supported.
type OrderField string

const (
	OrderFieldHash                     OrderField = "hash"
	OrderFieldChainID                  OrderField = "chainId"
	OrderFieldExchangeAddress          OrderField = "exchangeAddress"
	OrderFieldMakerAddress             OrderField = "makerAddress"
	OrderFieldMakerAssetData           OrderField = "makerAssetData"
	OrderFieldMakerAssetAmount         OrderField = "makerAssetAmount"
	OrderFieldMakerFeeAssetData        OrderField = "makerFeeAssetData"
	OrderFieldMakerFee                 OrderField = "makerFee"
	OrderFieldTakerAddress             OrderField = "takerAddress"
	OrderFieldTakerAssetData           OrderField = "takerAssetData"
	OrderFieldTakerAssetAmount         OrderField = "takerAssetAmount"
	OrderFieldTakerFeeAssetData        OrderField = "takerFeeAssetData"
	OrderFieldTakerFee                 OrderField = "takerFee"
	OrderFieldSenderAddress            OrderField = "senderAddress"
	OrderFieldFeeRecipientAddress      OrderField = "feeRecipientAddress"
	OrderFieldExpirationTimeSeconds    OrderField = "expirationTimeSeconds"
	OrderFieldSalt                     OrderField = "salt"
	OrderFieldFillableTakerAssetAmount OrderField = "fillableTakerAssetAmount"
)

var AllOrderField = []OrderField{
	OrderFieldHash,
	OrderFieldChainID,
	OrderFieldExchangeAddress,
	OrderFieldMakerAddress,
	OrderFieldMakerAssetData,
	OrderFieldMakerAssetAmount,
	OrderFieldMakerFeeAssetData,
	OrderFieldMakerFee,
	OrderFieldTakerAddress,
	OrderFieldTakerAssetData,
	OrderFieldTakerAssetAmount,
	OrderFieldTakerFeeAssetData,
	OrderFieldTakerFee,
	OrderFieldSenderAddress,
	OrderFieldFeeRecipientAddress,
	OrderFieldExpirationTimeSeconds,
	OrderFieldSalt,
	OrderFieldFillableTakerAssetAmount,
}

func (e OrderField) IsValid() bool {
	switch e {
	case OrderFieldHash, OrderFieldChainID, OrderFieldExchangeAddress, OrderFieldMakerAddress, OrderFieldMakerAssetData, OrderFieldMakerAssetAmount, OrderFieldMakerFeeAssetData, OrderFieldMakerFee, OrderFieldTakerAddress, OrderFieldTakerAssetData, OrderFieldTakerAssetAmount, OrderFieldTakerFeeAssetData, OrderFieldTakerFee, OrderFieldSenderAddress, OrderFieldFeeRecipientAddress, OrderFieldExpirationTimeSeconds, OrderFieldSalt, OrderFieldFillableTakerAssetAmount:
		return true
	}
	return false
}

func (e OrderField) String() string {
	return string(e)
}

func (e *OrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderField", str)
	}
	return nil
}

func (e OrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// A set of all possible codes included in RejectedOrderResult.
type RejectedOrderCode string

const (
	RejectedOrderCodeEthRPCRequestFailed              RejectedOrderCode = "ETH_RPC_REQUEST_FAILED"
	RejectedOrderCodeOrderHasInvalidMakerAssetAmount  RejectedOrderCode = "ORDER_HAS_INVALID_MAKER_ASSET_AMOUNT"
	RejectedOrderCodeOrderHasInvalidTakerAssetAmount  RejectedOrderCode = "ORDER_HAS_INVALID_TAKER_ASSET_AMOUNT"
	RejectedOrderCodeOrderExpired                     RejectedOrderCode = "ORDER_EXPIRED"
	RejectedOrderCodeOrderFullyFilled                 RejectedOrderCode = "ORDER_FULLY_FILLED"
	RejectedOrderCodeOrderCancelled                   RejectedOrderCode = "ORDER_CANCELLED"
	RejectedOrderCodeOrderUnfunded                    RejectedOrderCode = "ORDER_UNFUNDED"
	RejectedOrderCodeOrderHasInvalidMakerAssetData    RejectedOrderCode = "ORDER_HAS_INVALID_MAKER_ASSET_DATA"
	RejectedOrderCodeOrderHasInvalidMakerFeeAssetData RejectedOrderCode = "ORDER_HAS_INVALID_MAKER_FEE_ASSET_DATA"
	RejectedOrderCodeOrderHasInvalidTakerAssetData    RejectedOrderCode = "ORDER_HAS_INVALID_TAKER_ASSET_DATA"
	RejectedOrderCodeOrderHasInvalidTakerFeeAssetData RejectedOrderCode = "ORDER_HAS_INVALID_TAKER_FEE_ASSET_DATA"
	RejectedOrderCodeOrderHasInvalidSignature         RejectedOrderCode = "ORDER_HAS_INVALID_SIGNATURE"
	RejectedOrderCodeOrderMaxExpirationExceeded       RejectedOrderCode = "ORDER_MAX_EXPIRATION_EXCEEDED"
	RejectedOrderCodeInternalError                    RejectedOrderCode = "INTERNAL_ERROR"
	RejectedOrderCodeMaxOrderSizeExceeded             RejectedOrderCode = "MAX_ORDER_SIZE_EXCEEDED"
	RejectedOrderCodeOrderAlreadyStoredAndUnfillable  RejectedOrderCode = "ORDER_ALREADY_STORED_AND_UNFILLABLE"
	RejectedOrderCodeOrderForIncorrectChain           RejectedOrderCode = "ORDER_FOR_INCORRECT_CHAIN"
	RejectedOrderCodeIncorrectExchangeAddress         RejectedOrderCode = "INCORRECT_EXCHANGE_ADDRESS"
	RejectedOrderCodeSenderAddressNotAllowed          RejectedOrderCode = "SENDER_ADDRESS_NOT_ALLOWED"
	RejectedOrderCodeDatabaseFullOfOrders             RejectedOrderCode = "DATABASE_FULL_OF_ORDERS"
)

var AllRejectedOrderCode = []RejectedOrderCode{
	RejectedOrderCodeEthRPCRequestFailed,
	RejectedOrderCodeOrderHasInvalidMakerAssetAmount,
	RejectedOrderCodeOrderHasInvalidTakerAssetAmount,
	RejectedOrderCodeOrderExpired,
	RejectedOrderCodeOrderFullyFilled,
	RejectedOrderCodeOrderCancelled,
	RejectedOrderCodeOrderUnfunded,
	RejectedOrderCodeOrderHasInvalidMakerAssetData,
	RejectedOrderCodeOrderHasInvalidMakerFeeAssetData,
	RejectedOrderCodeOrderHasInvalidTakerAssetData,
	RejectedOrderCodeOrderHasInvalidTakerFeeAssetData,
	RejectedOrderCodeOrderHasInvalidSignature,
	RejectedOrderCodeOrderMaxExpirationExceeded,
	RejectedOrderCodeInternalError,
	RejectedOrderCodeMaxOrderSizeExceeded,
	RejectedOrderCodeOrderAlreadyStoredAndUnfillable,
	RejectedOrderCodeOrderForIncorrectChain,
	RejectedOrderCodeIncorrectExchangeAddress,
	RejectedOrderCodeSenderAddressNotAllowed,
	RejectedOrderCodeDatabaseFullOfOrders,
}

func (e RejectedOrderCode) IsValid() bool {
	switch e {
	case RejectedOrderCodeEthRPCRequestFailed, RejectedOrderCodeOrderHasInvalidMakerAssetAmount, RejectedOrderCodeOrderHasInvalidTakerAssetAmount, RejectedOrderCodeOrderExpired, RejectedOrderCodeOrderFullyFilled, RejectedOrderCodeOrderCancelled, RejectedOrderCodeOrderUnfunded, RejectedOrderCodeOrderHasInvalidMakerAssetData, RejectedOrderCodeOrderHasInvalidMakerFeeAssetData, RejectedOrderCodeOrderHasInvalidTakerAssetData, RejectedOrderCodeOrderHasInvalidTakerFeeAssetData, RejectedOrderCodeOrderHasInvalidSignature, RejectedOrderCodeOrderMaxExpirationExceeded, RejectedOrderCodeInternalError, RejectedOrderCodeMaxOrderSizeExceeded, RejectedOrderCodeOrderAlreadyStoredAndUnfillable, RejectedOrderCodeOrderForIncorrectChain, RejectedOrderCodeIncorrectExchangeAddress, RejectedOrderCodeSenderAddressNotAllowed, RejectedOrderCodeDatabaseFullOfOrders:
		return true
	}
	return false
}

func (e RejectedOrderCode) String() string {
	return string(e)
}

func (e *RejectedOrderCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RejectedOrderCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RejectedOrderCode", str)
	}
	return nil
}

func (e RejectedOrderCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// The direction to sort in. Ascending means lowest to highest. Descending means highest to lowest.
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
