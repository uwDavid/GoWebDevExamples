# 13 Setting Cookies
Cookies allow the server to preserve client state. 

# Steps to Preserve State
1. Server sends a cookie to the browser, to indicating the browser's session ID
2. Server stores this session ID, along with the client's state in database
3. When client revisites site and sends back the cookie, we can match the sessionID and retrieve the state for the client browser

# Setting a Cookie
We will use these functions to work with cookies. 
```go
// To set a cookie
func SetCookie(w ReponseWriter, cookie *Cookie)

// To retrieve cookie information
func (r *Request) Cookie(name string) (*Cookie, error)
```

Visit the Go docs to learn more about the Cookie Type.
Some useful items to note here is that: 
1. All cookies implements a String() method
2. Its value is stored as string
3. Max Age allow us to determine how long a cookie will last

Lastly, you can set multiple cookies on one page. 

```go
type Cookie {
	func (c *Cookie) String() string
}
// Cookie implements the String() method
// This is so that any function that needs to print the cookie can call it
type Cookie {
	Name string
	value string

	Path string
	Domain string
	Expires time.Time
	RawExpires string

	Max Age int  //last how many secondes, 0 = not specified, -1 = delete now
	Secure bool
	HttpOnly bool
	Raw string
	Unparsed []string
}
```