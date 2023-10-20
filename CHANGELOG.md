
# loginsrv changelog

## spezifisch's main branch

* add very basic CORS header support using `-login-allowed-origin` flag
* fix HTTP request timeout not being applied in oauth2/getAccessToken
* update builds to use go1.21
* various minor lint and test cleanups

## v1.3.0

Google will stop support for the Google+ APIs. So we changed loginsrv to use the standard oauth endpoints for Google login.
Please update loginsrv to v1.3.0 if you are using google login.

Since v1.3.0, loginsrv sets the secure flag for the login cookie. So, if you use HTTP fo connect with the browser, e.g. for testing, you browser will ignore the cookie.

* __*ATTENTION:*__ Added a config option to set the secure flag for cookies (default: -cookie-secure=true). If you run unsecure HTTP you have to set this option ot false!!!
* __Google OAuth provider now uses the google userinfo endpoint. No need to activate the google+ APIs anymore.__
* __Added Gitlab OAuth Provider__
* The GET endpoint now returns the user info if the call accepts JSON
* Default OAuth scopes for google and facebook provider. No need to configure them anymore.
* Caddy-plugin: let upstream middleware (e.g. fastcgi and cgi) know about authenticated user
* Caddy-plugin: fixed corner cases in handling of JWT_SECRET paramter for caddy
* Add viewport meta tag to get proper scaling on mobile

## v1.2.4

* Facebook OAuth provider
* Support for custom claims in a user file
* Support for elliptic curve signing methods
* Dynamic redirects
* Some minor cleanups

## v1.2.3

* Bitbucket OAuth provider
* Fix for default secret in caddy plugin
* Replacement of {user} placeholder in caddy logs
* Add domain to google oauth jwt

## v1.1.0

* Added Google OAuth Support
* New Provider: Httpupstream
* Support for multiple comma separated htpasswd files

## v1.1.0

* Official Caddyserver release
* Implement simple jwt refreshing
* Added refresh for htaccess file
* Added github login
* Implemented graceful shutdown
