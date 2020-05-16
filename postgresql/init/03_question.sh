# Questions table
psql -U postgres -d judge << "EOSQL"
create table questions(
  id SERIAL PRIMARY KEY,
  title varchar(128),
  body varchar(512),
  validation varchar(128),
  input varchar(128),
  output varchar(128),
  created_at timestamptz,
  updated_at timestamptz 
  )
EOSQL

psql -U postgres -d judge << "EOSQL"
insert into questions(title, body, validation, input, output, created_at, updated_at) 
values(
  '駄菓子屋さん',
  '1つX円のお菓子をY個買います。Z円出したときのお釣りを出力してください。',
  'Z > X*Y',
  'X Y Z',
  'お釣りを出力してください。',
  current_timestamp,
  current_timestamp
  )
EOSQL

psql -U postgres -d judge << "EOSQL"
insert into questions(title, body, validation, input, output, created_at, updated_at) 
values(
  '算数',
  'X+Y の計算をします。',
  '',
  'X Y',
  '計算の結果を出力してください。',
  current_timestamp,
  current_timestamp
  )
EOSQL

# 
psql -U postgres -d judge << "EOSQL"
create table sample(
  question_id SERIAL PRIMARY KEY,
  body varchar(512),
  validation varchar(128),
  input varchar(128),
  output varchar(128),
  created_at timestamptz,
  updated_at timestamptz 
  )
EOSQL