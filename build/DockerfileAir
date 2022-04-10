FROM golang:bullseye as builder
WORKDIR /app

RUN apt-get update --no-install-recommends
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

ENTRYPOINT ["/bin/air"]