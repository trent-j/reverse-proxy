ARG GO_VERSION=1.15

FROM octofactory.githubapp.com/github-docker/go/build:$GO_VERSION AS build

WORKDIR /app

COPY . .

RUN make build

FROM octofactory.githubapp.com/github-docker/go/release:$GO_VERSION

WORKDIR /app

COPY --from=build /app/bin ./

RUN apt-get update && \
    apt-get install -y ca-certificates

ENTRYPOINT bash -c "cp /tmp/certs/ghes-root-ca.pem /usr/local/share/ca-certificates/ && update-ca-certificates --fresh && ./reverse-proxy"