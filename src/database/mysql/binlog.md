# binlog

```text
sudo chmod 777 mysql

show master logs;
show binlog events in 'binlog.000005'\G


mysqldump -uroot --add-locks=0 --no-create-info --single-transaction  --set-gtid-purged=OFF db1 t --where="a>900" --result-file=/client_tmp/t.sql



select * from db1.t where a>900 into outfile 't.csv';
```

