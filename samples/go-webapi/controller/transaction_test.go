package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/ubozov/wallet-core/samples/go-webapi/middleware"
)

const (
	seed = "observe drum fault concert analyst old short plunge loan essence symbol invite"
)

var requestTests = []struct {
	url          string
	method       string
	bodyData     string
	expectedCode int
	response     string
	msg          string
}{
	{
		"/api/v1/sign_transaction/",
		"POST",
		`{
			"gate": "bitcoin",
			"tx": {
				"toAddress": "1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
				"changeAddress": "1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU",
				"byteFee": 1,
				"amount": 1000000,
				"utxo" : {
					"hash" : "fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f",
					"index": 0,
					"sequence": 4294967295,
					"amount": 625000000
				}
			}
		}
		`,
		http.StatusOK,
		`{"data":"01000000000101fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f0000000000ffffffff0240420f00000000001976a914769bdff96a02f9135a1d19b749db6a78fe07dc9088ac007c3125000000001976a9149e089b6889e032d46e3b915a3392edfd616fb1c488ac02473044022051274f712832cbe10f6cc49aad2bcb0512ff8dae3e7c8a7711519dfaa2991ccc02207bd5a196e8a0d8a061d4828ea4d6a104db181d89595d00aa44327371c0ebed09012103df8635c33eae028e20d912120ee1a7304c3bb1e454e69682b16fcf1e3a3128fd00000000","full_messages":["TX was signed successfully"],"success":true}`,
		"valid data and should return StatusOK",
	},
	{
		"/api/v1/sign_transaction/",
		"POST",
		`{
			"gate": "energi",
			"tx": {
				"toAddress": "1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
				"changeAddress": "1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU",
				"byteFee": 1,
				"amount": 1000000,
				"utxo" : {
					"hash" : "fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f",
					"index": 0,
					"sequence": 4294967295,
					"amount": 625000000
				}
			}
		}
		`,
		http.StatusInternalServerError,
		`{"success":false,"full_messages":null,"errors":"unsupported"}`,
		"invalid data and should return StatusInternalServerError",
	},
	{
		"/api/v1/sign_transaction/",
		"POST",
		``,
		http.StatusInternalServerError,
		`{"success":false,"full_messages":null,"errors":"EOF"}`,
		"invalid data and should return StatusInternalServerError",
	},
}

func Test_SignTransaction(t *testing.T) {
	asserts := assert.New(t)

	r := gin.New()

	r.Use(middleware.Mnemonic(seed))

	RegisterSignTransactionRoutes(r.Group("/api/v1/sign_transaction/"))

	for _, testData := range requestTests {
		bodyData := testData.bodyData
		req, err := http.NewRequest(testData.method, testData.url, bytes.NewBufferString(bodyData))
		req.Header.Set("Content-Type", "application/json")
		asserts.NoError(err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		asserts.Equal(testData.expectedCode, w.Code, "Response Status - "+testData.msg)
		asserts.Equal(testData.response, w.Body.String(), "Response Content - "+testData.msg)
	}

}
