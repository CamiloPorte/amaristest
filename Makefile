runq1:
	@echo "Running question 1"
	@go run ./q1/main.go
	@echo "programa terminado"

runq2:
	@echo "Running question 2"
	@cd q2 && go run ./cmd/main.go
	@echo "programa terminado"

runq3:
	@echo "Running question 3"
	@go run ./q3/main.go
	@echo "programa terminado"

test:
	@echo "ejecutando test en todas las preguntas"
	@echo "pregunta 1"
	@cd q1 && go test ./...
	@echo "pregunta 2"
	@cd q2 && go test ./...
	@echo "pregunta 3"
	@cd q3 && go test ./...