# 16 Authentication - Login and Logout
We will continue with our signup example. 

# Login 
We will first initialize our database with a dummy user. 
Then use bcrypt's compare function to check user's login password against the hash in our database. 
```go
// To compare user entered password against our hashed password
func CompareHashAndPassword(hashedPassword, password []byte) error
```

After the user is validated, we will create a session for the user using cookie. 
We can then check against our session database, if there's a session created, then the user is logged in. 

# Permissions
A way to restrict content to users is to set up different roles. 
On request, we can identify the user's role from checking username against our database. 
If the user does not have the required role for a certain page, then we can deny access. 

# Logout
To log a user out, we need to ensure 2 things: 
1. Delete the session from our session database. 
2. Destory the cookie. 

To handle the situation where the user just walked away without logging out, we need to do a clean up of the session database and set a timer on the cookie.

** Need to fix secret page and permissions page.
superfluous response.WriteHeader call from main.secret (inout.go:61) and line 55... 