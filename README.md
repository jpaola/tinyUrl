# Shorten URL

The goal of this project is to use a free API service to shorten the URL of https://marketplace.digitalocean.com/ and output it either on the screen or via the terminal. The service in use is called https://cleanuri.com. This service does not require any authentication of any kind.

Using the API documentation found in https://cleanuri.com/docs the script calls on the POST method to the "https://cleanuri.com/api/v1/shorten" endpoint and the response should contain the shortened URL. The scripts will display the response via the terminal.


## Running the GOLANG solution

To run this script, you must have Golang installed in your system - https://go.dev/doc/install.

To run the golang script run the following command:

```bash
go run tinyUrl.go
```

## Running the JS solution
To run this script, you will need yarn, node, and axios.

Install and use latest node version - https://nodejs.org/en/download/

Install yarn - https://classic.yarnpkg.com/lang/en/docs/cli/install/

Install axios - https://axios-http.com/docs/intro or run

```bash
yarn add axios
```

To run the javascript file run the following command:

```bash
node tinyUrl.js
```
