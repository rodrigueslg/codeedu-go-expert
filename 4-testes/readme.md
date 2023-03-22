# Executando testes

* rodando testes  
`go test`

* rodando testes em modo verboso  
`go test -v`

# Cobertura de testes

* Executando testes exibindo a % de cobertura:  
`go test -coverprofile=coverage.out`  

* Exibindo em modo gráfico a cobertura de código:  
`go tool cover -html=coverage.out`

# Performance / Benchmarking

* Benchmark
`go test -bench=.`

* Benchmark exibindo alocação de memória  
`go test -bench=. -benchmem`

* Benchmark sem testes  
`go test -bench=. -run=^#`

* Benchmark sem testes repetindo 5 vezes  
`go test -bench=. -run=^# -count=5`

* Benchmark sem testes repetindo 5 vezes demorando 3 segundos em cada execução  
`go test -bench=. -run=^# -count=5 -benchtime=3s`

# Fuzzing

* fuzz  
`go test -fuzz=. -run=^#`

* retestar fuzz que falhou (substituir hash pelo hash do seu erro)  
`    go test -run=FuzzCalculateTax/5fb97e24f60a89626b3a21def0907a63de860058ee5b4522f4bcb12ef1299fdd` 

* fuzz com limite de tempo de 5 segundos  
`go test -fuzz=. -fuzztime=5s -run=^#`