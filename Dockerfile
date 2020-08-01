FROM scratch

COPY ./build/mockhttp-linux-* /usr/bin/mockhttp

CMD ["/usr/bin/mockhttp"]
