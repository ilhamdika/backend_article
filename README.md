# Backend Golang - Panduan Menjalankan Aplikasi

## 📌 Persyaratan

Sebelum menjalankan aplikasi, pastikan Anda sudah menginstal:

- **Golang** (versi 1.18 atau lebih baru)
- **Docker** (jika ingin menggunakan container)

---

## ⚡ Menjalankan Secara Lokal

### 1️⃣ **Clone Repository** (Opsional)

Jika Anda belum memiliki kode sumber, unduh dengan perintah berikut:

```sh
git clone https://github.com/ilhamdika/backend_article.git
cd backend-article
```

Lalu .env.example bisa di copy ke .env dan disitu ada dua Jika menggunakan local, Jika menggunakan docker, tinggal sesuaikan saja untuk sesuai dengan kemauan

```sh
note: buat databasenya terlebih dahulu
```

lalu tinggal sesuaikan namanya seperti di .env

---

### 2️⃣ **Install Dependensi**

```sh
go mod tidy
```

---

### 3️⃣ **Jalankan Aplikasi**

```sh
go run main.go
```

---

## 🚀 Menjalankan dengan Docker

### 1️⃣ **Menjalankan dengan Docker Compose**

Repositori ini sudah memiliki file docker-compose.yml yang dapat Anda jalankan dengan perintah berikut:

```sh
docker-compose up -d
```

---

## 🚀 Cek Status API

```sh
curl http://localhost:8080/article
```

karena ini berjalan di port 8080. atau mungkin bila anda ingin menggunakan docker bisa mengeksposenya dengan portnya sesuai dengan port yang diinginkan

---

### 📥 Import Koleksi ke Postman

di link berikut berisikan dokumentasi postman:

**\*https://documenter.getpostman.com/view/26658030/2sAYX9ogcX**

Jika ingin menggunakan API secara lokal, unduh file koleksi dan impor ke Postman:

**\*Backend Article.postman_collection.json**
lalu import ke postman

---

Terima kasih atas kunjungan Anda!
