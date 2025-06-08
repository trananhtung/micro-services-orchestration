# Inventory Service

A simple inventory management microservice using NestJS, TypeORM, and SQLite.

## Setup

```bash
npm install
```

## Run the Service

```bash
npm run start:dev
```

The service will run at: http://localhost:3000

## API Endpoints

- **GET** `/inventory` — Get all inventory items
- **GET** `/inventory/:id` — Get item by ID
- **GET** `/inventory/check-stock/:productName` — Check stock by product name
- **POST** `/inventory/update-stock`
  - Body:
    ```json
    {
      "productName": "Product Name",
      "quantity": 100
    }
    ```

## Swagger API Docs

After running the service, visit:
```
http://localhost:3000/api
```
To see and test the API documentation.

## Database Migration (TypeORM)

1. Build the project:
   ```bash
   npm run build
   ```
2. Generate a new migration:
   ```bash
   npx typeorm migration:generate -n MigrationName
   ```
3. Run migrations:
   ```bash
   npx typeorm migration:run
   ```

> Note: For TypeORM v0.3.x+, use a `data-source.ts` config file. See TypeORM docs for details.
