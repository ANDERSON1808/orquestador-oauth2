FROM golang:1.14 as build

ARG WORKDIR=/go/src/github.com/golang-oauth2-example/

WORKDIR "${WORKDIR}"

COPY go.mod go.sum "${WORKDIR}"/
RUN go mod download

COPY . "${WORKDIR}"

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -o /orquestador-oauth2

ENTRYPOINT ["/orquestador-oauth2"]

FROM scratch

WORKDIR /oauth2

# Copy the binary we built in the 'build' stage
COPY --from=build /orquestador-oauth2 .

# Copy across the CA certificates so that we can make TLS/SSL connections to things
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/oauth2/orquestador-oauth2"]
