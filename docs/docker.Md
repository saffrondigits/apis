# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port that your application is listening on (in your case, it's 8080)
EXPOSE 8080

# Start your application when the container starts
CMD ["./main"]

```
docker build -t your-app-name .
```



```
docker run -d -p 8080:8080 your-app-name
```
