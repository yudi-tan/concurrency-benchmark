# Overview

Comparison of different language/frameworks' performance on I/O heavy tasks vs CPU-intensive tasks.

# I/O Heavy Tasks (simulated by sleeping for 5 secs between requests)

Apache Bench Result with 100 calls and 100 concurrent requests.

Ranked in order (fastest to slowest):

# Golang Benchmark

![alt text](https://github.com/YudiTan/concurrency-benchmark/blob/master/ioheavy/golang.png)

# Node.js Benchmark

![alt text](https://github.com/YudiTan/concurrency-benchmark/blob/master/ioheavy/nodejs.png)

# Flask Benchmark

![alt text](https://github.com/YudiTan/concurrency-benchmark/blob/master/ioheavy/flaskpy.png)

# CPU-Intensive Tasks (simulated by running deeply-recursive fibonacci generator for the 45th fibonacci number)

Apache Bench Result with 5 calls and 5 concurrent requests.

Ranked in order (fastest to slowest):

# Golang Benchmark

![alt text](https://github.com/YudiTan/concurrency-benchmark/blob/master/cpuheavy/goabtest.png)

# Node.js with PM2 (multi-core/cluster mode) on 4 cores

![alt text](https://github.com/YudiTan/concurrency-benchmark/blob/master/cpuheavy/pm2abtest.png)

![alt text](https://github.com/YudiTan/concurrency-benchmark/blob/master/cpuheavy/pm2results.png)

# Node.js without PM2 (single-threaded)

![alt text](https://github.com/YudiTan/concurrency-benchmark/blob/master/cpuheavy/nonpm2results.png)

# Python Flask using pypy just-in-time compiler instead of cpython (using cpython, which is the default, would be too slow to even test) and running only 2 requests (any amount of requests above 3 time-outs the AB test)

![alt text](https://github.com/YudiTan/concurrency-benchmark/blob/master/cpuheavy/pythonabtest.png)
