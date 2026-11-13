# .NET dan Go

| C# / .NET | Go | Implikasi |
|---|---|---|
| Solution dan csproj | go.mod dan direktori paket | Modul mengelola identitas dan dependency; paket mengelompokkan kode. |
| Program.Main | func main dalam package main | Entry point berada di cmd/hello. |
| Static service method | Fungsi tingkat paket | Greeting tidak membutuhkan class atau interface. |
| internal access modifier | Direktori internal | Pembatasan import ditentukan struktur direktori. |
| xUnit Theory + InlineData | Table-driven test + t.Run | Data kasus diuji melalui loop eksplisit. |
| dotnet test | go test ./... | Tes ditemukan dari file _test.go. |
| dotnet build | go build | Binary berisi program tanpa runtime .NET terpisah. |

Contoh C#: `Greeting.Message("Rifki")`; contoh Go: `greeting.Message("Rifki")`. Kemiripan sintaks tidak berarti Go membutuhkan static class: `Message` adalah fungsi langsung dalam paket. String kosong adalah input valid yang diberi default; tidak perlu membuat error buatan untuk alur ini.

Konsep context, error wrapping, HTTP middleware, interface implisit, dan concurrency akan diperkenalkan ketika ada kebutuhan konkret.
