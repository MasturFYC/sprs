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
5118	Biaya Konsumsi	51	Biaya yg dikeluarkan karena suatu kegiatan yg dpt mengurangi atau menghabiskan barang dan jasa	'atau':12 'barang':14 'biaya':1,3 'dan':15 'dpt':10 'giat':8 'habis':13 'jasa':16 'karena':6 'keluar':5 'konsumsi':2 'suatu':7 'urang':11 'yg':4,9	t	t	3
5511	Piutang Jasa	55	Piutang diberikan kepada pihak Finance yg timbul karena perusahaan menerima job penarikan kendaraan sejumlah BT Matel\n	'arik':14 'beri':4 'bt':17 'erima':12 'finance':7 'jasa':2 'job':13 'karena':10 'matel':18 'ndara':15 'pada':5 'pihak':6 'piutang':1,3 'sejum':16 'timbul':9 'usaha':11 'yg':8	t	f	3
2311	Hutang Pajak	23	Pajak yg belum dibayar karena menunggu pembayaran dari tarikan	'bayar':6,9 'belum':5 'dari':10 'hutang':1 'karena':7 'pajak':2,3 'tari':11 'unggu':8 'yg':4	t	f	2
3111	Modal pak Kris	31	Modal yg masuk dari pak Kris	'dari':7 'kris':3,9 'masuk':6 'modal':1,4 'pak':2,8 'yg':5	t	t	2
6011	Pembayaran Pajak	60	Pajak Pertambahan Nilai	'bayar':1 'nila':5 'pajak':2,3 'tambah':4	t	f	3
5311	Biaya Gaji karyawan Tetap	53	Pencatatan data kompensasi karyawan seperti uang potongan dari setiap gaji dan pajak serta tunjangan karyawan tetap	'biaya':1 'catat':5 'dan':15 'dari':12 'data':6 'gaji':2,14 'karyaw':3,8,19 'kompensasi':7 'pajak':16 'potong':11 'sepert':9 'serta':17 'setiap':13 'tetap':4,20 'tunjang':18 'uang':10	t	t	3
5312	Biaya Gaji Karyawan Honorer	51	Pencatatan data kompensasi karyawan seperti uang potongan dari setiap gaji\ndan pajak serta tunjangan bukan karyawan tetap 	'biaya':1 'bukan':19 'catat':5 'dan':15 'dari':12 'data':6 'gaji':2,14 'honorer':4 'karyaw':3,8,20 'kompensasi':7 'pajak':16 'potong':11 'sepert':9 'serta':17 'setiap':13 'tetap':21 'tunjang':18 'uang':10	t	f	3
1111	Kas Kantor	11	Kas bendahara Kantor	'bendahara':4 'kantor':2,5 'kas':1,3	t	f	1
5211	Kasbon Cabang JTB	52	Biaya yg dikeluarkan untuk penarikan kendaraan yg tidak ada STNK	'ada':12 'arik':8 'biaya':4 'cabang':2 'jtb':3 'kasbon':1 'keluar':6 'ndara':9 'stnk':13 'tidak':11 'untuk':7 'yg':5,10	t	t	3
2211	BNI	22	\N	'bni':1	t	t	2
2212	SAMSAT	22	\N	'samsat':1	t	t	2
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
55	Piutang	Kelompok akun yg mencatat semua pengeluaran dalam bentuk piutang kepada pihak lain.	5
\.


--
-- Data for Name: actions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.actions (id, action_at, pic, descriptions, order_id, file_name) FROM stdin;
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
23	test	\N	c0-1
\.


--
-- Data for Name: finance_groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.finance_groups (id, name) FROM stdin;
1	BAF
2	CLIPAN
3	MTF
\.


--
-- Data for Name: finances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.finances (id, name, short_name, street, city, phone, cell, zip, email, group_id) FROM stdin;
1	Bussan Auto Finance	BAF	Jl. Jend. Sudirman	Indramayu	2569874545	65979	2598987	busan.123@gmail.com	1
2	Auto Discret Finance	Adira	Jl. Jend. Sudirman	Indramayu	2569874545	65979	2598987	adira.finance@gmail.com	1
5	OTO Kredit Motor	OTTO	\N	\N	\N	\N	\N	\N	1
6	COLLECTIUS	COL	\N	\N	\N	\N	\N	\N	1
7	Mandiri Utama Finance	MUF	\N	\N	\N	\N	\N	\N	1
8	FIF Group	FIF	\N	\N	\N	\N	\N	\N	1
9	Mitra Pinasthika Mustika Finance	MPMF	\N	\N	\N	\N	\N	\N	1
10	Top Finance Company	TFC	\N	\N	\N	\N	\N	\N	1
11	Kredit Plus	KP+	\N	\N	\N	\N	\N	\N	1
12	WOM Finance	WOMF	\N	\N	\N	\N	\N	\N	1
13	MEGAPARA	MPR	\N	\N	\N	\N	\N	\N	1
16	Safron Finance Karawang	SFI K	\N	\N	\N	\N	\N	\N	1
17	BFI Finance	BFI	\N	\N	\N	\N	\N	\N	1
20	Radana Finance	RAD	\N	\N	\N	\N	\N	\N	1
21	Mega Auto Central Finance	MACF	\N	\N	\N	\N	\N	\N	1
19	CLIPAN	CLIP	\N	\N	\N	\N	\N	\N	2
14	Clipan Karawang\n	CLIP K	\N	\N	\N	\N	\N	\N	2
15	Clipan Palembang	CLIP P	\N	\N	\N	\N	\N	\N	2
3	Mandiri Tunas Finance	MTF	\N	Cirebon	\N	\N	\N	\N	3
18	Mandiri Tunas Finance Semarang	MTF S	\N	\N	\N	\N	\N	\N	3
4	Clipan Bekasi	CLIP B	\N	\N	\N	\N	\N	\N	2
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
47	000000047	2022-03-16	2022-03-16	9000000.00	11.11	8000000.00	Mastur	test	15	3	t	0.00	9000000.00	'-03':3 '-16':4 '00':6 '000000047':1 '00z':7 '1623':20 '2006':24 '2022':2 '4':27 'ada':18 'bg':19 'clip':10 'clipan':8 'deddy':14 'honda':25 'indramayu':13 'jazz':22 'p':11 'palembang':9 'pf':21 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5
44	000000044	2022-03-12	2022-03-12	8000000.00	20.00	6400000.00	Mastur	test	17	3	t	0.00	8000000.00	'-03':3 '-12':4 '00':6 '000000044':1 '00z':7 '2014':23 '2022':2 '4':26 '8903':19 'ada':17 'bfi':8,10 'carry':21 'deddy':13 'e':18 'finance':9 'indramayu':12 'pp':20 'pranoto':14 'pusat':11,22 'r4':27 'roda':25 'stnk':16 'stnk-ada':15 'suzuk':24 't00':5
42	000000042	2022-02-16	2022-02-16	15000000.00	12.00	13200000.00	Mastur	test	18	3	t	0.00	15000000.00	'-02':3 '-16':4 '00':6 '000000042':1 '00z':7 '1000':25 '2020':27 '2022':2 '4':30 '9049':22 'ada':20 'brio':24 'deddy':16 'finance':10 'h':21 'honda':28 'indramayu':15 'mandir':8 'mtf':12 'pranoto':17 'pusat':14,26 'r4':31 'roda':29 's':13 'se':23 'semarang':11 'stnk':19 'stnk-ada':18 't00':5 'tunas':9
40	000000040	2022-01-17	2022-01-17	26000000.00	7.69	24000000.00	Mastur	test	14	3	t	0.00	26000000.00	'-01':3 '-17':4 '-3':23 '00':6 '000000040':1 '00z':7 '1164':20 '2017':25 '2022':2 '4':28 'ada':18 'clip':10 'clipan':8 'deddy':14 'er':22 'fq':21 'indramayu':13 'k':11 'karawang':9 'pranoto':15 'pusat':12,24 'r4':29 'roda':27 'stnk':17 'stnk-ada':16 'suzuk':26 't':19 't00':5
38	000000038	2021-12-29	2021-12-29	13000000.00	7.69	12000000.00	Mastur	test	14	3	t	0.00	13000000.00	'-12':3 '-29':4 '00':6 '000000038':1 '00z':7 '1412':20 '2004':24 '2021':2 '4':27 'ada':18 'carry':22 'clip':10 'clipan':8 'deddy':14 'indramayu':13 'k':11 'karawang':9 'km':21 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 'suzuk':25 't':19 't00':5
35	000000035	2022-03-18	2022-03-18	1800000.00	20.00	1440000.00	Mastur	test	2	1	t	0.00	1800000.00	'-03':3 '-15':21 '-18':4 '00':6 '000000035':1 '00z':7 '2':26 '2017':23 '2022':2 '2110':18 'ada':16 'adira':11 'auto':8 'discret':9 'finance':10 'jatibarang':12,22 'r':20 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 't00':5 'yamaha':24 'yv':19
33	000000033	2022-03-16	2022-03-16	1500000.00	20.00	1200000.00	Mastur	test	2	1	t	0.00	1500000.00	'-03':3 '-16':4 '00':6 '000000033':1 '00z':7 '2':25 '2019':22 '2022':2 '5856':18 'ada':16 'adira':11 'auto':8 'discret':9 'finance':10 'genio':20 'honda':23 'jatibarang':12,21 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 't00':5 'zt':19
31	000000031	2022-03-12	2022-03-12	1800000.00	20.00	1440000.00	Mastur	test	2	1	t	0.00	1800000.00	'-03':3 '-12':4 '-15':21 '00':6 '000000031':1 '00z':7 '2':26 '2017':23 '2022':2 '2391':18 'ada':16 'adira':11 'auto':8 'discret':9 'e':17 'finance':10 'jatibarang':12,22 'jm':19 'r':20 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'yamaha':24
29	000000029	2022-03-09	2022-03-09	1450000.00	20.00	1160000.00	Mastur	test	13	1	t	0.00	1450000.00	'-03':3 '-09':4 '00':6 '000000029':1 '00z':7 '2':23 '2015':20 '2022':2 '4544':16 'ada':14 'e':15 'gapara':8 'gudang':19 'jatibarang':10 'jd':17 'mio':18 'mpr':9 'r2':24 'roda':22 'stnk':13 'stnk-ada':12 'syaenudin':11 't00':5 'yamaha':21
26	000000026	2022-03-07	2022-03-07	850000.00	20.00	680000.00	Mastur	test	11	1	t	0.00	850000.00	'-03':3 '-07':4 '00':6 '000000026':1 '00z':7 '150':20 '2':25 '2014':22 '2022':2 '2891':17 'ada':15 'honda':23 'jatibarang':11,21 'kp':10 'kredit':8 'plus':9 'r2':26 'roda':24 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 't00':5 'vario':19 'wp':18
23	000000023	2022-02-23	2022-02-23	0.00	20.00	40000.00	Mastur	\N	12	1	t	0.00	0.00	'2021':1 '2815':6 'e':5 'gear':3 'pbx':7 'r2':8 'womf':4 'yamaha':2
24	000000024	2022-02-25	2022-02-25	1500000.00	20.00	1200000.00	Mastur	test	1	1	t	0.00	1500000.00	'-02':3 '-25':4 '00':6 '000000024':1 '00z':7 '2':26 '2018':23 '2022':2 '2146':18 'ada':16 'auto':9 'baf':11 'bussan':8 'e':17 'finance':10 'gudang':22 'jatibarang':12 'mio':20 'qaf':19 'r2':27 'roda':25 's':21 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'yamaha':24
21	000000021	2022-02-21	2022-02-21	950000.00	20.00	760000.00	Mastur	test	1	1	t	0.00	950000.00	'-02':3 '-21':4 '00':6 '000000021':1 '00z':7 '2':26 '2015':23 '2022':2 '6262':18 'ada':16 'auto':9 'b':17 'baf':11 'bussan':8 'finance':10 'gudang':22 'jatibarang':12 'm3':21 'mio':20 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'vky':19 'yamaha':24
15	000000015	2022-01-06	2022-01-06	900000.00	20.00	900000.00	Mastur	test	5	1	t	0.00	900000.00	'-01':3 '-06':4 '00':6 '000000015':1 '00z':7 '2':26 '2012':23 '2022':2 '4146':18 'ada':16 'gudang':22 'jatibarang':12 'jupiter':20 'ko':19 'kredit':9 'motor':10 'mx':21 'oto':8 'otto':11 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 't00':5 'yamaha':24
9	000000009	2022-03-18	2022-03-18	1300000.00	20.00	1040000.00	Mastur	test	2	3	t	0.00	1300000.00	'-03':3 '-18':4 '00':6 '000000009':1 '00z':7 '2':27 '2017':24 '2022':2 '6053':20 'ada':18 'adira':11 'auto':8 'beat':22 'deddy':14 'discret':9 'e':19 'finance':10 'honda':25 'indramayu':13 'pam':21 'pranoto':15 'pusat':12,23 'r2':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5
7	000000007	2022-03-15	2022-03-15	1700000.00	20.00	1360000.00	Mastur	test	1	3	t	0.00	1700000.00	'-03':3 '-15':4 '00':6 '000000007':1 '00z':7 '125':23 '2':28 '2019':25 '2022':2 '2033':20 'ada':18 'auto':9 'baf':11 'bussan':8 'deddy':14 'e':19 'finance':10 'fino':22 'indramayu':13 'pbj':21 'pranoto':15 'pusat':12,24 'r2':29 'roda':27 'stnk':17 'stnk-ada':16 't00':5 'yamaha':26
5	000000005	2022-02-24	2022-02-24	1500000.00	20.00	1200000.00	Mastur	test	7	3	t	0.00	1500000.00	'-02':3 '-24':4 '00':6 '000000005':1 '00z':7 '125':23 '2':28 '2017':25 '2022':2 '4096':20 'ada':18 'deddy':14 'e':19 'finance':10 'fino':22 'indramayu':13 'mandir':8 'muf':11 'paq':21 'pranoto':15 'pusat':12,24 'r2':29 'roda':27 'stnk':17 'stnk-ada':16 't00':5 'utama':9 'yamaha':26
22	000000022	2022-02-23	2022-02-23	950000.00	20.00	760000.00	Mastur	test	1	1	t	0.00	950000.00	'-02':3 '-23':4 '00':6 '000000022':1 '00z':7 '2':26 '2015':23 '2022':2 '2830':18 'ada':16 'auto':9 'baf':11 'bussan':8 'e':17 'finance':10 'gudang':22 'jatibarang':12 'm3':21 'mio':20 'qr':19 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'yamaha':24
20	000000020	2022-01-21	2022-01-21	900000.00	20.00	720000.00	Mastur	test	11	1	t	0.00	900000.00	'-01':3 '-21':4 '00':6 '000000020':1 '00z':7 '125':20 '2':25 '2013':22 '2022':2 '5253':17 'ada':15 'e':16 'gudang':21 'honda':23 'jatibarang':11 'kp':10 'kredit':8 'plus':9 'r2':26 'roda':24 'stnk':14 'stnk-ada':13 'syaenudin':12 't00':5 'ty':18 'vario':19
18	000000018	2022-01-18	2022-01-18	1000000.00	20.00	800000.00	Mastur	test	5	1	t	0.00	1000000.00	'-01':3 '-18':4 '00':6 '000000018':1 '00z':7 '2':26 '2015':23 '2022':2 '6716':18 'ada':16 'e':17 'gudang':22 'ix':19 'jatibarang':12 'kredit':9 'm3':21 'mio':20 'motor':10 'oto':8 'otto':11 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'yamaha':24
19	000000019	2022-01-26	2022-01-26	1000000.00	20.00	800000.00	Mastur	test	2	1	t	0.00	1000000.00	'-01':3 '-26':4 '00':6 '000000019':1 '00z':7 '2':25 '2018':22 '2022':2 '5638':18 'ada':16 'adira':11 'auto':8 'discret':9 'e':17 'finance':10 'honda':23 'jatibarang':12,21 'pav':19 'r2':26 'revo':20 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5
12	000000012	2021-11-04	2021-11-04	1000000.00	20.00	800000.00	Mastur	test	6	1	t	0.00	1000000.00	'-04':4 '-11':3 '00':6 '000000012':1 '00z':7 '2':24 '2016':21 '2021':2 '3479':16 'ada':14 'b':15 'col':9 'collectius':8 'gudang':20 'jatibarang':10 'm3':19 'mio':18 'r2':25 'roda':23 'stnk':13 'stnk-ada':12 'syaenudin':11 't00':5 'uju':17 'yamaha':22
13	000000013	2021-12-06	2021-12-06	1000000.00	20.00	800000.00	Mastur	test	6	1	t	0.00	1000000.00	'-06':4 '-12':3 '00':6 '000000013':1 '00z':7 '2':24 '2017':21 '2021':2 '2417':16 'ada':14 'col':9 'collectius':8 'e':15 'gudang':20 'jatibarang':10 'm3':19 'mio':18 'pao':17 'r2':25 'roda':23 'stnk':13 'stnk-ada':12 'syaenudin':11 't00':5 'yamaha':22
14	000000014	2021-12-07	2021-12-07	1500000.00	20.00	1200000.00	Mastur	test	9	1	t	0.00	1500000.00	'-07':4 '-12':3 '00':6 '000000014':1 '00z':7 '2':27 '2012':24 '2021':2 '3521':19 'ada':17 'finance':11 'fu':22 'gudang':23 'jatibarang':13 'kl':20 'mitra':8 'mpmf':12 'mustika':10 'pinasthika':9 'r2':28 'roda':26 'satria':21 'stnk':16 'stnk-ada':15 'suzuk':25 'syaenudin':14 't':18 't00':5
1	000000001	2021-12-15	2021-12-15	1300000.00	20.00	1040000.00	Mastur	test	5	3	t	0.00	1300000.00	'-12':3 '-15':4 '00':6 '000000001':1 '00z':7 '2':28 '2017':25 '2021':2 '5605':20 'ada':18 'deddy':14 'e':19 'indramayu':13 'kredit':9 'mio':22 'motor':10 'oto':8 'otto':11 'pas':21 'pranoto':15 'pusat':12,24 'r2':29 'roda':27 'stnk':17 'stnk-ada':16 't00':5 'yamaha':26 'z':23
2	000000002	2022-01-10	2022-01-10	1300000.00	20.00	1040000.00	Mastur	test	6	3	t	0.00	1300000.00	'-01':3 '-10':4 '00':6 '000000002':1 '00z':7 '2':25 '2016':22 '2022':2 '3977':18 'ada':16 'col':9 'collectius':8 'deddy':12 'e':17 'indramayu':11 'mio':20 'pac':19 'pranoto':13 'pusat':10,21 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 't00':5 'yamaha':23
3	000000003	2022-02-07	2022-02-07	1300000.00	20.00	1040000.00	Mastur	test	1	3	t	0.00	1300000.00	'-02':3 '-07':4 '00':6 '000000003':1 '00z':7 '2':28 '2018':25 '2022':2 '5125':20 'ada':18 'auto':9 'baf':11 'bussan':8 'deddy':14 'e':19 'finance':10 'indramayu':13 'm3':23 'mio':22 'pbc':21 'pranoto':15 'pusat':12,24 'r2':29 'roda':27 'stnk':17 'stnk-ada':16 't00':5 'yamaha':26
4	000000004	2022-02-18	2022-02-18	1200000.00	20.00	960000.00	Mastur	test	2	3	t	0.00	1200000.00	'-02':3 '-18':4 '00':6 '000000004':1 '00z':7 '2':27 '2015':24 '2022':2 '5080':20 'ada':18 'adira':11 'auto':8 'br':19 'deddy':14 'discret':9 'finance':10 'indramayu':13 'pranoto':15 'pusat':12,23 'py':21 'r2':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5 'vixion':22 'yamaha':25
6	000000006	2022-03-02	2022-03-02	1200000.00	20.00	960000.00	Mastur	test	2	3	t	0.00	1200000.00	'-02':4 '-03':3 '00':6 '000000006':1 '00z':7 '2':27 '2016':24 '2022':2 '2633':20 'ada':18 'adira':11 'auto':8 'beat':22 'deddy':14 'discret':9 'e':19 'finance':10 'honda':25 'indramayu':13 'pac':21 'pranoto':15 'pusat':12,23 'r2':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5
8	000000008	2022-03-17	2022-03-17	1700000.00	20.00	1360000.00	Mastur	test	2	3	t	0.00	1700000.00	'-03':3 '-15':23 '-17':4 '00':6 '000000008':1 '00z':7 '2':28 '2018':25 '2022':2 '6277':20 'ada':18 'adira':11 'auto':8 'deddy':14 'discret':9 'e':19 'finance':10 'indramayu':13 'jatibarang':24 'paz':21 'pranoto':15 'pusat':12 'r':22 'r2':29 'roda':27 'stnk':17 'stnk-ada':16 't00':5 'yamaha':26
10	000000010	2022-03-19	2022-03-19	850000.00	20.00	680000.00	Mastur	test	2	3	t	0.00	850000.00	'-03':3 '-19':4 '00':6 '000000010':1 '00z':7 '2':27 '2013':24 '2022':2 '5474':20 'ada':18 'adira':11 'auto':8 'beat':22 'deddy':14 'discret':9 'e':19 'finance':10 'honda':25 'indramayu':13 'pranoto':15 'pusat':12,23 'q':21 'r2':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5
11	000000011	2021-09-27	2021-09-27	900000.00	20.00	720000.00	Mastur	test	8	1	t	0.00	900000.00	'-09':3 '-27':4 '00':6 '000000011':1 '00z':7 '2':24 '2012':21 '2021':2 '4892':17 'ada':15 'e':16 'fif':8,10 'group':9 'gudang':20 'jatibarang':11 'jupiter':19 'r2':25 'roda':23 'stnk':14 'stnk-ada':13 'syaenudin':12 't00':5 'tk':18 'yamaha':22
16	000000016	2022-01-14	2022-01-14	1100000.00	20.00	880000.00	Mastur	test	5	1	t	0.00	1100000.00	'-01':3 '-14':4 '00':6 '000000016':1 '00z':7 '2':25 '2016':22 '2022':2 '3848':18 'ada':16 'beat':20 'e':17 'gudang':21 'honda':23 'jatibarang':12 'kredit':9 'motor':10 'oto':8 'otto':11 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'ub':19
17	000000017	2022-01-14	2022-01-14	750000.00	20.00	600000.00	Mastur	test	10	1	t	0.00	750000.00	'-01':3 '-125':22 '-14':4 '00':6 '000000017':1 '00z':7 '2':27 '2008':24 '2022':2 '3828':18 'ada':16 'company':10 'finance':9 'fw':19 'honda':25 'jatibarang':12 'pusat':23 'r2':28 'roda':26 'stnk':15 'stnk-ada':14 'supra':20 'syaenudin':13 't':17 't00':5 'tfc':11 'top':8 'x':21
25	000000025	2022-03-02	2022-03-02	1000000.00	20.00	800000.00	Mastur	test	2	1	t	0.00	1000000.00	'-02':4 '-03':3 '00':6 '000000025':1 '00z':7 '150':21 '2':26 '2015':23 '2022':2 '3812':18 'ada':16 'adira':11 'auto':8 'b':17 'discret':9 'finance':10 'honda':24 'jatibarang':12,22 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'ujy':19 'vario':20
28	000000028	2022-03-09	2022-03-09	1300000.00	20.00	1040000.00	Mastur	test	2	1	t	0.00	1300000.00	'-03':3 '-09':4 '00':6 '000000028':1 '00z':7 '2':26 '2017':23 '2022':2 '4487':18 'ada':16 'adira':11 'auto':8 'discret':9 'finance':10 'jatibarang':12,22 'mio':20 'pj':19 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 't00':5 'yamaha':24 'z':21
30	000000030	2022-03-12	2022-03-12	1450000.00	20.00	1160000.00	Mastur	test	2	1	t	0.00	1450000.00	'-03':3 '-12':4 '00':6 '000000030':1 '00z':7 '2':25 '2018':22 '2022':2 '3615':18 'ada':16 'adira':11 'auto':8 'discret':9 'finance':10 'honda':23 'jatibarang':12,21 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 't00':5 'verza':20 'zd':19
32	000000032	2022-03-14	2022-03-14	1200000.00	20.00	960000.00	Mastur	test	2	1	t	0.00	1200000.00	'-03':3 '-14':4 '00':6 '000000032':1 '00z':7 '2':25 '2017':22 '2022':2 '2191':18 'ada':16 'adira':11 'auto':8 'beat':20 'discret':9 'finance':10 'honda':23 'jatibarang':12,21 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 't00':5 'ys':19
58	000000058	2022-01-31	2022-01-31	21400000.00	6.54	20000000.00	Mastur	test	4	3	t	0.00	20000000.00	'-01':3 '-31':4 '00':6 '000000058':1 '00z':7 '2012':24 '2022':2 '2281':20 '4':27 'ada':18 'b':11,19 'bekasi':9 'clip':10 'clipan':8 'deddy':14 'indramayu':13 'ios':22 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'sbt':21 'stnk':17 'stnk-ada':16 't00':5 'toyota':25
56	000000056	2022-01-21	2022-01-21	22000000.00	9.09	20000000.00	Mastur	test	18	3	t	0.00	20000000.00	'-01':3 '-21':4 '00':6 '000000056':1 '00z':7 '2018':26 '2022':2 '4':29 '9086':22 'ada':20 'agya':24 'deddy':16 'finance':10 'h':21 'indramayu':15 'mandir':8 'mtf':12 'pranoto':17 'pusat':14,25 'r4':30 'roda':28 's':13 'semarang':11 'stnk':19 'stnk-ada':18 't00':5 'te':23 'toyota':27 'tunas':9
55	000000055	2022-01-17	2022-01-17	38400000.00	0.00	38400000.00	Mastur	test	14	3	t	0.00	38400000.00	'-01':3 '-17':4 '00':6 '000000055':1 '00z':7 '1729':20 '2018':24 '2022':2 '4':27 'ada':18 'bf':21 'brv':22 'clip':10 'clipan':8 'deddy':14 'honda':25 'indramayu':13 'k':11 'karawang':9 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 't':19 't00':5
53	000000053	2021-12-24	2021-12-24	21450000.00	9.09	19500000.00	Mastur	test	14	3	t	0.00	19500000.00	'-12':3 '-24':4 '00':6 '000000053':1 '00z':7 '1312':20 '2007':24 '2021':2 '4':27 'ada':18 'clip':10 'clipan':8 'd':19 'daihatsu':25 'deddy':14 'indramayu':13 'k':11 'karawang':9 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5 'wf':21 'xenia':22
51	000000051	2021-12-21	2021-12-21	25000000.00	8.00	23000000.00	Mastur	test	19	3	t	0.00	23000000.00	'-12':3 '-21':4 '00':6 '000000051':1 '00z':7 '1250':18 '2013':22 '2021':2 '4':25 'ada':16 'clip':9 'clipan':8 'deddy':12 'du':19 'indramayu':11 'ios':20 'pranoto':13 'pusat':10,21 'r4':26 'roda':24 'stnk':15 'stnk-ada':14 't':17 't00':5 'toyota':23
75	000000075	2022-02-07	2022-02-07	1600000.00	20.00	1280000.00	Mastur	test	13	1	t	0.00	1600000.00	'-02':3 '-07':4 '00':6 '000000075':1 '00z':7 '2':24 '2018':21 '2022':2 '4080':16 'ada':14 'e':15 'gapara':8 'jatibarang':10,20 'm3':19 'mio':18 'mpr':9 'r2':25 'roda':23 'stnk':13 'stnk-ada':12 'syaenudin':11 't00':5 'uo':17 'yamaha':22
68	000000068	2021-12-08	2021-12-08	1200000.00	20.00	960000.00	Mastur	test	20	1	t	0.00	1200000.00	'-08':4 '-12':3 '00':6 '000000068':1 '00z':7 '2':24 '2013':21 '2021':2 '3256':17 'ada':15 'b':16 'beat':19 'finance':9 'honda':22 'jatibarang':11,20 'pwy':18 'r2':25 'rad':10 'radana':8 'roda':23 'stnk':14 'stnk-ada':13 'syaenudin':12 't00':5
64	000000064	2021-09-13	2021-09-13	1900000.00	15.79	1600000.00	Mastur	test	7	1	t	0.00	1900000.00	'-09':3 '-13':4 '00':6 '000000064':1 '00z':7 '150':21 '2':26 '2018':23 '2021':2 '5097':18 'ada':16 'finance':10 'honda':24 'jatibarang':12,22 'mandir':8 'muf':11 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 't00':5 'utama':9 'vario':20 'zb':19
27	000000027	2022-03-09	2022-03-09	700000.00	20.00	560000.00	Mastur	test	2	1	t	0.00	700000.00	'-03':3 '-09':4 '00':6 '000000027':1 '00z':7 '2':25 '2014':22 '2022':2 '6819':18 'ada':16 'adira':11 'auto':8 'b':17 'discret':9 'finance':10 'jatibarang':12,21 'pzi':19 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'xeon':20 'yamaha':23
34	000000034	2022-03-17	2022-03-17	900000.00	20.00	720000.00	Mastur	test	2	1	t	0.00	900000.00	'-03':3 '-17':4 '00':6 '000000034':1 '00z':7 '2':25 '2012':22 '2022':2 '4593':18 'ada':16 'adira':11 'auto':8 'discret':9 'e':17 'finance':10 'jatibarang':12,21 'jupiter':20 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'tq':19 'yamaha':23
36	000000036	2022-03-18	2022-03-18	1500000.00	20.00	1200000.00	Mastur	test	1	1	t	0.00	1500000.00	'-03':3 '-18':4 '00':6 '000000036':1 '00z':7 '2':26 '2018':23 '2022':2 '5713':18 'ada':16 'auto':9 'baf':11 'bussan':8 'e':17 'finance':10 'jatibarang':12,22 'm3':21 'mio':20 'pav':19 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'yamaha':24
37	000000037	2021-12-02	2021-12-02	20000000.00	15.00	17000000.00	Mastur	test	14	3	t	0.00	20000000.00	'-02':4 '-12':3 '00':6 '000000037':1 '00z':7 '2006':24 '2021':2 '4':27 '8936':20 'ada':18 'b':19 'clip':10 'clipan':8 'deddy':14 'honda':25 'indramayu':13 'jazz':22 'k':11 'karawang':9 'no':21 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5
39	000000039	2022-01-14	2022-01-14	26000000.00	6.92	24200000.00	Mastur	test	18	3	t	0.00	26000000.00	'-01':3 '-14':4 '00':6 '000000039':1 '00z':7 '1000':25 '2016':27 '2022':2 '4':30 '8715':22 'ada':20 'brio':24 'deddy':16 'finance':10 'gp':23 'h':21 'honda':28 'indramayu':15 'mandir':8 'mtf':12 'pranoto':17 'pusat':14,26 'r4':31 'roda':29 's':13 'semarang':11 'stnk':19 'stnk-ada':18 't00':5 'tunas':9
41	000000041	2022-01-25	2022-01-25	26620000.00	9.09	24200000.00	Mastur	test	14	3	t	0.00	26620000.00	'-01':3 '-25':4 '00':6 '000000041':1 '00z':7 '1788':20 '2017':24 '2022':2 '4':27 'ada':18 'bc':21 'clip':10 'clipan':8 'deddy':14 'honda':25 'indramayu':13 'k':11 'karawang':9 'mobilio':22 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 't':19 't00':5
43	000000043	2022-03-09	2022-03-09	21000000.00	9.52	19000000.00	Mastur	test	16	3	t	0.00	21000000.00	'-03':3 '-09':4 '00':6 '000000043':1 '00z':7 '2020':25 '2022':2 '4':28 '8060':21 'ada':19 'carry':23 'deddy':15 'eg':22 'finance':9 'indramayu':14 'k':12 'karawang':10 'pranoto':16 'pusat':13,24 'r4':29 'roda':27 'safron':8 'sfi':11 'stnk':18 'stnk-ada':17 'suzuk':26 't':20 't00':5
45	000000045	2022-03-14	2022-03-14	8000000.00	20.00	6400000.00	Mastur	test	2	3	t	0.00	8000000.00	'-03':3 '-14':4 '00':6 '000000045':1 '00z':7 '2017':24 '2022':2 '4':27 '8013':20 'ada':18 'adira':11 'auto':8 'deddy':14 'discret':9 'e':19 'finance':10 'indramayu':13 'mitsubish':25 'pickup':22 'pranoto':15 'pusat':12,23 'qa':21 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5
46	000000046	2022-03-14	2022-03-14	23500000.00	10.64	21000000.00	Mastur	test	16	3	t	0.00	23500000.00	'-03':3 '-14':4 '-7':24 '00':6 '000000046':1 '00z':7 '1052':21 '2020':26 '2022':2 '4':29 'ada':19 'deddy':15 'finance':9 'indramayu':14 'k':12 'karawang':10 'pranoto':16 'pusat':13,25 'r4':30 'roda':28 'safron':8 'sfi':11 'stnk':18 'stnk-ada':17 'suzuk':27 't':20 't00':5 'ul':22 'xl':23
48	000000048	2022-03-18	2022-03-18	8000000.00	18.75	6500000.00	Mastur	test	2	3	t	0.00	8000000.00	'-03':3 '-18':4 '00':6 '000000048':1 '00z':7 '2018':24 '2022':2 '4':27 '938':20 'ada':18 'adira':11 'auto':8 'deddy':14 'discret':9 'e':19 'expander':22 'finance':10 'indramayu':13 'mitsubish':25 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5 'xy':21
57	000000057	2022-01-22	2022-01-22	8000000.00	0.00	8000000.00	Mastur	test	2	3	t	0.00	8000000.00	'-01':3 '-22':4 '00':6 '000000057':1 '00z':7 '1256':20 '2019':25 '2022':2 '4':28 'ada':18 'adira':11 'auto':8 'daihatsu':26 'deddy':14 'discret':9 'e':19 'finance':10 'grand':22 'indramayu':13 'max':23 'pranoto':15 'pusat':12,24 'qd':21 'r4':29 'roda':27 'stnk':17 'stnk-ada':16 't00':5
54	000000054	2021-12-28	2021-12-28	5400000.00	0.00	5400000.00	Mastur	test	2	3	t	0.00	5400000.00	'-12':3 '-28':4 '00':6 '000000054':1 '00z':7 '1242':20 '2012':24 '2021':2 '4':27 'ada':18 'adira':11 'auto':8 'd':19 'deddy':14 'discret':9 'finance':10 'indramayu':13 'ou':21 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5 'toyota':25 'vios':22
52	000000052	2021-12-23	2021-12-23	15000000.00	7.00	13950000.00	Mastur	test	14	3	t	0.00	13950000.00	'-12':3 '-23':4 '00':6 '000000052':1 '00z':7 '1000':23 '1184':20 '2018':25 '2021':2 '4':28 'ada':18 'brio':22 'clip':10 'clipan':8 'deddy':14 'ga':21 'honda':26 'indramayu':13 'k':11 'karawang':9 'pranoto':15 'pusat':12,24 'r4':29 'roda':27 'stnk':17 'stnk-ada':16 't':19 't00':5
50	000000050	2021-12-16	2021-12-16	34500000.00	9.10	31360000.00	Mastur	test	18	3	t	0.00	31360000.00	'-12':3 '-16':4 '00':6 '000000050':1 '00z':7 '1000':25 '2017':27 '2021':2 '4':30 '9442':22 'ada':20 'brio':24 'deddy':16 'finance':10 'h':21 'honda':28 'indramayu':15 'mandir':8 'mtf':12 'ng':23 'pranoto':17 'pusat':14,26 'r4':31 'roda':29 's':13 'semarang':11 'stnk':19 'stnk-ada':18 't00':5 'tunas':9
66	000000066	2021-12-06	2021-12-06	1300000.00	20.00	1040000.00	Mastur	test	12	1	t	0.00	1300000.00	'-06':4 '-12':3 '00':6 '000000066':1 '00z':7 '2':24 '2017':21 '2021':2 '6934':17 'ada':15 'beat':19 'finance':9 'honda':22 'jatibarang':11,20 'r2':25 'roda':23 'stnk':14 'stnk-ada':13 'syaenudin':12 't':16 't00':5 'wom':8 'womf':10 'yq':18
62	000000062	2022-01-19	2022-01-19	3600000.00	20.00	2880000.00	Mastur	test	1	3	t	0.00	3600000.00	'-01':3 '-19':4 '00':6 '000000062':1 '00z':7 '2':27 '2017':24 '2022':2 '3217':20 'ada':18 'auto':9 'baf':11 'bussan':8 'deddy':14 'e':19 'finance':10 'indramayu':13 'nmax':22 'par':21 'pranoto':15 'pusat':12,23 'r2':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5 'yamaha':25
59	000000059	2022-01-31	2022-01-31	19750000.00	0.00	19750000.00	Mastur	test	14	3	t	0.00	19750000.00	'-01':3 '-31':4 '00':6 '000000059':1 '00z':7 '1305':20 '2000':24 '2022':2 '4':27 'ada':18 'clip':10 'clipan':8 'deddy':14 'dl':21 'ertiga':22 'indramayu':13 'k':11 'karawang':9 'pranoto':15 'pusat':12,23 'r4':28 'roda':26 'stnk':17 'stnk-ada':16 'suzuk':25 't':19 't00':5
49	000000049	2021-11-22	2021-11-22	10000000.00	18.00	8200000.00	Mastur	test	17	3	t	0.00	8200000.00	'-11':3 '-22':4 '00':6 '000000049':1 '00z':7 '2012':23 '2021':2 '4':26 '8630':19 'ada':17 'bfi':8,10 'deddy':13 'finance':9 'h':18 'honda':24 'indramayu':12 'jazz':21 'pp':20 'pranoto':14 'pusat':11,22 'r4':27 'roda':25 'stnk':16 'stnk-ada':15 't00':5
74	000000074	2022-03-07	2022-03-07	1300000.00	58.46	540000.00	Mastur	test	2	1	t	0.00	1300000.00	'-03':3 '-07':4 '00':6 '000000074':1 '00z':7 '2':25 '2018':22 '2022':2 '6871':18 'ada':16 'adira':11 'auto':8 'cm':19 'discret':9 'e':17 'finance':10 'jatibarang':12,21 'mio':20 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'yamaha':23
73	000000073	2021-12-30	2021-12-30	1250000.00	20.00	1000000.00	Mastur	test	2	1	t	0.00	1250000.00	'-12':3 '-30':4 '00':6 '000000073':1 '00z':7 '2':26 '2015':23 '2021':2 '3351':18 'ada':16 'adira':11 'auto':8 'b':17 'beat':20 'discret':9 'finance':10 'honda':24 'jatibarang':12,22 'kuh':19 'pop':21 'r2':27 'roda':25 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5
72	000000072	2021-12-09	2021-12-09	1250000.00	20.00	1000000.00	Mastur	test	2	1	t	0.00	1250000.00	'-09':4 '-12':3 '00':6 '000000072':1 '00z':7 '2':25 '2016':22 '2021':2 '3430':18 'ada':16 'adira':11 'auto':8 'b':17 'beat':20 'discret':9 'ejx':19 'finance':10 'honda':23 'jatibarang':12,21 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5
71	000000071	2021-09-21	2021-09-21	1300000.00	20.00	1040000.00	Mastur	test	7	1	t	0.00	1300000.00	'-09':3 '-21':4 '00':6 '000000071':1 '00z':7 '2':27 '2017':24 '2021':2 '4261':18 'ada':16 'finance':10 'jatibarang':12,23 'king':22 'mandir':8 'muf':11 'mx':21 'mx-king':20 'r2':28 'roda':26 'stnk':15 'stnk-ada':14 'syaenudin':13 't':17 't00':5 'utama':9 'yamaha':25 'yq':19
70	000000070	2022-01-03	2022-01-03	1400000.00	20.00	1120000.00	Mastur	test	5	1	t	0.00	1400000.00	'-01':3 '-03':4 '00':6 '000000070':1 '00z':7 '2':25 '2018':22 '2022':2 '4654':18 'ada':16 'b':17 'beat':20 'fsg':19 'honda':23 'jatibarang':12,21 'kredit':9 'motor':10 'oto':8 'otto':11 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5
69	000000069	2021-12-09	2021-12-09	1400000.00	20.00	1120000.00	Mastur	test	2	1	t	0.00	1400000.00	'-09':4 '-12':3 '00':6 '000000069':1 '00z':7 '2':25 '2019':22 '2021':2 '3433':18 'ada':16 'adira':11 'auto':8 'b':17 'beat':20 'discret':9 'finance':10 'honda':23 'jatibarang':12,21 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'usn':19
67	000000067	2021-12-07	2021-12-07	2000000.00	20.00	1600000.00	Mastur	test	21	1	t	0.00	2000000.00	'-07':4 '-12':3 '00':6 '000000067':1 '00z':7 '2':26 '2020':23 '2021':2 '4845':19 'ada':17 'auto':9 'central':10 'finance':11 'iq':20 'jatibarang':13,22 'macf':12 'mega':8 'nmax':21 'r2':27 'roda':25 'stnk':16 'stnk-ada':15 'syaenudin':14 't':18 't00':5 'yamaha':24
65	000000065	2021-10-05	2021-10-05	3200000.00	0.00	3200000.00	Mastur	test	1	1	t	0.00	3200000.00	'-05':4 '-10':3 '00':6 '000000065':1 '00z':7 '2':25 '2019':22 '2021':2 '2676':18 'ada':16 'auto':9 'baf':11 'bussan':8 'e':17 'finance':10 'jatibarang':12,21 'nmax':20 'r2':26 'roda':24 'stnk':15 'stnk-ada':14 'syaenudin':13 't00':5 'ux':19 'yamaha':23
63	000000063	2021-12-24	2021-12-24	1800000.00	20.00	1440000.00	Mastur	test	5	3	t	0.00	1800000.00	'-12':3 '-24':4 '00':6 '000000063':1 '00z':7 '2':27 '2019':24 '2021':2 '2113':20 'ada':18 'deddy':14 'e':19 'honda':25 'indramayu':13 'kredit':9 'motor':10 'oto':8 'otto':11 'pbm':21 'pcx':22 'pranoto':15 'pusat':12,23 'r2':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5
61	000000061	2021-12-29	2021-12-29	1500000.00	20.00	1200000.00	Mastur	test	7	3	t	0.00	1500000.00	'-12':3 '-29':4 '00':6 '000000061':1 '00z':7 '2':27 '2015':24 '2021':2 '3310':20 'ada':18 'deddy':14 'e':19 'finance':10 'indramayu':13 'mandir':8 'muf':11 'pranoto':15 'pusat':12,23 'qr':21 'r2':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5 'utama':9 'vixion':22 'yamaha':25
60	000000060	2021-09-30	2021-09-30	920000.00	0.00	920000.00	Mastur	test	2	3	t	0.00	920000.00	'-09':3 '-30':4 '00':6 '000000060':1 '00z':7 '2':27 '2018':24 '2021':2 '6181':20 'ada':18 'adira':11 'auto':8 'beat':22 'deddy':14 'discret':9 'f':19 'fch':21 'finance':10 'honda':25 'indramayu':13 'pranoto':15 'pusat':12,23 'r2':28 'roda':26 'stnk':17 'stnk-ada':16 't00':5
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
83	12	TRX-Order	Piutang jasa COLLECTIUS (COL) Order SPK: /000000012	2022-03-22	Kendaraan R2 Yamaha Mio M3 , Nopol B 3479 UJU	'/000000012':1 '/ref-12':13 '3479':7 'b':6 'col':10 'collectius':9 'jatibarang':11 'm3':5 'mio':4 'order':16 'r2':2 'syaenudin':12 'trx':15 'trx-order':14 'uju':8 'yamaha':3
72	1	TRX-Order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000001	2022-03-22	Kendaraan R2 Yamaha Mio Z , Nopol E 5605 PAS	'/000000001':1 '/ref-1':17 '5605':7 'deddy':15 'e':6 'indramayu':14 'kredit':10 'mio':4 'motor':11 'order':20 'oto':9 'otto':12 'pas':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'yamaha':3 'z':5
73	2	TRX-Order	Piutang jasa COLLECTIUS (COL) Order SPK: /000000002	2022-03-22	Kendaraan R2 Yamaha Mio , Nopol E 3977 PAC	'/000000002':1 '/ref-2':14 '3977':6 'col':9 'collectius':8 'deddy':12 'e':5 'indramayu':11 'mio':4 'order':17 'pac':7 'pranoto':13 'pusat':10 'r2':2 'trx':16 'trx-order':15 'yamaha':3
74	3	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000003	2022-03-22	Kendaraan R2 Yamaha Mio M3 , Nopol E 5125 PBC	'/000000003':1 '/ref-3':17 '5125':7 'auto':10 'baf':12 'bussan':9 'deddy':15 'e':6 'finance':11 'indramayu':14 'm3':5 'mio':4 'order':20 'pbc':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'yamaha':3
75	4	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000004	2022-03-22	Kendaraan R2 Yamaha Vixion , Nopol BR 5080 PY	'/000000004':1 '/ref-4':16 '5080':6 'adira':11 'auto':8 'br':5 'deddy':14 'discret':9 'finance':10 'indramayu':13 'order':19 'pranoto':15 'pusat':12 'py':7 'r2':2 'trx':18 'trx-order':17 'vixion':4 'yamaha':3
76	5	TRX-Order	Piutang jasa Mandiri Utama Finance (MUF) Order SPK: /000000005	2022-03-22	Kendaraan R2 Yamaha Fino 125  , Nopol E 4096 PAQ	'/000000005':1 '/ref-5':17 '125':5 '4096':7 'deddy':15 'e':6 'finance':11 'fino':4 'indramayu':14 'mandir':9 'muf':12 'order':20 'paq':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'utama':10 'yamaha':3
77	6	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000006	2022-03-22	Kendaraan R2 Honda BEAT , Nopol E 2633 PAC	'/000000006':1 '/ref-6':16 '2633':6 'adira':11 'auto':8 'beat':4 'deddy':14 'discret':9 'e':5 'finance':10 'honda':3 'indramayu':13 'order':19 'pac':7 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17
78	7	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000007	2022-03-22	Kendaraan R2 Yamaha Fino 125  , Nopol E 2033 PBJ	'/000000007':1 '/ref-7':17 '125':5 '2033':7 'auto':10 'baf':12 'bussan':9 'deddy':15 'e':6 'finance':11 'fino':4 'indramayu':14 'order':20 'pbj':8 'pranoto':16 'pusat':13 'r2':2 'trx':19 'trx-order':18 'yamaha':3
79	8	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000008	2022-03-22	Kendaraan R2 Yamaha R-15 , Nopol E 6277 PAZ	'-15':5 '/000000008':1 '/ref-8':17 '6277':7 'adira':12 'auto':9 'deddy':15 'discret':10 'e':6 'finance':11 'indramayu':14 'order':20 'paz':8 'pranoto':16 'pusat':13 'r':4 'r2':2 'trx':19 'trx-order':18 'yamaha':3
80	9	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000009	2022-03-22	Kendaraan R2 Honda BEAT , Nopol E 6053 PAM	'/000000009':1 '/ref-9':16 '6053':6 'adira':11 'auto':8 'beat':4 'deddy':14 'discret':9 'e':5 'finance':10 'honda':3 'indramayu':13 'order':19 'pam':7 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17
81	10	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000010	2022-03-22	Kendaraan R2 Honda BEAT , Nopol E 5474 Q	'/000000010':1 '/ref-10':16 '5474':6 'adira':11 'auto':8 'beat':4 'deddy':14 'discret':9 'e':5 'finance':10 'honda':3 'indramayu':13 'order':19 'pranoto':15 'pusat':12 'q':7 'r2':2 'trx':18 'trx-order':17
82	11	TRX-Order	Piutang jasa FIF Group (FIF) Order SPK: /000000011	2022-03-22	Kendaraan R2 Yamaha Jupiter , Nopol E 4892 TK	'/000000011':1 '/ref-11':13 '4892':6 'e':5 'fif':8,10 'group':9 'jatibarang':11 'jupiter':4 'order':16 'r2':2 'syaenudin':12 'tk':7 'trx':15 'trx-order':14 'yamaha':3
84	13	TRX-Order	Piutang jasa COLLECTIUS (COL) Order SPK: /000000013	2022-03-22	Kendaraan R2 Yamaha Mio M3 , Nopol E 2417 PAO	'/000000013':1 '/ref-13':13 '2417':7 'col':10 'collectius':9 'e':6 'jatibarang':11 'm3':5 'mio':4 'order':16 'pao':8 'r2':2 'syaenudin':12 'trx':15 'trx-order':14 'yamaha':3
85	14	TRX-Order	Piutang jasa Mitra Pinasthika Mustika Finance (MPMF) Order SPK: /000000014	2022-03-22	Kendaraan R2 Suzuki Satria FU , Nopol T 3521 KL	'/000000014':1 '/ref-14':16 '3521':7 'finance':12 'fu':5 'jatibarang':14 'kl':8 'mitra':9 'mpmf':13 'mustika':11 'order':19 'pinasthika':10 'r2':2 'satria':4 'suzuk':3 'syaenudin':15 't':6 'trx':18 'trx-order':17
86	15	TRX-Order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000015	2022-03-22	Kendaraan R2 Yamaha Jupiter MX , Nopol T 4146 KO	'/000000015':1 '/ref-15':15 '4146':7 'jatibarang':13 'jupiter':4 'ko':8 'kredit':10 'motor':11 'mx':5 'order':18 'oto':9 'otto':12 'r2':2 'syaenudin':14 't':6 'trx':17 'trx-order':16 'yamaha':3
87	16	TRX-Order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000016	2022-03-22	Kendaraan R2 Honda BEAT , Nopol E 3848 UB	'/000000016':1 '/ref-16':14 '3848':6 'beat':4 'e':5 'honda':3 'jatibarang':12 'kredit':9 'motor':10 'order':17 'oto':8 'otto':11 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'ub':7
88	17	TRX-Order	Piutang jasa Top Finance Company (TFC) Order SPK: /000000017	2022-03-22	Kendaraan R2 Honda Supra X-125 , Nopol T 3828 FW	'-125':6 '/000000017':1 '/ref-17':16 '3828':8 'company':12 'finance':11 'fw':9 'honda':3 'jatibarang':14 'order':19 'r2':2 'supra':4 'syaenudin':15 't':7 'tfc':13 'top':10 'trx':18 'trx-order':17 'x':5
89	18	TRX-Order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000018	2022-03-22	Kendaraan R2 Yamaha Mio M3 , Nopol E 6716 IX	'/000000018':1 '/ref-18':15 '6716':7 'e':6 'ix':8 'jatibarang':13 'kredit':10 'm3':5 'mio':4 'motor':11 'order':18 'oto':9 'otto':12 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
90	19	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000019	2022-03-22	Kendaraan R2 Honda Revo , Nopol E 5638 PAV	'/000000019':1 '/ref-19':14 '5638':6 'adira':11 'auto':8 'discret':9 'e':5 'finance':10 'honda':3 'jatibarang':12 'order':17 'pav':7 'r2':2 'revo':4 'syaenudin':13 'trx':16 'trx-order':15
91	20	TRX-Order	Piutang jasa Kredit Plus (KP+) Order SPK: /000000020	2022-03-22	Kendaraan R2 Honda Vario 125 , Nopol E 5253 TY	'/000000020':1 '/ref-20':14 '125':5 '5253':7 'e':6 'honda':3 'jatibarang':12 'kp':11 'kredit':9 'order':17 'plus':10 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'ty':8 'vario':4
92	21	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000021	2022-03-22	Kendaraan R2 Yamaha Mio M3 , Nopol B 6262 VKY	'/000000021':1 '/ref-21':15 '6262':7 'auto':10 'b':6 'baf':12 'bussan':9 'finance':11 'jatibarang':13 'm3':5 'mio':4 'order':18 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'vky':8 'yamaha':3
94	24	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000024	2022-03-22	Kendaraan R2 Yamaha Mio S , Nopol E 2146 QAF	'/000000024':1 '/ref-24':15 '2146':7 'auto':10 'baf':12 'bussan':9 'e':6 'finance':11 'jatibarang':13 'mio':4 'order':18 'qaf':8 'r2':2 's':5 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
97	27	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000027	2022-03-22	Kendaraan R2 Yamaha Xeon , Nopol B 6819 PZI	'/000000027':1 '/ref-27':14 '6819':6 'adira':11 'auto':8 'b':5 'discret':9 'finance':10 'jatibarang':12 'order':17 'pzi':7 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'xeon':4 'yamaha':3
98	28	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000028	2022-03-22	Kendaraan R2 Yamaha Mio Z , Nopol T 4487 PJ	'/000000028':1 '/ref-28':15 '4487':7 'adira':12 'auto':9 'discret':10 'finance':11 'jatibarang':13 'mio':4 'order':18 'pj':8 'r2':2 'syaenudin':14 't':6 'trx':17 'trx-order':16 'yamaha':3 'z':5
99	29	TRX-Order	Piutang jasa MEGAPARA (MPR) Order SPK: /000000029	2022-03-22	Kendaraan R2 Yamaha Mio , Nopol E 4544 JD	'/000000029':1 '/ref-29':12 '4544':6 'e':5 'gapara':8 'jatibarang':10 'jd':7 'mio':4 'mpr':9 'order':15 'r2':2 'syaenudin':11 'trx':14 'trx-order':13 'yamaha':3
102	32	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000032	2022-03-22	Kendaraan R2 Honda BEAT , Nopol T 2191 YS	'/000000032':1 '/ref-32':14 '2191':6 'adira':11 'auto':8 'beat':4 'discret':9 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 't':5 'trx':16 'trx-order':15 'ys':7
93	22	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000022	2022-03-22	Kendaraan R2 Yamaha Mio M3 , Nopol E 2830 QR	'/000000022':1 '/ref-22':15 '2830':7 'auto':10 'baf':12 'bussan':9 'e':6 'finance':11 'jatibarang':13 'm3':5 'mio':4 'order':18 'qr':8 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
95	25	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000025	2022-03-22	Kendaraan R2 Honda Vario 150 , Nopol B 3812 UJY	'/000000025':1 '/ref-25':15 '150':5 '3812':7 'adira':12 'auto':9 'b':6 'discret':10 'finance':11 'honda':3 'jatibarang':13 'order':18 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'ujy':8 'vario':4
100	30	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000030	2022-03-22	Kendaraan R2 Honda Verza , Nopol T 3615 ZD	'/000000030':1 '/ref-30':14 '3615':6 'adira':11 'auto':8 'discret':9 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 't':5 'trx':16 'trx-order':15 'verza':4 'zd':7
101	31	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000031	2022-03-22	Kendaraan R2 Yamaha R-15 , Nopol E 2391 JM	'-15':5 '/000000031':1 '/ref-31':15 '2391':7 'adira':12 'auto':9 'discret':10 'e':6 'finance':11 'jatibarang':13 'jm':8 'order':18 'r':4 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
105	35	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000035	2022-03-22	Kendaraan R2 Yamaha R-15 , Nopol T 2110 YV	'-15':5 '/000000035':1 '/ref-35':15 '2110':7 'adira':12 'auto':9 'discret':10 'finance':11 'jatibarang':13 'order':18 'r':4 'r2':2 'syaenudin':14 't':6 'trx':17 'trx-order':16 'yamaha':3 'yv':8
96	26	TRX-Order	Piutang jasa Kredit Plus (KP+) Order SPK: /000000026	2022-03-22	Kendaraan R2 Honda Vario 150 , Nopol T 2891 WP	'/000000026':1 '/ref-26':14 '150':5 '2891':7 'honda':3 'jatibarang':12 'kp':11 'kredit':9 'order':17 'plus':10 'r2':2 'syaenudin':13 't':6 'trx':16 'trx-order':15 'vario':4 'wp':8
103	33	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000033	2022-03-22	Kendaraan R2 Honda GENIO , Nopol T 5856 ZT	'/000000033':1 '/ref-33':14 '5856':6 'adira':11 'auto':8 'discret':9 'finance':10 'genio':4 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 't':5 'trx':16 'trx-order':15 'zt':7
104	34	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000034	2022-03-22	Kendaraan R2 Yamaha Jupiter , Nopol E 4593 TQ	'/000000034':1 '/ref-34':14 '4593':6 'adira':11 'auto':8 'discret':9 'e':5 'finance':10 'jatibarang':12 'jupiter':4 'order':17 'r2':2 'syaenudin':13 'tq':7 'trx':16 'trx-order':15 'yamaha':3
106	36	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000036	2022-03-22	Kendaraan R2 Yamaha Mio M3 , Nopol E 5713 PAV	'/000000036':1 '/ref-36':15 '5713':7 'auto':10 'baf':12 'bussan':9 'e':6 'finance':11 'jatibarang':13 'm3':5 'mio':4 'order':18 'pav':8 'r2':2 'syaenudin':14 'trx':17 'trx-order':16 'yamaha':3
107	37	TRX-Order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000037	2022-03-22	Kendaraan R4 Honda Jazz , Nopol B 8936 NO	'/000000037':1 '/ref-37':16 '8936':6 'b':5 'clip':10 'clipan':8 'deddy':14 'honda':3 'indramayu':13 'jazz':4 'k':11 'karawang':9 'no':7 'order':19 'pranoto':15 'pusat':12 'r4':2 'trx':18 'trx-order':17
108	38	TRX-Order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000038	2022-03-22	Kendaraan R4 Suzuki Carry , Nopol T 1412 KM	'/000000038':1 '/ref-38':16 '1412':6 'carry':4 'clip':10 'clipan':8 'deddy':14 'indramayu':13 'k':11 'karawang':9 'km':7 'order':19 'pranoto':15 'pusat':12 'r4':2 'suzuk':3 't':5 'trx':18 'trx-order':17
109	39	TRX-Order	Piutang jasa Mandiri Tunas Finance Semarang (MTF S) Order SPK: /000000039	2022-03-22	Kendaraan R4 Honda Brio 1000 , Nopol H 8715 GP	'/000000039':1 '/ref-39':19 '1000':5 '8715':7 'brio':4 'deddy':17 'finance':11 'gp':8 'h':6 'honda':3 'indramayu':16 'mandir':9 'mtf':13 'order':22 'pranoto':18 'pusat':15 'r4':2 's':14 'semarang':12 'trx':21 'trx-order':20 'tunas':10
110	40	TRX-Order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000040	2022-03-22	Kendaraan R4 Suzuki ER-3 , Nopol T 1164 FQ	'-3':5 '/000000040':1 '/ref-40':17 '1164':7 'clip':11 'clipan':9 'deddy':15 'er':4 'fq':8 'indramayu':14 'k':12 'karawang':10 'order':20 'pranoto':16 'pusat':13 'r4':2 'suzuk':3 't':6 'trx':19 'trx-order':18
111	41	TRX-Order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000041	2022-03-22	Kendaraan R4 Honda Mobilio , Nopol T 1788 BC	'/000000041':1 '/ref-41':16 '1788':6 'bc':7 'clip':10 'clipan':8 'deddy':14 'honda':3 'indramayu':13 'k':11 'karawang':9 'mobilio':4 'order':19 'pranoto':15 'pusat':12 'r4':2 't':5 'trx':18 'trx-order':17
112	42	TRX-Order	Piutang jasa Mandiri Tunas Finance Semarang (MTF S) Order SPK: /000000042	2022-03-22	Kendaraan R4 Honda Brio 1000 , Nopol H 9049 SE	'/000000042':1 '/ref-42':19 '1000':5 '9049':7 'brio':4 'deddy':17 'finance':11 'h':6 'honda':3 'indramayu':16 'mandir':9 'mtf':13 'order':22 'pranoto':18 'pusat':15 'r4':2 's':14 'se':8 'semarang':12 'trx':21 'trx-order':20 'tunas':10
113	43	TRX-Order	Piutang jasa Safron Finance Karawang (SFI K) Order SPK: /000000043	2022-03-22	Kendaraan R4 Suzuki Carry , Nopol T 8060 EG	'/000000043':1 '/ref-43':17 '8060':6 'carry':4 'deddy':15 'eg':7 'finance':9 'indramayu':14 'k':12 'karawang':10 'order':20 'pranoto':16 'pusat':13 'r4':2 'safron':8 'sfi':11 'suzuk':3 't':5 'trx':19 'trx-order':18
114	44	TRX-Order	Piutang jasa BFI Finance (BFI) Order SPK: /000000044	2022-03-22	Kendaraan R4 Suzuki Carry , Nopol E 8903 PP	'/000000044':1 '/ref-44':15 '8903':6 'bfi':8,10 'carry':4 'deddy':13 'e':5 'finance':9 'indramayu':12 'order':18 'pp':7 'pranoto':14 'pusat':11 'r4':2 'suzuk':3 'trx':17 'trx-order':16
115	45	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000045	2022-03-22	Kendaraan R4 Mitsubishi Pickup , Nopol E 8013 QA	'/000000045':1 '/ref-45':16 '8013':6 'adira':11 'auto':8 'deddy':14 'discret':9 'e':5 'finance':10 'indramayu':13 'mitsubish':3 'order':19 'pickup':4 'pranoto':15 'pusat':12 'qa':7 'r4':2 'trx':18 'trx-order':17
116	46	TRX-Order	Piutang jasa Safron Finance Karawang (SFI K) Order SPK: /000000046	2022-03-22	Kendaraan R4 Suzuki XL-7 , Nopol T 1052 UL	'-7':5 '/000000046':1 '/ref-46':18 '1052':7 'deddy':16 'finance':10 'indramayu':15 'k':13 'karawang':11 'order':21 'pranoto':17 'pusat':14 'r4':2 'safron':9 'sfi':12 'suzuk':3 't':6 'trx':20 'trx-order':19 'ul':8 'xl':4
117	47	TRX-Order	Piutang jasa Clipan Palembang (CLIP P) Order SPK: /000000047	2022-03-22	Kendaraan R4 Honda Jazz , Nopol BG 1623 PF	'/000000047':1 '/ref-47':16 '1623':6 'bg':5 'clip':10 'clipan':8 'deddy':14 'honda':3 'indramayu':13 'jazz':4 'order':19 'p':11 'palembang':9 'pf':7 'pranoto':15 'pusat':12 'r4':2 'trx':18 'trx-order':17
118	48	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000048	2022-03-22	Kendaraan R4 Mitsubishi Expander , Nopol E 938 XY	'/000000048':1 '/ref-48':16 '938':6 'adira':11 'auto':8 'deddy':14 'discret':9 'e':5 'expander':4 'finance':10 'indramayu':13 'mitsubish':3 'order':19 'pranoto':15 'pusat':12 'r4':2 'trx':18 'trx-order':17 'xy':7
119	59	TRX-Order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000059	2022-03-22	Kendaraan R4 Suzuki ERTIGA , Nopol T 1305 DL	'/000000059':1 '/ref-59':16 '1305':6 'clip':10 'clipan':8 'deddy':14 'dl':7 'ertiga':4 'indramayu':13 'k':11 'karawang':9 'order':19 'pranoto':15 'pusat':12 'r4':2 'suzuk':3 't':5 'trx':18 'trx-order':17
120	58	TRX-Order	Piutang jasa Clipan Bekasi (CLIP B) Order SPK: /000000058	2022-03-22	Kendaraan R4 Toyota Terios , Nopol B 2281 SBT	'/000000058':1 '/ref-58':16 '2281':6 'b':5,11 'bekasi':9 'clip':10 'clipan':8 'deddy':14 'indramayu':13 'ios':4 'order':19 'pranoto':15 'pusat':12 'r4':2 'sbt':7 'toyota':3 'trx':18 'trx-order':17
121	57	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000057	2022-03-22	Kendaraan R4 Daihatsu Grand Max , Nopol E 1256 QD	'/000000057':1 '/ref-57':17 '1256':7 'adira':12 'auto':9 'daihatsu':3 'deddy':15 'discret':10 'e':6 'finance':11 'grand':4 'indramayu':14 'max':5 'order':20 'pranoto':16 'pusat':13 'qd':8 'r4':2 'trx':19 'trx-order':18
122	56	TRX-Order	Piutang jasa Mandiri Tunas Finance Semarang (MTF S) Order SPK: /000000056	2022-03-22	Kendaraan R4 Toyota AGYA , Nopol H 9086 TE	'/000000056':1 '/ref-56':18 '9086':6 'agya':4 'deddy':16 'finance':10 'h':5 'indramayu':15 'mandir':8 'mtf':12 'order':21 'pranoto':17 'pusat':14 'r4':2 's':13 'semarang':11 'te':7 'toyota':3 'trx':20 'trx-order':19 'tunas':9
123	55	TRX-Order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000055	2022-03-22	Kendaraan R4 Honda BRV , Nopol T 1729 BF	'/000000055':1 '/ref-55':16 '1729':6 'bf':7 'brv':4 'clip':10 'clipan':8 'deddy':14 'honda':3 'indramayu':13 'k':11 'karawang':9 'order':19 'pranoto':15 'pusat':12 'r4':2 't':5 'trx':18 'trx-order':17
124	54	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000054	2022-03-22	Kendaraan R4 Toyota Vios , Nopol D 1242 OU	'/000000054':1 '/ref-54':16 '1242':6 'adira':11 'auto':8 'd':5 'deddy':14 'discret':9 'finance':10 'indramayu':13 'order':19 'ou':7 'pranoto':15 'pusat':12 'r4':2 'toyota':3 'trx':18 'trx-order':17 'vios':4
125	53	TRX-Order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000053	2022-03-22	Kendaraan R4 Daihatsu Xenia , Nopol D 1312 WF	'/000000053':1 '/ref-53':16 '1312':6 'clip':10 'clipan':8 'd':5 'daihatsu':3 'deddy':14 'indramayu':13 'k':11 'karawang':9 'order':19 'pranoto':15 'pusat':12 'r4':2 'trx':18 'trx-order':17 'wf':7 'xenia':4
126	52	TRX-Order	Piutang jasa Clipan Karawang\n (CLIP K) Order SPK: /000000052	2022-03-22	Kendaraan R4 Honda Brio 1000 , Nopol T 1184 GA	'/000000052':1 '/ref-52':17 '1000':5 '1184':7 'brio':4 'clip':11 'clipan':9 'deddy':15 'ga':8 'honda':3 'indramayu':14 'k':12 'karawang':10 'order':20 'pranoto':16 'pusat':13 'r4':2 't':6 'trx':19 'trx-order':18
127	51	TRX-Order	Piutang jasa CLIPAN (CLIP) Order SPK: /000000051	2022-03-22	Kendaraan R4 Toyota Terios , Nopol T 1250 DU	'/000000051':1 '/ref-51':14 '1250':6 'clip':9 'clipan':8 'deddy':12 'du':7 'indramayu':11 'ios':4 'order':17 'pranoto':13 'pusat':10 'r4':2 't':5 'toyota':3 'trx':16 'trx-order':15
128	50	TRX-Order	Piutang jasa Mandiri Tunas Finance Semarang (MTF S) Order SPK: /000000050	2022-03-22	Kendaraan R4 Honda Brio 1000 , Nopol H 9442 NG	'/000000050':1 '/ref-50':19 '1000':5 '9442':7 'brio':4 'deddy':17 'finance':11 'h':6 'honda':3 'indramayu':16 'mandir':9 'mtf':13 'ng':8 'order':22 'pranoto':18 'pusat':15 'r4':2 's':14 'semarang':12 'trx':21 'trx-order':20 'tunas':10
129	49	TRX-Order	Piutang jasa BFI Finance (BFI) Order SPK: /000000049	2022-03-22	Kendaraan R4 Honda Jazz , Nopol H 8630 PP	'/000000049':1 '/ref-49':15 '8630':6 'bfi':8,10 'deddy':13 'finance':9 'h':5 'honda':3 'indramayu':12 'jazz':4 'order':18 'pp':7 'pranoto':14 'pusat':11 'r4':2 'trx':17 'trx-order':16
130	75	TRX-Order	Piutang jasa MEGAPARA (MPR) Order SPK: /000000075	2022-03-22	Kendaraan R2 Yamaha Mio M3 , Nopol E 4080 UO	'/000000075':1 '/ref-75':13 '4080':7 'e':6 'gapara':9 'jatibarang':11 'm3':5 'mio':4 'mpr':10 'order':16 'r2':2 'syaenudin':12 'trx':15 'trx-order':14 'uo':8 'yamaha':3
131	74	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000074	2022-03-22	Kendaraan R2 Yamaha Mio , Nopol E 6871 CM	'/000000074':1 '/ref-74':14 '6871':6 'adira':11 'auto':8 'cm':7 'discret':9 'e':5 'finance':10 'jatibarang':12 'mio':4 'order':17 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'yamaha':3
132	73	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000073	2022-03-22	Kendaraan R2 Honda Beat Pop , Nopol B 3351 KUH	'/000000073':1 '/ref-73':15 '3351':7 'adira':12 'auto':9 'b':6 'beat':4 'discret':10 'finance':11 'honda':3 'jatibarang':13 'kuh':8 'order':18 'pop':5 'r2':2 'syaenudin':14 'trx':17 'trx-order':16
133	72	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000072	2022-03-22	Kendaraan R2 Honda BEAT , Nopol B 3430 EJX	'/000000072':1 '/ref-72':14 '3430':6 'adira':11 'auto':8 'b':5 'beat':4 'discret':9 'ejx':7 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 'trx':16 'trx-order':15
134	71	TRX-Order	Piutang jasa Mandiri Utama Finance (MUF) Order SPK: /000000071	2022-03-22	Kendaraan R2 Yamaha MX-King , Nopol T 4261 YQ	'/000000071':1 '/ref-71':16 '4261':8 'finance':12 'jatibarang':14 'king':6 'mandir':10 'muf':13 'mx':5 'mx-king':4 'order':19 'r2':2 'syaenudin':15 't':7 'trx':18 'trx-order':17 'utama':11 'yamaha':3 'yq':9
135	70	TRX-Order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000070	2022-03-22	Kendaraan R2 Honda BEAT , Nopol B 4654 FSG	'/000000070':1 '/ref-70':14 '4654':6 'b':5 'beat':4 'fsg':7 'honda':3 'jatibarang':12 'kredit':9 'motor':10 'order':17 'oto':8 'otto':11 'r2':2 'syaenudin':13 'trx':16 'trx-order':15
136	69	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000069	2022-03-22	Kendaraan R2 Honda BEAT , Nopol B 3433 USN	'/000000069':1 '/ref-69':14 '3433':6 'adira':11 'auto':8 'b':5 'beat':4 'discret':9 'finance':10 'honda':3 'jatibarang':12 'order':17 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'usn':7
137	68	TRX-Order	Piutang jasa Radana Finance (RAD) Order SPK: /000000068	2022-03-22	Kendaraan R2 Honda BEAT , Nopol B 3256  PWY	'/000000068':1 '/ref-68':13 '3256':6 'b':5 'beat':4 'finance':9 'honda':3 'jatibarang':11 'order':16 'pwy':7 'r2':2 'rad':10 'radana':8 'syaenudin':12 'trx':15 'trx-order':14
138	67	TRX-Order	Piutang jasa Mega Auto Central Finance (MACF) Order SPK: /000000067	2022-03-22	Kendaraan R2 Yamaha NMax , Nopol T 4845 IQ	'/000000067':1 '/ref-67':15 '4845':6 'auto':9 'central':10 'finance':11 'iq':7 'jatibarang':13 'macf':12 'mega':8 'nmax':4 'order':18 'r2':2 'syaenudin':14 't':5 'trx':17 'trx-order':16 'yamaha':3
139	66	TRX-Order	Piutang jasa WOM Finance (WOMF) Order SPK: /000000066	2022-03-22	Kendaraan R2 Honda BEAT , Nopol T 6934 YQ	'/000000066':1 '/ref-66':13 '6934':6 'beat':4 'finance':9 'honda':3 'jatibarang':11 'order':16 'r2':2 'syaenudin':12 't':5 'trx':15 'trx-order':14 'wom':8 'womf':10 'yq':7
140	65	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000065	2022-03-22	Kendaraan R2 Yamaha NMax , Nopol E 2676 UX	'/000000065':1 '/ref-65':14 '2676':6 'auto':9 'baf':11 'bussan':8 'e':5 'finance':10 'jatibarang':12 'nmax':4 'order':17 'r2':2 'syaenudin':13 'trx':16 'trx-order':15 'ux':7 'yamaha':3
141	64	TRX-Order	Piutang jasa Mandiri Utama Finance (MUF) Order SPK: /000000064	2022-03-22	Kendaraan R2 Honda Vario 150 , Nopol T 5097 ZB	'/000000064':1 '/ref-64':15 '150':5 '5097':7 'finance':11 'honda':3 'jatibarang':13 'mandir':9 'muf':12 'order':18 'r2':2 'syaenudin':14 't':6 'trx':17 'trx-order':16 'utama':10 'vario':4 'zb':8
142	63	TRX-Order	Piutang jasa OTO Kredit Motor (OTTO) Order SPK: /000000063	2022-03-22	Kendaraan R2 Honda PCX , Nopol E 2113 PBM	'/000000063':1 '/ref-63':16 '2113':6 'deddy':14 'e':5 'honda':3 'indramayu':13 'kredit':9 'motor':10 'order':19 'oto':8 'otto':11 'pbm':7 'pcx':4 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17
143	62	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /000000062	2022-03-22	Kendaraan R2 Yamaha NMax , Nopol E 3217 PAR	'/000000062':1 '/ref-62':16 '3217':6 'auto':9 'baf':11 'bussan':8 'deddy':14 'e':5 'finance':10 'indramayu':13 'nmax':4 'order':19 'par':7 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17 'yamaha':3
144	61	TRX-Order	Piutang jasa Mandiri Utama Finance (MUF) Order SPK: /000000061	2022-03-22	Kendaraan R2 Yamaha Vixion , Nopol E 3310 QR	'/000000061':1 '/ref-61':16 '3310':6 'deddy':14 'e':5 'finance':10 'indramayu':13 'mandir':8 'muf':11 'order':19 'pranoto':15 'pusat':12 'qr':7 'r2':2 'trx':18 'trx-order':17 'utama':9 'vixion':4 'yamaha':3
145	60	TRX-Order	Piutang jasa Auto Discret Finance (Adira) Order SPK: /000000060	2022-03-22	Kendaraan R2 Honda BEAT , Nopol F 6181 FCH	'/000000060':1 '/ref-60':16 '6181':6 'adira':11 'auto':8 'beat':4 'deddy':14 'discret':9 'f':5 'fch':7 'finance':10 'honda':3 'indramayu':13 'order':19 'pranoto':15 'pusat':12 'r2':2 'trx':18 'trx-order':17
\.


--
-- Data for Name: trx_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trx_detail (id, code_id, trx_id, debt, cred) FROM stdin;
1	5511	72	1040000.00	0.00
2	1113	72	0.00	1040000.00
1	5511	73	1040000.00	0.00
2	1113	73	0.00	1040000.00
1	5511	74	1040000.00	0.00
2	1113	74	0.00	1040000.00
1	5511	75	960000.00	0.00
2	1113	75	0.00	960000.00
1	5511	76	1200000.00	0.00
2	1113	76	0.00	1200000.00
1	5511	77	960000.00	0.00
2	1113	77	0.00	960000.00
1	5511	78	1360000.00	0.00
2	1113	78	0.00	1360000.00
1	5511	79	1360000.00	0.00
2	1113	79	0.00	1360000.00
1	5511	80	1040000.00	0.00
2	1113	80	0.00	1040000.00
1	5511	81	680000.00	0.00
2	1113	81	0.00	680000.00
1	5511	82	720000.00	0.00
2	1113	82	0.00	720000.00
1	5511	83	800000.00	0.00
2	1113	83	0.00	800000.00
1	5511	84	800000.00	0.00
2	1113	84	0.00	800000.00
1	5511	85	1200000.00	0.00
2	1113	85	0.00	1200000.00
1	5511	86	900000.00	0.00
2	1113	86	0.00	900000.00
1	5511	87	880000.00	0.00
2	1113	87	0.00	880000.00
1	5511	88	600000.00	0.00
2	1113	88	0.00	600000.00
1	5511	89	800000.00	0.00
2	1113	89	0.00	800000.00
1	5511	90	800000.00	0.00
2	1113	90	0.00	800000.00
1	5511	91	720000.00	0.00
2	1113	91	0.00	720000.00
1	5511	92	760000.00	0.00
2	1113	92	0.00	760000.00
1	5511	93	760000.00	0.00
2	1113	93	0.00	760000.00
1	5511	94	1200000.00	0.00
2	1113	94	0.00	1200000.00
1	5511	95	800000.00	0.00
2	1113	95	0.00	800000.00
1	5511	96	680000.00	0.00
2	1113	96	0.00	680000.00
1	5511	97	560000.00	0.00
2	1113	97	0.00	560000.00
1	5511	98	1040000.00	0.00
2	1113	98	0.00	1040000.00
1	5511	99	1160000.00	0.00
2	1113	99	0.00	1160000.00
1	5511	100	1160000.00	0.00
2	1113	100	0.00	1160000.00
1	5511	101	1440000.00	0.00
2	1113	101	0.00	1440000.00
1	5511	102	960000.00	0.00
2	1113	102	0.00	960000.00
1	5511	103	1200000.00	0.00
2	1113	103	0.00	1200000.00
1	5511	104	720000.00	0.00
2	1113	104	0.00	720000.00
1	5511	105	1440000.00	0.00
2	1113	105	0.00	1440000.00
1	5511	106	1200000.00	0.00
2	1113	106	0.00	1200000.00
1	5511	107	17000000.00	0.00
2	1113	107	0.00	17000000.00
1	5511	108	12000000.00	0.00
2	1113	108	0.00	12000000.00
1	5511	109	24200000.00	0.00
2	1113	109	0.00	24200000.00
1	5511	110	24000000.00	0.00
2	1113	110	0.00	24000000.00
1	5511	111	24200000.00	0.00
2	1113	111	0.00	24200000.00
1	5511	112	13200000.00	0.00
2	1113	112	0.00	13200000.00
1	5511	113	19000000.00	0.00
2	1113	113	0.00	19000000.00
1	5511	114	6400000.00	0.00
2	1113	114	0.00	6400000.00
1	5511	115	6400000.00	0.00
2	1113	115	0.00	6400000.00
1	5511	116	21000000.00	0.00
2	1113	116	0.00	21000000.00
1	5511	117	8000000.00	0.00
2	1113	117	0.00	8000000.00
1	5511	118	6500000.00	0.00
2	1113	118	0.00	6500000.00
1	5511	119	19750000.00	0.00
2	1113	119	0.00	19750000.00
1	5511	120	20000000.00	0.00
2	1113	120	0.00	20000000.00
1	5511	121	8000000.00	0.00
2	1113	121	0.00	8000000.00
1	5511	122	20000000.00	0.00
2	1113	122	0.00	20000000.00
1	5511	123	38400000.00	0.00
2	1113	123	0.00	38400000.00
1	5511	124	5400000.00	0.00
2	1113	124	0.00	5400000.00
1	5511	125	19500000.00	0.00
2	1113	125	0.00	19500000.00
1	5511	126	13950000.00	0.00
2	1113	126	0.00	13950000.00
1	5511	127	23000000.00	0.00
2	1113	127	0.00	23000000.00
1	5511	128	31360000.00	0.00
2	1113	128	0.00	31360000.00
1	5511	129	8200000.00	0.00
2	1113	129	0.00	8200000.00
1	5511	130	1280000.00	0.00
2	1113	130	0.00	1280000.00
1	5511	131	540000.00	0.00
2	1113	131	0.00	540000.00
1	5511	132	1000000.00	0.00
2	1113	132	0.00	1000000.00
1	5511	133	1000000.00	0.00
2	1113	133	0.00	1000000.00
1	5511	134	1040000.00	0.00
2	1113	134	0.00	1040000.00
1	5511	135	1120000.00	0.00
2	1113	135	0.00	1120000.00
1	5511	136	1120000.00	0.00
2	1113	136	0.00	1120000.00
1	5511	137	960000.00	0.00
2	1113	137	0.00	960000.00
1	5511	138	1600000.00	0.00
2	1113	138	0.00	1600000.00
1	5511	139	1040000.00	0.00
2	1113	139	0.00	1040000.00
1	5511	140	3200000.00	0.00
2	1113	140	0.00	3200000.00
1	5511	141	1600000.00	0.00
2	1113	141	0.00	1600000.00
1	5511	143	2880000.00	0.00
2	1113	143	0.00	2880000.00
1	5511	142	1440000.00	0.00
2	1113	142	0.00	1440000.00
1	5511	145	920000.00	0.00
2	1113	145	0.00	920000.00
1	5511	144	1200000.00	0.00
2	1113	144	0.00	1200000.00
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
8	E 6277 PAZ	2018	\N	\N	\N	11	2
9	E 6053 PAM	2017	\N	\N	\N	6	1
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
25	B 3812 UJY	2015	\N	\N	\N	19	2
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
43	T 8060 EG	2020	\N	\N	\N	23	1
44	E 8903 PP	2014	\N	\N	\N	23	1
45	E 8013 QA	2017	\N	\N	\N	26	1
46	T 1052 UL	2020	\N	\N	\N	27	1
47	BG 1623 PF	2006	\N	\N	\N	22	1
48	E 938 XY	2018	\N	\N	\N	28	1
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
23	E 2815 PBX	2021	\N	\N	\N	17	4
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

SELECT pg_catalog.setval('public.action_id_seq', 7, true);


--
-- Name: branch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.branch_id_seq', 5, true);


--
-- Name: finance_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finance_groups_id_seq', 3, true);


--
-- Name: finance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finance_id_seq', 22, true);


--
-- Name: invoices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.invoices_id_seq', 5, true);


--
-- Name: merk_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.merk_id_seq', 17, true);


--
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_id_seq', 76, true);


--
-- Name: order_name_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_name_seq', 100, true);


--
-- Name: trx_detail_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_detail_seq', 1, false);


--
-- Name: trx_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_seq', 145, true);


--
-- Name: type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.type_id_seq', 39, true);


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

