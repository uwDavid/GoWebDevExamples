# 04 Retrieving Values from Request
In this example, we will look at how we can retrieve values from Request. 

There are 2 ways the Request can send values. 
1) If data is sent using POST request via a form => the data will be in the body of the Request
2) Data can also be embedded in the URL parameters of a GET request

To extract the values, we will utilize request.ParseForm() method. 

Example:
1. We use template.ExecuteTemplate() to display our html page
   To do so, we must first set up our template in "index.gohtml". 
   And then call template.ParseFiles() to parse "index.gohtml"

2. To pass the data to "index.gohtml" in our hotdog Handler. 
   We first have to call request.ParseForm().
   The variable names and values for ParseForm is set up in the form submission table. 
   After ParseForm(), we can then pass request.Form as data for ExecuteTemplate()
   Data is of the format map[string][]string - key returns a slice of string

Additional Note: 
The variable 'fname' has 2 values: 
1) There's a value in the URL <form action="/?fname=James" method="POST">  
2) The form also passed a value
