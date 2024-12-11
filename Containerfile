FROM golang:1.23.4 AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=js GOARCH=wasm go build -C /cmd/dummy_ai -o /app

FROM scratch

WORKDIR /
COPY --from=build /app .

EXPOSE 3000

CMD ["/app"]
