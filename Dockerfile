FROM alpine:latest

# Install Caddy
RUN apk add --no-cache caddy

# Copy Caddyfile and content
COPY Caddyfile /etc/caddy/Caddyfile
COPY gh /srv

# Expose the necessary port
EXPOSE 8080

# Run Caddy
CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]
