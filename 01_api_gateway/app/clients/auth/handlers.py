from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import JSONResponse
from app.settings import settings
from .schema import Token, AuthUser
import requests
import json
from fastapi.security import OAuth2PasswordBearer
from fastapi.security import OAuth2PasswordRequestForm


router_auth = APIRouter()


# @router_auth.get("/login", response_model=Token)
# async def login():
#     """Авторизация пользователя"""

#     try:
#         auth_token = requests.get(url=settings.LOGIN_URL, data=json.dumps({"email": "123", "password_hash": "123"}))
#     except Exception as ex:
#         raise HTTPException(status_code=500, detail="Internal server error")

#     return JSONResponse({"token": auth_token.json()})


oauth2_scheme = OAuth2PasswordBearer(tokenUrl="/login/token")



@router_auth.post(f"/token", response_model=Token)
async def login_for_access_token(form_data: OAuth2PasswordRequestForm = Depends()):

    auth_input = {"email": form_data.username, "password": form_data.password}

    try:
        auth_token = requests.post(url=settings.LOGIN_URL, json=auth_input)
        print(auth_token.json())
    except Exception as ex:
        raise HTTPException(status_code=500, detail=f"Internal server error {ex}")

    return auth_token.json()


@router_auth.post("/register")
async def register_user():
    pass


@router_auth.get("/refreshtokens")
async def refresh_tokens():
    pass