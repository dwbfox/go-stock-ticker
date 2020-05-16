# go-stock-ticker
---
A small CLI utility to check tickers, written in `Go`

## Usage:

#### Download and Build
1. clone repository:
```bash
$ git@github.com:dwbfox/go-stock-ticker.git
```

2. `cd` into project directory
```
$ cd ./go-stock-ticker
```

3. Build the app
```bash
$ make all
```

#### Run

Look up stock by symbol
```bash
$ ./bin/go-stock-ticker --symbol VOO,TSLA,SPY,F,AAPL,AMD,C

+--------+--------------------------------+-------------------------------+---------+---------+-----------+
| SYMBOL |              NAME              |           INDUSTRY            |   ASK   |   BID   |  VOLUME   |
+--------+--------------------------------+-------------------------------+---------+---------+-----------+
| VOO    | Vanguard Group, Inc. -         | InvestmentTrustsOrMutualFunds | $263.09 | $263.06 |   3386161 |
|        | Vanguard S&P 500 ETF           |                               |         |         |           |
| TSLA   | Tesla Inc                      | MotorVehicles                 | $799.19 | $799.21 |  10518428 |
| SPY    | SSgA Active Trust - SSGA SPDR  | InvestmentTrustsOrMutualFunds | $286.26 | $286.25 | 111146276 |
|        | S&P 500                        |                               |         |         |           |
| F      | Ford Motor Co.                 | MotorVehicles                 | $4.91   | $4.90   |  80544650 |
| AAPL   | Apple Inc                      | TelecommunicationsEquipment   | $307.76 | $307.70 |  41587094 |
| AMD    | Advanced Micro Devices Inc.    | Semiconductors                | $54.18  | $54.17  |  66950863 |
| C      | Citigroup Inc                  | FinancialConglomerates        | $41.93  | $41.92  |  28494867 |
+--------+--------------------------------+-------------------------------+---------+---------+-----------+
```

Look up stock by symbol and return as JSON
```bash
$ ./bin/go-stock-ticker --symbol VOO,TSLA,SPY,F,AAPL,AMD,C --json

[{"TradingHalted":false,"Low52Weeks":200.55,"High52Weeks":311.59,"AskSize":500,"Ask":263.09,"BidSize":800,"Bid":263.06,"PercentChangeFromPreviousClose":0.451,"ChangeFromPreviousClose":1.18,"PreviousClose":261.86,"Volume":3386161,"LastSize":21382,"Last":263.04,"Low":258.56,"High":263.11,"Close":263.04,"Open":259.54,"UTCOffset":-4,"Delay":0.0087893,"Outcome":"Success","Security":{"MostLiquidExchange":false,"CategoryOrIndustry":"InvestmentTrustsOrMutualFunds","MarketIdentificationCode":"ARCX","Market":"NYSEARCA","Name":"Vanguard Group, Inc. - Vanguard S\u0026P 500 ETF","Valoren":"22423967","ISIN":null,"Symbol":"VOO","CUSIP":null,"CIK":"0000102909"},"IdentifierType":"Symbol","Identifier":"VOO","LastMarketIdentificationCode":"ARCX","AskMarketIdentificationCode":"ARCX","BidMarketIdentificationCode":"ARCX","Currency":"USD","AskTime":"4:00:00 PM","AskDate":"5/15/2020","BidTime":"4:00:00 PM","BidDate":"5/15/2020","PreviousCloseDate":"5/14/2020","Time":"4:00:00 PM","Date":"5/15/2020","Identity":"Request","Message":"Delay times are 15 mins for NYSEARCA."},{"TradingHalted":false,"Low52Weeks":176.9919,"High52Weeks":968.9899,"AskSize":600,"Ask":799.19,"BidSize":1000,"Bid":799.21,"PercentChangeFromPreviousClose":-0.518,"ChangeFromPreviousClose":-4.16,"PreviousClose":803.33,"Volume":10518428,"LastSize":51267,"Last":799.17,"Low":786.552,"High":805.0486,"Close":799.17,"Open":790.35,"UTCOffset":-4,"Delay":0.0049241,"Outcome":"Success","Security":{"MostLiquidExchange":false,"CategoryOrIndustry":"MotorVehicles","MarketIden[...]
```