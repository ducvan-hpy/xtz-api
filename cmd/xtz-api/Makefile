GO = go

TARGET_NAME = xtzapi
DIST_DIR    = dist

CODEGEN_BIN = $(GOPATH)/bin/oapi-codegen
CODEGEN_VER = latest

OPENAPI_FILE = xtz_oapi.yml
CODEGEN_FILE = api/xtzapi.gen.go

all: build

run: $(TARGET_NAME)
	$(DIST_DIR)/$<

build: $(TARGET_NAME)

$(TARGET_NAME): codegen
	mkdir -p $(DIST_DIR)
	$(GO) build -o $(DIST_DIR)/$@

$(CODEGEN_BIN):
	$(GO) install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@$(CODEGEN_VER)

codegen: $(CODEGEN_BIN)
	$(RM) -f $(CODEGEN_FILE)
	oapi-codegen -package api -old-config-style -generate 'gin,types' $(OPENAPI_FILE) > $(CODEGEN_FILE)

clean:
	$(RM) $(CODEGEN_FILE)
	$(RM) -r $(DIST_DIR)

.PHONY: all build clean codegen run
