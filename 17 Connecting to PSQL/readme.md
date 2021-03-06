# 17 Connecting to PSQL
Now we will move onto working with SQL databases. For this example, I will use PostgresSQL. 

Here's a very helpful guide to working with SQL in Windows:
https://www.microfocus.com/documentation/idol/IDOL_12_0/MediaServer/Guides/html/English/Content/Getting_Started/Configure/_TRN_Set_up_PostgreSQL.htm

To start psql: 
```shell
psql -U postgres
```
Create a database: 
```sql
CREATE DATABASE test;
```
Create a user: 
```sql
CREATE USER tester WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE test TO tester;
```
List users: 
```sql
\du
```
To connect to database use /c in sql: 
```sql
\c test
```
# Generate Dummy Data
You can just copy and paste in prewritten commands. 
Or, you can write your SQL commands in a .sql file and run the file:  
```sql
\i path\person.sql
```

# Connecting to DB
We will have to use 2 packages to connect to postgres: 
Run $ go get [package_name]  
```Go
import (
	"database/sql"
	_"github.com/lib/pq"
)
```
Step 1: connect to database with sql.Open()
Step 2: db.Ping() to check if connection is live
