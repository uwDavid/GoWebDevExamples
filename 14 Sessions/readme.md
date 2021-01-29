# 14 Creating Sessions
Sessions allow us to store user's state, remembering the interactions they had with our website. 

# How do we do this? 
1. We first set-up an in-memory storage for unser info
2. When client browser visits a page for the 1st time, we will generate an unique id (called UUID) to the session
   and pass this UUID as a cookie to client's browser 
3. On server-side, store this UUID as session ID in a database along with other useful information
4. When the client revisits the website, we request for the cookie and find the corresponding info from our database. 

# UUIDs
We will use Google's UUID package. 
```
$ go get github.com/google/uuid
```
For additional information on this package, visit: 
https://pkg.go.dev/github.com/google/uuid

