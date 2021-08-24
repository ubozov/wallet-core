package crypto

import (
	"fmt"

	"go-webapi/crypto/bitcoin"
	"go-webapi/crypto/ethereum"
)

var fabric map[string]func(string, interface{}) (string, error)

func init() {
	fabric = make(map[string]func(string, interface{}) (string, error))

	fabric["bitcoin"] = bitcoin.Sign
	fabric["ethereum"] = ethereum.Sign
	// please, add supported coin here
}

func GetSigner(gate string) (func(string, interface{}) (string, error), error) {
	fn, ok := fabric[gate]
	if !ok {
		return nil, fmt.Errorf("unsupported")
	}
	return fn, nil
}
