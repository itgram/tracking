# docker build --build-arg GO_VERSION=1.19.3 --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" -t bluemill/service .
# docker run --rm --env PING_URLS=GGDDD bluemill/service

# argument for Go version
ARG GO_VERSION

################################################################
# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS build

ARG SSH_PRIVATE_KEY

# git required for go mod
# certs
# add a user here because addgroup and adduser are not available in scratch
RUN apk add --no-cache git && \
    apk add --no-cache openssh && \
    apk --no-cache add ca-certificates && \
    addgroup -S myapp && \
    adduser -S -u 10000 -g myapp myapp

# Working directory will be created if it does not exist
WORKDIR /src

# add ssh key to permit the download of private modules
RUN export GOPRIVATE=github.com/itgram && \
    git config --global url."git@github.com:itgram".insteadOf "https://github.com/itgram"

# add credentials on build
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && chmod 400 /root/.ssh/id_rsa

# make sure your domain is accepted
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

# We use go modules; copy go.mod and go.sum
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import code
COPY . .

# fix the "proto: file "gossip.proto" is already registered" error
# RUN export ESDB_DIR=$(go list -m -f "{{.Dir}}" github.com/EventStore/EventStore-Client-Go/v3)

# RUN echo "${ESDB_DIR}" > /root/dd

RUN go install github.com/tuimeo/go-proto-filename-prefixer@latest
RUN go-proto-filename-prefixer $(go list -m -f "{{.Dir}}" github.com/EventStore/EventStore-Client-Go/v3)/protos esdb.

# # Run tests
# # RUN CGO_ENABLED=0 go test -timeout 30s -v github.com/gbaeke/go-template/pkg/api

# Build the executable
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    # -ldflags "-extldflags -static -w -s -compressdwarf=true" \
    # -ldflags "-extldflags -static -compressdwarf=false -m=2 -l" \
    -ldflags "-extldflags -static -compressdwarf=false" \
    -tags musl \
    -o /app


################################################################
# STAGE 2: build the container to run
FROM scratch AS final

# add maintainer label
LABEL maintainer="hussein"

# copy compiled app
COPY --from=build /app /app

# copy ca certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# copy users from builder which contains myapp user
COPY --from=0 /etc/passwd /etc/passwd

USER myapp

# run binary
ENTRYPOINT ["/app"]
# CMD ["rest-server", "-env", "/api/env.example"]