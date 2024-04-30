from fastapi import FastAPI
from app.routes import routes
from app.config.config import ConfigLoader
from contextlib import asynccontextmanager


app = FastAPI()
# app.include_router(routes.routes)

config = ConfigLoader()


app.include_router(routes.routes)