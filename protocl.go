package proto

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/wzshiming/ffmt"
)

var f = ffmt.NewOptional(15, ffmt.StlyeP, ffmt.CanRowSpan|ffmt.CanFilterDuplicate)

func A() {
	ffmt.Space = '-'
	fs := token.NewFileSet()

	a, err := parser.ParseDir(fs, "./testd", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range a {
		for _, v2 := range v.Files {
			for _, d := range v2.Decls {
				Decl(d)
			}
		}
	}
}

func Decl(gd ast.Decl) {
	switch gd.(type) {
	default:
		fmt.Println("type err")
	case *ast.BadDecl: // 坏包
		fmt.Println("Bad")
	case *ast.GenDecl: // 定义
		fmt.Println("Gen")
		//f.Print(gd)
	case *ast.FuncDecl: // 函数
		fmt.Println("Func")
		method := gd.(*ast.FuncDecl)

		// 获得注释
		doc := ""
		if method.Doc != nil {
			doc = method.Doc.Text()
		}

		// 获得类型名
		recv := ""
		if method.Recv != nil {
			if len(method.Recv.List) != 0 {
				ty := method.Recv.List[0].Type.(*ast.StarExpr)
				ident := ty.X.(*ast.Ident)
				recv = ident.Name
			}
		}

		// 获得方法名
		name := method.Name.Name

		// 获得参数

		//f.Print(method)
		fmt.Println(doc)
		if recv == "" {
			fmt.Println(name)
		} else {
			fmt.Println(recv + "." + name)
		}
		f.Print(GetField(method.Type.Params))
		f.Print(GetField(method.Type.Results))
	}
}

func GetField(fl *ast.FieldList) (ret []*pair, types map[*ast.Object]bool) {
	if fl == nil {
		return
	}
	types = map[*ast.Object]bool{}
	for k, v := range fl.List {
		switch v.Type.(type) {
		case *ast.Ident:
			typ := v.Type.(*ast.Ident)
			if len(v.Names) == 0 {
				ret = append(ret, &pair{
					Key: typ.String() + fmt.Sprint(k),
					Typ: typ.String(),
				})
			} else {
				for _, v2 := range v.Names {
					ret = append(ret, &pair{
						Key: v2.Name,
						Typ: typ.String(),
					})
				}
			}

		case *ast.StarExpr:
			typ := v.Type.(*ast.StarExpr)
			typ2 := typ.X.(*ast.Ident)
			if len(v.Names) == 0 {
				ret = append(ret, &pair{
					Key: typ2.String() + fmt.Sprint(k),
					Typ: typ2.String(),
				})
			} else {
				for _, v2 := range v.Names {
					ret = append(ret, &pair{
						Key: v2.Name,
						Typ: typ2.String(),
					})
				}
			}
			if typ2.Obj != nil {
				types[typ2.Obj] = true
			}
		}
	}
	return
}
