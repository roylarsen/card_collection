FROM golang:latest as build

WORKDIR /app

COPY . .

CMD go mod download

RUN go test

RUN CGO_ENABLED=0 go build -o lookup .

FROM alpine

COPY --from=build /app/lookup /

EXPOSE 1323

CMD ["./lookup"]