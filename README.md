

# 🚀 DevOps Intern Assignment – Nginx Reverse Proxy + Docker Compose

This project demonstrates a microservices architecture using **Docker Compose**, with a reverse proxy configured through **Nginx**, routing requests to two backend services: one written in **Golang** and the other in **Python (Flask)**.

# Watch Demo 
🎥 [Watch Demo Video](https://youtu.be/XRF6bJpgLtc)


---

## 📁 Project Structure

```bash
.
├── docker-compose.yml
├── nginx
│   ├── Dockerfile
│   └── nginx.conf
├── service_1
│   ├── Dockerfile
│   └── main.go
├── service_2
│   ├── Dockerfile
│   ├── app.py
│   └── requirements.txt
└── README.md
````

---

## 🎯 Features

* ✅ **Nginx Reverse Proxy** running inside Docker
* ✅ Routes:

  * `/service1/*` → Go backend (Service 1)
  * `/service2/*` → Python/Flask backend (Service 2)
* ✅ Single accessible port: `localhost:8080`
* ✅ Real-time request logging
* ✅ Health checks for all services
* ✅ Bridge networking

---

## 🔧 How to Run

> Prerequisites:
>
> * Docker 🐳
> * Docker Compose

```bash
# Clone the repo
git clone https://github.com/aniketpuro/docker_-_ngnix_reverse_proxy.git
cd docker_&_ngnix_reverse_proxy

# Run the project
docker-compose down --remove-orphans && docker-compose up --build
```

📸 *Screenshot of terminal with containers starting successfully:*

![Containers Running](https://github.com/user-attachments/assets/c32072e7-067c-4b33-81b0-93a144d0ceee)


---

## 🌐 Access Services

| Route                                  | Description             | Example Output                       |
| -------------------------------------- | ----------------------- | ------------------------------------ |
| `http://localhost:8080/service1/hello` | Routes to Go service    | `{"message":"Hello from Service 1"}` |
| `http://localhost:8080/service2/hello` | Routes to Flask service | `{"message":"Hello from Service 2"}` |
| `http://localhost:8080/service1/ping`  | Healthcheck endpoint    | `{"status":"ok", "service":"1"}`     |
| `http://localhost:8080/service2/ping`  | Healthcheck endpoint    | `{"status":"ok", "service":"2"}`     |

📸 *Screenshot of browser hitting `/service2/hello`:*

![Service 2 Hello](https://github.com/user-attachments/assets/b72ce3e9-180d-4d25-a179-2a0e24107425)


---

## 🧠 Explanation of Routing

Nginx listens on port `8080` and routes based on path prefixes:

```nginx
location /service1/ {
    proxy_pass http://service1/;
}

location /service2/ {
    proxy_pass http://service2/;
}
```

* `service1` and `service2` are **Docker service names**.
* It uses the `proxy_pass` directive to forward traffic inside the Docker network.
* Access logs include IP, time, path, and duration.

📸 *Screenshot of `nginx.conf`:*

![Nginx Config](https://github.com/user-attachments/assets/02917a55-14c1-486f-af41-c6ca9ff79fae)


---

## 🩺 Health Checks

Each service includes a Docker-level health check:

* **Go Service (Service 1):**

  ```Dockerfile
  HEALTHCHECK CMD curl --fail http://localhost:8001/ping || exit 1
  ```

* **Flask Service (Service 2):**

  ```Dockerfile
  HEALTHCHECK CMD curl --fail http://localhost:8002/ping || exit 1
  ```

* **Nginx:**

  ```Dockerfile
  HEALTHCHECK CMD curl --fail http://localhost || exit 1
  ```

📸 *Screenshot of `docker ps` showing health status:*

![Docker Health](https://github.com/user-attachments/assets/d1db70a0-cca6-474a-b1e3-68a055ef7e74)


---

## 📝 Logs Preview

All services log incoming requests:

**Service 1 (Go)**

```text
➡️  /hello hit from 172.19.0.1:43792
```

**Service 2 (Flask)**

```text
➡️  /hello hit from 172.19.0.1
```

**Nginx Access Log**

```text
172.19.0.1 - - [24/Jun/2025:14:22:46 +0000] "GET /service1/hello HTTP/1.1" 200
```

📸 *Screenshot of terminal logs:*

![Logs](https://github.com/user-attachments/assets/ab55fa3d-1dcf-4129-817c-526c5aa1cf00)


---

## ✅ Bonus Features Implemented

* [x] Docker Health Checks
* [x] Real-time logging in both Go and Python
* [x] Nginx running inside Docker with custom Dockerfile
* [x] Bridge networking (default in Docker Compose)

---

## 💬 Final Thoughts

This project helped me practically implement reverse proxying, container health management, and inter-service communication using Docker networking. I’m excited to learn more and contribute in real-world DevOps environments.

---

## 🙏 Thank You
