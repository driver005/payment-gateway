package migrate

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"

	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/pop/v6"
	"github.com/pkg/errors"
)

var SQLTemplateFuncs = map[string]interface{}{
	"identifier": Identifier,
}

var identifierPattern = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*$")

func Identifier(i string) (string, error) {
	if !identifierPattern.MatchString(i) {
		return "", fmt.Errorf("invalid SQL identifier '%s'", i)
	}
	return i, nil
}

func ParameterizedMigrationContent(params map[string]interface{}) func(mf Migration, c *pop.Connection, r []byte, usingTemplate bool) (string, error) {
	return func(mf Migration, c *pop.Connection, b []byte, usingTemplate bool) (string, error) {
		content := ""
		if usingTemplate {
			t := template.New("migration")
			t.Funcs(SQLTemplateFuncs)
			t, err := t.Parse(string(b))
			if err != nil {
				return "", errors.Wrapf(err, "could not parse template %s", mf.Path)
			}
			var bb bytes.Buffer
			err = t.Execute(&bb, struct {
				IsSQLite       bool
				IsCockroach    bool
				IsMySQL        bool
				IsMariaDB      bool
				IsPostgreSQL   bool
				DialectDetails *pop.ConnectionDetails
				Parameters     map[string]interface{}
			}{
				IsSQLite:       c.Dialect.Name() == "sqlite3",
				IsCockroach:    c.Dialect.Name() == "cockroach",
				IsMySQL:        c.Dialect.Name() == "mysql",
				IsMariaDB:      c.Dialect.Name() == "mariadb",
				IsPostgreSQL:   c.Dialect.Name() == "postgres",
				DialectDetails: c.Dialect.Details(),
				Parameters:     params,
			})
			if err != nil {
				return "", errors.Wrapf(err, "could not execute migration template %s", mf.Path)
			}
			content = bb.String()
		} else {
			content = string(b)
		}

		if mf.Type == "fizz" {
			var err error
			content, err = fizz.AString(content, c.Dialect.FizzTranslator())
			if err != nil {
				return "", errors.Wrapf(err, "could not fizz the migration %s", mf.Path)
			}
		}

		return content, nil
	}
}
