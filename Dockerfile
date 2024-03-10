# Use a specific version of the golang image
FROM golang:1.22.1

# Set the working directory inside the container to /app
WORKDIR /app

# Copy the Go module files first to take advantage of Docker cache
# Assuming the go.mod and go.sum files are in the src directory alongside your source code
COPY src/go.mod src/go.sum ./

# Download the dependencies specified in go.mod and go.sum
RUN go mod download

# Now copy the rest of the source code from the src directory to the working directory
COPY src/ .

# Build the application to the ./main in the working directory
RUN go build -o main .

# Expose port 8888 on which the application will run
EXPOSE 8888

# Command to run the executable
CMD ["./main"]
