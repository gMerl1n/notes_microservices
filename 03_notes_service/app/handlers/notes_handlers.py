from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import JSONResponse
from services.notes_services import INoteService
from handlers.schema_request import NoteCreateRequest
from container import container
from sqlalchemy.ext.asyncio import AsyncSession
from settings.async_session import get_async_session


router_notes = APIRouter()


@router_notes.get("/get_notes")
async def get_note_by_id(id_note: int,
                         async_session: AsyncSession = Depends(get_async_session),
                         notes_service: INoteService = Depends(container.get_notes_service)):
    pass


@router_notes.post("/create_note")
async def create_note(note: NoteCreateRequest,
                      async_session: AsyncSession = Depends(get_async_session),
                      notes_service: INoteService = Depends(container.get_notes_service)):
    pass