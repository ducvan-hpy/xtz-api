DIST_DIR = dist

XTZ_PATH = cmd/xtz-api

all: build-xtz-api

build-xtz-api: $(DIST_DIR)
	$(MAKE) -C $(XTZ_PATH)
	cp $(XTZ_PATH)/dist/* $(DIST_DIR)

$(DIST_DIR):
	mkdir $@

clean:
	$(MAKE) -C $(XTZ_PATH) $@
	$(RM) -r $(DIST_DIR)

run:
	$(MAKE) -C $(XTZ_PATH) $@

.PHONY: all build-xtz-api clean run
