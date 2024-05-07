



# ================== Auth-service api ===============

LOGIN_URL = "http://localhost:10001/api/auth/login"
REGISTER = "http://localhost:10001/api/auth/register"
REFRESHTOKENS = "http://localhost:10001/api/auth/refreshtokens"


# =================== JWT =========================


SECRET_KEY: str = "$3cr3t"
ALGORITHM: str = "HS256"


# =================== Notes-service gRPC ===================

NOTES_GRPC_SERVER_ADDR = "0.0.0.0:50052"