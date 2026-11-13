# Catatan Pembelajaran

## Roadmap bertahap

1. Modul, paket, CLI, dan table-driven tests.
2. HTTP server berbasis net/http, endpoint health/greet, graceful shutdown, dan httptest.
3. Konfigurasi environment, log/slog, timeout, dan graceful shutdown.
4. REST API sederhana dengan penyimpanan in-memory dan validasi input.
5. context.Context, cancellation, serta explicit dependency injection sesuai kebutuhan.
6. Persistensi SQL, migrasi, dan integration tests setelah kebutuhan database ditetapkan.
7. Middleware dan autentikasi setelah kebutuhan keamanan ditetapkan.

Setiap tahap dijelaskan, diimplementasikan, diformat, diuji, diperiksa go vet, didokumentasikan, di-commit, dan di-push secara terpisah. Jika push gagal, tahap berikutnya berhenti.

## Step 01 — Modul dan program pertama

**Objective:** membuktikan alur edit–test–build–run tanpa framework.

**Equivalent in ASP.NET Core:** `dotnet new console` dan proyek xUnit biasanya menghasilkan beberapa proyek dalam solution. Di Go, satu `go.mod` mendefinisikan modul, sementara direktori mendefinisikan paket.

**Implementation in Golang:** `cmd/hello` berisi entry point. `internal/greeting` menyediakan fungsi `Message(name string) string`. Fungsi murni ini tidak membutuhkan struct, interface, atau dependency injection. Paket `internal` membatasi import dari luar subtree modul.

**Key Differences:** Go menggunakan fungsi tingkat paket; tidak semua perilaku memerlukan service class. File `_test.go` hidup bersama paket dan tidak masuk binary produksi. Subtest table-driven mencakup nama biasa, kosong, whitespace, dan Unicode. Tidak diperlukan goroutine untuk pekerjaan sinkron kecil ini.

**Verification:** jalankan perintah README. `go test -count=1 -v ./...` harus mengeksekusi TestMessage beserta lima subtest. `go vet` memeriksa pola mencurigakan; keberhasilannya bukan pengganti tes perilaku. Jalankan binary untuk membuktikan entry point tersambung ke paket.

**What You Learned:** module path, package boundary, exported function, zero-dependency CLI, flag, standard testing, dan perbedaan fungsi dengan service class.

## Step 02 — HTTP server dan handler

`net/http` menyediakan router dan server di standard library. `httptest` memungkinkan handler diuji tanpa membuka port. `http.Server.Shutdown` menerima `context.Context`, sehingga graceful shutdown memiliki batas waktu eksplisit.

## Step 03 — Konfigurasi

Environment variable adalah padanan sederhana `appsettings` plus options binding. Go tidak membutuhkan configuration framework untuk dua nilai; fungsi `Load` mengembalikan struct dan error secara eksplisit.
