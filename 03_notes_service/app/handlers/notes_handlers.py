from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import JSONResponse
from services.notes_services import INoteService
from handlers.schema_request import NoteCreateRequest
from container.container import container
from sqlalchemy.ext.asyncio import AsyncSession
from settings.async_session import get_async_session

router_notes = APIRouter()


@router_notes.get("/get_note")
async def get_note_by_id(note_id: int,
                         async_session: AsyncSession = Depends(get_async_session),
                         notes_service: INoteService = Depends(container.get_notes_service)) -> JSONResponse:

    note = await notes_service.get_note_by_id(async_session=async_session, note_id=note_id)
    if note is None:
        raise HTTPException(status_code=400, detail="Something wrong")

    return JSONResponse(content=note, status_code=200)


@router_notes.post("/create_note")
async def create_note(note: NoteCreateRequest,
                      async_session: AsyncSession = Depends(get_async_session),
                      notes_service: INoteService = Depends(container.get_notes_service)) -> JSONResponse:

    note_id = await notes_service.save_note(async_session=async_session, note=note.model_dump())
    if note_id is None:
        raise HTTPException(status_code=500, detail="Something wrong")

    return JSONResponse(content=note_id, status_code=201)
