FROM golang:1.22.1 as build
RUN mkdir /src
ADD ./dynamicWelcomeService/*.go /src
ADD ./dynamicWelcomeService/go.mod /src
ADD ./dynamicWelcomeService/go.sum /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o dynamicWelcomeService
RUN chmod +x dynamicWelcomeService

FROM ghcr.io/chainguard-images/static:latest
COPY --from=build /src/dynamicWelcomeService /usr/local/bin/dynamicWelcomeService
EXPOSE 8080
CMD ["/usr/local/bin/dynamicWelcomeService"]
