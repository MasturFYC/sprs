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
11	Kas	Kelompok akun yg berfungsi mencatat perubahan uang seperti penerimaan atau pengeluaran. termasuk akun kas, seperti cek, giro.	1
\.


--
-- Data for Name: actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.actions (id, action_at, pic, descriptions, order_id, file_name) FROM stdin;
10	2022-03-25	Test JPG / JPEG / PNG	Test upload and download image	10	d9f143ce.jpg
11	2022-03-25	Test PDF	Test upload and download pdf	10	d855835b.pdf
12	2022-04-09	stnk	fidusia	130	86f1bfc6.jpg
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
8	BIMA
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
23	Bima Finance	BIMA	\N	\N	\N	\N	\N	\N	8
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
11	129
\.


--
-- Data for Name: invoices; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.invoices (id, invoice_at, payment_term, due_at, salesman, finance_id, account_id, subtotal, ppn, tax, total, memo, token) FROM stdin;
11	2022-04-06	1	2022-04-06	Gondrong	2	1113	1400000.00	0.00	0.00	1400000.00	\N	'/id-0':2 '4630':8 'adira':6 'auto':3 'beat':10 'discret':4 'e':7 'finance':5 'gondrong':1 'pax':9
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
42	Norman Brio	\N	\N	\N	\N	\N	6
50	Norman Brio 1	\N	\N	\N	\N	\N	4
56	Norman Agya	\N	\N	\N	\N	\N	5
51	Bahrudin Terios DU	\N	\N	\N	\N	\N	7
53	Bahrudin Xenia	\N	\N	\N	\N	\N	8
55	Bahrudin BRV	\N	\N	\N	\N	\N	9
59	Bahrudin Ertiga	\N	\N	\N	\N	\N	10
52	Samuel Nus	\N	\N	\N	\N	\N	11
58	Ferdinan	\N	\N	\N	\N	\N	12
49	BFI JAZZ	\N	\N	\N	\N	\N	13
54	Adira Vios	\N	\N	\N	\N	\N	14
57	Adira Grandmax	\N	\N	\N	\N	\N	15
96	Sofyan/ Bang Kei	\N	\N	\N	\N	\N	16
124	Pa gineng	\N	\N	\N	\N	\N	23
65	Hilang	\N	\N	\N	\N	\N	34
66	Mastukin	\N	\N	\N	\N	\N	35
69	Wa Jileng (anak)	\N	\N	\N	\N	\N	37
71	Andre	\N	\N	\N	\N	\N	38
72	Tanggung Jawab Cabang (JTB)	\N	\N	\N	\N	\N	39
73	Tanggung Jawab Cabang (JTB) 2	\N	\N	\N	\N	\N	40
75	Eksternal Ali 	\N	\N	\N	\N	\N	42
18	Eksternal Bule	\N	\N	\N	\N	\N	43
67	Wa Jileng	\N	\N	\N	\N	\N	36
84	Eksternal Jatibarang	\N	\N	\N	\N	\N	44
60	Hilang 	\N	\N	\N	\N	\N	45
61	Opick Admin	\N	\N	\N	\N	\N	46
62	Tim Om Saja	\N	\N	\N	\N	\N	47
121	Opick Gudang	\N	\N	\N	\N	\N	48
133	Clip Cikarang	\N	\N	\N	\N	\N	50
\.


--
-- Data for Name: loans; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.loans (id, name, street, city, phone, cell, zip, persen, serial_num) FROM stdin;
4	Abdul sholeh	\N	\N	\N	\N	\N	0.00	18
12	An Arsim	\N	\N	\N	\N	\N	0.00	27
15	pak ipeng	\N	\N	\N	\N	\N	0.00	30
16	PA Agus BAF	\N	\N	\N	\N	\N	0.00	31
18	PELUS AN CARDI	\N	\N	\N	\N	\N	5.00	33
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
133	000000270	2022-04-26	2022-04-26	55000000.00	10.79	49066000.00	Opick	test	19	3	t	0.00	55000000.00	'-1293':17 '-2022':4 '000000270':1 '2014':24 '26':2 '4':27 'ada':15 'apr':3 'cabang':8 'clip':7 'clipan':6 'daihatsu':25 'deddy':11 'finance':5 'grand':19 'gudang':21 'indramayu':10 'max':20 'pranoto':12 'pusat':9,22 'r4':28 'roda':26 'stnk':14 'stnk-ada':13 't':16 'tahun':23 'tn':18
77	000000103	2022-03-19	2022-03-26	1500000.00	20.00	1200000.00	Opick	\N	21	1	t	0.00	1500000.00	'-2022':4 '-5826':18 '000000103':1 '125':21 '19':2 '2':28 '2019':25 'ada':16 'auto':7 'cabang':11 'central':8 'cq':19 'e':17 'finance':5,9 'gudang':22 'honda':26 'jatibarang':12,23 'macf':10 'mar':3 'mega':6 'r2':29 'roda':27 'stnk':15 'stnk-ada':14 'syaenudin':13 'tahun':24 'vario':20
42	000000042	2022-02-16	2022-02-16	15000000.00	12.00	13200000.00	Mastur	\N	18	3	t	0.00	15000000.00	'-2022':4 '-9049':21 '000000042':1 '1000':24 '16':2 '2020':28 '4':31 'ada':19 'brio':23 'cabang':12 'deddy':15 'feb':3 'finance':5,8 'gudang':25 'h':20 'honda':29 'indramayu':14 'mandir':6 'mtf':10 'pranoto':16 'pusat':13,26 'r4':32 'roda':30 's':11 'se':22 'semarang':9 'stnk':18 'stnk-ada':17 'tahun':27 'tunas':7
83	000000109	2022-03-22	2022-03-26	1600000.00	20.00	1280000.00	Opick	\N	2	1	t	0.00	1600000.00	'-2022':4 '-3802':17 '000000109':1 '2':26 '2019':23 '22':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'r2':27 'roda':25 'scoopy':19 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':22 'zr':18
21	000000021	2022-02-21	2022-02-21	950000.00	20.00	760000.00	Mastur	\N	1	1	t	0.00	950000.00	'-2022':4 '-6262':17 '000000021':1 '2':27 '2015':24 '21':2 'ada':15 'auto':7 'b':16 'baf':9 'bussan':6 'cabang':10 'feb':3 'finance':5,8 'gudang':21,22 'jatibarang':11 'm3':20 'mio':19 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'vky':18 'yamaha':25
24	000000024	2022-02-25	2022-02-25	1500000.00	20.00	1200000.00	Mastur	\N	1	1	t	0.00	1500000.00	'-2022':4 '-2146':17 '000000024':1 '2':27 '2018':24 '25':2 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'feb':3 'finance':5,8 'gudang':21,22 'jatibarang':11 'mio':19 'qaf':18 'r2':28 'roda':26 's':20 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
65	000000065	2021-10-05	2021-10-05	3200000.00	0.00	3200000.00	Mastur	\N	1	1	t	0.00	3200000.00	'-2021':4 '-2676':17 '000000065':1 '05':2 '2':26 '2019':23 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'finance':5,8 'gudang':20 'jatibarang':11,21 'nmax':19 'okt':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'ux':18 'yamaha':24
4	000000004	2022-02-18	2022-02-18	1200000.00	20.00	960000.00	Mastur	\N	2	3	t	0.00	1200000.00	'-2022':4 '-5080':19 '000000004':1 '18':2 '2':28 '2015':25 'ada':17 'adira':9 'auto':6 'br':18 'cabang':10 'deddy':13 'discret':7 'feb':3 'finance':5,8 'gudang':22 'indramayu':12 'pranoto':14 'pusat':11,23 'py':20 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'vixion':21 'yamaha':26
6	000000006	2022-03-02	2022-03-02	1200000.00	20.00	960000.00	Mastur	\N	2	3	t	0.00	1200000.00	'-2022':4 '-2633':19 '000000006':1 '02':2 '2':28 '2016':25 'ada':17 'adira':9 'auto':6 'beat':21 'cabang':10 'deddy':13 'discret':7 'e':18 'finance':5,8 'gudang':22 'honda':26 'indramayu':12 'mar':3 'pac':20 'pranoto':14 'pusat':11,23 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24
10	000000010	2022-03-19	2022-03-19	850000.00	20.00	680000.00	Mastur	\N	2	3	t	0.00	850000.00	'-2022':4 '-5474':19 '000000010':1 '19':2 '2':28 '2013':25 'ada':17 'adira':9 'auto':6 'beat':21 'cabang':10 'deddy':13 'discret':7 'e':18 'finance':5,8 'gudang':22 'honda':26 'indramayu':12 'mar':3 'pranoto':14 'pusat':11,23 'q':20 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24
99	000000139	2022-03-28	2022-03-28	16000000.00	18.75	13000000.00	Opick	\N	16	3	t	0.00	16000000.00	'-2022':4 '-8066':19 '000000139':1 '2020':25 '28':2 '4':28 'ada':17 'cabang':10 'carry':21 'deddy':13 'eg':20 'finance':5,7 'gudang':22 'indonesia':8 'indramayu':12 'mar':3 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'sfi':9 'stnk':16 'stnk-ada':15 'suzuk':6,26 't':18 'tahun':24
123	000000209	2022-03-31	2022-03-31	11500000.00	10.43	10300000.00	Opick	\N	14	3	t	0.00	11500000.00	'-1561':19 '-2022':4 '000000209':1 '2007':25 '31':2 '4':28 'ada':17 'b':18 'cabang':10 'clip':8 'clipan':6 'deddy':13 'finance':5 'gudang':22 'indramayu':12 'k':9 'karawang':7 'kbj':20 'mar':3 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'toyota':26 'vios':21
57	000000057	2022-01-22	2022-01-22	8000000.00	0.00	8000000.00	Mastur	\N	2	3	t	0.00	8000000.00	'-1256':19 '-2022':4 '000000057':1 '2019':26 '22':2 '4':29 'ada':17 'adira':9 'auto':6 'cabang':10 'daihatsu':27 'deddy':13 'discret':7 'e':18 'finance':5,8 'grand':21 'gudang':23 'indramayu':12 'jan':3 'max':22 'pranoto':14 'pusat':11,24 'qd':20 'r4':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25
18	000000018	2022-01-18	2022-01-18	1000000.00	20.00	800000.00	Mastur	\N	5	1	t	0.00	1000000.00	'-2022':4 '-6716':17 '000000018':1 '18':2 '2':27 '2015':24 'ada':15 'cabang':10 'e':16 'finance':5 'gudang':21,22 'ix':18 'jan':3 'jatibarang':11 'kredit':7 'm3':20 'mio':19 'motor':8 'oto':6 'otto':9 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
101	000000146	2022-03-26	2022-03-28	1000000.00	20.00	800000.00	Opick	\N	2	1	t	0.00	1000000.00	'-2022':4 '-6987':17 '000000146':1 '2':27 '2013':24 '26':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'fu':20 'gudang':21 'jatibarang':11,22 'mar':3 'r2':28 'roda':26 'satria':19 'sd':18 'stnk':14 'stnk-ada':13 'suzuk':25 'syaenudin':12 'tahun':23
107	000000179	2022-03-30	2022-03-30	1000000.00	20.00	800000.00	Opick	\N	12	1	t	0.00	1000000.00	'-2022':4 '-5359':16 '000000179':1 '125':19 '2':26 '2016':23 '30':2 'ada':14 'cabang':9 'e':15 'finance':5,7 'fino':18 'gudang':20 'jatibarang':10,21 'mar':3 'pai':17 'r2':27 'roda':25 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':22 'wom':6 'womf':8 'yamaha':24
114	000000186	2022-04-01	2022-04-01	1300000.00	20.00	1040000.00	Opick	\N	7	1	t	0.00	1300000.00	'-2022':4 '-4418':17 '000000186':1 '01':2 '125':20 '2':27 '2017':24 'ada':15 'apr':3 'cabang':10 'finance':5,8 'fino':19 'gudang':21 'jatibarang':11,22 'mandir':6 'muf':9 'pf':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':23 'utama':7 'yamaha':25
115	000000187	2022-04-01	2022-04-01	1200000.00	20.00	960000.00	Opick	\N	16	1	t	0.00	1200000.00	'-2022':4 '-4080':17 '000000187':1 '01':2 '2':26 '2013':23 'ada':15 'apr':3 'cabang':10 'finance':5,7 'gudang':20 'indonesia':8 'jatibarang':11,21 'nex':19 'r2':27 'roda':25 'sfi':9 'stnk':14 'stnk-ada':13 'suzuk':6,24 'syaenudin':12 't':16 'tahun':22 'we':18
62	000000062	2022-01-19	2022-01-19	3600000.00	20.00	2880000.00	Mastur	\N	1	3	t	0.00	3600000.00	'-2022':4 '-3217':19 '000000062':1 '19':2 '2':28 '2017':25 'ada':17 'auto':7 'baf':9 'bussan':6 'cabang':10 'deddy':13 'e':18 'finance':5,8 'gudang':22 'indramayu':12 'jan':3 'nmax':21 'par':20 'pranoto':14 'pusat':11,23 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'yamaha':26
40	000000040	2022-01-17	2022-01-17	26000000.00	7.69	24000000.00	Mastur	\N	14	3	t	0.00	26000000.00	'-1164':19 '-2022':4 '000000040':1 '17':2 '2017':25 '4':28 'ada':17 'cabang':10 'clip':8 'clipan':6 'deddy':13 'ertiga':21 'finance':5 'fq':20 'gudang':22 'indramayu':12 'jan':3 'k':9 'karawang':7 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'suzuk':26 't':18 'tahun':24
41	000000041	2022-01-25	2022-01-25	26620000.00	9.09	24200000.00	Mastur	\N	14	3	t	0.00	26620000.00	'-1788':19 '-2022':4 '000000041':1 '2017':25 '25':2 '4':28 'ada':17 'bc':20 'cabang':10 'clip':8 'clipan':6 'deddy':13 'finance':5 'gudang':22 'honda':26 'indramayu':12 'jan':3 'k':9 'karawang':7 'mobilio':21 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 't':18 'tahun':24
121	000000199	2022-03-30	2022-03-30	1600000.00	20.00	1280000.00	Opick	\N	7	3	t	0.00	1600000.00	'-2022':4 '-4172':19 '000000199':1 '2':30 '2022':27 '30':2 'ada':17 'b':18 'cabang':10 'deddy':13 'finance':5,8 'fpb':20 'gudang':24 'indramayu':12 'mandir':6 'mar':3 'muf':9 'pranoto':14 'pusat':11,25 'r2':31 'ride':23 'roda':29 'stnk':16 'stnk-ada':15 'tahun':26 'utama':7 'x':22 'x-ride':21 'yamaha':28
7	000000007	2022-03-15	2022-03-15	1700000.00	20.00	1360000.00	Mastur	\N	1	3	t	0.00	1700000.00	'-2022':4 '-2033':19 '000000007':1 '125':22 '15':2 '2':29 '2019':26 'ada':17 'auto':7 'baf':9 'bussan':6 'cabang':10 'deddy':13 'e':18 'finance':5,8 'fino':21 'gudang':23 'indramayu':12 'mar':3 'pbj':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'yamaha':27
15	000000015	2022-01-06	2022-01-06	900000.00	20.00	900000.00	Mastur	\N	5	1	t	0.00	900000.00	'-2022':4 '-4146':17 '000000015':1 '06':2 '2':27 '2012':24 'ada':15 'cabang':10 'finance':5 'gudang':21,22 'jan':3 'jatibarang':11 'jupiter':19 'ko':18 'kredit':7 'motor':8 'mx':20 'oto':6 'otto':9 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':23 'yamaha':25
50	000000050	2021-12-16	2021-12-16	34500000.00	9.10	31360000.00	Mastur	\N	18	3	t	0.00	34500000.00	'-2021':4 '-9442':21 '000000050':1 '1000':24 '16':2 '2017':28 '4':31 'ada':19 'brio':23 'cabang':12 'deddy':15 'des':3 'finance':5,8 'gudang':25 'h':20 'honda':29 'indramayu':14 'mandir':6 'mtf':10 'ng':22 'pranoto':16 'pusat':13,26 'r4':32 'roda':30 's':11 'semarang':9 'stnk':18 'stnk-ada':17 'tahun':27 'tunas':7
53	000000053	2021-12-24	2021-12-24	21450000.00	9.09	19500000.00	Mastur	\N	14	3	t	0.00	21450000.00	'-1312':19 '-2021':4 '000000053':1 '2007':25 '24':2 '4':28 'ada':17 'cabang':10 'clip':8 'clipan':6 'd':18 'daihatsu':26 'deddy':13 'des':3 'finance':5 'gudang':22 'indramayu':12 'k':9 'karawang':7 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'wf':20 'xenia':21
49	000000049	2021-11-22	2021-11-22	10000000.00	18.00	8200000.00	Mastur	\N	17	3	t	0.00	10000000.00	'-2021':4 '-8630':18 '000000049':1 '2012':24 '22':2 '4':27 'ada':16 'bfi':6,8 'cabang':9 'deddy':12 'finance':5,7 'gudang':21 'h':17 'honda':25 'indramayu':11 'jazz':20 'nop':3 'pp':19 'pranoto':13 'pusat':10,22 'r4':28 'roda':26 'stnk':15 'stnk-ada':14 'tahun':23
124	000000227	2022-04-02	2022-04-02	465000.00	19.35	375000.00	Opick	\N	2	3	t	0.00	465000.00	'-2000':19 '-2022':4 '000000227':1 '02':2 '2022':25 '4':28 'ada':17 'adira':9 'apr':3 'auto':6 'cabang':10 'deddy':13 'discret':7 'e':18 'finance':5,8 'gudang':22 'indramayu':12 'ios':21 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'toyota':26 'xx':20
66	000000066	2021-12-06	2021-12-06	1300000.00	20.00	1040000.00	Mastur	\N	12	1	t	0.00	1300000.00	'-2021':4 '-6934':16 '000000066':1 '06':2 '2':25 '2017':22 'ada':14 'beat':18 'cabang':9 'des':3 'finance':5,7 'gudang':19 'honda':23 'jatibarang':10,20 'r2':26 'roda':24 'stnk':13 'stnk-ada':12 'syaenudin':11 't':15 'tahun':21 'wom':6 'womf':8 'yq':17
26	000000026	2022-03-07	2022-03-07	850000.00	20.00	680000.00	Mastur	\N	11	1	t	0.00	850000.00	'-2022':4 '-2891':16 '000000026':1 '07':2 '150':19 '2':26 '2014':23 'ada':14 'cabang':9 'finance':5 'gudang':20 'honda':24 'jatibarang':10,21 'kp':8 'kredit':6 'mar':3 'plus':7 'r2':27 'roda':25 'stnk':13 'stnk-ada':12 'syaenudin':11 't':15 'tahun':22 'vario':18 'wp':17
33	000000033	2022-03-16	2022-03-16	1500000.00	20.00	1200000.00	Mastur	\N	2	1	t	0.00	1500000.00	'-2022':4 '-5856':17 '000000033':1 '16':2 '2':26 '2019':23 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'genio':19 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':22 'zt':18
131	000000245	2022-04-06	2022-04-06	1300000.00	20.00	1040000.00	Opick	\N	8	3	f	200000.00	1500000.00	'-2022':4 '-5491':19 '000000245':1 '06':2 '125':22 '2':29 '2022':26 'ada':17 'apr':3 'cabang':9 'deddy':12 'e':18 'fif':6,8 'finance':5 'group':7 'gudang':23 'honda':27 'indramayu':11 'pbl':20 'pranoto':13 'pusat':10,24 'r2':30 'roda':28 'stnk':15 'stnk-tidak-ada':14 'tahun':25 'tidak':16 'vario':21
96	000000132	2022-03-21	2022-03-28	13300000.00	7.52	12300000.00	Opick	\N	19	3	t	0.00	13300000.00	'-1301':17 '-2022':4 '000000132':1 '1000':20 '2022':24 '21':2 '4':27 'ada':15 'b':16 'brio':19 'cabang':8 'clip':7 'clipan':6 'deddy':11 'finance':5 'gudang':21 'honda':25 'indramayu':10 'mar':3 'pranoto':12 'pusat':9,22 'r4':28 'roda':26 'stnk':14 'stnk-ada':13 'tahun':23 'uzr':18
125	000000230	2022-03-29	2022-03-29	4250000.00	29.41	3000000.00	Opick	\N	2	3	t	0.00	4250000.00	'-1411':19 '-2022':4 '000000230':1 '2020':25 '29':2 '4':28 'ada':17 'adira':9 'auto':6 'avanza':21 'cabang':10 'deddy':13 'discret':7 'e':18 'finance':5,8 'gudang':22 'indramayu':12 'mar':3 'pranoto':14 'pusat':11,23 'pv':20 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'toyota':26
71	000000071	2021-09-21	2021-09-21	1300000.00	20.00	1040000.00	Mastur	\N	7	1	t	0.00	1300000.00	'-2021':4 '-4261':17 '000000071':1 '2':28 '2017':25 '21':2 'ada':15 'cabang':10 'finance':5,8 'gudang':22 'jatibarang':11,23 'king':21 'mandir':6 'muf':9 'mx':20 'mx-king':19 'r2':29 'roda':27 'sep':3 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':24 'utama':7 'yamaha':26 'yq':18
69	000000069	2021-12-09	2021-12-09	1400000.00	20.00	1120000.00	Mastur	\N	2	1	t	0.00	1400000.00	'-2021':4 '-3433':17 '000000069':1 '09':2 '2':26 '2019':23 'ada':15 'adira':9 'auto':6 'b':16 'beat':19 'cabang':10 'des':3 'discret':7 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'usn':18
72	000000072	2021-12-09	2021-12-09	1250000.00	20.00	1000000.00	Mastur	\N	2	1	t	0.00	1250000.00	'-2021':4 '-3430':17 '000000072':1 '09':2 '2':26 '2016':23 'ada':15 'adira':9 'auto':6 'b':16 'beat':19 'cabang':10 'des':3 'discret':7 'ejx':18 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
13	000000013	2021-12-06	2021-12-06	1000000.00	20.00	800000.00	Mastur	\N	6	1	t	0.00	1000000.00	'-2021':4 '-2417':15 '000000013':1 '06':2 '2':25 '2017':22 'ada':13 'cabang':8 'col':7 'collectius':6 'des':3 'e':14 'finance':5 'gudang':19,20 'jatibarang':9 'm3':18 'mio':17 'pao':16 'r2':26 'roda':24 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':21 'yamaha':23
20	000000020	2022-01-21	2022-01-21	900000.00	20.00	720000.00	Mastur	\N	11	1	t	0.00	900000.00	'-2022':4 '-5253':16 '000000020':1 '125':19 '2':26 '2013':23 '21':2 'ada':14 'cabang':9 'e':15 'finance':5 'gudang':20,21 'honda':24 'jan':3 'jatibarang':10 'kp':8 'kredit':6 'plus':7 'r2':27 'roda':25 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':22 'ty':17 'vario':18
29	000000029	2022-03-09	2022-03-09	1450000.00	20.00	1160000.00	Mastur	\N	13	1	t	0.00	1450000.00	'-2022':4 '-4544':15 '000000029':1 '09':2 '2':24 '2015':21 'ada':13 'cabang':8 'e':14 'finance':5 'gapara':6 'gudang':18,19 'jatibarang':9 'jd':16 'mar':3 'mio':17 'mpr':7 'r2':25 'roda':23 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':20 'yamaha':22
31	000000031	2022-03-12	2022-03-12	1800000.00	20.00	1440000.00	Mastur	\N	2	1	t	0.00	1800000.00	'-15':20 '-2022':4 '-2391':17 '000000031':1 '12':2 '2':27 '2017':24 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':21 'jatibarang':11,22 'jm':18 'mar':3 'r':19 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
22	000000022	2022-02-23	2022-02-23	950000.00	20.00	760000.00	Mastur	\N	1	1	t	0.00	950000.00	'-2022':4 '-2830':17 '000000022':1 '2':27 '2015':24 '23':2 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'feb':3 'finance':5,8 'gudang':21,22 'jatibarang':11 'm3':20 'mio':19 'qr':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
5	000000005	2022-02-24	2022-02-24	1500000.00	20.00	1200000.00	Mastur	\N	7	3	t	0.00	1500000.00	'-2022':4 '-4096':19 '000000005':1 '125':22 '2':29 '2017':26 '24':2 'ada':17 'cabang':10 'deddy':13 'e':18 'feb':3 'finance':5,8 'fino':21 'gudang':23 'indramayu':12 'mandir':6 'muf':9 'paq':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'utama':7 'yamaha':27
3	000000003	2022-02-07	2022-02-07	1300000.00	20.00	1040000.00	Mastur	\N	1	3	t	0.00	1300000.00	'-2022':4 '-5125':19 '000000003':1 '07':2 '2':29 '2018':26 'ada':17 'auto':7 'baf':9 'bussan':6 'cabang':10 'deddy':13 'e':18 'feb':3 'finance':5,8 'gudang':23 'indramayu':12 'm3':22 'mio':21 'pbc':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'yamaha':27
132	000000265	2022-04-09	2022-04-09	1600000.00	20.00	1280000.00	Opick	\N	5	1	t	0.00	1600000.00	'-2022':4 '-4415':17 '000000265':1 '09':2 '2':26 '2019':23 'ada':15 'apr':3 'cabang':10 'e':16 'finance':5 'gudang':20 'jatibarang':11,21 'jz':18 'kredit':7 'motor':8 'nmax':19 'oto':6 'otto':9 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'yamaha':24
74	000000074	2022-03-07	2022-03-07	1300000.00	20.00	1040000.00	Mastur	\N	2	1	t	0.00	1300000.00	'-2022':4 '-6871':17 '000000074':1 '07':2 '2':26 '2018':23 'ada':15 'adira':9 'auto':6 'cabang':10 'cm':18 'discret':7 'e':16 'finance':5,8 'gudang':20 'jatibarang':11,21 'mar':3 'mio':19 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'yamaha':24
84	000000110	2022-03-22	2022-03-26	1050000.00	20.00	840000.00	Opick	\N	2	1	t	0.00	1050000.00	'-2022':4 '-5462':17 '000000110':1 '2':26 '2014':23 '22':2 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'qm':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
11	000000011	2021-09-27	2021-09-27	900000.00	20.00	720000.00	Mastur	\N	8	1	t	0.00	900000.00	'-2021':4 '-4892':16 '000000011':1 '2':25 '2012':22 '27':2 'ada':14 'cabang':9 'e':15 'fif':6,8 'finance':5 'group':7 'gudang':19,20 'jatibarang':10 'jupiter':18 'r2':26 'roda':24 'sep':3 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':21 'tk':17 'yamaha':23
80	000000106	2022-03-19	2022-03-26	1200000.00	20.00	960000.00	Opick	\N	2	1	t	0.00	1200000.00	'-2022':4 '-5737':17 '000000106':1 '125':20 '19':2 '2':27 '2019':24 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'fino':19 'gudang':21 'jatibarang':11,22 'mar':3 'pbo':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
58	000000058	2022-01-31	2022-01-31	21400000.00	6.54	20000000.00	Mastur	\N	4	3	t	0.00	21400000.00	'-2022':4 '-2281':19 '000000058':1 '2012':25 '31':2 '4':28 'ada':17 'b':9,18 'bekasi':7 'cabang':10 'clip':8 'clipan':6 'deddy':13 'finance':5 'gudang':22 'indramayu':12 'ios':21 'jan':3 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'sbt':20 'stnk':16 'stnk-ada':15 'tahun':24 'toyota':26
59	000000059	2022-01-31	2022-01-31	19750000.00	0.00	19750000.00	Mastur	\N	14	3	t	0.00	19750000.00	'-1305':19 '-2022':4 '000000059':1 '2000':25 '31':2 '4':28 'ada':17 'cabang':10 'clip':8 'clipan':6 'deddy':13 'dl':20 'ertiga':21 'finance':5 'gudang':22 'indramayu':12 'jan':3 'k':9 'karawang':7 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'suzuk':26 't':18 'tahun':24
16	000000016	2022-01-14	2022-01-14	1000000.00	20.00	800000.00	Mastur	\N	5	1	t	0.00	1000000.00	'-2022':4 '-3848':17 '000000016':1 '14':2 '2':26 '2016':23 'ada':15 'beat':19 'cabang':10 'e':16 'finance':5 'gudang':20,21 'honda':24 'jan':3 'jatibarang':11 'kredit':7 'motor':8 'oto':6 'otto':9 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'ub':18
19	000000019	2022-01-26	2022-01-26	1000000.00	20.00	800000.00	Mastur	\N	2	1	t	0.00	1000000.00	'-2022':4 '-5638':17 '000000019':1 '2':26 '2018':23 '26':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jan':3 'jatibarang':11,21 'pav':18 'r2':27 'revo':19 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
79	000000105	2022-03-19	2022-03-26	1400000.00	20.00	1120000.00	Opick	\N	2	1	t	0.00	1400000.00	'-2022':4 '-2282':17 '000000105':1 '19':2 '2':26 '2018':23 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'pbc':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
55	000000055	2022-01-17	2022-01-17	38400000.00	0.00	38400000.00	Mastur	\N	14	3	t	0.00	38400000.00	'-1729':19 '-2022':4 '000000055':1 '17':2 '2018':25 '4':28 'ada':17 'bf':20 'brv':21 'cabang':10 'clip':8 'clipan':6 'deddy':13 'finance':5 'gudang':22 'honda':26 'indramayu':12 'jan':3 'k':9 'karawang':7 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 't':18 'tahun':24
17	000000017	2022-01-14	2022-01-14	750000.00	20.00	600000.00	Mastur	\N	10	1	t	0.00	750000.00	'-125':21 '-2022':4 '-3828':17 '000000017':1 '14':2 '2':28 '2008':25 'ada':15 'cabang':10 'company':8 'finance':5,7 'fw':18 'gudang':22 'honda':26 'jan':3 'jatibarang':11 'pusat':23 'r2':29 'roda':27 'stnk':14 'stnk-ada':13 'supra':19 'syaenudin':12 't':16 'tahun':24 'tfc':9 'top':6 'x':20
32	000000032	2022-03-14	2022-03-14	1200000.00	20.00	960000.00	Mastur	\N	2	1	t	0.00	1200000.00	'-2022':4 '-2191':17 '000000032':1 '14':2 '2':26 '2017':23 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':22 'ys':18
28	000000028	2022-03-09	2022-03-09	1300000.00	20.00	1040000.00	Mastur	\N	2	1	t	0.00	1300000.00	'-2022':4 '-4487':17 '000000028':1 '09':2 '2':27 '2017':24 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'gudang':21 'jatibarang':11,22 'mar':3 'mio':19 'pj':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':23 'yamaha':25 'z':20
30	000000030	2022-03-12	2022-03-12	1450000.00	20.00	1160000.00	Mastur	\N	2	1	t	0.00	1450000.00	'-2022':4 '-3615':17 '000000030':1 '12':2 '2':26 '2018':23 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 'tahun':22 'verza':19 'zd':18
100	000000143	2022-03-28	2022-03-28	30500000.00	11.48	27000000.00	Opick	\N	3	3	t	0.00	30500000.00	'-2022':4 '-9844':19 '000000143':1 '2016':25 '28':2 '4':28 'ada':17 'box':21 'cabang':10 'deddy':13 'e':18 'finance':5,8 'gudang':22 'hb':20 'indramayu':12 'mandir':6 'mar':3 'mitsubish':26 'mtf':9 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'tunas':7
108	000000180	2022-03-31	2022-03-31	1450000.00	20.00	1160000.00	Opick	\N	1	1	t	0.00	1450000.00	'-2022':4 '-5712':17 '000000180':1 '125':20 '2':27 '2019':24 '31':2 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'finance':5,8 'fino':19 'gudang':21 'jatibarang':11,22 'mar':3 'pbg':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
109	000000181	2022-03-31	2022-03-31	1200000.00	20.00	960000.00	Opick	\N	2	1	t	0.00	1200000.00	'-2022':4 '-2575':17 '000000181':1 '2':26 '2017':23 '31':2 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mar':3 'pas':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
119	000000191	2022-04-01	2022-04-01	1200000.00	20.00	960000.00	Opick	\N	2	1	t	0.00	1200000.00	'-2022':4 '-6945':17 '000000191':1 '01':2 '2':26 '2018':23 'ada':15 'adira':9 'apr':3 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'pbc':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
127	000000237	2022-04-02	2022-04-02	1000000.00	20.00	800000.00	Opick	\N	6	1	t	0.00	1000000.00	'-2022':4 '-3907':15 '000000237':1 '02':2 '125':18 '2':25 '2022':22 'a':14 'ada':13 'apr':3 'cabang':8 'col':7 'collectius':6 'finance':5 'fino':17 'gudang':19 'jatibarang':9,20 'pbf':16 'r2':26 'roda':24 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':21 'yamaha':23
38	000000038	2021-12-29	2021-12-29	13000000.00	7.69	12000000.00	Mastur	\N	14	3	t	0.00	13000000.00	'-1412':19 '-2021':4 '000000038':1 '2004':25 '29':2 '4':28 'ada':17 'cabang':10 'carry':21 'clip':8 'clipan':6 'deddy':13 'des':3 'finance':5 'gudang':22 'indramayu':12 'k':9 'karawang':7 'km':20 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'suzuk':26 't':18 'tahun':24
126	000000232	2022-03-29	2022-03-29	400000.00	20.00	320000.00	Opick	\N	2	3	t	0.00	400000.00	'-1456':19 '-2022':4 '000000232':1 '2020':25 '29':2 '4':28 'ada':17 'adira':9 'auto':6 'avanza':21 'cabang':10 'deddy':13 'discret':7 'e':18 'finance':5,8 'gudang':22 'indramayu':12 'mar':3 'pranoto':14 'pusat':11,23 'r4':29 'rl':20 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'toyota':26
67	000000067	2021-12-07	2021-12-07	2000000.00	20.00	1600000.00	Mastur	\N	21	1	t	0.00	2000000.00	'-2021':4 '-4845':18 '000000067':1 '07':2 '2':27 '2020':24 'ada':16 'auto':7 'cabang':11 'central':8 'des':3 'finance':5,9 'gudang':21 'iq':19 'jatibarang':12,22 'macf':10 'mega':6 'nmax':20 'r2':28 'roda':26 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 'tahun':23 'yamaha':25
73	000000073	2021-12-30	2021-12-30	1250000.00	20.00	1000000.00	Mastur	\N	2	1	t	0.00	1250000.00	'-2021':4 '-3351':17 '000000073':1 '2':27 '2015':24 '30':2 'ada':15 'adira':9 'auto':6 'b':16 'beat':19 'cabang':10 'des':3 'discret':7 'finance':5,8 'gudang':21 'honda':25 'jatibarang':11,22 'kuh':18 'pop':20 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23
12	000000012	2021-11-04	2021-11-04	1000000.00	20.00	800000.00	Mastur	\N	6	1	t	0.00	1000000.00	'-2021':4 '-3479':15 '000000012':1 '04':2 '2':25 '2016':22 'ada':13 'b':14 'cabang':8 'col':7 'collectius':6 'finance':5 'gudang':19,20 'jatibarang':9 'm3':18 'mio':17 'nop':3 'r2':26 'roda':24 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':21 'uju':16 'yamaha':23
14	000000014	2021-12-07	2021-12-07	1500000.00	20.00	1200000.00	Mastur	\N	9	1	t	0.00	1500000.00	'-2021':4 '-3521':18 '000000014':1 '07':2 '2':28 '2012':25 'ada':16 'cabang':11 'des':3 'finance':5,9 'fu':21 'gudang':22,23 'jatibarang':12 'kl':19 'mitra':6 'mpmf':10 'mustika':8 'pinasthika':7 'r2':29 'roda':27 'satria':20 'stnk':15 'stnk-ada':14 'suzuk':26 'syaenudin':13 't':17 'tahun':24
61	000000061	2021-12-29	2021-12-29	1500000.00	20.00	1200000.00	Mastur	\N	7	3	t	0.00	1500000.00	'-2021':4 '-3310':19 '000000061':1 '2':28 '2015':25 '29':2 'ada':17 'cabang':10 'deddy':13 'des':3 'e':18 'finance':5,8 'gudang':22 'indramayu':12 'mandir':6 'muf':9 'pranoto':14 'pusat':11,23 'qr':20 'r2':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'utama':7 'vixion':21 'yamaha':26
37	000000037	2021-12-02	2021-12-02	18700000.00	9.09	17000000.00	Mastur	\N	14	3	t	0.00	18700000.00	'-2021':4 '-8936':19 '000000037':1 '02':2 '2006':25 '4':28 'ada':17 'b':18 'cabang':10 'clip':8 'clipan':6 'deddy':13 'des':3 'finance':5 'gudang':22 'honda':26 'indramayu':12 'jazz':21 'k':9 'karawang':7 'no':20 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24
1	000000001	2021-12-15	2021-12-15	1300000.00	20.00	1040000.00	Mastur	\N	5	3	t	0.00	1300000.00	'-2021':4 '-5605':19 '000000001':1 '15':2 '2':29 '2017':26 'ada':17 'cabang':10 'deddy':13 'des':3 'e':18 'finance':5 'gudang':23 'indramayu':12 'kredit':7 'mio':21 'motor':8 'oto':6 'otto':9 'pas':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'yamaha':27 'z':22
39	000000039	2022-01-14	2022-01-14	26000000.00	6.92	24200000.00	Mastur	\N	18	3	t	0.00	26000000.00	'-2022':4 '-8715':21 '000000039':1 '1000':24 '14':2 '2016':28 '4':31 'ada':19 'brio':23 'cabang':12 'deddy':15 'finance':5,8 'gp':22 'gudang':25 'h':20 'honda':29 'indramayu':14 'jan':3 'mandir':6 'mtf':10 'pranoto':16 'pusat':13,26 'r4':32 'roda':30 's':11 'semarang':9 'stnk':18 'stnk-ada':17 'tahun':27 'tunas':7
2	000000002	2022-01-10	2022-01-10	1300000.00	20.00	1040000.00	Mastur	\N	6	3	t	0.00	1300000.00	'-2022':4 '-3977':17 '000000002':1 '10':2 '2':26 '2016':23 'ada':15 'cabang':8 'col':7 'collectius':6 'deddy':11 'e':16 'finance':5 'gudang':20 'indramayu':10 'jan':3 'mio':19 'pac':18 'pranoto':12 'pusat':9,21 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'tahun':22 'yamaha':24
51	000000051	2021-12-21	2021-12-21	25000000.00	8.00	23000000.00	Mastur	\N	19	3	t	0.00	25000000.00	'-1250':17 '-2021':4 '000000051':1 '2013':23 '21':2 '4':26 'ada':15 'cabang':8 'clip':7 'clipan':6 'deddy':11 'des':3 'du':18 'finance':5 'gudang':20 'indramayu':10 'ios':19 'pranoto':12 'pusat':9,21 'r4':27 'roda':25 'stnk':14 'stnk-ada':13 't':16 'tahun':22 'toyota':24
56	000000056	2022-01-21	2022-01-21	22000000.00	9.09	20000000.00	Mastur	\N	18	3	t	0.00	22000000.00	'-2022':4 '-9086':21 '000000056':1 '2018':27 '21':2 '4':30 'ada':19 'agya':23 'cabang':12 'deddy':15 'finance':5,8 'gudang':24 'h':20 'indramayu':14 'jan':3 'mandir':6 'mtf':10 'pranoto':16 'pusat':13,25 'r4':31 'roda':29 's':11 'semarang':9 'stnk':18 'stnk-ada':17 'tahun':26 'te':22 'toyota':28 'tunas':7
103	000000175	2022-03-28	2022-03-28	1000000.00	20.00	800000.00	Opick	\N	23	1	t	0.00	1000000.00	'-2022':4 '-3541':16 '000000175':1 '125':19 '2':26 '2014':23 '28':2 'ada':14 'bima':6,8 'cabang':9 'e':15 'finance':5,7 'gudang':20 'honda':24 'jatibarang':10,21 'mar':3 'qn':17 'r2':27 'roda':25 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':22 'vario':18
110	000000182	2022-03-31	2022-03-31	1600000.00	20.00	1280000.00	Opick	\N	2	1	t	0.00	1600000.00	'-2022':4 '-2298':17 '000000182':1 '2':26 '2016':23 '31':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'jatibarang':11,21 'je':18 'mar':3 'nmax':19 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'yamaha':24
105	000000177	2022-03-29	2022-03-29	1300000.00	20.00	1040000.00	Opick	\N	11	1	t	0.00	1300000.00	'-2022':4 '-4027':16 '000000177':1 '2':27 '2018':24 '29':2 'ada':14 'cabang':9 'd':15 'finance':5 'gudang':21 'jatibarang':10,22 'king':20 'kp':8 'kredit':6 'mar':3 'mx':19 'mx-king':18 'plus':7 'r2':28 'roda':26 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':23 'ucb':17 'yamaha':25
113	000000185	2022-03-31	2022-03-31	1100000.00	20.00	880000.00	Opick	\N	2	1	t	0.00	1100000.00	'-2022':4 '-4479':17 '000000185':1 '125':20 '2':27 '2018':24 '31':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'fino':19 'gudang':21 'jatibarang':11,22 'mar':3 'paw':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
60	000000060	2021-09-30	2021-09-30	920000.00	0.00	920000.00	Mastur	\N	2	3	t	0.00	920000.00	'-2021':4 '-6181':19 '000000060':1 '2':28 '2018':25 '30':2 'ada':17 'adira':9 'auto':6 'beat':21 'cabang':10 'deddy':13 'discret':7 'f':18 'fch':20 'finance':5,8 'gudang':22 'honda':26 'indramayu':12 'pranoto':14 'pusat':11,23 'r2':29 'roda':27 'sep':3 'stnk':16 'stnk-ada':15 'tahun':24
111	000000183	2022-03-31	2022-03-31	1300000.00	20.00	1040000.00	Opick	\N	11	1	t	0.00	1300000.00	'-2022':4 '-4733':16 '000000183':1 '2':26 '2015':23 '31':2 'ada':14 'cabang':9 'e':15 'finance':5 'gudang':20 'jatibarang':10,21 'jb':17 'kp':8 'kredit':6 'm3':19 'mar':3 'mio':18 'plus':7 'r2':27 'roda':25 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':22 'yamaha':24
118	000000190	2022-04-01	2022-04-01	1300000.00	20.00	1040000.00	Opick	\N	7	1	t	0.00	1300000.00	'-2022':4 '-3637':17 '000000190':1 '01':2 '2':26 '2019':23 'ada':15 'apr':3 'beat':19 'cabang':10 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'mandir':6 'muf':9 'qah':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'utama':7
75	000000075	2022-02-07	2022-02-07	1600000.00	20.00	1280000.00	Mastur	\N	13	1	t	0.00	1600000.00	'-2022':4 '-4080':15 '000000075':1 '07':2 '2':25 '2018':22 'ada':13 'cabang':8 'e':14 'feb':3 'finance':5 'gapara':6 'gudang':19 'jatibarang':9,20 'm3':18 'mio':17 'mpr':7 'r2':26 'roda':24 'stnk':12 'stnk-ada':11 'syaenudin':10 'tahun':21 'uo':16 'yamaha':23
128	000000238	2022-02-02	2022-02-02	1300000.00	20.00	1040000.00	Opick	\N	11	1	t	0.00	1300000.00	'-2022':4 '-2392':16 '000000238':1 '02':2 '2':25 '2017':22 'ada':14 'beat':18 'cabang':9 'cm':17 'feb':3 'finance':5 'gudang':19 'honda':23 'jatibarang':10,20 'kp':8 'kredit':6 'plus':7 'r2':26 'roda':24 'stnk':13 'stnk-ada':12 'syaenudin':11 'tahun':21 'z':15
85	000000111	2022-03-24	2022-03-26	1400000.00	20.00	1120000.00	Opick	\N	2	1	t	0.00	1400000.00	'-2022':4 '-2867':17 '000000111':1 '2':27 '2019':24 '24':2 'ada':15 'adira':9 'auf':18 'auto':6 'cabang':10 'discret':7 'finance':5,8 'g':16 'gudang':21 'jatibarang':11,22 'm3':20 'mar':3 'mio':19 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
27	000000027	2022-03-09	2022-03-09	700000.00	20.00	560000.00	Mastur	\N	2	1	t	0.00	700000.00	'-2022':4 '-6819':17 '000000027':1 '09':2 '2':26 '2014':23 'ada':15 'adira':9 'auto':6 'b':16 'cabang':10 'discret':7 'finance':5,8 'gudang':20 'jatibarang':11,21 'mar':3 'pzi':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'xeon':19 'yamaha':24
94	000000119	2022-03-23	2022-03-26	1600000.00	20.00	1280000.00	Opick	\N	1	3	t	0.00	1600000.00	'-2022':4 '-2753':19 '000000119':1 '125':22 '2':29 '2018':26 '23':2 'ada':17 'auto':7 'baf':9 'bussan':6 'cabang':10 'deddy':13 'e':18 'finance':5,8 'fino':21 'gudang':23 'indramayu':12 'mar':3 'pba':20 'pranoto':14 'pusat':11,24 'r2':30 'roda':28 'stnk':16 'stnk-ada':15 'tahun':25 'yamaha':27
130	000000244	2022-04-04	2022-04-04	1200000.00	20.00	960000.00	Opick	\N	8	3	f	200000.00	1400000.00	'-2022':4 '-6314':19 '000000244':1 '04':2 '2':29 '2019':26 'ada':17 'apr':3 'beat':21 'cabang':9 'deddy':12 'e':18 'fif':6,8 'finance':5 'group':7 'gudang':23 'honda':27 'indramayu':11 'pop':22 'pranoto':13 'pusat':10,24 'qr':20 'r2':30 'roda':28 'stnk':15 'stnk-tidak-ada':14 'tahun':25 'tidak':16
34	000000034	2022-03-17	2022-03-17	900000.00	20.00	720000.00	Mastur	\N	2	1	t	0.00	900000.00	'-2022':4 '-4593':17 '000000034':1 '17':2 '2':26 '2012':23 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'jatibarang':11,21 'jupiter':19 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22 'tq':18 'yamaha':24
52	000000052	2021-12-23	2021-12-23	15000000.00	7.00	13950000.00	Mastur	\N	14	3	t	0.00	15000000.00	'-1184':19 '-2021':4 '000000052':1 '1000':22 '2018':26 '23':2 '4':29 'ada':17 'brio':21 'cabang':10 'clip':8 'clipan':6 'deddy':13 'des':3 'finance':5 'ga':20 'gudang':23 'honda':27 'indramayu':12 'k':9 'karawang':7 'pranoto':14 'pusat':11,24 'r4':30 'roda':28 'stnk':16 'stnk-ada':15 't':18 'tahun':25
54	000000054	2021-12-28	2021-12-28	5400000.00	0.00	5400000.00	Mastur	\N	2	3	t	0.00	5400000.00	'-1242':19 '-2021':4 '000000054':1 '2012':25 '28':2 '4':28 'ada':17 'adira':9 'auto':6 'cabang':10 'd':18 'deddy':13 'des':3 'discret':7 'finance':5,8 'gudang':22 'indramayu':12 'ou':20 'pranoto':14 'pusat':11,23 'r4':29 'roda':27 'stnk':16 'stnk-ada':15 'tahun':24 'toyota':26 'vios':21
36	000000036	2022-03-18	2022-03-18	1500000.00	20.00	1200000.00	Mastur	\N	1	1	t	0.00	1500000.00	'-2022':4 '-5713':17 '000000036':1 '18':2 '2':27 '2018':24 'ada':15 'auto':7 'baf':9 'bussan':6 'cabang':10 'e':16 'finance':5,8 'gudang':21 'jatibarang':11,22 'm3':20 'mar':3 'mio':19 'pav':18 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'yamaha':25
104	000000176	2022-03-28	2022-03-28	1300000.00	20.00	1040000.00	Opick	\N	2	1	t	0.00	1300000.00	'-2022':4 '-4530':17 '000000176':1 '2':26 '2021':23 '28':2 'ada':15 'adira':9 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'jv':18 'mar':3 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
112	000000184	2022-03-31	2022-03-31	900000.00	20.00	720000.00	Opick	\N	2	1	t	0.00	900000.00	'-2022':4 '-2859':17 '000000184':1 '2':27 '2013':24 '31':2 'ada':15 'adira':9 'auto':6 'cabang':10 'discret':7 'e':16 'finance':5,8 'gt':20 'gudang':21 'jatibarang':11,22 'mar':3 'mio':19 'r2':28 'roda':26 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':23 'xh':18 'yamaha':25
129	000000239	2022-04-04	2022-04-04	1400000.00	20.00	1120000.00	Opick	\N	2	1	t	0.00	1400000.00	'-2022':4 '-4630':17 '000000239':1 '04':2 '2':26 '2018':23 'ada':15 'adira':9 'apr':3 'auto':6 'beat':19 'cabang':10 'discret':7 'e':16 'finance':5,8 'gudang':20 'honda':24 'jatibarang':11,21 'pax':18 'r2':27 'roda':25 'stnk':14 'stnk-ada':13 'syaenudin':12 'tahun':22
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
299	50	trx-lent	Kurangan Pencairan	2022-04-04	\N	'cair':2 'kurang':1 'norman':3
304	56	trx-lent	Kurangan Pencairan	2022-04-04	\N	'cair':2 'kurang':1 'norman':3
310	42	trx-lent	Kurangan Pencairan	2022-04-04	\N	'brio':4 'cair':2 'kurang':1 'norman':3
300	51	trx-lent	Kurangan Pencairan	2022-04-04	\N	'bahrudin':3 'cair':2 'du':5 'ios':4 'kurang':1
302	53	trx-lent	Kurangan Pencairan	2022-04-04	\N	'bahrudin':3 'cair':2 'kurang':1 'xenia':4
303	55	trx-lent	Kurangan Pencairan	2022-04-04	\N	'bahrudin':3 'brv':4 'cair':2 'kurang':1
306	59	trx-lent	Kurangan Pencairan	2022-04-04	\N	'bahrudin':3 'cair':2 'ertiga':4 'kurang':1
301	52	trx-lent	Kurangan Pencairan	2022-04-04	\N	'cair':2 'kurang':1 'nus':4 'samuel':3
305	58	trx-lent	Kurangan Pencairan	2022-04-04	\N	'cair':2 'ferdin':3 'kurang':1
307	49	trx-lent	Kurangan Tebus Unit	2022-04-04	\N	'bfi':4 'jazz':5 'kurang':1 'tebus':2 'unit':3
308	54	trx-lent	Kekurangan Tebus Unit	2022-04-04	\N	'adira':4 'kurang':1 'tebus':2 'unit':3 'vios':5
309	57	trx-lent	Kurangan Tebus Unit	2022-04-04	\N	'adira':4 'grandmax':5 'kurang':1 'tebus':2 'unit':3
311	96	trx-lent	Kurangan Pencairan	2022-04-04	\N	'bang':4 'cair':2 'kei':5 'kurang':1 'sofyan':3
314	56	trx-cicilan	Cicilan ke 1	2021-12-21	Cicilan Norman Agya	'1':3 'cicil':1 'ke':2
313	50	trx-cicilan	gagal Pajak cicilan ke 1	2021-12-16	Cicilan Norman Brio 1	'1':5,9 'brio':8 'cicil':3,6 'gagal':1 'ke':4 'norman':7 'pajak':2
312	50	trx-cicilan	bayar cicilan ke-1	2021-12-16	Cicilan Norman Brio 1	'-1':4 '1':8 'bayar':1 'brio':7 'cicil':2,5 'ke':3 'norman':6
315	42	trx-cicilan	Cicilan ke 1	2022-02-16	Cicilan Norman Brio	'1':3 'cicil':1 'ke':2
316	51	trx-cicilan	Cicilan ke 1	2021-12-21	Cicilan Bahrudin Terios DU	'1':3 'cicil':1 'ke':2
317	53	trx-cicilan	Cicilan ke 1	2021-12-21	Cicilan Bahrudin Xenia	'1':3 'cicil':1 'ke':2
318	55	trx-cicilan	Cicilan ke 1	2022-01-17	Cicilan Bahrudin BRV	'1':3 'cicil':1 'ke':2
319	59	trx-cicilan	Cicilan ke 1	2022-01-31	Cicilan Bahrudin Ertiga	'1':3 'bahrudin':5 'cicil':1,4 'ertiga':6 'ke':2
320	52	trx-cicilan	cicilan ke 1	2021-12-23	Cicilan Samuel Nus	'1':3 'cicil':1 'ke':2
321	58	trx-cicilan	Cicilan 1	2021-01-31	Cicilan Ferdinan	'1':2 'cicil':1
322	49	trx-cicilan	Cicilan 1	2021-11-22	Cicilan BFI JAZZ	'1':2 'cicil':1
324	57	trx-cicilan	Cicilan 1	2022-01-22	Cicilan Adira Grandmax	'1':2 'cicil':1
323	54	trx-cicilan	Cicilan 1	2021-12-28	Cicilan Adira Vios	'1':2 'adira':4 'cicil':1,3 'vios':5
328	4	trx-loan	Abdul Sholeh	2022-04-04	\N	'abdul':1,3 'sholeh':2,4
325	96	trx-cicilan	Cicilan 1	2022-03-21	Cicilan Sofyan/ Bang Kei	'1':2 'bang':5 'cicil':1,3 'kei':6 'sofyan':4
334	124	trx-lent	kekurangan pencairan	2022-04-05	\N	'cair':2 'gineng':4 'kurang':1 'pa':3
335	124	trx-cicilan	cicilan 1	2022-04-05	Cicilan Pa gineng	'1':2 'cicil':1
340	12	trx-loan	kekurangan pencairan 	2022-01-22	\N	'an':3 'arsim':4 'cair':2 'kurang':1
341	12	trx-angsuran	cicll 1	2022-04-05	Angsuran An Arsim	'1':2 'cicll':1
344	15	trx-loan	pinjeman pribadi 	2022-01-22	\N	'ipeng':4 'pak':3 'pinjem':1 'pribad':2
345	16	trx-loan	pinjeman pribadi 	2022-01-22	\N	'agus':4 'baf':5 'pa':3 'pinjem':1 'pribad':2
347	18	trx-loan	DANA PINJAMAN	2022-01-22	\N	'an':4 'cardi':5 'dana':1 'pelus':3 'pinjam':2
349	65	trx-lent	unit hilang	2022-04-05	\N	'hilang':2,3 'unit':1
350	66	trx-lent	Pinjam Pakai	2022-04-05	\N	'mastukin':3 'paka':2 'pinjam':1
351	67	trx-lent	Pinjam Pakai	2022-04-05	\N	'jiring':4 'paka':2 'pinjam':1 'wa':3
352	69	trx-lent	Pinjam Pakai	2022-04-05	\N	'anak':5 'jileng':4 'paka':2 'pinjam':1 'wa':3
348	71	trx-lent	Pinjam Pakai	2022-04-05	\N	'andre':3 'paka':2 'pinjam':1
353	72	trx-lent	KF Adira JTB	2022-04-05	\N	'adira':2 'cabang':6 'jawab':5 'jtb':3,7 'kf':1 'tanggung':4
354	73	trx-lent	KF Adira	2022-04-05	\N	'2':7 'adira':2 'cabang':5 'jawab':4 'jtb':6 'kf':1 'tanggung':3
436	133	trx-lent	qweqweqweqwe	2022-04-09	\N	'dasd':2 'qweqweqweqwe':1 'werqwerqwer':3
355	75	trx-lent	Pinjem pakai	2022-04-05	\N	'ali':4 'eksternal':3 'paka':2 'pinjem':1
357	18	trx-lent	Pinjem Pakai	2022-04-05	\N	'bule':4 'eksternal':3 'paka':2 'pinjem':1
358	84	trx-lent	Pinjem Pakai	2022-04-05	\N	'eksternal':3 'jatibarang':4 'paka':2 'pinjem':1
361	84	trx-cicilan	Cicilan 1	2021-12-01	Cicilan Eksternal Jatibarang	'1':2 'cicil':1
362	18	trx-cicilan	Cicilan 1	2021-12-01	Cicilan Eksternal Bule	'1':2 'cicil':1
437	133	trx-cicilan	Cicilan 1	2022-04-09	Cicilan Clip Cikarang	'1':2 'cicil':1
408	60	trx-lent	Unit Hilang	2022-04-05	\N	'hilang':2,3 'unit':1
409	61	trx-lent	Pinjam pakai	2022-04-05	\N	'admin':4 'opick':3 'paka':2 'pinjam':1
410	62	trx-lent	Kurangan Tebus Unit	2022-04-05	\N	'kurang':1 'om':5 'saja':6 'tebus':2 'tim':4 'unit':3
412	121	trx-lent	Pinjam Unit	2022-04-05	\N	'gudang':4 'opick':3 'pinjam':1 'unit':2
421	62	trx-cicilan	cicilan 1	2021-12-01	Cicilan Tim Om Saja	'1':2 'cicil':1
\.


--
-- Data for Name: trx_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trx_detail (id, code_id, trx_id, debt, cred) FROM stdin;
2	1113	299	0.00	31360000.00
2	1113	300	0.00	23000000.00
2	1113	301	0.00	13950000.00
2	1113	302	0.00	19500000.00
2	1113	303	0.00	38400000.00
2	1113	304	0.00	20000000.00
2	1113	305	0.00	20000000.00
2	1113	306	0.00	19750000.00
2	1113	307	0.00	8200000.00
2	1113	308	0.00	5400000.00
2	1113	309	0.00	8000000.00
2	1113	310	0.00	13200000.00
2	1113	311	0.00	12300000.00
1	1113	314	19600000.00	0.00
2	4113	314	0.00	19600000.00
1	1113	313	150000.00	0.00
2	4113	313	0.00	150000.00
1	1113	312	17000000.00	0.00
2	4113	312	0.00	17000000.00
1	1113	315	8350000.00	0.00
2	4113	315	0.00	8350000.00
1	1113	316	20000000.00	0.00
2	4113	316	0.00	20000000.00
1	1113	317	19000000.00	0.00
2	4113	317	0.00	19000000.00
1	1113	318	30000000.00	0.00
2	4113	318	0.00	30000000.00
1	1113	319	16000000.00	0.00
2	4113	319	0.00	16000000.00
1	1113	320	12500000.00	0.00
2	4113	320	0.00	12500000.00
1	1113	321	20000000.00	0.00
2	4113	321	0.00	20000000.00
1	1113	322	5000000.00	0.00
2	4113	322	0.00	5000000.00
1	1113	324	3637000.00	0.00
2	4113	324	0.00	3637000.00
1	1113	323	2000000.00	0.00
2	4113	323	0.00	2000000.00
1	5512	328	11665000.00	0.00
2	1113	328	0.00	11665000.00
1	1113	325	12500000.00	0.00
2	4113	325	0.00	12500000.00
2	1113	334	0.00	375000.00
1	1113	335	196000.00	0.00
2	4113	335	0.00	196000.00
1	5512	340	10200000.00	0.00
2	1113	340	0.00	10200000.00
1	1113	341	7910000.00	0.00
2	4112	341	0.00	7910000.00
1	5512	344	2000000.00	0.00
2	1113	344	0.00	2000000.00
1	5512	345	2000000.00	0.00
2	1113	345	0.00	2000000.00
1	5512	347	6000000.00	0.00
2	1113	347	0.00	6000000.00
2	1113	348	0.00	1040000.00
2	1113	349	0.00	3200000.00
2	1113	350	0.00	1040000.00
2	1113	351	0.00	1600000.00
2	1113	352	0.00	1120000.00
2	1113	353	0.00	1000000.00
2	1113	354	0.00	1000000.00
2	1113	355	0.00	1280000.00
2	1113	357	0.00	800000.00
2	1113	358	0.00	840000.00
1	1113	361	500000.00	0.00
2	4113	361	0.00	500000.00
1	1113	362	400000.00	0.00
2	4113	362	0.00	400000.00
1	5513	299	31360000.00	0.00
1	5513	300	23000000.00	0.00
1	5513	301	13950000.00	0.00
1	5513	302	19500000.00	0.00
1	5513	436	49066000.00	0.00
2	1113	436	0.00	49066000.00
1	1113	437	47500000.00	0.00
2	4113	437	0.00	47500000.00
2	1113	408	0.00	920000.00
2	1113	409	0.00	1200000.00
2	1113	410	0.00	2880000.00
2	1113	412	0.00	1280000.00
1	1113	421	1300000.00	0.00
2	4113	421	0.00	1300000.00
1	5513	303	38400000.00	0.00
1	5513	304	20000000.00	0.00
1	5513	305	20000000.00	0.00
1	5513	306	19750000.00	0.00
1	5513	307	8200000.00	0.00
1	5513	308	5400000.00	0.00
1	5513	309	8000000.00	0.00
1	5513	310	13200000.00	0.00
1	5513	311	12300000.00	0.00
1	5513	334	375000.00	0.00
1	5513	348	1040000.00	0.00
1	5513	349	3200000.00	0.00
1	5513	350	1040000.00	0.00
1	5513	351	1600000.00	0.00
1	5513	352	1120000.00	0.00
1	5513	353	1000000.00	0.00
1	5513	354	1000000.00	0.00
1	5513	355	1280000.00	0.00
1	5513	357	800000.00	0.00
1	5513	358	840000.00	0.00
1	5513	408	920000.00	0.00
1	5513	409	1200000.00	0.00
1	5513	410	2880000.00	0.00
1	5513	412	1280000.00	0.00
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
43	AEROX	2	2
44	Box	3	1
45	Mio GT	2	2
46	Nex	2	12
47	X-Ride	2	2
40	CB-150	2	13
42	Sonic	2	13
41	Scoopy	2	13
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
65	E 2676 UX	2019	\N	\N	\N	35	2
66	T 6934 YQ	2017	\N	\N	\N	6	2
67	T 4845 IQ	2020	\N	\N	\N	35	2
69	B 3433 USN	2019	\N	\N	\N	6	2
71	T 4261 YQ	2017	\N	\N	\N	38	2
72	B 3430 EJX	2016	\N	\N	\N	6	2
73	B 3351 KUH	2015	\N	\N	\N	37	2
74	E 6871 CM	2018	\N	\N	\N	8	2
75	E 4080 UO	2018	\N	\N	\N	9	2
77	E 5826 CQ	2019	\N	\N	\N	2	2
79	E 2282 PBC	2018	\N	\N	\N	6	2
80	E 5737 PBO	2019	\N	\N	\N	1	2
83	T 3802 ZR	2019	\N	\N	\N	41	2
84	E 5462 QM	2014	\N	\N	\N	6	2
85	G 2867 AUF	2019	\N	\N	\N	9	2
94	E 2753 PBA	2018	\N	\N	\N	1	1
96	B 1301 UZR	2022	\N	\N	\N	3	1
99	T 8066 EG	2020	\N	\N	\N	23	1
100	E 9844 HB	2016	\N	\N	\N	44	1
101	E 6987 SD	2013	\N	\N	\N	13	2
103	E 3541 QN	2014	\N	\N	\N	2	2
104	E 4530 JV	2021	\N	\N	\N	6	2
105	D 4027 UCB	2018	\N	\N	\N	38	2
107	E 5359 PAI	2016	\N	\N	\N	1	2
108	E 5712 PBG	2019	\N	\N	\N	1	2
109	E 2575 PAS	2017	\N	\N	\N	6	2
110	E 2298 JE	2016	\N	\N	\N	35	2
111	E 4733 JB	2015	\N	\N	\N	9	2
112	E 2859 XH	2013	\N	\N	\N	45	2
113	E 4479 PAW	2018	\N	\N	\N	1	2
114	T 4418 PF	2017	\N	\N	\N	1	2
115	T 4080 WE	2013	\N	\N	\N	46	2
118	E 3637 QAH	2019	\N	\N	\N	6	2
119	E 6945 PBC	2018	\N	\N	\N	6	2
121	B 4172 FPB	2022	\N	\N	\N	47	1
123	B 1561 KBJ	2007	\N	\N	\N	31	1
124	E 2000 XX	2022	\N	\N	\N	29	1
125	E 1411 PV	2020	\N	\N	\N	4	1
126	E 1456 RL	2020	\N	\N	\N	4	1
127	A 3907 PBF	2022	\N	\N	\N	1	2
128	Z 2392 CM	2017	\N	\N	\N	6	2
129	E 4630 PAX	2018	\N	\N	\N	6	2
130	E 6314 QR	2019	\N	\N	\N	37	1
131	E 5491 PBL	2022	\N	\N	\N	2	1
132	E 4415 JZ	2019	\N	\N	\N	35	2
133	T 1293 TN	2014	\N	\N	\N	34	1
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
7	gudang adira	jatibarang
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

SELECT pg_catalog.setval('public.action_id_seq', 12, true);


--
-- Name: branch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.branch_id_seq', 5, true);


--
-- Name: finance_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finance_groups_id_seq', 8, true);


--
-- Name: finance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finance_id_seq', 23, true);


--
-- Name: invoices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.invoices_id_seq', 11, true);


--
-- Name: lents_id_sequence; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lents_id_sequence', 50, true);


--
-- Name: loans_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.loans_id_seq', 18, true);


--
-- Name: merk_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.merk_id_seq', 17, true);


--
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_id_seq', 133, true);


--
-- Name: order_name_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_name_seq', 272, true);


--
-- Name: trx_detail_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_detail_seq', 1, false);


--
-- Name: trx_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_seq', 437, true);


--
-- Name: type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.type_id_seq', 47, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: warehouse_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.warehouse_id_seq', 7, true);


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

