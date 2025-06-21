# justfile

set export:
PROTO_DIR :="examples/protos"
FLATBUFFER_DIR := "examples/flatbuffers"
GO_DIR := "examples/go/fb"
GO_MODULE := "example.com/buffman/fb"

build:
    go build -o buffman main.go

convert:
    ./buffman convert flatbuffers -I "{{PROTO_DIR}}" -o "{{FLATBUFFER_DIR}}"

generate-go:
    ./buffman generate flatbuffers -I "{{FLATBUFFER_DIR}}" -o {{GO_DIR}} -l go -m "{{GO_MODULE}}"

rm:
    rm -rf example/flatbuffers