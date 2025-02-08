import json
import os
from pydantic import BaseModel
from pathlib import Path
from dotenv import load_dotenv
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine

load_dotenv()

BASE_DIR = Path(__file__).resolve().parent.parent

POSTGRES_HOST = os.environ.get("POSTGRES_HOST")
POSTGRES_PORT = os.environ.get("POSTGRES_PORT")
POSTGRES_DB = os.environ.get("POSTGRES_DB")
POSTGRES_USER = os.environ.get("POSTGRES_USER")
POSTGRES_PASSWORD = os.environ.get("POSTGRES_PASSWORD")

DATABASE_URL_POSTGRES = f"postgresql+asyncpg://{POSTGRES_USER}:{POSTGRES_PASSWORD}@{POSTGRES_HOST}:{POSTGRES_PORT}/{POSTGRES_DB}?async_fallback=True"

engine = create_async_engine(DATABASE_URL_POSTGRES, echo=False, future=True)
SessionLocal = sessionmaker(autoflush=False, bind=engine, class_=AsyncSession)

config_path = os.path.join(BASE_DIR, 'settings', "config.json")

with open(config_path) as file:
    config = json.load(file)
    server_config = config["server"]


class ServerConfig(BaseModel):
    port: int
    log_level: str


class Settings:
    server_config = ServerConfig(port=server_config["port"], log_level=server_config["log_level"])


config = Settings()
