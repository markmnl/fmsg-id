/****************************************************************
 *
 * PostgreSQL database objects data definition for fmsgid
 *
 ****************************************************************/

-- database with encoding UTF8 should already be created and connected

create extension if not exists citext;

create table if not exists address (
    address 					citext 	primary key,
    display_name 				text,
    accepting_new 				bool	not null default true,
    limit_recv_size_total 		bigint	not null default -1,
    limit_recv_size_per_msg 	bigint	not null default -1,
    limit_recv_size_per_1d 		bigint	not null default -1,
    limit_recv_count_per_1d 	bigint	not null default -1,
    limit_send_size_total 		bigint	not null default -1,
    limit_send_size_per_msg 	bigint	not null default -1,
    limit_send_size_per_1d 		bigint	not null default -1,
    limit_send_count_per_1d 	bigint	not null default -1
);

-- TODO consider time-series datastore e.g. TimescaleDB, InfluxDb ...
create table if not exists address_tx (
	address	citext					not null references address (address),
	ts		timestamptz				not null,
	op		varchar(5)				not null, -- send, recv
	size	int						not null,
	primary key (address, ts)
);