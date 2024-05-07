FROM python:3.10-slim

WORKDIR /api_gateway


ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1
ENV PYTHONPATH "${PYTHONPATH}:/api_gateway"

RUN apt-get update
RUN apt-get install -y python3-dev gcc libc-dev libffi-dev
RUN apt-get -y install libpq-dev gcc

COPY requirements.txt .
RUN pip install --upgrade pip
RUN pip install --no-cache-dir -r requirements.txt

COPY . .