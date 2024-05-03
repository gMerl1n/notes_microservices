from grpc import aio
from app.protos.genproto import notes_pb2_grpc
from app.notes.handlers import NoteService




async def run_server(address):
    # Здесь получаем асинхронный сервер
    server = aio.server()
    print('START SERVER')
    # Регистрируем наш Todo сервер в aio сервере
    notes_pb2_grpc.add_NoteServicer_to_server(NoteService(), server)
    # Теперь этот сервер необходимо зарегистрировать по какому-то адресу
    server.add_insecure_port(address)
    print('START SERVER')
    await server.start()
    await server.wait_for_termination()
