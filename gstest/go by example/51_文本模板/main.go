package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1") //我们可以创建一个新模板，并从字符串解析其正文, 模板是静态文本和包含在”动作“中用于动态插入内容的混合体。
	t1, err := t1.Parse("Value is{{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n")) //我们可以使用 template.Must 函数，在 Parse 错误时返回 panic。 这对于在全局作用域中初始化的模板非常有用。

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"go",
		"Rust",
		"C++",
		"C#",
	})
}
