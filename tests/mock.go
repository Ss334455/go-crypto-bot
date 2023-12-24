package tests

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/open-soft/go-crypto-bot/exchange_context/model"
)

type ExchangeRepositoryMock struct {
	mock.Mock
}

func (m *ExchangeRepositoryMock) CreateSwapPair(swapPair model.SwapPair) (*int64, error) {
	args := m.Called(swapPair)
	id := int64(args.Int(1))
	return &id, args.Error(0)
}
func (m *ExchangeRepositoryMock) UpdateSwapPair(swapPair model.SwapPair) error {
	args := m.Called(swapPair)
	return args.Error(0)
}
func (m *ExchangeRepositoryMock) GetSwapPairs() []model.SwapPair {
	args := m.Called()
	return args.Get(0).([]model.SwapPair)
}
func (m *ExchangeRepositoryMock) GetSwapPairsByBaseAsset(baseAsset string) []model.SwapPair {
	args := m.Called(baseAsset)
	return args.Get(0).([]model.SwapPair)
}
func (m *ExchangeRepositoryMock) GetSwapPairsByQuoteAsset(quoteAsset string) []model.SwapPair {
	args := m.Called(quoteAsset)
	return args.Get(0).([]model.SwapPair)
}
func (m *ExchangeRepositoryMock) GetSwapPair(symbol string) (model.SwapPair, error) {
	args := m.Called(symbol)
	return args.Get(0).(model.SwapPair), args.Error(0)
}

type SwapRepositoryMock struct {
	mock.Mock
	savedChain model.SwapChainEntity
	swapAction model.SwapAction
}

func (m *SwapRepositoryMock) GetSwapChain(hash string) (model.SwapChainEntity, error) {
	args := m.Called(hash)
	return args.Get(0).(model.SwapChainEntity), args.Error(1)
}
func (m *SwapRepositoryMock) CreateSwapChain(swapChain model.SwapChainEntity) (*int64, error) {
	m.savedChain = swapChain
	args := m.Called(swapChain)
	id := int64(args.Int(0))
	return &id, args.Error(1)
}
func (m *SwapRepositoryMock) UpdateSwapChain(swapChain model.SwapChainEntity) error {
	args := m.Called(swapChain)
	return args.Error(0)
}
func (m *SwapRepositoryMock) SaveSwapChainCache(asset string, entity model.SwapChainEntity) {
	m.Called(asset, entity)
}
func (m *SwapRepositoryMock) GetSwapPairBySymbol(symbol string) (model.SwapPair, error) {
	args := m.Called(symbol)
	return args.Get(0).(model.SwapPair), args.Error(1)
}
func (s *SwapRepositoryMock) GetActiveSwapAction(order model.Order) (model.SwapAction, error) {
	args := s.Called(order)
	return args.Get(0).(model.SwapAction), args.Error(1)
}
func (s *SwapRepositoryMock) UpdateSwapAction(action model.SwapAction) error {
	s.swapAction = action
	args := s.Called(action)
	return args.Error(0)
}
func (s *SwapRepositoryMock) GetSwapChainById(id int64) (model.SwapChainEntity, error) {
	args := s.Called(id)
	return args.Get(0).(model.SwapChainEntity), args.Error(1)
}

type OrderUpdaterMock struct {
	mock.Mock
	order model.Order
}

func (s *OrderUpdaterMock) Update(order model.Order) error {
	s.order = order
	args := s.Called(order)
	return args.Error(0)
}

type BalanceServiceMock struct {
	mock.Mock
}

func (b *BalanceServiceMock) GetAssetBalance(asset string, cache bool) (float64, error) {
	args := b.Called(asset, cache)
	return args.Get(0).(float64), args.Error(1)
}
func (b *BalanceServiceMock) InvalidateBalanceCache(asset string) {
	_ = b.Called(asset)
}

type ExchangeOrderAPIMock struct {
	mock.Mock
}

func (b *ExchangeOrderAPIMock) LimitOrder(symbol string, quantity float64, price float64, operation string, timeInForce string) (model.BinanceOrder, error) {
	args := b.Called(symbol, quantity, price, operation, timeInForce)
	return args.Get(0).(model.BinanceOrder), args.Error(1)
}
func (b *ExchangeOrderAPIMock) QueryOrder(symbol string, orderId int64) (model.BinanceOrder, error) {
	args := b.Called(symbol, orderId)
	return args.Get(0).(model.BinanceOrder), args.Error(1)
}
func (b *ExchangeOrderAPIMock) CancelOrder(symbol string, orderId int64) (model.BinanceOrder, error) {
	args := b.Called(symbol, orderId)
	return args.Get(0).(model.BinanceOrder), args.Error(1)
}

type TimeServiceMock struct {
	mock.Mock
}

func (t *TimeServiceMock) WaitSeconds(seconds int64) {
	_ = t.Called(seconds)
}
func (t *TimeServiceMock) GetNowDiffMinutes(unixTime int64) float64 {
	args := t.Called(unixTime)
	return args.Get(0).(float64)
}

type ExchangePriceStorageMock struct {
	mock.Mock
}

func (e *ExchangePriceStorageMock) GetLastKLine(symbol string) *model.KLine {
	args := e.Called(symbol)
	kLine := args.Get(0).(*model.KLine)
	return kLine
}
func (e *ExchangePriceStorageMock) GetPeriodMinPrice(symbol string, period int64) float64 {
	args := e.Called(symbol, period)
	return args.Get(0).(float64)
}
func (e *ExchangePriceStorageMock) GetDepth(symbol string) model.Depth {
	args := e.Called(symbol)
	return args.Get(0).(model.Depth)
}
func (e *ExchangePriceStorageMock) SetDepth(depth model.Depth) {
	_ = e.Called(depth)
}

type OrderCachedReaderMock struct {
	mock.Mock
}

func (o *OrderCachedReaderMock) GetOpenedOrderCached(symbol string, operation string) (model.Order, error) {
	args := o.Called(symbol, operation)
	return args.Get(0).(model.Order), args.Error(1)
}

type FrameServiceMock struct {
	mock.Mock
}

func (f *FrameServiceMock) GetFrame(symbol string, interval string, limit int64) model.Frame {
	args := f.Called(symbol, interval, limit)
	return args.Get(0).(model.Frame)
}

type ExchangePriceAPIMock struct {
	mock.Mock
}

func (e *ExchangePriceAPIMock) GetDepth(symbol string) (model.OrderBook, error) {
	args := e.Called(symbol)
	return args.Get(0).(model.OrderBook), args.Error(1)
}
func (e *ExchangePriceAPIMock) GetKLines(symbol string, interval string, limit int64) []model.KLineHistory {
	args := e.Called(symbol, interval, limit)
	return args.Get(0).([]model.KLineHistory)
}
func (e *ExchangePriceAPIMock) GetKLinesCached(symbol string, interval string, limit int64) []model.KLineHistory {
	args := e.Called(symbol, interval, limit)
	return args.Get(0).([]model.KLineHistory)
}