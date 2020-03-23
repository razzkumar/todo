FROM golang:alpine as builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .


## copy only build file
#FROM scratch
FROM scratch

WORKDIR /api

COPY --from=builder /build/main .
# Export necessary port
EXPOSE 8888

# Command to run when starting the container
CMD ["./main"]
