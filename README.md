# Coupon Service API (POC)

## Description

This project implements a RESTful API service for managing coupons. It provides endpoints for creating, retrieving, updating, and deleting coupons. Built with Go, it leverages the power of the `mux` router and includes unit tests to ensure reliability.

NOTE: This repo is just a POC not a complete project so please be kind.

## Installation

To install and run this project, follow these steps:

1. If we have taskfile installed just run ```task build_run```
2. Manually build and run Dockerfile (check Taskfile.yaml for example)
3. Compile manually and run compiled binary *(probably the least cool)*

## API Endpoints

This service provides several RESTful endpoints for managing coupons and applying discounts. Below are the available endpoints with examples on how to use them.

### Health Check

- **Endpoint**: `GET /api/v1/`
- **Description**: Checks the health of the service.
- **Example**:
```bash
curl -X GET http://localhost:8080/api/v1/
```

### Save Coupon
- **Endpoint**: `GET /api/v1/coupon/save`
- **Description**: Saves a new coupon.
- **Payload Example**:
```json
{
  "id": "283b8e96-a5f4-4389-a82f-5ec46399de37",
  "code": "NEWYEAR",
  "discount": 10,
  "expiry": "2023-12-31"
}
```
- **Example**:
```bash
curl -X POST http://localhost:8080/api/v1/coupon/save \
     -H "Content-Type: application/json" \
     -d '{"id": "283b8e96-a5f4-4389-a82f-5ec46399de37", "code": "NEWYEAR", "discount": 10, "expiry": "2023-12-31"}'
```

### Find All Coupons
- **Endpoint**: `GET /api/v1/coupon/findAll`
- **Description**: Retrieves all available coupons.
- **Example**:
```bash
curl -X GET http://localhost:8080/api/v1/coupon/findAll
```

### Apply Discount
- **Endpoint**: `POST /api/v1/coupon/apply-discount`
- **Description**: Applies a discount code to a basket and returns the discounted value.
- **Payload Example**:
```json
{
  "code": "NEWYEAR",
  "basketValue": 100
}
```
- **Example**:
```bash
curl -X POST http://localhost:8080/api/v1/coupon/apply-discount \
     -H "Content-Type: application/json" \
     -d '{"code": "NEWYEAR", "basketValue": 100}'
```

Each of these endpoints serves a specific function within the coupon management system, from health checks to creating, listing, and applying coupons. Ensure your server is running on localhost:8080 or adjust the URL in the examples accordingly.
