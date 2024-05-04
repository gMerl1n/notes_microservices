import os
import logging
from app.settings import settings



def init_logger(name: str, level: int = logging.INFO):

    if isinstance(level, str):
        level = getattr(logging, level)

    logger = logging.getLogger(name)
    logger.setLevel(level)
    console_log = logging.StreamHandler()

    logs_path = os.path.join(settings.BASE_DIR, "logger", "logs.log")

    file_log = logging.FileHandler(filename=logs_path, mode='a')
    console_log.setFormatter(logging.Formatter(
        '%(asctime)s, — %(levelname)s — module: %(name)s — %(message)s'))
    file_log.setFormatter(logging.Formatter(
        '%(asctime)s, — %(levelname)s — module: %(name)s — %(message)s'))
    

    logger.addHandler(console_log)
    logger.addHandler(file_log)

    return logger