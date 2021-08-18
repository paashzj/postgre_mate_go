CREATE DATABASE ttbb encoding 'UTF8';
ALTER DATABASE ttbb set search_path to public;
CREATE role sh login password  'ttlovezj';
GRANT connect on DATABASE ttbb to sh;
CREATE DATABASE thingsboard encoding 'UTF8';
ALTER DATABASE thingsboard set search_path to public;
CREATE role thingsboard login password  'thingsboard';
GRANT connect on DATABASE thingsboard to thingsboard;