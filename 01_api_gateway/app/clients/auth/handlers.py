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
        #print(auth_input.dict())
        auth_token = requests.post(url=settings.LOGIN_URL, json=auth_input)
        print(auth_token.json())
    except Exception as ex:
        raise HTTPException(status_code=500, detail=f"Internal server error {ex}")

    return auth_token.json()

    # user = await authenticate_user(form_data.username, form_data.password)
    # if not user:
    #     raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="Incorrect username or password",)

    # # Если пользователь найден по почте, то создаем для него токен
    # access_token_expires = timedelta(minutes=settings.ACCESS_TOKEN_EXPIRE_MINUTES)

    # access_token = create_access_token(
    #     data={"sub": user.email, "other_custom_data": [1, 2, 3, 4]},
    #     expires_delta=access_token_expires,
    # )

    # return {"access_token": access_token, "token_type": "bearer"}

# @router_auth.get("/")
# async def test_for_auth(userUUID = Depends(get_userUUID_from_token)):

#     """Получить информацию о пользователе по его id"""

#     return JSONResponse({"userUUID": userUUID})