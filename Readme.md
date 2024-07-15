    Delivery App API Documentation

1.      GET All Products
   Endpoint: GET /products
   Description: Retrieve a list of all products.
   Example: GET http://localhost:8080/products

2.     Create New Product
   Endpoint: POST /product
   Description: Create a new product.
   Example:
   POST http://localhost:8080/product
   
   Body (JSON):

            {
               "isAvailable": true,
               "productName": "New Product",
               "description": "Description of the new product",
               "actualPrice": 50.00,
               "offerPrice": 40.00,
               "productType": "New Type",
               "shopID": "123456"
            }
            

3.     Get Product by ID
   Endpoint: GET /product/{id}
   Description: Retrieve a specific product by its ID.
   Example: GET http://localhost:8080/product/65ee8c9267930b8b07c4fdef

4.     Delete Product
   Endpoint: DELETE /products/{id}
   Description: Delete a product by its ID.
   Example: DELETE http://localhost:8080/products/65ec1a342ec45b9ddf2bd5c1

5.     Update Product
   Endpoint: PUT /products/{id}
   Description: Update a product by its ID.
   Example:
   PUT http://localhost:8080/products/65ec1a342ec45b9ddf2bd5c1
   
   Body (JSON):
   
         {
            "isAvailable": false,
            "productName": "Updated Product Name",
            "description": "Updated Description",
            "actualPrice": 150,
            "offerPrice": 120,
            "productType": "Updated Product Type",
            "shopID": "Updated Shop ID"
         }


