# Aplikasi Todo 📝  
Aplikasi Todo sederhana dan efisien yang dibangun menggunakan **Golang** untuk membantu mengelola daftar tugas dengan mudah. Proyek ini merupakan bagian dari praktikum Laboratorium Informatika Universitas Gunadarma.  

## Fitur 🚀  
- ✅ Menambahkan, melihat, mengedit, dan menghapus tugas.  
- 🔍 Filter tugas berdasarkan status penyelesaian.  
- 📦 Ringan dan cepat, memanfaatkan performa Golang.  
- 📁 Penyimpanan data menggunakan JSON atau database (contoh: SQLite).  
- 🖥️ Antarmuka yang sederhana dan ramah pengguna melalui Command Line Interface (CLI).

## Persyaratan 🛑
- Golang (versi 1.18 atau lebih baru)
- Thunder Client (ekstensi VS Code) untuk pengujian API.
- Opsional: SQLite atau database lain jika menggunakan penyimpanan database.


## Instalasi ⚙️  
1. Clone repository ini:  
   ```bash
   git clone https://github.com/fidhera/todoApp-golang.git
   cd todoApp-golang

2. Instal dependensi (jika diperlukan):
   ```bash
   go mod tidy
   
3. Jalankan aplikasi:
   ```bash
   cd todo-app-golang

## Penggunaan API dengan ekstensi Thunder Client VS Code 🛠️:
1. Jalankan aplikasi:
   ```bash
   go run main.go

2. Nanti akan muncul seperti ini:
   
   Connected to database

![image](https://github.com/user-attachments/assets/9ee93beb-598e-45d5-96ba-f3016671363f)


4. Copy url **http://127.0.0.1:2597**.
   
6. Gunakan Thunder Client (ekstensi VS Code) untuk melakukan CRUD pada data Todo:
   - Klik New Request.
   - Masukkan atau paste url tadi pada box url.
   - Klik Send.
     
7. Endpoint API:
   - GET /todos: Mendapatkan daftar semua tugas.
   - GET /todos/{id}: Mendapatkan detail tugas berdasarkan ID.
   - POST /todos: Menambahkan tugas baru.
     **Body:**
     ```bash
     {
     "title": "Belajar Golang",
     "description": "Menyelesaikan proyek Todo App",
     "status": "pending"
     }

   - PUT /todos/{id}: Memperbarui tugas berdasarkan ID.
     **Body:**
     ```bash
     {
     "title": "Belajar Golang Lanjutan",
     "description": "Menyelesaikan tugas lanjutan",
     "status": "completed"
     }

   - DELETE /todos/{id}: Menghapus tugas berdasarkan ID.

8. Tambahkan koleksi API di Thunder Client untuk mempermudah pengujian endpoint.

7. Ikuti petunjuk pada layar untuk menambahkan, melihat, mengedit, atau menghapus tugas.
