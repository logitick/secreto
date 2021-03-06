[![Build Status](https://travis-ci.org/logitick/secreto.svg?branch=master)](https://travis-ci.org/logitick/secreto)
[![Test Coverage](https://api.codeclimate.com/v1/badges/974721adb0268898ebff/test_coverage)](https://codeclimate.com/github/logitick/secreto/test_coverage)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/1ec5b7fde1154cf8be51f634e5782a8d)](https://app.codacy.com/app/paul_38/secreto?utm_source=github.com&utm_medium=referral&utm_content=logitick/secreto&utm_campaign=Badge_Grade_Settings)
# secreto 
a cli tool to encode and decode kubernetes secrets


## How to use

### decoding a secrets.yml to its literal values
<details>
  <summary>./encoded_secrets.yml</summary>

```yml
apiVersion: v1
kind: Secret
metadata:
  name: database-secret-config
type: Opaque
data:
  username: QXp1cmVEaWFtb25k
  password: aHVudGVyMg==
```  
</details>

```bash
$ secreto decode ./encoded_secrets.yml 
apiVersion: v1
kind: Secret
metadata:
  name: database-secret-config
type: Opaque
data:
  username: AzureDiamond
  password: hunter2

# save to a file
$ secreto decode ./encoded_secrets.yml > secrets.yml
```

### encoding a literal secrets.yml to base64
<details>
  <summary>./secrets.yml</summary>

```yml
apiVersion: v1
kind: Secret
metadata:
  name: database-secret-config
type: Opaque
data:
  username: AzureDiamond
  password: hunter2
```  
</details>

```bash
$ secreto encode ./secrets.yml
apiVersion: v1
kind: Secret
metadata:
  name: database-secret-config
type: Opaque
data:
  username: QXp1cmVEaWFtb25k
  password: aHVudGVyMg==
```

## Roadmap
- encryption & decryption of the secret values to make it safe to store in VCS