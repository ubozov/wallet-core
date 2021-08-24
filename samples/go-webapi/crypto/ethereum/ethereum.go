package ethereum

// #cgo CFLAGS: -I../../../../include
// #cgo LDFLAGS: -L../../../../build -L../../../../build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWAnySigner.h>
// #include <TrustWalletCore/TWMnemonic.h>
import "C"

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"google.golang.org/protobuf/proto"

	"github.com/ubozov/wallet-core/samples/go-webapi/types"
)

type transaction struct {
	Nonce     int    `json:"nonce"`
	GasPrice  int    `json:"gasPrice"`
	GasLimit  int    `json:"gasLimit"`
	ToAddress string `json:"toAddress"`
	Amount    int    `json:"amount`
}

func Sign(seed string, in interface{}) (string, error) {
	jsonString, err := json.Marshal(in)
	if err != nil {
		return "", err
	}

	tx := transaction{}
	if err := json.Unmarshal(jsonString, &tx); err != nil {
		return "", err
	}

	fmt.Println("==> calling wallet core from go")
	str := types.TWStringCreateWithGoString(seed)
	emtpy := types.TWStringCreateWithGoString("")
	defer C.TWStringDelete(str)
	defer C.TWStringDelete(emtpy)

	fmt.Println("==> mnemonic is valid: ", C.TWMnemonicIsValid(str))

	wallet := C.TWHDWalletCreateWithMnemonic(str, emtpy)
	defer C.TWHDWalletDelete(wallet)

	key := C.TWHDWalletGetKeyForCoin(wallet, C.TWCoinTypeEthereum)
	keyData := C.TWPrivateKeyData(key)
	defer C.TWDataDelete(keyData)

	fmt.Println("<== ethereum private key: ", types.TWDataHexString(keyData))

	address := C.TWHDWalletGetAddressForCoin(wallet, C.TWCoinTypeEthereum)
	defer C.TWStringDelete(address)
	fmt.Println("<== ethereum address: ", types.TWStringGoString(address))

	input := SigningInput{
		Nonce:      []byte(strconv.Itoa(tx.Nonce)),
		GasPrice:   []byte(strconv.Itoa(tx.GasPrice)),
		GasLimit:   []byte(strconv.Itoa(tx.GasLimit)),
		ToAddress:  tx.ToAddress,
		PrivateKey: types.TWDataGoBytes(keyData),
		Transaction: &Transaction{
			TransactionOneof: &Transaction_Transfer_{
				Transfer: &Transaction_Transfer{
					Amount: []byte(strconv.Itoa(tx.Amount)),
				},
			},
		},
	}

	inputBytes, err := proto.Marshal(&input)
	if err != nil {
		return "", err
	}
	inputData := types.TWDataCreateWithGoBytes(inputBytes)
	defer C.TWDataDelete(inputData)

	outputData := C.TWAnySignerSign(inputData, C.TWCoinTypeEthereum)
	defer C.TWDataDelete(outputData)

	var output SigningOutput
	if err := proto.Unmarshal(types.TWDataGoBytes(outputData), &output); err != nil {
		return "", err
	}
	fmt.Println("<== ethereum signed tx: ", hex.EncodeToString(output.Encoded))

	return hex.EncodeToString(output.Encoded), nil
}