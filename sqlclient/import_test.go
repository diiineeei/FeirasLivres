package sqlclient

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	err := CreateTable()
	if err != nil {
		t.Fatalf("Erro ao criar tabela %v", err)
	}
	_, err = Import()
	if err != nil {
		t.Fatalf("Erro ao importar dados %v", err)
	}
}

func TestImport(t *testing.T) {
	err := CreateTable()
	if err != nil {
		t.Fatalf("Erro ao criar tabela %v", err)
	}
	_, err = Import()
	if err != nil {
		t.Fatalf("Erro ao importar dados %v", err)
	}

	db, err := Connect()
	if err != nil {
		t.Fatalf("conectar ao banco de dados %v", err)
	}

	rows, err := db.Query("SELECT COUNT(*) FROM " + TABLENAME)
	if err != nil {
		t.Fatalf("executar query %v", err)
	}

	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			t.Fatal(err)
		}
	}
	if count < 879 {
		t.Fatal("numero de linhas diferente do esperado")
	}
}
