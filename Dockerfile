# build stage
FROM golang:1.20-alpine AS build-env
RUN apk add --update make
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app
RUN CGO_ENABLED=0 GOOS=linux make

# final stage
FROM alpine:3.18
LABEL maintainer="m.vorobev"
WORKDIR /app
COPY --from=build-env /go/src/app/bin/app /app/
RUN chmod +x /app/app
CMD ["/app/app"]
