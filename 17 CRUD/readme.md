# 17 CRUD Operation
Now we will move onto working with SQL databases. For this example, I will use PostgresSQL. 

Here's a very helpful guide to working with SQL in Windows:
https://www.microfocus.com/documentation/idol/IDOL_12_0/MediaServer/Guides/html/English/Content/Getting_Started/Configure/_TRN_Set_up_PostgreSQL.htm

To start psql: 
```shell
psql -U postgres
```
Create a database. 
Once database is created use /c to connect to the database. 
Then to run the .sql file to generate dummy data:
```shell
\i path/filename.sql
```

# Connecting to DB
We will have to use 2 packages to connect to postgres: 
```Go
import (
	"database/sql"
	_"github.com/lib/pq"
)
```
Step 1: connect to database with sql.Open()
Step 2: db.Ping() to check if connection is live
