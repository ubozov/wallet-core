## Build

```
make build
```

## Run

```
make run
``` 
or 
```
.go-webapi --config ./go-webapi.config.json
```


## Test

```
curl --location --request POST 'http://localhost:2322/api/v1/sign_transaction/' \
--header 'Content-Type: application/json' \
--data-raw '{
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
'
```