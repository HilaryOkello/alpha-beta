# Alpha-Beta (Vaccine Distribution and Administration Tracking Platform)

## Overview

This project is a platform designed to track the distribution and administration of vaccines from the manufacturer to the patient, leveraging blockchain for transparency and Go APIs for data integration and analytics. The platform aims to address challenges in Kenya's vaccine distribution system, ensuring secure, transparent, and efficient vaccine tracking.

## Features

- **Blockchain Transparency**: Immutable records of vaccine production, distribution, and administration.
- **Data Integration**: Integration with existing systems using Go APIs.
- **Analytics**: Real-time data analytics to monitor vaccine distribution and identify bottlenecks.
- **Security**: Data encryption and access control to protect sensitive information.
- **Compliance**: Automated compliance reporting using smart contracts.
- **User Interfaces**: Provides interfaces for different stakeholders including manufacturers, distributors and health facilities.

## Architecture

The platform is built using the following components:

- **Blockchain Network**: For transparent and secure data recording.
- **Go APIs**: For data integration and analytics.
- **Frontend**: A web interface for stakeholders to interact with the platform.
- **Backend**: A server-side application to handle business logic and API requests.
- **Database**: For storing non-blockchain data related to vaccines, users, and transaction logs.

## Installation

### Prerequisites

- Go (1.16+)
- HTML/CSS
- Node.js (14+)
- PostgreSQL

### Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/HilaryOkello/alpha-beta.git
   cd alpha-beta
2. **Run the command below to retrieve information on the blockchain**
    ```go
    go run main.go 
    ```
## Usage
**Manufacturers**

Log in to the platform to record vaccine production details.
Use the provided APIs to integrate with existing production systems.

**Distributors**

Log in to the platform to track vaccine shipments.
Use the provided APIs to update shipment status and conditions.

**Health Facilities:**

Log in to the platform to confirm receipt of vaccines.
Record vaccine administration details.

**Regulators:**

Access real-time data and analytics to monitor vaccine distribution.
 Use the audit trail to verify compliance and investigate issues.

## API Documentation

The API documentation is available at main.go . It includes:
- Endpoint Descriptions: Information on each endpoint, including purpose and usage.
- Request and Response Formats: Examples of how to structure requests and responses.

## Contributing

We welcome contributions from the community! Please read our Contributing Guidelines for more information on how to get involved.

## License

This project is licensed under the [MIT](https://opensource.org/license/mit) License. See the LICENSE file for details.

## Contributors

- [Hilary](https://github.com/HilaryOkello)
- [Joan](https://github.com/Joan2509)
- [Joab](https://github.com/JoabOwala)
- [Cynthia](https://github.com/CynthiaOketch)
- [Antony](https://github.com/antmusumba) 

