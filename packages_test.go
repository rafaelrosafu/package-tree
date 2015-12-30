package main

import (
	"reflect"
	"testing"
)

func TestParsePackageFromLine(t *testing.T) {
	lineWithoutDependencies := "a:"
	expectedPackage := &Package{
		Name:         "a",
		Processed:    true,
		Dependencies: make([]*Package, 0),
	}

	pkg, err := ParsePackageFromLine(lineWithoutDependencies)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if !reflect.DeepEqual(pkg, expectedPackage) {
		t.Errorf("Couldn't parse package without dependencies: %v != %v", *pkg, expectedPackage)
	}

	lineWithDependencies := "abcde:  autoconf  automake  cd-discid "
	expectedPackage = &Package{
		Name:      "abcde",
		Processed: true,
		Dependencies: []*Package{
			MakeUnprocessedPackage("autoconf"),
			MakeUnprocessedPackage("automake"),
			MakeUnprocessedPackage("cd-discid"),
		},
	}

	pkg, err = ParsePackageFromLine(lineWithDependencies)

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if !reflect.DeepEqual(pkg, expectedPackage) {
		t.Errorf("Couldn't parse package with dependencies: %v != %v", *pkg, *expectedPackage)
	}

	brokenLine := "missing tokens"
	_, err = ParsePackageFromLine(brokenLine)

	if err == nil {
		t.Error("Didn't throw error on broken line")
	}
}
