FROM python:2-alpine

RUN apk --update add git

RUN git clone https://github.com/Hadesy2k/sqliv.git
WORKDIR /sqliv
RUN pip install -r requirements.txt && python setup.py -i

ENTRYPOINT ["python","sqliv.py"]
CMD ["--help"]