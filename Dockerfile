FROM alpine:3.21

# Update and upgrade
RUN apk update && apk upgrade

# Install dependencies:
# - Golang
RUN apk add --no-cache go

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
