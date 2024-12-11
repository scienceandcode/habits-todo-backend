# Nome do diretório de dados do PostgreSQL
POSTGRES_DATA_DIR := .docker/data/postgres

# Comando para criar o diretório e aplicar permissões
.PHONY: setup
setup:
	@echo "📁 Criando o diretório $(POSTGRES_DATA_DIR) e configurando as permissões..."
	mkdir -p $(POSTGRES_DATA_DIR)
	chown -R $(shell id -u):$(shell id -g) $(POSTGRES_DATA_DIR)
	chmod -R 755 $(POSTGRES_DATA_DIR)
	@echo "✅ Tudo pronto! O diretório foi criado e as permissões foram configuradas."

# Limpa o diretório (opcional)
.PHONY: clean
clean:
	@echo "🗑️ Removendo o diretório $(POSTGRES_DATA_DIR)..."
	rm -rf $(POSTGRES_DATA_DIR)
