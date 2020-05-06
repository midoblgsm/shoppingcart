FROM golang as build
RUN mkdir shoppingcart
WORKDIR /shoppingcart
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/shoppingcart

FROM scratch
COPY --from=build /go/bin/shoppingcart /go/bin/shoppingcart
ENTRYPOINT ["/go/bin/shoppingcart"]