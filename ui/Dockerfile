FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json .
RUN npm install

COPY . .
COPY ./.env.placeholder ./.env

RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

COPY 99-replace-placeholders.sh /docker-entrypoint.d/
COPY nginx.conf /etc/nginx/nginx.conf
RUN chmod +x /docker-entrypoint.d/99-replace-placeholders.sh

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
