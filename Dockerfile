FROM fedora:32 AS builder
RUN dnf install -y go && dnf clean all
ADD . /app/
WORKDIR /app
RUN go get -d -v 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN ls

FROM scatch
MAINTAINER Michael Scherer <misc@redhat.com>
COPY --from=builder /app/app /not-all-bot
CMD ["/not-all-bot"]
