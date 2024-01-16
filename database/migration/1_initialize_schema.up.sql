-- Table: public.users

-- DROP TABLE IF EXISTS public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL,
    account character varying(64) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    password character varying(64) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    name character varying(16) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    avatar character varying(64) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    created_date timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;



-- Table: public.groups

-- DROP TABLE IF EXISTS public.groups;

CREATE TABLE IF NOT EXISTS public.groups
(
    id bigint NOT NULL,
    group_name character varying(64) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    created_date timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT groups_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.groups
    OWNER to postgres;



-- Table: public.user_group

-- DROP TABLE IF EXISTS public.user_group;

CREATE TABLE IF NOT EXISTS public.user_group
(
    id SERIAL NOT NULL,
    group_id bigint NOT NULL,
    user_id bigint NOT NULL,
    CONSTRAINT user_group_pkey PRIMARY KEY (id),
    CONSTRAINT fk_user_group_group FOREIGN KEY (group_id)
        REFERENCES public.groups (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_user_group_user FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_group
    OWNER to postgres;



-- Table: public.activities

-- DROP TABLE IF EXISTS public.activities;

CREATE TABLE IF NOT EXISTS public.activities
(
    id bigint NOT NULL,
    group_id bigint NOT NULL,
    creator_id bigint NOT NULL,
    activity_name character varying(64) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    created_date timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT activities_pkey PRIMARY KEY (id),
    CONSTRAINT fk_activities_group FOREIGN KEY (group_id)
        REFERENCES public.groups (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_activities_user FOREIGN KEY (creator_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.activities
    OWNER to postgres;



-- Table: public.schedules

-- DROP TABLE IF EXISTS public.schedules;

CREATE TABLE IF NOT EXISTS public.schedules
(
    id bigint NOT NULL,
    activity_id bigint NOT NULL,
    name character varying(64) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    comment character varying(64) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    start_date timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_date timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_date timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT schedules_pkey PRIMARY KEY (id),
    CONSTRAINT fk_schedules_activity FOREIGN KEY (activity_id)
        REFERENCES public.activities (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.schedules
    OWNER to postgres;


 -- Table: public.user_schedule

-- DROP TABLE IF EXISTS public.user_schedule;

CREATE TABLE IF NOT EXISTS public.user_schedule
(
    id SERIAL NOT NULL,
    schedule_id bigint NOT NULL,
    user_id bigint NOT NULL,
    CONSTRAINT user_schedule_pkey PRIMARY KEY (id),
    CONSTRAINT fk_user_schedule_schedule FOREIGN KEY (schedule_id)
        REFERENCES public.schedules (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_user_schedule_user FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.user_schedule
    OWNER to postgres;   