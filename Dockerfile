FROM golang:1.21-alpine as build

WORKDIR /app

COPY . .

RUN go build -o /bin/server .

FROM scratch

COPY --from=build /bin/server /bin/server

EXPOSE 8080

CMD ["/bin/server"]
