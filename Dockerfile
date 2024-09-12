# Import Base Image Builder
FROM golang:alpine3.20 as builder

# Setting Working Directory
WORKDIR /app

# Copy Source Code
COPY . .

# Download Depencies
RUN go mod download

# Compile Project
RUN go build -o main


# Runner Image Golang
FROM alpine:3.20.3

# Setting Working Directory
WORKDIR /app

# Setting ENV Golang APP
ENV APP_MODE=production
ENV HOST=0.0.0.0
ENV PORT=3000

ENV DB_HOST=127.0.0.1
ENV DB_PORT=3306
ENV DB_USERNAME=root
ENV DB_PASSWORD=
ENV DB_NAME=

ENV JWT_SECRET=
ENV JWT_EXPIRED_MINUTE=5

# Copy Compile Golang App
COPY --from=builder /app/main /app/main

# Expose Port
EXPOSE 3000

# Running Project
CMD [ "/app/main" ]