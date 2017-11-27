FROM python:2.7-slim

RUN apt-get update && apt-get install -y git \
										python2.7-dev  \
										python-pip \
										python-docutils \
										perl \
										nmap \
										sslscan

WORKDIR /opt
RUN git clone https://github.com/golismero/golismero.git
COPY user.conf ~/.golismero/user.conf
WORKDIR golismero/
RUN pip install -r requirements.txt && pip install -r requirements_unix.txt && ln -s /opt/golismero/golismero.py /usr/bin/golismero
ENTRYPOINT ["python", "golismero.py"]
CMD ["--help"]