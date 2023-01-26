package main

///*
import (
	"flag"
	"fmt"

	"github.com/pjmd89/gogql/cmd/authorization"
	"github.com/pjmd89/gogql/cmd/newproject"
	"github.com/pjmd89/gogql/lib/generate"
	"github.com/pjmd89/gogql/lib/generate/gqltypes"
	"github.com/pjmd89/gogql/lib/gql"
	"github.com/pjmd89/goutils/systemutils"
)

var (
	what     = "all"
	schema   = "schema"
	driverDB = "mongo"
)

func main() {
	flag.StringVar(&what, "what", what, "que quieres generar (all, auth, project)")
	flag.StringVar(&schema, "schema", schema, "ruta de la carpeta donde se encuentra el schema GQL")
	flag.StringVar(&driverDB, "driverDB", driverDB, "Manejador de base de datos (none, mongo)")
	flag.Parse()
	fs := &systemutils.FS{}
	gql := gql.Init(fs, schema)

	gqlGenerate := generate.NewGqlGenerate(gql.GetSchema(), schema)
	newproject.Generate(gqlGenerate, driver(driverDB))
	authorization.Generate(gqlGenerate)
	fmt.Println(gql, gqlGenerate)
}
func driver(driverDB string) (r gqltypes.DriverDB) {
	switch driverDB {
	case "mongo":
		r = gqltypes.DRIVERDB_MONGO
	default:
		r = gqltypes.DRIVERDB_NONE
	}
	return
}

//*/
