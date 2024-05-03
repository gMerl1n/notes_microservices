from typing import List, Optional
from pydantic import BaseModel


class Category(BaseModel):

    name_category: str


class Categories(BaseModel):

    list_categories: List[Category]


class Note(BaseModel):

    note_uuid: str
    category_id: Optional[int]
    title: str
    body: str
    update_at: Optional[float]
    created_at: float


class ListNotes(BaseModel):

    notes: List[Note]


class ListNotesByCategory(BaseModel):

    category: Category
    list_notes: ListNotes
