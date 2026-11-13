# Belajar Golang

Pembelajaran Go idiomatis untuk developer senior .NET. Tahap pertama menyediakan CLI kecil, paket internal, dan table-driven tests. Belum ada HTTP server atau database.

## Prasyarat

Go 1.27.4 (linux/amd64 pada cloud environment ini). Toolchain cloud dipasang di `/workspace/.toolchains/go`; aktifkan dengan:

```sh
export PATH="/workspace/.toolchains/go/bin:$PATH"
export GOCACHE=/workspace/.cache/go-build
export GOPATH=/workspace/.cache/go
```

Di mesin lain, gunakan instalasi Go Anda sendiri.

## Menjalankan dan memverifikasi

Dari root repositori:

```sh
go run ./cmd/hello
go run ./cmd/hello -name Rifki
gofmt -l cmd internal
go test -count=1 -v ./...
go vet ./...
go build -o /tmp/belajar-golang-hello ./cmd/hello
/tmp/belajar-golang-hello -name Rifki
```

Output bernama: `Hello, Rifki!`. Tidak ada dependency eksternal, secret, atau layanan pendukung. Jangan commit `.env` maupun database lokal.

## Panduan

- [Roadmap dan catatan belajar](docs/LEARNING.md)
- [Perbandingan .NET dan Go](docs/DOTNET-VS-GOLANG.md)
- [Keputusan arsitektur](docs/ARCHITECTURE.md)
- [Milestone dan checkpoint Git](docs/PROGRESS.md)
