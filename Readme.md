# Fullstack Golang NextJs Boilerplate

This is a boilerplate for a fullstack application using Golang for the backend and NextJs for the frontend.

### About the Backend
The backend uses Golang Chi Router and takes an OpenAPI Spec First approach to development. 

#### Features
- Golang
- Zap Logging
- Chi Router
- OpenAPI Spec
- Swagger UI
- Redocly UI
- Spectral Spec Linting
- GolangCI-Lint
- Testify Test

#### What is OpenAPI Spec First Approach?
It's where you design the specifications for the APIs before development and center your development process around the OpenApi Spec. 

Meaning you start by writing the OpenAPI Spec and then generate the server stubs from the spec. This way you're not writing out all the controllers and models for the requests and responses manually.

#### Benefits of OpenAPI Spec First Approach
This approach is great for teams where the frontend and backend are developed in parallel. The frontend team can start developing the UI based on the spec while the backend team can start developing the backend based on the spec. 
It also enforces consistency across functions because you use the same spec for autogenerating code for testing, API gateways, postman collections, etc. 

Checkout the resources below on this methodology:

https://blog.apideck.com/spec-driven-development-part-1

https://tanzu.vmware.com/developer/guides/api-first-development/

### About the Frontend

#### Features
- Typescript
- NextJs
- ESLint
- TailwindCSS




