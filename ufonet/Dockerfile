FROM python:2
RUN apt-get update && apt-get install -y python-pycurl python-geoip python-whois python-crypto python-requests openssl
RUN pip install geoip requests pycrypto
RUN git clone https://github.com/epsylon/ufonet.git
WORKDIR /ufonet
RUN python setup.py install
EXPOSE 80 443
ENTRYPOINT ["python", "ufonet"]
CMD ["--help"]
