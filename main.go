package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	text_template "text/template"
)

type Metrics struct {
	PluginID      string
	MetricsSource string
	ScopeName     string
	TimeStamp     string
	InfaceName    string
	MetricsKey    string
	MetricsValue  string
}

func main() {
	outPath := flag.String("out-path", "./output/metrics.dat", "metrics output path")
	tmplPath := flag.String("tmpl-path", "./metrics.dat", "metrics template file path")
	flag.Parse()
	fmt.Printf("metrics template file path: %s, %s\n", *tmplPath, *outPath)

	tmpl := text_template.New("metrics.dat")
	t, err := tmpl.ParseFiles("./metrics.dat")
	if err != nil {
		fmt.Printf("Parse template file error: %v\n", err)
		return
	}
	m := Metrics{
		PluginID:      "nino",
		MetricsSource: "10.120.195.2",
		ScopeName:     "sc",
		InfaceName:    "inface_name",
		MetricsKey:    "cpu_total",
	}

	buffer := new(bytes.Buffer)
	if err := t.Execute(buffer, m); err != nil {
		fmt.Printf("Execute get template error: %v\n", err)
		return
	}

	err = ioutil.WriteFile(*outPath, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Content Files is: %s\n", string(buffer.Bytes()))
		return
	}
}
