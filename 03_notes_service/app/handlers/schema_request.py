from pydantic import BaseModel


class NoteCreateRequest(BaseModel):
    user_id: int
    role_id: int
    category_name: str
    title: str
    body: str
