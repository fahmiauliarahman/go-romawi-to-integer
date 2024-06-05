### Golang Romawi To Integer

1. Clone this repository
2. Run `go mod tidy`

### Test Result (without goroutine)
```
go test -run TestRomawiToInteger -v

=== RUN   TestRomawiToIntegerGoroutine_ReturnErrIfEmptyString
--- PASS: TestRomawiToIntegerGoroutine_ReturnErrIfEmptyString (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckRomawiValidity
--- PASS: TestRomawiToIntegerGoroutine_CheckRomawiValidity (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckQuadSymbol
--- PASS: TestRomawiToIntegerGoroutine_CheckQuadSymbol (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckDoubleVLD
--- PASS: TestRomawiToIntegerGoroutine_CheckDoubleVLD (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckAbnormalRomawi
--- PASS: TestRomawiToIntegerGoroutine_CheckAbnormalRomawi (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckNormalRomawi
--- PASS: TestRomawiToIntegerGoroutine_CheckNormalRomawi (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckMixedRomawi
--- PASS: TestRomawiToIntegerGoroutine_CheckMixedRomawi (0.00s)
=== RUN   TestRomawiToInteger_ReturnErrIfEmptyString
--- PASS: TestRomawiToInteger_ReturnErrIfEmptyString (0.00s)
=== RUN   TestRomawiToInteger_CheckRomawiValidity
--- PASS: TestRomawiToInteger_CheckRomawiValidity (0.00s)
=== RUN   TestRomawiToInteger_CheckQuadSymbol
--- PASS: TestRomawiToInteger_CheckQuadSymbol (0.00s)
=== RUN   TestRomawiToInteger_CheckDoubleVLD
--- PASS: TestRomawiToInteger_CheckDoubleVLD (0.00s)
=== RUN   TestRomawiToInteger_CheckAbnormalRomawi
--- PASS: TestRomawiToInteger_CheckAbnormalRomawi (0.00s)
=== RUN   TestRomawiToInteger_CheckNormalRomawi
--- PASS: TestRomawiToInteger_CheckNormalRomawi (0.00s)
=== RUN   TestRomawiToInteger_CheckMixedRomawi
--- PASS: TestRomawiToInteger_CheckMixedRomawi (0.00s)
PASS
ok      github.com/fahmiauliarahman/go_romawi_to_integer        0.256s
```

### Test Result (with goroutine)
```
go test -run TestRomawiToIntegerGoroutine -v

=== RUN   TestRomawiToIntegerGoroutine_ReturnErrIfEmptyString
--- PASS: TestRomawiToIntegerGoroutine_ReturnErrIfEmptyString (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckRomawiValidity
--- PASS: TestRomawiToIntegerGoroutine_CheckRomawiValidity (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckQuadSymbol
--- PASS: TestRomawiToIntegerGoroutine_CheckQuadSymbol (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckDoubleVLD
--- PASS: TestRomawiToIntegerGoroutine_CheckDoubleVLD (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckAbnormalRomawi
--- PASS: TestRomawiToIntegerGoroutine_CheckAbnormalRomawi (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckNormalRomawi
--- PASS: TestRomawiToIntegerGoroutine_CheckNormalRomawi (0.00s)
=== RUN   TestRomawiToIntegerGoroutine_CheckMixedRomawi
--- PASS: TestRomawiToIntegerGoroutine_CheckMixedRomawi (0.00s)
PASS
ok      github.com/fahmiauliarahman/go_romawi_to_integer        0.201s
```

### Benchmark Result
```
go test -bench=. -run=^$ -benchmem -v

goos: windows
goarch: amd64
pkg: github.com/fahmiauliarahman/go_romawi_to_integer
cpu: 12th Gen Intel(R) Core(TM) i5-12500H
BenchmarkRomawiToIntegerGoroutine_I
BenchmarkRomawiToIntegerGoroutine_I-16                            315339              3672 ns/op             328 B/op          7 allocs/op
BenchmarkRomawiToIntegerGoroutine_MMM
BenchmarkRomawiToIntegerGoroutine_MMM-16                          292467              4521 ns/op             328 B/op          7 allocs/op
BenchmarkRomawiToIntegerGoroutine_MMMDCCCLXXXVIII
BenchmarkRomawiToIntegerGoroutine_MMMDCCCLXXXVIII-16              220267              5245 ns/op             328 B/op          7 allocs/op
BenchmarkRomawiToInteger_I
BenchmarkRomawiToInteger_I-16                                    6745620               163.1 ns/op             0 B/op          0 allocs/op
BenchmarkRomawiToInteger_MMM
BenchmarkRomawiToInteger_MMM-16                                  4841301               248.8 ns/op             0 B/op          0 allocs/op
BenchmarkRomawiToInteger_MMMDCCCLXXXVIII
BenchmarkRomawiToInteger_MMMDCCCLXXXVIII-16                      1518942               781.1 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/fahmiauliarahman/go_romawi_to_integer        8.781s
```