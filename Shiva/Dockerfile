FROM alpine:latest

RUN apk add --update python py-pip git

RUN git clone https://github.com/UltimateHackers/Shiva.git
WORKDIR Shiva
RUN pip install -r requirements.txt
ENTRYPOINT ["python", "shiva.py"]
