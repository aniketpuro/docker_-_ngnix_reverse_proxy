FROM python:3.11-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
RUN apt-get update && apt-get install -y curl

HEALTHCHECK --interval=10s --timeout=2s --start-period=1s --retries=2 \
  CMD curl --fail http://localhost:8002/ping || exit 1

CMD ["python", "app.py"]
