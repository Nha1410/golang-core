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

### Generate and set a personal access token

1. Generate a personal access token from [Github > Fine-grained personal access tokens](https://github.com/settings/tokens?type=beta) (You need to wait for the approval by the resource owner)
    - Token name: Read and write fikaigo-backend-common
    - Expiration: 90 days
    - Description: FIKAIGO backend common module
    - Resource Owner: Sumitomo-corporation
    - Repository access: Only select repository > fikaigo-backend-common
    - Repository permission: Contents > Read and Write
2. Copy the `.env.example` file and create a file named `.env`
3. Put your personal access token after the string `GIT_HUB_APP_TOKEN=`

example:
```
GIT_HUB_APP_TOKEN=your_personal_access_token
```


### Create containers for development

```
make setup-devbox
```

### GCP Application Default Credentials
When running in a local environment, please do application-default login in advance.
docker compose reads the credentials created by application-default login.
```
cloud auth application-default login
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






