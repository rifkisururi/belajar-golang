# Arsitektur

```text
cmd/hello/main.go -> internal/greeting
cmd/server/main.go -> internal/httpserver -> internal/greeting
```

`cmd/hello` menangani flag CLI dan output. `internal/greeting` menangani normalisasi nama dan pembuatan pesan; tidak memiliki I/O sehingga mudah diuji.

Belum diperlukan layer controller/service/repository, DI container, atau interface. Tambahkan abstraksi ketika ada perilaku yang perlu diganti atau dependency nyata. HTTP dan persistensi akan diperkenalkan bertahap, bukan dipaksakan ke contoh CLI.

Go 1.27.4 dipilih dari distribusi stabil resmi yang tersedia saat setup dan diverifikasi SHA-256. Tidak ada dependency eksternal sehingga belum ada go.sum. `go mod tidy` akan menjaga deklarasi dependency sesuai import.

Server menggunakan `http.ServeMux` dan handler function. Shutdown dipicu signal OS dan dibatasi timeout melalui `context`; tidak ada framework atau container.
