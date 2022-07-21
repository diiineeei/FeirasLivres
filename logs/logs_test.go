package logs

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	const testCaseLog = "Teste Log"

	Print(testCaseLog)

	content, err := ioutil.ReadFile(LogFile)

	if err != nil {
		t.Fatal("Erro ao ler arquivo de log")
	}

	if !strings.Contains(string(content), testCaseLog) {
		t.Fatal("Arquivo de log invalido")
	}

}
