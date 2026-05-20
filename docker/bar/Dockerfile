FROM golang:1.24 AS build-stage

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build ./cmd/bar

FROM gcr.io/distroless/base-debian11 AS run-stage

COPY --from=build-stage ./app/bar ./bar

EXPOSE 8080

ENTRYPOINT ["./bar"]
