import uuid
from typing import List, Optional
from pydantic import BaseModel


class Category(BaseModel):

    category_name: str


class Categories(BaseModel):

    list_categories: List[Category]


class Note(BaseModel):

    user_uuid: uuid.UUID
    category_name: Optional[str] = None
    title: str
    body: str
    update_at: Optional[float] = None
    created_at: Optional[float] = None


# class NoteResponse(BaseModel):

#     note_uuid: str

class ListNotes(BaseModel):

    notes: List[Note]


class ListNotesByCategory(BaseModel):

    category: Category
    list_notes: ListNotes


class NotesByCategory(BaseModel):

    category_id: int