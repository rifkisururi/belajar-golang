# Progress

## Step 01 — Modul, CLI, dan tes

- Commit yang direncanakan: `feat(step-01): initialize Go module and tested CLI`.
- Branch: `learning/step-01`.
- Cakupan: go.mod, CLI, paket greeting, lima subtest, dan dokumentasi awal.
- Verifikasi wajib: gofmt, go mod tidy, go test ./..., go vet ./..., build, dan pemeriksaan output CLI.
- Checkpoint: gunakan `git log -1 --format='%h %s'` untuk hash commit tahap ini dan `git ls-remote origin refs/heads/learning/step-01` untuk membandingkan remote setelah push. Hasil push dilaporkan terpisah; dokumen ini tidak mengasumsikan push sudah berhasil.

Tahap selanjutnya hanya dimulai setelah checkpoint tahap ini berhasil diuji dan di-push.

## Step 03 — Konfigurasi environment

- `internal/config` memuat `APP_ADDR` dan `APP_LOG_LEVEL` dengan default aman.
- Commit: `feat(step-03): add environment configuration`.

## Step 02 — HTTP server dan handler

- Implementasi: `internal/httpserver` dengan `/healthz` dan `/greet`, plus `cmd/server`.
- Verification: table-driven `httptest`, `go test`, `go vet`, build, dan smoke request lokal.
- Commit: `feat(step-02): add net/http server and handlers`.
