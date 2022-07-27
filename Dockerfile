FROM golang:1.18 as build
WORKDIR /go/src/app

RUN apt update && apt install curl jq sed gcc -y

# Install go-swagger cli
RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') \
  && curl -o /usr/local/bin/swagger -L'#' "$download_url" \
  && chmod +x /usr/local/bin/swagger \
  && swagger version

COPY . .
COPY todo.db /go/bin/todo.db

RUN go mod download
RUN go fmt ./cmd 
RUN go vet -v ./cmd

# Generating api documentation
RUN swagger generate spec -o /go/bin/swagger.yaml -w ./cmd -m && \
    sed -i 's/:5001/:5555/g' /go/bin/swagger.yaml

RUN CGO_ENABLED=1 go build -o /go/bin/app ./cmd

FROM gcr.io/distroless/base
ENV ENV=Development
EXPOSE 80

COPY --from=build /go/bin /bin

CMD ["/bin/app"]