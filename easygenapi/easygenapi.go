////////////////////////////////////////////////////////////////////////////
// Porgram: EasyGen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

package easygenapi

import (
	"bytes"
	"flag"
	"fmt"
	ht "html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	tt "text/template"

	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "EasyGen" // os.Args[0]

// The Options structure holds the values for/from commandline
type Options struct {
	HTML         bool
	TemplateStr  string
	TemplateFile string
}

// common type for a *(text|html).Template value
type template interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Name() string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds all paramters from the command line
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {
	flag.BoolVar(&Opts.HTML, "html", false, "treat the template file as html instead of text")
	flag.StringVar(&Opts.TemplateStr, "ts", "", "template string (in text)")
	flag.StringVar(&Opts.TemplateFile, "tf", "", ".tmpl template file name (default: same as .yaml file)")
}

// The Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr, "\nUsage:\n %s [flags] YamlFileName\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nYamlFileName: The name for the .yaml data and .tmpl template file\n\tOnly the name part, without extension. Can include the path as well.\n")
	os.Exit(0)
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Generate will produce output from the template according to driving data
func Generate(HTML bool, fileName string) string {
	source, err := ioutil.ReadFile(fileName + ".yaml")
	checkError(err)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(source, &m)
	checkError(err)

	// template file name
	fileNameT := fileName
	if len(Opts.TemplateFile) > 0 {
		fileNameT = Opts.TemplateFile
	}

	t, err := parseFiles(HTML, fileNameT+".tmpl")
	checkError(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, m)
	checkError(err)

	return buf.String()
}

// parseFiles, intialization. By Matt Harden @gmail.com
func parseFiles(HTML bool, filenames ...string) (template, error) {
	tname := filepath.Base(filenames[0])

	if HTML {
		// use html template
		t, err := ht.ParseFiles(filenames...)
		return t, err
	}

	// use text template
	funcMap := tt.FuncMap{
		"minus1": minus1,
	}

	if len(Opts.TemplateStr) > 0 {
		t, err := tt.New("TT").Funcs(funcMap).Parse(Opts.TemplateStr)
		return t, err
	}

	t, err := tt.New(tname).Funcs(funcMap).ParseFiles(filenames...)
	return t, err
}

// Exit if error occurs
func checkError(err error) {
	if err != nil {
		fmt.Printf("[%s] Fatal error - %v", progname, err.Error())
		os.Exit(1)
	}
}
