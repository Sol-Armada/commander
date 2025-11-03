FROM node:latest AS frontend
WORKDIR /app
COPY package.json package-lock.json ./
RUN yarn install
COPY . .
RUN yarn build

FROM golang:latest AS backend
WORKDIR /app

COPY ./server/go.mod ./server/go.sum ./
RUN go mod download

COPY --from=frontend /app/server/dist ./dist
RUN ls -lah ./dist
COPY ./server .
RUN go build -o ./bin/commander .


FROM ubuntu:latest
WORKDIR /app
COPY --from=backend /app/bin/commander .
# CMD [ "/app/commander" ]
