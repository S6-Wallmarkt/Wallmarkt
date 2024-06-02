FROM node:20.10-bookworm-slim as build

WORKDIR /app

COPY package.json ./
COPY pnpm-lock.yaml ./

# pnpm must be installed as it doesn't come with the default image
RUN npm i -g pnpm
RUN pnpm i
COPY . ./

RUN pnpm build

FROM node:20.10-bookworm-slim

WORKDIR /app
COPY --from=build /app .

EXPOSE 3000

CMD ["node","build"]