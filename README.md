# coinconv
The sample tool to get quotes from CoinMarketCap service.<br>

## notes
The tool uses the sandbox of the CoinMarketCap service and testing key.

## usage
1. cointool 123.45 BTC USD - to get quote for conversion of 123.45 BTC to USD
2. cointool 123.45 BTC USD,EUR,GBP - to get quotes for conversion of 123.45 BTC to USD, EUR, and GBP

Getting quotes only for single currency is possible.

## gow to compile (Windows)
1. `cd cmd`
2. `go build -o bin/coinconv.exe main.go` 




