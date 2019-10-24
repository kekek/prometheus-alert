FROM golang:1.13

MAINTAINER jinyf, jinyf@ktkt.com

ENV PROPATH /project
ENV PATH $PROPATH/bin:$PATH

RUN mkdir -p "$PROPATH" && chmod -R 777 "$PROPATH"

WORKDIR $PROPATH

VOLUME ["/project"]

CMD ["make","build", "run"]