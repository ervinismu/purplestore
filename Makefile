help: ## This help dialog
help h:
	@IFS=$$'\n' ; \
	help_lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##/:/'`); \
	printf "%-30s %s\n" "target" "help" ; \
	printf "%-30s %s\n" "------" "----" ; \
	for help_line in $${help_lines[@]}; do \
		IFS=$$':' ; \
		help_split=($$help_line) ; \
		help_command=`echo $${help_split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		help_info=`echo $${help_split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf '\033[36m'; \
		printf "%-30s %s" $$help_command ; \
		printf '\033[0m'; \
		printf "%s\n" $$help_info; \
	done

#===================#
#== Env Variables ==#
#===================#
DOCKER_COMPOSE_FILE ?= docker-compose.dev.yml


#========================#
#== DATABASE MIGRATION ==#
#========================#

migration-up: ## Run migrations UP
migration-up mu:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate up

migration-down: ## Rollback migrations against non test DB
migration-down md:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate down 1

migration-create: ## Create a DB migration files e.g `make migrate-create name=migration-name`
migration-create mc:
	docker compose -f ${DOCKER_COMPOSE_FILE} --profile tools run --rm migrate create -ext sql -dir /migrations $(name)


#==========================#
#== CONTAINER INTERACTION ==#
#==========================#

shell-db: ## Enter to database console
shell-db:
	docker compose -f ${DOCKER_COMPOSE_FILE} exec db psql -U postgres -d postgres

shell-api: ## Enter to api console
shell-api:
	docker compose -f ${DOCKER_COMPOSE_FILE} exec api /bin/sh


run : ## running docker compose services
run :
	docker compose -f ${DOCKER_COMPOSE_FILE} up
