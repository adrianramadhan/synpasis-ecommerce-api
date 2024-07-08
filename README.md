# E-Commerce Backend API

This project is a RESTful API for an e-commerce platform, developed as part of the backend internship technical test for PT Synapsis Sinergi Digital.

## Features

- Product listing by category
- Shopping cart management (add, view, delete items)
- Checkout and payment processing
- User authentication (login and registration)

## Tech Stack

- Go (Golang)
- Echo framework
- GORM (ORM)
- PostgreSQL
- Docker

## Getting Started

### Prerequisites

- Go 1.16+
- Docker and Docker Compose
- PostgreSQL (if running locally)

### Installation

1. Clone the repository:
```
git clone https://github.com/adrianramadhan/synpasis-ecommerce-api
```

2. Navigate to the project directory:
```
cd synpasis-ecommerce-api
```
3. Install dependencies:
```
go mod tidy
```
4. Set up environment variables:
```
cp .env.example .env
```
5. Migrate Database Schema:
```
go run main.go migrate
```
6. Start REST Server
```
go run main.go rest
```

### Running with Docker
1. Build and run the containers:
```
docker-compose up --build

```
2. The API will be available at `http://localhost:8080`

## Database Schema
![synapsis-ecommerce-erd](https://github.com/adrianramadhan/synpasis-ecommerce-api/assets/59206760/5782b168-447f-412a-ad1e-dd875e9640ad)
Entity Relationships:
User - Cart: One-to-Many (one user can have multiple shopping carts)
User - Order: One-to-Many (one user can create multiple orders)
Category - Product: One-to-Many (one category can have multiple products)
Cart - CartItem: One-to-Many (one cart can have multiple items)
Product - CartItem: One-to-Many (one product can be in multiple cart items)
Cart - Order: One-to-One (one cart can become one order)
Order - Payment: One-to-One (one order has one payment)

Business Process Flow:
User registers and logs into the application.
User views Products based on Categories.
User adds Products to the Cart, creating CartItems.
User can view and edit CartItems within their Cart.
When the User checks out, the Cart is converted into an Order.
User makes a Payment for that Order.
Order and Payment statuses are updated according to the transaction progress

## Docker Image

The Docker image for this project is available on Docker Hub:
```
https://hub.docker.com/r/adrianramadhan/ecommerce-api/tags
```

## Contributing
This project is for demonstration and technical test purposes. However, feedback and suggestions are welcome.

## Contact
Adrian Putra Ramadhan - adrianramadhan881@gmail.com
