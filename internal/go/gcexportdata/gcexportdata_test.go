package gcexportdata_test

import (
	"go/token"
	"go/types"
	"log"
	"os"
	"testing"

	"honnef.co/go/tools/internal/go/gcexportdata"
	"honnef.co/go/tools/internal/go/gcimporter"
)

// Test to ensure that gcexportdata can read files produced by App
// Engine Go runtime v1.6.
func TestAppEngine16(t *testing.T) {
	// Open and read the file.
	f, err := os.Open("testdata/errors-ae16.a")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := gcimporter.GetExportData(f)
	if err != nil {
		log.Fatalf("reading export data: %v", err)
	}

	// Decode the export data.
	fset := token.NewFileSet()
	imports := make(map[string]*types.Package)
	pkg, err := gcexportdata.Read(b, fset, imports, "errors")
	if err != nil {
		log.Fatal(err)
	}

	// Print package information.
	got := pkg.Scope().Lookup("New").Type().String()
	want := "func(text string) error"
	if got != want {
		t.Errorf("New.Type = %s, want %s", got, want)
	}
}
