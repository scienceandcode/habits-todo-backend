# Nome do diretório de dados
DATA_DIR := .docker/data

# Nome do arquivo de configuração de ambiente
ENV_FILE := .env
ENV_EXAMPLE_FILE := .env.example

# Comando para criar o diretório e aplicar permissões
.PHONY: setup
setup: create-env create-dir
	@echo "✅ Tudo pronto! O diretório foi criado, o arquivo .env foi configurado e as permissões foram aplicadas."

# Alvo para criar o diretório
.PHONY: create-dir
create-dir:
	@echo "📁 Criando o diretório $(DATA_DIR) e configurando as permissões..."
	mkdir -p $(DATA_DIR)
	chown -R $(shell id -u):$(shell id -g) $(DATA_DIR)
	chmod -R 755 $(DATA_DIR)

# Alvo para copiar o arquivo .env.example para .env, se não existir
.PHONY: create-env
create-env:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "📄 Criando o arquivo $(ENV_FILE) a partir de $(ENV_EXAMPLE_FILE)..."; \
		cp $(ENV_EXAMPLE_FILE) $(ENV_FILE); \
	else \
		echo "📄 O arquivo $(ENV_FILE) já existe. Nenhuma ação necessária."; \
	fi

# Limpa o diretório (opcional)
.PHONY: clean
clean:
	@echo "🗑️ Removendo o diretório $(POSTGRES_DATA_DIR) e o arquivo $(ENV_FILE)..."
	rm -rf $(POSTGRES_DATA_DIR)
	rm -f $(ENV_FILE)
