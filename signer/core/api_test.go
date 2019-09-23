// Copyright 2018 The go-ethereum Authors
<<<<<<< HEAD
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.
//
package core
=======
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core_test
>>>>>>> upstream/master

import (
	"bytes"
	"context"
<<<<<<< HEAD
	"errors"
=======
>>>>>>> upstream/master
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

<<<<<<< HEAD
	"github.com/Onther-Tech/go-ethereum/accounts/keystore"
	"github.com/Onther-Tech/go-ethereum/cmd/utils"
	"github.com/Onther-Tech/go-ethereum/common"
	"github.com/Onther-Tech/go-ethereum/common/hexutil"
	"github.com/Onther-Tech/go-ethereum/core/types"
	"github.com/Onther-Tech/go-ethereum/internal/ethapi"
	"github.com/Onther-Tech/go-ethereum/rlp"
)

//Used for testing
type HeadlessUI struct {
	controller chan string
}

func (ui *HeadlessUI) OnInputRequired(info UserInputRequest) (UserInputResponse, error) {
	return UserInputResponse{}, errors.New("not implemented")
}

func (ui *HeadlessUI) OnSignerStartup(info StartupInfo) {
}

func (ui *HeadlessUI) OnApprovedTx(tx ethapi.SignTransactionResult) {
	fmt.Printf("OnApproved()\n")
}

func (ui *HeadlessUI) ApproveTx(request *SignTxRequest) (SignTxResponse, error) {

	switch <-ui.controller {
	case "Y":
		return SignTxResponse{request.Transaction, true, <-ui.controller}, nil
	case "M": //Modify
		old := big.Int(request.Transaction.Value)
		newVal := big.NewInt(0).Add(&old, big.NewInt(1))
		request.Transaction.Value = hexutil.Big(*newVal)
		return SignTxResponse{request.Transaction, true, <-ui.controller}, nil
	default:
		return SignTxResponse{request.Transaction, false, ""}, nil
	}
}

func (ui *HeadlessUI) ApproveSignData(request *SignDataRequest) (SignDataResponse, error) {
	if "Y" == <-ui.controller {
		return SignDataResponse{true, <-ui.controller}, nil
	}
	return SignDataResponse{false, ""}, nil
}

func (ui *HeadlessUI) ApproveExport(request *ExportRequest) (ExportResponse, error) {
	return ExportResponse{<-ui.controller == "Y"}, nil

}

func (ui *HeadlessUI) ApproveImport(request *ImportRequest) (ImportResponse, error) {
	if "Y" == <-ui.controller {
		return ImportResponse{true, <-ui.controller, <-ui.controller}, nil
	}
	return ImportResponse{false, "", ""}, nil
}

func (ui *HeadlessUI) ApproveListing(request *ListRequest) (ListResponse, error) {
	switch <-ui.controller {
	case "A":
		return ListResponse{request.Accounts}, nil
	case "1":
		l := make([]Account, 1)
		l[0] = request.Accounts[1]
		return ListResponse{l}, nil
	default:
		return ListResponse{nil}, nil
	}
}

func (ui *HeadlessUI) ApproveNewAccount(request *NewAccountRequest) (NewAccountResponse, error) {
	if "Y" == <-ui.controller {
		return NewAccountResponse{true, <-ui.controller}, nil
	}
	return NewAccountResponse{false, ""}, nil
}

func (ui *HeadlessUI) ShowError(message string) {
=======
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/ethereum/go-ethereum/signer/fourbyte"
	"github.com/ethereum/go-ethereum/signer/storage"
)

//Used for testing
type headlessUi struct {
	approveCh chan string // to send approve/deny
	inputCh   chan string // to send password
}

func (ui *headlessUi) OnInputRequired(info core.UserInputRequest) (core.UserInputResponse, error) {
	input := <-ui.inputCh
	return core.UserInputResponse{Text: input}, nil
}

func (ui *headlessUi) OnSignerStartup(info core.StartupInfo)        {}
func (ui *headlessUi) RegisterUIServer(api *core.UIServerAPI)       {}
func (ui *headlessUi) OnApprovedTx(tx ethapi.SignTransactionResult) {}

func (ui *headlessUi) ApproveTx(request *core.SignTxRequest) (core.SignTxResponse, error) {

	switch <-ui.approveCh {
	case "Y":
		return core.SignTxResponse{request.Transaction, true}, nil
	case "M": // modify
		// The headless UI always modifies the transaction
		old := big.Int(request.Transaction.Value)
		newVal := big.NewInt(0).Add(&old, big.NewInt(1))
		request.Transaction.Value = hexutil.Big(*newVal)
		return core.SignTxResponse{request.Transaction, true}, nil
	default:
		return core.SignTxResponse{request.Transaction, false}, nil
	}
}

func (ui *headlessUi) ApproveSignData(request *core.SignDataRequest) (core.SignDataResponse, error) {
	approved := "Y" == <-ui.approveCh
	return core.SignDataResponse{approved}, nil
}

func (ui *headlessUi) ApproveListing(request *core.ListRequest) (core.ListResponse, error) {
	approval := <-ui.approveCh
	//fmt.Printf("approval %s\n", approval)
	switch approval {
	case "A":
		return core.ListResponse{request.Accounts}, nil
	case "1":
		l := make([]accounts.Account, 1)
		l[0] = request.Accounts[1]
		return core.ListResponse{l}, nil
	default:
		return core.ListResponse{nil}, nil
	}
}

func (ui *headlessUi) ApproveNewAccount(request *core.NewAccountRequest) (core.NewAccountResponse, error) {
	if "Y" == <-ui.approveCh {
		return core.NewAccountResponse{true}, nil
	}
	return core.NewAccountResponse{false}, nil
}

func (ui *headlessUi) ShowError(message string) {
>>>>>>> upstream/master
	//stdout is used by communication
	fmt.Fprintln(os.Stderr, message)
}

<<<<<<< HEAD
func (ui *HeadlessUI) ShowInfo(message string) {
=======
func (ui *headlessUi) ShowInfo(message string) {
>>>>>>> upstream/master
	//stdout is used by communication
	fmt.Fprintln(os.Stderr, message)
}

func tmpDirName(t *testing.T) string {
	d, err := ioutil.TempDir("", "eth-keystore-test")
	if err != nil {
		t.Fatal(err)
	}
	d, err = filepath.EvalSymlinks(d)
	if err != nil {
		t.Fatal(err)
	}
	return d
}

<<<<<<< HEAD
func setup(t *testing.T) (*SignerAPI, chan string) {

	controller := make(chan string, 20)

	db, err := NewAbiDBFromFile("../../cmd/clef/4byte.json")
	if err != nil {
		utils.Fatalf(err.Error())
	}
	var (
		ui  = &HeadlessUI{controller}
		api = NewSignerAPI(
			1,
			tmpDirName(t),
			true,
			ui,
			db,
			true, true)
	)
	return api, controller
}
func createAccount(control chan string, api *SignerAPI, t *testing.T) {

	control <- "Y"
	control <- "a_long_password"
=======
func setup(t *testing.T) (*core.SignerAPI, *headlessUi) {
	db, err := fourbyte.New()
	if err != nil {
		t.Fatal(err.Error())
	}
	ui := &headlessUi{make(chan string, 20), make(chan string, 20)}
	am := core.StartClefAccountManager(tmpDirName(t), true, true, "")
	api := core.NewSignerAPI(am, 1337, true, ui, db, true, &storage.NoStorage{})
	return api, ui

}
func createAccount(ui *headlessUi, api *core.SignerAPI, t *testing.T) {
	ui.approveCh <- "Y"
	ui.inputCh <- "a_long_password"
>>>>>>> upstream/master
	_, err := api.New(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	// Some time to allow changes to propagate
	time.Sleep(250 * time.Millisecond)
}

<<<<<<< HEAD
func failCreateAccountWithPassword(control chan string, api *SignerAPI, password string, t *testing.T) {

	control <- "Y"
	control <- password
	control <- "Y"
	control <- password
	control <- "Y"
	control <- password

	acc, err := api.New(context.Background())
	if err == nil {
		t.Fatal("Should have returned an error")
	}
	if acc.Address != (common.Address{}) {
=======
func failCreateAccountWithPassword(ui *headlessUi, api *core.SignerAPI, password string, t *testing.T) {

	ui.approveCh <- "Y"
	// We will be asked three times to provide a suitable password
	ui.inputCh <- password
	ui.inputCh <- password
	ui.inputCh <- password

	addr, err := api.New(context.Background())
	if err == nil {
		t.Fatal("Should have returned an error")
	}
	if addr != (common.Address{}) {
>>>>>>> upstream/master
		t.Fatal("Empty address should be returned")
	}
}

<<<<<<< HEAD
func failCreateAccount(control chan string, api *SignerAPI, t *testing.T) {
	control <- "N"
	acc, err := api.New(context.Background())
	if err != ErrRequestDenied {
		t.Fatal(err)
	}
	if acc.Address != (common.Address{}) {
=======
func failCreateAccount(ui *headlessUi, api *core.SignerAPI, t *testing.T) {
	ui.approveCh <- "N"
	addr, err := api.New(context.Background())
	if err != core.ErrRequestDenied {
		t.Fatal(err)
	}
	if addr != (common.Address{}) {
>>>>>>> upstream/master
		t.Fatal("Empty address should be returned")
	}
}

<<<<<<< HEAD
func list(control chan string, api *SignerAPI, t *testing.T) []common.Address {
	control <- "A"
	list, err := api.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	return list
=======
func list(ui *headlessUi, api *core.SignerAPI, t *testing.T) ([]common.Address, error) {
	ui.approveCh <- "A"
	return api.List(context.Background())

>>>>>>> upstream/master
}

func TestNewAcc(t *testing.T) {
	api, control := setup(t)
	verifyNum := func(num int) {
<<<<<<< HEAD
		if list := list(control, api, t); len(list) != num {
=======
		list, err := list(control, api, t)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if len(list) != num {
>>>>>>> upstream/master
			t.Errorf("Expected %d accounts, got %d", num, len(list))
		}
	}
	// Testing create and create-deny
	createAccount(control, api, t)
	createAccount(control, api, t)
	failCreateAccount(control, api, t)
	failCreateAccount(control, api, t)
	createAccount(control, api, t)
	failCreateAccount(control, api, t)
	createAccount(control, api, t)
	failCreateAccount(control, api, t)
<<<<<<< HEAD

=======
>>>>>>> upstream/master
	verifyNum(4)

	// Fail to create this, due to bad password
	failCreateAccountWithPassword(control, api, "short", t)
	failCreateAccountWithPassword(control, api, "longerbutbad\rfoo", t)
<<<<<<< HEAD

=======
>>>>>>> upstream/master
	verifyNum(4)

	// Testing listing:
	// Listing one Account
<<<<<<< HEAD
	control <- "1"
=======
	control.approveCh <- "1"
>>>>>>> upstream/master
	list, err := api.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 {
		t.Fatalf("List should only show one Account")
	}
	// Listing denied
<<<<<<< HEAD
	control <- "Nope"
=======
	control.approveCh <- "Nope"
>>>>>>> upstream/master
	list, err = api.List(context.Background())
	if len(list) != 0 {
		t.Fatalf("List should be empty")
	}
<<<<<<< HEAD
	if err != ErrRequestDenied {
=======
	if err != core.ErrRequestDenied {
>>>>>>> upstream/master
		t.Fatal("Expected deny")
	}
}

<<<<<<< HEAD
func TestSignData(t *testing.T) {
	api, control := setup(t)
	//Create two accounts
	createAccount(control, api, t)
	createAccount(control, api, t)
	control <- "1"
	list, err := api.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := common.NewMixedcaseAddress(list[0])

	control <- "Y"
	control <- "wrongpassword"
	h, err := api.Sign(context.Background(), a, []byte("EHLO world"))
	if h != nil {
		t.Errorf("Expected nil-data, got %x", h)
	}
	if err != keystore.ErrDecrypt {
		t.Errorf("Expected ErrLocked! %v", err)
	}
	control <- "No way"
	h, err = api.Sign(context.Background(), a, []byte("EHLO world"))
	if h != nil {
		t.Errorf("Expected nil-data, got %x", h)
	}
	if err != ErrRequestDenied {
		t.Errorf("Expected ErrRequestDenied! %v", err)
	}
	control <- "Y"
	control <- "a_long_password"
	h, err = api.Sign(context.Background(), a, []byte("EHLO world"))
	if err != nil {
		t.Fatal(err)
	}
	if h == nil || len(h) != 65 {
		t.Errorf("Expected 65 byte signature (got %d bytes)", len(h))
	}
}
func mkTestTx(from common.MixedcaseAddress) SendTxArgs {
=======
func mkTestTx(from common.MixedcaseAddress) core.SendTxArgs {
>>>>>>> upstream/master
	to := common.NewMixedcaseAddress(common.HexToAddress("0x1337"))
	gas := hexutil.Uint64(21000)
	gasPrice := (hexutil.Big)(*big.NewInt(2000000000))
	value := (hexutil.Big)(*big.NewInt(1e18))
	nonce := (hexutil.Uint64)(0)
	data := hexutil.Bytes(common.Hex2Bytes("01020304050607080a"))
<<<<<<< HEAD
	tx := SendTxArgs{
=======
	tx := core.SendTxArgs{
>>>>>>> upstream/master
		From:     from,
		To:       &to,
		Gas:      gas,
		GasPrice: gasPrice,
		Value:    value,
		Data:     &data,
		Nonce:    nonce}
	return tx
}

func TestSignTx(t *testing.T) {
	var (
		list      []common.Address
		res, res2 *ethapi.SignTransactionResult
		err       error
	)

	api, control := setup(t)
	createAccount(control, api, t)
<<<<<<< HEAD
	control <- "A"
=======
	control.approveCh <- "A"
>>>>>>> upstream/master
	list, err = api.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	a := common.NewMixedcaseAddress(list[0])

	methodSig := "test(uint)"
	tx := mkTestTx(a)

<<<<<<< HEAD
	control <- "Y"
	control <- "wrongpassword"
=======
	control.approveCh <- "Y"
	control.inputCh <- "wrongpassword"
>>>>>>> upstream/master
	res, err = api.SignTransaction(context.Background(), tx, &methodSig)
	if res != nil {
		t.Errorf("Expected nil-response, got %v", res)
	}
	if err != keystore.ErrDecrypt {
		t.Errorf("Expected ErrLocked! %v", err)
	}
<<<<<<< HEAD
	control <- "No way"
=======
	control.approveCh <- "No way"
>>>>>>> upstream/master
	res, err = api.SignTransaction(context.Background(), tx, &methodSig)
	if res != nil {
		t.Errorf("Expected nil-response, got %v", res)
	}
<<<<<<< HEAD
	if err != ErrRequestDenied {
		t.Errorf("Expected ErrRequestDenied! %v", err)
	}
	control <- "Y"
	control <- "a_long_password"
=======
	if err != core.ErrRequestDenied {
		t.Errorf("Expected ErrRequestDenied! %v", err)
	}
	// Sign with correct password
	control.approveCh <- "Y"
	control.inputCh <- "a_long_password"
>>>>>>> upstream/master
	res, err = api.SignTransaction(context.Background(), tx, &methodSig)

	if err != nil {
		t.Fatal(err)
	}
	parsedTx := &types.Transaction{}
	rlp.Decode(bytes.NewReader(res.Raw), parsedTx)

	//The tx should NOT be modified by the UI
	if parsedTx.Value().Cmp(tx.Value.ToInt()) != 0 {
		t.Errorf("Expected value to be unchanged, expected %v got %v", tx.Value, parsedTx.Value())
	}
<<<<<<< HEAD
	control <- "Y"
	control <- "a_long_password"
=======
	control.approveCh <- "Y"
	control.inputCh <- "a_long_password"
>>>>>>> upstream/master

	res2, err = api.SignTransaction(context.Background(), tx, &methodSig)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(res.Raw, res2.Raw) {
		t.Error("Expected tx to be unmodified by UI")
	}

	//The tx is modified by the UI
<<<<<<< HEAD
	control <- "M"
	control <- "a_long_password"
=======
	control.approveCh <- "M"
	control.inputCh <- "a_long_password"
>>>>>>> upstream/master

	res2, err = api.SignTransaction(context.Background(), tx, &methodSig)
	if err != nil {
		t.Fatal(err)
	}
	parsedTx2 := &types.Transaction{}
	rlp.Decode(bytes.NewReader(res.Raw), parsedTx2)

	//The tx should be modified by the UI
	if parsedTx2.Value().Cmp(tx.Value.ToInt()) != 0 {
		t.Errorf("Expected value to be unchanged, got %v", parsedTx.Value())
	}
	if bytes.Equal(res.Raw, res2.Raw) {
		t.Error("Expected tx to be modified by UI")
	}

}
<<<<<<< HEAD

/*
func TestAsyncronousResponses(t *testing.T){

	//Set up one account
	api, control := setup(t)
	createAccount(control, api, t)

	// Two transactions, the second one with larger value than the first
	tx1 := mkTestTx()
	newVal := big.NewInt(0).Add((*big.Int) (tx1.Value), big.NewInt(1))
	tx2 := mkTestTx()
	tx2.Value = (*hexutil.Big)(newVal)

	control <- "W" //wait
	control <- "Y" //
	control <- "a_long_password"
	control <- "Y" //
	control <- "a_long_password"

	var err error

	h1, err := api.SignTransaction(context.Background(), common.HexToAddress("1111"), tx1, nil)
	h2, err := api.SignTransaction(context.Background(), common.HexToAddress("2222"), tx2, nil)


	}
*/
=======
>>>>>>> upstream/master
