--
-- PostgreSQL database dump
--

-- Dumped from database version 12.10 (Debian 12.10-1.pgdg100+1)
-- Dumped by pg_dump version 12.10 (Debian 12.10-1.pgdg100+1)

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

SET default_table_access_method = heap;

--
-- Name: acc_code; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.acc_code (
    id smallint NOT NULL,
    name character varying(50) NOT NULL,
    type_id smallint NOT NULL,
    descriptions character varying(256),
    token_name tsvector,
    is_active boolean DEFAULT true NOT NULL,
    is_auto_debet boolean DEFAULT true,
    receivable_option smallint DEFAULT 0
);


ALTER TABLE public.acc_code OWNER TO postgres;

--
-- Name: acc_group; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.acc_group (
    id smallint NOT NULL,
    name character varying(50) NOT NULL,
    descriptions character varying(256)
);


ALTER TABLE public.acc_group OWNER TO postgres;

--
-- Name: acc_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.acc_type (
    id smallint NOT NULL,
    name character varying(50) NOT NULL,
    descriptions character varying(256),
    group_id smallint DEFAULT 1 NOT NULL
);


ALTER TABLE public.acc_type OWNER TO postgres;

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

--
-- Name: actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.actions (
    id integer DEFAULT nextval('public.action_id_seq'::regclass) NOT NULL,
    action_at date NOT NULL,
    pic character varying(50) NOT NULL,
    descriptions text,
    order_id integer NOT NULL,
    file_name character varying(128)
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
    agreement_number character varying(50),
    payment_type character varying(25) NOT NULL
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- Name: finance_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.finance_groups (
    id smallint NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.finance_groups OWNER TO postgres;

--
-- Name: finance_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.finance_groups_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.finance_groups_id_seq OWNER TO postgres;

--
-- Name: finance_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.finance_groups_id_seq OWNED BY public.finance_groups.id;


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
    email character varying(50),
    group_id smallint DEFAULT 0 NOT NULL
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
-- Name: invoice_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.invoice_details (
    invoice_id integer NOT NULL,
    order_id integer NOT NULL
);


ALTER TABLE public.invoice_details OWNER TO postgres;

--
-- Name: invoices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.invoices (
    id integer NOT NULL,
    invoice_at date NOT NULL,
    payment_term smallint DEFAULT 0 NOT NULL,
    due_at date NOT NULL,
    salesman character varying(50) NOT NULL,
    finance_id smallint NOT NULL,
    account_id smallint DEFAULT 0 NOT NULL,
    subtotal numeric(12,2) DEFAULT 0 NOT NULL,
    ppn numeric(8,2) DEFAULT 0 NOT NULL,
    tax numeric(12,2) DEFAULT 0 NOT NULL,
    total numeric(12,2) DEFAULT 0 NOT NULL,
    memo character varying(256),
    token tsvector
);


ALTER TABLE public.invoices OWNER TO postgres;

--
-- Name: invoices_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.invoices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.invoices_id_seq OWNER TO postgres;

--
-- Name: invoices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.invoices_id_seq OWNED BY public.invoices.id;


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
-- Name: lents_id_sequence; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lents_id_sequence
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lents_id_sequence OWNER TO postgres;

--
-- Name: lents; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lents (
    order_id integer NOT NULL,
    name character varying(50) NOT NULL,
    street character varying(128),
    city character varying(50),
    phone character varying(25),
    cell character varying(25),
    zip character varying(6),
    serial_num integer DEFAULT nextval('public.lents_id_sequence'::regclass) NOT NULL
);


ALTER TABLE public.lents OWNER TO postgres;

--
-- Name: loans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.loans (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    street character varying(128),
    city character varying(50),
    phone character varying(25),
    cell character varying(25),
    zip character varying(6),
    persen numeric(8,2),
    serial_num integer DEFAULT nextval('public.lents_id_sequence'::regclass) NOT NULL
);


ALTER TABLE public.loans OWNER TO postgres;

--
-- Name: loans_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.loans_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.loans_id_seq OWNER TO postgres;

--
-- Name: loans_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.loans_id_seq OWNED BY public.loans.id;


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
-- Name: order_name_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_name_seq
    AS integer
    START WITH 100
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_name_seq OWNER TO postgres;

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
    user_name character varying(50) NOT NULL,
    verified_by character varying(50),
    finance_id smallint DEFAULT 0 NOT NULL,
    branch_id smallint DEFAULT 0 NOT NULL,
    is_stnk boolean DEFAULT true NOT NULL,
    stnk_price numeric(12,2) DEFAULT 0 NOT NULL,
    matrix numeric(12,2) DEFAULT 0 NOT NULL,
    token tsvector
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
-- Name: trx_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.trx_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.trx_seq OWNER TO postgres;

--
-- Name: trx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trx (
    id integer DEFAULT nextval('public.trx_seq'::regclass) NOT NULL,
    ref_id integer DEFAULT 0 NOT NULL,
    division character varying(25) DEFAULT 'trx-umum'::character varying NOT NULL,
    descriptions character varying(128) NOT NULL,
    trx_date date DEFAULT now() NOT NULL,
    memo character varying(256),
    trx_token tsvector
);


ALTER TABLE public.trx OWNER TO postgres;

--
-- Name: trx_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trx_detail (
    id integer NOT NULL,
    code_id smallint NOT NULL,
    trx_id integer NOT NULL,
    debt numeric(12,2),
    cred numeric(12,2)
);


ALTER TABLE public.trx_detail OWNER TO postgres;

--
-- Name: trx_detail_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.trx_detail_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.trx_detail_seq OWNER TO postgres;

--
-- Name: trx_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trx_type (
    id smallint NOT NULL,
    name character varying(50) NOT NULL,
    descriptions character varying(256)
);


ALTER TABLE public.trx_type OWNER TO postgres;

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
    color character varying(50),
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
-- Name: finance_groups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.finance_groups ALTER COLUMN id SET DEFAULT nextval('public.finance_groups_id_seq'::regclass);


--
-- Name: invoices id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices ALTER COLUMN id SET DEFAULT nextval('public.invoices_id_seq'::regclass);


--
-- Name: loans id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.loans ALTER COLUMN id SET DEFAULT nextval('public.loans_id_seq'::regclass);


--
-- Name: merks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.merks ALTER COLUMN id SET DEFAULT nextval('public.merk_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: acc_code; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.acc_code (id, name, type_id, descriptions, token_name, is_active, is_auto_debet, receivable_option) FROM stdin;
5112	Biaya Listrik	51	Biaya pemakaian listrik	'biaya':1,3 'listrik':2,5 'pakai':4	t	t	3
3211	Prive Pak Kris	32	Pengambilan modal, pinjam kas oleh pak Kris	'ambil':4 'kas':7 'kris':3,10 'modal':5 'oleh':8 'pak':2,9 'pinjam':6 'prive':1	t	t	3
4111	Pendapatan Invoice	41	Penarikan dana dari pihak Finance karena adanya ...	'ada':9 'arik':3 'dana':4 'dapat':1 'dari':5 'finance':7 'invoice':2 'karena':8 'pihak':6	t	f	2
5111	Biaya Transport	51	Biaya transportasi karyawan	'biaya':1,3 'karyaw':5 'transport':2 'transportasi':4	t	t	3
5113	Biaya Telephone dan Fax	51	Biaya telephone dan faximile ke telkomsel	'biaya':1,5 'dan':3,7 'fax':4 'faximile':8 'ke':9 'telephone':2,6 'telkomsel':10	t	f	3
5115	Biaya Pos dan Materai	51	Biaya pengiriman surat dan pembelian materai.	'beli':9 'biaya':1,5 'dan':3,8 'irim':6 'matera':4,10 'pos':2 'surat':7	t	f	3
5116	Biaya ATK	51	Biaya alat tulis kantor termasuk termasuk peralatan seperti komputer, meja, kursi, lemari	'alat':4,9 'atk':2 'biaya':1,3 'kantor':6 'komputer':11 'kursi':13 'lemar':14 'masuk':7,8 'meja':12 'sepert':10 'tulis':5	t	t	3
5411	Upah Tenaga Kerja	54	Biaya overhead perusahaan yg dikeluarkan untuk memayar upah karena mengerjakan sesuatu	'biaya':4 'erja':13 'karena':12 'keluar':8 'kerja':3 'overhead':5 'payar':10 'sesuatu':14 'tenaga':2 'untuk':9 'upah':1,11 'usaha':6 'yg':7	t	t	3
5119	Biaya Lain-lain	51	Biaya yg terdiri dari bermacam transaksi serta tidak tercantum pada salah satu perkiraan yang terdapat dalam transaksi perushaan	'biaya':1,5 'cantum':13 'dalam':20 'dapat':19 'dari':8 'diri':7 'kira':17 'lain':3,4 'lain-lain':2 'macam':9 'pada':14 'salah':15 'satu':16 'serta':11 'tidak':12 'transaksi':10,21 'usha':22 'yang':18 'yg':6	t	t	3
5117	Biaya Servis	51	Biaya service kendaraan, AC, komputer, dll.	'ac':6 'biaya':1,3 'dll':8 'komputer':7 'ndara':5 'service':4 'servis':2	t	t	3
5114	Biaya Internet	51	Biaya jaringan internet ke Biznet	'biaya':1,3 'biznet':7 'internet':2,5 'jaring':4 'ke':6	t	t	3
1112	Bank BCA 3039203040	11	An Sarana Padma Ridho Sepuh	'3039203040':3 'an':4 'bank':1 'bca':2 'padma':6 'ridho':7 'sarana':5 'sepuh':8	t	f	1
1113	BANK MANDIRI 1340000006105	11	An Sarana Padma Ridho Sepuh	'1340000006105':3 'an':4 'bank':1 'mandir':2 'padma':6 'ridho':7 'sarana':5 'sepuh':8	t	f	1
2211	BNI	22	Utang usaha di bank BNI	'bank':5 'bni':1,6 'di':4 'usaha':3 'utang':2	t	t	2
2212	SAMSAT	22	Utang usaha di Samsat	'di':4 'samsat':1,5 'usaha':3 'utang':2	t	t	2
5118	Biaya Konsumsi	51	Biaya yg dikeluarkan karena suatu kegiatan yg dpt mengurangi atau menghabiskan barang dan jasa	'atau':12 'barang':14 'biaya':1,3 'dan':15 'dpt':10 'giat':8 'habis':13 'jasa':16 'karena':6 'keluar':5 'konsumsi':2 'suatu':7 'urang':11 'yg':4,9	t	t	3
5511	Piutang Jasa	55	Piutang diberikan kepada pihak Finance yg timbul karena perusahaan menerima job penarikan kendaraan sejumlah BT Matel\n	'arik':14 'beri':4 'bt':17 'erima':12 'finance':7 'jasa':2 'job':13 'karena':10 'matel':18 'ndara':15 'pada':5 'pihak':6 'piutang':1,3 'sejum':16 'timbul':9 'usaha':11 'yg':8	t	f	3
2311	Hutang Pajak	23	Pajak yg belum dibayar karena menunggu pembayaran dari tarikan	'bayar':6,9 'belum':5 'dari':10 'hutang':1 'karena':7 'pajak':2,3 'tari':11 'unggu':8 'yg':4	t	f	2
4211	BAYAR PIUTANG	42	\N	'bayar':1 'piutang':2	t	t	2
5611	SELI	56	\N	'seli':1	t	t	3
5513	Piutang Unit	55	\N	'piutang':1 'unit':2	t	f	3
3111	Modal pak Kris	31	Modal yg masuk dari pak Kris	'dari':7 'kris':3,9 'masuk':6 'modal':1,4 'pak':2,8 'yg':5	t	t	2
6011	Pembayaran Pajak	60	Pajak Pertambahan Nilai	'bayar':1 'nila':5 'pajak':2,3 'tambah':4	t	f	3
5512	Pinjaman	55	\N	'pinjam':1	t	f	3
5311	Biaya Gaji karyawan Tetap	53	Pencatatan data kompensasi karyawan seperti uang potongan dari setiap gaji dan pajak serta tunjangan karyawan tetap	'biaya':1 'catat':5 'dan':15 'dari':12 'data':6 'gaji':2,14 'karyaw':3,8,19 'kompensasi':7 'pajak':16 'potong':11 'sepert':9 'serta':17 'setiap':13 'tetap':4,20 'tunjang':18 'uang':10	t	t	3
5312	Biaya Gaji Karyawan Honorer	51	Pencatatan data kompensasi karyawan seperti uang potongan dari setiap gaji\ndan pajak serta tunjangan bukan karyawan tetap 	'biaya':1 'bukan':19 'catat':5 'dan':15 'dari':12 'data':6 'gaji':2,14 'honorer':4 'karyaw':3,8,20 'kompensasi':7 'pajak':16 'potong':11 'sepert':9 'serta':17 'setiap':13 'tetap':21 'tunjang':18 'uang':10	t	f	3
1111	Kas Kantor	11	Kas bendahara Kantor	'bendahara':4 'kantor':2,5 'kas':1,3	t	f	1
4113	Cicilan Unit	41	\N	'cicil':1 'unit':2	t	f	2
5211	Kasbon Cabang JTB	52	Biaya yg dikeluarkan untuk penarikan kendaraan yg tidak ada STNK	'ada':12 'arik':8 'biaya':4 'cabang':2 'jtb':3 'kasbon':1 'keluar':6 'ndara':9 'stnk':13 'tidak':11 'untuk':7 'yg':5,10	t	t	3
4112	Angsuran Piutang	41	\N	'angsur':1 'piutang':2	t	f	2
\.


--
-- Data for Name: acc_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.acc_group (id, name, descriptions) FROM stdin;
1	Harta (Aktiva)	Segala sesuatu yang berhubungan dengan asset perusahaan (+)
2	Utang (Kewajiban)	Segala sesuatu yang menjadi kewajiban perusahaan yang harus dibayarkan kepada pihak luar dalam periode tertentu (-).
3	Modal	Kekayaan perusahaan yang menjadi bagian dari pemilik perusahaan (-).
4	Pendapatan	Segala sesuatu yang diterima oleh perusahaan, baik yang didapat dari hasil operasional perusahaan (misalnya, bengkel mendapat pendapatan jasa servis kendaraan) dan kegiatan di luar operasional perusahaan (misalnya, bunga bank) (-)
5	Beban	Biaya-biaya yang dikeluarkan perusahaan dalam kegiatan operasionalnya untuk mendapatkan penghasilan. Contoh: beban air, listrik, dan telepon (+).
\.


--
-- Data for Name: acc_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.acc_type (id, name, descriptions, group_id) FROM stdin;
14	Peralatan	Kelompok akun yg digunakan untuk mencatat barang atau tempat yang digunakan perusahaan untuk mendukung jalannya pekerjaan.	1
11	Kas	Kelompok akun yg berfungsi mencatat perubahan uang seperti penerimaan atau pengeluaran. termasuk akun kas, seperti cek, giro.	1
15	Perlengkapan	Barang2 yg bisa dipakai berulang-ulang dan habis, bentuknya relatif kecil umunya untuk melengkapi kebutuhan perusahaan	1
21	Hutang Usaha	Pinjaman uang dari pihak yang dilakukan seseorang kepada perusahaan, tidak hanya uang bisa juga barang atau jasa	2
22	Hutang Bank	Pinjaman yang diberikan oleh bank kepada perusahaan yg harus dibayar oleh perusahaan selama periode berikut bunganya	2
23	Hutang Lainnya	Semua utang yg diklasifikasikan sbg utang yg tidak lancar.	2
31	Modal Usaha	Segala sesuatu yang dipergunakan untuk membangun atau memulai sebuah usaha, tergantung pada jenis usaha yang dijalankan	3
32	Prive	Investor yang menarik kembali modal atau aset mereka dari suatu perusahaan dgn maksud untuk digunakan sendiri oleh pemiliknya	3
41	Pendapatan Jasa	Arus masuk bruto dari kegiatan usaha yg mengakibatkan kenaikan ekuitas yang tidak berasal dari kontribusi modal	4
42	Pendapatan Lainnya	Arus masuk bruto dari kegiatan di luar usaha yg mengakibatkan kenaikan ekuitas yang tidak berasal dari kontribusi modal	4
51	Biaya Kantor	Kelompok akun yg diakibatkan adanya pembayaran listrik, internet, PDAM, telephone.	5
52	Biaya Operasional	Biaya yg dikeluarkan untuk aktivitas harian perusahaan seperti komisi, transportasi, sewa, perbaikan, pajak	5
53	Biaya Gaji	Biaya yg dikeluarkan untuk gaji atau tunjangan	5
54	Biaya Tenaga Kerja	Biaya yg dikeluarkan untuk membayar upah tenaga kerja	5
60	Pajak	Kontribusi wajib kepada negara yang terutang oleh orang pribadi atau badan yang bersifat memaksa berdasarkan Undang-undang	5
56	KASBON	\N	5
55	Piutang	Kelompok akun yg mencatat semua pengeluaran dalam bentuk piutang kepada pihak lain.	5
\.


--
-- Data for Name: actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.actions (id, action_at, pic, descriptions, order_id, file_name) FROM stdin;
10	2022-03-25	Test JPG / JPEG / PNG	Test upload and download image	10	d9f143ce.jpg
11	2022-03-25	Test PDF	Test upload and download pdf	10	d855835b.pdf
\.


--
-- Data for Name: branchs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.branchs (id, name, street, city, phone, cell, zip, head_branch, email) FROM stdin;
3	Pusat Indramayu	\N	\N	\N	\N	\N	Deddy Pranoto	\N
1	Jatibarang	Jl. Pasar Sepur	Jatibarang	08596522323	012454787	45616	Syaenudin	mastur.st12@gmail.com
4	Karawang	\N	\N	\N	\N	\N	Gugur Junaedi	\N
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (order_id, name, agreement_number, payment_type) FROM stdin;
\.


--
-- Data for Name: finance_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.finance_groups (id, name) FROM stdin;
1	BAF
2	CLIPAN
3	MTF
4	WOM
5	FIF
6	OTTO
7	ADIRA
\.


--
-- Data for Name: finances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.finances (id, name, short_name, street, city, phone, cell, zip, email, group_id) FROM stdin;
1	Bussan Auto Finance	BAF	Jl. Jend. Sudirman	Indramayu	2569874545	65979	2598987	busan.123@gmail.com	1
6	COLLECTIUS	COL	\N	\N	\N	\N	\N	\N	1
7	Mandiri Utama Finance	MUF	\N	\N	\N	\N	\N	\N	1
9	Mitra Pinasthika Mustika Finance	MPMF	\N	\N	\N	\N	\N	\N	1
10	Top Finance Company	TFC	\N	\N	\N	\N	\N	\N	1
11	Kredit Plus	KP+	\N	\N	\N	\N	\N	\N	1
13	MEGAPARA	MPR	\N	\N	\N	\N	\N	\N	1
17	BFI Finance	BFI	\N	\N	\N	\N	\N	\N	1
20	Radana Finance	RAD	\N	\N	\N	\N	\N	\N	1
21	Mega Auto Central Finance	MACF	\N	\N	\N	\N	\N	\N	1
19	CLIPAN	CLIP	\N	\N	\N	\N	\N	\N	2
14	Clipan Karawang\n	CLIP K	\N	\N	\N	\N	\N	\N	2
15	Clipan Palembang	CLIP P	\N	\N	\N	\N	\N	\N	2
3	Mandiri Tunas Finance	MTF	\N	Cirebon	\N	\N	\N	\N	3
18	Mandiri Tunas Finance Semarang	MTF S	\N	\N	\N	\N	\N	\N	3
4	Clipan Bekasi	CLIP B	\N	\N	\N	\N	\N	\N	2
12	WOM Finance	WOMF	\N	\N	\N	\N	\N	\N	4
8	FIF Group	FIF	\N	\N	\N	\N	\N	\N	5
5	OTO Kredit Motor	OTTO	\N	\N	\N	\N	\N	\N	6
2	Auto Discret Finance	Adira	Jl. Jend. Sudirman	Indramayu	2569874545	65979	2598987	adira.finance@gmail.com	7
16	SUZUKI FINANCE INDONESIA	SFI 	\N	\N	\N	\N	\N	\N	1
\.


--
-- Data for Name: home_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.home_addresses (order_id, street, region, city, phone, zip) FROM stdin;
\.


--
-- Data for Name: invoice_details; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.invoice_details (invoice_id, order_id) FROM stdin;
\.


--
-- Data for Name: invoices; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.invoices (id, invoice_at, payment_term, due_at, salesman, finance_id, account_id, subtotal, ppn, tax, total, memo, token) FROM stdin;
\.


--
-- Data for Name: ktp_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ktp_addresses (order_id, street, region, city, phone, zip) FROM stdin;
\.


--
-- Data for Name: lents; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lents (order_id, name, street, city, phone, cell, zip, serial_num) FROM stdin;
\.


--
-- Data for Name: loans; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.loans (id, name, street, city, phone, cell, zip, persen, serial_num) FROM stdin;
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
15	Toyota
16	Daihatsu
\.


--
-- Data for Name: office_addresses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.office_addresses (order_id, street, region, city, phone, zip) FROM stdin;
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, name, order_at, printed_at, bt_finance, bt_percent, bt_matel, user_name, verified_by, finance_id, branch_id, is_stnk, stnk_price, matrix, token) FROM stdin;
18	000000018	2022-01-18	2022-01-18	1000000.00	20.00	800000.00	Mastur	test	5	1	t	0.00	1000000.00	'-2022':4 '-6716':17 '000000018':1 '18':2 '2':27 '2015':24 'ada':15 'cabang':10 'e':16 'finance':5 'gudang':21,22 'ix':18 'jan':3 'jatibarang':11 'kredit':7 'm3':20 'mio':19 'motor':8 'oto':6 'otto':9 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
57	000000057	2022-01-22	2022-01-22	8000000.00	0.00	8000000.00	Mastur	test	2	3	t	0.00	8000000.00	'-1256':19 '-2022':4 '000000057':1 '2019':26 '22':2 '4':29 'ada':17 'adira':9 'auto':6 'cabang':10 'daihatsu':27 'deddy':13 'discret':7 'e':18 'finance':5,8 'grand':21 'gudang':23 'indramayu':12 'jan':3 'max':22 'pranoto':14 'pusat':11,24 'qd':20 'r4':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25
77	000000103	2022-03-19	2022-03-26	1500000.00	20.00	1200000.00	Opick	test	21	1	t	0.00	1500000.00	'-2022':4 '-5826':18 '000000103':1 '125':21 '19':2 '2':28 '2019':25 'ada':16 'auto':7 'cabang':11 'central':8 'cq':19 'e':17 'finance':5,9 'gudang':22 'honda':26 'jatibarang':12,23 'macf':10 'mar':3 'mega':6 'r2':29 'roda':27 'stnk':15 'stnk-ada':14 'syaenudin':13 'tahun':24 'vario':20
15	000000015	2022-01-06	2022-01-06	900000.00	20.00	900000.00	Mastur	test	5	1	t	0.00	900000.00	'-2022':4 '-4146':17 '000000015':1 '06':2 '2':27 '2012':24 'ada':15 'cabang':10 'finance':5 'gudang':21,22 'jan':3 'jatibarang':11 'jupiter':19 'ko':18 'kredit':7 'motor':8 'mx':20 'oto':6 'otto':9 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':23 'yamaha':25
21	000000021	2022-02-21	2022-02-21	950000.00	20.00	760000.00	Mastur	test	1	1	t	0.00	950000.00	'-2022':4 '-6262':17 '000000021':1 '2':27 '2015':24 '21':2 'ada':15 'auto':7 'b':16 'baf':9 'bussan':6 'cabang':10 'feb':3 'finance':5,8 'gudang':21,22 'jatibarang':11 'm3':20 'mio':19 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'vky':18 'yamaha':25
24	000000024	2022-02-25	2022-02-25	1500000.00	20.00	1200000.00	Mastur	test	1	1	t	0.00	1500000.00	'-2022':4 '-2146':17 '000000024':1 '2':27 '2018':24 '25':2 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'feb':3 'finance':5,8 'gudang':21,22 'jatibarang':11 'mio':19 'qaf':18 'r2':28 'roda':26 's':20 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
26	000000026	2022-03-07	2022-03-07	850000.00	20.00	680000.00	Mastur	test	11	1	t	0.00	850000.00	'-2022':4 '-2891':16 '000000026':1 '07':2 '150':19 '2':26 '2014':23 'ada':14 'cabang':9 'finance':5 'gudang':20 'honda':24 'jatibarang':10,21 'kp':8 'kredit':6 'mar':3 'plus':7 'r2':27 'roda':25 'stnk':13 'stnk-ada':12 'syaenudin':11 't':15 'tahun':22 'vario':18 'wp':17
7	000000007	2022-03-15	2022-03-15	1700000.00	20.00	1360000.00	Mastur	test	1	3	t	0.00	1700000.00	'-2022':4 '-2033':19 '000000007':1 '125':22 '15':2 '2':29 '2019':26 'ada':17 'auto':7 'baf':9 'bussan':6 'cabang':10 'deddy':13 'e':18 'finance':5,8 'fino':21 'gudang':23 'indramayu':12 'mar':3 'pbj':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'yamaha':27
33	000000033	2022-03-16	2022-03-16	1500000.00	20.00	1200000.00	Mastur	test	2	1	t	0.00	1500000.00	'-2022':4 '-5856':17 '000000033':1 '16':2 '2':26 '2019':23 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'genio':19 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':22 'zt':18
97	000000135	2022-03-21	2022-03-28	17000000.00	10.00	15300000.00	Opick	test	19	3	t	0.00	17000000.00	'-1942':17 '-2022':4 '000000135':1 '2005':23 '21':2 '4':26 'ada':15 'b':16 'cabang':8 'clip':7 'clipan':6 'deddy':11 'evk':18 'finance':5 'gudang':20 'honda':24 'indramayu':10 'jazz':19 'mar':3 'pranoto':12 'pusat':9,21 'r4':27 'roda':25 'stnk':14 'stnk-ada':13 'tahun':22
78	000000104	2022-03-19	2022-03-26	1200000.00	20.00	960000.00	Opick	test	17	1	t	0.00	1200000.00	'-2022':4 '-6505':16 '000000104':1 '19':2 '2':25 '2019':22 'ada':14 'beat':18 'bfi':6,8 'cabang':9 'e':15 'finance':5,7 'gudang':19 'honda':23 'jatibarang':10,20 'mar':3 'qr':17 'r2':26 'roda':24 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':21
89	000000114	2022-03-25	2022-03-26	1300000.00	20.00	1040000.00	Opick	test	12	1	t	0.00	1300000.00	'-2022':4 '-3561':16 '000000114':1 '2':25 '2019':22 '25':2 'ada':14 'beat':18 'cabang':9 'e':15 'finance':5,7 'gudang':19 'honda':23 'jatibarang':10,20 'mar':3 'pbi':17 'r2':26 'roda':24 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':21 'wom':6 'womf':8
11	000000011	2021-09-27	2021-09-27	900000.00	20.00	720000.00	Mastur	test	8	1	t	0.00	900000.00	'-2021':4 '-4892':16 '000000011':1 '2':25 '2012':22 '27':2 'ada':14 'cabang':9 'e':15 'fif':6,8 'finance':5 'group':7 'gudang':19,20 'jatibarang':10 'jupiter':18 'r2':26 'roda':24 'sep':3 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':21 'tk':17 'yamaha':23
4	000000004	2022-02-18	2022-02-18	1200000.00	20.00	960000.00	Mastur	test	2	3	t	0.00	1200000.00	'-2022':4 '-5080':19 '000000004':1 '18':2 '2':28 '2015':25 'ada':17 'adira':9 'auto':6 'br':18 'cabang':10 'deddy':13 'discret':7 'feb':3 'finance':5,8 'gudang':22 'indramayu':12 'pranoto':14 'pusat':11,23 'py':20 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'vixion':21 'yamaha':26
6	000000006	2022-03-02	2022-03-02	1200000.00	20.00	960000.00	Mastur	test	2	3	t	0.00	1200000.00	'-2022':4 '-2633':19 '000000006':1 '02':2 '2':28 '2016':25 'ada':17 'adira':9 'auto':6 'beat':21 'cabang':10 'deddy':13 'discret':7 'e':18 'finance':5,8 'gudang':22 'honda':26 'indramayu':12 'mar':3 'pac':20 'pranoto':14 'pusat':11,23 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24
74	000000074	2022-03-07	2022-03-07	1300000.00	58.46	540000.00	Mastur	test	2	1	t	0.00	1300000.00	'-2022':4 '-6871':17 '000000074':1 '07':2 '2':26 '2018':23 'ada':15 'adira':9 'auto':6 'cabang':10 'cm':18 'discret':7 'e':16 'finance':5,8 'gudang':20 'jatibarang':11,21 'mar':3 'mio':19 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'yamaha':24
10	000000010	2022-03-19	2022-03-19	850000.00	20.00	680000.00	Mastur	test	2	3	t	0.00	850000.00	'-2022':4 '-5474':19 '000000010':1 '19':2 '2':28 '2013':25 'ada':17 'adira':9 'auto':6 'beat':21 'cabang':10 'deddy':13 'discret':7 'e':18 'finance':5,8 'gudang':22 'honda':26 'indramayu':12 'mar':3 'pranoto':14 'pusat':11,23 'q':20 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24
99	000000139	2022-03-28	2022-03-28	16000000.00	18.75	13000000.00	Opick	test	16	3	t	0.00	16000000.00	'-2022':4 '-8066':19 '000000139':1 '2020':25 '28':2 '4':28 'ada':17 'cabang':10 'carry':21 'deddy':13 'eg':20 'finance':5,7 'gudang':22 'indonesia':8 'indramayu':12 'mar':3 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'sfi':9 'stnk':16 'stnk-ada':15 'suzuk':6,26 't':18 'tahun':24
80	000000106	2022-03-19	2022-03-26	1200000.00	20.00	960000.00	Opick	test	2	1	t	0.00	1200000.00	'-2022':4 '-5737':17 '000000106':1 '125':20 '19':2 '2':27 '2019':24 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'fino':19 'gudang':21 'jatibarang':11,22 'mar':3 'pbo':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
84	000000110	2022-03-22	2022-03-26	1050000.00	20.00	840000.00	Opick	test	2	1	t	0.00	1050000.00	'-2022':4 '-5462':17 '000000110':1 '2':26 '2014':23 '22':2 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'qm':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
16	000000016	2022-01-14	2022-01-14	1000000.00	20.00	800000.00	Mastur	test	5	1	t	0.00	1000000.00	'-2022':4 '-3848':17 '000000016':1 '14':2 '2':26 '2016':23 'ada':15 'beat':19 'cabang':10 'e':16 'finance':5 'gudang':20,21 'honda':24 'jan':3 'jatibarang':11 'kredit':7 'motor':8 'oto':6 'otto':9 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'ub':18
17	000000017	2022-01-14	2022-01-14	750000.00	20.00	600000.00	Mastur	test	10	1	t	0.00	750000.00	'-125':21 '-2022':4 '-3828':17 '000000017':1 '14':2 '2':28 '2008':25 'ada':15 'cabang':10 'company':8 'finance':5,7 'fw':18 'gudang':22 'honda':26 'jan':3 'jatibarang':11 'pusat':23 'r2':29 'roda':27 'stnk':14 'stnk-ada':13 'supra':19 'syaenudin':12 't':16 'tahun':24 'tfc':9 'top':6 'x':20
100	000000143	2022-03-28	2022-03-28	35000000.00	28.57	25000000.00	Opick	test	3	3	t	0.00	35000000.00	'-2022':4 '-9844':19 '000000143':1 '2016':25 '28':2 '4':28 'ada':17 'box':21 'cabang':10 'deddy':13 'e':18 'finance':5,8 'gudang':22 'hb':20 'indramayu':12 'mandir':6 'mar':3 'mitsubish':26 'mtf':9 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'tunas':7
58	000000058	2022-01-31	2022-01-31	21400000.00	6.54	20000000.00	Mastur	test	4	3	t	0.00	21400000.00	'-2022':4 '-2281':19 '000000058':1 '2012':25 '31':2 '4':28 'ada':17 'b':9,18 'bekasi':7 'cabang':10 'clip':8 'clipan':6 'deddy':13 'finance':5 'gudang':22 'indramayu':12 'ios':21 'jan':3 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'sbt':20 'stnk':16 'stnk-ada':15 'tahun':24 'toyota':26
28	000000028	2022-03-09	2022-03-09	1300000.00	20.00	1040000.00	Mastur	test	2	1	t	0.00	1300000.00	'-2022':4 '-4487':17 '000000028':1 '09':2 '2':27 '2017':24 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'gudang':21 'jatibarang':11,22 'mar':3 'mio':19 'pj':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':23 'yamaha':25 'z':20
30	000000030	2022-03-12	2022-03-12	1450000.00	20.00	1160000.00	Mastur	test	2	1	t	0.00	1450000.00	'-2022':4 '-3615':17 '000000030':1 '12':2 '2':26 '2018':23 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':22 'verza':19 'zd':18
32	000000032	2022-03-14	2022-03-14	1200000.00	20.00	960000.00	Mastur	test	2	1	t	0.00	1200000.00	'-2022':4 '-2191':17 '000000032':1 '14':2 '2':26 '2017':23 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':22 'ys':18
81	000000107	2022-03-21	2022-03-26	800000.00	20.00	640000.00	Opick	test	12	1	t	0.00	800000.00	'-2022':4 '-2764':16 '000000107':1 '2':25 '2014':22 '21':2 'ada':14 'cabang':9 'cb150':18 'e':15 'finance':5,7 'gudang':19 'honda':23 'jatibarang':10,20 'mar':3 'r2':26 'roda':24 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':21 'wom':6 'womf':8 'xl':17
83	000000109	2022-03-22	2022-03-26	1600000.00	20.00	1280000.00	Opick	test	2	1	t	0.00	1600000.00	'-2022':4 '-3802':17 '000000109':1 '2':26 '2019':23 '22':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'r2':27 'roda':25 'scoopy':19 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':22 'zr':18
49	000000049	2021-11-22	2021-11-22	10000000.00	18.00	8200000.00	Mastur	test	17	3	t	0.00	10000000.00	'-2021':4 '-8630':18 '000000049':1 '2012':24 '22':2 '4':27 'ada':16 'bfi':6,8 'cabang':9 'deddy':12 'finance':5,7 'gudang':21 'h':17 'honda':25 'indramayu':11 'jazz':20 'nop':3 'pp':19 'pranoto':13 'pusat':10,22 'r4':28 'roda':26 'stnk':15 'stnk-ada':14 'tahun':23
66	000000066	2021-12-06	2021-12-06	1300000.00	20.00	1040000.00	Mastur	test	12	1	t	0.00	1300000.00	'-2021':4 '-6934':16 '000000066':1 '06':2 '2':25 '2017':22 'ada':14 'beat':18 'cabang':9 'des':3 'finance':5,7 'gudang':19 'honda':23 'jatibarang':10,20 'r2':26 'roda':24 'stnk':13 'stnk-ada':12 'syaenudin':11 't':15 'tahun':21 'wom':6 'womf':8 'yq':17
50	000000050	2021-12-16	2021-12-16	34500000.00	9.10	31360000.00	Mastur	test	18	3	t	0.00	34500000.00	'-2021':4 '-9442':21 '000000050':1 '1000':24 '16':2 '2017':28 '4':31 'ada':19 'brio':23 'cabang':12 'deddy':15 'des':3 'finance':5,8 'gudang':25 'h':20 'honda':29 'indramayu':14 'mandir':6 'mtf':10 'ng':22 'pranoto':16 'pusat':13,26 'r4':32 'roda':30 's':11 'semarang':9 'stnk':18 'stnk-ada':17 'tahun':27 'tunas':7
53	000000053	2021-12-24	2021-12-24	21450000.00	9.09	19500000.00	Mastur	test	14	3	t	0.00	21450000.00	'-1312':19 '-2021':4 '000000053':1 '2007':25 '24':2 '4':28 'ada':17 'cabang':10 'clip':8 'clipan':6 'd':18 'daihatsu':26 'deddy':13 'des':3 'finance':5 'gudang':22 'indramayu':12 'k':9 'karawang':7 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'wf':20 'xenia':21
40	000000040	2022-01-17	2022-01-17	26000000.00	7.69	24000000.00	Mastur	test	14	3	t	0.00	26000000.00	'-1164':19 '-2022':4 '000000040':1 '17':2 '2017':25 '4':28 'ada':17 'cabang':10 'clip':8 'clipan':6 'deddy':13 'ertiga':21 'finance':5 'fq':20 'gudang':22 'indramayu':12 'jan':3 'k':9 'karawang':7 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'suzuk':26 't':18 'tahun':24
41	000000041	2022-01-25	2022-01-25	26620000.00	9.09	24200000.00	Mastur	test	14	3	t	0.00	26620000.00	'-1788':19 '-2022':4 '000000041':1 '2017':25 '25':2 '4':28 'ada':17 'bc':20 'cabang':10 'clip':8 'clipan':6 'deddy':13 'finance':5 'gudang':22 'honda':26 'indramayu':12 'jan':3 'k':9 'karawang':7 'mobilio':21 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 't':18 'tahun':24
101	000000146	2022-03-26	2022-03-28	1000000.00	20.00	800000.00	Opick	test	2	1	t	0.00	1000000.00	'-2022':4 '-6987':17 '000000146':1 '2':27 '2013':24 '26':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'fu':20 'gudang':21 'jatibarang':11,22 'mar':3 'r2':28 'roda':26 'satria':19 'sd':18 'stnk':14 'stnk-ada':13 'suzuk':25 'syaenudin':12 'tahun':23
102	000000147	2022-03-28	2022-03-28	1600000.00	20.00	1280000.00	Opick	test	1	1	t	0.00	1600000.00	'-2022':4 '-6645':17 '000000147':1 '125':20 '2':27 '2018':24 '28':2 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'finance':5,8 'fino':19 'gudang':21 'jatibarang':11,22 'mar':3 'pav':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
62	000000062	2022-01-19	2022-01-19	3600000.00	20.00	2880000.00	Mastur	test	1	3	t	0.00	3600000.00	'-2022':4 '-3217':19 '000000062':1 '19':2 '2':28 '2017':25 'ada':17 'auto':7 'baf':9 'bussan':6 'cabang':10 'deddy':13 'e':18 'finance':5,8 'gudang':22 'indramayu':12 'jan':3 'nmax':21 'par':20 'pranoto':14 'pusat':11,23 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'yamaha':26
59	000000059	2022-01-31	2022-01-31	19750000.00	0.00	19750000.00	Mastur	test	14	3	t	0.00	19750000.00	'-1305':19 '-2022':4 '000000059':1 '2000':25 '31':2 '4':28 'ada':17 'cabang':10 'clip':8 'clipan':6 'deddy':13 'dl':20 'ertiga':21 'finance':5 'gudang':22 'indramayu':12 'jan':3 'k':9 'karawang':7 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'suzuk':26 't':18 'tahun':24
42	000000042	2022-02-16	2022-02-16	15000000.00	12.00	13200000.00	Mastur	test	18	3	t	0.00	15000000.00	'-2022':4 '-9049':21 '000000042':1 '1000':24 '16':2 '2020':28 '4':31 'ada':19 'brio':23 'cabang':12 'deddy':15 'feb':3 'finance':5,8 'gudang':25 'h':20 'honda':29 'indramayu':14 'mandir':6 'mtf':10 'pranoto':16 'pusat':13,26 'r4':32 'roda':30 's':11 'se':22 'semarang':9 'stnk':18 'stnk-ada':17 'tahun':27 'tunas':7
88	000000113	2022-03-24	2022-03-26	1250000.00	20.00	1000000.00	Opick	test	11	1	t	0.00	1250000.00	'-2022':4 '-4454':16 '000000113':1 '2':25 '2018':22 '24':2 'ada':14 'cabang':9 'finance':5 'gudang':19 'honda':23 'ih':17 'jatibarang':10,20 'kp':8 'kredit':6 'mar':3 'plus':7 'r2':26 'roda':24 'sonic':18 'stnk':13 'stnk-ada':12 'syaenudin':11 't':15 'tahun':21
90	000000115	2022-03-25	2022-03-26	1450000.00	20.00	1160000.00	Opick	test	13	1	t	0.00	1450000.00	'-2022':4 '-5948':15 '000000115':1 '2':24 '2011':21 '25':2 'ada':13 'cabang':8 'e':14 'finance':5 'gapara':6 'gudang':18 'jatibarang':9,19 'lz':16 'mar':3 'mpr':7 'r2':25 'roda':23 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':20 'xeon':17 'yamaha':22
93	000000118	2022-03-25	2022-03-26	1250000.00	20.00	1000000.00	Opick	test	2	1	t	0.00	1250000.00	'-2022':4 '-4560':17 '000000118':1 '2':26 '2015':23 '25':2 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'qv':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
95	000000120	2022-03-26	2022-03-26	1700000.00	20.00	1360000.00	Opick	test	2	3	t	0.00	1700000.00	'-2000':19 '-2022':4 '000000120':1 '2':28 '2021':25 '26':2 'ada':17 'adira':9 'auto':6 'cabang':10 'deddy':13 'discret':7 'e':18 'finance':5,8 'gudang':22 'indramayu':12 'mar':3 'nmax':21 'pranoto':14 'pusat':11,23 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'xx':20 'yamaha':26
98	000000136	2022-03-28	2022-03-28	10000000.00	15.00	8500000.00	Opick	test	3	3	t	0.00	10000000.00	'-2022':4 '-8072':19 '000000136':1 '2022':26 '28':2 '4':29 'ada':17 'cabang':10 'daihatsu':27 'deddy':13 'e':18 'finance':5,8 'grand':21 'gudang':23 'indramayu':12 'mandir':6 'mar':3 'max':22 'mtf':9 'pranoto':14 'pusat':11,24 'qb':20 'r4':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'tunas':7
71	000000071	2021-09-21	2021-09-21	1300000.00	20.00	1040000.00	Mastur	test	7	1	t	0.00	1300000.00	'-2021':4 '-4261':17 '000000071':1 '2':28 '2017':25 '21':2 'ada':15 'cabang':10 'finance':5,8 'gudang':22 'jatibarang':11,23 'king':21 'mandir':6 'muf':9 'mx':20 'mx-king':19 'r2':29 'roda':27 'sep':3 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':24 'utama':7 'yamaha':26 'yq':18
69	000000069	2021-12-09	2021-12-09	1400000.00	20.00	1120000.00	Mastur	test	2	1	t	0.00	1400000.00	'-2021':4 '-3433':17 '000000069':1 '09':2 '2':26 '2019':23 'ada':15 'adira':9 'auto':6 'b':16 'beat':19 'cabang':10 'des':3 'discret':7 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'usn':18
72	000000072	2021-12-09	2021-12-09	1250000.00	20.00	1000000.00	Mastur	test	2	1	t	0.00	1250000.00	'-2021':4 '-3430':17 '000000072':1 '09':2 '2':26 '2016':23 'ada':15 'adira':9 'auto':6 'b':16 'beat':19 'cabang':10 'des':3 'discret':7 'ejx':18 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
20	000000020	2022-01-21	2022-01-21	900000.00	20.00	720000.00	Mastur	test	11	1	t	0.00	900000.00	'-2022':4 '-5253':16 '000000020':1 '125':19 '2':26 '2013':23 '21':2 'ada':14 'cabang':9 'e':15 'finance':5 'gudang':20,21 'honda':24 'jan':3 'jatibarang':10 'kp':8 'kredit':6 'plus':7 'r2':27 'roda':25 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':22 'ty':17 'vario':18
3	000000003	2022-02-07	2022-02-07	1300000.00	20.00	1040000.00	Mastur	test	1	3	t	0.00	1300000.00	'-2022':4 '-5125':19 '000000003':1 '07':2 '2':29 '2018':26 'ada':17 'auto':7 'baf':9 'bussan':6 'cabang':10 'deddy':13 'e':18 'feb':3 'finance':5,8 'gudang':23 'indramayu':12 'm3':22 'mio':21 'pbc':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'yamaha':27
22	000000022	2022-02-23	2022-02-23	950000.00	20.00	760000.00	Mastur	test	1	1	t	0.00	950000.00	'-2022':4 '-2830':17 '000000022':1 '2':27 '2015':24 '23':2 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'feb':3 'finance':5,8 'gudang':21,22 'jatibarang':11 'm3':20 'mio':19 'qr':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
5	000000005	2022-02-24	2022-02-24	1500000.00	20.00	1200000.00	Mastur	test	7	3	t	0.00	1500000.00	'-2022':4 '-4096':19 '000000005':1 '125':22 '2':29 '2017':26 '24':2 'ada':17 'cabang':10 'deddy':13 'e':18 'feb':3 'finance':5,8 'fino':21 'gudang':23 'indramayu':12 'mandir':6 'muf':9 'paq':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'utama':7 'yamaha':27
29	000000029	2022-03-09	2022-03-09	1450000.00	20.00	1160000.00	Mastur	test	13	1	t	0.00	1450000.00	'-2022':4 '-4544':15 '000000029':1 '09':2 '2':24 '2015':21 'ada':13 'cabang':8 'e':14 'finance':5 'gapara':6 'gudang':18,19 'jatibarang':9 'jd':16 'mar':3 'mio':17 'mpr':7 'r2':25 'roda':23 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':20 'yamaha':22
31	000000031	2022-03-12	2022-03-12	1800000.00	20.00	1440000.00	Mastur	test	2	1	t	0.00	1800000.00	'-15':20 '-2022':4 '-2391':17 '000000031':1 '12':2 '2':27 '2017':24 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':21 'jatibarang':11,22 'jm':18 'mar':3 'r':19 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
35	000000035	2022-03-18	2022-03-18	1800000.00	20.00	1440000.00	Mastur	test	2	1	t	0.00	1800000.00	'-15':20 '-2022':4 '-2110':17 '000000035':1 '18':2 '2':27 '2017':24 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'gudang':21 'jatibarang':11,22 'mar':3 'r':19 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':23 'yamaha':25 'yv':18
96	000000132	2022-03-21	2022-03-28	14000000.00	12.14	12300000.00	Opick	test	19	3	t	0.00	14000000.00	'-1301':17 '-2022':4 '000000132':1 '1000':20 '2022':24 '21':2 '4':27 'ada':15 'b':16 'brio':19 'cabang':8 'clip':7 'clipan':6 'deddy':11 'finance':5 'gudang':21 'honda':25 'indramayu':10 'mar':3 'pranoto':12 'pusat':9,22 'r4':28 'roda':26 'stnk':14 'stnk-ada':13 'tahun':23 'uzr':18
12	000000012	2021-11-04	2021-11-04	1000000.00	20.00	800000.00	Mastur	test	6	1	t	0.00	1000000.00	'-2021':4 '-3479':15 '000000012':1 '04':2 '2':25 '2016':22 'ada':13 'b':14 'cabang':8 'col':7 'collectius':6 'finance':5 'gudang':19,20 'jatibarang':9 'm3':18 'mio':17 'nop':3 'r2':26 'roda':24 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':21 'uju':16 'yamaha':23
14	000000014	2021-12-07	2021-12-07	1500000.00	20.00	1200000.00	Mastur	test	9	1	t	0.00	1500000.00	'-2021':4 '-3521':18 '000000014':1 '07':2 '2':28 '2012':25 'ada':16 'cabang':11 'des':3 'finance':5,9 'fu':21 'gudang':22,23 'jatibarang':12 'kl':19 'mitra':6 'mpmf':10 'mustika':8 'pinasthika':7 'r2':29 'roda':27 'satria':20 'stnk':15 'stnk-ada':14 'suzuk':26 'syaenudin':13 't':17 'tahun':24
1	000000001	2021-12-15	2021-12-15	1300000.00	20.00	1040000.00	Mastur	test	5	3	t	0.00	1300000.00	'-2021':4 '-5605':19 '000000001':1 '15':2 '2':29 '2017':26 'ada':17 'cabang':10 'deddy':13 'des':3 'e':18 'finance':5 'gudang':23 'indramayu':12 'kredit':7 'mio':21 'motor':8 'oto':6 'otto':9 'pas':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'yamaha':27 'z':22
2	000000002	2022-01-10	2022-01-10	1300000.00	20.00	1040000.00	Mastur	test	6	3	t	0.00	1300000.00	'-2022':4 '-3977':17 '000000002':1 '10':2 '2':26 '2016':23 'ada':15 'cabang':8 'col':7 'collectius':6 'deddy':11 'e':16 'finance':5 'gudang':20 'indramayu':10 'jan':3 'mio':19 'pac':18 'pranoto':12 'pusat':9,21 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'tahun':22 'yamaha':24
19	000000019	2022-01-26	2022-01-26	1000000.00	20.00	800000.00	Mastur	test	2	1	t	0.00	1000000.00	'-2022':4 '-5638':17 '000000019':1 '2':26 '2018':23 '26':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jan':3 'jatibarang':11,21 'pav':18 'r2':27 'revo':19 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
79	000000105	2022-03-19	2022-03-26	1400000.00	20.00	1120000.00	Opick	test	2	1	t	0.00	1400000.00	'-2022':4 '-2282':17 '000000105':1 '19':2 '2':26 '2018':23 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'pbc':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
82	000000108	2022-03-22	2022-03-26	1200000.00	20.00	960000.00	Opick	test	2	1	t	0.00	1200000.00	'-2022':4 '-5117':17 '000000108':1 '2':26 '2018':23 '22':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'pba':18 'r2':27 'roda':25 'scoopy':19 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
64	000000064	2021-09-13	2021-09-13	1900000.00	15.79	1600000.00	Mastur	test	7	1	t	0.00	1900000.00	'-2021':4 '-5097':17 '000000064':1 '13':2 '150':20 '2':27 '2018':24 'ada':15 'cabang':10 'finance':5,8 'gudang':21 'honda':25 'jatibarang':11,22 'mandir':6 'muf':9 'r2':28 'roda':26 'sep':3 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':23 'utama':7 'vario':19 'zb':18
85	000000111	2022-03-24	2022-03-26	1400000.00	20.00	1120000.00	Opick	test	2	1	t	0.00	1400000.00	'-2022':4 '-2867':17 '000000111':1 '2':27 '2019':24 '24':2 'ada':15 'adira':9 'auf':18 'auto':6 'cabang':10 'discret':7 'finance':5,8 'g':16 'gudang':21 'jatibarang':11,22 'm3':20 'mar':3 'mio':19 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
68	000000068	2021-12-08	2021-12-08	1200000.00	20.00	960000.00	Mastur	test	20	1	t	0.00	1200000.00	'-2021':4 '-3256':16 '000000068':1 '08':2 '2':25 '2013':22 'ada':14 'b':15 'beat':18 'cabang':9 'des':3 'finance':5,7 'gudang':19 'honda':23 'jatibarang':10,20 'pwy':17 'r2':26 'rad':8 'radana':6 'roda':24 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':21
51	000000051	2021-12-21	2021-12-21	25000000.00	8.00	23000000.00	Mastur	test	19	3	t	0.00	25000000.00	'-1250':17 '-2021':4 '000000051':1 '2013':23 '21':2 '4':26 'ada':15 'cabang':8 'clip':7 'clipan':6 'deddy':11 'des':3 'du':18 'finance':5 'gudang':20 'indramayu':10 'ios':19 'pranoto':12 'pusat':9,21 'r4':27 'roda':25 'stnk':14 'stnk-ada':13 't':16 'tahun':22 'toyota':24
56	000000056	2022-01-21	2022-01-21	22000000.00	9.09	20000000.00	Mastur	test	18	3	t	0.00	22000000.00	'-2022':4 '-9086':21 '000000056':1 '2018':27 '21':2 '4':30 'ada':19 'agya':23 'cabang':12 'deddy':15 'finance':5,8 'gudang':24 'h':20 'indramayu':14 'jan':3 'mandir':6 'mtf':10 'pranoto':16 'pusat':13,25 'r4':31 'roda':29 's':11 'semarang':9 'stnk':18 'stnk-ada':17 'tahun':26 'te':22 'toyota':28 'tunas':7
75	000000075	2022-02-07	2022-02-07	1600000.00	20.00	1280000.00	Mastur	test	13	1	t	0.00	1600000.00	'-2022':4 '-4080':15 '000000075':1 '07':2 '2':25 '2018':22 'ada':13 'cabang':8 'e':14 'feb':3 'finance':5 'gapara':6 'gudang':19 'jatibarang':9,20 'm3':18 'mio':17 'mpr':7 'r2':26 'roda':24 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':21 'uo':16 'yamaha':23
27	000000027	2022-03-09	2022-03-09	700000.00	20.00	560000.00	Mastur	test	2	1	t	0.00	700000.00	'-2022':4 '-6819':17 '000000027':1 '09':2 '2':26 '2014':23 'ada':15 'adira':9 'auto':6 'b':16 'cabang':10 'discret':7 'finance':5,8 'gudang':20 'jatibarang':11,21 'mar':3 'pzi':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'xeon':19 'yamaha':24
34	000000034	2022-03-17	2022-03-17	900000.00	20.00	720000.00	Mastur	test	2	1	t	0.00	900000.00	'-2022':4 '-4593':17 '000000034':1 '17':2 '2':26 '2012':23 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'jatibarang':11,21 'jupiter':19 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'tq':18 'yamaha':24
37	000000037	2021-12-02	2021-12-02	20000000.00	15.00	17000000.00	Mastur	test	14	3	t	0.00	20000000.00	'-2021':4 '-8936':19 '000000037':1 '02':2 '2006':25 '4':28 'ada':17 'b':18 'cabang':10 'clip':8 'clipan':6 'deddy':13 'des':3 'finance':5 'gudang':22 'honda':26 'indramayu':12 'jazz':21 'k':9 'karawang':7 'no':20 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24
52	000000052	2021-12-23	2021-12-23	15000000.00	7.00	13950000.00	Mastur	test	14	3	t	0.00	15000000.00	'-1184':19 '-2021':4 '000000052':1 '1000':22 '2018':26 '23':2 '4':29 'ada':17 'brio':21 'cabang':10 'clip':8 'clipan':6 'deddy':13 'des':3 'finance':5 'ga':20 'gudang':23 'honda':27 'indramayu':12 'k':9 'karawang':7 'pranoto':14 'pusat':11,24 'r4':30 'roda':28 'stnk':16 'stnk-ada':15 't':18 'tahun':25
54	000000054	2021-12-28	2021-12-28	5400000.00	0.00	5400000.00	Mastur	test	2	3	t	0.00	5400000.00	'-1242':19 '-2021':4 '000000054':1 '2012':25 '28':2 '4':28 'ada':17 'adira':9 'auto':6 'cabang':10 'd':18 'deddy':13 'des':3 'discret':7 'finance':5,8 'gudang':22 'indramayu':12 'ou':20 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'toyota':26 'vios':21
39	000000039	2022-01-14	2022-01-14	26000000.00	6.92	24200000.00	Mastur	test	18	3	t	0.00	26000000.00	'-2022':4 '-8715':21 '000000039':1 '1000':24 '14':2 '2016':28 '4':31 'ada':19 'brio':23 'cabang':12 'deddy':15 'finance':5,8 'gp':22 'gudang':25 'h':20 'honda':29 'indramayu':14 'jan':3 'mandir':6 'mtf':10 'pranoto':16 'pusat':13,26 'r4':32 'roda':30 's':11 'semarang':9 'stnk':18 'stnk-ada':17 'tahun':27 'tunas':7
36	000000036	2022-03-18	2022-03-18	1500000.00	20.00	1200000.00	Mastur	test	1	1	t	0.00	1500000.00	'-2022':4 '-5713':17 '000000036':1 '18':2 '2':27 '2018':24 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'finance':5,8 'gudang':21 'jatibarang':11,22 'm3':20 'mar':3 'mio':19 'pav':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
94	000000119	2022-03-23	2022-03-26	1600000.00	20.00	1280000.00	Opick	test	1	3	t	0.00	1600000.00	'-2022':4 '-2753':19 '000000119':1 '125':22 '2':29 '2018':26 '23':2 'ada':17 'auto':7 'baf':9 'bussan':6 'cabang':10 'deddy':13 'e':18 'finance':5,8 'fino':21 'gudang':23 'indramayu':12 'mar':3 'pba':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'yamaha':27
87	000000112	2022-03-24	2022-03-26	1000000.00	20.00	800000.00	Opick	test	17	1	t	0.00	1000000.00	'-2022':4 '-4857':16 '000000112':1 '2':25 '2016':22 '24':2 'ada':14 'beat':18 'bfi':6,8 'cabang':9 'e':15 'finance':5,7 'gudang':19 'honda':23 'jatibarang':10,20 'mar':3 'pag':17 'r2':26 'roda':24 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':21
60	000000060	2021-09-30	2021-09-30	920000.00	0.00	920000.00	Mastur	test	2	3	t	0.00	920000.00	'-2021':4 '-6181':19 '000000060':1 '2':28 '2018':25 '30':2 'ada':17 'adira':9 'auto':6 'beat':21 'cabang':10 'deddy':13 'discret':7 'f':18 'fch':20 'finance':5,8 'gudang':22 'honda':26 'indramayu':12 'pranoto':14 'pusat':11,23 'r2':29 'roda':27 'sep':3 'stnk':16 'stnk-ada':15 'tahun':24
65	000000065	2021-10-05	2021-10-05	3200000.00	0.00	3200000.00	Mastur	test	1	1	t	0.00	3200000.00	'-2021':4 '-2676':17 '000000065':1 '05':2 '2':26 '2019':23 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'finance':5,8 'gudang':20 'jatibarang':11,21 'nmax':19 'okt':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'ux':18 'yamaha':24
13	000000013	2021-12-06	2021-12-06	1000000.00	20.00	800000.00	Mastur	test	6	1	t	0.00	1000000.00	'-2021':4 '-2417':15 '000000013':1 '06':2 '2':25 '2017':22 'ada':13 'cabang':8 'col':7 'collectius':6 'des':3 'e':14 'finance':5 'gudang':19,20 'jatibarang':9 'm3':18 'mio':17 'pao':16 'r2':26 'roda':24 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':21 'yamaha':23
67	000000067	2021-12-07	2021-12-07	2000000.00	20.00	1600000.00	Mastur	test	21	1	t	0.00	2000000.00	'-2021':4 '-4845':18 '000000067':1 '07':2 '2':27 '2020':24 'ada':16 'auto':7 'cabang':11 'central':8 'des':3 'finance':5,9 'gudang':21 'iq':19 'jatibarang':12,22 'macf':10 'mega':6 'nmax':20 'r2':28 'roda':26 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 'tahun':23 'yamaha':25
63	000000063	2021-12-24	2021-12-24	1800000.00	20.00	1440000.00	Mastur	test	5	3	t	0.00	1800000.00	'-2021':4 '-2113':19 '000000063':1 '2':28 '2019':25 '24':2 'ada':17 'cabang':10 'deddy':13 'des':3 'e':18 'finance':5 'gudang':22 'honda':26 'indramayu':12 'kredit':7 'motor':8 'oto':6 'otto':9 'pbm':20 'pcx':21 'pranoto':14 'pusat':11,23 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24
38	000000038	2021-12-29	2021-12-29	13000000.00	7.69	12000000.00	Mastur	test	14	3	t	0.00	13000000.00	'-1412':19 '-2021':4 '000000038':1 '2004':25 '29':2 '4':28 'ada':17 'cabang':10 'carry':21 'clip':8 'clipan':6 'deddy':13 'des':3 'finance':5 'gudang':22 'indramayu':12 'k':9 'karawang':7 'km':20 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'suzuk':26 't':18 'tahun':24
61	000000061	2021-12-29	2021-12-29	1500000.00	20.00	1200000.00	Mastur	test	7	3	t	0.00	1500000.00	'-2021':4 '-3310':19 '000000061':1 '2':28 '2015':25 '29':2 'ada':17 'cabang':10 'deddy':13 'des':3 'e':18 'finance':5,8 'gudang':22 'indramayu':12 'mandir':6 'muf':9 'pranoto':14 'pusat':11,23 'qr':20 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'utama':7 'vixion':21 'yamaha':26
73	000000073	2021-12-30	2021-12-30	1250000.00	20.00	1000000.00	Mastur	test	2	1	t	0.00	1250000.00	'-2021':4 '-3351':17 '000000073':1 '2':27 '2015':24 '30':2 'ada':15 'adira':9 'auto':6 'b':16 'beat':19 'cabang':10 'des':3 'discret':7 'finance':5,8 'gudang':21 'honda':25 'jatibarang':11,22 'kuh':18 'pop':20 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23
70	000000070	2022-01-03	2022-01-03	1400000.00	20.00	1120000.00	Mastur	test	5	1	t	0.00	1400000.00	'-2022':4 '-4654':17 '000000070':1 '03':2 '2':26 '2018':23 'ada':15 'b':16 'beat':19 'cabang':10 'finance':5 'fsg':18 'gudang':20 'honda':24 'jan':3 'jatibarang':11,21 'kredit':7 'motor':8 'oto':6 'otto':9 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
55	000000055	2022-01-17	2022-01-17	38400000.00	0.00	38400000.00	Mastur	test	14	3	t	0.00	38400000.00	'-1729':19 '-2022':4 '000000055':1 '17':2 '2018':25 '4':28 'ada':17 'bf':20 'brv':21 'cabang':10 'clip':8 'clipan':6 'deddy':13 'finance':5 'gudang':22 'honda':26 'indramayu':12 'jan':3 'k':9 'karawang':7 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 't':18 'tahun':24
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
\.


--
-- Data for Name: trx; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trx (id, ref_id, division, descriptions, trx_date, memo, trx_token) FROM stdin;
207	64	trx-order	Piutang jasa Mandiri Utama Finance (MUF) Order SPK: /000000064	2022-03-28	Kendaraan R2 Honda Vario 150 , Nopol T 5097 ZB	'/000000064':1 '/ref-64':15 '150':5 '5097':7 'finance':11 'honda':3 'jatibarang':13 'mandir':9 'muf':12 'order':18 'r2':2 'syaenudin':14 't':6 'trx':17 'trx-order':16 'utama':10 'vario':4 'zb':8
208	71	trx-order	Piutang jasa Mandiri Utama Finance (MUF) Order SPK: /000000071	2022-03-28	Kendaraan R2 Yamaha MX-King , Nopol T 4261 YQ	'/000000071':1 '/ref-71':16 '4261':8 'finance':12 'jatibarang':14 'king':6 'mandir':10 'muf':13 'mx':5 'mx-king':4 'order':19 'r2':2 'syaenudin':15 't':7 'trx':18 'trx-order':17 'utama':11 'yamaha':3 'yq':9
209	11	trx-order	Piutang jasa FIF Group (FIF) Order SPK: /000000011	2022-03-28	Kendaraan R2 Yamaha Jupiter , Nopol E 4892 TK	'/000000011':1 '/ref-11':13 '4892':6 'e':5 'fif':8,10 'group':9 'jatibarang':11 'jupiter':4 'order':16 'r2':2 'syaenudin':12 'tk':7 'trx':15 'trx-order':14 'yamaha':3
210	60	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000060	2022-03-28	Kendaraan R2 Honda BEAT , Nopol F 6181 FCH	'/000000060':1 '/ref-60':16 '6181':6 'adira':11 'auto':8 'beat':4 'deddy':14 'discret':9 'f':5 'fch':7 'finance':10 'honda':3 'indramayu':13 'order':19 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17
211	65	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000065	2022-03-28	Kendaraan R2 Yamaha NMax , Nopol E 2676 UX	'/000000065':1 '/ref-65':14 '2676':6 'auto':9 'baf':11 'bussan':8 'e':5 'finance':10 'jatibarang':12 'nmax':4 'order':17 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'ux':7 'yamaha':3
212	12	trx-order	Piutang jasa COLLECTIUS (COL) Order SPK: /000000012	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol B 3479 UJU	'/000000012':1 '/ref-12':13 '3479':7 'b':6 'col':10 'collectius':9 'jatibarang':11 'm3':5 'mio':4 'order':16 'r2':2 'syaenudin':12 'trx':15 'trx-order':14 'uju':8 'yamaha':3
213	49	trx-order	Piutang jasa BFI Finance (BFI) Order SPK: /000000049	2022-03-28	Kendaraan R4 Honda Jazz , Nopol H 8630 PP	'/000000049':1 '/ref-49':15 '8630':6 'bfi':8,10 'deddy':13 'finance':9 'h':5 'honda':3 'indramayu':12 'jazz':4 'order':18 'pp':7 'pranoto':14 'pusat':11 'r4':2 'trx':17 'trx-order':16
214	37	trx-order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000037	2022-03-28	Kendaraan R4 Honda Jazz , Nopol B 8936 NO	'/000000037':1 '/ref-37':16 '8936':6 'b':5 'clip':10 'clipan':8 'deddy':14 'honda':3 'indramayu':13 'jazz':4 'k':11 'karawang':9 'no':7 'order':19 'pranoto':15 'pusat':12 'r4':2 'trx':18 'trx-order':17
215	13	trx-order	Piutang jasa COLLECTIUS (COL) Order SPK: /000000013	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol E 2417 PAO	'/000000013':1 '/ref-13':13 '2417':7 'col':10 'collectius':9 'e':6 'jatibarang':11 'm3':5 'mio':4 'order':16 'pao':8 'r2':2 'syaenudin':12 'trx':15 'trx-order':14 'yamaha':3
216	66	trx-order	Piutang jasa WOM Finance (WOMF) Order SPK: /000000066	2022-03-28	Kendaraan R2 Honda BEAT , Nopol T 6934 YQ	'/000000066':1 '/ref-66':13 '6934':6 'beat':4 'finance':9 'honda':3 'jatibarang':11 'order':16 'r2':2 'syaenudin':12 't':5 'trx':15 'trx-order':14 'wom':8 'womf':10 'yq':7
217	14	trx-order	Piutang jasa Mitra Pinasthika Mustika Finance (MPMF) Order SPK: /000000014	2022-03-28	Kendaraan R2 Suzuki Satria FU , Nopol T 3521 KL	'/000000014':1 '/ref-14':16 '3521':7 'finance':12 'fu':5 'jatibarang':14 'kl':8 'mitra':9 'mpmf':13 'mustika':11 'order':19 'pinasthika':10 'r2':2 'satria':4 'suzuk':3 'syaenudin':15 't':6 'trx':18 'trx-order':17
218	67	trx-order	Piutang jasa Mega Auto Central Finance (MACF) Order SPK: /000000067	2022-03-28	Kendaraan R2 Yamaha NMax , Nopol T 4845 IQ	'/000000067':1 '/ref-67':15 '4845':6 'auto':9 'central':10 'finance':11 'iq':7 'jatibarang':13 'macf':12 'mega':8 'nmax':4 'order':18 'r2':2 'syaenudin':14 't':5 'trx':17 'trx-order':16 'yamaha':3
219	68	trx-order	Piutang jasa Radana Finance (RAD) Order SPK: /000000068	2022-03-28	Kendaraan R2 Honda BEAT , Nopol B 3256  PWY	'/000000068':1 '/ref-68':13 '3256':6 'b':5 'beat':4 'finance':9 'honda':3 'jatibarang':11 'order':16 'pwy':7 'r2':2 'rad':10 'radana':8 'syaenudin':12 'trx':15 'trx-order':14
220	69	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000069	2022-03-28	Kendaraan R2 Honda BEAT , Nopol B 3433 USN	'/000000069':1 '/ref-69':14 '3433':6 'adira':11 'auto':8 'b':5 'beat':4 'discret':9 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'usn':7
221	72	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000072	2022-03-28	Kendaraan R2 Honda BEAT , Nopol B 3430 EJX	'/000000072':1 '/ref-72':14 '3430':6 'adira':11 'auto':8 'b':5 'beat':4 'discret':9 'ejx':7 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 'trx':16 'trx-order':15
222	1	trx-order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000001	2022-03-28	Kendaraan R2 Yamaha Mio Z , Nopol E 5605 PAS	'/000000001':1 '/ref-1':17 '5605':7 'deddy':15 'e':6 'indramayu':14 'kredit':10 'mio':4 'motor':11 'order':20 'oto':9 'otto':12 'pas':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'yamaha':3 'z':5
223	50	trx-order	Piutang jasa Mandiri Tunas Finance Semarang (MTF S) Order SPK: /000000050	2022-03-28	Kendaraan R4 Honda Brio 1000 , Nopol H 9442 NG	'/000000050':1 '/ref-50':19 '1000':5 '9442':7 'brio':4 'deddy':17 'finance':11 'h':6 'honda':3 'indramayu':16 'mandir':9 'mtf':13 'ng':8 'order':22 'pranoto':18 'pusat':15 'r4':2 's':14 'semarang':12 'trx':21 'trx-order':20 'tunas':10
224	51	trx-order	Piutang jasa CLIPAN (CLIP) Order SPK: /000000051	2022-03-28	Kendaraan R4 Toyota Terios , Nopol T 1250 DU	'/000000051':1 '/ref-51':14 '1250':6 'clip':9 'clipan':8 'deddy':12 'du':7 'indramayu':11 'ios':4 'order':17 'pranoto':13 'pusat':10 'r4':2 't':5 'toyota':3 'trx':16 'trx-order':15
225	52	trx-order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000052	2022-03-28	Kendaraan R4 Honda Brio 1000 , Nopol T 1184 GA	'/000000052':1 '/ref-52':17 '1000':5 '1184':7 'brio':4 'clip':11 'clipan':9 'deddy':15 'ga':8 'honda':3 'indramayu':14 'k':12 'karawang':10 'order':20 'pranoto':16 'pusat':13 'r4':2 't':6 'trx':19 'trx-order':18
226	53	trx-order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000053	2022-03-28	Kendaraan R4 Daihatsu Xenia , Nopol D 1312 WF	'/000000053':1 '/ref-53':16 '1312':6 'clip':10 'clipan':8 'd':5 'daihatsu':3 'deddy':14 'indramayu':13 'k':11 'karawang':9 'order':19 'pranoto':15 'pusat':12 'r4':2 'trx':18 'trx-order':17 'wf':7 'xenia':4
227	63	trx-order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000063	2022-03-28	Kendaraan R2 Honda PCX , Nopol E 2113 PBM	'/000000063':1 '/ref-63':16 '2113':6 'deddy':14 'e':5 'honda':3 'indramayu':13 'kredit':9 'motor':10 'order':19 'oto':8 'otto':11 'pbm':7 'pcx':4 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17
228	54	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000054	2022-03-28	Kendaraan R4 Toyota Vios , Nopol D 1242 OU	'/000000054':1 '/ref-54':16 '1242':6 'adira':11 'auto':8 'd':5 'deddy':14 'discret':9 'finance':10 'indramayu':13 'order':19 'ou':7 'pranoto':15 'pusat':12 'r4':2 'toyota':3 'trx':18 'trx-order':17 'vios':4
231	73	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000073	2022-03-28	Kendaraan R2 Honda Beat Pop , Nopol B 3351 KUH	'/000000073':1 '/ref-73':15 '3351':7 'adira':12 'auto':9 'b':6 'beat':4 'discret':10 'finance':11 'honda':3 'jatibarang':13 'kuh':8 'order':18 'pop':5 'r2':2 'syaenudin':14 'trx':17 'trx-order':16
232	70	trx-order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000070	2022-03-28	Kendaraan R2 Honda BEAT , Nopol B 4654 FSG	'/000000070':1 '/ref-70':14 '4654':6 'b':5 'beat':4 'fsg':7 'honda':3 'jatibarang':12 'kredit':9 'motor':10 'order':17 'oto':8 'otto':11 'r2':2 'syaenudin':13 'trx':16 'trx-order':15
233	15	trx-order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000015	2022-03-28	Kendaraan R2 Yamaha Jupiter MX , Nopol T 4146 KO	'/000000015':1 '/ref-15':15 '4146':7 'jatibarang':13 'jupiter':4 'ko':8 'kredit':10 'motor':11 'mx':5 'order':18 'oto':9 'otto':12 'r2':2 'syaenudin':14 't':6 'trx':17 'trx-order':16 'yamaha':3
237	39	trx-order	Piutang jasa Mandiri Tunas Finance Semarang (MTF S) Order SPK: /000000039	2022-03-28	Kendaraan R4 Honda Brio 1000 , Nopol H 8715 GP	'/000000039':1 '/ref-39':19 '1000':5 '8715':7 'brio':4 'deddy':17 'finance':11 'gp':8 'h':6 'honda':3 'indramayu':16 'mandir':9 'mtf':13 'order':22 'pranoto':18 'pusat':15 'r4':2 's':14 'semarang':12 'trx':21 'trx-order':20 'tunas':10
238	40	trx-order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000040	2022-03-28	Kendaraan R4 Suzuki ERTIGA , Nopol T 1164 FQ	'/000000040':1 '/ref-40':16 '1164':6 'clip':10 'clipan':8 'deddy':14 'ertiga':4 'fq':7 'indramayu':13 'k':11 'karawang':9 'order':19 'pranoto':15 'pusat':12 'r4':2 'suzuk':3 't':5 'trx':18 'trx-order':17
246	19	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000019	2022-03-28	Kendaraan R2 Honda Revo , Nopol E 5638 PAV	'/000000019':1 '/ref-19':14 '5638':6 'adira':11 'auto':8 'discret':9 'e':5 'finance':10 'honda':3 'jatibarang':12 'order':17 'pav':7 'r2':2 'revo':4 'syaenudin':13 'trx':16 'trx-order':15
229	38	trx-order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000038	2022-03-28	Kendaraan R4 Suzuki Carry , Nopol T 1412 KM	'/000000038':1 '/ref-38':16 '1412':6 'carry':4 'clip':10 'clipan':8 'deddy':14 'indramayu':13 'k':11 'karawang':9 'km':7 'order':19 'pranoto':15 'pusat':12 'r4':2 'suzuk':3 't':5 'trx':18 'trx-order':17
234	2	trx-order	Piutang jasa COLLECTIUS (COL) Order SPK: /000000002	2022-03-28	Kendaraan R2 Yamaha Mio , Nopol E 3977 PAC	'/000000002':1 '/ref-2':14 '3977':6 'col':9 'collectius':8 'deddy':12 'e':5 'indramayu':11 'mio':4 'order':17 'pac':7 'pranoto':13 'pusat':10 'r2':2 'trx':16 'trx-order':15 'yamaha':3
244	57	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000057	2022-03-28	Kendaraan R4 Daihatsu Grand Max , Nopol E 1256 QD	'/000000057':1 '/ref-57':17 '1256':7 'adira':12 'auto':9 'daihatsu':3 'deddy':15 'discret':10 'e':6 'finance':11 'grand':4 'indramayu':14 'max':5 'order':20 'pranoto':16 'pusat':13 'qd':8 'r4':2 'trx':19 'trx-order':18
245	41	trx-order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000041	2022-03-28	Kendaraan R4 Honda Mobilio , Nopol T 1788 BC	'/000000041':1 '/ref-41':16 '1788':6 'bc':7 'clip':10 'clipan':8 'deddy':14 'honda':3 'indramayu':13 'k':11 'karawang':9 'mobilio':4 'order':19 'pranoto':15 'pusat':12 'r4':2 't':5 'trx':18 'trx-order':17
248	59	trx-order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000059	2022-03-28	Kendaraan R4 Suzuki ERTIGA , Nopol T 1305 DL	'/000000059':1 '/ref-59':16 '1305':6 'clip':10 'clipan':8 'deddy':14 'dl':7 'ertiga':4 'indramayu':13 'k':11 'karawang':9 'order':19 'pranoto':15 'pusat':12 'r4':2 'suzuk':3 't':5 'trx':18 'trx-order':17
249	3	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000003	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol E 5125 PBC	'/000000003':1 '/ref-3':17 '5125':7 'auto':10 'baf':12 'bussan':9 'deddy':15 'e':6 'finance':11 'indramayu':14 'm3':5 'mio':4 'order':20 'pbc':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'yamaha':3
250	75	trx-order	Piutang jasa MEGAPARA (MPR) Order SPK: /000000075	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol E 4080 UO	'/000000075':1 '/ref-75':13 '4080':7 'e':6 'gapara':9 'jatibarang':11 'm3':5 'mio':4 'mpr':10 'order':16 'r2':2 'syaenudin':12 'trx':15 'trx-order':14 'uo':8 'yamaha':3
252	4	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000004	2022-03-28	Kendaraan R2 Yamaha Vixion , Nopol BR 5080 PY	'/000000004':1 '/ref-4':16 '5080':6 'adira':11 'auto':8 'br':5 'deddy':14 'discret':9 'finance':10 'indramayu':13 'order':19 'pranoto':15 'pusat':12 'py':7 'r2':2 'trx':18 'trx-order':17 'vixion':4 'yamaha':3
230	61	trx-order	Piutang jasa Mandiri Utama Finance (MUF) Order SPK: /000000061	2022-03-28	Kendaraan R2 Yamaha Vixion , Nopol E 3310 QR	'/000000061':1 '/ref-61':16 '3310':6 'deddy':14 'e':5 'finance':10 'indramayu':13 'mandir':8 'muf':11 'order':19 'pranoto':15 'pusat':12 'qr':7 'r2':2 'trx':18 'trx-order':17 'utama':9 'vixion':4 'yamaha':3
235	16	trx-order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000016	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 3848 UB	'/000000016':1 '/ref-16':14 '3848':6 'beat':4 'e':5 'honda':3 'jatibarang':12 'kredit':9 'motor':10 'order':17 'oto':8 'otto':11 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'ub':7
236	17	trx-order	Piutang jasa Top Finance Company (TFC) Order SPK: /000000017	2022-03-28	Kendaraan R2 Honda Supra X-125 , Nopol T 3828 FW	'-125':6 '/000000017':1 '/ref-17':16 '3828':8 'company':12 'finance':11 'fw':9 'honda':3 'jatibarang':14 'order':19 'r2':2 'supra':4 'syaenudin':15 't':7 'tfc':13 'top':10 'trx':18 'trx-order':17 'x':5
239	55	trx-order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000055	2022-03-28	Kendaraan R4 Honda BRV , Nopol T 1729 BF	'/000000055':1 '/ref-55':16 '1729':6 'bf':7 'brv':4 'clip':10 'clipan':8 'deddy':14 'honda':3 'indramayu':13 'k':11 'karawang':9 'order':19 'pranoto':15 'pusat':12 'r4':2 't':5 'trx':18 'trx-order':17
240	18	trx-order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000018	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol E 6716 IX	'/000000018':1 '/ref-18':15 '6716':7 'e':6 'ix':8 'jatibarang':13 'kredit':10 'm3':5 'mio':4 'motor':11 'order':18 'oto':9 'otto':12 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
241	62	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000062	2022-03-28	Kendaraan R2 Yamaha NMax , Nopol E 3217 PAR	'/000000062':1 '/ref-62':16 '3217':6 'auto':9 'baf':11 'bussan':8 'deddy':14 'e':5 'finance':10 'indramayu':13 'nmax':4 'order':19 'par':7 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17 'yamaha':3
242	20	trx-order	Piutang jasa Kredit Plus (KP+) Order SPK: /000000020	2022-03-28	Kendaraan R2 Honda Vario 125 , Nopol E 5253 TY	'/000000020':1 '/ref-20':14 '125':5 '5253':7 'e':6 'honda':3 'jatibarang':12 'kp':11 'kredit':9 'order':17 'plus':10 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'ty':8 'vario':4
243	56	trx-order	Piutang jasa Mandiri Tunas Finance Semarang (MTF S) Order SPK: /000000056	2022-03-28	Kendaraan R4 Toyota AGYA , Nopol H 9086 TE	'/000000056':1 '/ref-56':18 '9086':6 'agya':4 'deddy':16 'finance':10 'h':5 'indramayu':15 'mandir':8 'mtf':12 'order':21 'pranoto':17 'pusat':14 'r4':2 's':13 'semarang':11 'te':7 'toyota':3 'trx':20 'trx-order':19 'tunas':9
247	58	trx-order	Piutang jasa Clipan Bekasi (CLIP B) Order SPK: /000000058	2022-03-28	Kendaraan R4 Toyota Terios , Nopol B 2281 SBT	'/000000058':1 '/ref-58':16 '2281':6 'b':5,11 'bekasi':9 'clip':10 'clipan':8 'deddy':14 'indramayu':13 'ios':4 'order':19 'pranoto':15 'pusat':12 'r4':2 'sbt':7 'toyota':3 'trx':18 'trx-order':17
251	42	trx-order	Piutang jasa Mandiri Tunas Finance Semarang (MTF S) Order SPK: /000000042	2022-03-28	Kendaraan R4 Honda Brio 1000 , Nopol H 9049 SE	'/000000042':1 '/ref-42':19 '1000':5 '9049':7 'brio':4 'deddy':17 'finance':11 'h':6 'honda':3 'indramayu':16 'mandir':9 'mtf':13 'order':22 'pranoto':18 'pusat':15 'r4':2 's':14 'se':8 'semarang':12 'trx':21 'trx-order':20 'tunas':10
253	21	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000021	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol B 6262 VKY	'/000000021':1 '/ref-21':15 '6262':7 'auto':10 'b':6 'baf':12 'bussan':9 'finance':11 'jatibarang':13 'm3':5 'mio':4 'order':18 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'vky':8 'yamaha':3
254	22	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000022	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol E 2830 QR	'/000000022':1 '/ref-22':15 '2830':7 'auto':10 'baf':12 'bussan':9 'e':6 'finance':11 'jatibarang':13 'm3':5 'mio':4 'order':18 'qr':8 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
255	5	trx-order	Piutang jasa Mandiri Utama Finance (MUF) Order SPK: /000000005	2022-03-28	Kendaraan R2 Yamaha Fino 125  , Nopol E 4096 PAQ	'/000000005':1 '/ref-5':17 '125':5 '4096':7 'deddy':15 'e':6 'finance':11 'fino':4 'indramayu':14 'mandir':9 'muf':12 'order':20 'paq':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'utama':10 'yamaha':3
256	24	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000024	2022-03-28	Kendaraan R2 Yamaha Mio S , Nopol E 2146 QAF	'/000000024':1 '/ref-24':15 '2146':7 'auto':10 'baf':12 'bussan':9 'e':6 'finance':11 'jatibarang':13 'mio':4 'order':18 'qaf':8 'r2':2 's':5 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
257	6	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000006	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 2633 PAC	'/000000006':1 '/ref-6':16 '2633':6 'adira':11 'auto':8 'beat':4 'deddy':14 'discret':9 'e':5 'finance':10 'honda':3 'indramayu':13 'order':19 'pac':7 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17
258	26	trx-order	Piutang jasa Kredit Plus (KP+) Order SPK: /000000026	2022-03-28	Kendaraan R2 Honda Vario 150 , Nopol T 2891 WP	'/000000026':1 '/ref-26':14 '150':5 '2891':7 'honda':3 'jatibarang':12 'kp':11 'kredit':9 'order':17 'plus':10 'r2':2 'syaenudin':13 't':6 'trx':16 'trx-order':15 'vario':4 'wp':8
259	74	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000074	2022-03-28	Kendaraan R2 Yamaha Mio , Nopol E 6871 CM	'/000000074':1 '/ref-74':14 '6871':6 'adira':11 'auto':8 'cm':7 'discret':9 'e':5 'finance':10 'jatibarang':12 'mio':4 'order':17 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'yamaha':3
260	27	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000027	2022-03-28	Kendaraan R2 Yamaha Xeon , Nopol B 6819 PZI	'/000000027':1 '/ref-27':14 '6819':6 'adira':11 'auto':8 'b':5 'discret':9 'finance':10 'jatibarang':12 'order':17 'pzi':7 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'xeon':4 'yamaha':3
261	28	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000028	2022-03-28	Kendaraan R2 Yamaha Mio Z , Nopol T 4487 PJ	'/000000028':1 '/ref-28':15 '4487':7 'adira':12 'auto':9 'discret':10 'finance':11 'jatibarang':13 'mio':4 'order':18 'pj':8 'r2':2 'syaenudin':14 't':6 'trx':17 'trx-order':16 'yamaha':3 'z':5
262	29	trx-order	Piutang jasa MEGAPARA (MPR) Order SPK: /000000029	2022-03-28	Kendaraan R2 Yamaha Mio , Nopol E 4544 JD	'/000000029':1 '/ref-29':12 '4544':6 'e':5 'gapara':8 'jatibarang':10 'jd':7 'mio':4 'mpr':9 'order':15 'r2':2 'syaenudin':11 'trx':14 'trx-order':13 'yamaha':3
263	30	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000030	2022-03-28	Kendaraan R2 Honda Verza , Nopol T 3615 ZD	'/000000030':1 '/ref-30':14 '3615':6 'adira':11 'auto':8 'discret':9 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 't':5 'trx':16 'trx-order':15 'verza':4 'zd':7
264	31	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000031	2022-03-28	Kendaraan R2 Yamaha R-15 , Nopol E 2391 JM	'-15':5 '/000000031':1 '/ref-31':15 '2391':7 'adira':12 'auto':9 'discret':10 'e':6 'finance':11 'jatibarang':13 'jm':8 'order':18 'r':4 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
265	32	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000032	2022-03-28	Kendaraan R2 Honda BEAT , Nopol T 2191 YS	'/000000032':1 '/ref-32':14 '2191':6 'adira':11 'auto':8 'beat':4 'discret':9 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 't':5 'trx':16 'trx-order':15 'ys':7
266	7	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000007	2022-03-28	Kendaraan R2 Yamaha Fino 125  , Nopol E 2033 PBJ	'/000000007':1 '/ref-7':17 '125':5 '2033':7 'auto':10 'baf':12 'bussan':9 'deddy':15 'e':6 'finance':11 'fino':4 'indramayu':14 'order':20 'pbj':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'yamaha':3
267	33	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000033	2022-03-28	Kendaraan R2 Honda GENIO , Nopol T 5856 ZT	'/000000033':1 '/ref-33':14 '5856':6 'adira':11 'auto':8 'discret':9 'finance':10 'genio':4 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 't':5 'trx':16 'trx-order':15 'zt':7
268	34	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000034	2022-03-28	Kendaraan R2 Yamaha Jupiter , Nopol E 4593 TQ	'/000000034':1 '/ref-34':14 '4593':6 'adira':11 'auto':8 'discret':9 'e':5 'finance':10 'jatibarang':12 'jupiter':4 'order':17 'r2':2 'syaenudin':13 'tq':7 'trx':16 'trx-order':15 'yamaha':3
269	35	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000035	2022-03-28	Kendaraan R2 Yamaha R-15 , Nopol T 2110 YV	'-15':5 '/000000035':1 '/ref-35':15 '2110':7 'adira':12 'auto':9 'discret':10 'finance':11 'jatibarang':13 'order':18 'r':4 'r2':2 'syaenudin':14 't':6 'trx':17 'trx-order':16 'yamaha':3 'yv':8
270	36	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000036	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol E 5713 PAV	'/000000036':1 '/ref-36':15 '5713':7 'auto':10 'baf':12 'bussan':9 'e':6 'finance':11 'jatibarang':13 'm3':5 'mio':4 'order':18 'pav':8 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
271	10	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000010	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 5474 Q	'/000000010':1 '/ref-10':16 '5474':6 'adira':11 'auto':8 'beat':4 'deddy':14 'discret':9 'e':5 'finance':10 'honda':3 'indramayu':13 'order':19 'pranoto':15 'pusat':12 'q':7 'r2':2 'trx':18 'trx-order':17
272	77	trx-order	Piutang jasa Mega Auto Central Finance (MACF) Order SPK: /000000103	2022-03-28	Kendaraan R2 Honda Vario 125 , Nopol E 5826 CQ	'/000000103':1 '/ref-77':16 '125':5 '5826':7 'auto':10 'central':11 'cq':8 'e':6 'finance':12 'honda':3 'jatibarang':14 'macf':13 'mega':9 'order':19 'r2':2 'syaenudin':15 'trx':18 'trx-order':17 'vario':4
273	78	trx-order	Piutang jasa BFI Finance (BFI) Order SPK: /000000104	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 6505 QR	'/000000104':1 '/ref-78':13 '6505':6 'beat':4 'bfi':8,10 'e':5 'finance':9 'honda':3 'jatibarang':11 'order':16 'qr':7 'r2':2 'syaenudin':12 'trx':15 'trx-order':14
274	79	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000105	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 2282 PBC	'/000000105':1 '/ref-79':14 '2282':6 'adira':11 'auto':8 'beat':4 'discret':9 'e':5 'finance':10 'honda':3 'jatibarang':12 'order':17 'pbc':7 'r2':2 'syaenudin':13 'trx':16 'trx-order':15
275	80	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000106	2022-03-28	Kendaraan R2 Yamaha Fino 125  , Nopol E 5737 PBO	'/000000106':1 '/ref-80':15 '125':5 '5737':7 'adira':12 'auto':9 'discret':10 'e':6 'finance':11 'fino':4 'jatibarang':13 'order':18 'pbo':8 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
276	81	trx-order	Piutang jasa WOM Finance (WOMF) Order SPK: /000000107	2022-03-28	Kendaraan R2 Honda cb150 , Nopol E 2764 XL	'/000000107':1 '/ref-81':13 '2764':6 'cb150':4 'e':5 'finance':9 'honda':3 'jatibarang':11 'order':16 'r2':2 'syaenudin':12 'trx':15 'trx-order':14 'wom':8 'womf':10 'xl':7
277	96	trx-order	Piutang jasa CLIPAN (CLIP) Order SPK: /000000132	2022-03-28	Kendaraan R4 Honda Brio 1000 , Nopol B 1301 UZR	'/000000132':1 '/ref-96':15 '1000':5 '1301':7 'b':6 'brio':4 'clip':10 'clipan':9 'deddy':13 'honda':3 'indramayu':12 'order':18 'pranoto':14 'pusat':11 'r4':2 'trx':17 'trx-order':16 'uzr':8
278	97	trx-order	Piutang jasa CLIPAN (CLIP) Order SPK: /000000135	2022-03-28	Kendaraan R4 Honda Jazz , Nopol B 1942 EVK	'/000000135':1 '/ref-97':14 '1942':6 'b':5 'clip':9 'clipan':8 'deddy':12 'evk':7 'honda':3 'indramayu':11 'jazz':4 'order':17 'pranoto':13 'pusat':10 'r4':2 'trx':16 'trx-order':15
279	82	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000108	2022-03-28	Kendaraan R2 Honda scoopy , Nopol E 5117 PBA	'/000000108':1 '/ref-82':14 '5117':6 'adira':11 'auto':8 'discret':9 'e':5 'finance':10 'honda':3 'jatibarang':12 'order':17 'pba':7 'r2':2 'scoopy':4 'syaenudin':13 'trx':16 'trx-order':15
280	83	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000109	2022-03-28	Kendaraan R2 Honda scoopy , Nopol T 3802 ZR	'/000000109':1 '/ref-83':14 '3802':6 'adira':11 'auto':8 'discret':9 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'scoopy':4 'syaenudin':13 't':5 'trx':16 'trx-order':15 'zr':7
281	84	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000110	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 5462 QM	'/000000110':1 '/ref-84':14 '5462':6 'adira':11 'auto':8 'beat':4 'discret':9 'e':5 'finance':10 'honda':3 'jatibarang':12 'order':17 'qm':7 'r2':2 'syaenudin':13 'trx':16 'trx-order':15
282	94	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000119	2022-03-28	Kendaraan R2 Yamaha Fino 125  , Nopol E 2753 PBA	'/000000119':1 '/ref-94':17 '125':5 '2753':7 'auto':10 'baf':12 'bussan':9 'deddy':15 'e':6 'finance':11 'fino':4 'indramayu':14 'order':20 'pba':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'yamaha':3
285	88	trx-order	Piutang jasa Kredit Plus (KP+) Order SPK: /000000113	2022-03-28	Kendaraan R2 Honda sonic , Nopol T 4454 IH	'/000000113':1 '/ref-88':13 '4454':6 'honda':3 'ih':7 'jatibarang':11 'kp':10 'kredit':8 'order':16 'plus':9 'r2':2 'sonic':4 'syaenudin':12 't':5 'trx':15 'trx-order':14
286	89	trx-order	Piutang jasa WOM Finance (WOMF) Order SPK: /000000114	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 3561 PBI	'/000000114':1 '/ref-89':13 '3561':6 'beat':4 'e':5 'finance':9 'honda':3 'jatibarang':11 'order':16 'pbi':7 'r2':2 'syaenudin':12 'trx':15 'trx-order':14 'wom':8 'womf':10
288	93	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000118	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 4560 QV	'/000000118':1 '/ref-93':14 '4560':6 'adira':11 'auto':8 'beat':4 'discret':9 'e':5 'finance':10 'honda':3 'jatibarang':12 'order':17 'qv':7 'r2':2 'syaenudin':13 'trx':16 'trx-order':15
289	95	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000120	2022-03-28	Kendaraan R2 Yamaha NMax , Nopol E 2000 XX	'/000000120':1 '/ref-95':16 '2000':6 'adira':11 'auto':8 'deddy':14 'discret':9 'e':5 'finance':10 'indramayu':13 'nmax':4 'order':19 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17 'xx':7 'yamaha':3
283	85	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000111	2022-03-28	Kendaraan R2 Yamaha Mio M3 , Nopol G 2867 AUF	'/000000111':1 '/ref-85':15 '2867':7 'adira':12 'auf':8 'auto':9 'discret':10 'finance':11 'g':6 'jatibarang':13 'm3':5 'mio':4 'order':18 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
284	87	trx-order	Piutang jasa BFI Finance (BFI) Order SPK: /000000112	2022-03-28	Kendaraan R2 Honda BEAT , Nopol E 4857 PAG	'/000000112':1 '/ref-87':13 '4857':6 'beat':4 'bfi':8,10 'e':5 'finance':9 'honda':3 'jatibarang':11 'order':16 'pag':7 'r2':2 'syaenudin':12 'trx':15 'trx-order':14
291	98	trx-order	Piutang jasa Mandiri Tunas Finance (MTF) Order SPK: /000000136	2022-03-28	Kendaraan R4 Daihatsu Grand Max , Nopol E 8072 QB	'/000000136':1 '/ref-98':17 '8072':7 'daihatsu':3 'deddy':15 'e':6 'finance':11 'grand':4 'indramayu':14 'mandir':9 'max':5 'mtf':12 'order':20 'pranoto':16 'pusat':13 'qb':8 'r4':2 'trx':19 'trx-order':18 'tunas':10
292	99	trx-order	Piutang jasa SUZUKI FINANCE INDONESIA (SFI ) Order SPK: /000000139	2022-03-28	Kendaraan R4 Suzuki Carry , Nopol T 8066 EG	'/000000139':1 '/ref-99':16 '8066':6 'carry':4 'deddy':14 'eg':7 'finance':9 'indonesia':10 'indramayu':13 'order':19 'pranoto':15 'pusat':12 'r4':2 'sfi':11 'suzuk':3,8 't':5 'trx':18 'trx-order':17
294	102	trx-order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000147	2022-03-28	Kendaraan R2 Yamaha Fino 125  , Nopol E 6645 PAV	'/000000147':1 '/ref-102':15 '125':5 '6645':7 'auto':10 'baf':12 'bussan':9 'e':6 'finance':11 'fino':4 'jatibarang':13 'order':18 'pav':8 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
287	90	trx-order	Piutang jasa MEGAPARA (MPR) Order SPK: /000000115	2022-03-28	Kendaraan R2 Yamaha Xeon , Nopol E 5948 LZ	'/000000115':1 '/ref-90':12 '5948':6 'e':5 'gapara':8 'jatibarang':10 'lz':7 'mpr':9 'order':15 'r2':2 'syaenudin':11 'trx':14 'trx-order':13 'xeon':4 'yamaha':3
290	101	trx-order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000146	2022-03-28	Kendaraan R2 Suzuki Satria FU , Nopol E 6987 SD	'/000000146':1 '/ref-101':15 '6987':7 'adira':12 'auto':9 'discret':10 'e':6 'finance':11 'fu':5 'jatibarang':13 'order':18 'r2':2 'satria':4 'sd':8 'suzuk':3 'syaenudin':14 'trx':17 'trx-order':16
293	100	trx-order	Piutang jasa Mandiri Tunas Finance (MTF) Order SPK: /000000143	2022-03-28	Kendaraan R4 Mitsubishi Box , Nopol E 9844 HB	'/000000143':1 '/ref-100':16 '9844':6 'box':4 'deddy':14 'e':5 'finance':10 'hb':7 'indramayu':13 'mandir':8 'mitsubish':3 'mtf':11 'order':19 'pranoto':15 'pusat':12 'r4':2 'trx':18 'trx-order':17 'tunas':9
\.


--
-- Data for Name: trx_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trx_detail (id, code_id, trx_id, debt, cred) FROM stdin;
1	5511	207	1600000.00	0.00
2	1113	207	0.00	1600000.00
1	5511	208	1040000.00	0.00
2	1113	208	0.00	1040000.00
1	5511	209	720000.00	0.00
2	1113	209	0.00	720000.00
1	5511	210	920000.00	0.00
2	1113	210	0.00	920000.00
1	5511	211	3200000.00	0.00
2	1113	211	0.00	3200000.00
1	5511	212	800000.00	0.00
2	1113	212	0.00	800000.00
1	5511	213	8200000.00	0.00
2	1113	213	0.00	8200000.00
1	5511	214	17000000.00	0.00
2	1113	214	0.00	17000000.00
1	5511	215	800000.00	0.00
2	1113	215	0.00	800000.00
1	5511	216	1040000.00	0.00
2	1113	216	0.00	1040000.00
1	5511	217	1200000.00	0.00
2	1113	217	0.00	1200000.00
1	5511	218	1600000.00	0.00
2	1113	218	0.00	1600000.00
1	5511	219	960000.00	0.00
2	1113	219	0.00	960000.00
1	5511	220	1120000.00	0.00
2	1113	220	0.00	1120000.00
1	5511	221	1000000.00	0.00
2	1113	221	0.00	1000000.00
1	5511	222	1040000.00	0.00
2	1113	222	0.00	1040000.00
1	5511	223	31360000.00	0.00
2	1113	223	0.00	31360000.00
1	5511	224	23000000.00	0.00
2	1113	224	0.00	23000000.00
1	5511	225	13950000.00	0.00
2	1113	225	0.00	13950000.00
1	5511	226	19500000.00	0.00
2	1113	226	0.00	19500000.00
1	5511	227	1440000.00	0.00
2	1113	227	0.00	1440000.00
1	5511	228	5400000.00	0.00
2	1113	228	0.00	5400000.00
1	5511	229	12000000.00	0.00
2	1113	229	0.00	12000000.00
1	5511	230	1200000.00	0.00
2	1113	230	0.00	1200000.00
1	5511	231	1000000.00	0.00
2	1113	231	0.00	1000000.00
1	5511	232	1120000.00	0.00
2	1113	232	0.00	1120000.00
1	5511	233	900000.00	0.00
2	1113	233	0.00	900000.00
1	5511	234	1040000.00	0.00
2	1113	234	0.00	1040000.00
1	5511	235	800000.00	0.00
2	1113	235	0.00	800000.00
1	5511	236	600000.00	0.00
2	1113	236	0.00	600000.00
1	5511	237	24200000.00	0.00
2	1113	237	0.00	24200000.00
1	5511	238	24000000.00	0.00
2	1113	238	0.00	24000000.00
1	5511	239	38400000.00	0.00
2	1113	239	0.00	38400000.00
1	5511	240	800000.00	0.00
2	1113	240	0.00	800000.00
1	5511	241	2880000.00	0.00
2	1113	241	0.00	2880000.00
1	5511	242	720000.00	0.00
2	1113	242	0.00	720000.00
1	5511	243	20000000.00	0.00
2	1113	243	0.00	20000000.00
1	5511	244	8000000.00	0.00
2	1113	244	0.00	8000000.00
1	5511	245	24200000.00	0.00
2	1113	245	0.00	24200000.00
1	5511	246	800000.00	0.00
2	1113	246	0.00	800000.00
1	5511	247	20000000.00	0.00
2	1113	247	0.00	20000000.00
1	5511	248	19750000.00	0.00
2	1113	248	0.00	19750000.00
1	5511	249	1040000.00	0.00
2	1113	249	0.00	1040000.00
1	5511	250	1280000.00	0.00
2	1113	250	0.00	1280000.00
1	5511	251	13200000.00	0.00
2	1113	251	0.00	13200000.00
1	5511	252	960000.00	0.00
2	1113	252	0.00	960000.00
1	5511	253	760000.00	0.00
2	1113	253	0.00	760000.00
1	5511	254	760000.00	0.00
2	1113	254	0.00	760000.00
1	5511	255	1200000.00	0.00
2	1113	255	0.00	1200000.00
1	5511	256	1200000.00	0.00
2	1113	256	0.00	1200000.00
1	5511	257	960000.00	0.00
2	1113	257	0.00	960000.00
1	5511	258	680000.00	0.00
2	1113	258	0.00	680000.00
1	5511	259	540000.00	0.00
2	1113	259	0.00	540000.00
1	5511	260	560000.00	0.00
2	1113	260	0.00	560000.00
1	5511	261	1040000.00	0.00
2	1113	261	0.00	1040000.00
1	5511	262	1160000.00	0.00
2	1113	262	0.00	1160000.00
1	5511	263	1160000.00	0.00
2	1113	263	0.00	1160000.00
1	5511	264	1440000.00	0.00
2	1113	264	0.00	1440000.00
1	5511	265	960000.00	0.00
2	1113	265	0.00	960000.00
1	5511	266	1360000.00	0.00
2	1113	266	0.00	1360000.00
1	5511	267	1200000.00	0.00
2	1113	267	0.00	1200000.00
1	5511	268	720000.00	0.00
2	1113	268	0.00	720000.00
1	5511	269	1440000.00	0.00
2	1113	269	0.00	1440000.00
1	5511	270	1200000.00	0.00
2	1113	270	0.00	1200000.00
1	5511	271	680000.00	0.00
2	1113	271	0.00	680000.00
1	5511	272	1200000.00	0.00
2	1113	272	0.00	1200000.00
1	5511	273	960000.00	0.00
2	1113	273	0.00	960000.00
1	5511	274	1120000.00	0.00
2	1113	274	0.00	1120000.00
1	5511	275	960000.00	0.00
2	1113	275	0.00	960000.00
1	5511	276	640000.00	0.00
2	1113	276	0.00	640000.00
1	5511	277	12300000.00	0.00
2	1113	277	0.00	12300000.00
1	5511	278	15300000.00	0.00
2	1113	278	0.00	15300000.00
1	5511	279	960000.00	0.00
2	1113	279	0.00	960000.00
1	5511	280	1280000.00	0.00
2	1113	280	0.00	1280000.00
1	5511	281	840000.00	0.00
2	1113	281	0.00	840000.00
1	5511	282	1280000.00	0.00
2	1113	282	0.00	1280000.00
1	5511	283	1120000.00	0.00
2	1113	283	0.00	1120000.00
1	5511	284	800000.00	0.00
2	1113	284	0.00	800000.00
1	5511	285	1000000.00	0.00
2	1113	285	0.00	1000000.00
1	5511	286	1040000.00	0.00
2	1113	286	0.00	1040000.00
1	5511	288	1000000.00	0.00
2	1113	288	0.00	1000000.00
1	5511	289	1360000.00	0.00
2	1113	289	0.00	1360000.00
1	5511	287	1160000.00	0.00
2	1113	287	0.00	1160000.00
1	5511	291	8500000.00	0.00
2	1113	291	0.00	8500000.00
1	5511	292	13000000.00	0.00
2	1113	292	0.00	13000000.00
1	5511	290	800000.00	0.00
2	1113	290	0.00	800000.00
1	5511	293	25000000.00	0.00
2	1113	293	0.00	25000000.00
1	5511	294	1280000.00	0.00
2	1113	294	0.00	1280000.00
\.


--
-- Data for Name: trx_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trx_type (id, name, descriptions) FROM stdin;
1	Pendapatan	Arus kas dari aktivitas operasi seperti tagihan / invoice 
2	Pengeluaran	Arus kas karena adanya operasional, gaji karyawan, biaya tetap
3	Financing	Arus kas aktivitas pendanaan, segala aktivitas kas yg mempengaruhi posisi modal dan piutang pelanggan
\.


--
-- Data for Name: types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.types (id, name, wheel_id, merk_id) FROM stdin;
2	Vario 125	2	13
3	Brio 1000	3	13
1	Fino 125 	2	2
4	Avanza	3	15
5	GENIO	2	13
6	BEAT	2	13
7	Mio Z	2	2
8	Mio	2	2
9	Mio M3	2	2
10	Vixion	2	2
11	R-15	2	2
12	Jupiter	2	2
13	Satria FU	2	12
14	Jupiter MX	2	2
15	Supra X-125	2	13
16	Revo	2	13
17	Gear	2	2
18	Mio S	2	2
19	Vario 150	2	13
20	Xeon	2	2
21	Verza	2	13
22	Jazz	3	13
23	Carry	3	12
25	Mobilio	3	13
26	Pickup	3	1
27	XL-7	3	12
28	Expander	3	1
29	Terios	3	15
30	Xenia	3	16
31	Vios	3	15
24	ERTIGA	3	12
32	BRV	3	13
33	AGYA	3	15
34	Grand Max	3	16
35	NMax	2	2
36	PCX	2	13
37	Beat Pop	2	13
38	MX-King	2	2
40	cb150	2	13
41	scoopy	2	13
42	sonic	2	13
43	AEROX	2	2
44	Box	3	1
\.


--
-- Data for Name: units; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.units (order_id, nopol, year, frame_number, machine_number, color, type_id, warehouse_id) FROM stdin;
1	E 5605 PAS	2017	\N	\N	\N	7	1
2	E 3977 PAC	2016	\N	\N	\N	8	1
3	E 5125 PBC	2018	\N	\N	\N	9	1
4	BR 5080 PY	2015	\N	\N	\N	10	1
5	E 4096 PAQ	2017	\N	\N	\N	1	1
6	E 2633 PAC	2016	\N	\N	\N	6	1
7	E 2033 PBJ	2019	\N	\N	\N	1	1
10	E 5474 Q	2013	\N	\N	\N	6	1
11	E 4892 TK	2012	\N	\N	\N	12	3
12	B 3479 UJU	2016	\N	\N	\N	9	3
13	E 2417 PAO	2017	\N	\N	\N	9	3
14	T 3521 KL	2012	\N	\N	\N	13	3
15	T 4146 KO	2012	\N	\N	\N	14	3
16	E 3848 UB	2016	\N	\N	\N	6	3
17	T 3828 FW	2008	\N	\N	\N	15	1
18	E 6716 IX	2015	\N	\N	\N	9	3
19	E 5638 PAV	2018	\N	\N	\N	16	2
20	E 5253 TY	2013	\N	\N	\N	2	3
21	B 6262 VKY	2015	\N	\N	\N	9	3
22	E 2830 QR	2015	\N	\N	\N	9	3
24	E 2146 QAF	2018	\N	\N	\N	18	3
26	T 2891 WP	2014	\N	\N	\N	19	2
27	B 6819 PZI	2014	\N	\N	\N	20	2
28	T 4487 PJ	2017	\N	\N	\N	7	2
29	E 4544 JD	2015	\N	\N	\N	8	3
30	T 3615 ZD	2018	\N	\N	\N	21	2
31	E 2391 JM	2017	\N	\N	\N	11	2
32	T 2191 YS	2017	\N	\N	\N	6	2
33	T 5856 ZT	2019	\N	\N	\N	5	2
34	E 4593 TQ	2012	\N	\N	\N	12	2
35	T 2110 YV	2017	\N	\N	\N	11	2
36	E 5713 PAV	2018	\N	\N	\N	9	2
37	B 8936 NO	2006	\N	\N	\N	22	1
38	T 1412 KM	2004	\N	\N	\N	23	1
39	H 8715 GP	2016	\N	\N	\N	3	1
40	T 1164 FQ	2017	\N	\N	\N	24	1
41	T 1788 BC	2017	\N	\N	\N	25	1
42	H 9049 SE	2020	\N	\N	\N	3	1
49	H 8630 PP	2012	\N	\N	\N	22	1
50	H 9442 NG	2017	\N	\N	\N	3	1
51	T 1250 DU	2013	\N	\N	\N	29	1
52	T 1184 GA	2018	\N	\N	\N	3	1
53	D 1312 WF	2007	\N	\N	\N	30	1
54	D 1242 OU	2012	\N	\N	\N	31	1
55	T 1729 BF	2018	\N	\N	\N	32	1
56	H 9086 TE	2018	\N	\N	\N	33	1
57	E 1256 QD	2019	\N	\N	\N	34	1
58	B 2281 SBT	2012	\N	\N	\N	29	1
59	T 1305 DL	2000	\N	\N	\N	24	1
60	F 6181 FCH	2018	\N	\N	\N	6	1
61	E 3310 QR	2015	\N	\N	\N	10	1
62	E 3217 PAR	2017	\N	\N	\N	35	1
63	E 2113 PBM	2019	\N	\N	\N	36	1
64	T 5097 ZB	2018	\N	\N	\N	19	2
65	E 2676 UX	2019	\N	\N	\N	35	2
66	T 6934 YQ	2017	\N	\N	\N	6	2
67	T 4845 IQ	2020	\N	\N	\N	35	2
68	B 3256  PWY	2013	\N	\N	\N	6	2
69	B 3433 USN	2019	\N	\N	\N	6	2
70	B 4654 FSG	2018	\N	\N	\N	6	2
71	T 4261 YQ	2017	\N	\N	\N	38	2
72	B 3430 EJX	2016	\N	\N	\N	6	2
73	B 3351 KUH	2015	\N	\N	\N	37	2
74	E 6871 CM	2018	\N	\N	\N	8	2
75	E 4080 UO	2018	\N	\N	\N	9	2
77	E 5826 CQ	2019	\N	\N	\N	2	2
78	E 6505 QR	2019	\N	\N	\N	6	2
79	E 2282 PBC	2018	\N	\N	\N	6	2
80	E 5737 PBO	2019	\N	\N	\N	1	2
81	E 2764 XL	2014	\N	\N	\N	40	2
82	E 5117 PBA	2018	\N	\N	\N	41	2
83	T 3802 ZR	2019	\N	\N	\N	41	2
84	E 5462 QM	2014	\N	\N	\N	6	2
85	G 2867 AUF	2019	\N	\N	\N	9	2
87	E 4857 PAG	2016	\N	\N	\N	6	2
88	T 4454 IH	2018	\N	\N	\N	42	2
89	E 3561 PBI	2019	\N	\N	\N	6	2
90	E 5948 LZ	2011	\N	\N	\N	20	2
93	E 4560 QV	2015	\N	\N	\N	6	2
94	E 2753 PBA	2018	\N	\N	\N	1	1
95	E 2000 XX	2021	\N	\N	\N	35	1
96	B 1301 UZR	2022	\N	\N	\N	3	1
97	B 1942 EVK	2005	\N	\N	\N	22	1
98	E 8072 QB	2022	\N	\N	\N	34	1
99	T 8066 EG	2020	\N	\N	\N	23	1
100	E 9844 HB	2016	\N	\N	\N	44	1
101	E 6987 SD	2013	\N	\N	\N	13	2
102	E 6645 PAV	2018	\N	\N	\N	1	2
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
2	Jatibarang	Jatibarang
1	Pusat	Indramayu
4	KURANG 300	Sementara
3	GUDANG	Sementara
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

SELECT pg_catalog.setval('public.action_id_seq', 11, true);


--
-- Name: branch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.branch_id_seq', 5, true);


--
-- Name: finance_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finance_groups_id_seq', 7, true);


--
-- Name: finance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finance_id_seq', 22, true);


--
-- Name: invoices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.invoices_id_seq', 10, true);


--
-- Name: lents_id_sequence; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lents_id_sequence', 1, false);


--
-- Name: loans_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.loans_id_seq', 1, true);


--
-- Name: merk_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.merk_id_seq', 17, true);


--
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_id_seq', 102, true);


--
-- Name: order_name_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_name_seq', 152, true);


--
-- Name: trx_detail_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_detail_seq', 1, false);


--
-- Name: trx_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_seq', 295, true);


--
-- Name: type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.type_id_seq', 44, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: warehouse_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.warehouse_id_seq', 6, true);


--
-- Name: wheel_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wheel_id_seq', 6, true);


--
-- Name: acc_code acc_code_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.acc_code
    ADD CONSTRAINT acc_code_name_key UNIQUE (name);


--
-- Name: acc_code acc_code_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.acc_code
    ADD CONSTRAINT acc_code_pkey PRIMARY KEY (id);


--
-- Name: acc_group acc_group_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.acc_group
    ADD CONSTRAINT acc_group_name_key UNIQUE (name);


--
-- Name: acc_group acc_group_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.acc_group
    ADD CONSTRAINT acc_group_pkey PRIMARY KEY (id);


--
-- Name: acc_type acc_type_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.acc_type
    ADD CONSTRAINT acc_type_name_key UNIQUE (name);


--
-- Name: acc_type acc_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.acc_type
    ADD CONSTRAINT acc_type_pkey PRIMARY KEY (id);


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
-- Name: finance_groups finance_groups_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.finance_groups
    ADD CONSTRAINT finance_groups_name_key UNIQUE (name);


--
-- Name: finance_groups finance_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.finance_groups
    ADD CONSTRAINT finance_groups_pkey PRIMARY KEY (id);


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
-- Name: invoice_details invoice_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_details
    ADD CONSTRAINT invoice_details_pkey PRIMARY KEY (invoice_id, order_id);


--
-- Name: invoices invoices_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT invoices_pkey PRIMARY KEY (id);


--
-- Name: ktp_addresses ktp_addresses_order_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ktp_addresses
    ADD CONSTRAINT ktp_addresses_order_id_key UNIQUE (order_id);


--
-- Name: lents lents_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lents
    ADD CONSTRAINT lents_pkey PRIMARY KEY (order_id);


--
-- Name: loans loans_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.loans
    ADD CONSTRAINT loans_pkey PRIMARY KEY (id);


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
-- Name: trx_detail trx_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_detail
    ADD CONSTRAINT trx_detail_pkey PRIMARY KEY (trx_id, id);


--
-- Name: trx trx_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx
    ADD CONSTRAINT trx_pkey PRIMARY KEY (id);


--
-- Name: trx_type trx_type_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_type
    ADD CONSTRAINT trx_type_name_key UNIQUE (name);


--
-- Name: trx_type trx_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_type
    ADD CONSTRAINT trx_type_pkey PRIMARY KEY (id);


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
-- Name: acc_code_token; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX acc_code_token ON public.acc_code USING gin (token_name);


--
-- Name: acc_code_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX acc_code_type ON public.acc_code USING btree (type_id);


--
-- Name: gx_invoices_token; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX gx_invoices_token ON public.invoices USING gin (token);


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
-- Name: ix_acc_type_group; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_acc_type_group ON public.acc_type USING btree (group_id);


--
-- Name: ix_finance_group; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_finance_group ON public.finances USING btree (group_id);


--
-- Name: ix_gin_trx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_gin_trx ON public.trx USING gin (trx_token);


--
-- Name: ix_invoice_account; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_invoice_account ON public.invoices USING btree (account_id);


--
-- Name: ix_invoice_detail_invoice; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_invoice_detail_invoice ON public.invoice_details USING btree (invoice_id);


--
-- Name: ix_invoice_detail_order; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_invoice_detail_order ON public.invoice_details USING btree (order_id);


--
-- Name: ix_invoice_finance; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_invoice_finance ON public.invoices USING btree (finance_id);


--
-- Name: ix_lent_order; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_lent_order ON public.lents USING btree (order_id);


--
-- Name: ix_lents_serial; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_lents_serial ON public.lents USING btree (serial_num);


--
-- Name: ix_loans_serial; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_loans_serial ON public.loans USING btree (serial_num);


--
-- Name: ix_order_token; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_order_token ON public.orders USING gin (token);


--
-- Name: ix_trx_detail_acc_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_trx_detail_acc_code ON public.trx_detail USING btree (code_id);


--
-- Name: ix_trx_detail_trx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_trx_detail_trx ON public.trx_detail USING btree (trx_id);


--
-- Name: acc_code acc_code_acc_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.acc_code
    ADD CONSTRAINT acc_code_acc_type_id_fkey FOREIGN KEY (type_id) REFERENCES public.acc_type(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: acc_type acc_type_group_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.acc_type
    ADD CONSTRAINT acc_type_group_fkey FOREIGN KEY (group_id) REFERENCES public.acc_group(id);


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
-- Name: finances fkey_finance_group; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.finances
    ADD CONSTRAINT fkey_finance_group FOREIGN KEY (group_id) REFERENCES public.finance_groups(id);


--
-- Name: invoice_details fkey_invdetail_invoice; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_details
    ADD CONSTRAINT fkey_invdetail_invoice FOREIGN KEY (invoice_id) REFERENCES public.invoices(id) ON DELETE CASCADE;


--
-- Name: invoice_details fkey_invdetail_order; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoice_details
    ADD CONSTRAINT fkey_invdetail_order FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- Name: invoices fkey_invoice_account; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT fkey_invoice_account FOREIGN KEY (account_id) REFERENCES public.acc_code(id);


--
-- Name: invoices fkey_invoice_finance; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.invoices
    ADD CONSTRAINT fkey_invoice_finance FOREIGN KEY (finance_id) REFERENCES public.finances(id);


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
-- Name: lents lents_order_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lents
    ADD CONSTRAINT lents_order_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id);


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
-- Name: trx_detail trx_detail_acc_code_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_detail
    ADD CONSTRAINT trx_detail_acc_code_id_fkey FOREIGN KEY (code_id) REFERENCES public.acc_code(id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: trx_detail trx_detail_trx_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trx_detail
    ADD CONSTRAINT trx_detail_trx_id_fkey FOREIGN KEY (trx_id) REFERENCES public.trx(id) ON UPDATE CASCADE ON DELETE CASCADE;


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

