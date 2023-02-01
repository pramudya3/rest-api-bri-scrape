create table accounts(
	id int not null auto_increment primary key,
	bank varchar(50),
	username varchar(100) unique,
	password varchar(255),
	account_number varchar(255),
	check_interval int,
	time_active int,
	auto_logout bool
)engine=innodb;