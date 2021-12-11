module myapp

go 1.17

replace example => ../celeritas

require (
	example v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.7
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.1.0 // indirect
	github.com/alexedwards/scs/v2 v2.5.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/upper/db/v4 v4.2.1 // indirect
)
