FROM alpine

COPY gopath/bin/gcp-cd-codelab /go/bin/gcp-cd-codelab
# Copy the content of the /image directory into the newly built container for the images to remain accessible
COPY ./image /image
ENTRYPOINT /go/bin/gcp-cd-codelab
