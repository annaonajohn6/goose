package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// RunMigration demonstrates the correct pattern to avoid connection leaks.
func RunMigration(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Ensure rollback is called. If Commit() is called, Rollback() does nothing.
	defer tx.Rollback()

	// Example of executing a query with context
	rows, err := tx.QueryContext(ctx, "SELECT 1")
	if err != nil {
		return err
	}
	// Ensure rows are closed to release the connection back to the pool.
	defer rows.Close()

	// ... migration logic ...

	return tx.Commit()
}

func main() {
	fmt.Println("Database migration runner initialized.")
}