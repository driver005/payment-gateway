package helper

import "fmt"

const (
	// DriverMySQL is the mysql driver name.
	DriverMySQL = "mysql"

	// DriverPostgreSQL is the postgres driver name.
	DriverPostgreSQL = "postgres"

	// DriverCockroachDB is the cockroach driver name.
	DriverCockroachDB = "cockroach"

	// UnknownDriver is the driver name if the driver is unknown.
	UnknownDriver = "unknown"
)

// Canonicalize returns constants DriverMySQL, DriverPostgreSQL, DriverCockroachDB, UnknownDriver, depending on `database`.
func Canonicalize(database string) string {
	switch database {
	case "mysql":
		return DriverMySQL
	case "pgx", "pq", "postgres":
		return DriverPostgreSQL
	case "cockroach":
		return DriverCockroachDB
	default:
		return UnknownDriver
	}
}

// MustCanonicalize returns constants DriverMySQL, DriverPostgreSQL, DriverCockroachDB or fatals.
func MustCanonicalize(database string) string {
	d := Canonicalize(database)
	if d == UnknownDriver {
		fmt.Printf("Unknown database driver: %s", database)
	}
	return d
}