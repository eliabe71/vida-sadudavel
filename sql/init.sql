CREATE TABLE medics (
	id serial primary key,
	crm text NOT NULL unique,
	name text,
	lastname text,
	state text,
	city text,
	areaofoccupation text
);

create table Client (
	id serial primary key,
	name text,
	lastname text,
	state text,
	city text, 
	email text NOT NULL unique,
	password text
);
create table clerk (
	id serial  primary key, 
	name text,
	email text unique not null,
	senha text
);
create table appointment(
	id serial primary key ,
	price float,
	medics_fk text references medics(crm),
	client_fk text references Client(email),
	status boolean,
	effected boolean
);
create table employee (
	 id serial primary key,
	cleck_fk int references clerk(id),
	medic_fk text references medics(crm)
);