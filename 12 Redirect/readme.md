# 12 Redirect
In the official RFC 2616 lists 300's codes as redirection codes:

301 moved permanently - the browser remembers this redirect
303 see other - this redirect will always change the request method to GET
307 temporary redirect - preserves the request method 

There are 2 ways to handle redirects: 
1. use http.ResponseWriter.Header().Set()
```go
w.Header().Set("Location", "/")
w.WriteHeader(http.StatusSeeOther) 
```

2. use http.Redirect()


We will just focus on using http.Redirect()