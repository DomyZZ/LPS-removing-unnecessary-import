package testmocks

import (
	"context"
	"math/big"
	"time"

	"github.com/rsksmart/liquidity-provider-server/connectors"
	"github.com/rsksmart/liquidity-provider-server/pegout"

	"github.com/btcsuite/btcd/chaincfg"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rsksmart/liquidity-provider-server/connectors/bindings"
	"github.com/rsksmart/liquidity-provider/types"
	"github.com/stretchr/testify/mock"
)

type RskMock struct {
	mock.Mock
	QuoteHash string
}

func (m *RskMock) GetMinimumLockTxValue() (*big.Int, error) {
	args := m.Called()
	return args.Get(0).(*big.Int), args.Error(1)
}

func (m *RskMock) GetLbcBalance(addr string) (*big.Int, error) {
	args := m.Called(addr)
	return args.Get(0).(*big.Int), args.Error(1)
}

func (m *RskMock) GetAvailableLiquidity(addr string) (*big.Int, error) {
	args := m.Called(addr)
	return args.Get(0).(*big.Int), args.Error(1)
}

func (m *RskMock) GetCollateral(addr string) (*big.Int, *big.Int, error) {
	m.Called(addr)
	return big.NewInt(10), big.NewInt(10), nil
}

func (m *RskMock) RegisterProvider(opts *bind.TransactOpts) error {
	m.Called(opts)
	return nil
}

func (m *RskMock) AddCollateral(opts *bind.TransactOpts) error {
	m.Called(opts)
	return nil
}

func (m *RskMock) GetRequiredBridgeConfirmations() int64 {
	m.Called()
	return 0
}

func (m *RskMock) GetChainId() (*big.Int, error) {
	m.Called()
	return big.NewInt(0), nil
}

func (m *RskMock) ParseQuote(q *types.Quote) (bindings.LiquidityBridgeContractQuote, error) {
	m.Called(q)
	return bindings.LiquidityBridgeContractQuote{}, nil
}

func (m *RskMock) RegisterPegIn(opt *bind.TransactOpts, q bindings.LiquidityBridgeContractQuote, signature []byte, tx []byte, pmt []byte, height *big.Int) (*gethTypes.Transaction, error) {
	m.Called(opt, q, signature, tx, pmt, height)
	return nil, nil
}

func (m *RskMock) RegisterPegInWithoutTx(q bindings.LiquidityBridgeContractQuote, signature []byte, tx []byte, pmt []byte, height *big.Int) error {
	m.Called(q, signature, tx, pmt, height)
	return nil
}

func (m *RskMock) CallForUser(opt *bind.TransactOpts, q bindings.LiquidityBridgeContractQuote) (*gethTypes.Transaction, error) {
	m.Called(opt, q)
	return nil, nil
}

func (m *RskMock) Connect(endpoint string, chainId *big.Int) error {
	m.Called(endpoint, chainId)
	return nil
}

func (m *RskMock) CheckConnection() error {
	args := m.Called()
	return args.Error(0)
}

func (m *RskMock) Close() {
	m.Called()
}

func (m *RskMock) EstimateGas(addr string, value *big.Int, data []byte) (uint64, error) {
	m.Called(addr, value, data)
	return 10000, nil
}

func (m *RskMock) GasPrice() (*big.Int, error) {
	m.Called()
	return big.NewInt(100000), nil
}
func (m *RskMock) HashQuote(q *types.Quote) (string, error) {
	m.Called(q)
	return "", nil
}
func (m *RskMock) GetFedSize() (int, error) {
	args := m.Called()
	return args.Int(0), nil
}
func (m *RskMock) GetFedThreshold() (int, error) {
	args := m.Called()
	return args.Int(0), nil
}

func (m *RskMock) GetFedPublicKey(index int) (string, error) {
	args := m.Called(index)
	return args.String(), nil
}
func (m *RskMock) GetFedAddress() (string, error) {
	args := m.Called()
	return args.String(), nil
}
func (m *RskMock) GetActiveFederationCreationBlockHeight() (int, error) {
	args := m.Called()
	return args.Int(0), nil
}

func (m *RskMock) GetLBCAddress() string {
	args := m.Called()
	return args.String()
}

func (m *RskMock) GetTxStatus(ctx context.Context, tx *gethTypes.Transaction) (bool, error) {
	m.Called(ctx, tx)
	return false, nil
}

func (m *RskMock) FetchFederationInfo() (*connectors.FedInfo, error) {
	args := m.Called()
	return args.Get(0).(*connectors.FedInfo), args.Error(1)
}

func (m *RskMock) AddQuoteToWatch(hash string, interval time.Duration, exp time.Time, w connectors.QuotePegOutWatcher, cb func(w connectors.QuotePegOutWatcher)) error {
	return nil
}

func (m *RskMock) HashPegOutQuote(q *pegout.Quote) (string, error) {
	return m.QuoteHash, nil
}

func (m *RskMock) GetProviders() ([]bindings.LiquidityBridgeContractProvider, error) {
	m.Called()

	return nil, nil
}

func (m *RskMock) GetActivePowpegRedeemScript() ([]byte, error) {
	return nil, nil
}

func (b *RskMock) GetDerivedBitcoinAddress(fedInfo *connectors.FedInfo, btcParams chaincfg.Params, userBtcRefundAddr []byte, lbcAddress []byte, lpBtcAddress []byte, derivationArgumentsHash []byte) (string, error) {
	b.Called(fedInfo, nil, userBtcRefundAddr, lbcAddress, lpBtcAddress, derivationArgumentsHash)
	return "", nil
}
