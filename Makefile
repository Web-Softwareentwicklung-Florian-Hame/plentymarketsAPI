export GOBIN = $(PWD)/bin
OS = $(shell uname -s | tr A-Z a-z)

MAKEFILES_DIR = makefiles
MAKEFILES = $(shell find $(MAKEFILES_DIR) -name '*.mk')

include $(MAKEFILES)