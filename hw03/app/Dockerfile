FROM golang:1.7.3
WORKDIR /app
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest  
COPY --from=0 /app/app .
EXPOSE 8000
CMD ["./app"]
