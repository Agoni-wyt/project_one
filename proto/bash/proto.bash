protoc --proto_path=./proto --go_out=./proto/player_pb --go_opt=paths=source_relative --go-grpc_out=./proto/player_pb --go-grpc_opt=paths=source_relative ./proto/player.proto


protoc --proto_path=./proto --go_out=./proto/chat_pb --go_opt=paths=source_relative --go-grpc_out=./proto/chat_pb --go-grpc_opt=paths=source_relative ./proto/chat.proto
