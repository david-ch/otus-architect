FROM golang:1.14.2
WORKDIR /app
COPY *.go ./
RUN go get github.com/gorilla/mux && go get github.com/jackc/pgx && CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest  
COPY --from=0 /app/app .
EXPOSE 8000
CMD ["./app"]
