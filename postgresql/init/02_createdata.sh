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

psql -U postgres -d judge << "EOSQL"
create table questions(
  id SERIAL PRIMARY KEY,
  body varchar(512),
  validation varchar(128),
  input varchar(128),
  output varchar(128),
  created_at timestamptz,
  updated_at timestamptz 
  )
EOSQL

psql -U postgres -d judge << "EOSQL"
insert into questions(body, validation, input, output, created_at, updated_at) 
values(
  '1つX円のお菓子をY個買います。Z円出したときのお釣りを出力してください。',
  'Z > X*Y',
  'X Y Z',
  'お釣りを出力してください。',
  current_timestamp,
  current_timestamp
  )
EOSQL

psql -U postgres -d judge << "EOSQL"
insert into questions(body, validation, input, output, created_at, updated_at) 
values(
  'X+Y の計算をします。',
  '',
  'X Y',
  '計算の結果を出力してください。',
  current_timestamp,
  current_timestamp
  )
EOSQL