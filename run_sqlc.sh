pg_dump $db_url > db/internals/schema.sql --schema-only
sqlc generate
