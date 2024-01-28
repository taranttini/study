# Setup go

## to run

`go mod tidy`

# Setup initial to use mysql

## to run

`docker compose up -d`

## to access bash on mysql

`docker exec -t -i mysql bash`

## to connect on mysql

`mysql -uroot -proot`

## change database

`use goexpert;`

## to create initial base

```sql
create table products (
    id varchar(255), 
    name varchar(80),
    price decimal(10,2),
    primary key (id)
);
```

## to show tables

`show tables`

## to show table detail

`describe TABLE_NAME`

