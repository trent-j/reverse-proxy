ARG GO_VERSION=1.15

FROM octofactory.githubapp.com/github-docker/go/build:$GO_VERSION AS build

WORKDIR /app

COPY . .

RUN make build

FROM octofactory.githubapp.com/github-docker/go/release:$GO_VERSION

WORKDIR /app

COPY --from=build /app/bin ./
COPY --from=build /app/bp-dev.crt ./bp-dev.crt
COPY --from=build /app/bp-dev.key ./bp-dev.key
COPY --from=build /app/root-ca.pem /usr/local/share/ca-certificates/root-ca.pem

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates --fresh

ENTRYPOINT ["./reverse-proxy"]
