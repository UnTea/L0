# WRK tests
## Results for ```wrk -t12 -c400 --latency -d1m http://127.0.0.1:8080```

| Thread Stats | Avg      | Stdev    | Max   | +/- Stdev |
|--------------|----------|----------|-------|-----------|
| **Latency**  | 139.03ms | 201.77ms | 1.99s | 84.99%    |
| **Req/Sec**  | 0.89k    | 459.84   | 3.47k | 71.14%    |

| Latency | Distribution |
|---------|--------------|
| 50%     | 17.34ms      |
| 75%     | 212.35ms     |
| 90%     | 367.75ms     |
| 99%     | 752.27ms     |

- **640245** requests in **1.00m**, **826.12MB** read
- Socket errors: connect **0**, read **0**, write **0**, timeout **6**
- Requests/sec: **10655.20**
- Transfer/sec: **13.75MB**

## Results for ```wrk -t12 -c400 --latency -d1m http://127.0.0.1:8080/orders/a3c325b7f2a82c5_cucumber```

| Thread Stats | Avg      | Stdev    | Max   | +/- Stdev |
|--------------|----------|----------|-------|-----------|
| **Latency**  | 187.10ms | 257.88ms | 2.00s | 84.64%    |
| **Req/Sec**  | 449.93   | 212.36   | 2.04k | 71.64%    |

| Latency | Distribution |
|---------|--------------|
| 50%     | 44.03ms      |
| 75%     | 275.15ms     |
| 90%     | 504.85ms     |
| 99%     | 1.04s        |

- **322431** requests in **1.00m**, **1.10GB** read
- Socket errors: connect **0**, read **0**, write **0**, timeout **52**
- Requests/sec: **5366.38**
- Transfer/sec: **18.74MB**

## Results for ```wrk -t12 -c400 --latency -d1m http://127.0.0.1:8080/orders/f363bc34229456_tomato```

| Thread Stats | Avg      | Stdev    | Max   | +/- Stdev |
|--------------|----------|----------|-------|-----------|
| **Latency**  | 194.16ms | 267.44ms | 2.00s | 84.15%    |
| **Req/Sec**  | 434.77   | 194.10   | 1.94k | 72.60%    |

| Latency | Distribution |
|---------|--------------|
| 50%     | 46.48ms      |
| 75%     | 324.80ms     |
| 90%     | 567.78ms     |
| 99%     | 1.15s        |

- **311538** requests in **1.00m**, **1.22GB** read
- Socket errors: connect **0**, read **0**, write **0**, timeout **77**
- Requests/sec: **5184.16**
- Transfer/sec: **20.87MB**

## Results for ```wrk -t12 -c400 --latency -d1m http://127.0.0.1:8080/orders/b563feb7b2b84b6_test```


| Thread Stats | Avg      | Stdev    | Max   | +/- Stdev |
|--------------|----------|----------|-------|-----------|
| **Latency**  | 167.38ms | 229.49ms | 1.99s | 84.29%    |
| **Req/Sec**  | 482.06   | 195.42   | 2.03k | 73.01%    |

| Latency | Distribution |
|---------|--------------|
| 50%     | 45.83ms      |
| 75%     | 278.71ms     |
| 90%     | 507.84ms     |
| 99%     | 1.05s        |

- **345325** requests in **1.00m**, **1.17GB** read
- Socket errors: connect **0**, read **0**, write **0**, timeout **17**
- Requests/sec: **5749.98**
- Transfer/sec: **19.99MB**
