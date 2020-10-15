--
-- PostgreSQL database dump
--

-- Dumped from database version 10.6 (Debian 10.6-1.pgdg90+1)
-- Dumped by pg_dump version 12.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

--
-- Name: additive; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.additive (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    note text NOT NULL
);




--
-- Name: additive_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.additive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: additive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.additive_id_seq OWNED BY public.additive.id;


--
-- Name: additive_inventory; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.additive_inventory (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    purchase_date timestamp with time zone NOT NULL,
    expiry_date timestamp with time zone NOT NULL,
    cost double precision NOT NULL,
    weight double precision NOT NULL,
    additive_id integer NOT NULL,
    supplier_id integer NOT NULL
);




--
-- Name: additive_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.additive_inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: additive_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.additive_inventory_id_seq OWNED BY public.additive_inventory.id;


--
-- Name: auth_group; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_group (
    id integer NOT NULL,
    name character varying(150) NOT NULL
);




--
-- Name: auth_group_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: auth_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_group_id_seq OWNED BY public.auth_group.id;


--
-- Name: auth_group_permissions; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_group_permissions (
    id integer NOT NULL,
    group_id integer NOT NULL,
    permission_id integer NOT NULL
);




--
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_group_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: auth_group_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_group_permissions_id_seq OWNED BY public.auth_group_permissions.id;


--
-- Name: auth_permission; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_permission (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    content_type_id integer NOT NULL,
    codename character varying(100) NOT NULL
);




--
-- Name: auth_permission_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_permission_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: auth_permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_permission_id_seq OWNED BY public.auth_permission.id;


--
-- Name: auth_user; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_user (
    id integer NOT NULL,
    password character varying(128) NOT NULL,
    last_login timestamp with time zone,
    is_superuser boolean NOT NULL,
    username character varying(150) NOT NULL,
    first_name character varying(150) NOT NULL,
    last_name character varying(150) NOT NULL,
    email character varying(254) NOT NULL,
    is_staff boolean NOT NULL,
    is_active boolean NOT NULL,
    date_joined timestamp with time zone NOT NULL
);



--
-- Name: auth_user_groups; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_user_groups (
    id integer NOT NULL,
    user_id integer NOT NULL,
    group_id integer NOT NULL
);



--
-- Name: auth_user_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_user_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: auth_user_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_user_groups_id_seq OWNED BY public.auth_user_groups.id;


--
-- Name: auth_user_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: auth_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_user_id_seq OWNED BY public.auth_user.id;


--
-- Name: auth_user_user_permissions; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.auth_user_user_permissions (
    id integer NOT NULL,
    user_id integer NOT NULL,
    permission_id integer NOT NULL
);



--
-- Name: auth_user_user_permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.auth_user_user_permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: auth_user_user_permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.auth_user_user_permissions_id_seq OWNED BY public.auth_user_user_permissions.id;


--
-- Name: django_admin_log; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.django_admin_log (
    id integer NOT NULL,
    action_time timestamp with time zone NOT NULL,
    object_id text,
    object_repr character varying(200) NOT NULL,
    action_flag smallint NOT NULL,
    change_message text NOT NULL,
    content_type_id integer,
    user_id integer NOT NULL,
    CONSTRAINT django_admin_log_action_flag_check CHECK ((action_flag >= 0))
);



--
-- Name: django_admin_log_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.django_admin_log_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: django_admin_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.django_admin_log_id_seq OWNED BY public.django_admin_log.id;


--
-- Name: django_content_type; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.django_content_type (
    id integer NOT NULL,
    app_label character varying(100) NOT NULL,
    model character varying(100) NOT NULL
);



--
-- Name: django_content_type_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.django_content_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: django_content_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.django_content_type_id_seq OWNED BY public.django_content_type.id;


--
-- Name: django_migrations; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.django_migrations (
    id integer NOT NULL,
    app character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    applied timestamp with time zone NOT NULL
);



--
-- Name: django_migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.django_migrations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: django_migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.django_migrations_id_seq OWNED BY public.django_migrations.id;


--
-- Name: django_session; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.django_session (
    session_key character varying(40) NOT NULL,
    session_data text NOT NULL,
    expire_date timestamp with time zone NOT NULL
);



--
-- Name: fragrance; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.fragrance (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    note text NOT NULL
);



--
-- Name: fragrance_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.fragrance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: fragrance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.fragrance_id_seq OWNED BY public.fragrance.id;


--
-- Name: fragrance_inventory; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.fragrance_inventory (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    purchase_date timestamp with time zone NOT NULL,
    expiry_date timestamp with time zone NOT NULL,
    cost double precision NOT NULL,
    weight double precision NOT NULL,
    fragrance_id integer NOT NULL,
    supplier_id integer NOT NULL
);



--
-- Name: fragrance_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.fragrance_inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: fragrance_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.fragrance_inventory_id_seq OWNED BY public.fragrance_inventory.id;


--
-- Name: lipid; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.lipid (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    lauric integer NOT NULL,
    myristic integer NOT NULL,
    palmitic integer NOT NULL,
    stearic integer NOT NULL,
    ricinoleic integer NOT NULL,
    oleic integer NOT NULL,
    linoleic integer NOT NULL,
    linolenic integer NOT NULL,
    hardness integer NOT NULL,
    cleansing integer NOT NULL,
    conditioning integer NOT NULL,
    bubbly integer NOT NULL,
    creamy integer NOT NULL,
    iodine integer NOT NULL,
    ins integer NOT NULL,
    inci_name character varying(100) NOT NULL,
    family character varying(50) NOT NULL,
    naoh double precision NOT NULL
);



--
-- Name: lipid_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.lipid_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: lipid_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.lipid_id_seq OWNED BY public.lipid.id;


--
-- Name: lipid_inventory; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.lipid_inventory (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    purchase_date timestamp with time zone NOT NULL,
    expiry_date timestamp with time zone NOT NULL,
    cost double precision NOT NULL,
    weight double precision NOT NULL,
    sap double precision NOT NULL,
    naoh double precision NOT NULL,
    koh double precision NOT NULL,
    grams_per_liter double precision NOT NULL,
    lipid_id integer NOT NULL,
    supplier_id integer NOT NULL
);



--
-- Name: lipid_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.lipid_inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: lipid_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.lipid_inventory_id_seq OWNED BY public.lipid_inventory.id;


--
-- Name: lye; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.lye (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    kind character varying(4) NOT NULL,
    name character varying(100) NOT NULL,
    note text NOT NULL
);



--
-- Name: lye_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.lye_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: lye_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.lye_id_seq OWNED BY public.lye.id;


--
-- Name: lye_inventory; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.lye_inventory (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    purchase_date timestamp with time zone NOT NULL,
    expiry_date timestamp with time zone NOT NULL,
    cost double precision NOT NULL,
    weight double precision NOT NULL,
    concentration double precision NOT NULL,
    lye_id integer NOT NULL,
    supplier_id integer NOT NULL
);



--
-- Name: lye_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.lye_inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: lye_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.lye_inventory_id_seq OWNED BY public.lye_inventory.id;


--
-- Name: recipe; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    note text NOT NULL
);




--
-- Name: recipe_additive; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_additive (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    percentage double precision NOT NULL,
    additive_id integer NOT NULL,
    recipe_id integer NOT NULL
);



--
-- Name: recipe_additive_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_additive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: recipe_additive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_additive_id_seq OWNED BY public.recipe_additive.id;


--
-- Name: recipe_batch; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    tag character varying(16) NOT NULL,
    production_date timestamp with time zone NOT NULL,
    sellable_date timestamp with time zone NOT NULL,
    note text NOT NULL,
    lipid_weight double precision NOT NULL,
    production_weight double precision NOT NULL,
    cured_weight double precision NOT NULL,
    recipe_id integer NOT NULL
);




--
-- Name: recipe_batch_additive; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_additive (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight double precision NOT NULL,
    cost double precision NOT NULL,
    additive_id integer NOT NULL,
    batch_id integer NOT NULL
);




--
-- Name: recipe_batch_additive_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_additive_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_batch_additive_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_additive_id_seq OWNED BY public.recipe_batch_additive.id;


--
-- Name: recipe_batch_fragrance; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_fragrance (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight double precision NOT NULL,
    cost double precision NOT NULL,
    fragrance_id integer NOT NULL,
    batch_id integer NOT NULL
);




--
-- Name: recipe_batch_fragrance_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_fragrance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_batch_fragrance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_fragrance_id_seq OWNED BY public.recipe_batch_fragrance.id;


--
-- Name: recipe_batch_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_batch_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_id_seq OWNED BY public.recipe_batch.id;


--
-- Name: recipe_batch_lipid; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_lipid (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight double precision NOT NULL,
    cost double precision NOT NULL,
    lipid_id integer NOT NULL,
    batch_id integer NOT NULL
);




--
-- Name: recipe_batch_lipid_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_lipid_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_batch_lipid_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_lipid_id_seq OWNED BY public.recipe_batch_lipid.id;


--
-- Name: recipe_batch_lye; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_lye (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    weight double precision NOT NULL,
    discount double precision NOT NULL,
    cost double precision NOT NULL,
    lye_id integer NOT NULL,
    batch_id integer NOT NULL
);




--
-- Name: recipe_batch_lye_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_lye_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_batch_lye_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_lye_id_seq OWNED BY public.recipe_batch_lye.id;


--
-- Name: recipe_batch_note; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_batch_note (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    note text NOT NULL,
    link character varying(255) NOT NULL,
    batch_id integer NOT NULL
);




--
-- Name: recipe_batch_note_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_batch_note_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_batch_note_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_batch_note_id_seq OWNED BY public.recipe_batch_note.id;


--
-- Name: recipe_fragrance; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_fragrance (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    percentage double precision NOT NULL,
    fragrance_id integer NOT NULL,
    recipe_id integer NOT NULL
);




--
-- Name: recipe_fragrance_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_fragrance_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_fragrance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_fragrance_id_seq OWNED BY public.recipe_fragrance.id;


--
-- Name: recipe_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_id_seq OWNED BY public.recipe.id;


--
-- Name: recipe_lipid; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_lipid (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    percentage double precision NOT NULL,
    lipid_id integer NOT NULL,
    recipe_id integer NOT NULL
);




--
-- Name: recipe_lipid_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_lipid_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_lipid_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_lipid_id_seq OWNED BY public.recipe_lipid.id;


--
-- Name: recipe_step; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.recipe_step (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    num integer NOT NULL,
    note text NOT NULL,
    recipe_id integer NOT NULL
);




--
-- Name: recipe_step_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.recipe_step_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: recipe_step_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.recipe_step_id_seq OWNED BY public.recipe_step.id;


--
-- Name: supplier; Type: TABLE; Schema: public; Owner: soap
--

CREATE TABLE public.supplier (
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    website character varying(255) NOT NULL,
    note text NOT NULL
);




--
-- Name: supplier_id_seq; Type: SEQUENCE; Schema: public; Owner: soap
--

CREATE SEQUENCE public.supplier_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




--
-- Name: supplier_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: soap
--

ALTER SEQUENCE public.supplier_id_seq OWNED BY public.supplier.id;


--
-- Name: additive id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive ALTER COLUMN id SET DEFAULT nextval('public.additive_id_seq'::regclass);


--
-- Name: additive_inventory id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive_inventory ALTER COLUMN id SET DEFAULT nextval('public.additive_inventory_id_seq'::regclass);


--
-- Name: auth_group id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group ALTER COLUMN id SET DEFAULT nextval('public.auth_group_id_seq'::regclass);


--
-- Name: auth_group_permissions id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_group_permissions_id_seq'::regclass);


--
-- Name: auth_permission id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_permission ALTER COLUMN id SET DEFAULT nextval('public.auth_permission_id_seq'::regclass);


--
-- Name: auth_user id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user ALTER COLUMN id SET DEFAULT nextval('public.auth_user_id_seq'::regclass);


--
-- Name: auth_user_groups id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups ALTER COLUMN id SET DEFAULT nextval('public.auth_user_groups_id_seq'::regclass);


--
-- Name: auth_user_user_permissions id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions ALTER COLUMN id SET DEFAULT nextval('public.auth_user_user_permissions_id_seq'::regclass);


--
-- Name: django_admin_log id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_admin_log ALTER COLUMN id SET DEFAULT nextval('public.django_admin_log_id_seq'::regclass);


--
-- Name: django_content_type id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_content_type ALTER COLUMN id SET DEFAULT nextval('public.django_content_type_id_seq'::regclass);


--
-- Name: django_migrations id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_migrations ALTER COLUMN id SET DEFAULT nextval('public.django_migrations_id_seq'::regclass);


--
-- Name: fragrance id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance ALTER COLUMN id SET DEFAULT nextval('public.fragrance_id_seq'::regclass);


--
-- Name: fragrance_inventory id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance_inventory ALTER COLUMN id SET DEFAULT nextval('public.fragrance_inventory_id_seq'::regclass);


--
-- Name: lipid id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid ALTER COLUMN id SET DEFAULT nextval('public.lipid_id_seq'::regclass);


--
-- Name: lipid_inventory id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid_inventory ALTER COLUMN id SET DEFAULT nextval('public.lipid_inventory_id_seq'::regclass);


--
-- Name: lye id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye ALTER COLUMN id SET DEFAULT nextval('public.lye_id_seq'::regclass);


--
-- Name: lye_inventory id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye_inventory ALTER COLUMN id SET DEFAULT nextval('public.lye_inventory_id_seq'::regclass);


--
-- Name: recipe id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe ALTER COLUMN id SET DEFAULT nextval('public.recipe_id_seq'::regclass);


--
-- Name: recipe_additive id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive ALTER COLUMN id SET DEFAULT nextval('public.recipe_additive_id_seq'::regclass);


--
-- Name: recipe_batch id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_id_seq'::regclass);


--
-- Name: recipe_batch_additive id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_additive_id_seq'::regclass);


--
-- Name: recipe_batch_fragrance id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_fragrance_id_seq'::regclass);


--
-- Name: recipe_batch_lipid id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_lipid_id_seq'::regclass);


--
-- Name: recipe_batch_lye id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_lye_id_seq'::regclass);


--
-- Name: recipe_batch_note id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_note ALTER COLUMN id SET DEFAULT nextval('public.recipe_batch_note_id_seq'::regclass);


--
-- Name: recipe_fragrance id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance ALTER COLUMN id SET DEFAULT nextval('public.recipe_fragrance_id_seq'::regclass);


--
-- Name: recipe_lipid id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid ALTER COLUMN id SET DEFAULT nextval('public.recipe_lipid_id_seq'::regclass);


--
-- Name: recipe_step id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_step ALTER COLUMN id SET DEFAULT nextval('public.recipe_step_id_seq'::regclass);


--
-- Name: supplier id; Type: DEFAULT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.supplier ALTER COLUMN id SET DEFAULT nextval('public.supplier_id_seq'::regclass);


--
-- Name: additive_inventory additive_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive_inventory
    ADD CONSTRAINT additive_inventory_pkey PRIMARY KEY (id);


--
-- Name: additive additive_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive
    ADD CONSTRAINT additive_pkey PRIMARY KEY (id);


--
-- Name: auth_group auth_group_name_key; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_name_key UNIQUE (name);


--
-- Name: auth_group_permissions auth_group_permissions_group_id_permission_id_0cd325b0_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_permission_id_0cd325b0_uniq UNIQUE (group_id, permission_id);


--
-- Name: auth_group_permissions auth_group_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_pkey PRIMARY KEY (id);


--
-- Name: auth_group auth_group_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group
    ADD CONSTRAINT auth_group_pkey PRIMARY KEY (id);


--
-- Name: auth_permission auth_permission_content_type_id_codename_01ab375a_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_codename_01ab375a_uniq UNIQUE (content_type_id, codename);


--
-- Name: auth_permission auth_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_pkey PRIMARY KEY (id);


--
-- Name: auth_user_groups auth_user_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_pkey PRIMARY KEY (id);


--
-- Name: auth_user_groups auth_user_groups_user_id_group_id_94350c0c_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_user_id_group_id_94350c0c_uniq UNIQUE (user_id, group_id);


--
-- Name: auth_user auth_user_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_pkey PRIMARY KEY (id);


--
-- Name: auth_user_user_permissions auth_user_user_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_pkey PRIMARY KEY (id);


--
-- Name: auth_user_user_permissions auth_user_user_permissions_user_id_permission_id_14a6b632_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_user_id_permission_id_14a6b632_uniq UNIQUE (user_id, permission_id);


--
-- Name: auth_user auth_user_username_key; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user
    ADD CONSTRAINT auth_user_username_key UNIQUE (username);


--
-- Name: django_admin_log django_admin_log_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_pkey PRIMARY KEY (id);


--
-- Name: django_content_type django_content_type_app_label_model_76bd3d3b_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_app_label_model_76bd3d3b_uniq UNIQUE (app_label, model);


--
-- Name: django_content_type django_content_type_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_content_type
    ADD CONSTRAINT django_content_type_pkey PRIMARY KEY (id);


--
-- Name: django_migrations django_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_migrations
    ADD CONSTRAINT django_migrations_pkey PRIMARY KEY (id);


--
-- Name: django_session django_session_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_session
    ADD CONSTRAINT django_session_pkey PRIMARY KEY (session_key);


--
-- Name: fragrance_inventory fragrance_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance_inventory
    ADD CONSTRAINT fragrance_inventory_pkey PRIMARY KEY (id);


--
-- Name: fragrance fragrance_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance
    ADD CONSTRAINT fragrance_pkey PRIMARY KEY (id);


--
-- Name: lipid_inventory lipid_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid_inventory
    ADD CONSTRAINT lipid_inventory_pkey PRIMARY KEY (id);


--
-- Name: lipid lipid_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid
    ADD CONSTRAINT lipid_pkey PRIMARY KEY (id);


--
-- Name: lye_inventory lye_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye_inventory
    ADD CONSTRAINT lye_inventory_pkey PRIMARY KEY (id);


--
-- Name: lye lye_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye
    ADD CONSTRAINT lye_pkey PRIMARY KEY (id);


--
-- Name: recipe_additive recipe_additive_additive_id_676ee4ae_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive
    ADD CONSTRAINT recipe_additive_additive_id_676ee4ae_uniq UNIQUE (additive_id);


--
-- Name: recipe_additive recipe_additive_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive
    ADD CONSTRAINT recipe_additive_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_additive recipe_batch_additive_additive_id_2361c625_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive
    ADD CONSTRAINT recipe_batch_additive_additive_id_2361c625_uniq UNIQUE (additive_id);


--
-- Name: recipe_batch_additive recipe_batch_additive_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive
    ADD CONSTRAINT recipe_batch_additive_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_fragrance recipe_batch_fragrance_fragrance_id_c6546fe6_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance
    ADD CONSTRAINT recipe_batch_fragrance_fragrance_id_c6546fe6_uniq UNIQUE (fragrance_id);


--
-- Name: recipe_batch_fragrance recipe_batch_fragrance_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance
    ADD CONSTRAINT recipe_batch_fragrance_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_lipid recipe_batch_lipid_lipid_id_d8453e8a_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid
    ADD CONSTRAINT recipe_batch_lipid_lipid_id_d8453e8a_uniq UNIQUE (lipid_id);


--
-- Name: recipe_batch_lipid recipe_batch_lipid_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid
    ADD CONSTRAINT recipe_batch_lipid_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_lye recipe_batch_lye_lye_id_784d24d9_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye
    ADD CONSTRAINT recipe_batch_lye_lye_id_784d24d9_uniq UNIQUE (lye_id);


--
-- Name: recipe_batch_lye recipe_batch_lye_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye
    ADD CONSTRAINT recipe_batch_lye_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch_note recipe_batch_note_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_note
    ADD CONSTRAINT recipe_batch_note_pkey PRIMARY KEY (id);


--
-- Name: recipe_batch recipe_batch_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch
    ADD CONSTRAINT recipe_batch_pkey PRIMARY KEY (id);


--
-- Name: recipe_fragrance recipe_fragrance_fragrance_id_0dcf3ef6_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance
    ADD CONSTRAINT recipe_fragrance_fragrance_id_0dcf3ef6_uniq UNIQUE (fragrance_id);


--
-- Name: recipe_fragrance recipe_fragrance_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance
    ADD CONSTRAINT recipe_fragrance_pkey PRIMARY KEY (id);


--
-- Name: recipe_lipid recipe_lipid_lipid_id_76650ba8_uniq; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid
    ADD CONSTRAINT recipe_lipid_lipid_id_76650ba8_uniq UNIQUE (lipid_id);


--
-- Name: recipe_lipid recipe_lipid_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid
    ADD CONSTRAINT recipe_lipid_pkey PRIMARY KEY (id);


--
-- Name: recipe recipe_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe
    ADD CONSTRAINT recipe_pkey PRIMARY KEY (id);


--
-- Name: recipe_step recipe_step_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_step
    ADD CONSTRAINT recipe_step_pkey PRIMARY KEY (id);


--
-- Name: supplier supplier_pkey; Type: CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.supplier
    ADD CONSTRAINT supplier_pkey PRIMARY KEY (id);


--
-- Name: additive_inventory_additive_id_390dfc35; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX additive_inventory_additive_id_390dfc35 ON public.additive_inventory USING btree (additive_id);


--
-- Name: additive_inventory_supplier_id_dc5c2c7b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX additive_inventory_supplier_id_dc5c2c7b ON public.additive_inventory USING btree (supplier_id);


--
-- Name: auth_group_name_a6ea08ec_like; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_group_name_a6ea08ec_like ON public.auth_group USING btree (name varchar_pattern_ops);


--
-- Name: auth_group_permissions_group_id_b120cbf9; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_group_permissions_group_id_b120cbf9 ON public.auth_group_permissions USING btree (group_id);


--
-- Name: auth_group_permissions_permission_id_84c5c92e; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_group_permissions_permission_id_84c5c92e ON public.auth_group_permissions USING btree (permission_id);


--
-- Name: auth_permission_content_type_id_2f476e4b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_permission_content_type_id_2f476e4b ON public.auth_permission USING btree (content_type_id);


--
-- Name: auth_user_groups_group_id_97559544; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_groups_group_id_97559544 ON public.auth_user_groups USING btree (group_id);


--
-- Name: auth_user_groups_user_id_6a12ed8b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_groups_user_id_6a12ed8b ON public.auth_user_groups USING btree (user_id);


--
-- Name: auth_user_user_permissions_permission_id_1fbb5f2c; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_user_permissions_permission_id_1fbb5f2c ON public.auth_user_user_permissions USING btree (permission_id);


--
-- Name: auth_user_user_permissions_user_id_a95ead1b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_user_permissions_user_id_a95ead1b ON public.auth_user_user_permissions USING btree (user_id);


--
-- Name: auth_user_username_6821ab7c_like; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX auth_user_username_6821ab7c_like ON public.auth_user USING btree (username varchar_pattern_ops);


--
-- Name: django_admin_log_content_type_id_c4bce8eb; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX django_admin_log_content_type_id_c4bce8eb ON public.django_admin_log USING btree (content_type_id);


--
-- Name: django_admin_log_user_id_c564eba6; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX django_admin_log_user_id_c564eba6 ON public.django_admin_log USING btree (user_id);


--
-- Name: django_session_expire_date_a5c62663; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX django_session_expire_date_a5c62663 ON public.django_session USING btree (expire_date);


--
-- Name: django_session_session_key_c0390e0f_like; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX django_session_session_key_c0390e0f_like ON public.django_session USING btree (session_key varchar_pattern_ops);


--
-- Name: fragrance_inventory_fragrance_id_9f202030; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX fragrance_inventory_fragrance_id_9f202030 ON public.fragrance_inventory USING btree (fragrance_id);


--
-- Name: fragrance_inventory_supplier_id_90b4caaf; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX fragrance_inventory_supplier_id_90b4caaf ON public.fragrance_inventory USING btree (supplier_id);


--
-- Name: lipid_inventory_lipid_id_4c032624; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX lipid_inventory_lipid_id_4c032624 ON public.lipid_inventory USING btree (lipid_id);


--
-- Name: lipid_inventory_supplier_id_62914da1; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX lipid_inventory_supplier_id_62914da1 ON public.lipid_inventory USING btree (supplier_id);


--
-- Name: lye_inventory_lye_id_5e9da65f; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX lye_inventory_lye_id_5e9da65f ON public.lye_inventory USING btree (lye_id);


--
-- Name: lye_inventory_supplier_id_78e9941f; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX lye_inventory_supplier_id_78e9941f ON public.lye_inventory USING btree (supplier_id);


--
-- Name: recipe_additive_recipe_id_48b68995; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_additive_recipe_id_48b68995 ON public.recipe_additive USING btree (recipe_id);


--
-- Name: recipe_batch_additive_batch_id_d26265ff; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_additive_batch_id_d26265ff ON public.recipe_batch_additive USING btree (batch_id);


--
-- Name: recipe_batch_fragrance_batch_id_a36d24d9; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_fragrance_batch_id_a36d24d9 ON public.recipe_batch_fragrance USING btree (batch_id);


--
-- Name: recipe_batch_lipid_batch_id_b292008e; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_lipid_batch_id_b292008e ON public.recipe_batch_lipid USING btree (batch_id);


--
-- Name: recipe_batch_lye_batch_id_db6fa60b; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_lye_batch_id_db6fa60b ON public.recipe_batch_lye USING btree (batch_id);


--
-- Name: recipe_batch_note_batch_id_46a82deb; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_note_batch_id_46a82deb ON public.recipe_batch_note USING btree (batch_id);


--
-- Name: recipe_batch_recipe_id_60edb3ae; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_batch_recipe_id_60edb3ae ON public.recipe_batch USING btree (recipe_id);


--
-- Name: recipe_fragrance_recipe_id_404ec2b0; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_fragrance_recipe_id_404ec2b0 ON public.recipe_fragrance USING btree (recipe_id);


--
-- Name: recipe_lipid_recipe_id_d2a52df1; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_lipid_recipe_id_d2a52df1 ON public.recipe_lipid USING btree (recipe_id);


--
-- Name: recipe_step_recipe_id_bb16b3a0; Type: INDEX; Schema: public; Owner: soap
--

CREATE INDEX recipe_step_recipe_id_bb16b3a0 ON public.recipe_step USING btree (recipe_id);


--
-- Name: additive_inventory additive_inventory_additive_id_390dfc35_fk_additive_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive_inventory
    ADD CONSTRAINT additive_inventory_additive_id_390dfc35_fk_additive_id FOREIGN KEY (additive_id) REFERENCES public.additive(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: additive_inventory additive_inventory_supplier_id_dc5c2c7b_fk_supplier_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.additive_inventory
    ADD CONSTRAINT additive_inventory_supplier_id_dc5c2c7b_fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES public.supplier(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_group_permissions auth_group_permissio_permission_id_84c5c92e_fk_auth_perm; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissio_permission_id_84c5c92e_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_group_permissions auth_group_permissions_group_id_b120cbf9_fk_auth_group_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_group_permissions
    ADD CONSTRAINT auth_group_permissions_group_id_b120cbf9_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_permission auth_permission_content_type_id_2f476e4b_fk_django_co; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_permission
    ADD CONSTRAINT auth_permission_content_type_id_2f476e4b_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_user_groups auth_user_groups_group_id_97559544_fk_auth_group_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_group_id_97559544_fk_auth_group_id FOREIGN KEY (group_id) REFERENCES public.auth_group(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_user_groups auth_user_groups_user_id_6a12ed8b_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_groups
    ADD CONSTRAINT auth_user_groups_user_id_6a12ed8b_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_user_user_permissions auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permi_permission_id_1fbb5f2c_fk_auth_perm FOREIGN KEY (permission_id) REFERENCES public.auth_permission(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: auth_user_user_permissions auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.auth_user_user_permissions
    ADD CONSTRAINT auth_user_user_permissions_user_id_a95ead1b_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: django_admin_log django_admin_log_content_type_id_c4bce8eb_fk_django_co; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_content_type_id_c4bce8eb_fk_django_co FOREIGN KEY (content_type_id) REFERENCES public.django_content_type(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: django_admin_log django_admin_log_user_id_c564eba6_fk_auth_user_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.django_admin_log
    ADD CONSTRAINT django_admin_log_user_id_c564eba6_fk_auth_user_id FOREIGN KEY (user_id) REFERENCES public.auth_user(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: fragrance_inventory fragrance_inventory_fragrance_id_9f202030_fk_fragrance_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance_inventory
    ADD CONSTRAINT fragrance_inventory_fragrance_id_9f202030_fk_fragrance_id FOREIGN KEY (fragrance_id) REFERENCES public.fragrance(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: fragrance_inventory fragrance_inventory_supplier_id_90b4caaf_fk_supplier_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.fragrance_inventory
    ADD CONSTRAINT fragrance_inventory_supplier_id_90b4caaf_fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES public.supplier(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: lipid_inventory lipid_inventory_lipid_id_4c032624_fk_lipid_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid_inventory
    ADD CONSTRAINT lipid_inventory_lipid_id_4c032624_fk_lipid_id FOREIGN KEY (lipid_id) REFERENCES public.lipid(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: lipid_inventory lipid_inventory_supplier_id_62914da1_fk_supplier_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lipid_inventory
    ADD CONSTRAINT lipid_inventory_supplier_id_62914da1_fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES public.supplier(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: lye_inventory lye_inventory_lye_id_5e9da65f_fk_lye_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye_inventory
    ADD CONSTRAINT lye_inventory_lye_id_5e9da65f_fk_lye_id FOREIGN KEY (lye_id) REFERENCES public.lye(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: lye_inventory lye_inventory_supplier_id_78e9941f_fk_supplier_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.lye_inventory
    ADD CONSTRAINT lye_inventory_supplier_id_78e9941f_fk_supplier_id FOREIGN KEY (supplier_id) REFERENCES public.supplier(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_additive recipe_additive_additive_id_676ee4ae_fk_additive_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive
    ADD CONSTRAINT recipe_additive_additive_id_676ee4ae_fk_additive_id FOREIGN KEY (additive_id) REFERENCES public.additive(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_additive recipe_additive_recipe_id_48b68995_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_additive
    ADD CONSTRAINT recipe_additive_recipe_id_48b68995_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_additive recipe_batch_additive_additive_id_2361c625_fk_additive_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive
    ADD CONSTRAINT recipe_batch_additive_additive_id_2361c625_fk_additive_id FOREIGN KEY (additive_id) REFERENCES public.additive(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_additive recipe_batch_additive_batch_id_d26265ff_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_additive
    ADD CONSTRAINT recipe_batch_additive_batch_id_d26265ff_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_fragrance recipe_batch_fragrance_batch_id_a36d24d9_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance
    ADD CONSTRAINT recipe_batch_fragrance_batch_id_a36d24d9_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_fragrance recipe_batch_fragrance_fragrance_id_c6546fe6_fk_fragrance_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_fragrance
    ADD CONSTRAINT recipe_batch_fragrance_fragrance_id_c6546fe6_fk_fragrance_id FOREIGN KEY (fragrance_id) REFERENCES public.fragrance(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_lipid recipe_batch_lipid_batch_id_b292008e_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid
    ADD CONSTRAINT recipe_batch_lipid_batch_id_b292008e_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_lipid recipe_batch_lipid_lipid_id_d8453e8a_fk_lipid_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lipid
    ADD CONSTRAINT recipe_batch_lipid_lipid_id_d8453e8a_fk_lipid_id FOREIGN KEY (lipid_id) REFERENCES public.lipid(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_lye recipe_batch_lye_batch_id_db6fa60b_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye
    ADD CONSTRAINT recipe_batch_lye_batch_id_db6fa60b_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_lye recipe_batch_lye_lye_id_784d24d9_fk_lye_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_lye
    ADD CONSTRAINT recipe_batch_lye_lye_id_784d24d9_fk_lye_id FOREIGN KEY (lye_id) REFERENCES public.lye(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch_note recipe_batch_note_batch_id_46a82deb_fk_recipe_batch_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch_note
    ADD CONSTRAINT recipe_batch_note_batch_id_46a82deb_fk_recipe_batch_id FOREIGN KEY (batch_id) REFERENCES public.recipe_batch(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_batch recipe_batch_recipe_id_60edb3ae_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_batch
    ADD CONSTRAINT recipe_batch_recipe_id_60edb3ae_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_fragrance recipe_fragrance_fragrance_id_0dcf3ef6_fk_fragrance_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance
    ADD CONSTRAINT recipe_fragrance_fragrance_id_0dcf3ef6_fk_fragrance_id FOREIGN KEY (fragrance_id) REFERENCES public.fragrance(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_fragrance recipe_fragrance_recipe_id_404ec2b0_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_fragrance
    ADD CONSTRAINT recipe_fragrance_recipe_id_404ec2b0_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_lipid recipe_lipid_lipid_id_76650ba8_fk_lipid_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid
    ADD CONSTRAINT recipe_lipid_lipid_id_76650ba8_fk_lipid_id FOREIGN KEY (lipid_id) REFERENCES public.lipid(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_lipid recipe_lipid_recipe_id_d2a52df1_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_lipid
    ADD CONSTRAINT recipe_lipid_recipe_id_d2a52df1_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- Name: recipe_step recipe_step_recipe_id_bb16b3a0_fk_recipe_id; Type: FK CONSTRAINT; Schema: public; Owner: soap
--

ALTER TABLE ONLY public.recipe_step
    ADD CONSTRAINT recipe_step_recipe_id_bb16b3a0_fk_recipe_id FOREIGN KEY (recipe_id) REFERENCES public.recipe(id) DEFERRABLE INITIALLY DEFERRED;


--
-- PostgreSQL database dump complete
--

