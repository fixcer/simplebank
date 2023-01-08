create table if not exists accounts (
    id bigserial primary key,
    owner varchar not null,
    balance bigint not null,
    currency varchar not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create table if not exists entries (
    id bigserial primary key,
    account_id bigint not null references accounts(id),
    amount bigint not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create table if not exists transfers (
    id bigserial primary key,
    from_account_id bigint not null references accounts(id),
    to_account_id bigint not null references accounts(id),
    amount bigint not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create index if not exists accounts_owner_idx on accounts (owner);
create index if not exists entries_account_id_idx on entries(account_id);
create index if not exists transfers_from_account_id_idx on transfers(from_account_id);
create index if not exists transfers_to_account_id_idx on transfers(to_account_id);
create index if not exists transfers_from_account_id_to_account_id_idx on transfers(from_account_id, to_account_id);
