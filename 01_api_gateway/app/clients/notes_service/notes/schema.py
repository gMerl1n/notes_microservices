import uuid
from typing import List, Optional
from pydantic import BaseModel


# Notes pydantic models

class Note(BaseModel):

    category_name: Optional[str] = None
    title: str
    body: str
    update_at: Optional[float] = None
    created_at: Optional[float] = None


class NoteCreate(Note):
    pass


class NoteGet(BaseModel):

    note_uuid: str


class ListNotesGet(BaseModel):

    user_uuid: str


class ListNotesByCategory(BaseModel):

    category: str
    list_notes: List[Note]


class NotesByCategory(BaseModel):

    user_uuid: str
    category_id: int


class NoteUpdate(BaseModel):

    user_uuid: str
    note_uuid: str
    category_name: Optional[str] = None
    title: Optional[str] = None
    body: Optional[str] = None


class NoteDelete(BaseModel):

    note_uuid: str


# Category pydantic models
    

class Category(BaseModel):

    category_name: str


class Categories(BaseModel):

    list_categories: List[Category]


class CategoryCreate(BaseModel):

    user_uuid: str
    category_name: str


class CategoryGet(BaseModel):

    user_uuid: str