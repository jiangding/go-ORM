package orm

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestNewEngine(t *testing.T) {
	t.Helper()
	engine, err := NewDB("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	defer engine.Close()


}