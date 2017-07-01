FROM golang:onbuild

RUN curl https://glide.sh/get | sh
RUN make install

EXPOSE 3000
