FROM --platform=$BUILDPLATFORM golang:1.23.4-alpine AS build
WORKDIR /app
COPY . .
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN go mod tidy && go mod download
RUN GOOS=$(echo $TARGETPLATFORM | cut -d'/' -f1) \
    GOARCH=$(echo $TARGETPLATFORM | cut -d'/' -f2) \
    go build -ldflags "-s -w" -o /semdgo ./cmd/server

FROM scratch
COPY --from=build /semdgo /semdgo
COPY --from=build /app/templates /templates
EXPOSE 80
CMD ["/semdgo"]