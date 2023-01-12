
# Product Catalog and Shopping Cart Service API

This Project is to create the APIs which are used to process the shopping cart service using CRUD Operations.
The API returns the details of the products into the shopping cart.


## Getting started
### Prerequisites

- Go 1.19.3 (should still be backwards compatible with earlier versions)
- Postman-The collections will need an environment setup with `scheme`, `port` and `host` variables setup with values of `http`, `8000` and `localhost` respectively.
- PostgreSql database
## Environment Setup

Command to start working with MySQL
```bash
$ mysql -u root -p
```
Command to run golang code
```bash
$ go run main.go (Run the code using the following command)
```
Commands to upload code to github

```bash
git init -b main (Initialize the local directory as a Git repository)
```
```bash
git add . && git commit -m "initial commit"(Stage and commit all the files in your project)
```
## API Reference

### Products
#### Request Body
##### FORM-VALUE:
- `id` (int): ID of the specific product.
- `name` (string): The name of the product.
- `specification` (string): The Specifications of the product.
- `sku` (string): Stock Keeping Unit number of the product.
- `category_id` (string): The product's category ID which needs to be present in the Category Table.
- `price` (int): The price of the product.

#### Insert Product
##### Adds a product to database

```http
  POST /create
```

#### Get Products
##### Recieve all products 

```http
  GET /products
```

#### Get Product
##### Return a Specific Product using name of the product

```http
  GET /product/{name}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of item to fetch |


#### Update Product
##### Updates product detail of item depending on product_id
```http
  PUT /updateproduct
```
#### Request Body
##### FORM-VALUE:
- `name` (string): The name of the product.
- `specification` (string): The Specifications of the product.
- `sku` (string): Stock Keeping Unit number of the product.
- `category_id` (string): The product's category ID which needs to be present in the Category Table.
- `price` (int): The price of the product.


#### Delete Product
##### Delete product details based on productname

```http
  Delete /deleteproduct
```

#### Request Body
##### FORM-VALUE:
`product` (string): The name of the product.


### Category_master

#### Insert Category Name

```http
  POST /createcategory
```
#### Request Body
##### FORM-VALUE:
- `name` (string): Category name of the corresponding product

#### Get All Categories

```http
  GET /categories
```
#### Get Category
##### Return a Specific Category using name of the category

```http
  GET /getcategory/{name}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of item to fetch |


#### Delete Category

```http
  Delete /deletecategory
```

#### Request Body
##### FORM-VALUE:
- `name` (string): Category name of the corresponding product

### Inventory

#### Insert inventory details

```http
  POST /createinventory
```

#### Request Body
##### FORM-VALUE:
- `product` (string): Name of specific product.
- `quantity` (int): Quantity of the specific product present inside the inventory.

#### Get Inventory details

```http
  GET /productsinventory
```

#### Get Single Product Inventory details

```http
  GET /productsinventory
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of item to fetch |

#### Update inventory detail

```http
  PUT /updateproductinventory
```
#### Request Body
##### FORM-VALUE:
- `product` (string): Name of specific product.
- `quantity` (int): Quantity of the specific product present inside the inventory.

#### Delete inventory Detail

```http
  Delete /deleteproductinventory
```
#### Request Body
##### FORM-VALUE:
- `product` (string): Name of specific product.


### Cart 

#### Get cart details

```http
  GET /getcart  
```

#### Insert Product items to cart

```http
  POST /addcart  (To add required product to cart)
```
#### Request Body
##### FORM-VALUE:
- `product` (string): Name of specific product.
- `quantity` (int): Quantity of the product 

#### Delete product item from cart
```http
  DELETE /deletecart  (Delete specific product from cart)
```
#### Request Body
##### FORM-VALUE:
- `product` (string): Name of specific product.

## List of packages

- encoding/json -Package json implements encoding and decoding of JSON as defined in RFC 7159.
- fmt-Package fmt implements formatted I/O with functions analogous to C's printf and scanf. 
- io/ioutil-Package ioutil implements some I/O utility functions.
- net/http-Package http provides HTTP client and server implementations.
- gorilla mux implements a request router and dispatcher for matching incoming requests to their respective handler.