FROM harbor.star-bit.io/golang/golang:latest
WORKDIR /go/src/rarity-backend
COPY . .
RUN go build .
EXPOSE 8844
ENTRYPOINT [ "./rarity-backend" ]