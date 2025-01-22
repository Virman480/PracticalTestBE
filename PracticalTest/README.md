
# Proyek API Konsumsi dan Pemesanan Ruangan (Golang)

Proyek ini adalah API yang dibangun menggunakan **Golang** untuk mengelola dan menampilkan data pemesanan ruangan serta biaya konsumsi yang terkait. API ini mengambil data dari dua sumber eksternal (API booking dan API master konsumsi), menggabungkannya, dan menghitung total biaya konsumsi. Outputnya disiapkan dalam format JSON yang cocok untuk digunakan pada dashboard.

---

## Fitur
1. Mengambil Data Pemesanan Ruangan: 
   - API mengambil data dari API eksternal berisi ruangan, tanggal booking, jumlah peserta, dan konsumsi.
2. Mengambil Data Master Konsumsi:
   - API mengambil data dari API eksternal berisi jenis konsumsi (seperti Snack Siang, Makan Siang) dan harga maksimumnya.
3. Penggabungan Data:
   - Data pemesanan ruangan dan master konsumsi digabungkan berdasarkan nama konsumsi.
4. Perhitungan Biaya Konsumsi:
   - Menghitung biaya konsumsi berdasarkan jumlah peserta dan harga maksimum konsumsi:
     ```
     biaya = maxPrice * participants
     ```
5. Output JSON:
   - Menyediakan output JSON dalam format yang cocok untuk dashboard.

---

## Teknologi yang Digunakan
- *Golang*: Bahasa pemrograman utama untuk membangun API.
- *Fiber*: Framework web untuk menangani HTTP request.
- *Resty*: HTTP Client untuk mengambil data dari API eksternal.
- *JSON*: Format data output API.



## Struktur Proyek

practical/
├── booking.go             # File untuk mengambil data booking dari API eksternal
├── konsumsi.go            # File untuk mengambil data master konsumsi dari API eksternal
├── main.go                # File utama untuk menjalankan server dan endpoint API
├── go.mod                 # File konfigurasi modul Golang



## Cara Menjalankan Proyek

### 1. Persiapan
1. Pastikan *Golang* sudah terinstal.
   - Unduh Golang: [https://golang.org/dl](https://golang.org/dl)
2. Pastikan koneksi internet aktif untuk mengambil data dari API eksternal.

### 2. Langkah Menjalankan
1. Clone repository proyek ini:
   - git clone <URL_REPOSITORY>
   - cd practical

2. Install dependensi:
   - go mod tidy


3. Jalankan server:
   - go run .

4. Server akan berjalan di:
   http://localhost:8000

---

## **Endpoint API**

### 1. Endpoint: `/dashboard`
Deskripsi: Menampilkan data pemesanan ruangan dan detail konsumsi.

- *Method*: `GET`  
- *URL*: `http://localhost:8000/dashboard`

## Contoh Respons:
json
[
    {
        "roomName": "Ruang Borobudur",
        "officeName": "UID JAYA",
        "bookingDate": "2024-01-03T07:25:52.737Z",
        "participants": 62,
        "consumptionDetails": [
            {
                "name": "Snack Siang",
                "cost": 1240000
            },
            {
                "name": "Makan Siang",
                "cost": 1860000
            },
            {
                "name": "Snack Sore",
                "cost": 1240000
            }
        ],
        "totalCost": 4340000
    }
]


## Penjelasan Logika API

1. Pengambilan Data
   - *Booking*: API mengambil data pemesanan ruangan dari URL berikut:
     https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList

   - *Master Konsumsi*: API mengambil data master konsumsi dari URL berikut:
     https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi

2. Penggabungan Data
   - Untuk setiap ruangan dalam booking, data konsumsi dicocokkan dengan data master konsumsi berdasarkan nama konsumsi.

3. Perhitungan Biaya
   - Biaya konsumsi dihitung dengan rumus:
     cost = maxPrice * participants

4. Output JSON
   - Setiap ruangan dalam booking memiliki:
     - *roomName*: Nama ruangan.
     - *officeName*: Nama kantor pemesan.
     - *bookingDate*: Tanggal pemesanan.
     - *participants*: Jumlah peserta.
     - *consumptionDetails*: Detail biaya konsumsi (nama konsumsi dan biaya).
     - *totalCost*: Total biaya konsumsi untuk ruangan tersebut.

---

## Struktur Data

### *1. Data Pemesanan Ruangan (Booking)**
Contoh Data dari API Booking:
```json
[
    {
        "id": "1",
        "roomName": "Ruang Borobudur",
        "officeName": "UID JAYA",
        "bookingDate": "2024-01-03T07:25:52.737Z",
        "startTime": "2024-01-04T09:00:00.000Z",
        "endTime": "2024-01-04T16:00:00.000Z",
        "participants": 62,
        "listConsumption": [
            { "name": "Snack Siang" },
            { "name": "Makan Siang" },
            { "name": "Snack Sore" }
        ]
    }
]
```

### 2. Data Master Konsumsi
Contoh Data dari API Master Konsumsi:
```json
[
    {
        "id": "1",
        "name": "Snack Siang",
        "maxPrice": 20000
    },
    {
        "id": "2",
        "name": "Makan Siang",
        "maxPrice": 30000
    },
    {
        "id": "3",
        "name": "Snack Sore",
        "maxPrice": 20000
    }
]
```

---

## Pengembangan Selanjutnya

1. Filter Data:
   - Tambahkan fitur filter berdasarkan tanggal atau ruangan.
2. Pagination:
   - Jika data terlalu besar, tambahkan fitur pagination untuk membagi data menjadi halaman.
3. Unit Testing:
   - Gunakan library seperti `testify` untuk menulis unit test.
4. Error Handling yang Lebih Baik:
   - Tangani berbagai error dari API eksternal, seperti timeout atau respons kosong.

---

## **Lisensi**

Proyek ini bebas digunakan dan dimodifikasi untuk keperluan apapun.
