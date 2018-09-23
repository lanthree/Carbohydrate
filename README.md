# Carbohydrate

+ install mysql & go-lang
+ rebuild databases
	1. `cd ./Data`
	2. `mysql -u$(username) -p$(passwd) < CanEat.sql`
	3. `python readfromxls.py`
+ install go deps
	1. `go get -u github.com/go-sql-driver/mysql`
	2. `go get -u github.com/gin-gonic/gin`
	3. `go get -u github.com/go-ego/riot`
