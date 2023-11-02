FROM node:18-alpine as build
WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN npm install -g pnpm
RUN pnpm install
COPY . /app
RUN pnpm run build

FROM node:alpine
COPY --from=build /app/dist ./dist
RUN npm install -g serve
CMD ["serve", "-s", "dist"]