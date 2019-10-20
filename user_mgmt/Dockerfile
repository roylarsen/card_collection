FROM golang:latest as build

WORKDIR /app

COPY . .

CMD go mod download

#RUN go test

RUN CGO_ENABLED=0 go build -o user_mgmt .

FROM alpine

COPY --from=build /app/user_mgmt .

EXPOSE 1323

CMD ["/user_mgmt"]