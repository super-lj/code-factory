# CodeFactoryCI

This repo contains the code for CodeFactoryCI.

## File Structure
* `ci`
Contains the code for core CI backend services implemented in Go.
* `idl`
Contains the IDL code for Thrift interfaces between CI backend and web application.
* `web`
  * `backend`
  Contains the Go code of web backend for frontend (BFF) that communicates with CI backend using Thrift RPC, transforms the data and serves through GraphQL to frontend.
  * `frontend`
  Contains the frontend code of web application, mainly using React, Material-UI and Apollo (GraphQL client).