### A Quick Step-by-Step Guide:

1. Clone this repo.

2. run go run main.go This will start a sever at localhost:8080

3. Using the terminal, you can test the API by running, `curl http://localhost:8000/users`

4. You can send a POST request by, `curl -i -X POST -d '{"name":"Mia","age":25}' http://localhost:8000/users/create`

5. You can even use Postman to test the API.