// validar cobetura dos testes

go test -coverprofile=coverage.out

// verificar linhas nao cobertas nos testes

go test -coverprofile=coverage.out && go tool cover -html=coverage.out