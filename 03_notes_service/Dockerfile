FROM python:3.10-slim

WORKDIR /usr/local/notes_service/


ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1


RUN apt-get update
RUN apt-get install -y python3-dev gcc libc-dev libffi-dev
RUN apt-get -y install libpq-dev gcc

COPY requirements.txt /usr/local/notes_service/
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt

COPY . /usr/local/notes_service/

ENV PYTHONPATH "${PYTHONPATH}:/usr/local/notes_service/"
