
start:
	uvicorn app.cmd.main:app --reload


createproto:
    # работает из .app/clients/notes/protos/proto  
	python3 -m grpc_tools.protoc --proto_path=. *.proto --python_out=../genproto --grpc_python_out=../genproto --pyi_out=../genproto