create database trans_level;

create table record
(
    id   int         not null primary key auto_increment,
    name varchar(20) not null
);
use trans_level;
select * from record;


insert into  record values (1,'hello mysql');
insert into record values (2,'hello my gopher');
SELECT @@global.transaction_ISOLATION;

show variables  like  'trans%';
# READ-COMMITTED

begin ;
insert into record values (2,'hello my gopher');
commit ;


set global transaction_isolation='repeatable-read';
update record set name='commit' where id=1;