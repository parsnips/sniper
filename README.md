Sniper
------------

TastyWorks apis on the command line.

```
mparsons@snipsbook:~/go/src/github.com/parsnips/sniper$ ./sniper
There are many little commands here.  We can open quote streamer websocket, call search apis.. etc etc.

Usage:
  sniper [command]

Available Commands:
  getStreamerToken Retrieve a streamer token for live quote data
  help             Help about any command
  optionChain      Get the option chain for an underlying
  quotes           Open a websocket to receive quotes for a set of underlyings
  search           Search for an underlying by symbol
  session          Retrieve a session token from tastyworks

Flags:
      --config string   config file (default is $HOME/.snipe.yaml)
  -h, --help            help for sniper
  -t, --toggle          Help message for toggle

Use "sniper [command] --help" for more information about a command.
```

TastyNotes
---------------------


## Get a session token

request:

```
curl 'https://api.tastyworks.com/sessions' -H 'Origin: https://trade.tastyworks.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'Authorization: null' -H 'Content-Type: application/json' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw/login' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36' -H 'Connection: keep-alive' -H 'Accept-Version: v1' --data-binary '{"login":"parsnips@gmail.com","password":"password"}' --compressed
```


response body:

```
{
  "data": {
    "user": {
      "email": "parsnips@gmail.com",
      "username": "parsnips",
      "external-id": "U0000634596"
    },
    "session-token": "rJvtTwp7U4bGzD7q95-8IKjvHBR-8rMkpCVXryotBLx1BnpKMUyLag+C"
  },
  "api-version": "v1",
  "context": "/sessions"
}
```

## Validate


request:

```
curl 'https://api.tastyworks.com/sessions/validate' -X POST -H 'Authorization: su646fo0_Y5GhXBKVq2FysdrY-2JTfEHy86Fc3BHuA0eOwXnXgwqqA+C' -H 'Origin: https://trade.tastyworks.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw' -H 'Connection: keep-alive' -H 'Accept-Version: v1' -H 'Content-Length: 0' --compressed
```

response:

```
{
  "data": {
    "email": "parsnips@gmail.com",
    "username": "parsnips",
    "external-id": "U0000634596",
    "id": 57529
  },
  "api-version": "v1",
  "context": "/sessions/validate"
}
```


## Quote Streamer Tokens

Looks like this is what you use to open up the quote stream

request:

```
curl 'https://api.tastyworks.com/quote-streamer-tokens' -H 'Authorization: su646fo0_Y5GhXBKVq2FysdrY-2JTfEHy86Fc3BHuA0eOwXnXgwqqA+C' -H 'Origin: https://trade.tastyworks.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw' -H 'Connection: keep-alive' -H 'Accept-Version: v1' --compressed
```

response:

```
{
  "data": {
    "token": "dGFzdHksbGl2ZSwsMTUyMjk1ODgyOSwxNTIyODcyNDI5LFUwMDAwNjM0NTk2.Y0dq3VlEVI0jN1ANP5ZuV5YPTG0GZVc1iUSHNkhw_xM",
    "streamer-url": "tasty.dxfeed.com:7301",
    "websocket-url": "https://tasty.dxfeed.com/live",
    "level": "live"
  },
  "context": "/quote-streamer-tokens"
}
```


when subscribed to streamer.tastyworks.com:


first frame
```
{"action":"account-subscribe","value":["5WT54885"],"request-id":null,"auth-token":"su646fo0_Y5GhXBKVq2FysdrY-2JTfEHy86Fc3BHuA0eOwXnXgwqqA+C"}
```

heartbeats:

```
{"action":"heartbeat","request-id":null,"auth-token":"su646fo0_Y5GhXBKVq2FysdrY-2JTfEHy86Fc3BHuA0eOwXnXgwqqA+C"}
```

### Quote Stream

#### Request frames
```
[{"ext":{"com.devexperts.auth.AuthToken":"dGFzdHksbGl2ZSwsMTUyMjk1ODgyOSwxNTIyODcyNDI5LFUwMDAwNjM0NTk2.Y0dq3VlEVI0jN1ANP5ZuV5YPTG0GZVc1iUSHNkhw_xM"},"id":"1","version":"1.0","minimumVersion":"1.0","channel":"/meta/handshake","supportedConnectionTypes":["websocket","long-polling","callback-polling"],"advice":{"timeout":60000,"interval":0}}]
```

```
[{"id":"2","channel":"/service/sub","data":{"reset":true},"clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

```
[{"id":"3","channel":"/meta/connect","connectionType":"websocket","advice":{"timeout":0},"clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

```
[{"id":"4","channel":"/meta/connect","connectionType":"websocket","clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

```
[{"id":"5","channel":"/service/sub","data":{"addTimeSeries":{"Candle":[{"eventSymbol":"AAPL{=1d}","fromTime":1520197632460}]}},"clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

```
[{"id":"6","channel":"/service/sub","data":{"add":{"Trade":["SPY","/ES","VIX","/NQ","AAPL"],"Quote":["SPY","/ES","VIX","/NQ","AAPL"],"Summary":["SPY","/ES","VIX","/NQ","AAPL"],"Profile":["SPY","/ES","VIX","/NQ","AAPL"]}},"clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

```
[{"id":"7","channel":"/service/sub","data":{"removeTimeSeries":{"Candle":["AAPL{=1d}"]}},"clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

```
[{"id":"8","channel":"/service/sub","data":{"add":{"Summary":[".NFLX180518C305",".NFLX180518C310",".MSFT180518P85",".MSFT180518P87.5",".AAPL180518C180",".AAPL180518C175","NFLX","MSFT"],"Quote":[".NFLX180518C305",".NFLX180518C310",".MSFT180518P85",".MSFT180518P87.5",".AAPL180518C180",".AAPL180518C175","NFLX","MSFT"],"Greeks":[".NFLX180518C305",".NFLX180518C310",".MSFT180518P85",".MSFT180518P87.5",".AAPL180518C180",".AAPL180518C175"],"Trade":["NFLX","MSFT"],"Profile":["NFLX","MSFT"]}},"clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

```
[{"id":"9","channel":"/service/sub","data":{"add":{"Trade":["QQQ","BA","/ZB","/GC","AMZN","IYR","TSLA"],"Quote":["QQQ","BA","/ZB","/GC","AMZN","IYR","TSLA"],"Summary":["QQQ","BA","/ZB","/GC","AMZN","IYR","TSLA"],"Profile":["QQQ","BA","/ZB","/GC","AMZN","IYR","TSLA"]}},"clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

```
[{"id":"10","channel":"/meta/connect","connectionType":"websocket","clientId":"34k11jrbfdqraxj7d1ivux821vpr1i"}]
```

#### Responses

Looks like ticker?
```
[{"data":["Candle",["AAPL{=d}",0,0,6540376198348800000,153024,164.88,172.01,164.77,171.61,34556981,168.96691,14576949,15666416,0,6540376198348800000,1522800000000,"DEFAULT"]],"channel":"/service/timeSeriesData"}]
```

A trade?
```
[{"data":["Trade",["AAPL",0,6540685435994112000,0,"Q",171.61,1867298,34556983,5.848872E9,0,1522872000000000000,"ZERO_UP",false,1522872000000]],"channel":"/service/data"}]
```

A quote?
```
[{"data":["Quote",["AAPL",0,0,1522874096000,"P",171.53,5,1522874002000,"Q",171.6,2,0,1522874096000000000,1522874096000,"/ZB",0,0,1522874096000,"G",145.5,230,1522874097000,"G",145.53125,325,0,1522874097000000000,1522874097000]],"channel":"/service/data"}]
```

## Accounts

request:

```
curl 'https://api.tastyworks.com/customers/me/accounts' -H 'Authorization: su646fo0_Y5GhXBKVq2FysdrY-2JTfEHy86Fc3BHuA0eOwXnXgwqqA+C' -H 'Origin: https://trade.tastyworks.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw' -H 'Connection: keep-alive' -H 'Accept-Version: v1' --compressed
```

response:

```
{"data":{"items":[{"authority-level":"owner","account":{"account-number":"5WT54885","external-id":"A0000634702","opened-at":"2018-02-28T06:33:40.020+00:00","nickname":"Individual","account-type-name":"Individual","day-trader-status":false,"is-firm-error":false,"is-firm-proprietary":false,"margin-or-cash":"Margin","is-foreign":false,"funding-date":"2018-03-15","investment-objective":"GROWTH"}}]},"context":"/customers/me/accounts"}
```

## Positions

request:

```
curl 'https://api.tastyworks.com/accounts/5WT54885/positions' -H 'Authorization: su646fo0_Y5GhXBKVq2FysdrY-2JTfEHy86Fc3BHuA0eOwXnXgwqqA+C' -H 'Origin: https://trade.tastyworks.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw' -H 'Connection: keep-alive' -H 'Accept-Version: v1' --compressed
```

response:

```
{"data":{"items":[{"account-number":"5WT54885","symbol":"NFLX  180518C00305000","instrument-type":"Equity Option","underlying-symbol":"NFLX","quantity":1,"quantity-direction":"Short","close-price":"0.0","average-open-price":"13.63","average-yearly-market-close-price":"13.63","mark":"1240.0","cost-effect":"Debit","is-suppressed":false,"is-frozen":false,"realized-day-gain":"0.0","realized-day-gain-effect":"None","realized-day-gain-date":"2018-04-04","created-at":"2018-04-04T16:26:34.527+00:00","updated-at":"2018-04-04T17:37:47.806+00:00"},{"account-number":"5WT54885","symbol":"NFLX  180518C00310000","instrument-type":"Equity Option","underlying-symbol":"NFLX","quantity":1,"quantity-direction":"Long","close-price":"0.0","average-open-price":"11.88","average-yearly-market-close-price":"11.88","mark":"1077.5","cost-effect":"Credit","is-suppressed":false,"is-frozen":false,"realized-day-gain":"0.0","realized-day-gain-effect":"None","realized-day-gain-date":"2018-04-04","created-at":"2018-04-04T16:26:34.447+00:00","updated-at":"2018-04-04T17:37:47.807+00:00"},{"account-number":"5WT54885","symbol":"MSFT  180518P00085000","instrument-type":"Equity Option","underlying-symbol":"MSFT","quantity":1,"quantity-direction":"Long","close-price":"0.0","average-open-price":"2.27","average-yearly-market-close-price":"2.27","mark":"206.0","cost-effect":"Credit","is-suppressed":false,"is-frozen":false,"realized-day-gain":"0.0","realized-day-gain-effect":"None","realized-day-gain-date":"2018-04-04","created-at":"2018-04-04T15:37:09.306+00:00","updated-at":"2018-04-04T17:37:47.807+00:00"},{"account-number":"5WT54885","symbol":"MSFT  180518P00087500","instrument-type":"Equity Option","underlying-symbol":"MSFT","quantity":1,"quantity-direction":"Short","close-price":"0.0","average-open-price":"3.14","average-yearly-market-close-price":"3.14","mark":"290.0","cost-effect":"Debit","is-suppressed":false,"is-frozen":false,"realized-day-gain":"0.0","realized-day-gain-effect":"None","realized-day-gain-date":"2018-04-04","created-at":"2018-04-04T15:37:09.393+00:00","updated-at":"2018-04-04T17:37:47.807+00:00"},{"account-number":"5WT54885","symbol":"AAPL  180518C00180000","instrument-type":"Equity Option","underlying-symbol":"AAPL","quantity":1,"quantity-direction":"Long","close-price":"0.0","average-open-price":"2.82","average-yearly-market-close-price":"2.82","mark":"263.5","cost-effect":"Credit","is-suppressed":false,"is-frozen":false,"realized-day-gain":"0.0","realized-day-gain-effect":"None","realized-day-gain-date":"2018-04-04","created-at":"2018-04-04T16:23:42.759+00:00","updated-at":"2018-04-04T17:37:47.808+00:00"},{"account-number":"5WT54885","symbol":"AAPL  180518C00175000","instrument-type":"Equity Option","underlying-symbol":"AAPL","quantity":1,"quantity-direction":"Short","close-price":"0.0","average-open-price":"4.6","average-yearly-market-close-price":"4.6","mark":"430.0","cost-effect":"Debit","is-suppressed":false,"is-frozen":false,"realized-day-gain":"0.0","realized-day-gain-effect":"None","realized-day-gain-date":"2018-04-04","created-at":"2018-04-04T16:23:42.836+00:00","updated-at":"2018-04-04T17:37:47.808+00:00"}]},"api-version":"v1","context":"/accounts/5WT54885/positions"}
```


## Balances


request:

```
curl 'https://api.tastyworks.com/accounts/5WT54885/balances' -H 'Authorization: su646fo0_Y5GhXBKVq2FysdrY-2JTfEHy86Fc3BHuA0eOwXnXgwqqA+C' -H 'Origin: https://trade.tastyworks.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw' -H 'Connection: keep-alive' -H 'Accept-Version: v1' --compressed
```

response:

```
{"data":{"account-number":"5WT54885","cash-balance":"2472.458","long-equity-value":"0.0","short-equity-value":"0.0","long-derivative-value":"1547.0","short-derivative-value":"1960.0","long-futures-value":"0.0","short-futures-value":"0.0","debit-margin-balance":"0.0","long-margineable-value":"0.0","short-margineable-value":"0.0","margin-equity":"0.0","equity-buying-power":"2444.916","derivative-buying-power":"1222.458","day-trading-buying-power":"0.0","futures-margin-requirement":"0.0","available-trading-funds":"0.0","maintenance-requirement":"0.0","maintenance-call-value":"0.0","reg-t-call-value":"0.0","day-trading-call-value":"0.0","day-equity-call-value":"0.0","net-liquidating-value":"2059.458","cash-available-to-withdraw":"2039.31","day-trade-excess":"2039.31","pending-cash":"0.0","snapshot-date":"2018-04-04"},"api-version":"v1","context":"/accounts/5WT54885/balances"}
```


## Trading Status

request:

```
curl 'https://api.tastyworks.com/accounts/5WT54885/trading-status' -H 'Authorization: su646fo0_Y5GhXBKVq2FysdrY-2JTfEHy86Fc3BHuA0eOwXnXgwqqA+C' -H 'Origin: https://trade.tastyworks.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw' -H 'Connection: keep-alive' -H 'Accept-Version: v1' --compressed
```

response:

```
{"data":{"id":48450,"account-number":"5WT54885","can-trade-options":false,"options-level":"No Restrictions","option-exposure-amount":"0.0","is-frozen":false,"is-closing-only":false,"is-in-margin-call":false,"is-pattern-day-trader":false,"day-trade-count":4,"is-in-day-trade-equity-maintenance-call":false,"equities-margin-calculation-type":"Reg T","has-intraday-equities-margin":false,"is-closed":false,"is-futures-enabled":false,"is-futures-intra-day-enabled":false,"futures-margin-rate-multiplier":"0.0","is-futures-closing-only":false,"short-calls-enabled":true},"api-version":"v1","context":"/accounts/5WT54885/trading-status"}
```


## AN TSLA QUOTE

[{"data":["Quote",[".TSLA180615C225",0,0,1526414395000,"Z",60.05,4,1526414395000,"Z",61.75,11,0,1526414395000000000,1526414395000]],"channel":"/service/data"}]

[
  {
    "data": [
      "Quote",
      [
        ".TSLA180615C225",
        0,
        0,
        1526414395000,
        "Z",
        60.05,
        4,
        1526414395000,
        "Z",
        61.75,
        11,
        0,
        1526414395000000000,
        1526414395000
      ]
    ],
    "channel": "/service/data"
  }
]


[
  {
    "data": [
      "Quote",
      [
        ".TSLA180615C320",
        0,
        0,
        1526414400000,
        "N",
        2.93,
        1,
        1526414396000,
        "N",
        3.1,
        7,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615C315",
        0,
        0,
        1526414399000,
        "X",
        3.75,
        30,
        1526414399000,
        "N",
        3.95,
        6,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615P270",
        0,
        0,
        1526414399000,
        "X",
        9.35,
        12,
        1526414400000,
        "N",
        9.6,
        5,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615P265",
        0,
        0,
        1526414399000,
        "X",
        7.85,
        27,
        1526414400000,
        "N",
        8.1,
        5,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615C310",
        0,
        0,
        1526414399000,
        "I",
        4.85,
        7,
        1526414398000,
        "N",
        5.05,
        4,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615C305",
        0,
        0,
        1526414399000,
        "X",
        6.15,
        13,
        1526414392000,
        "N",
        6.35,
        3,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615P260",
        0,
        0,
        1526414400000,
        "H",
        6.6,
        2,
        1526414400000,
        "N",
        6.8,
        5,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615P300",
        0,
        0,
        1526414399000,
        "N",
        23.7,
        3,
        1526414398000,
        "H",
        24.2,
        1,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615P255",
        0,
        0,
        1526414399000,
        "X",
        5.5,
        25,
        1526414399000,
        "X",
        5.75,
        25,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615P250",
        0,
        0,
        1526414399000,
        "Z",
        4.65,
        20,
        1526414400000,
        "N",
        4.85,
        5,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615P245",
        0,
        0,
        1526414398000,
        "N",
        3.9,
        3,
        1526414399000,
        "Z",
        4.1,
        43,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615P240",
        0,
        0,
        1526414397000,
        "N",
        3.3,
        2,
        1526414397000,
        "X",
        3.45,
        25,
        0,
        1526414397000000000,
        1526414397000,
        ".TSLA180615P295",
        0,
        0,
        1526414399000,
        "N",
        20.55,
        5,
        1526414399000,
        "X",
        21.15,
        22,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615C340",
        0,
        0,
        1526414400000,
        "N",
        1,
        5,
        1526414400000,
        "N",
        1.09,
        2,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615C335",
        0,
        0,
        1526414390000,
        "N",
        1.28,
        6,
        1526414399000,
        "X",
        1.43,
        22,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615P290",
        0,
        0,
        1526414399000,
        "Z",
        17.7,
        13,
        1526414399000,
        "N",
        18.25,
        6,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615C290",
        0,
        0,
        1526414400000,
        "N",
        11.7,
        1,
        1526414399000,
        "Q",
        12,
        4,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615C330",
        0,
        0,
        1526414400000,
        "N",
        1.74,
        2,
        1526414400000,
        "Z",
        1.87,
        4,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615C325",
        0,
        0,
        1526414400000,
        "N",
        2.27,
        2,
        1526414396000,
        "N",
        2.39,
        3,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615P280",
        0,
        0,
        1526414399000,
        "X",
        12.95,
        25,
        1526414400000,
        "N",
        13.35,
        6,
        0,
        1526414400000000000,
        1526414400000,
        ".TSLA180615C280",
        0,
        0,
        1526414399000,
        "H",
        16.75,
        2,
        1526414396000,
        "N",
        17,
        1,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180615P275",
        0,
        0,
        1526414399000,
        "X",
        11,
        25,
        1526414399000,
        "X",
        11.35,
        38,
        0,
        1526414399000000000,
        1526414399000
      ]
    ],
    "channel": "/service/data"
  }
]



[
  {
    "data": [
      "Quote",
      [
        ".TSLA180720C250",
        0,
        0,
        1526414395000,
        "X",
        43.5,
        20,
        1526414398000,
        "H",
        44.2,
        10,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720P350",
        0,
        0,
        1526414398000,
        "Z",
        68.75,
        9,
        1526414349000,
        "Z",
        70,
        2,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720C350",
        0,
        0,
        1526414388000,
        "Z",
        3.05,
        32,
        1526414393000,
        "X",
        3.3,
        56,
        0,
        1526414393000000000,
        1526414393000,
        ".TSLA180720C245",
        0,
        0,
        1526414397000,
        "X",
        47.25,
        29,
        1526414393000,
        "X",
        48.55,
        24,
        0,
        1526414397000000000,
        1526414397000,
        ".TSLA180720P345",
        0,
        0,
        1526414399000,
        "Z",
        64.4,
        10,
        1526414399000,
        "Z",
        65.95,
        12,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C345",
        0,
        0,
        1526414393000,
        "Z",
        3.65,
        33,
        1526414395000,
        "Z",
        3.9,
        60,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720C240",
        0,
        0,
        1526414396000,
        "Z",
        51.1,
        7,
        1526414385000,
        "X",
        52.4,
        24,
        0,
        1526414396000000000,
        1526414396000,
        ".TSLA180720P340",
        0,
        0,
        1526414392000,
        "Z",
        60.1,
        10,
        1526414392000,
        "Z",
        61.55,
        12,
        0,
        1526414392000000000,
        1526414392000,
        ".TSLA180720C340",
        0,
        0,
        1526414399000,
        "Z",
        4.3,
        29,
        1526414356000,
        "X",
        4.55,
        44,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720P335",
        0,
        0,
        1526414398000,
        "Z",
        56,
        12,
        1526414399000,
        "Z",
        57.4,
        13,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C295",
        0,
        0,
        1526414399000,
        "B",
        16.65,
        5,
        1526414398000,
        "Z",
        17.2,
        20,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720P395",
        0,
        0,
        1526414398000,
        "Z",
        110.85,
        8,
        1526414398000,
        "Z",
        112.9,
        11,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720P330",
        0,
        0,
        1526414397000,
        "X",
        52.05,
        29,
        1526414399000,
        "X",
        53.15,
        33,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C395",
        0,
        0,
        1526414355000,
        "Z",
        0.6,
        20,
        1526414378000,
        "X",
        0.79,
        40,
        0,
        1526414378000000000,
        1526414378000,
        ".TSLA180720C290",
        0,
        0,
        1526414399000,
        "X",
        18.9,
        7,
        1526414398000,
        "Z",
        19.3,
        15,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720P390",
        0,
        0,
        1526414395000,
        "Z",
        106.05,
        8,
        1526414395000,
        "Z",
        107.95,
        11,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720P325",
        0,
        0,
        1526414396000,
        "X",
        48.1,
        22,
        1526414399000,
        "X",
        49.15,
        33,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C390",
        0,
        0,
        1526414367000,
        "Z",
        0.72,
        19,
        1526414385000,
        "Z",
        0.92,
        37,
        0,
        1526414385000000000,
        1526414385000,
        ".TSLA180720P385",
        0,
        0,
        1526414398000,
        "Z",
        101.2,
        8,
        1526414398000,
        "Z",
        103.05,
        11,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720C285",
        0,
        0,
        1526414398000,
        "T",
        21.3,
        6,
        1526414398000,
        "Z",
        21.8,
        2,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720C325",
        0,
        0,
        1526414391000,
        "Z",
        7.1,
        3,
        1526414393000,
        "X",
        7.35,
        28,
        0,
        1526414393000000000,
        1526414393000,
        ".TSLA180720P280",
        0,
        0,
        1526414398000,
        "B",
        20.65,
        5,
        1526414399000,
        "Z",
        21.1,
        22,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720P320",
        0,
        0,
        1526414397000,
        "X",
        44.35,
        27,
        1526414399000,
        "X",
        45.3,
        33,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C385",
        0,
        0,
        1526414370000,
        "Z",
        0.87,
        19,
        1526414395000,
        "X",
        1.06,
        33,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720C280",
        0,
        0,
        1526414396000,
        "X",
        23.9,
        20,
        1526414395000,
        "Z",
        24.7,
        20,
        0,
        1526414396000000000,
        1526414396000,
        ".TSLA180720P380",
        0,
        0,
        1526414395000,
        "Z",
        96.4,
        8,
        1526414395000,
        "Z",
        98.3,
        11,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720P275",
        0,
        0,
        1526414399000,
        "X",
        18.5,
        27,
        1526414399000,
        "Z",
        18.95,
        24,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720P315",
        0,
        0,
        1526414399000,
        "X",
        40.8,
        20,
        1526414384000,
        "X",
        41.6,
        24,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C380",
        0,
        0,
        1526414355000,
        "Z",
        1.03,
        19,
        1526414392000,
        "Z",
        1.24,
        37,
        0,
        1526414392000000000,
        1526414392000,
        ".TSLA180720C275",
        0,
        0,
        1526414398000,
        "X",
        26.75,
        25,
        1526414356000,
        "X",
        27.5,
        20,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720P375",
        0,
        0,
        1526414395000,
        "Z",
        91.65,
        8,
        1526414395000,
        "Z",
        93.5,
        11,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720C315",
        0,
        0,
        1526414395000,
        "Z",
        9.55,
        23,
        1526414395000,
        "Z",
        9.95,
        25,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720P310",
        0,
        0,
        1526414398000,
        "A",
        37.35,
        22,
        1526414398000,
        "A",
        38.25,
        21,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720C375",
        0,
        0,
        1526414393000,
        "Z",
        1.24,
        30,
        1526414393000,
        "Z",
        1.45,
        40,
        0,
        1526414393000000000,
        1526414393000,
        ".TSLA180720C270",
        0,
        0,
        1526414394000,
        "X",
        29.7,
        20,
        1526414396000,
        "N",
        30.6,
        4,
        0,
        1526414396000000000,
        1526414396000,
        ".TSLA180720P370",
        0,
        0,
        1526414395000,
        "Z",
        86.95,
        8,
        1526414395000,
        "Z",
        88.75,
        11,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720P265",
        0,
        0,
        1526414398000,
        "M",
        14.85,
        5,
        1526414399000,
        "B",
        15.15,
        1,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C310",
        0,
        0,
        1526414398000,
        "Z",
        11.05,
        18,
        1526414395000,
        "Z",
        11.5,
        25,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720P305",
        0,
        0,
        1526414395000,
        "Z",
        34.1,
        13,
        1526414395000,
        "X",
        34.85,
        24,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720C370",
        0,
        0,
        1526414397000,
        "Z",
        1.49,
        30,
        1526414376000,
        "X",
        1.7,
        36,
        0,
        1526414397000000000,
        1526414397000,
        ".TSLA180720C265",
        0,
        0,
        1526414395000,
        "A",
        32.95,
        25,
        1526414398000,
        "Z",
        33.9,
        11,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720P365",
        0,
        0,
        1526414395000,
        "Z",
        82.3,
        8,
        1526414395000,
        "Z",
        84.05,
        11,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720C305",
        0,
        0,
        1526414399000,
        "Z",
        12.7,
        18,
        1526414397000,
        "Z",
        13.2,
        22,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720P260",
        0,
        0,
        1526414399000,
        "A",
        13.2,
        15,
        1526414395000,
        "Z",
        13.55,
        32,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C365",
        0,
        0,
        1526414391000,
        "Z",
        1.79,
        30,
        1526414376000,
        "Z",
        2.01,
        60,
        0,
        1526414391000000000,
        1526414391000,
        ".TSLA180720C260",
        0,
        0,
        1526414396000,
        "A",
        36.25,
        20,
        1526414399000,
        "Z",
        37.3,
        13,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720P255",
        0,
        0,
        1526414396000,
        "B",
        11.8,
        5,
        1526414398000,
        "Z",
        12.1,
        29,
        0,
        1526414398000000000,
        1526414398000,
        ".TSLA180720P360",
        0,
        0,
        1526414399000,
        "Z",
        77.75,
        9,
        1526414399000,
        "X",
        79.25,
        28,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C300",
        0,
        0,
        1526414399000,
        "B",
        14.6,
        5,
        1526414399000,
        "Z",
        15.1,
        21,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C360",
        0,
        0,
        1526414396000,
        "Z",
        2.15,
        30,
        1526414396000,
        "Z",
        2.36,
        40,
        0,
        1526414396000000000,
        1526414396000,
        ".TSLA180720P355",
        0,
        0,
        1526414398000,
        "Z",
        73.2,
        9,
        1526414399000,
        "X",
        74.7,
        29,
        0,
        1526414399000000000,
        1526414399000,
        ".TSLA180720C255",
        0,
        0,
        1526414397000,
        "X",
        39.8,
        30,
        1526414385000,
        "X",
        40.95,
        24,
        0,
        1526414397000000000,
        1526414397000,
        ".TSLA180720P250",
        0,
        0,
        1526414395000,
        "Z",
        10.5,
        12,
        1526414395000,
        "Z",
        10.8,
        35,
        0,
        1526414395000000000,
        1526414395000,
        ".TSLA180720C355",
        0,
        0,
        1526414393000,
        "Z",
        2.58,
        31,
        1526414372000,
        "N",
        2.78,
        7,
        0,
        1526414393000000000,
        1526414393000
      ]
    ],
    "channel": "/service/data"
  }
]


### Quote Request

[
  {
    "id": "85",
    "channel": "/service/sub",
    "data": {
      "remove": {
        "Summary": [
          ".TSLA180720C225",
          ".TSLA180720P225",
          ".TSLA180720C230",
          ".TSLA180720P230",
          ".TSLA180720C235",
          ".TSLA180720C220",
          ".TSLA180720P220",
          ".TSLA180720C205",
          ".TSLA180720P205",
          ".TSLA180720C210",
          ".TSLA180720P210",
          ".TSLA180720C215",
          ".TSLA180720P215",
          ".TSLA180720C185",
          ".TSLA180720P185",
          ".TSLA180720C190",
          ".TSLA180720P190",
          ".TSLA180720C195",
          ".TSLA180720P195",
          ".TSLA180720C200",
          ".TSLA180720P200",
          ".TSLA180720C180",
          ".TSLA180720P180",
          ".TSLA180720C175",
          ".TSLA180720P175"
        ],
        "Quote": [
          ".TSLA180720C225",
          ".TSLA180720P225",
          ".TSLA180720C230",
          ".TSLA180720P230",
          ".TSLA180720C235",
          ".TSLA180720C220",
          ".TSLA180720P220",
          ".TSLA180720C205",
          ".TSLA180720P205",
          ".TSLA180720C210",
          ".TSLA180720P210",
          ".TSLA180720C215",
          ".TSLA180720P215",
          ".TSLA180720C185",
          ".TSLA180720P185",
          ".TSLA180720C190",
          ".TSLA180720P190",
          ".TSLA180720C195",
          ".TSLA180720P195",
          ".TSLA180720C200",
          ".TSLA180720P200",
          ".TSLA180720C180",
          ".TSLA180720P180",
          ".TSLA180720C175",
          ".TSLA180720P175"
        ],
        "Greeks": [
          ".TSLA180720C225",
          ".TSLA180720P225",
          ".TSLA180720C230",
          ".TSLA180720P230",
          ".TSLA180720C235",
          ".TSLA180720C220",
          ".TSLA180720P220",
          ".TSLA180720C205",
          ".TSLA180720P205",
          ".TSLA180720C210",
          ".TSLA180720P210",
          ".TSLA180720C215",
          ".TSLA180720P215",
          ".TSLA180720C185",
          ".TSLA180720P185",
          ".TSLA180720C190",
          ".TSLA180720P190",
          ".TSLA180720C195",
          ".TSLA180720P195",
          ".TSLA180720C200",
          ".TSLA180720P200",
          ".TSLA180720C180",
          ".TSLA180720P180",
          ".TSLA180720C175",
          ".TSLA180720P175"
        ]
      }
    },
    "clientId": "2rdyxys929iw2anx17oy19hndf53s"
  }
]



### Public Watchlists

curl 'https://trade.dough.com/api/public_watchlists?include_synthetic=true' -H 'X-Tastyworks: web' -H 'Origin: https://trade.tastyworks.com' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.167 Safari/537.36' -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw' -H 'Authorization: O98IM5IblqpWpaTmZw32-XF_YIB8eups83XDTY8CF9SBbfgpKD75Cg+C' -H 'Connection: keep-alive' --compressed


### Search

curl 'https://trade.dough.com/api/stocks/search' -H 'Origin: https://trade.tastyworks.com' -H 'X-Tastyworks: web' -H 'Accept-Encoding: gzip, deflate, br' -H 'Accept-Language: en-US,en;q=0.9' -H 'Authorization: O98IM5IblqpWpaTmZw32-XF_YIB8eups83XDTY8CF9SBbfgpKD75Cg+C' -H 'Content-Type: application/json' -H 'Accept: application/json, text/javascript, */*; q=0.01' -H 'Referer: https://trade.tastyworks.com/tw' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.167 Safari/537.36' -H 'Connection: keep-alive' --data-binary '{"symbols":["BABA","BIDU","C","CMG","EEM","EWW","FXI","GDX","GDXJ","GOOG","IBM","RUT","SPX","XLE","XLU","XOP"]}' --compressed

```
{"stocks":[{"symbol":"BABA","beta":"1.532","volatility_index":"29.37270928418381","description":"Alibaba Group Holding Limited American Depositary Shares each representing one Ordinary share","is_notable":false,"tos_volatility_percentile":"27.28","volatility_percentile":"27.28","liquidity_value":"452410.211312423578333333333333333333333333","liquidity_rank":"0.180053622535535393535838471180147484","liquidity_rating":4,"volatility_index_5_day_change":"-0.9771752260837161","has_weekly_options":true,"exchange":"NYSE","dividend_rate_per_share":null,"dividend_ex_date":null,"research_team":{"corr_SPY_3month":"0.67"},"notabilities":[],"expiration_volatilities":[{"expiration_date":"2018-06-15","option_chain_type":"standard","implied_volatility":"29.62961639473293"},{"expiration_date":"2018-06-08","option_chain_type":"standard","implied_volatility":"28.40186324394539"},{"expiration_date":"2018-06-29","option_chain_type":"standard","implied_volatility":"28.74102988277659"},{"expiration_date":"2019-06-21","option_chain_type":"standard","implied_volatility":"32.81149085020584"},{"expiration_date":"2020-01-17","option_chain_type":"standard","implied_volatility":"33.45916196871956"},{"expiration_date":"2019-01-18","option_chain_type":"standard","implied_volatility":"31.52582144845838"},{"expiration_date":"2018-07-20","option_chain_type":"standard","implied_volatility":"29.05083450646336"},{"expiration_date":"2018-10-19","option_chain_type":"standard","implied_volatility":"30.44087554488424"},{"expiration_date":"2018-05-18","option_chain_type":"standard","implied_volatility":"46.69727581595708"},{"expiration_date":"2018-05-25","option_chain_type":"standard","implied_volatility":"30.38102750711411"},{"expiration_date":"2018-06-01","option_chain_type":"standard","implied_volatility":"28.30904067983377"},{"expiration_date":"2018-06-22","option_chain_type":"standard","implied_volatility":"28.44331046868492"}],"earnings":{"expected_report_date":"2018-05-04","is_estimated":false,"time_of_day_code":"BTO","late_flag":0,"quarter_end_date":"2018-03-01","actual_eps":"0.39","consensus_estimate":"0.65","is_visible":true},"implied_volatility_percentile_input":{"snapshot_at":"2018-05-15","standard_deviation_30d":"4.652847468884848","mean_30d":"36.24016543072656","standard_deviation_60d":"4.245107207756023","mean_60d":"38.01665047019169","standard_deviation_6m":"4.90651919483004","mean_6m":"34.45627308985695","standard_deviation_1y":"4.561619869929783","mean_1y":"33.29326015141296"}}]}
```
