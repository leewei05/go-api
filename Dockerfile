# golang:alpine is 370MB
# golang:1.13 is 803MB
# Of course we will choose golang:alpine, we want our Docker image to be as small as possible
FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set our building directory
WORKDIR /build

# Copy and download dependency using go mod
COPY . . 
RUN go mod download

# Build our go code into binary
RUN go build -o main .

# Build our final image using an explicitly empty image, scratch
FROM scratch 

# Set our execution directory and copy our binary
WORKDIR /app
COPY --from=builder /build/main /app/main
COPY --from=builder /build/config.env /app/config.env

# Expose our HTTP server to port 8080
EXPOSE 8080

# Command to run when starting the container
CMD ["/app/main"]


