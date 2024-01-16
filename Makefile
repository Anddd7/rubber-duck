COMPONENT_DIRS := $(wildcard *)

doc:
	@echo "Generating documentation"
	@for dir in $(COMPONENT_DIRS); do \
		if [ ! -d $$dir ]; then continue; fi; \
		echo "Generating documentation for $$dir"; \
		$(MAKE) -s -C $$dir doc; \
	done
