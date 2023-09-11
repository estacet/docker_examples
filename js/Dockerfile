FROM node:20 as build

WORKDIR /app

COPY . .

RUN npm install && npm run build

FROM node:20-alpine

WORKDIR /app

COPY --from=build /app/index.js /app/index.js

CMD ["node", "/app/index.js"]

