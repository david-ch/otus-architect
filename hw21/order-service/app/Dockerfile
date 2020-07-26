FROM golang:1.14.2-alpine

# confluent go kafka client needs to be built in -alpine, with gcc, with -tags musl flag

# Step 1. Download dependencies, so docker can cache them
WORKDIR /app
ENV GO111MODULE=on
RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      musl-dev
COPY go.mod .
COPY go.sum .
RUN go mod download

# Step 2. Build app
ADD . ./
RUN GOOS=linux go build -tags musl -o app .

# Image for publishing
FROM alpine:latest  
COPY --from=0 /app/app .
EXPOSE 8000
CMD ["./app"]
