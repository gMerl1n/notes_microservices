from fastapi import Depends, HTTPException, status
from fastapi.security import OAuth2PasswordBearer
from jose import JWTError, jwt
from app.settings import settings
# from app.clients.auth.schema import User


oauth2_scheme = OAuth2PasswordBearer(tokenUrl="/login/token")


async def get_user_uuid_from_token(token: str = Depends(oauth2_scheme)) -> str:

    credentials_exception = HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Could not validate credentials",
    )

    try:
        payload = jwt.decode(
            token, settings.SECRET_KEY, algorithms=[settings.ALGORITHM]
        )
        userUUID: str = payload.get("sub")
        print("userUUID/extracted is ", userUUID)
        if userUUID is None:
            raise credentials_exception
    except JWTError:
        raise credentials_exception
    
    return userUUID