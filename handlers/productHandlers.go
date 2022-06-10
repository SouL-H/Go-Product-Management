package handlers

import (
	. "Product/helpers"
	. "Product/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var productData = make(map[string]Product)
var flag int = 0

//HTTP POST - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	CheckErr(err)

	product.CreatedOn = time.Now()
	flag++
	product.ID = flag
	key := strconv.Itoa(flag)
	productData[key] = product
	x, err := json.Marshal(product)
	CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //İslem basarili
	w.Write(x)

}

//HTTP GET - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, v := range productData {
		products = append(products, v)
	}

	x, err := json.Marshal(products)
	CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //İslem basarili
	w.Write(x)

}

//HTTP GET - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for _, data := range productData {
		if data.ID == key {
			product = data
		}
	}
	x, err := json.Marshal(product)
	CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //İslem basarili
	w.Write(x)

}

//HTTP PUT - /api/productss{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {

	var err error
	vars := mux.Vars(r)
	key := vars["id"]
	var proUpd Product
	err = json.NewDecoder(r.Body).Decode(&proUpd)
	CheckErr(err)
	if _, data := productData[key]; data { //data var ise
		proUpd.ID, _ = strconv.Atoi(key)
		proUpd.ChangedOn = time.Now()
		delete(productData, key)  //Old data sil
		productData[key] = proUpd //New data ekle

	} else {
		log.Printf("Değer bulunamadı: %s", key)

	}
	w.WriteHeader(http.StatusOK) //İşlem başarılı
}

//HTTP DELETE - /api/productss/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, found := productData[key]; found {
		delete(productData, key)

	} else {
		log.Printf("Değer bulunamadı: %s", key)

	}
	w.WriteHeader(http.StatusOK) //İşlem başarılı

}
