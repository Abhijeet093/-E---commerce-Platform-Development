import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the product into the database and handle any errors

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products", CreateProductHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
