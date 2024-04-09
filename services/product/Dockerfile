#Dockerfile for the product service

#Use Golang image
FROM golang:1.22.1

#Create app directory
WORKDIR /app

#Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

#Copy source code
COPY . .

#Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /product-service ./cmd/main.go

EXPOSE 8081
#Run 
CMD ["/product-service"]