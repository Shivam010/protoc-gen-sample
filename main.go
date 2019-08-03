package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
)

//TODO: GO111MODULE=on go build -o bin/protoc-gen-sample
//go:generate protoc -I . type.proto --go_out=plugins=grpc:.
//go:generate protoc -I . check.proto --plugin=bin/protoc-gen-sample --sample_out=:.

func main() {
	pgs.Init().RegisterModule(&cMod{&pgs.ModuleBase{}}).Render()
}

type cMod struct {
	*pgs.ModuleBase
}

func (m *cMod) Name() string {
	return "Sample"
}

func (m *cMod) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, fl := range targets {
		for _, msg := range fl.Messages() {

			content := ""

			for _, f := range msg.Fields() {

				if f.Name().String() == "field_1" {
					opt := f.Descriptor().GetOptions()
					ext, _ := proto.GetExtension(opt, E_Rule1)
					content += fmt.Sprintln(f.Name().String(), ext.(*Rule1).Type)
				} else {
					opt := f.Descriptor().GetOptions()
					ext, _ := proto.GetExtension(opt, E_Rule2)
					content += fmt.Sprintln(f.Name().String(), ext.(*Rule2).GetType())
				}
			}

			m.OverwriteGeneratorFile("sample.txt", content)
		}
	}
	return m.Artifacts()
}
