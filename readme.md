# GraphQL Server for Microservices

Untuk melanjutkan development untuk server graphql berbasih golang ini dapat melakukan step sebagaimana berikut:

1. Update file `schema.graphql` untuk membuat schema berdasarkan kebutuhan
2. Update code menggunakan command: `go run github.com/99designs/gqlgen generate`
3. Update resolver untuk implementasi
4. Jalankan server dengan command `go run server/server.go`
