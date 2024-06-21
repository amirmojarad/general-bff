# General BFF

Welcome to the General BFF project, a Go-based service designed to streamline your API interactions using the Echo web framework. This project provides a centralized interface for managing multiple microservices, enabling seamless and efficient communication.

## Configuration
The configuration file is written in YAML format. Below is an example configuration that outlines the necessary settings.
```yaml
version: 1.0

middlewares:
  prefixes:
    api: "/api"

services:
  - name: [YOUR_SERVICE_NAME]
    enabled: [true | false]
    service_url: [YOUR_SERVICE_URL]
    paths:
      - GET: "/v1/products"
      - GET: "/v1/product/:id"
      - POST: "/v1/product"
```

### Parameters
- **middlewares**
  - **prefixes**: define url prefixes that you want to use in this service and wants to remove in redirect to source service.
- **services**
  - name: your service name
  - enabled: clear :)
  - service_url: base url of your service
  - paths: your service paths and routes