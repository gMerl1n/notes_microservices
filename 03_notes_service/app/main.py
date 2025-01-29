import logging
import uvicorn
from settings.settings import settings
from fastapi import FastAPI



# Remove all handlers associated with the root logger object.
for handler in logging.root.handlers[:]:
    logging.root.removeHandler(handler)


logging.basicConfig(
    format='%(asctime)s - %(message)s | %(levelname)s ',
    datefmt='%d-%b-%y %H:%M:%S',
    level=logging.INFO
)


app = FastAPI()


if __name__ == "__main__":
    logging.info(f'Start server {settings.server_config.port}')
    uvicorn.run("main:app",
                host="0.0.0.0",
                port=settings.server_config.port,
                log_level=settings.server_config.log_level)
