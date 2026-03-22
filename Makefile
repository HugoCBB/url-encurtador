GREEN  := \033[1;32m
BLUE   := \033[1;34m
YELLOW := \033[1;33m
RESET  := \033[0m

.PHONY: run stop clean

run:
	@echo "$(GREEN)1. Subindo infraestrutura (Docker)...$(RESET)"
	@docker compose up -d
	
	@echo "$(YELLOW)2. Aguardando Redis ficar pronto...$(RESET)"
	@sleep 2
	
	@echo "$(BLUE)3. Iniciando servidor Go...$(RESET)"
	@go run ./cmd/server/main.go

stop:
	@echo "$(YELLOW)Parando containers...$(RESET)"
	@docker compose down

clean:
	@echo "$(YELLOW)Limpando binários e cache...$(RESET)"
	@go clean
	@docker compose down -v