# url-shortener
A url shortener microservice made in Golang using hexagonal service architecture principles. Most of the code is from https://github.com/tensor-programming/hex-microservice as part of his Golang video series on Youtube. This is mostly for learning and reference purposes.

## Getting Started
1. Clone repo
2. `cp .env-example .env`
3. `docker-compose up --build`
4. Test in postman or curl by sending a POST request to `localhost:8000` with `Content-Type: application/json` header set:
```json
{
    "url": "https://github.com/edwintcloud?tab=repositories"
}
```