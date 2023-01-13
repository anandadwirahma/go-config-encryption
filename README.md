# go-config-encryption

### Prerequisites
1. Create a secret key and put it in the SECRETKEY variable in the .env file.
```sh
SECRETKEY="253AC3AB47D6422622A4864B3EB53"
```
2. Hash the secret key, then take the last 32 digits. Use it as your secret key to encrypt your config later on.
![img.png](https://ibb.co/TYXtx8H)

   <br />The last 32 digits are **67d81e2c5717548a4ee1bd1e81395746**<br /> 
   Use it to encrypt your signature.<br />

   _You can use the following tools to generate it easily. [SHA-256 Generator](https://emn178.github.io/online-tools/sha256.html)_<br />


### Usage

1. Change the data type of the attribute which wants to use encryption format to be EncryptedValue type.
```sh
    type Config struct {
      Application Application `json:"application"`
      Database    Database    `json:"database"`
    }

    type Application struct {
      ClientID  EncryptedValue `json:"client_id"`
      SecretKey EncryptedValue `json:"secret_key"`
    }

    type Database struct {
      Name     EncryptedValue  `json:"name"`
      Username EncryptedValue  `json:"username"`
      Password EncryptedValue  `json:"password"`
      Host     string          `json:"host"`
      Port     int             `json:"port"`
    }
```
2. Encrypt your config value using the secret key in **_Prerequisites point 2_**, then put it into .env file.<br />
   _You can use this tool to encrypt your value: [AES-256-CBC Encryptor](https://encode-decode.com/aes-256-cbc-encrypt-online/)<br />
![img_1.png](img_1.png)_
   <br />Encrypted config example:
```sh
SECRETKEY="253AC3AB47D6422622A4864B3EB53"

APPLICATION_CLIENTID="ENC(971p+xPuEEuXlE3UIKd1y8cwZ7YmR8Iw5zEFLmaRRGs=)"
APPLICATION_SECRETKEY="ENC(pN8SXPZKHpBOdCjhozHJYHyueKPgfCSFHEaTXr5wri0=)"

DATABASE_NAME="ENC(BCFmDsKvzOwwsWEMH4BxIg==)"
DATABASE_USERNAME="ENC(xd4tUTVt6aTTnaGsQVpcEw==)"
DATABASE_PASSWORD="ENC(iuGDWYv/mu+ZtBGH3Oz4aQ==)"
DATABASE_HOST="localhost"
DATABASE_PORT=8080
```

If you want to use plain text configuration, just put the value directly(without **ENC() text pattern**). <br />

```shell
SECRETKEY="253AC3AB47D6422622A4864B3EB53"

APPLICATION_CLIENTID="this is client ID"
APPLICATION_SECRETKEY="this is secret key"

DATABASE_NAME="db_name"
DATABASE_USERNAME="db_username"
DATABASE_PASSWORD="db_password"
DATABASE_HOST="localhost"
DATABASE_PORT=8080
```

