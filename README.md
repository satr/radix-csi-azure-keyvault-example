# Example of an app with use of CSI Azure KeyVault driver

[Radix supports connection]([https://radix.equinor.com/guides/azure-key-vaults/) to Azure Key vaults. They are available as secrets within Radix application components - content is mapped to replica environment variables and files.

WARNING! The application shows content of following environment variables: 
`CONNECTION-STRING`, `DB_USER`, `DB_PASS`, `DB_QA_USER`
Avoid connecting real secrets to these variables for this example!

To run locally run the command
```
docker-compose up --build
```
and open a URL [http://localhost:8082](http://localhost:8082) in a browser. 
