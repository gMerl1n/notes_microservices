from fastapi import APIRouter, HTTPException
from fastapi.responses import JSONResponse
from app.config.config import config
import requests
import json


router_auth = APIRouter()


@router_auth.get("/login")
async def login():
    """Авторизация пользователя"""

    try:
        login_url = config.get_login_url()
        auth_token = requests.get(url=login_url, data=json.dumps({"email": "123", "password_hash": "123"}))
    except Exception as ex:
        raise HTTPException(status_code=500, detail="Internal server error")

    return JSONResponse({"token": auth_token.json()})