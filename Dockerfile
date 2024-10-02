FROM node:22.8.0-alpine3.19 AS build-frontend
WORKDIR /app
COPY ./frontend .
RUN npm install
ENV NODE_ENV=production
RUN npm run build

FROM golang:1.23.1-alpine3.19 AS build-backend
WORKDIR /app
COPY . .
COPY --from=build-frontend /app/dist /app/frontend/dist
RUN go build -o main *.go

FROM alpine:3.19 AS server
WORKDIR /app
COPY --from=build-backend /app/ .
RUN chmod +x main
EXPOSE 8080
CMD ["./main"]
