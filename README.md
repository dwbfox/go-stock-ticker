# go-stock-ticker
---
A small CLI utility to check stock tickers, written in `Go`

## Usage:

### Build 
```bash
$ make
```

### Run
```bash
$ ./build/go-stock-ticker TSLA AMD SPY VOO

+--------+---------+---------+----------+
| SYMBOL |   ASK   |   BID   |  VOLUME  |
+--------+---------+---------+----------+
| TSLA   | $819.32 | $819.00 | 16130087 |
| AMD    | $53.11  | $53.10  | 40774186 |
| SPY    | $292.47 | $292.46 | 76622128 |
| VOO    | $268.76 | $268.75 |  3206956 |
+--------+---------+---------+----------+
```