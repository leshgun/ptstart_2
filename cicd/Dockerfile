FROM python:3.12.4-alpine

ENV LISTEN_IP=0.0.0.0
ENV LISTEN_PORT=11111

COPY app /app
WORKDIR /app

RUN python3 --version
# RUN pip3 install flask
RUN pip install -r requirements.txt

CMD [ "sh", "/app/start.sh" ]