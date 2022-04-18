pginit:
	PGDATA=postgres-data/ pg_ctl init

pgstart:
	PGDATA=postgres-data/ pg_ctl start

pgstop:
	PGDATA=postgres-data/ pg_ctl stop
