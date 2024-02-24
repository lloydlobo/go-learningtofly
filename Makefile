ROOTDIR := .
PKGDIR := $(ROOTDIR)/pkg
BUILD_CMD := go build
PATH_FILE_WASM := assets/json.wasm 

GOFLAGS := "-s -w"

# Build all packages
all: build

# Build individual packages (Windows)
#
# With flags:
# 	go build -ldflags="$(GOFLAGS)" "$(PKGDIR)/%%d"
build:
	@echo "Building packages in $(PKGDIR)..."
	@for /f "tokens=*" %%d in ('dir /ad /b "$(PKGDIR)"') do ( \
		go build "$(PKGDIR)/%%d" \
		)

# ^ older version for above...
# | build:
# | 	@echo "Building packages in $(PKGDIR)..."
# | 	@$(foreach dir,$(wildcard $(PKGDIR)/*), \
# | 		$(BUILD_CMD) $(dir);)

# ^ older version for above...
# | .PHONY: $(PKGDIR)/*
# | $(PKGDIR)/*:
# | 	@echo "Building package $(@)"
# | 	cd $(@) && go build $(GOFLAGS)

clean:
	@echo "Cleaning build artifacts..."
	$(MAKE) -C $(PKGDIR) clean

test:
	@echo "Running tests for all packages..."
	$(MAKE) -C $(PKGDIR) test

.PHONY: $(PKGDIR)/*_test
$(PKGDIR)/*_test:
	@echo "Running tests for package $(@F)"
	cd $(@D) && go test

# Help message
help:
	@echo "Available commands:"
	@echo "  all                  Build all packages."
	@echo "  build                Build all packages."
	@echo "  clean                Clean build artifacts."
	@echo "  test                 Run tests for all packages."
	@echo "  help                 Show this help message."



