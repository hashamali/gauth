# gjwt
[![godoc](https://godoc.org/github.com/hashamali/gjwt?status.svg)](http://godoc.org/github.com/hashamali/gjwt)
[![tests](https://img.shields.io/github/workflow/status/hashamali/gjwt/tests?label=tests&style=flat-square)](https://github.com/hashamali/gjwt/actions?query=workflow%3Atests)
[![sec](https://img.shields.io/github/workflow/status/hashamali/gjwt/security?label=security&style=flat-square)](https://github.com/hashamali/gjwt/actions?query=workflow%3Asecurity)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/gjwt)](https://goreportcard.com/report/github.com/hashamali/gjwt)
[![license](https://badgen.net/github/license/hashamali/gjwt)](https://opensource.org/licenses/MIT)

Handles authentication using JWTs.

#### Types

* `JWTAuth`: Does JWT authentication. Contains the secret key used to sign JWTs as well as the expiration time.

#### Methods

* `JWTAuth.Create`: Creates a JWT token with the provided metadata.
* `JWTAuth.Extract`: Extracts metadata from the given JWT token.

#### Testing

`make test`
