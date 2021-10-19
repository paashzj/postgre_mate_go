FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o postgre_mate .


FROM ttbb/postgre:nake

LABEL maintainer="shoothzj@gmail.com"

COPY --chown=sh:sh docker-build /opt/sh/postgre/mate

COPY --from=build --chown=sh:sh /opt/sh/compile/pkg/postgre_mate /opt/sh/postgre/mate/postgre_mate

USER sh
CMD ["/usr/local/bin/dumb-init", "bash", "-vx", "/opt/sh/postgre/mate/scripts/start.sh"]