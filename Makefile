SHELL := /bin/zsh

.PHONY: deploy
deploy: book
	@echo "====> deploying to github"
	rm -rf /tmp/books
	go build tool/*.go
	git worktree add -f /tmp/books gh-pages
	rm -rf /tmp/books/*
	cp -rp book/* /tmp/books/
	cd /tmp/books && \
        git add -A && \
        git commit -m "deployed on $(shell date) by ${USER}" && \
        git push origin gh-pages
build:
	rm -rf book && go  run github.com/dtest11/alg/generator
run:
	rm -rf book && go  run github.com/dtest11/alg/generator
	cd book && python3 -m http.server
