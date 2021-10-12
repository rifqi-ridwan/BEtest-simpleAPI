# Backend Test Simple REST API
This is a simple REST API project for Learning purpose

## Prerequisite
- Go `1.16`
- MySQL `5.7`

### Database Migration

1. Migrate database
```bash
go run db/migrate.go migrate
```

## How to Run

1. Install Go, [read the docs](https://golang.org/doc/install)

2. Install MySQL, [read the docs](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en)

3. Install Go dependencies
```bash
go mod download
```

4. Copy environments (then edit as you need, e.g. `MYSQL_PASSWORD`)
```bash
cp .env.example .env
```

5. Run server
```bash
go run main.go
```

## How to use

### Routing

```
GET    /product
GET    /product/:id
GET    /product?name=productname
GET    /product?category=categoryname
POST   /product/create
PATCH  /product/update
DELETE /product/:id
```

### Data structure
#### Product
```json
{
  "id": 1,
  "name": "test product 1",
  "price": 10000,
  "category_id": 1,
  "image": ""
}
```

#### Category
Once you do the migration it will create 2 category, so you can ignore it..
