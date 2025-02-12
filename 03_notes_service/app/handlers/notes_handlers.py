from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import JSONResponse
from sqlalchemy.ext.asyncio import AsyncSession

from services.notes_services import INoteService
from handlers.schema_request import (
    NoteCreateRequest,
    NoteGetRequestById,
    NoteRemoveRequestById,
    NotesGetRequest,
    NotesRemoveRequest
)
from container.container import container
from settings.async_session import get_async_session

router_notes = APIRouter()


@router_notes.post("/create_note")
async def create_note(note: NoteCreateRequest,
                      async_session: AsyncSession = Depends(get_async_session),
                      notes_service: INoteService = Depends(container.get_notes_service)) -> JSONResponse:

    note_id = await notes_service.save_note(async_session=async_session,
                                            title=note.title,
                                            body=note.body,
                                            category_name=note.category_name,
                                            user_id=note.user_id)
    if note_id is None:
        raise HTTPException(status_code=500, detail="Something wrong")

    return JSONResponse(content=note_id, status_code=201)


@router_notes.post("/get_note")
async def get_note_by_id(note_get_request: NoteGetRequestById,
                         async_session: AsyncSession = Depends(get_async_session),
                         notes_service: INoteService = Depends(container.get_notes_service)) -> JSONResponse:

    note = await notes_service.get_note_by_id(async_session=async_session,
                                              note_id=note_get_request.note_id,
                                              user_id=note_get_request.user_id)
    if note is None:
        raise HTTPException(status_code=400, detail="Something wrong")

    return JSONResponse(content=note.to_dict(), status_code=200)


@router_notes.post("/get_notes")
async def get_notes(notes_get_request: NotesGetRequest,
                    async_session: AsyncSession = Depends(get_async_session),
                    notes_service: INoteService = Depends(container.get_notes_service)) -> JSONResponse:

    notes = await notes_service.get_all_notes(async_session=async_session, user_id=notes_get_request.user_id)
    if notes is None:
        raise HTTPException(status_code=400, detail="Not found")

    return JSONResponse(content=[n.to_dict() for n in notes], status_code=200)


@router_notes.delete("/remove_note_by_id")
async def remove_note_by_id(note_remove_by_id_request: NoteRemoveRequestById,
                            async_session: AsyncSession = Depends(get_async_session),
                            notes_service: INoteService = Depends(container.get_notes_service)) -> JSONResponse:

    removed_note_id = await notes_service.remove_note_by_id(async_session=async_session,
                                                            note_id=note_remove_by_id_request.note_id,
                                                            user_id=note_remove_by_id_request.user_id)
    if removed_note_id is None:
        raise HTTPException(status_code=400, detail="Not found")

    return JSONResponse(content=removed_note_id, status_code=200)


@router_notes.delete("/remove_notes")
async def remove_notes(notes_remove_request: NotesRemoveRequest,
                       async_session: AsyncSession = Depends(get_async_session),
                       notes_service: INoteService = Depends(container.get_notes_service)) -> JSONResponse:

    removed_notes_ids = await notes_service.remove_all_notes(async_session=async_session,
                                                             user_id=notes_remove_request.user_id)
    if removed_notes_ids is None:
        raise HTTPException(status_code=500, detail="Something went wrong")

    return JSONResponse(content=removed_notes_ids, status_code=200)
