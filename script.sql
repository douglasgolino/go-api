create table product(
	id Serial primary key,
	product_name varchar(50) not null,
	price numeric (10,2) not null
);

insert into product (product_name, price) values('Sushi', 100);
insert into product (product_name, price) values('Temaki', 20);

