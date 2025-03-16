package main

///*
import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/pjmd89/gogql/cmd/authorization"
	"github.com/pjmd89/gogql/cmd/newproject"
	"github.com/pjmd89/gogql/lib/generate"
	"github.com/pjmd89/gogql/lib/generate/gqltypes"
	"github.com/pjmd89/gogql/lib/gql"
	"github.com/pjmd89/goutils/systemutils"
)

var (
	what        = "all"
	projectName = "my-project"
	schema      = "schema"
	driverDB    = "mongo"
)

func main() {
	flag.StringVar(&what, "what", what, "que quieres generar (all, auth, project)")
	flag.StringVar(&projectName, "projectName", projectName, "nombre del proyecto")
	flag.StringVar(&schema, "schema", schema, "ruta de la carpeta donde se encuentra el schema GQL")
	flag.StringVar(&driverDB, "driverDB", driverDB, "Manejador de base de datos (none, mongo)")
	flag.Parse()
	fs := &systemutils.FS{}
	gql := gql.Init(fs, schema)
	fmt.Printf("Iniciando proyecto %s\n", projectName)
	exec.Command("bash", "-c", "go mod init "+projectName).Run()

	gqlGenerate := generate.NewGqlGenerate(gql.GetSchema(), schema)
	fmt.Println("Generando proyecto")
	newproject.Generate(gqlGenerate, driver(driverDB))
	fmt.Println("Generando autorización")
	authorization.Generate(gqlGenerate)
	exec.Command("bash", "-c", "go get -d github.com/pjmd89/gogql").Run()
	exec.Command("bash", "-c", "go get -d github.com/pjmd89/goutils").Run()
	exec.Command("bash", "-c", "go get -d github.com/pjmd89/mongomodel").Run()
	exec.Command("bash", "-c", "go mod tidy && go mod vendor").Run()
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
