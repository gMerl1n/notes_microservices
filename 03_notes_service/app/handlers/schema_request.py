from pydantic import BaseModel


class NoteCreateRequest(BaseModel):
    user_id: int
    category_name: str
    title: str
    body: str


class NoteGetRequest(BaseModel):
    note_id: int
    user_id: int

class CategoryCreateRequest(BaseModel):
    category_name: str
    user_id: int


class CategoryGetRequest(BaseModel):
    category_id: int
    user_id: int


class CategoryRemoveRequest(BaseModel):
    category_id: int
    user_id: int
