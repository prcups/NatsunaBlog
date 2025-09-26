.PHONY: build deploy

build: backend/blog frontend/.output

backend/blog:
	cd backend && go build

frontend/.output:
	cd frontend && npm run build

deploy: build
	if [ -d output ]; then rm -r output; fi
	mkdir -p output/backend/static
	cp -r backend/{blog,config} output/backend
	cp -r frontend/.output output/frontend
	cp -r frontend/{ecosystem.config.js,assets} output/frontend/
	tar cf output.tar.gz output
	rm -r output
