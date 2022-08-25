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