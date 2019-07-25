FROM alpine

COPY gopath/bin/gcp-cd-codelab /go/bin/gcp-cd-codelab
COPY ./image/Linty6.jpg /image/Linty6.jpg
ENTRYPOINT /go/bin/gcp-cd-codelab
