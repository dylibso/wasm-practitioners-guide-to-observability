FROM golang:latest as build
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /host


FROM scratch
WORKDIR /
COPY --from=build /host .
COPY --from=build /tmp /tmp
EXPOSE 3000
CMD ["/host"]