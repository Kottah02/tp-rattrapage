package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Product structure
type Product struct {
	ID          int
	Name        string
	ImagePath   string
	Price       float64
	SalePrice   float64
	OnSale      bool
	Stock       int
	Description string
	Sizes       []string // Ajout du champ Sizes pour les tailles disponibles
}

// Mocked list of products
var products = []Product{
	{ID: 1, Name: "Palace Pull à Capuche Unisexe Chasseur", ImagePath: "/images/p1.png", Price: 120, SalePrice: 90, OnSale: true, Stock: 20, Description: "Un pull confortable pour toutes les saisons.", Sizes: []string{"XS", "S", "M", "L", "XL"}},
	{ID: 2, Name: "Palace Pull à Capuche Marine", ImagePath: "/images/p2.png", Price: 140, OnSale: false, Stock: 15, Description: "Pull marin de haute qualité.", Sizes: []string{"S", "M", "L", "XL"}},
	{ID: 3, Name: "Palace Pull Crew Passeport Noir", ImagePath: "/images/p3.png", Price: 130, OnSale: false, Stock: 10, Description: "Élégant et moderne, parfait pour toutes les occasions.", Sizes: []string{"XS", "M", "L"}},
	{ID: 4, Name: "Palace Washed Terry 1/4 Pullover Hoodie", ImagePath: "/images/p4.png", Price: 135, Stock: 12, OnSale: false, Description: "Confortable hoodie délavé pour un look unique.", Sizes: []string{"S", "M", "L", "XL"}},
	{ID: 5, Name: "Palace Pantalon Bossy Jean Bleu", ImagePath: "/images/p5.png", Price: 125, Stock: 15, OnSale: true, SalePrice: 110, Description: "Pantalon en jean de haute qualité."}, // Pas de tailles car ce n'est pas un pull
}

// Fonction pour servir la page d'accueil
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "Erreur d'exécution du template", http.StatusInternalServerError)
	}
}

// Fonction pour servir la page des détails du produit
func productDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère l'ID du produit depuis la requête
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de produit invalide", http.StatusBadRequest)
		return
	}

	// Trouve le produit correspondant à l'ID
	var selectedProduct *Product
	for _, product := range products {
		if product.ID == id {
			selectedProduct = &product
			break
		}
	}

	// Si le produit n'est pas trouvé
	if selectedProduct == nil {
		http.Error(w, "Produit non trouvé", http.StatusNotFound)
		return
	}

	// Chargement du template pour la page de détails
	tmpl, err := template.ParseFiles("templates/product_details.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Exécution du template avec les données du produit sélectionné
	err = tmpl.Execute(w, selectedProduct)
	if err != nil {
		http.Error(w, "Erreur d'exécution du template", http.StatusInternalServerError)
	}
}

func main() {
	// Routage
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/product", productDetailHandler)

	// Serveurs de fichiers statiques
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("public/images"))))

	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
