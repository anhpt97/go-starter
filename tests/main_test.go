package tests

import (
	"go-starter/i18n"
	"go-starter/lib"
	"net/http/httptest"
	"os"
	"testing"
)

var env lib.Env
var db lib.Db
var w *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	i18n.NewI18n("../locales")
	env = lib.NewEnv("../.env")
	db = lib.NewDb(env)
	w = httptest.NewRecorder()
	os.Exit(m.Run())
}
