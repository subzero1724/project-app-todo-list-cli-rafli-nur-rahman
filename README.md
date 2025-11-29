# 📝 project-app-todo-list-cli-rafli-nur-rahman

A simple and clean Command Line Interface (CLI) To-Do List Application built using Golang.  
This project allows users to manage tasks directly from the terminal with full CRUD features, search, update, and JSON-based storage.

Project dibuat untuk memenuhi tugas Golang dengan ketentuan:  
✔ Operator  
✔ Variabel  
✔ Function  
✔ Slice  
✔ Error Handling  
✔ File JSON  
✔ CLI Flag / Cobra  
✔ Tampilan tabel  
✔ Validasi input  

---
## 📂 Struktur Folder
```
project-app-todo-list-cli-rafli-nur-rahman/
│
├── cmd/                # Semua perintah CLI
│   ├── add.go
│   ├── delete.go
│   ├── done.go
│   ├── display.go
│   ├── list.go
│   ├── search.go
│   ├── update.go
│   └── root.go
│
├── model/              # Struktur data Task
│   └── task.go
│
├── service/            # Logic bisnis
│   └── task_service.go
│
├── utils/              # Helper JSON
│   └── json.go
│
├── data/               # File data JSON
│   └── tasks.json
│
├── main.go
├── README.md
├── go.sum
└── go.mod

```

---

## 🚀 Cara Menjalankan

### 1️⃣ Clone repository
```
git clone https://github.com/subzero1724/project-app-todo-list-cli-rafli-nur-rahman.git
```

### 2️⃣ Masuk ke folder project
```
cd project-app-todo-list-cli-rafli-nur-rahman
```

### 3️⃣ Install dependency Cobra
```
go mod tidy
```

### 4️⃣ Jalankan aplikasi
```
go run .
```

---

## 🧪 Contoh Penggunaan

### ➕ Tambah task
```
go run . add "Kerjakan Tugas" -d "Tugas Golang"
```

### 📋 Lihat daftar task
```
go run . list
```

### ✔ Tandai sebagai selesai
```
go run . done [id]
```

### 🗑 Hapus task
```
go run . delete [id]
```

### 🔍 Cari task
```
go run . search [Keyword]
```

### ✏ Update task
todo, done
```
go run . update --id=1 --status=completed --priority=high
```

---

## ⚙️ Teknologi yang Digunakan
- Golang 1.21+
- Cobra CLI
- go-pretty table
- JSON file storage
