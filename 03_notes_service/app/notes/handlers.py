from app.protos.genproto import notes_pb2_grpc, notes_pb2
from google.protobuf.json_format import MessageToDict
from .schema import Note, Category, NotesByCategory, NoteUpdate
from .services import create_note, create_category, get_list_notes, get_note_by_id, get_notes_by_category, update_note
from app.settings import settings
from app.logger.logg import init_logger
from datetime import datetime


logger = init_logger(name="handlers")


# rpc CreateNote (CreateNoteRequest) returns (CreateNoteResponse);
#     rpc GetNote (GetNoteRequest) returns (GetNoteResponse);
#     rpc GetNotes (GetListNotesRequest) returns (GetListNotesResponse);
#     rpc CreateCategory (CreateCategoryRequest) returns (CreateCategoryResponse);
#     rpc GetNotesByCategory (GetNotesByCategoryRequest) returns (GetNotesByCategoryResponse);
#     rpc UpdateNote (UpdateNoteRequest) returns (UpdateNoteResponse);
#     rpc DeleteNote (DeleteNoteRequest) returns (DeleteNoteReponse);
# }


class NoteService(notes_pb2_grpc.NoteServicer):

    # rpc CreateNote (CreateNoteRequest) returns (CreateNoteResponse);    
    async def CreateNote(self, request, context):
        
        logger.info("CreateNote: request from api has received")
        note = MessageToDict(request, preserving_proto_field_name=True)
        serialized_note = Note.model_validate(note)
        
        try:
            note_uuid = await create_note(note=serialized_note, async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"CreateNote: note has not been added. Error: {str(ex)}")

        return notes_pb2.CreateNoteResponse(note_uuid=str(note_uuid))
    
    # rpc GetNote (GetNoteRequest) returns (GetNoteResponse);
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
    
    # rpc GetNotes (GetListNotesRequest) returns (GetListNotesResponse);
    async def GetNotes(self, request, context):

        logger.info("GetNotes: request from api has received")
        
        try:
            notes = await get_list_notes(async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"GetNotes: notes has not been added. Error: {str(ex)}")
        return notes_pb2.GetListNotesResponse(notes=notes)


    # rpc CreateCategory (CreateCategoryRequest) returns (CreateCategoryResponse);
    async def CreateCategory(self, request, context):

        logger.info("CreateCategory: request from api has received")
        category = MessageToDict(request, preserving_proto_field_name=True)

        serialized_category = Category.model_validate(category)

        try:
            category_id = await create_category(category=serialized_category, async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"CreateCategory: category has not been added. Error: {str(ex)}")

        return notes_pb2.CreateCategoryResponse(category_id=category_id)
    

    #rpc GetNotesByCategory (GetNotesByCategoryRequest) returns (GetNotesByCategoryResponse);
    async def GetNotesByCategory(self, request, context):

        logger.info("GetNotesByCategory: request from api has received")
        
        category = MessageToDict(request, preserving_proto_field_name=True)
        serialized_category = NotesByCategory.model_validate(category)

        try:
            notes_by_category, count_notes = await get_notes_by_category(category_id=serialized_category, async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"GetNotesByCategory: notes has not been added. Error: {str(ex)}")
        
        return notes_pb2.GetNotesByCategoryResponse(count_notes_by_cat=count_notes, notes=notes_by_category) 


    #rpc UpdateNote (UpdateNoteRequest) returns (UpdateNoteResponse);
    async def UpdateNote(self, request, context):
        
        logger.info("UpdateNote: request from api has received")

        data_to_update = MessageToDict(request, preserving_proto_field_name=True)
        serialized_category = NoteUpdate.model_validate(data_to_update)

        try:
            note_uuid = await update_note(note_uuid=data_to_update["note_uuid"],
                                    params_to_update=serialized_category,
                                    async_session=settings.async_session)
        except Exception as ex:
            logger.warning(f"UpdateNote: notes has not been updated. Error: {str(ex)}")

        return notes_pb2.UpdateNoteResponse(note_uuid=note_uuid) 


    #rpc DeleteNote (DeleteNoteRequest) returns (DeleteNoteReponse);
    async def DeleteNote(self, request, context):
        pass