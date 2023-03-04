# Purplestore

## Prerequisites
- install [Docker](https://docs.docker.com/engine/install/)

##

## Migration Commands

1. Create migration file
```
make migrate-create name=migration_name
```

2. Migrate Down / Rollback migration
```
make migrate-down
```
3. Migrate UP / Applt migration to database
```
make migrate-up
```

## Docker Compose Commands

1. Running docker compose
```
docker-compose up -d
```
2. List all running container
```
docker ps
```
