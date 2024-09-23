# Registro 📚

`Registro` is a backend service that utilizes Kafka for real-time event processing and a blockchain
ledger for secure, immutable audit trails. The project aims to provide a robust and scalable
system for tracking events with integrity and accountability.

## ✨Features

-[*] Real-time event production and consumption using Apache Kafka. -[*] Immutable record of events stored in a blockchain-like data structure. -[*] Endpoints for producing events, retrieving the blockchain state, and validating ledger integrity.

## API Endpoints

- Produce Event
  - Endpoint: `POST/produce`
  - Description: Sends an event to Kafka to be recorded in the ledger.
  - Request Body:
  ```
  {
      "topic": "libro-events",
      "message": "Your event message"
  }
  ```
- Get Blockchain

  - Endpoint: `GET/events`
  - Description: Retrieves the current state of the blockchain ledger.

- Validate Blockchain
  - Endpoint: `GET/blockchain`
  - Description: Validates the integrity of the blockchain ledger.
