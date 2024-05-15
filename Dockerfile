FROM golang:bullseye

WORKDIR /usr/src/app
ENV PORT 3000
ENV HOST 0.0.0.0

COPY . .

# Ensure the .env file is copied (if it exists)
COPY .env .env

RUN go build -o bin .

ENTRYPOINT ["./bin"]