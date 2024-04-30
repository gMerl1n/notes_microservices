from fastapi import APIRouter, HTTPException
from fastapi.responses import JSONResponse
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