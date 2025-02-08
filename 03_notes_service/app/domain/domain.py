from dataclasses import dataclass, field
from typing import Optional


@dataclass
class NoteEntity:
    id: Optional[int] = field(init=False, default=None)
    category_id: int
    user_id: int
    title: str
    body: str
    update_at: int
    created_at: int

    def to_dict(self):
        return self.__dict__

    @classmethod
    def to_model(cls, dict_obj):
        return cls(**dict_obj)


@dataclass
class CategoryEntity:
    id: Optional[int] = field(init=False, default=None)
    category_name: str
    user_id: str
    update_at: int
    created_at: int

    def to_dict(self):
        return self.__dict__

    @classmethod
    def to_model(cls, dict_obj):
        return cls(**dict_obj)