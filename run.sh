cd protoc-gen-gql && go build -gcflags="all=-N -l" . &&
mv protoc-gen-gql ../../../go/bin/protoc-gen-gql &&
cd ..  &&
protoc --gql_out=. --go_out=. --go_opt=paths=source_relative test/test.proto