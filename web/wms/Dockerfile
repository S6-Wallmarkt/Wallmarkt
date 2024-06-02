FROM node:20.10-bookworm-slim as build

WORKDIR /app

COPY package.json ./
COPY pnpm-lock.yaml ./

# pnpm must be installed as it doesn't come with the default image
RUN npm i -g pnpm
RUN pnpm i
COPY . ./

# Accept build arguments for environment variables
ARG AUTH0_DOMAIN
ARG AUTH0_CLIENT_ID
ARG AUTH0_AUDIENCE
ARG GATEWAY_BASE


# Set environment variables during the build process
ENV VITE_AUTH0_DOMAIN=$AUTH0_DOMAIN
ENV VITE_AUTH0_CLIENT_ID=$AUTH0_CLIENT_ID
ENV VITE_AUTH0_AUDIENCE=$AUTH0_AUDIENCE
ENV VITE_GATEWAY_BASE=$GATEWAY_BASE

RUN pnpm build

FROM node:20.10-bookworm-slim

ENV VITE_AUTH0_DOMAIN=${AUTH0_DOMAIN}
ENV VITE_AUTH0_CLIENT_ID=${AUTH0_CLIENT_ID}
ENV VITE_AUTH0_AUDIENCE=${AUTH0_AUDIENCE}
ENV VITE_GATEWAY_BASE=${GATEWAY_BASE}

WORKDIR /app
COPY --from=build /app .

EXPOSE 3000

CMD ["node","build"]