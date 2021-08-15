# DATABASE
DB_USER=postgres
DB_PASSWORD=bismillah
DB_HOST=localhost
DB_PORT=5432
DB_NAME=mufid
DB_SSL=disable

install:
	cd .. && go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate && \
	go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && \
	go get github.com/vektra/mockery/v2/.../ && \
	cd ${PROJECT_NAME} && swag init

local:
	air -c .air.toml

test:
	go test -v -cover -coverprofile=cover.out ./...

migrate-up:
	migrate -source file:./schema/migration/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} up

migrate-down:
	migrate -source file:./schema/migration/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} down 1

migrate-drop:
	migrate -source file:./schema/migration/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} drop

mockery-usecase:
	cd service/auth/usecase && mockery --name=Usecase --output=../mocks 
	
mockery-repo:
	cd service/auth/store && mockery --name=Repository --output=../mocks

rebase-dev:
	git checkout development && git pull origin development && git checkout @{-1} && git rebase development

rebase-release:
	git checkout release && git pull origin release && git checkout @{-1} && git rebase release

rebase-master:
	git checkout master && git pull origin master && git checkout @{-1} && git rebase master

push:
	go fmt ./... && git push origin HEAD

push-force:
	go fmt ./... && git push -f origin HEAD
