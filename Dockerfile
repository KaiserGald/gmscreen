FROM golang:latest
ARG app_name
ARG src_path
ARG install_path
ENV GOBIN /go/bin
ENV BINARY_NAME $app_name
ENV INSTALLPATH $install_path
ENV APPROOT $src_path/$BINARY_NAME
RUN mkdir /srv/$BINARY_NAME
RUN mkdir -p $src_path/$BINARY_NAME
ADD . $src_path/$BINARY_NAME
RUN cd $src_path/$BINARY_NAME && make all
ENTRYPOINT $INSTALLPATH/$BINARY_NAME -c
EXPOSE 8080
EXPOSE 443
EXPOSE 8081
