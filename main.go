package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 5000, Stok: 10},
	{ID: 2, Nama: "Popmie", Harga: 10000, Stok: 30},
	{ID: 3, Nama: "Produk C", Harga: 15000, Stok: 20},
}

func getProdukByID(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		for _, p := range produk {
			if p.ID == id {
				json.NewEncoder(w).Encode(p)
				w.Header().Set("Content-Type", "application/json")
				return
			}
		}

		http.Error(w, "Produk Belum Ada", http.StatusNotFound)
}

func updateProduk(w http.ResponseWriter, r *http.Request) {
    //PUT localhost:8080/api/produk/{id}
    idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
			http.Error(w, "Invalid Produk  ID", http.StatusBadRequest)
			return
		}
    var updatedProduk Produk
    err = json.NewDecoder(r.Body).Decode(&updatedProduk)
    if err != nil {
        http.Error(w, "Invalid Request", http.StatusBadRequest)
        return
    }
    for i := range produk {
        if produk[i].ID == id {
            updatedProduk.ID = id
            produk[i] = updatedProduk
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updatedProduk)
            return
        }
    }
    http.Error(w, "Produk Belum Ada", http.StatusNotFound)
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
    //DELETE localhost:8080/api/produk/{id}
    idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
            http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
            return
        }
    for i,p  := range produk {
        if p .ID == id {
            produk = append(produk[:i], produk[i+1:]...)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(map[string]string{
                "message": "Sukses Delete",
            })

            return
        }
    }
    http.Error(w, "Produk Belum Ada", http.StatusNotFound)
}


func main() {
	//GET localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            getProdukByID(w, r)
        } else if r.Method == "PUT" {
            updateProduk(w, r)
        } else if r.Method == "DELETE" {
            deleteProduk(w, r)
        }
	})
	//GET localhost:8080/api/produk
	//POST localhost:8080/api/produk
	//DELETE localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}

			//masukin data ke dalam variable produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)
		}
	})
	//localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "Ok",
			"message": "API Running",
		})
	})
	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start")
	}
}
