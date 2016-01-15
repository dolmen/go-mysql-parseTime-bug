


Test case for https://github.com/go-sql-driver/mysql
6fd058ce0d6b7ee43174e80d5a3e7f483c4dfbe5

Shows that parseTime option is not able to roundtrip time.Time values
with sub-second precision.

Server: mysql 5.6


