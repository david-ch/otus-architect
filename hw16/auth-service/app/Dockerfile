FROM golang:1.14.2

# Step 1. Download dependencies, so docker can cache them
WORKDIR /app
ENV GO111MODULE=on
ENV PORT=9000
COPY go.mod .
COPY go.sum .
RUN go mod download

# Step 2. Build app
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Image for publishing
FROM alpine:latest  
COPY --from=0 /app/app .
EXPOSE 8000
CMD ["./app"]