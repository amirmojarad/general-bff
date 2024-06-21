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


## Features
1. **Service Management**
    - Easily manage multiple services with customizable configurations.
    - Enable or disable services as required
2. **Path Configuration**
    - Define supported paths for each service with HTTP methods.
    - Streamlined routing to different service endpoints.
3. **Middleware Support**
    - Remove custom prefixes to standardize request handling.
## Future Enhancements   
- [ ] Add support for authentication mechanisms (e.g., JWT, OAuth2).
- [ ] Implement rate limiting to prevent abuse and ensure fair usage.
- [ ] Integrate comprehensive logging and monitoring solutions for better observability.
- [ ] Implement circuit breaker patterns to handle service failures gracefully.
- [ ] Enable dynamic reloading of configuration without restarting the service.

## Getting Started
### Prerequisites
* Go (version 1.15 or later)
* Echo

## Installation
1. Clone the Repository:
    ```shell
      git clone https://github.com/your-repo/bff-project.git
      cd bff-project
    ```
2. Install Dependencies
    ```shell
    go mod tidy
    ```
3. Setup env
    ```
    # App
    APP_NAME=general-bff
    APP_ENV=[devlopment | production]
    APP_PORT=[YOUR_APP_PORT]

    SERVICES_YAML_FILE_PATH=path/to/services.yml
    ```
4. Run the Application
    ```shell
    go run main.go
    ```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact
For any inquiries, please reach out to amirmojarad13@gmail.com.






