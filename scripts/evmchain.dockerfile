
FROM node:18-alpine AS builder

WORKDIR /app
COPY ./contracts/package*.json ./

RUN npm install

# Copy project files
COPY ./contracts/ .

RUN npx hardhat compile

# Final image for running the EVM chain
FROM node:18-alpine

WORKDIR /app

COPY --from=builder /app /app

EXPOSE 8545

CMD ["npx", "hardhat", "node"]
