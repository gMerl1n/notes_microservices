import grpc
from app.settings import settings
from app.clients.notes_service.protos.genproto import notes_pb2_grpc


async def grpc_notes_service_client():
    # Открываем канал и указываем, на каком порту
    channel = grpc.aio.insecure_channel(settings.NOTES_GRPC_SERVER_ADDR)
    # Создаем и возвращаем клиент
    client = notes_pb2_grpc.NoteStub(channel)
    return client