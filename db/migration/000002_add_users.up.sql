create table if not exists users (
    username varchar primary key not null,
    password varchar not null,
    full_name varchar not null,
    email varchar not null,
    password_changed_at timestamptz not null default '0001-01-01 00:00:00Z',
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

alter table accounts add foreign key (owner) references users (username);

alter table accounts add constraint owner_currency_key unique (owner, currency);
