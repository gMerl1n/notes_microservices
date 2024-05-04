import asyncio
from app.settings import settings
from app.server import notes_server


if __name__ == '__main__':
    asyncio.run(notes_server.run_server(settings.NOTES_GRPC_SERVER_ADDR))


import sys

for i in sys.path:
    print(i)