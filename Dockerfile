FROM scratch

COPY ./build/mockhttp-linux-* /usr/bin/mockhttp
RUN chmod 755 /usr/bin/mockhttp

ENTRYPOINT ["/usr/bin/mockhttp"]
