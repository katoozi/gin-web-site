SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;
CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
SET search_path = public, pg_catalog;
SET default_tablespace = '';
SET default_with_oids = false;

CREATE TABLE IF NOT EXISTS "user" (
    "id" serial not null PRIMARY KEY,
    "first_name" varchar(30),
    "last_name" varchar(150),
    "password" varchar(130) not null,
    "last_login" timestamptz default now(),
    "date_joined" timestamptz default now(),
    "username" varchar(150) unique not null,
    "email" varchar(254),
    "is_active" boolean not null default 'true',
    "is_staff" boolean not null default 'false',
    "is_superuser" boolean not null default 'false'
);

CREATE TABLE IF NOT EXISTS "session" (
    "session_key" varchar(40) not null primary key,
    "session_data" text not null,
    "expire_date" timestamptz not null
);
