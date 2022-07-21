package feiraslivres

import (
	"testesrod/sqlclient"
	"testing"
)

func TestInsertFeira(t *testing.T) {
	err := sqlclient.CreateTable()
	if err != nil {
		t.Fatalf("Erro ao criar tabela %v", err)
	}

	testCase := FeiraLivreInsert{
		LONGITUDE:  -46550164,
		LATITUDE:   -23558733,
		SETCENS:    355030885000091,
		AREAP:      3550308005040,
		CODDIST:    87,
		DISTRITO:   "VILA FORMOSA",
		CODSUBPREF: 26,
		SUBPREFE:   "ARICANDUVA-FORMOSA-CARRAO",
		REGIAO5:    "Leste",
		REGIAO8:    "Leste 1",
		NOME_FEIRA: "VILA FORMOSA",
		REGISTRO:   "4041-0",
		LOGRADOURO: "RUA MARAGOJIPE",
		NUMERO:     "S/N",
		BAIRRO:     "VL FORMOSA",
		REFERENCIA: "TV RUA PRETORIA",
	}

	err = InsertFeira(testCase)
	if err != nil {
		t.Errorf("erro ao executar insert %v", err)
	}

	feira, err := SelectFeira("VILA FORMOSA", "Leste", "VILA FORMOSA", "VL FORMOSA")
	if err != nil {
		t.Fatalf("erro ao consultar feira %v", err)
	}

	if len(feira) != 1 {
		t.Errorf("quantidade de registros diferente da esperada")
	}

	if feira[0].LONGITUDE != testCase.LONGITUDE {
		t.Fatal("Erro parametro LONGITUDE diferente do esperado")
	}
	if feira[0].LATITUDE != testCase.LATITUDE {
		t.Fatal("Erro parametro LATITUDE diferente do esperado")
	}
	if feira[0].SETCENS != testCase.SETCENS {
		t.Fatal("Erro parametro SETCENS diferente do esperado")
	}
	if feira[0].AREAP != testCase.AREAP {
		t.Fatal("Erro parametro AREAP diferente do esperado")
	}
	if feira[0].CODDIST != testCase.CODDIST {
		t.Fatal("Erro parametro CODDIST diferente do esperado")
	}
	if feira[0].DISTRITO != testCase.DISTRITO {
		t.Fatal("Erro parametro DISTRITO diferente do esperado")
	}
	if feira[0].CODSUBPREF != testCase.CODSUBPREF {
		t.Fatal("Erro parametro CODSUBPREF diferente do esperado")
	}
	if feira[0].SUBPREFE != testCase.SUBPREFE {
		t.Fatal("Erro parametro SUBPREFE diferente do esperado")
	}
	if feira[0].REGIAO5 != testCase.REGIAO5 {
		t.Fatal("Erro parametro REGIAO5 diferente do esperado")
	}
	if feira[0].REGIAO8 != testCase.REGIAO8 {
		t.Fatal("Erro parametro REGIAO8 diferente do esperado")
	}
	if feira[0].NOME_FEIRA != testCase.NOME_FEIRA {
		t.Fatal("Erro parametro NOME_FEIRA diferente do esperado")
	}
	if feira[0].REGISTRO != testCase.REGISTRO {
		t.Fatal("Erro parametro REGISTRO diferente do esperado")
	}
	if feira[0].LOGRADOURO != testCase.LOGRADOURO {
		t.Fatal("Erro parametro LOGRADOURO diferente do esperado")
	}
	if feira[0].NUMERO != testCase.NUMERO {
		t.Fatal("Erro parametro NUMERO diferente do esperado")
	}
	if feira[0].BAIRRO != testCase.BAIRRO {
		t.Fatal("Erro parametro BAIRRO diferente do esperado")
	}
	if feira[0].REFERENCIA != testCase.REFERENCIA {
		t.Fatal("Erro parametro REFERENCIA diferente do esperado")
	}
}

func TestDeleteFeira(t *testing.T) {
	err := sqlclient.CreateTable()
	if err != nil {
		t.Fatalf("Erro ao criar tabela %v", err)
	}

	testCase := FeiraLivreInsert{
		LONGITUDE:  -46550164,
		LATITUDE:   -23558733,
		SETCENS:    355030885000091,
		AREAP:      3550308005040,
		CODDIST:    87,
		DISTRITO:   "VILA FORMOSA",
		CODSUBPREF: 26,
		SUBPREFE:   "ARICANDUVA-FORMOSA-CARRAO",
		REGIAO5:    "Leste",
		REGIAO8:    "Leste 1",
		NOME_FEIRA: "Feira para Deletar",
		REGISTRO:   "4041-0",
		LOGRADOURO: "RUA MARAGOJIPE",
		NUMERO:     "S/N",
		BAIRRO:     "VL FORMOSA",
		REFERENCIA: "TV RUA PRETORIA",
	}

	err = InsertFeira(testCase)
	if err != nil {
		t.Errorf("erro ao executar insert %v", err)
	}

	err = DeleteFeira(IDFeiraLivre{ID: 1})
	if err != nil {
		t.Errorf("erro ao deletar feira %v", err)
	}

	feira, err := SelectFeira("VILA FORMOSA", "Leste", "Feira para Deletar", "VL FORMOSA")
	if err != nil {
		t.Fatalf("erro ao consultar feira %v", err)
	}

	if len(feira) != 0 {
		t.Errorf("quantidade de registros diferente da esperada")
	}
}

func TestUpdateFeira(t *testing.T) {
	err := sqlclient.CreateTable()
	if err != nil {
		t.Fatalf("Erro ao criar tabela %v", err)
	}

	testCaseInsert := FeiraLivreInsert{
		LONGITUDE:  -46550164,
		LATITUDE:   -23558733,
		SETCENS:    355030885000091,
		AREAP:      3550308005040,
		CODDIST:    87,
		DISTRITO:   "VILA FORMOSA",
		CODSUBPREF: 26,
		SUBPREFE:   "ARICANDUVA-FORMOSA-CARRAO",
		REGIAO5:    "Leste",
		REGIAO8:    "Leste 1",
		NOME_FEIRA: "Feira Antes de Atualizar",
		REGISTRO:   "4041-0",
		LOGRADOURO: "RUA MARAGOJIPE",
		NUMERO:     "S/N",
		BAIRRO:     "VL FORMOSA",
		REFERENCIA: "TV RUA PRETORIA",
	}

	err = InsertFeira(testCaseInsert)
	if err != nil {
		t.Errorf("erro ao executar insert %v", err)
	}

	testCase := FeiraLivre{
		ID:         1,
		LONGITUDE:  -46550164,
		LATITUDE:   -23558733,
		SETCENS:    355030885000091,
		AREAP:      3550308005040,
		CODDIST:    87,
		DISTRITO:   "VILA FORMOSA",
		CODSUBPREF: 26,
		SUBPREFE:   "ARICANDUVA-FORMOSA-CARRAO",
		REGIAO5:    "Leste",
		REGIAO8:    "Leste 1",
		NOME_FEIRA: "Feira Pos Atualizar",
		REGISTRO:   "4041-0",
		LOGRADOURO: "RUA MARAGOJIPE",
		NUMERO:     "S/N",
		BAIRRO:     "VL FORMOSA",
		REFERENCIA: "TV RUA PRETORIA",
	}
	err = UpdateFeira(testCase)
	if err != nil {
		t.Errorf("erro ao atualizar feira %v", err)
	}

	feira, err := SelectFeira("VILA FORMOSA", "Leste", "Feira Pos Atualizar", "VL FORMOSA")
	if err != nil {
		t.Fatalf("erro ao consultar feira %v", err)
	}

	if len(feira) != 1 {
		t.Errorf("quantidade de registros diferente da esperada")
	}

	if feira[0].LONGITUDE != testCase.LONGITUDE {
		t.Fatal("Erro parametro LONGITUDE diferente do esperado")
	}
	if feira[0].LATITUDE != testCase.LATITUDE {
		t.Fatal("Erro parametro LATITUDE diferente do esperado")
	}
	if feira[0].SETCENS != testCase.SETCENS {
		t.Fatal("Erro parametro SETCENS diferente do esperado")
	}
	if feira[0].AREAP != testCase.AREAP {
		t.Fatal("Erro parametro AREAP diferente do esperado")
	}
	if feira[0].CODDIST != testCase.CODDIST {
		t.Fatal("Erro parametro CODDIST diferente do esperado")
	}
	if feira[0].DISTRITO != testCase.DISTRITO {
		t.Fatal("Erro parametro DISTRITO diferente do esperado")
	}
	if feira[0].CODSUBPREF != testCase.CODSUBPREF {
		t.Fatal("Erro parametro CODSUBPREF diferente do esperado")
	}
	if feira[0].SUBPREFE != testCase.SUBPREFE {
		t.Fatal("Erro parametro SUBPREFE diferente do esperado")
	}
	if feira[0].REGIAO5 != testCase.REGIAO5 {
		t.Fatal("Erro parametro REGIAO5 diferente do esperado")
	}
	if feira[0].REGIAO8 != testCase.REGIAO8 {
		t.Fatal("Erro parametro REGIAO8 diferente do esperado")
	}
	if feira[0].NOME_FEIRA != testCase.NOME_FEIRA {
		t.Fatal("Erro parametro NOME_FEIRA diferente do esperado")
	}
	if feira[0].REGISTRO != testCase.REGISTRO {
		t.Fatal("Erro parametro REGISTRO diferente do esperado")
	}
	if feira[0].LOGRADOURO != testCase.LOGRADOURO {
		t.Fatal("Erro parametro LOGRADOURO diferente do esperado")
	}
	if feira[0].NUMERO != testCase.NUMERO {
		t.Fatal("Erro parametro NUMERO diferente do esperado")
	}
	if feira[0].BAIRRO != testCase.BAIRRO {
		t.Fatal("Erro parametro BAIRRO diferente do esperado")
	}
	if feira[0].REFERENCIA != testCase.REFERENCIA {
		t.Fatal("Erro parametro REFERENCIA diferente do esperado")
	}
}
