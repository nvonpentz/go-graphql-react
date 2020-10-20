drop trigger update_updated_at_on_users on users;

drop table users;

drop function update_updated_at();

drop extension "uuid-ossp";
