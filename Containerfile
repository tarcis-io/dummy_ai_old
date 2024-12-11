FROM golang:1.23.4 AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app ./cmd/dummy_ai

FROM scratch

WORKDIR /
COPY --from=build /app/app /app

EXPOSE 3000

CMD ["/app"]
