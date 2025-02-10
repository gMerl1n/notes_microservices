from pydantic import BaseModel


class NoteCreateRequest(BaseModel):
    user_id: int
    category_name: str
    title: str
    body: str


class CategoryCreateRequest(BaseModel):
    category_name: str
    user_id: int