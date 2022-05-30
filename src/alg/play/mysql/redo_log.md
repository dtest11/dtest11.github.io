```sql

CREATE TABLE `T` (
  `ID` int(11) NOT NULL,
  `c` int(11) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


```

mac mysql data location:

/usr/local/var/mysql

/usr/local/etc/my.cnf

```sql

Show variables like '%slow_query%'; -- can use this query for all variables

//first step
set global log_output='TABLE'; -- Turn on slow logging and log to mysql.slow_log table
set global long_query_time=2; -- Set queries longer than 2 seconds to slow queries
set global slow_query_log='ON';-- Turn on slow logging

/ / The second step Run a slower function, execute the following statement
select convert(sql_text using utf8) sql_text from mysql.slow_log -- Query slow sql logs
//The third step, remember to close the log.
set global slow_query_log='OFF'; -- If you don't have to remember to close the log

```


mysqldump -uroot -p -v binlog.000003

https://dev.mysql.com/doc/employee/en/employees-installation.html
