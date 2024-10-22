## Fika backend individually company managed

## Tech stack

Basically, follow the latest version.

- paas: gcp(cloud run, cloud sql, and...)
- lang: Golang
- db: postgres
- fw: fiber
- orm: Gorm, gorm.io
- container: docker
- ci: GitHub action
- API Method: restAPI(with openAPI)

## Development environment setup

### Create containers for development

```
make setup-devbox
```

### Run project

```
docker compose up -d --build
```

## Other Development Tools

### How to openAPI generate
```
make oapi-codegen
```

### How to database specification


### How to generate database model for gorm

```
make gorm-gen
```






