FROM alpine:edge


RUN apk --update add --no-cache python2 py2-requests py2-pip py2-lxml py2-requests openssl ca-certificates
RUN apk --update add --virtual build-dependencies python2-dev build-base wget git \
  && git clone https://github.com/1N3/BlackWidow.git
WORKDIR BlackWidow

RUN cp blackwidow /usr/bin/blackwidow && cp injectx.py /usr/bin/injectx.py

RUN pip2 install -r requirements.txt
ENTRYPOINT ["python2", "blackwidow"]
CMD ["--help"]