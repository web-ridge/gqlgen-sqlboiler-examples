module github.com/web-ridge/gqlgen-sqlboiler-examples

go 1.14

replace github.com/web-ridge/gqlgen-sqlboiler/v2 => ../gqlgen-sqlboiler

replace github.com/web-ridge/sqlboiler-graphql-schema => ../sqlboiler-graphql-schema

replace github.com/web-ridge/utils-go/boilergql => ../utils-go/boilergql

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/ericlagergren/decimal v0.0.0-20191206042408-88212e6cfca9
	github.com/friendsofgo/errors v0.9.2
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/joho/godotenv v1.3.0
	github.com/kat-co/vala v0.0.0-20170210184112-42e1d8b61f12
	github.com/lib/pq v1.2.1-0.20191011153232-f91d3411e481
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/rs/zerolog v1.20.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.4.0
	github.com/vektah/gqlparser/v2 v2.0.1
	github.com/volatiletech/null/v8 v8.1.0
	github.com/volatiletech/randomize v0.0.1
	github.com/volatiletech/sqlboiler/v4 v4.2.0
	github.com/volatiletech/strmangle v0.0.1
	github.com/web-ridge/gqlgen-sqlboiler/v2 v2.0.0-00010101000000-000000000000 // indirect
	github.com/web-ridge/utils-go/api v0.0.0-20201017151809-22aa4aa28bbb
	github.com/web-ridge/utils-go/boilergql v0.0.0-20201017192349-9efd036e9e61
	github.com/zenazn/goji v0.9.0 // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
)
