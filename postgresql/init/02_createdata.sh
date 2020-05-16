#!/bin/bash

psql -U postgres -d judge << "EOSQL"
ALTER DATABASE judge SET timezone TO 'Asia/Tokyo';
EOSQL

psql -U postgres -d judge << "EOSQL"
create table users(
  id SERIAL PRIMARY KEY,
  name varchar(128),
  email varchar(128),
  created_at timestamptz,
  updated_at timestamptz 
  )
EOSQL

psql -U postgres -d judge << "EOSQL"
  insert into users (name, email, created_at, updated_at )values ('Rito', 'rito@example.com', '2019-01-21 12:00:00', '2019-01-21 12:00:00')
EOSQL

psql -U postgres -d judge << "EOSQL"
  insert into users (name, email, created_at, updated_at )values ('Minami', 'south@example.com', '2019-01-21 12:00:00', '2019-01-21 12:00:00')
EOSQL
