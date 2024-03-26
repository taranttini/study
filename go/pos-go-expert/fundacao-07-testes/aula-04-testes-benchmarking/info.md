// validar cobetura dos testes

go test -coverprofile=coverage.out

// verificar linhas nao cobertas nos testes

go test -coverprofile=coverage.out && go tool cover -html=coverage.out

// para rodar o benchmark 

go test -bench=.


// para rodar o benchmark testes que comecem com um valor especifico

go test -bench=. -run=^#


// para rodar o benchmark 10 vezes

go test -bench=. -run=^# -count=10

// para rodar o benchmark 10 vezes rodar 3s de cada teste

go test -bench=. -run=^# -count=10 -benchtime=3s

// go help text - listar opcoes de teste

go help text

// go help text - listar uso de memoria

go test -bench=. -run=^# -benchmem