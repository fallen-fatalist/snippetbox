package postgres

import (
	"context"
	"database/sql"
	"time"
)

// Initialize the PostgreSQL database
func OpenDB(dsn string) (*sql.DB, error) {

	// Retry logic: attempt to connect multiple times
	maxRetries := 10                 // Try 6 times (30 seconds total if we wait 5 seconds between retries)
	retryInterval := 5 * time.Second // Retry every 5 seconds
	var db *sql.DB
	var err error

	// Try connecting up to maxRetries times
	for i := 0; i < maxRetries; i++ {
		// Use sql.Open() to create an empty connection pool, using the DSN from the config
		// struct.
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			// Create a context with a timeout for the Ping operation
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Try to ping the database to check if it's available
			err = db.PingContext(ctx)
			// if err == nil {
			// 	// If the ping is successful, use the connection
			// 	postgresDB = db
			// 	slog.Info("PostgreSQL database connection established")
			// 	return postgresDB, nil
			// }
			if err != nil {
				db.Close()
				return nil, err
			}
		}

		// If any error occurs, log it and retry after a delay
		//slog.Errorf("Failed to connect to PostgreSQL, retrying in %v... (attempt %d/%d)", retryInterval, i+1, maxRetries)
		time.Sleep(retryInterval)
	}

	// Return the sql.DB connection pool.
	return db, nil
}
