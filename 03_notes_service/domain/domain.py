from dataclasses import dataclass


@dataclass
class Note:
    note_id: int
    category_id: int
    user_id: int
    title: str
    body: str
    update_at: float
    created_at: float


@dataclass
class Category:
    category_id: int
    category_name: str
    user_id: str
    update_at: float
    created_at: float