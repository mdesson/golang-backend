# Start from latest golang base image
FROM golang:latest

# Maintainer Info
LABEL maintainer="Michael Desson <mdesson@gmail.com>"

# set current workign directory insie container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source from current directory to container's workign directory
COPY . .

# Build go app
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the app
CMD ["./golang-backend"]
