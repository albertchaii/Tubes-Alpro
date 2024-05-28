package main

import "fmt"

type Tenant struct {
	ID              int
	nama            string
	nTransaksi      int
	TotalPendapatan float64
}

const NMAX = 10

type daftarTenant [NMAX]Tenant

var indexTenant int = 0
var nextTenantID int = 1

// function tambah tenant
func tambahTenant(A *daftarTenant) {
	if indexTenant <= NMAX {
		fmt.Print("Tambah tenant baru: ")
		fmt.Scan(&A[indexTenant].nama)
		A[indexTenant].ID = nextTenantID // Assign array tenant dengan nextTenantID untuk menentukan ID tenant yang unik
		nextTenantID++                   // Increment ID tanant
		indexTenant++                    // Increment indeks array tenant
		fmt.Println(" ")
		fmt.Println("Tenant berhasil ditambahkan")
	}

}

//Indek search untuk hapus data
func seqSearch(A daftarTenant, hapusID int) int {

	for i := 0; i < indexTenant; i++ {
		if A[i].ID == hapusID {
			return i
		}
	}
	return -1
}

//function untuk menghapus tenant
func hapusTenant(A *daftarTenant, hapusID int) {
	fmt.Print("Masukan id tenant yang akan dihapus: ")
	fmt.Scan(&hapusID)
	index := seqSearch(*A, hapusID) // sequantial search digunakan untuk mencari indeks tenant yang akan dihapus

	if index == -1 {
		fmt.Println("Tenant tidak ditemukan")
	} else if A[index].ID == hapusID {
		for i := index; i < indexTenant-1; i++ {
			A[i] = A[i+1]
		}
		indexTenant--
		fmt.Println(" ")
		fmt.Println("Tenant berhasil dihapus")
	}

}

// index search untuk mengedit data dan mencatat transaksi
func binarySearch(A daftarTenant, editTenantID int) int {
	left := 0
	right := indexTenant - 1
	found := -1

	for left <= right && found == -1 {
		mid := (left + right) / 2
		if editTenantID < A[mid].ID {
			right = mid - 1
		} else if editTenantID > A[mid].ID {
			left = mid + 1
		} else {
			found = mid // Tenant found
		}
	}
	return found
}

//function untuk mengubah nama tenant berdasarkan id tenant
func editTenant(A *daftarTenant, editTenantID int, newName string) {

	fmt.Print("Masukan id tenant yang akan diedit: ")
	fmt.Scan(&editTenantID)
	found := binarySearch(*A, editTenantID) // binary search digunakan untuk mencari indeks tenant yang akan diedit

	if found == -1 {
		fmt.Println(" ")
		fmt.Println("Id tenant tidak ditemukan!")
	} else {
		fmt.Print("Masukan nama tenant baru: ")
		fmt.Scan(&newName)
		(*A)[found].nama = newName
		fmt.Println(" ")
		fmt.Println("Tenant berhasil diubah")
	}

}

// function untuk mencatat transaksi yang dilakukan oleh setiap tenant
func catatTransaksi(A *daftarTenant, tenantID int, KomisiAdmin *float64) {
	var jumlah float64
	fmt.Print("Masukan Id tenant yang akan dicatat transaksinya: ")
	fmt.Scan(&tenantID)
	fmt.Print("Masukan jumlah transaksi: ")
	fmt.Scan(&jumlah)
	found := binarySearch(*A, tenantID) // binary search digunakan untuk mencari indeks tenant yang akan dicatat transaksi
	if found == -1 {
		fmt.Println("Tenant tidak ditemukan!")
	} else {
		(*A)[found].nTransaksi++
		(*A)[found].TotalPendapatan += jumlah
		*KomisiAdmin += jumlah * 0.25
		fmt.Println(" ")
		fmt.Println("Transaksi berhasil dicatat")
	}
}

// Sorting secara descending
func selectionSort(A *daftarTenant) {
	for i := 0; i < indexTenant-1; i++ {
		maxIdx := i
		for j := i + 1; j < indexTenant; j++ {
			if (*A)[j].nTransaksi > (*A)[maxIdx].nTransaksi {
				maxIdx = j
			}
		}
		// Swap the entire structs
		temp := (*A)[i]
		(*A)[i] = (*A)[maxIdx]
		(*A)[maxIdx] = temp
	}
}

// Sorting secara Ascending
func InsertionSort(A *daftarTenant) {
	for pass := 1; pass < indexTenant; pass++ {
		temp := (*A)[pass]
		i := pass
		for i > 0 && (*A)[i-1].nTransaksi > temp.nTransaksi {
			(*A)[i] = (*A)[i-1]
			i--
		}
		(*A)[i] = temp
	}
}

//function untuk print data tenant
func cetakTenant(A daftarTenant, KomisiAdmin float64) {

	var pilihsort int

	fmt.Println("1. Urutkan berdasarkan jumlah transaksi (Descending)")
	fmt.Println("2. Urutkan berdasarkan jumlah transaksi (Ascending)")
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&pilihsort)
	fmt.Println(" ")

	if pilihsort == 1 {
		selectionSort(&A)
	} else if pilihsort == 2 {
		InsertionSort(&A)
	} else {
		fmt.Println(" ")
		fmt.Println("Pilihan tidak valid !")
	}

	for i := 0; i < indexTenant; i++ {
		fmt.Printf("ID: %d, Nama: %s, Jumlah Transaksi: %d, Jumlah Uang: %.2f\n",
			A[i].ID, A[i].nama, A[i].nTransaksi, A[i].TotalPendapatan)
	}
	fmt.Printf("Total uang yang diperoleh admin: %.2f\n", KomisiAdmin)
}

func main() {

	var tenant daftarTenant
	var Name string
	var pilihan int
	var id int
	var komisi float64

	for {
		fmt.Println("==========================")
		fmt.Println("    Menu Kantin Tel-U     ")
		fmt.Println("==========================")
		fmt.Println("1. Tambah Tenant")
		fmt.Println("2. Edit Tenant")
		fmt.Println("3. Hapus Tenant")
		fmt.Println("4. Catat Transaksi")
		fmt.Println("5. Tampilkan Tenant")
		fmt.Println("6. Keluar")
		fmt.Println("==========================")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println(" ")

		if pilihan == 1 {
			if indexTenant == NMAX {
				fmt.Println("Data tenant telah penuh.")
			} else {
				tambahTenant(&tenant)
			}
		} else if pilihan == 2 {
			editTenant(&tenant, id, Name)
		} else if pilihan == 3 {
			hapusTenant(&tenant, id)
		} else if pilihan == 4 {
			catatTransaksi(&tenant, id, &komisi)
		} else if pilihan == 5 {
			cetakTenant(tenant, komisi)
		} else if pilihan == 6 {
			break
		} else {
			fmt.Println("Pilihan tidak valid !")
		}
	}
}
