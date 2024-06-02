-- +migrate Up
-- please read this article to understand why we use VARCHAR(191)
-- https://www.grouparoo.com/blog/varchar-191#why-varchar-and-not-text

create TABLE public.users
(
    id bigserial NOT NULL,
    mobile character varying(191) NOT NULL,
    name character varying(191),
    email character varying(191),
    password character varying(191) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now(),
    PRIMARY KEY (id),
    UNIQUE(mobile, email)
);

-- +migrate Down
drop table IF EXISTS users;
