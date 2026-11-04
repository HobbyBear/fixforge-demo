package infra

import (
	"fmt"
	"testing"
)

func TestParseSql(t *testing.T) {
	sql := `
create index t_agency_email_index
    on t_agency (email);

`
	fmt.Println(SqlMonitor.parseTable(sql))
}

func TestHelloWorld(t *testing.T) {
	fmt.Println("hello world")
}

func TestHelloWorld2(t *testing.T) {
	fmt.Println("hello world2")
}
