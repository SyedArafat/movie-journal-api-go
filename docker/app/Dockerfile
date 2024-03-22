ARG GO_VERSION=1.22
ARG GO_IMAGE_TYPE=alpine

FROM golang:${GO_VERSION}-${GO_IMAGE_TYPE} as dev

ARG MAINTAINER="Sayed Yeamin Arafat <sayed.yeamin@brainstation-23.com>"

LABEL org.opencontainers.image.authors=${MAINTAINER}

ARG TIMEZONE="Asia/Dhaka"
ARG GOOS="linux"
ARG GOARCH="amd64"
ARG CGO_ENABLED=0

ARG HTTP_PROXY
ARG HTTPS_PROXY
ARG NO_PROXY

ARG HOST_OS

ENV http_proxy="${HTTP_PROXY}" \
    https_proxy="${HTTPS_PROXY}" \
    no_proxy="${NO_PROXY}" \
    HOST_OS="${HOST_OS}"

ENV GO111MODULE=on \
    CGO_ENABLED="${CGO_ENABLED}" \
    GOOS="${GOOS}" \
    GOARCH="${GOARCH}" \
    TZ="${TIMEZONE}"

ARG INSTALL_EXTRA_APK_PACKAGES="curl nano busybox-extras bash git tzdata"
ARG INSTALL_GO_DEV_TOOLS="github.com/cosmtrek/air@v1.51.0"

RUN echo "--- Set Timezone ---" \
        && echo "${TZ}" > /etc/timezone \
        && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo "-- Install Alpine Packages --" \
        && apk update \
        && apk upgrade -U -a \
        && if [ ! -z "${INSTALL_EXTRA_APK_PACKAGES}" ]; then \
            apk add --no-cache ${INSTALL_EXTRA_APK_PACKAGES}; \
        fi \
    && if [ ! -z "${INSTALL_GO_DEV_TOOLS}" ]; then \
            echo "-- Install Development Tools --" \
            && go install ${INSTALL_GO_DEV_TOOLS}; \
       fi \
    && echo "--- Clean Up ---" \
        && rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

WORKDIR /usr/src/app

# COPY go.mod, go.sum and download the dependencies
COPY ./src/go.* ./
RUN go mod download && go mod verify

# COPY All things inside the project and build
COPY ./src .

# Update go.mod file
RUN go mod tidy

RUN go build -o /goapp ./cmd/main/

COPY ./docker/app/dev.entrypoint.sh /

RUN chmod a+x /dev.entrypoint.sh

CMD ["sh", "/dev.entrypoint.sh"]

FROM alpine:latest as prod

ARG TIMEZONE="Asia/Dhaka"
ARG GOOS="linux"
ARG GOARCH="amd64"
ARG CGO_ENABLED=0

ARG ESSINTIAL_APK_PACKAGES_FOR_PROD="curl busybox-extras bash tzdata"

RUN echo "--- Set Timezone ---" \
        && echo "${TZ}" > /etc/timezone \
        && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo "-- Install Alpine Packages --" \
        && apk update \
        && apk upgrade -U -a \
        && if [ ! -z "${ESSINTIAL_APK_PACKAGES_FOR_PROD}" ]; then \
            apk add --no-cache ${ESSINTIAL_APK_PACKAGES_FOR_PROD}; \
        fi \
    && echo "--- Clean Up ---" \
        && rm -rf /var/cache/apk/* /tmp/* /var/tmp/*

ENV GO111MODULE=on \
    CGO_ENABLED="${CGO_ENABLED}" \
    GOOS="${GOOS}" \
    GOARCH="${GOARCH}" \
    TZ="${TIMEZONE}"

COPY --from=dev /goapp /goapp

COPY ./docker/app/entrypoint.sh /

RUN chmod a+x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
