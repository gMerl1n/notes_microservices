from pydantic import BaseModel


class NoteCreateRequest(BaseModel):
    user_id: int
    category_name: str
    title: str
    body: str


class NoteGetRequestById(BaseModel):
    note_id: int
    user_id: int


class NotesGetRequest(BaseModel):
    user_id: int


class NoteRemoveRequestById(BaseModel):
    note_id: int
    user_id: int


class NotesRemoveRequest(BaseModel):
    user_id: int


class CategoryCreateRequest(BaseModel):
    category_name: str
    user_id: int


class CategoryGetRequestById(BaseModel):
    category_id: int
    user_id: int


class CategoriesGetRequest(BaseModel):
    user_id: int


class CategoryRemoveRequestById(BaseModel):
    category_id: int
    user_id: int
