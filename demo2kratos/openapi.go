package demo2kratos

import (
	"embed"

	"github.com/go-xlan/yaml-go-edit/yamlv3edit"
	"github.com/yyle88/rese"
	"gopkg.in/yaml.v3"
)

//go:embed openapi.yaml
var files embed.FS

func GetOpenapiContent(docTitle string) []byte {
	// 读取文档的内容
	content := rese.A1(files.ReadFile("openapi.yaml"))
	// 设置文档的标题
	content = yamlv3edit.ChangeYamlFieldValue(content, []string{"info", "title"}, func(node *yaml.Node) {
		if node.Value == "" {
			node.SetString(docTitle)
		}
	})
	return content
}
