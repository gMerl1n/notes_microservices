from pydantic import BaseModel


class AuthUser(BaseModel):

    email: str
    password: str


class Token(BaseModel):

    AccessToken: str
    RefreshToken: str