from app.protos.genproto import notes_pb2_grpc, notes_pb2
from google.protobuf.json_format import MessageToDict
from .schema import Note
from .services import create_note
from app.settings.async_session import get_async_session


class NoteService(notes_pb2_grpc.NoteServicer):
    
    async def CreateNote(self, request, context):

        note = MessageToDict(request, preserving_proto_field_name=True)
        pydantic_note = Note.model_validate(note)

        note_uuid = await create_note(note=pydantic_note, async_session=get_async_session())

        return notes_pb2.CreateNoteResponse(noteUUID=note_uuid)
    

