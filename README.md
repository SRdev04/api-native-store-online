# Golang CRUD API with MySQL

This project is a simple CRUD API developed in Golang, featuring HTTP routing, middleware, and MySQL integration. It provides functionality for managing products, shopping carts, shipping details, and orders.

## Features

- Create, Read, Update, and Delete (CRUD) operations for Products, Cart, Shipping, and Order entities.
- Golang HTTP router for handling API endpoints.
- Middleware for additional functionalities such as API key validation.
- MySQL database for persistent storage.

## Requirements

- [Golang](https://golang.org/) installed on your machine.
- MySQL server for database storage.
- Necessary Golang packages and dependencies, which can be installed using `go get`.

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/SRdev04/api-native-store-online
    ```

2. Navigate to the project directory:

    ```bash
    cd your_project
    ```

3. Install dependencies:

    ```bash
    go get -u ./...
    ```

4. Set up the MySQL database by creating the necessary tables. You can find the SQL scripts in the `database` directory.

5. Configure the database connection in the `config/config.go` file.

6. Run the application:

    ```bash
    go run main.go
    ```

7. Access the API at [http://localhost:3000](http://localhost:3000).

## API Endpoints

- Products:
  - `GET /products`: Retrieve all products.
  - `GET /products/{id}`: Retrieve a specific product by ID.
  - `POST /products`: Create a new product.
  - `PUT /products/{id}`: Update a product by ID.
  - `DELETE /products/{id}`: Delete a product by ID.

- Cart, Shipping, and Order:
  - Similar CRUD endpoints for Cart, Shipping, and Order entities.

## Middleware

The project includes a basic middleware for API key validation. Modify the middleware logic as needed in `main.go`.

## Contributing

Feel free to contribute to the project by opening issues or submitting pull requests. Your feedback and contributions are highly appreciated.

## License

This project is licensed under the [MIT License](LICENSE).
