export GOBIN = $(PWD)/bin
GQLGEN = $(GOBIN)/gqlgen
export SCRIPTS_ROOT_DIR= ${PWD}/scripts

OS = $(shell uname -s | tr A-Z a-z)

MAKEFILES_DIR = makefiles
MAKEFILES = $(shell find $(MAKEFILES_DIR) -name '*.mk')

include $(MAKEFILES)