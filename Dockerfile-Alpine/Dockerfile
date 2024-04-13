# Dockerfile Alpine
FROM golang:1.22.1-alpine as build
RUN mkdir /src
ADD ./dynamicWelcomeService/*.go /src
ADD ./dynamicWelcomeService/go.mod /src
ADD ./dynamicWelcomeService/go.sum /src
WORKDIR /src
RUN go get -d -v -t
RUN GOOS=linux go build -v -o dynamicWelcomeService
RUN chmod +x dynamicWelcomeService

FROM alpine:3.19.1
COPY --from=build /src/dynamicWelcomeService /usr/local/bin/dynamicWelcomeService
EXPOSE 8080
CMD ["dynamicWelcomeService"]