FROM golang:1.23

WORKDIR /go/src

# OPENAPI_GENERATOR_CLI_VERSION=latest \
ENV OAPI_CODEGEN_VERSION=latest \
    GENTOOL_VERSION=latest \
    AIR_VERSION=latest

RUN go install github.com/air-verse/air@${AIR_VERSION} && \
    go install gorm.io/gen/tools/gentool@${GENTOOL_VERSION} && \
    go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@${OAPI_CODEGEN_VERSION}

ENV DEV_ENV=development

CMD ["air"]
