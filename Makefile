DB_PORT?=1234
TMP_CONTAINER_NAME?=tmp_mysql

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

## Server
start: ## start container.
	docker run --name $(TMP_CONTAINER_NAME) --rm -it --mount type=bind,source=${CURDIR},target=/myapp -p ${DB_PORT}:3306 -e MYSQL_ROOT_PASSWORD=mysql mysql:8.0.30
attach: ## attach container
	docker exec -it $(TMP_CONTAINER_NAME) mysql -u root -pmysql test_db

## Init
init: ## init database
	docker exec -it $(TMP_CONTAINER_NAME) bash -c "mysql -u root -pmysql < /myapp/sql/create_database.sql"
	docker exec -it $(TMP_CONTAINER_NAME) bash -c "mysql -u root -pmysql test_db < /myapp/sql/init_table.sql"

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-30s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
