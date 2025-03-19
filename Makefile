# Nome do diretório de dados
DATA_DIR := .docker/data

# Nome do arquivo de configuração de ambiente
ENV_FILE := .env
ENV_EXAMPLE_FILE := .env.example

# Nome do diretório GOPATH
GOPATH_DIR := $(shell grep ^GOPATH= $(ENV_FILE) | cut -d '=' -f 2)

# Verificação do Go
check-go:
	@which go > /dev/null || { echo "Go não encontrado! Instale o Go antes de continuar."; exit 1; }

# Comando para criar o diretório e aplicar permissões
.PHONY: setup
setup: check-go create-env create-dir set-gopath install-wire
	@echo "✅ Tudo pronto! O diretório foi criado, o arquivo .env foi configurado, o GOPATH foi configurado e o Wire foi instalado."

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

# Alvo para configurar o GOPATH
.PHONY: set-gopath
set-gopath:
	@if [ -z "$(GOPATH_DIR)" ]; then \
		echo "GOPATH não encontrado no arquivo .env"; \
		exit 1; \
	fi
	@echo "🔧 Configurando o GOPATH..."
	@if [ -n "$(shell echo $$SHELL | grep -E 'zsh')" ]; then \
		echo "export GOPATH=$(GOPATH_DIR)" >> ~/.zshrc; \
		echo "🚀 Lembre-se de reiniciar o terminal ou rodar 'source ~/.zshrc'"; \
	elif [ -n "$(shell echo $$SHELL | grep -E 'bash')" ]; then \
		echo "export GOPATH=$(GOPATH_DIR)" >> ~/.bashrc; \
		echo "🚀 Lembre-se de reiniciar o terminal ou rodar 'source ~/.bashrc'"; \
	else \
		echo "Shell não suportado. Configure o GOPATH manualmente."; \
	fi
	@echo "GOPATH configurado para $(GOPATH_DIR)"

# Alvo para instalar o Wire
.PHONY: install-wire
install-wire:
	@echo "🔧 Instalando o Wire..."
	$(shell $GOPATH/bin/go get github.com/google/wire/cmd/wire) 
	@echo "🔧 Wire instalado com sucesso!"

# Limpa o diretório (opcional)
.PHONY: clean
clean:
	@echo "🗑️ Removendo o diretório $(POSTGRES_DATA_DIR) e o arquivo $(ENV_FILE)..."
	rm -rf $(POSTGRES_DATA_DIR)
	rm -f $(ENV_FILE)
