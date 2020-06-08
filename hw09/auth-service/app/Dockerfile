FROM golang:1.14.2
WORKDIR /app
ADD . ./
RUN ls -l
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest  
COPY --from=0 /app/app .
EXPOSE 8000
CMD ["./app"]
