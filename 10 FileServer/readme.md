# 10 FileServer
The previous example will only serve 1 single file. 
FileServer() will allow us to serve multiple files, but we will need to tweak it with StripPrefix() for it to work properly. 

First a basic example: 
```go
func main(){
	http.Handle("/", http.FileServer(http.Dir(".")))
			//Anything at root, the handler will be FileServer
			//FileServer takes a directory, "." means current directory
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8008", nil)
}

func dog(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`) //this will look for a toby route
	// it will be found in the "/" => and will serve toby.jpg
}
```

But we only want to serve the assets and not the source code. 
So we will have to modify it with StripPrefix(): 
```go
func main(){
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(
		http.Dir("./assets"))))
// http.Handle takes a route, and a handler
// "resources/" => anything down this path is caught using trailing /
// stripPrefix strips off the "/resources" and we are left with /toby.jpg
// stripPrefix takes a prefix to strip off, and a handler
// here, fileServer will serve everything in ./assets => looks for toby.jpg
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg>`)
}
```