# 03 net/http Package
The most critical item to understand about the http package is the Handler interface.
```
type Handler interface{
    ServerHTTP(ResponseWriter, *Request)
}
```
There are 2 available functions for us to listen for requests. 
Both of these functions will require a Handler type that implements the ServerHTTP() method. 
```
func ListenAndServe(addr string, ahandler Handler) error

func ListenAndServe(addr, certFile, keyFile string, ahandler Handler) error
```
The ResponseWriter and Request type/interface detail is available on official Go documentations. 

Here's a example of how polymorphism is in effect. 
We will define a variable of int type, and implement the Handler interface in it. 
We will then pass this variable as a Handler in http.ListenAndServe() function. 

To turn on the server 'go run' the file. 
To exit the server press Ctrl+C. 