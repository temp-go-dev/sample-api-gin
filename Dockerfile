FROM golang:1.12 AS build
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN ls -la /build
RUN go get -d -v
RUN go build -a -installsuffix cgo -o sample-api-gin .
RUN ls -la /build/sample-api-gin

FROM scratch AS runtime
COPY --from=build /build/sample-api-gin ./
COPY --from=build /build/environment.env ./
EXPOSE 8080/tcp
ENTRYPOINT ["./sample-api-gin"]
