USER_BIN_DIR := $(shell echo $$HOME)/bin
USER_ZSH_COMPLETION_DIR := $(shell echo $$HOME)/.zsh/completions

# install shell & kubectl completion to executable path
install:
	cp kubectl* $(USER_BIN_DIR)

# install zsh completion 
install_zsh_completion:
	cp _kubectl* $(USER_ZSH_COMPLETION_DIR)