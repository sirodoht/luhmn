pginit:
	PGDATA=postgres-data/ pg_ctl init
	PGDATA=postgres-data/ pg_ctl start
	createuser luhmn
	psql -U sirodoht -d postgres -c "ALTER USER luhmn CREATEDB;"
	psql -U luhmn -d postgres -c "CREATE DATABASE luhmn;"
	psql -U luhmn -d luhmn -f schema.sql

pgstart:
	PGDATA=postgres-data/ pg_ctl start

pgstop:
	PGDATA=postgres-data/ pg_ctl stop
