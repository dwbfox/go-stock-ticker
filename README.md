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
$ ./build/go-stock-ticker --symbol TSLA,AAPL

+--------+------+---------+---------+----------+
| SYMBOL | NAME |   ASK   |   BID   |  VOLUME  |
+--------+------+---------+---------+----------+
| TSLA   |      | $819.32 | $819.00 | 16130087 |
| AAPL   |      | $310.04 | $310.00 | 33511985 |
+--------+------+---------+---------+----------+
```