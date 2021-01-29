# 15 Authentication - Sign Up
For authentication, we will need a user to sign-up with a username and a password. 
But our server should NEVER store the user's password directly into our database. 

# Password Salting and Hashing
Instead of saving the password directly into our database, we will have to encrypt it. 
But just encrypting it is not enough, because a hacker can still reverse engineer the most common passwords. 
To counter this, we will have to 'salt' the password by adding additional phrases to each password. 
So that similar passwords will have different hash. 

To do this, we will use the standard 'bcrypt' package:
https://pkg.go.dev/golang.org/x/crypto/bcrypt

```
$ go get golang.org/x/crypto/bcrypt
```
The 2 functions we will use: 
```go
// To generate hash from password
// Note: the hashed password is a slice of byte
func GenerateFromPassword(password []byte, cost int) ([]byte, error)

// To compare user entered password against our hashed password
func CompareHashAndPassword(hashedPassword, password []byte) error
```
In case we will need to increase hash cost in the future.
```go
func Cost(hashedPassword []byte) (int, error)
```
