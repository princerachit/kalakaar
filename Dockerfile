FROM golang

ARG BINARY_NAME=kalakaar
COPY bin/$BINARY_NAME /app/kalakaar
WORKDIR /app
CMD ["./kalakaar"]
