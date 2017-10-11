#DOES NOT WORK
#todo
FROM python:2.7-alpine

WORKDIR /root

RUN apk --update add git py-lxml python-dev
RUN git clone https://github.com/DataSploit/datasploit.git datasploit

WORKDIR datasploit
RUN apk add --update --no-cache g++ gcc libxslt-dev libxml2-dev
RUN pip install -r requirements.txt
CMD python datasploit_config.py
#CMD ["mv", "config_sample.py", "config.py"]
ENTRYPOINT ["python", "datasploit.py"]

#CMD ["--help"]