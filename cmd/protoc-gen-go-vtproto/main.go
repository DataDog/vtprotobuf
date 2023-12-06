package main

import (
	"flag"
	"strings"

	_ "github.com/DataDog/vtprotobuf/features/clone"
	_ "github.com/DataDog/vtprotobuf/features/equal"
	_ "github.com/DataDog/vtprotobuf/features/grpc"
	_ "github.com/DataDog/vtprotobuf/features/marshal"
	_ "github.com/DataDog/vtprotobuf/features/pool"
	_ "github.com/DataDog/vtprotobuf/features/size"
	_ "github.com/DataDog/vtprotobuf/features/unmarshal"
	"github.com/DataDog/vtprotobuf/generator"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var cfg generator.Config
	var features string
	var f flag.FlagSet

	f.BoolVar(&cfg.AllowEmpty, "allow-empty", false, "allow generation of empty files")
	cfg.Poolable = make(generator.ObjectSet)
	f.Var(&cfg.Poolable, "pool", "use memory pooling for this object")
	f.BoolVar(&cfg.Wrap, "wrap", false, "generate wrapper types")
	f.StringVar(&features, "features", "all", "list of features to generate (separated by '+')")

	protogen.Options{ParamFunc: f.Set}.Run(func(plugin *protogen.Plugin) error {
		gen, err := generator.NewGenerator(plugin, strings.Split(features, "+"), &cfg)
		if err != nil {
			return err
		}
		gen.Generate()
		return nil
	})
}
