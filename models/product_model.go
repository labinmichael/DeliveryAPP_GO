package models

import (
	"context"
	"errors"
	"log"
	"temp/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ErrProductNotFound is returned when a requested product is not found in the database.
var ErrProductNotFound = errors.New("product not found")

var (
	Db_name         = "my_database"
	Collection_name = "products"
)

// Product represents a product entity in the system.
type Product struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	IsAvailable bool      `json:"isAvailable"`
	ProductName string    `json:"productName"`
	Description string    `json:"description"`
	ActualPrice float64   `json:"actualPrice"`
	OfferPrice  float64   `json:"offerPrice"`
	ProductType string    `json:"productType"`
	ShopID      string    `json:"shopID"`
	CreatedAt   time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

// Access the "products" collection
collection := client.Database(Db_name).Collection(Collection_name)

// GetAllProducts retrieves all products from the database.
func GetAllProducts() ([]Product, error) {
	// Initialize MongoDB client
	client, err := config.InitMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve all products
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var products []Product
	if err := cursor.All(ctx, &products); err != nil {
		log.Fatal(err)
	}

	return products, nil
}

// GetProductByID retrieves a product by its ID.

func GetProductByID(id string) (Product, error) {
	client, err := config.InitMongoClient()
	if err != nil {
		return Product{}, err
	}
	defer client.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Product{}, err
	}

	var product Product
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return Product{}, ErrProductNotFound
		}
		return Product{}, err
	}

	return product, nil
}

// CreateProduct creates a new product.
func CreateProduct(newProduct Product) (string, error) {
	// Initialize MongoDB client
	client, err := config.InitMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())


	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set creation and update timestamps
	newProduct.CreatedAt = time.Now()
	newProduct.UpdatedAt = time.Now()

	// Insert new product into the collection
	result, err := collection.InsertOne(ctx, newProduct)
	if err != nil {
		log.Fatal(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// UpdateProduct updates an existing product.
func UpdateProduct(id string, updatedProduct Product) error {
	// Initialize MongoDB client
	client, err := config.InitMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert ID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	// Set update timestamp
	updatedProduct.UpdatedAt = time.Now()

	// Update product in the collection
	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": updatedProduct})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// DeleteProduct deletes a product by its ID.
func DeleteProduct(id string) error {
	// Initialize MongoDB client
	client, err := config.InitMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert ID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	// Delete product from the collection
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
