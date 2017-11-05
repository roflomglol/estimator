FROM golang:onbuild

RUN ["apt-get", "update"]

EXPOSE 3001
