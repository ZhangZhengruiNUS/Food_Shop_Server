# Development instructions

>**Recommended development OS version: Ubuntu 22.04.3 LTS**  
**For Windows, using WSL is highly recommended.**  
**For other OS, such as MacOS, some details may be different, please use Google/Chatgpt to solve it**  

![Backend Development Pattern](img/design_pic2.jpg)

_Continuously updated..._

## 1. Preparation

### 1.1 IDE

VSCode (recommended), other IDE is also acceptable

### 1.2 Clone Code

```bash
git clone git@github.com:ZhangZhengruiNUS/Food_Shop_Server.git
```

If using the HTTP protocol, please modify the link

### 1.3 Tools Installation

- Install gvm
_Go Version Manager_

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

you can see all the go versions you can install by:

```bash
gvm listall
```

- Install Go

```bash
gvm install go1.21.3
```

Set the version of 1.21.3 as default

```bash
gvm use go1.21.3 --default
```

- Install gopls
_Provide "Code Autocompletion", "Code Navigation", etc._

```bash
go get -u golang.org/x/tools/gopls
```

- Install sqlc
_Generate Go Code via sqlc script_

```bash
sudo snap install sqlc --channel=stable/1.22.0
```

- Install viper
_Reading configurations_

```bash
go get github.com/spf13/viper
```

- Install mock
_Provide mock implementations of interfaces for unit testing._

```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

### 1.4 Database

Use Docker image to get a Postgresql DB (recommended) or download a Postgresql DB to your local environment

#### 1.4.1 Install Docker

For WSL & MacOS, download Docker Desktop directly from the official website

#### 1.4.2 Docker Image

Use the following commands in the terminal

```bash
docker pull postgres:16-alpine
```

#### 1.4.3 Run PostgreSQL DB

- Install make

Use the following commands in the terminal

```bash
sudo apt update
sudo apt install make
```

- Use "make" to run docker

In the terminal, navigate to the project root directory and use the following commands

```bash
make postgres
```

This will trigger the instruction corresponding to "postgres" in "Makefile"

- Create database

In the terminal, navigate to the project root directory and use the following commands

```bash
make createdb
```

This will trigger the instruction corresponding to "createdb" in the Makefile

- Create init tables of database

In the terminal, navigate to the project root directory and use the following commands

```bash
make migrateup
```

This will trigger the instruction corresponding to "createdb" in the Makefile
It will execute the SQL scripts in the latest **0XXXXX_XXXX.up.sql** file under the Food_Shop_Server/db/migration

## 2. Common rules

- Use CamelCase to name variable, like this:

>userName, inventoryId

- Write more comments to show indicate the work of the regions, like this:

![comments example](img/1697780205697.jpg)

Especially, add structured comments before **function** and **struct definition**, like this:

![comments example](img/1697780076090.jpg)

- Write "fmt" log at the beginning and end of the handler function, like this:

![comments example](img/1697780038549.jpg)

## 3. Use Git to start team development

>_You can also use the git shortcut button in the IDE, as long as you ensure that your actions are correct_

### 3.1 Use branch to handle major changes

- When you want to start your current work for the first time, please create a new branch, and it is recommended to summarize and name it with the approximate work you will complete this time.

```bash
git checkout -b my-work
```

- If you are unsure of the current branch, please use the following command to confirm

```bash
git branch
```

(All branches will be listed and marked with * next to the current branch)

- If the current branch is not the feature/my feature branch you created you are preparing to write code, please use the following command to switch to that branch

```bash
git checkout my-work
```

- When completing this stage of code writing

Use the following command to get latest code on main branch **(Important)**
_Please do this frequently to ensure your code has the latest version of the main branch_

```bash
git fetch origin
git merge origin/main
```

If there's no more modification, use the following command to add your changes to the local staging area

```bash
git add .
```

_This will add all modified files to the staging area. If you only want to add specific files, you can Replace "." with the file name._

- Submit your changes using the following command

```bash
git commit -m "A brief description of the work"
```

- Push to remote repository

```bash
git push origin my-work
```

- Pull request to merge branch

>a. Enter the GitHub repository page, open the Pull Request page, Click the "New Pull Request" button  
b. Select the branch to merge  
c. Compare changes  
d. Fill in the modified information  
e. Click the "Create Pull Request" button  
f. Approve and merge the pull request on your own

### 3.2 Directly handle small changes in the main branch

>_Be cautious when making modifications on the main branch. Ensure that no other person is making changes that conflict with your changes at the same time._

- First, make sure you are on the main branch

```bash
git checkout main
```

- Get the latest code **(Important)**

```bash
git pull origin main
```

- When completing modification, use the following command to add your changes to the local staging area

```bash
git add .
```

- Submit your changes using the following command

```bash
git commit -m "A brief description of the work"
```

- Push to remote repository

```bash
git push origin main
```

## 4. Start Develop

### 4.1 Run Project

- When you first clone the code or pull the new version of the main branch code, you need to check and download the external libraries required for the project

```bash
go mod download
```

- When you finish some code, you want to test it by Postman or other tools, navigate to the project root directory and use the following commands to run the server

```bash
make server
```

This will trigger the instruction corresponding to "createdb" in "Makefile":
>go run main.go

- Verify

Download the Postman tool for service verification

![verify example](img/5249398125634204d8c40607c285935.png)

- Config file

In "app.env"

![config file](img/1697875304184.png)

### 4.2 Code

#### 4.2.1 Router

When you want to add a new route, modify "NewServer" function in "handlers/server.go"

![server.go example](img/1697479547897.png)

#### 4.2.2 Service

When you want to add a service in a new field, add a file under "handlers"

![add new service](img/1697479982877.png)

When you need to print and check the data from the frontend, you can print it like this:

```go
bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
fmt.Println(string(bodyBytes))
ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
```

#### 4.2.3 Sql

##### 4.2.3.1 Non-Transactions

>Only in the following two cases, you can use this method:  
a. Change(Create/Update/Delete) a single record in a single table, without the need to verify the data that could be modified by other transactions  
b. Only read tables

- Write the sqlc scripts in the "db/query/table_name.sql"(follow the syntax of sqlc), like this:
_(ref: <https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html>)_

![sqlc scripts example](img/1697650119551.jpg)

>-- name: GoFunctionName -> The Go function name you want generate automatically  
:one -> Return only one record  
:many -> Return some records  
:exec -> Return zero record

When you need to query two or more tables, write like this:

```sql
-- name: ListInventoriesByUserID :many
SELECT t1.item_id, COALESCE(t2.item_name, ' '), COALESCE(t2.describe, ' '), t1.quantity, COALESCE(t2.pic_path, ' ')  FROM inventories t1
LEFT JOIN items t2 ON t1.item_id = t2.item_id
WHERE user_id = $3
ORDER BY t1.item_id
LIMIT $1
OFFSET $2;
```

- in the terminal, navigate to the project root directory and use the following commands

```bash
make sqlc
```

This will generate the affected files under "./db/sqlc", then you can find the sqlc sctipts corresponding Go functions there and use them in the handler services

##### 4.2.3.2 Transactions

>In the following two cases, but not only the two cases, you should use this method:  
a. You want to change(Create/Update/Delete) more than one table  
b. You want to change(Create/Update/Delete) one table with more than one record  
c. You want to change(Create/Update/Delete) a single record in a single table, with the need to verify the data that could be modified by other transactions

- The same way as "Non-Transactions" to write Sqlc sctipts and generate Go functions for a single table.
_**Note: If you want to ensure that the data you need to verify remains unchanged before making the change(Create/Update/Delete), you should use an "Explicit Lock" by adding "... FOR UPDATE" to your SQL query**_

- Use "ExecTx" function in the "db/sqlc/store.go", put the steps in its call back function, like this:

![sqlc scripts example](img/1697722482517.jpg)

### 3.4 Database Modification

>In this way, every change to the database structure is recorded in detail and can be easily applied or rolled back.

When you want to change the structure of the tables

- In the terminal, navigate to the project root directory and use the following commands

```bash
make migratecreate name=migration_name
```

_describe the changes you want to make in the tables in the "name" parameter_
This will trigger the instruction corresponding to "migratecreate" in the Makefile:
>migrate create -ext sql -dir db/migrations -seq $(name)

This command will generate **0XXXXX_migration_name.up.sql** & **0XXXXX_migration_name.down.sql** file under the Red_Pocket_Game_Server/db/migration (the serail number will automatically increase)

- Write the difference script in the 0XXXXX_migration_name.up.sql file and write rollback scripts in the 0XXXXX_migration_name.up.sql
_if the changes are quite significant, you can try to use some database compare tools to automatically generate the scripts_

>**Important: Do consider the impact of these changes on existing data**

- In the terminal, navigate to the project root directory and use the following commands to apply the changes

```bash
make migrateup
```

If you want to rollback the changes, use the following commands

```bash
make migratedown
```

After you run the command, if there's any mistake in the scripts, the migration will stop and when you run the migraion command again, it will probably show the error message:
>error: Dirty database version 1. Fix and force version.

This is because before the migration runs, each database will be set with a dirty flag, if the migration fails, the dirty state will persist. This will prevent future migration operations.
You can update the change the version of the database into the lastest version forcefully to solve this, use the following commands

```bash
make migrateforce version=latest version number
```

_put the latest version number in the "version" parameter_
This will trigger the instruction corresponding to "migrateforce" in the Makefile:
>migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable" force $(version)

Then you remove the dirty flag, you can continue the migration operations.

- check if your "db/query/table.sql" need to change

If you make some changes in the files
In the terminal, navigate to the project root directory and use the following commands to apply the changes

```bash
make sqlc
```

This will regenerate the affected files under "./db/sqlc"

![account.sql.go](img/1697375254302.png)

### 3.5 Test

## 5. Dockerfile

## 6. CICD
