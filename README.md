## gauth

[![CircleCI](https://circleci.com/gh/hashamali/gauth/tree/master.svg?style=svg)](https://circleci.com/gh/hashamali/gauth/tree/master)

Handles authentication using basic auth or JWTs.

#### Types

* `JWTAuth`: Does JWT authentication. Contains the secret key used to sign JWTs as well as the expiration time.
* `BasicAuth`: Defines an interface for handling basic auth.
* `StaticBasicAuth`: `BasicAuth` implementation that uses a static username and password.

#### Methods

* `GetJWTAuth`: Generates a new `JWTAuth`.
* `GetStaticBasicAuth`: Generates a new `StaticBasicAuth`.
* `JWTAuth.Create`: Creates a JWT token with the provided metadata.
* `JWTAuth.Extract`: Extracts metadata from the given JWT token.
* `StaticBasicAuth.Validate`: Validates the given username and password with it's static values.

#### Testing

`make test`
