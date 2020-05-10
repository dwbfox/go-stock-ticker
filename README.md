# go-stock-ticker
---
A small CLI utility to check stock tickers, written in `Go`

## Usage:

#### Download and Build
1. clone repository:
```bash
$ git@github.com:dwbfox/go-stock-ticker.git
```

2. cd into project directory
```
$ cd ./go-stock-ticker
```

3. Build the app
```bash
$ make all
```

#### Run
```bash
$ ./build/go-stock-ticker --symbol VOO,TSLA,SPY,F,AAPL,AMD,C

+--------+---------+---------+-----------+
| SYMBOL |   ASK   |   BID   |  VOLUME   |
+--------+---------+---------+-----------+
| VOO    | $268.76 | $268.75 |   3206956 |
| TSLA   | $819.32 | $819.00 |  16130087 |
| SPY    | $292.47 | $292.46 |  76622128 |
| F      | $5.24   | $5.23   | 101333782 |
| AAPL   | $310.04 | $310.00 |  33511985 |
| AMD    | $53.11  | $53.10  |  40774186 |
| C      | $46.32  | $46.31  |  22919094 |
+--------+---------+---------+-----------+
```
