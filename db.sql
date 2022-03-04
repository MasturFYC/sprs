--
-- PostgreSQL database dump
--

-- Dumped from database version 13.5 (Debian 13.5-0+deb11u1)
-- Dumped by pg_dump version 13.5 (Debian 13.5-0+deb11u1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'SQL_ASCII';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: action_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.action_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.action_id_seq OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.actions (
    id integer DEFAULT nextval('public.action_id_seq'::regclass) NOT NULL,
    action_at date NOT NULL,
    code character varying(50) NOT NULL,
    pic character varying(50) NOT NULL,
    descriptions character varying(128),
    order_id integer NOT NULL
);


ALTER TABLE public.actions OWNER TO postgres;

--
-- Name: branch_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.branch_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.branch_id_seq OWNER TO postgres;

--
-- Name: branchs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.branchs (
    id smallint DEFAULT nextval('public.branch_id_seq'::regclass) NOT NULL,
    name character varying(50) NOT NULL,
    street character varying(128),
    city character varying(50),
    phone character varying(25),
    cell character varying(25),
    zip character varying(10),
    head_branch character varying(50) NOT NULL,
    email character varying(50)
);


ALTER TABLE public.branchs OWNER TO postgres;

--
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    order_id integer NOT NULL,
    name character varying(50) NOT NULL,
    agreement_number character varying(25),
    payment_type character varying(25) NOT NULL
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- Name: finance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.finance_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.finance_id_seq OWNER TO postgres;

--
-- Name: finances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.finances (
    id smallint DEFAULT nextval('public.finance_id_seq'::regclass) NOT NULL,
    name character varying(50) NOT NULL,
    short_name character varying(25) NOT NULL,
    street character varying(128),
    city character varying(50),
    phone character varying(25),
    cell character varying(25),
    zip character varying(10),
    email character varying(50)
);


ALTER TABLE public.finances OWNER TO postgres;

--
-- Name: home_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.home_addresses (
    order_id integer NOT NULL,
    street character varying(128),
    region character varying(50),
    city character varying(50),
    phone character varying(25),
    zip character varying(10)
);


ALTER TABLE public.home_addresses OWNER TO postgres;

--
-- Name: ktp_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ktp_addresses (
    order_id integer NOT NULL,
    street character varying(128),
    region character varying(50),
    city character varying(50),
    phone character varying(25),
    zip character varying(10)
);


ALTER TABLE public.ktp_addresses OWNER TO postgres;

--
-- Name: merks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.merks (
    id smallint NOT NULL,
    name character varying(25) NOT NULL
);


ALTER TABLE public.merks OWNER TO postgres;

--
-- Name: merk_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.merk_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.merk_id_seq OWNER TO postgres;

--
-- Name: merk_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.merk_id_seq OWNED BY public.merks.id;


--
-- Name: office_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.office_addresses (
    order_id integer NOT NULL,
    street character varying(128),
    region character varying(50),
    city character varying(50),
    phone character varying(25),
    zip character varying(10)
);


ALTER TABLE public.office_addresses OWNER TO postgres;

--
-- Name: order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_id_seq OWNER TO postgres;

--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id integer DEFAULT nextval('public.order_id_seq'::regclass) NOT NULL,
    name character varying(50) NOT NULL,
    order_at date NOT NULL,
    printed_at date NOT NULL,
    bt_finance numeric(12,2) DEFAULT 0 NOT NULL,
    bt_percent numeric(5,2) DEFAULT 0 NOT NULL,
    bt_matel numeric(12,2) DEFAULT 0 NOT NULL,
    ppn numeric(5,2) DEFAULT 0 NOT NULL,
    user_name character varying(50) NOT NULL,
    verified_by character varying(50),
    validated_by character varying(50),
    finance_id smallint DEFAULT 0 NOT NULL,
    branch_id smallint DEFAULT 0 NOT NULL,
    nominal numeric(12,2) DEFAULT 0 NOT NULL,
    subtotal numeric(12,2) DEFAULT 0 NOT NULL,
    is_stnk boolean DEFAULT true NOT NULL,
    stnk_price numeric(12,2) DEFAULT 0 NOT NULL
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: post_addresses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_addresses (
    order_id integer NOT NULL,
    street character varying(128),
    region character varying(50),
    city character varying(50),
    phone character varying(25),
    zip character varying(10)
);


ALTER TABLE public.post_addresses OWNER TO postgres;

--
-- Name: receivables; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.receivables (
    order_id integer NOT NULL,
    covenant_at date,
    due_at date,
    mortgage_by_month numeric(12,2) DEFAULT 0 NOT NULL,
    mortgage_receivable numeric(12,2) DEFAULT 0 NOT NULL,
    running_fine numeric(12,2) DEFAULT 0 NOT NULL,
    rest_fine numeric(12,2) DEFAULT 0 NOT NULL,
    bill_service numeric(12,2) DEFAULT 0 NOT NULL,
    pay_deposit numeric(12,2) DEFAULT 0 NOT NULL,
    rest_receivable numeric(12,2) DEFAULT 0 NOT NULL,
    rest_base numeric(12,2) DEFAULT 0 NOT NULL,
    day_period smallint DEFAULT 0 NOT NULL,
    mortgage_to smallint DEFAULT 0 NOT NULL,
    day_count integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.receivables OWNER TO postgres;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tasks (
    order_id integer NOT NULL,
    descriptions character varying(128),
    period_from date NOT NULL,
    period_to date NOT NULL,
    recipient_name character varying(50) NOT NULL,
    recipient_position character varying(50) NOT NULL,
    giver_position character varying(50) NOT NULL,
    giver_name character varying(50) NOT NULL
);


ALTER TABLE public.tasks OWNER TO postgres;

--
-- Name: type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.type_id_seq OWNER TO postgres;

--
-- Name: types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.types (
    id integer DEFAULT nextval('public.type_id_seq'::regclass) NOT NULL,
    name character varying(50) NOT NULL,
    wheel_id smallint DEFAULT 0 NOT NULL,
    merk_id smallint DEFAULT 0 NOT NULL
);


ALTER TABLE public.types OWNER TO postgres;

--
-- Name: units; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.units (
    order_id integer NOT NULL,
    nopol character varying(15) NOT NULL,
    year integer DEFAULT 0 NOT NULL,
    frame_number character varying(25),
    machine_number character varying(25),
    bpkb_name character varying(50),
    color character varying(50),
    dealer character varying(50),
    surveyor character varying(50),
    type_id integer DEFAULT 0 NOT NULL,
    warehouse_id smallint DEFAULT 0 NOT NULL
);


ALTER TABLE public.units OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    email character varying(128) NOT NULL,
    password character varying(50) NOT NULL,
    role character varying(25) NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: warehouse_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.warehouse_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.warehouse_id_seq OWNER TO postgres;

--
-- Name: warehouses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.warehouses (
    id smallint DEFAULT nextval('public.warehouse_id_seq'::regclass) NOT NULL,
    name character varying(50) NOT NULL,
    descriptions character varying(128)
);


ALTER TABLE public.warehouses OWNER TO postgres;

--
-- Name: wheel_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wheel_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wheel_id_seq OWNER TO postgres;

--
-- Name: wheels; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wheels (
    id smallint DEFAULT nextval('public.wheel_id_seq'::regclass) NOT NULL,
    name character varying(10) NOT NULL,
    short_name character varying(5) NOT NULL
);


ALTER TABLE public.wheels OWNER TO postgres;

--
-- Name: merks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.merks ALTER COLUMN id SET DEFAULT nextval('public.merk_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.actions (id, action_at, code, pic, descriptions, order_id) FROM stdin;
\.


--
-- Data for Name: branchs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.branchs (id, name, street, city, phone, cell, zip, head_branch, email) FROM stdin;
1	Jatibarang	Jl. Pasar Sepur	Jatibarang	08596522323	012454787	45616	Mastur	mastur.st12@gmail.com
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (order_id, name, agreement_number, payment_type) FROM stdin;
3	Wardiman	dd	CO-1
\.


--
-- Data for Name: finances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.finances (id, name, short_name, street, city, phone, cell, zip, email) FROM stdin;
1	Bussan Auto Finance	BAF	Jl. Jend. Sudirman	Indramayu	2569874545	65979	2598987	busan.123@gmail.com
2	Auto Discret Finance	Adira	Jl. Jend. Sudirman	Indramayu	2569874545	65979	2598987	adira.finance@gmail.com
\.


--
-- Data for Name: home_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.home_addresses (order_id, street, region, city, phone, zip) FROM stdin;
\.


--
-- Data for Name: ktp_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ktp_addresses (order_id, street, region, city, phone, zip) FROM stdin;
\.


--
-- Data for Name: merks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.merks (id, name) FROM stdin;
2	Yamaha
12	Suzuki
13	Honda
1	Mitsubishi
14	Hyundai
\.


--
-- Data for Name: office_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.office_addresses (order_id, street, region, city, phone, zip) FROM stdin;
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel, ppn, user_name, verified_by, validated_by, finance_id, branch_id, nominal, subtotal, is_stnk, stnk_price) FROM stdin;
3	x/3336/2022/XII/5556	2022-03-03	2022-03-05	1800000.00	30.00	1260000.00	0.00	Opick	\N	\N	1	1	0.00	540000.00	t	0.00
\.


--
-- Data for Name: post_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.post_addresses (order_id, street, region, city, phone, zip) FROM stdin;
\.


--
-- Data for Name: receivables; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.receivables (order_id, covenant_at, due_at, mortgage_by_month, mortgage_receivable, running_fine, rest_fine, bill_service, pay_deposit, rest_receivable, rest_base, day_period, mortgage_to, day_count) FROM stdin;
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tasks (order_id, descriptions, period_from, period_to, recipient_name, recipient_position, giver_position, giver_name) FROM stdin;
3	qwewqe	2022-03-05	2022-03-07	Kepala Cabang PT SPRS	Field Collector	Branch Head	Abdul Rahman
\.


--
-- Data for Name: types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.types (id, name, wheel_id, merk_id) FROM stdin;
2	Vario 125	2	13
3	Brio 1000	3	13
1	Fino 125	2	2
\.


--
-- Data for Name: units; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.units (order_id, nopol, year, frame_number, machine_number, bpkb_name, color, dealer, surveyor, type_id, warehouse_id) FROM stdin;
3	E 5985 FGD	2022							2	2
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, role) FROM stdin;
\.


--
-- Data for Name: warehouses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.warehouses (id, name, descriptions) FROM stdin;
1	Gudang Pusat	Indramayu
2	Gudang Patrol	Patrol
\.


--
-- Data for Name: wheels; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wheels (id, name, short_name) FROM stdin;
2	Roda 2	R2
6	Roda 3	R3
3	Roda 4	R4
\.


--
-- Name: action_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.action_id_seq', 1, false);


--
-- Name: branch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.branch_id_seq', 2, true);


--
-- Name: finance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finance_id_seq', 2, true);


--
-- Name: merk_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.merk_id_seq', 14, true);


--
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_id_seq', 7, true);


--
-- Name: type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.type_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: warehouse_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.warehouse_id_seq', 2, true);


--
-- Name: wheel_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wheel_id_seq', 6, true);


--
-- Name: branchs cabangs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.branchs
    ADD CONSTRAINT cabangs_pkey PRIMARY KEY (id);


--
-- Name: customers customers_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_order_id_key UNIQUE (order_id);


--
-- Name: finances finances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.finances
    ADD CONSTRAINT finances_pkey PRIMARY KEY (id);


--
-- Name: warehouses gudangs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.warehouses
    ADD CONSTRAINT gudangs_pkey PRIMARY KEY (id);


--
-- Name: home_addresses home_addresses_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.home_addresses
    ADD CONSTRAINT home_addresses_order_id_key UNIQUE (order_id);


--
-- Name: ktp_addresses ktp_addresses_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ktp_addresses
    ADD CONSTRAINT ktp_addresses_order_id_key UNIQUE (order_id);


--
-- Name: merks merk_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.merks
    ADD CONSTRAINT merk_pkey PRIMARY KEY (id);


--
-- Name: office_addresses office_addresses_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.office_addresses
    ADD CONSTRAINT office_addresses_order_id_key UNIQUE (order_id);


--
-- Name: orders orders_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_name_key UNIQUE (name);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: post_addresses post_addresses_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_addresses
    ADD CONSTRAINT post_addresses_order_id_key UNIQUE (order_id);


--
-- Name: wheels roda_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wheels
    ADD CONSTRAINT roda_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_order_id_key UNIQUE (order_id);


--
-- Name: actions tindakans_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT tindakans_pkey PRIMARY KEY (id);


--
-- Name: receivables tunggakans_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receivables
    ADD CONSTRAINT tunggakans_order_id_key UNIQUE (order_id);


--
-- Name: types types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.types
    ADD CONSTRAINT types_pkey PRIMARY KEY (id);


--
-- Name: units units_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_order_id_key UNIQUE (order_id);


--
-- Name: units uq_unit_nopol; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT uq_unit_nopol UNIQUE (nopol);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_key UNIQUE (name);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idq_branch_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idq_branch_name ON public.branchs USING btree (name);


--
-- Name: idq_cabang_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idq_cabang_name ON public.types USING btree (name);


--
-- Name: idq_finance_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idq_finance_name ON public.finances USING btree (name);


--
-- Name: idq_gudang_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idq_gudang_name ON public.warehouses USING btree (name);


--
-- Name: idq_merk_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idq_merk_name ON public.merks USING btree (name);


--
-- Name: idq_roda_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idq_roda_name ON public.wheels USING btree (name);


--
-- Name: idq_type_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idq_type_name ON public.types USING btree (name);


--
-- Name: idx_type_merk; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_type_merk ON public.types USING btree (merk_id);


--
-- Name: idx_type_roda; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_type_roda ON public.types USING btree (wheel_id);


--
-- Name: actions actions_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT actions_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: customers customers_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: types fk_type_merk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.types
    ADD CONSTRAINT fk_type_merk FOREIGN KEY (merk_id) REFERENCES public.merks(id) ON DELETE RESTRICT;


--
-- Name: types fk_type_roda; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.types
    ADD CONSTRAINT fk_type_roda FOREIGN KEY (wheel_id) REFERENCES public.wheels(id) ON DELETE RESTRICT;


--
-- Name: home_addresses home_addresses_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.home_addresses
    ADD CONSTRAINT home_addresses_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: ktp_addresses ktp_addresses_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ktp_addresses
    ADD CONSTRAINT ktp_addresses_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: office_addresses office_addresses_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.office_addresses
    ADD CONSTRAINT office_addresses_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: orders orders_cabang_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_cabang_id_fkey FOREIGN KEY (branch_id) REFERENCES public.branchs(id);


--
-- Name: orders orders_finance_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_finance_id_fkey FOREIGN KEY (finance_id) REFERENCES public.finances(id);


--
-- Name: post_addresses post_addresses_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_addresses
    ADD CONSTRAINT post_addresses_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: receivables receivables_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receivables
    ADD CONSTRAINT receivables_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: tasks tasks_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: units units_gudang_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_gudang_id_fkey FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(id) ON DELETE RESTRICT;


--
-- Name: units units_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: units units_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.units
    ADD CONSTRAINT units_type_id_fkey FOREIGN KEY (type_id) REFERENCES public.types(id) ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

