



.PHONY: image-tag
image-tag:
	@echo $(shell sed -n 2p docker/emissary.docker.tag.remote)