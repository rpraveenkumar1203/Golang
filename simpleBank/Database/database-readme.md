Here's a nicely formatted `README.md` file based on the steps you've provided for installing and running PostgreSQL using Docker:

---

# 🚀 PostgreSQL Setup with Docker

This guide provides a step-by-step tutorial to download and run PostgreSQL 12 using Docker on your system.

## 🐳 Docker Installation

Ensure Docker is installed on your machine. If you don’t have Docker, you can install it from the official website:

- [Docker Installation Guide](https://docs.docker.com/get-docker/)

## 📦 Download PostgreSQL Image

To get started, pull the PostgreSQL 12 Alpine image from Docker Hub using the following command:

```bash
docker pull postgres:12-alpine
```

### Output:
```bash
12-alpine: Pulling from library/postgres
43c4264eed91: Pull complete 
e7368e03b632: Pull complete 
... [truncated for brevity]
Status: Downloaded newer image for postgres:12-alpine
docker.io/library/postgres:12-alpine
```

## 🔧 Initialize PostgreSQL Container

Run the following command to initialize the PostgreSQL container with environment variables for user and password. We will expose the default PostgreSQL port (5432).

```bash
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

### Common Commands

- **Check running containers:**

    ```bash
    docker ps
    ```

- **List Docker images:**

    ```bash
    docker images
    ```

- **View container logs:**

    ```bash
    docker logs postgres12
    ```

## 🗄️ Accessing PostgreSQL

To access the PostgreSQL database using the terminal, use the following command to connect via `psql`:

```bash
docker exec -it postgres12 psql -U root
```

Once inside the `psql` shell, you can run SQL commands such as:

```sql
SELECT NOW();
```

### Output:
```bash
              now              
-------------------------------
 2024-09-23 17:17:02.765667+00
(1 row)
```

## 📝 Additional Notes

- Ensure that port `5432` is available on your system or adjust the port mapping accordingly.
- You can stop or start the container using:

    ```bash
    docker stop postgres12
    docker start postgres12
    ```

## 📚 Resources

- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Docker PostgreSQL Image](https://hub.docker.com/_/postgres)

---

