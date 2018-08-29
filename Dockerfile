FROM debian:stretch-slim

ARG bin
COPY $bin  /usr/local/bin/secreto
ADD secreto-linux-x86_64  /usr/bin/secreto
RUN chmod +x /usr/bin/secreto
ENTRYPOINT [ "secreto help" ]
