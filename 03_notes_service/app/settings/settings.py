import os
from dotenv import load_dotenv
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine

load_dotenv()


POSTGRES_HOST = os.environ.get("POSTGRES_HOST")
POSTGRES_PORT = os.environ.get("POSTGRES_PORT")
POSTGRES_DB = os.environ.get("POSTGRES_DB")
POSTGRES_USER = os.environ.get("POSTGRES_USER")
POSTGRES_PASSWORD = os.environ.get("POSTGRES_PASSWORD")

DATABASE_URL_POSTGRES = f"postgresql+asyncpg://{POSTGRES_USER}:{POSTGRES_PASSWORD}@{POSTGRES_HOST}:{POSTGRES_PORT}/{POSTGRES_DB}?async_fallback=True"

engine = create_async_engine(DATABASE_URL_POSTGRES, echo=False, future=True)
async_session = sessionmaker(autoflush=False, bind=engine, class_=AsyncSession)


NOTES_GRPC_SERVER_ADDR = "0.0.0.0:50052"