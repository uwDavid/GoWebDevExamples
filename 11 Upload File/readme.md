# 11 Uploading File to Server
Here we will create a way for client to upload file to server. 

We just have to use request.FormFile() instead of FormValue() to "catch" the file. 
Step 1: Implement Handler, such that if there's a POST request, catch it use request.FormFile()
Step 2: Use ioutil.ReadAll() to read the file.
        Note the ReadAll(), returns a byte slice, but it will have a toString() method. 
Step 3: We can choose to save the file on the server using os.Create()
        Then we can pass the byte slice, into the file Writer.
        We can also destinate save folder in this step by embedding filepath.Join() inside os.Create()

On the HTML side, we just need to have a place for users to upload the file. 
This is done via a submission form. 
There are 3 enctype available: 
1. enctype = "application/x-www-form-urlencoded"
   This is the default enctype. It provides a key-value pair for form data.

2. enctype = "multipart/form-data"
   This is the enctype we will use to handle file uploads. 

3. enctype = "text/plain" 
   This is pretty much used for debugging. 