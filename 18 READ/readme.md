# 18 CRUD Operations - READ
Now that we have our drivers and can connect to the database, let's see how we can made updates to the database. 

CRUD stands for: 
C - Create
R - Read
U - Update
D - Delete

# Set User tester 
In order for our queries to run, we need to give our tester user superuser permissions in psql: 
```sql
ALTER USER tester WITH superuser;
```

# Go File Setup
Step 1: Set up Person type
To process queries from our psql data base, it is a good idea to first set up a struct to ingest all the incoming data. 
In this example, we have set up our struct so that it matches exactly with the columns we expect to get from our query. 

IMPORTANT NOTE: 
The email field/column in the database is allowed to be empty. 
To accomodate for this the email field in our Person struct uses the type NullString, instead of a string. 
Try running this code with email being just a String, and you can see the pointer error that we run into when dealing with potentiall empty fields. 

Step 2: connect to database with sql.Open()
Step 3: db.Ping() to check if connection is live

# Read using Query()
To read from the database we will use the Query() method in the "database/sql" package. 

Step 1: Query() will return a pointer to a type Rows (plural)
Step 2: Use Scan() on the Rows pointer to get 1 row of data
        Note: the order here matches the column order returned by the query

This for Rows.Next() pattern is very well documented here:
https://pkg.go.dev/database/sql#example-DB.Query-MultipleResultSets

# Read using QueryRow()
Similar to Query, but this will only return 1 row of data. 

Step 1: QueryRow() will return a poitner to a Row type (singular)
Step 2: row.Scan() on the row pointer to get our data

Note: we pass in the id parameter to postgres query by using $1. 

