# Go Gin Repository Pattern

Boilerplate Go adalah template proyek Go yang mengimplementasikan clean architecture dengan repository pattern. Proyek ini menggunakan Gin Gonic sebagai web framework dan mendukung REST API serta gRPC.

## Fitur

- Clean Architecture dengan Repository Pattern
- REST API menggunakan Gin Gonic
- gRPC support
- Swagger documentation
- Database migration
- Environment configuration menggunakan Viper
- Unit testing
- Makefile untuk kemudahan pengembangan

## Prasyarat

- Go 1.21 atau lebih baru
- PostgreSQL
- Make
- Swag (untuk Swagger documentation)
- golang-migrate (untuk database migration)

## Instalasi

1. Clone repository:
```bash
git clone https://github.com/sekolahmu/boilerplate-go.git
cd boilerplate-go
```

2. Install dependencies:
```bash
make deps
```

3. Copy file environment:
```bash
cp .env.example .env
```

4. Sesuaikan konfigurasi di file `.env` dengan environment Anda

5. Jalankan migrasi database:
```bash
make migrate-up
```

6. Generate Swagger documentation:
```bash
make swagger
```

7. Build dan jalankan aplikasi:
```bash
make build
make run
```

## Struktur Proyek

```
.
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── domain/
│   │   └── entity/
│   ├── repository/
│   │   └── interface/
│   ├── usecase/
│   │   └── interface/
│   ├── delivery/
│   │   ├── http/
│   │   └── grpc/
│   └── middleware/
├── pkg/
│   ├── logger/
│   └── utils/
├── docs/
├── migrations/
├── test/
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Penggunaan

### API Endpoints

Setelah aplikasi berjalan, Anda dapat mengakses API di `http://localhost:8080/api/v1`:

- `POST /users` - Membuat user baru
- `GET /users/:id` - Mendapatkan user berdasarkan ID
- `PUT /users/:id` - Mengupdate user
- `DELETE /users/:id` - Menghapus user
- `GET /users` - Mendapatkan daftar user dengan pagination

### Swagger Documentation

Dokumentasi API tersedia di `http://localhost:8080/swagger/index.html`

## Testing

Untuk menjalankan test:

```bash
make test
```

## Migrasi Database

Untuk menjalankan migrasi database:

```bash
make migrate-up
```

Untuk rollback migrasi:

```bash
make migrate-down
```

## Kontribusi

1. Fork repository
2. Buat branch fitur (`git checkout -b feature/amazing-feature`)
3. Commit perubahan (`git commit -m 'Add some amazing feature'`)
4. Push ke branch (`git push origin feature/amazing-feature`)
5. Buat Pull Request

## Lisensi

Proyek ini dilisensikan di bawah MIT License - lihat file [LICENSE](LICENSE) untuk detail. 
