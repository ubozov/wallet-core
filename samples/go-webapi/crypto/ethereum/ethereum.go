package ethereum

// #cgo CFLAGS: -I../../../../include
// #cgo LDFLAGS: -L../../../../build -L../../../../build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWEthereumScript.h>
// #include <TrustWalletCore/TWAnySigner.h>
// #include <TrustWalletCore/TWMnemonic.h>
import "C"

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"

	"go-webapi/crypto/proto/ethereum/pb"
	"go-webapi/types"
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

	key := C.TWHDWalletGetKeyForCoin(wallet, C.TWCoinTypeEthereum)
	keyData := C.TWPrivateKeyData(key)
	defer C.TWDataDelete(keyData)

	fmt.Println("<== ethereum private key: ", types.TWDataHexString(keyData))

	address := C.TWHDWalletGetAddressForCoin(wallet, C.TWCoinTypeEthereum)
	defer C.TWStringDelete(address)
	fmt.Println("<== ethereum address: ", types.TWStringGoString(address))

	script := C.TWEthereumScriptLockScriptForAddress(address, C.TWCoinTypeEthereum)
	scriptData := C.TWEthereumScriptData(script)
	defer C.TWEthereumScriptDelete(script)
	defer C.TWDataDelete(scriptData)
	fmt.Println("<== ethereum address lock script: ", types.TWDataHexString(scriptData))

	input := pb.SigningInput{
		Amount:        tx.Amount,
		ByteFee:       tx.Fee,
		ToAddress:     tx.ToAddress,
		ChangeAddress: tx.ChangeAddress,
		PrivateKey:    [][]byte{types.TWDataGoBytes(keyData)},
		CoinType:      60,
	}

	inputBytes, err := proto.Marshal(&input)
	if err != nil {
		return "", err
	}
	inputData := types.TWDataCreateWithGoBytes(inputBytes)
	defer C.TWDataDelete(inputData)

	outputData := C.TWAnySignerSign(inputData, C.TWCoinTypeEthereum)
	defer C.TWDataDelete(outputData)

	var output pb.SigningOutput
	if err := proto.Unmarshal(types.TWDataGoBytes(outputData), &output); err != nil {
		return "", err
	}
	fmt.Println("<== ethereum signed tx: ", hex.EncodeToString(output.Encoded))

	return hex.EncodeToString(output.Encoded), nil
}
