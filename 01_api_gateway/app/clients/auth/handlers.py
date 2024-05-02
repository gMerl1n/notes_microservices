from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import JSONResponse
from app.clients.auth.service import get_userUUID_from_token
from app.settings import settings
import requests
import json


router_auth = APIRouter()


@router_auth.get("/login")
async def login():
    """Авторизация пользователя"""

    try:
        auth_token = requests.get(url=settings.LOGIN_URL, data=json.dumps({"email": "123", "password_hash": "123"}))
    except Exception as ex:
        raise HTTPException(status_code=500, detail="Internal server error")

    return JSONResponse({"token": auth_token.json()})



@router_auth.get("/")
async def test_for_auth(userUUID = Depends(get_userUUID_from_token)):

    """Получить информацию о пользователе по его id"""

    return JSONResponse({"userUUID": userUUID})