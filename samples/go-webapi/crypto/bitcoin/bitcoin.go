package bitcoin

// #cgo CFLAGS: -I../../../../include
// #cgo LDFLAGS: -L../../../../build -L../../../../build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWBitcoinScript.h>
// #include <TrustWalletCore/TWAnySigner.h>
// #include <TrustWalletCore/TWMnemonic.h>
import "C"

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/ubozov/wallet-core/samples/go-webapi/types"
)

type transaction struct {
	UTXO struct {
		Hash     string `json:"hash"`
		Index    uint32 `json:"index"`
		Sequence uint32 `json:"sequence"`
		Amount   int64  `json:"amount"`
	} `json:"utxo"`
	ToAddress     string `json:"toAddress"`
	ChangeAddress string `json:"changeAddress`
	Fee           int64  `json:"byteFee`
	Amount        int64  `json:"amount`
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

	key := C.TWHDWalletGetKeyForCoin(wallet, C.TWCoinTypeBitcoin)
	keyData := C.TWPrivateKeyData(key)
	defer C.TWDataDelete(keyData)

	fmt.Println("<== bitcoin private key: ", types.TWDataHexString(keyData))

	// pubKey, err := hex.DecodeString("0288be7586c41a0498c1f931a0aaf08c15811ee2651a5fe0fa213167dcaba59ae8")
	// if err != nil {
	// 	return "", err
	// }
	// pubKeyData := types.TWDataCreateWithGoBytes(pubKey)
	// defer C.TWDataDelete(pubKeyData)
	// fmt.Println("==> bitcoin public key is valid: ", C.TWPublicKeyIsValid(pubKeyData, C.TWPublicKeyTypeSECP256k1))

	address := C.TWHDWalletGetAddressForCoin(wallet, C.TWCoinTypeBitcoin)
	defer C.TWStringDelete(address)
	fmt.Println("<== bitcoin address: ", types.TWStringGoString(address))

	script := C.TWBitcoinScriptLockScriptForAddress(address, C.TWCoinTypeBitcoin)
	scriptData := C.TWBitcoinScriptData(script)
	defer C.TWBitcoinScriptDelete(script)
	defer C.TWDataDelete(scriptData)
	fmt.Println("<== bitcoin address lock script: ", types.TWDataHexString(scriptData))

	utxoHash, err := hex.DecodeString(tx.UTXO.Hash)
	if err != nil {
		return "", err
	}

	utxo := UnspentTransaction{
		OutPoint: &OutPoint{
			Hash:     utxoHash,
			Index:    tx.UTXO.Index,
			Sequence: tx.UTXO.Sequence,
		},
		Amount: tx.UTXO.Amount,
		Script: types.TWDataGoBytes(scriptData),
	}

	input := SigningInput{
		HashType:      1, // TWBitcoinSigHashTypeAll
		Amount:        tx.Amount,
		ByteFee:       tx.Fee,
		ToAddress:     tx.ToAddress,
		ChangeAddress: tx.ChangeAddress,
		PrivateKey:    [][]byte{types.TWDataGoBytes(keyData)},
		Utxo:          []*UnspentTransaction{&utxo},
		CoinType:      0, // TWCoinTypeBitcoin
	}

	inputBytes, err := proto.Marshal(&input)
	if err != nil {
		return "", err
	}
	inputData := types.TWDataCreateWithGoBytes(inputBytes)
	defer C.TWDataDelete(inputData)

	outputData := C.TWAnySignerSign(inputData, C.TWCoinTypeBitcoin)
	defer C.TWDataDelete(outputData)

	var output SigningOutput
	if err := proto.Unmarshal(types.TWDataGoBytes(outputData), &output); err != nil {
		return "", err
	}
	fmt.Println("<== bitcoin signed tx: ", hex.EncodeToString(output.Encoded))

	return hex.EncodeToString(output.Encoded), nil
}
