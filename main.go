package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	f "github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "goFaunaStore/docs" // docs is generated by Swag CLI, you have to import it.
)

var dbClient *f.FaunaClient


// @title Shop API
// @version 1.0
// @description This is a sample shop service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
	secret := os.Getenv("FAUNA_SECRET_KEY")
	customHeader := map[string]string{"X-Fauna-Source": "fauna-shopapp-golang"}

	dbClient = f.NewFaunaClient(secret, f.Headers(customHeader))

	router := mux.NewRouter()
	router.HandleFunc("/categories", getCategories).Methods("GET")
	router.HandleFunc("/categories", createCategories).Methods("POST")
	router.HandleFunc("/product", createProduct).Methods("POST")
	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

type Product struct {
	Name           string `fauna:"name"`
	Price          string `fauna:"price"`
	Weight         int    `fauna:"weight"`
	Sku            string `fauna:"sku"`
	CategoriesRefs []string
	DateTime       time.Time
}

type Categories struct {
	Name string `fauna:"name"`
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body Product true "Create product"
// @Success 200 {object} Product
// @Router /product [post]
func createProduct(w http.ResponseWriter, r *http.Request) {

	var product Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	product.DateTime = time.Now()
	w.Header().Set("Content-Type", "application/json")
	_, err := dbClient.Query(
		f.Create(
			f.Collection("products"),
			f.Obj{"data": f.Obj{
				"name":       product.Name,
				"categories": f.Map(
					product.CategoriesRefs,
					f.Lambda("c", f.Ref(f.Collection("categories"), f.Var("c")))),
			}}))
	if err != nil {
		panic(err)
	}
	respondWithJSON(w, http.StatusCreated, product)
}

// CreateCategories godoc
// @Summary Create a new categories
// @Description Create new categories with the input payload
// @Tags categories
// @Accept  json
// @Produce  json
// @Param categories body Categories true "Create categories"
// @Success 200 {object} Categories
// @Router /categories [post]
func createCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var categories Categories
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&categories); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err := dbClient.Query(
		f.Create(f.Collection("categories"), f.Obj{"data": &categories}))
	if err != nil {
		panic(err)
	}
	respondWithJSON(w, http.StatusCreated, categories)
}

// GetCategories godoc
// @Summary Get list of all categories
// @Description Get details of all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} Categories
// @Router /categories [get]
func getCategories(w http.ResponseWriter, r *http.Request) {
	type Categories struct {
		Categories f.Value
	}

	w.Header().Set("Content-Type", "application/json")

	res, err := dbClient.Query(
		f.SelectAll(f.Arr{"data", "data", "name"},
			f.Map(
				f.Paginate(f.Documents(f.Collection("categories"))),
				f.Lambda(
					"categoriesRef",
					f.Let().
						Bind("categories", f.Get(f.Var("categoriesRef"))).
						In(f.Var("categories"))))))

	if err != nil {
		panic(err)
	}

	categories := Categories{
		Categories: res,
	}

	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		panic(err)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

