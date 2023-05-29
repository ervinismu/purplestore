# Purplestore

## Prerequisites
- install [Docker](https://docs.docker.com/engine/install/)

## Project Structures
- Layout

```
.
├── app.env.sample ( contains application configuration )
├── cmd ( main binary )
├── Makefile ( simplify project commands )
├── docs ( swagger documentation )
├── db
│   └── migrations ( database migrations )
└── internal
    ├── app
    │   ├── controllers ( request response handler )
    │   ├── models ( all about database table )
    │   ├── repository ( database/cache operation )
    │   ├── router ( http router )
    │   ├── schema ( request/response schema )
    │   └── service ( business logic )
    └── pkg ( private lib )
```
- Application flow
```
router --> middleware --> controllers(use schema) --> service --> repository(user model)
```

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
