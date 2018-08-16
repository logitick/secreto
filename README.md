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