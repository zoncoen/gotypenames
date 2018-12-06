package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
)

var (
	app          = kingpin.New("gotypenames", "Go tool to print all declared type names.")
	filename     = app.Flag("filename", "Target filename.").Short('f').Required().String()
	onlyExported = app.Flag("only-exported", "Print only exported type name.").Default("false").Bool()

	allTypes = []string{"primitive", "array", "map", "func", "struct", "interface", "chan"}
	types    = app.Flag("types", fmt.Sprintf("Filter by type. (%s)", strings.Join(allTypes, ", "))).Default(allTypes...).Enums(allTypes...)
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, *filename, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf(`failed to parse "%s": %s`, *filename, err)
	}

	fil := newFilter(*onlyExported, *types)
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if fil.shouldPrint(x) {
				fmt.Println(x.Name.String())
			}
		}
		return true
	})
}

type filter struct {
	onlyExported bool
	types        map[string]struct{}
}

func newFilter(onlyExported bool, types []string) *filter {
	m := map[string]struct{}{}
	for _, k := range types {
		m[k] = struct{}{}
	}
	return &filter{
		onlyExported: onlyExported,
		types:        m,
	}
}

func (f *filter) shouldPrint(x *ast.TypeSpec) bool {
	if *onlyExported && !x.Name.IsExported() {
		return false
	}
	if _, ok := x.Type.(*ast.Ident); ok {
		if _, ok := f.types["primitive"]; ok {
			return true
		}
	}
	if _, ok := x.Type.(*ast.ArrayType); ok {
		if _, ok := f.types["array"]; ok {
			return true
		}
	}
	if _, ok := x.Type.(*ast.MapType); ok {
		if _, ok := f.types["map"]; ok {
			return true
		}
	}
	if _, ok := x.Type.(*ast.FuncType); ok {
		if _, ok := f.types["func"]; ok {
			return true
		}
	}
	if _, ok := x.Type.(*ast.StructType); ok {
		if _, ok := f.types["struct"]; ok {
			return true
		}
	}
	if _, ok := x.Type.(*ast.InterfaceType); ok {
		if _, ok := f.types["interface"]; ok {
			return true
		}
	}
	if _, ok := x.Type.(*ast.ChanType); ok {
		if _, ok := f.types["chan"]; ok {
			return true
		}
	}
	return false
}
