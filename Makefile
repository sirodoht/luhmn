pginit:
	PGDATA=postgres-data/ pg_ctl init
	PGDATA=postgres-data/ pg_ctl start
	createuser fdm
	psql -U sirodoht -d postgres -c "ALTER USER fdm CREATEDB;"
	psql -U fdm -d postgres -c "CREATE DATABASE fdm;"
	psql -U fdm -d fdm -f schema.sql

pgstart:
	PGDATA=postgres-data/ pg_ctl start

pgstop:
	PGDATA=postgres-data/ pg_ctl stop
