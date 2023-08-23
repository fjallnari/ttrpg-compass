FROM node:20

WORKDIR /app

RUN npm install -g pnpm

COPY package*.json ./pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

COPY . .

CMD pnpm run build && node build