FROM alpine:latest

RUN apk --update add nmap \
                     nmap-scripts \
                     git \
                     bash

RUN git clone https://github.com/scipag/vulscan.git /usr/share/nmap/scripts/vulscan
RUN git clone https://github.com/vulnersCom/nmap-vulners.git /usr/share/nmap/scripts/vulners
WORKDIR /usr/share/nmap/scripts/vulscan

#Update CVE databases
CMD ["/bin/bash","updateFiles.sh"]
WORKDIR /usr/share/nmap/scripts
ENTRYPOINT ["nmap"]
CMD ["-h"]

