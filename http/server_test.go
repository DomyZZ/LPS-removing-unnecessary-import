package http

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rsksmart/liquidity-provider-server/http/testmocks"
	"github.com/rsksmart/liquidity-provider/providers"
	"github.com/rsksmart/liquidity-provider/types"
	"github.com/stretchr/testify/assert"
	http2 "github.com/stretchr/testify/http"
	"github.com/stretchr/testify/mock"
)

type LiquidityProviderMock struct {
	address string
}

func (lp LiquidityProviderMock) SignTx(common.Address, *gethTypes.Transaction) (*gethTypes.Transaction, error) {
	return nil, nil
}

func (lp LiquidityProviderMock) Address() string {
	return lp.address
}

func (lp LiquidityProviderMock) GetQuote(q types.Quote, gas uint64, gasPrice big.Int) *types.Quote {
	return nil
}

func (lp LiquidityProviderMock) SignHash(hash []byte) ([]byte, error) {
	return nil, nil
}

var providerMocks = []LiquidityProviderMock{
	{address: "123"},
	{address: "0x00d80aA033fb51F191563B08Dc035fA128e942C5"},
}

var testQuotes = []*types.Quote{
	{
		FedBTCAddr:         "mnxKdPFrYqLSUy2oP1eno8n5X8AwkcnPjk",
		LBCAddr:            "2ff74F841b95E000625b3A77fed03714874C4fEa",
		LPRSKAddr:          "0x00d80aA033fb51F191563B08Dc035fA128e942C5",
		BTCRefundAddr:      "mnxKdPFrYqLSUy2oP1eno8n5X8AwkcnPjk",
		RSKRefundAddr:      "0x5F3b836CA64DA03e613887B46f71D168FC8B5Bdf",
		LPBTCAddr:          "2NDjJznHgtH1rzq63eeFG3SiDi5wxE25FSz",
		CallFee:            *big.NewInt(250),
		PenaltyFee:         *big.NewInt(5000),
		ContractAddr:       "0x87136cf829edaF7c46Eb943063369a1C8D4f9085",
		Data:               "",
		GasLimit:           6000000,
		Nonce:              rand.Int(),
		Value:              *big.NewInt(250),
		AgreementTimestamp: 0,
		TimeForDeposit:     3600,
		CallTime:           3600,
		Confirmations:      10,
	},
}

func testGetProviderByAddress(t *testing.T) {
	var liquidityProviders []providers.LiquidityProvider
	for _, mock := range providerMocks {
		liquidityProviders = append(liquidityProviders, mock)
	}

	for _, tt := range liquidityProviders {
		result := getProviderByAddress(liquidityProviders, tt.Address())
		assert.EqualValues(t, tt.Address(), result.Address())
	}
}

func testAcceptQuoteComplete(t *testing.T) {
	for _, quote := range testQuotes {
		rsk := new(testmocks.RskMock)
		btc := new(testmocks.BtcMock)
		db := testmocks.NewDbMock(quote)

		srv := New(rsk, btc, db)
		for _, lp := range providerMocks {
			rsk.On("GetCollateral", lp.address).Times(1).Return(big.NewInt(10), big.NewInt(10))
			srv.AddProvider(lp)
		}
		w := http2.TestResponseWriter{}
		hash := "555c9cfba7638a40a71a17a34fef0c3e192c1fbf4b311ad6e2ae288e97794228"
		body := fmt.Sprintf("{\"quoteHash\":\"%v\"}", hash)

		btcRefAddr, lpBTCAddr, lbcAddr, err := decodeAddresses(quote.BTCRefundAddr, quote.LPBTCAddr, quote.LBCAddr)
		if err != nil {
			t.Errorf("couldn't decode addresses. error: %v", err)
		}
		req, err := http.NewRequest("POST", "acceptQuote", bytes.NewReader([]byte(body)))
		if err != nil {
			t.Errorf("couldn't instantiate request. error: %v", err)
		}
		hashBytes, err := hex.DecodeString(hash)
		if err != nil {
			t.Errorf("couldn't decode hash. error: %v", err)
		}

		db.On("GetQuote", hash).Times(1).Return(quote)
		btc.On("GetDerivedBitcoinAddress", btcRefAddr, lbcAddr, lpBTCAddr, hashBytes).Times(1).Return("")
		btc.On("AddAddressWatcher", "", time.Minute, mock.AnythingOfType("*http.BTCAddressWatcher")).Times(1).Return("")

		srv.acceptQuoteHandler(&w, req)
		db.AssertExpectations(t)
		btc.AssertExpectations(t)
		rsk.AssertExpectations(t)
	}
}

func TestLiquidityProviderServer(t *testing.T) {
	t.Run("get provider by address", testGetProviderByAddress)
	t.Run("accept quote", testAcceptQuoteComplete)
}
