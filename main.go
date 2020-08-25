package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
)

//go:generate go build -o bin/protoc-gen-sample
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

				opt := f.Descriptor().GetOptions()
				res := &String{}

				rule1, _ := proto.GetExtension(opt, E_Rule1)
				if rule, ok := rule1.(*Rule1); ok && rule != nil {
					res = rule.GetType()
				}

				rule2, _ := proto.GetExtension(opt, E_Rule2)
				if rule, ok := rule2.(*Rule2); ok && rule != nil {
					res = rule.GetType()
				}

				content += fmt.Sprintln(f.Name().String(), res)
			}

			m.OverwriteGeneratorFile("sample.txt", content)
		}
	}
	return m.Artifacts()
}
