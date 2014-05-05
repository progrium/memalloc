FROM flynn/busybox
MAINTAINER Jeff Lindsay <progrium@gmail.com>

ADD ./build/mem /bin/mem

ENTRYPOINT ["/bin/mem"]
