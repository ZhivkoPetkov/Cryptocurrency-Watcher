# Cryptocurrency-Watcher
Cryptocurrency price watcher, developed with GO. The Kraken API is used for fetching of the actual prices.

* The price is in USD
* The difference is calculated based on the current price and the opening price of the day


## Example result
![alt text](http://i.ibb.co/QP0Br9T/Screenshot-2021-10-24-151319.png)

## RUN

To run the application:

```
CD home
go run .
```

If you need to make it as an .exe file
```
CD home
go build .
```

### If you need to exclude any crypto currency, remove the corresponding row in pairs.txt file:

```
XBTUSD,XXBTZUSD,Bitcoin
ETHUSD,XETHZUSD,Etherium
DOTUSD,DOTUSD,Polkadot
ADAUSD,ADAUSD,Cardano
SOLUSD,SOLUSD,Solana
XRPUSD,XXRPZUSD,Ripple
DOGEUSD,XDGUSD,DogeCoin
UNIUSD,UNIUSD,UniSwap
LINKUSD,LINKUSD,ChainLink
LTCUSD,XLTCZUSD,LiteCoin
```

