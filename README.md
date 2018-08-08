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
  username: bGV0bWVpbgo=
  password: cGFzc3dvcmQxMjMK
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
  username: letmein
  password: password123

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
  username: letmein
  password: password123
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
  username: bGV0bWVpbgo=
  password: cGFzc3dvcmQxMjMK
```

## Roadmap
- encryption & decryption of the secret values to make it safe to store in VCS