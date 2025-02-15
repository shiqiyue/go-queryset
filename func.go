package go_queryset

import (
	"context"
	"github.com/shiqiyue/go-queryset/internal/parser"
	"github.com/shiqiyue/go-queryset/internal/queryset/generator"
	"log"
)

func Gen(inFile string) {
	ctx := context.Background()
	g := generator.Generator{
		StructsParser: &parser.Structs{},
	}

	outFile := "autogenerated_" + inFile

	if err := g.Generate(ctx, inFile, outFile); err != nil {
		log.Fatalf("can't generate query sets: %s", err)
	}
}
