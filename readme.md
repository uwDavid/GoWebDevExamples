# Go Web App Concepts
This folder shows some fundamental concepts of web server programming in Go. 
It aims to use small code snippets to demonstrate how to use Go's libraries. 

# 01-Using Template package
This will demonstrate how we can pass variables to templates.
And use templates to generate our desired html file. 

Step 1: Parse file - this will take our template file, tpl.gothml
        The extension name ".gohtml" can be anything here.
        There are a few methods available for us: template.ParseFiles(), ParseGlob
        Visit the official documentation on how to set up templates.

Step 2: Create a file using os.Create() method for our "index.html"

Step 3: Execute the file - this will convert the template into html. 

# 02-Starting a Server
