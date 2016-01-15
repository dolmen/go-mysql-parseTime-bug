


Test case for https://github.com/go-sql-driver/mysql
6fd058ce0d6b7ee43174e80d5a3e7f483c4dfbe5

Shows that parseTime option is not able to roundtrip time.Time values
with sub-second precision.

Server: mysql 5.6


Example output (expecting time.Time for both column):
```
2016/01/15 15:57:38 SELECT NOW(), ? FROM DUAL
column 0: time.Time 2016-01-15 15:57:39 +0000 UTC
column 1: []uint8 [50 48 49 54 45 48 49 45 49 53 32 49 52 58 53 55 58 51 56 46 56 49 55 51 57] 2016-01-15 14:57:38.81739

```
