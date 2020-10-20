create extension "uuid-ossp";

create function update_updated_at()
  returns trigger
as
$body$
  begin
    new.updated_at = current_timestamp;
    return new;
  end;
$body$
language plpgsql;

create table users(
  id            uuid primary key not null default uuid_generate_v4(),
  created_at    timestamp with time zone not null default current_timestamp,
  updated_at    timestamp with time zone not null default current_timestamp,
  name text     not null,
  email text    not null,
  password text not null
);

create trigger update_updated_at_on_users
  before update on users
  for each row
  execute procedure update_updated_at();
