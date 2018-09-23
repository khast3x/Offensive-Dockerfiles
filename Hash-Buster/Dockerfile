FROM python:3-alpine

LABEL name CMSeeK
LABEL src "https://github.com/s0md3v/Hash-Buster"
LABEL creator s0md3v
LABEL dockerfile_maintenance khast3x
LABEL desc "Crack hashes in seconds"

RUN apk add --no-cache git py3-pip && git clone https://github.com/s0md3v/Hash-Buster.git
WORKDIR Hash-Buster

# Lighter alternative to installing make to run makefile
COPY hash.py /usr/local/bin/
RUN chmod +x /usr/local/bin/hash.py

RUN pip3 install requests

ENTRYPOINT [ "hash.py" ]
CMD [ "-h" ]