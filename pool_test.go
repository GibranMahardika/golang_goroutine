package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		}, // membuat membuat function yang return nya interface kosong untuk default value, yang pada dijalankan itu akan keluar New bukan nul
	}
	group := &sync.WaitGroup{}

	pool.Put("Gibran")
	pool.Put("Gibs")
	pool.Put("aselole")
	pool.Put("kang_seblak")
	pool.Put("abang_tamvan")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get() // ini untuk mengambil data yang ada di pool PUT kemudian ditaro di variable data
			fmt.Println(data)  // ini untuk print variable data yang terdapat pool Get
			group.Wait()       // ini akan menunggu proses seluruh data yang ada di PUT
			// time.Sleep(1 * time.Second) // syntax ini sama kegunaannya dengan sync.Waitgroup, jadi pilih salah satu saja, atau proses akan lebih lama
			pool.Put(data) // lalu dikembalikan lagi kedalam variable data1
		}()
	}

	group.Wait()
	// time.Sleep(11 * time.Second)
	fmt.Println("Selesai")

	// Hasil dari ini akan random penempatannya, tetapi semua data akan ditampilkan, data yang ditampilkan berupa yang ada diatas.
}
