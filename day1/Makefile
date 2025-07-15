run:
	@go run main.go list.go solver.go

test:  
	@go test ./... -v -cover
benchmark:  
	@go test ./... -bench=. -benchmem

.PHONY: run test
