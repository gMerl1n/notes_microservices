from app.protos.genproto import notes_pb2_grpc, notes_pb2
from google.protobuf.json_format import MessageToDict
from .schema import Note, Category
from .services import create_note, create_category, get_list_notes, get_note_by_id
from app.settings import settings
from app.logger.logg import init_logger
from datetime import datetime


logger = init_logger(name="handlers")


class NoteService(notes_pb2_grpc.NoteServicer):

    
    async def CreateNote(self, request, context):
        
        logger.info("CreateNote: request from api has received")
        note = MessageToDict(request, preserving_proto_field_name=True)
        serialized_note = Note.model_validate(note)
        
        try:
            note_uuid = await create_note(note=serialized_note, async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"CreateNote: note has not been added. Error: {str(ex)}")

        #logger.info("CreateNote: response with note uuid from note service has sent")
        return notes_pb2.CreateNoteResponse(note_uuid=str(note_uuid))
    

    async def CreateCategory(self, request, context):

        logger.info("CreateCategory: request from api has received")
        category = MessageToDict(request, preserving_proto_field_name=True)

        serialized_category = Category.model_validate(category)

        try:
            category_id = await create_category(category=serialized_category, async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"CreateCategory: category has not been added. Error: {str(ex)}")

        return notes_pb2.CreateCategoryResponse(category_id=category_id)
    

    async def GetNotes(self, request, context):

        logger.info("GetNotes: request from api has received")
        
        try:
            notes = await get_list_notes(async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"GetNotes: notes has not been added. Error: {str(ex)}")
        return notes_pb2.GetListNotesResponse(notes=notes)


    async def GetNote(self, request, context):
        
        note = MessageToDict(request, preserving_proto_field_name=True)
        try:
            note = await get_note_by_id(uuid=note["note_uuid"], async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"GetNote: note has not been added. Error: {str(ex)}")
        return notes_pb2.GetNoteResponse(note_uuid=str(note.note_uuid), 
                                         title=note.title,
                                         body=note.body,
                                         update_at=datetime.timestamp(note.update_at) if note.update_at is not None else None,
                                         created_at=datetime.timestamp(note.created_at)
        )