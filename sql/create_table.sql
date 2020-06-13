CREATE TABLE product(
	id uuid,
	name character varying(1000) not null,
	type character varying(1000) not null,

	create_time timestamp with time zone not null default current_timestamp,
	update_time timestamp with time zone not null default current_timestamp,
	CONSTRAINT "p_pk" PRIMARY KEY (id)
);
