from pydantic import typing
from fastapi import APIRouter, Depends, HTTPException
from grpc.aio import AioRpcError
from app.clients.notes_service.protos.genproto import notes_pb2
from app.clients.notes_service.client_grpc import grpc_notes_service_client
from app.clients.auth.service import get_user_uuid_from_token
from google.protobuf.json_format import MessageToDict
from fastapi.responses import JSONResponse
from .schema import (
    NoteCreate, 
    Category, 
    NotesByCategory, 
    NoteUpdate, 
    ListNotesGet, 
    CategoryCreate, 
    CategoryGet
)


router_notes = APIRouter()


@router_notes.post("/")
async def create_note(new_note: NoteCreate, 
                      user_uuid_from_token = Depends(get_user_uuid_from_token),
                      client: typing.Any = Depends(grpc_notes_service_client)):
    
    try:
        data = new_note.dict(exclude_none=True)
        response = await client.CreateNote(notes_pb2.CreateNoteRequest(user_uuid=user_uuid_from_token,**data))
    except AioRpcError as ex:
        raise HTTPException(status_code=500, detail=f"error: {ex.details()}")
    
    return JSONResponse(MessageToDict(response, preserving_proto_field_name=True))


@router_notes.get("/{uuid}")
async def get_note(note_uuid: str, 
                   user_uuid_from_token = Depends(get_user_uuid_from_token),
                   client: typing.Any = Depends(grpc_notes_service_client)):
    
    try:
        # data = new_note.dict(exclude_none=True)
        response = await client.GetNote(notes_pb2.GetNoteRequest(user_uuid=user_uuid_from_token, note_uuid=note_uuid))
    except AioRpcError as ex:
        raise HTTPException(status_code=500, detail=f"error: {ex.details()}")
    
    return JSONResponse(MessageToDict(response, preserving_proto_field_name=True))


@router_notes.get("/listnotes")
async def get_list_notes(user_uuid_from_token = Depends(get_user_uuid_from_token),
                         client: typing.Any = Depends(grpc_notes_service_client)):
    
    try:
        # data = new_note.dict(exclude_none=True)
        response = await client.GetNotes(notes_pb2.GetListNotesRequest(user_uuid=user_uuid_from_token))
    except AioRpcError as ex:
        raise HTTPException(status_code=500, detail=f"error: {ex.details()}")
    
    return JSONResponse(MessageToDict(response, preserving_proto_field_name=True))


@router_notes.get("/notesbycategory/{uuid}")
async def get_list_notes_by_category(category_id: int, 
                                     user_uuid_from_token = Depends(get_user_uuid_from_token),
                                     client: typing.Any = Depends(grpc_notes_service_client)):
    
    try:
        # data = new_note.dict(exclude_none=True)
        response = await client.GetNotesByCategory(notes_pb2.GetNotesByCategoryRequest(user_uuid=user_uuid_from_token, category_id=category_id))
    except AioRpcError as ex:
        raise HTTPException(status_code=500, detail=f"error: {ex.details()}")
    
    return JSONResponse(MessageToDict(response, preserving_proto_field_name=True))


@router_notes.patch("/updatenote/{uuid}")
async def update_note(note_to_update: NoteUpdate, 
                      user_uuid_from_token = Depends(get_user_uuid_from_token),
                      client: typing.Any = Depends(grpc_notes_service_client)):
    try:
        data = note_to_update.dict(exclude_none=True)
        response = await client.UpdateNote(notes_pb2.UpdateNoteRequest(user_uuid=user_uuid_from_token, note_to_update=data))
    except AioRpcError as ex:
        raise HTTPException(status_code=500, detail=f"error: {ex.details()}")
    
    return JSONResponse(MessageToDict(response, preserving_proto_field_name=True))


@router_notes.delete("/{uuid}")
async def delete_note(note_uuid: str, 
                      user = Depends(get_user_uuid_from_token),
                      client: typing.Any = Depends(grpc_notes_service_client)):
    pass
