FROM python:2.7

RUN apt-get update && apt-get install -y \
        git


RUN git clone https://github.com/anilbaranyelken/tulpar.git

WORKDIR tulpar

RUN pip install --requirement requirements

ENTRYPOINT ["python", "tulpar.py"]
CMD ["-h"]