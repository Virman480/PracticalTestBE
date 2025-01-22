
## **Fitur**
1. **Mengambil Data Pemesanan Ruangan**: 
   - API mengambil data dari API eksternal yang berisi ruangan, tanggal booking, jumlah peserta, dan konsumsi.
2. **Mengambil Data Master Konsumsi**:
   - API mengambil data dari API eksternal yang berisi jenis konsumsi (seperti Snack Siang, Makan Siang) dan harga maksimumnya.
3. **Penggabungan Data**:
   - Data pemesanan ruangan dan master konsumsi digabungkan berdasarkan nama konsumsi.
4. **Perhitungan Biaya Konsumsi**:
   - Menghitung biaya konsumsi berdasarkan jumlah peserta dan harga maksimum konsumsi:
     ```
     biaya = maxPrice * participants
     ```
5. **Filter Berdasarkan Tanggal**:
   - Data pemesanan dapat difilter menggunakan parameter `startDate` dan `endDate`.
6. **Pagination**:
   - Data dapat dipisahkan per halaman menggunakan parameter `page` dan `limit`.
7. **Output JSON**:
   - Menyediakan output JSON dalam format yang cocok untuk dashboard.

---

## **Teknologi yang Digunakan**
- **Golang**: Bahasa pemrograman utama untuk membangun API.
- **Fiber**: Framework web untuk menangani HTTP request.
- **Resty**: HTTP Client untuk mengambil data dari API eksternal.
- **JSON**: Format data output API.

---

## **Struktur Proyek**

```
practical/
├── booking.go             # File untuk mengambil data booking dari API eksternal
├── konsumsi.go            # File untuk mengambil data master konsumsi dari API eksternal
├── main.go                # File utama untuk menjalankan server dan endpoint API
├── go.mod                 # File konfigurasi modul Golang
```

---

## **Cara Menjalankan Proyek**

### **1. Persiapan**
1. Pastikan **Golang** sudah terinstal di komputer Anda.  
   - Unduh Golang di: [https://golang.org/dl](https://golang.org/dl)
2. Pastikan koneksi internet aktif untuk mengambil data dari API eksternal.

### **2. Langkah Menjalankan**
1. Clone repository proyek ini:
   ```bash
   git clone https://github.com/Virman480/PracticalTestBE.git
   cd practical
   ```

2. Install dependensi:
   ```bash
   go mod tidy
   ```

3. Jalankan server:
   ```bash
   go run .
   ```

4. Server akan berjalan di:
   ```
   http://localhost:3000
   ```

---

## **Endpoint API**

### **1. Endpoint: `/dashboard`**
**Deskripsi**: Menampilkan data pemesanan ruangan dan detail konsumsi.

- **Method**: `GET`  
- **URL**: `http://localhost:3000/dashboard`

#### **Parameter Query**:
1. **`startDate`** (opsional): Filter data pemesanan berdasarkan tanggal mulai (format: `YYYY-MM-DD`).
2. **`endDate`** (opsional): Filter data pemesanan berdasarkan tanggal akhir (format: `YYYY-MM-DD`).
3. **`page`** (opsional): Nomor halaman untuk pagination (default: `1`).
4. **`limit`** (opsional): Jumlah data per halaman untuk pagination (default: `10`).

#### **Contoh Request**:
- **Tanpa Filter atau Pagination**:
  ```
  GET http://localhost:3000/dashboard
  ```
- **Dengan Filter Tanggal**:
  ```
  GET http://localhost:3000/dashboard?startDate=2024-01-01&endDate=2024-01-05
  ```
- **Dengan Pagination**:
  ```
  GET http://localhost:3000/dashboard?page=1&limit=5
  ```

#### **Contoh Respons**:
```json
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
```

---

## **Penjelasan Logika API**

1. **Pengambilan Data**:
   - Data **booking** diambil dari URL berikut:
     ```
     https://66876cc30bc7155dc017a662.mockapi.io/api/dummy-data/bookingList
     ```
   - Data **master konsumsi** diambil dari URL berikut:
     ```
     https://6686cb5583c983911b03a7f3.mockapi.io/api/dummy-data/masterJenisKonsumsi
     ```

2. **Penggabungan Data**:
   - Untuk setiap ruangan dalam booking, data konsumsi dicocokkan dengan data master konsumsi berdasarkan nama konsumsi.

3. **Perhitungan Biaya**:
   - Biaya konsumsi dihitung menggunakan formula:
     ```
     cost = maxPrice * participants
     ```

4. **Filter Tanggal**:
   - Pemesanan difilter berdasarkan tanggal `bookingDate` menggunakan parameter `startDate` dan `endDate`.

5. **Pagination**:
   - Data dibagi per halaman berdasarkan parameter `page` dan `limit`.

---

## **Struktur Data**

### **1. Data Pemesanan Ruangan (Booking)**
**Contoh Data dari API Booking**:
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

### **2. Data Master Konsumsi**
**Contoh Data dari API Master Konsumsi**:
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

## **Pengembangan Selanjutnya**
1. **Unit Testing**:
   - Gunakan library seperti `testify` untuk menulis unit test pada setiap fungsi.
2. **Error Handling**:
   - Tambahkan penanganan error lebih baik untuk skenario seperti timeout, respons kosong, atau data tidak valid.
3. **Fitur Pencarian**:
   - Tambahkan fitur pencarian ruangan berdasarkan nama ruangan atau kantor.

---
