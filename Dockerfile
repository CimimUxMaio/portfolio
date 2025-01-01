FROM golang:1.23.4-alpine

# Update and upgrade
RUN apk update && apk upgrade

# Copy project files
COPY . /app

# Set working directory
WORKDIR /app

# Build the project
RUN go install && go build -o ./builds/portfolio .

# Set production mode environment variable
ENV GIN_MODE=release

# Run the application
CMD ["./builds/portfolio"]
