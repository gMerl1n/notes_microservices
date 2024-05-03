
createproto:
    # работает из .app/protos/proto  
	python3 -m grpc_tools.protoc --proto_path=. *.proto --python_out=../genproto --grpc_python_out=../genproto --pyi_out=../genproto

start:
	python3 -m app.cmd.main