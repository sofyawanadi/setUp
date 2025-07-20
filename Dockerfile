# Gunakan image Golang sebagai base build
FROM golang:1.23.6 AS builder

# Set direktori kerja di dalam container
WORKDIR /app

# Copy semua file dari project ke dalam container
COPY . .

# Download dependensi
RUN go mod download

# Build binary aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# Gunakan image kecil untuk menjalankan binary (lebih ringan)
FROM alpine:latest

# Install sertifikat root agar koneksi TLS tidak error (opsional)
RUN apk --no-cache add ca-certificates

# Direktori kerja di container final
WORKDIR /root/

# Copy binary dari stage builder
COPY --from=builder /app/main .

# Jalankan binary
CMD ["./main"]