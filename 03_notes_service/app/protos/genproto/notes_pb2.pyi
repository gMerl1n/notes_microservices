from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateNoteRequest(_message.Message):
    __slots__ = ("user_uuid", "category_name", "title", "body")
    USER_UUID_FIELD_NUMBER: _ClassVar[int]
    CATEGORY_NAME_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    user_uuid: str
    category_name: str
    title: str
    body: str
    def __init__(self, user_uuid: _Optional[str] = ..., category_name: _Optional[str] = ..., title: _Optional[str] = ..., body: _Optional[str] = ...) -> None: ...

class CreateNoteResponse(_message.Message):
    __slots__ = ("note_uuid",)
    NOTE_UUID_FIELD_NUMBER: _ClassVar[int]
    note_uuid: str
    def __init__(self, note_uuid: _Optional[str] = ...) -> None: ...

class GetNoteRequest(_message.Message):
    __slots__ = ("note_uuid",)
    NOTE_UUID_FIELD_NUMBER: _ClassVar[int]
    note_uuid: str
    def __init__(self, note_uuid: _Optional[str] = ...) -> None: ...

class GetNoteResponse(_message.Message):
    __slots__ = ("note_uuid", "title", "body", "update_at", "created_at")
    NOTE_UUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    UPDATE_AT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    note_uuid: str
    title: str
    body: str
    update_at: float
    created_at: float
    def __init__(self, note_uuid: _Optional[str] = ..., title: _Optional[str] = ..., body: _Optional[str] = ..., update_at: _Optional[float] = ..., created_at: _Optional[float] = ...) -> None: ...

class GetListNotesRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetListNotesResponse(_message.Message):
    __slots__ = ("notes",)
    NOTES_FIELD_NUMBER: _ClassVar[int]
    notes: _containers.RepeatedCompositeFieldContainer[GetNoteResponse]
    def __init__(self, notes: _Optional[_Iterable[_Union[GetNoteResponse, _Mapping]]] = ...) -> None: ...

class CreateCategoryRequest(_message.Message):
    __slots__ = ("category_name",)
    CATEGORY_NAME_FIELD_NUMBER: _ClassVar[int]
    category_name: str
    def __init__(self, category_name: _Optional[str] = ...) -> None: ...

class CreateCategoryResponse(_message.Message):
    __slots__ = ("category_id",)
    CATEGORY_ID_FIELD_NUMBER: _ClassVar[int]
    category_id: int
    def __init__(self, category_id: _Optional[int] = ...) -> None: ...

class GetNotesByCategoryRequest(_message.Message):
    __slots__ = ("category_id",)
    CATEGORY_ID_FIELD_NUMBER: _ClassVar[int]
    category_id: int
    def __init__(self, category_id: _Optional[int] = ...) -> None: ...

class GetNotesByCategoryResponse(_message.Message):
    __slots__ = ("category_name", "notes")
    CATEGORY_NAME_FIELD_NUMBER: _ClassVar[int]
    NOTES_FIELD_NUMBER: _ClassVar[int]
    category_name: str
    notes: _containers.RepeatedCompositeFieldContainer[GetNoteResponse]
    def __init__(self, category_name: _Optional[str] = ..., notes: _Optional[_Iterable[_Union[GetNoteResponse, _Mapping]]] = ...) -> None: ...

class UpdateNoteRequest(_message.Message):
    __slots__ = ("note_uuid", "title", "body")
    NOTE_UUID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    BODY_FIELD_NUMBER: _ClassVar[int]
    note_uuid: str
    title: str
    body: str
    def __init__(self, note_uuid: _Optional[str] = ..., title: _Optional[str] = ..., body: _Optional[str] = ...) -> None: ...

class UpdateNoteResponse(_message.Message):
    __slots__ = ("note_uuid",)
    NOTE_UUID_FIELD_NUMBER: _ClassVar[int]
    note_uuid: str
    def __init__(self, note_uuid: _Optional[str] = ...) -> None: ...

class DeleteNoteRequest(_message.Message):
    __slots__ = ("note_uuid",)
    NOTE_UUID_FIELD_NUMBER: _ClassVar[int]
    note_uuid: str
    def __init__(self, note_uuid: _Optional[str] = ...) -> None: ...

class DeleteNoteReponse(_message.Message):
    __slots__ = ("note_uuid",)
    NOTE_UUID_FIELD_NUMBER: _ClassVar[int]
    note_uuid: str
    def __init__(self, note_uuid: _Optional[str] = ...) -> None: ...
