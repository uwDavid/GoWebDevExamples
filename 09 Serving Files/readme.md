# 09 Serving Files
There are many ways in which we can serve files to our http Request.

# Method 1: Using io.copy()
```go
func main(){
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dogPic(w http.ResponseWriter, req *http.Request){
	f, err = os.Open("toby.jpg")  //returns pointer to a file
	if err!=nil{
		http.Error(w, "file not found", 404)
		return 
	}
	defer f.Close()
	io.Copy(w, f) //serving file with io.copy
}
```

# Method 2: Using ServeContent or ServeFile
We perfer using ServeFile() because it's alot easier than ServeContent()
```go
func main(){
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil) 
}

func dog(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg">`) 
}

func dogPic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "toby.jpg")  // alot simpler than ServeContent
}
//Note: toby.jpq is in the same folder
```