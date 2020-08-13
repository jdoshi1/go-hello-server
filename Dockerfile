
################ Build container stage ###############
# Start from the latest golang base image
FROM golang:latest as build

# Add Maintainer Info
LABEL maintainer="Jinesh Doshi<jdoshi1@asu.edu>"

# Set the current working dir inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-hello-server .

############### Run container stage ##################
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary from build stage
COPY --from=build /app/go-hello-server .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./go-hello-server" ]