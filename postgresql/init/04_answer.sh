# Answer table
psql -U postgres -d judge << "EOSQL"
create table answers(
  id SERIAL PRIMARY KEY,
  question_id integer,
  answer varchar(1024),
  status varchar(128),
  result varchar(1024),
  detail varchar(1024),
  created_at timestamptz,
  updated_at timestamptz,
  FOREIGN KEY (question_id) REFERENCES questions(id)
  )
EOSQL

psql -U postgres -d judge << "EOSQL"
insert into answers(question_id, answer, status, result, detail, created_at, updated_at) 
values(
  1,
  'print "YES"',
  'status',
  'result',
  'detail',
  current_timestamp,
  current_timestamp
)
EOSQL
