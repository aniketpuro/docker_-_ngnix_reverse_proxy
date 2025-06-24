# ==========================
# 🔌 nginx/Dockerfile
# ==========================
FROM nginx:1.25-alpine
WORKDIR /etc/nginx

# ✅ Install curl
RUN apk add --no-cache curl

COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80

# ✅ Add healthcheck
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl -f http://localhost || exit 1

CMD echo "[nginx] 🟢 Reverse proxy started on port 8080" && nginx -g 'daemon off;'
