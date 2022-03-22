--
-- PostgreSQL database dump
--

-- Dumped from database version 12.10 (Ubuntu 12.10-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.2 (Ubuntu 14.2-1.pgdg20.04+1)

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
    receivable_option smallint DEFAULT 0,
    is_auto_debet boolean DEFAULT false
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
    subtotal numeric(12,2) DEFAULT 0 NOT NULL,
    ppn numeric(8,2) DEFAULT 0 NOT NULL,
    tax numeric(12,2) DEFAULT 0 NOT NULL,
    total numeric(12,2) DEFAULT 0 NOT NULL,
    account_id smallint DEFAULT 0 NOT NULL,
    memo character varying(256),
    token tsvector
);


ALTER TABLE public.invoices OWNER TO postgres;

--
-- Name: invoices_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.invoices_id_seq
    AS integer
    START WITH 100
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 100;


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

COPY public.acc_code (id, name, type_id, descriptions, token_name, is_active, receivable_option, is_auto_debet) FROM stdin;
1111	Kas Kecil	11	Kas bendahara Kantor qwewqe 	'bendahara':4 'kantor':5 'kas':1,3 'kecil':2 'qwewqe':6	t	1	f
5311	Biaya Gaji karyawan Tetap	53	Pencatan data kompensasi karyawan seperti uang potongan dari setiap gaji dan pajak serta tunjangan karyawan tetap	'biaya':1 'catan':5 'dan':15 'dari':12 'data':6 'gaji':2,14 'karyaw':3,8,19 'kompensasi':7 'pajak':16 'potong':11 'sepert':9 'serta':17 'setiap':13 'tetap':4,20 'tunjang':18 'uang':10	t	3	t
3211	Prive Pak Kris	32	Pengambilan modal, pinjam kas oleh pak Kris	'ambil':4 'kas':7 'kris':3,10 'modal':5 'oleh':8 'pak':2,9 'pinjam':6 'prive':1	t	3	f
5111	Biaya Transport	51	Biaya transportasi karyawan	'biaya':1,3 'karyaw':5 'transport':2 'transportasi':4	t	3	f
5113	Biaya Telephone dan Fax	51	Biaya telephone dan faximile ke telkomsel	'biaya':1,5 'dan':3,7 'fax':4 'faximile':8 'ke':9 'telephone':2,6 'telkomsel':10	t	3	f
5117	Biaya Servis	51	Biaya service kendaraan, AC, komputer, dll.	'ac':6 'biaya':1,3 'dll':8 'komputer':7 'ndara':5 'service':4 'servis':2	t	3	f
5118	Biaya Konsumsi	51	Biaya yg dikeluarkan karena suatu kegiatan yg dpt mengurangi atau menghabiskan barang dan jasa	'atau':12 'barang':14 'biaya':1,3 'dan':15 'dpt':10 'giat':8 'habis':13 'jasa':16 'karena':6 'keluar':5 'konsumsi':2 'suatu':7 'urang':11 'yg':4,9	t	3	f
5115	Biaya Pos dan Materai	51	Biaya pengiriman surat dan pembelian materai.	'beli':9 'biaya':1,5 'dan':3,8 'irim':6 'matera':4,10 'pos':2 'surat':7	t	3	t
5312	Biaya Gaji Karyawan Honorer	51	Pencatan data kompensasi karyawan seperti uang potongan dari setiap gaji\ndan pajak serta tunjangan bukan karyawan tetap 	'biaya':1 'bukan':19 'catan':5 'dan':15 'dari':12 'data':6 'gaji':2,14 'honorer':4 'karyaw':3,8,20 'kompensasi':7 'pajak':16 'potong':11 'sepert':9 'serta':17 'setiap':13 'tetap':21 'tunjang':18 'uang':10	t	3	f
6011	Pembayaran Pajak	60	Pajak Pertambahan Nilai	'bayar':1 'nila':5 'pajak':2,3 'tambah':4	t	3	f
5211	Biaya STNK	52	Biaya yg dikeluarkan untuk penarikan kendaraan yg tidak ada STNK	'ada':11 'arik':7 'biaya':1,3 'keluar':5 'ndara':8 'stnk':2,12 'tidak':10 'untuk':6 'yg':4,9	t	3	t
5112	Biaya Listrik	51	Biaya pemakaian listrik	'biaya':1,3 'listrik':2,5 'pakai':4	t	3	t
5114	Biaya Internet	51	Biaya jaringan internet ke Biznet	'biaya':1,3 'biznet':7 'internet':2,5 'jaring':4 'ke':6	t	3	t
5116	Biaya ATK	51	Biaya alat tulis kantor termasuk termasuk peralatan seperti komputer, meja, kursi, lemari	'alat':4,9 'atk':2 'biaya':1,3 'kantor':6 'komputer':11 'kursi':13 'lemar':14 'masuk':7,8 'meja':12 'sepert':10 'tulis':5	t	3	t
5119	Biaya Lain-lain	51	Biaya yg terdiri dari bermacam transaksi serta tidak tercantum pada salah satu perkiraan yang terdapat dalam transaksi perusahaan	'biaya':1,5 'cantum':13 'dalam':20 'dapat':19 'dari':8 'diri':7 'kira':17 'lain':3,4 'lain-lain':2 'macam':9 'pada':14 'salah':15 'satu':16 'serta':11 'tidak':12 'transaksi':10,21 'usaha':22 'yang':18 'yg':6	t	3	t
5411	Upah Tenaga Kerja	54	Biaya overhead perusahaan yg dikeluarkan untuk memayar upah karena mengerjakan sesuatu	'biaya':4 'erja':13 'karena':12 'keluar':8 'kerja':3 'overhead':5 'payar':10 'sesuatu':14 'tenaga':2 'untuk':9 'upah':1,11 'usaha':6 'yg':7	t	3	f
3111	Modal Pak Kris	31	Modal yg diterima dari pak Kris	'dari':7 'kris':3,9 'modal':1,4 'pak':2,8 'terima':6 'yg':5	t	2	t
2311	Hutang Pajak	23	Pajak yg belum dibayar karena menunggu pembayaran dari tarikan	'bayar':6,9 'belum':5 'dari':10 'hutang':1 'karena':7 'pajak':2,3 'tari':11 'unggu':8 'yg':4	t	2	f
1112	Bank BCA 0856212654	11	Atas nama Opik	'0856212654':3 'atas':4 'bank':1 'bca':2 'nama':5 'opik':6	t	1	f
5511	Piutang Jasa	55	Pendanaan yg dikeluarkan untuk operasi penarikan berdasarkan SPK dari Finance sejumlah BT Matel	'arik':8 'bt':14 'dana':3 'dari':11 'dasar':9 'finance':12 'jasa':2 'keluar':5 'matel':15 'operasi':7 'piutang':1 'sejum':13 'spk':10 'untuk':6 'yg':4	t	3	f
4111	Pendapatan Invoice	41	Penarikan dana dari pihak Finance karena adanya ...	'ada':9 'arik':3 'dana':4 'dapat':1 'dari':5 'finance':7 'invoice':2 'karena':8 'pihak':6	t	2	f
\.


--
-- Data for Name: acc_group; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.acc_group (id, name, descriptions) FROM stdin;
1	Harta	Segala sesuatu yang berhubungan dengan asset perusahaan
3	Modal	Kekayaan perusahaan yang menjadi bagian dari pemilik perusahaan.
4	Pendapatan	Segala sesuatu yang diterima oleh perusahaan, baik yang didapat dari hasil operasional perusahaan (misalnya, bengkel mendapat pendapatan jasa servis kendaraan) dan kegiatan di luar operasional perusahaan (misalnya, bunga bank)
5	Beban	Biaya-biaya yang dikeluarkan perusahaan dalam kegiatan operasionalnya untuk mendapatkan penghasilan. Contoh: beban air, listrik, dan telepon.
2	Utang	Segala sesuatu yang menjadi kewajiban perusahaan yang harus dibayarkan kepada pihak luar dalam periode tertentu.
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
55	Piutang	Piutang adalah tagihan yang dipinjamkan kepada pelanggan dan wajib dilunasi paling lama satu tahun atau menurut kesepakatan	5
13	Persediaan	Kelompok akun yg digunakan untuk mencatat persediaan bahan baku yang menunggu penggunaannya dalam suatu proses produksi.	1
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
1	Jatibarang	Jl. Pasar Sepur	Jatibarang	08596522323	012454787	45616	Mastur	mastur.st12@gmail.com
3	Pusat Indramayu	\N	\N	\N	\N	\N	Deddy Pranoto	\N
4	Karawang	\N	\N	\N	\N	\N	Gugur Junaedi	\N
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (order_id, name, agreement_number, payment_type) FROM stdin;
\.


--
-- Data for Name: finances; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.finances (id, name, short_name, street, city, phone, cell, zip, email) FROM stdin;
1	Bussan Auto Finance	BAF	Jl. Jend. Sudirman	Indramayu	2569874545	65979	2598987	busan.123@gmail.com
2	Auto Discret Finance	Adira	Jl. Jend. Sudirman	Indramayu	2569874545	65979	2598987	adira.finance@gmail.com
3	Mandiri Tunas Finance	MTF	\N	Cirebon	\N	\N	\N	\N
4	Clipan Bekasi	CLIP B	\N	\N	\N	\N	\N	\N
5	OTO Kredit Motor	OTTO	\N	\N	\N	\N	\N	\N
6	COLLECTIUS	COL	\N	\N	\N	\N	\N	\N
7	Mandiri Utama Finance	MUF	\N	\N	\N	\N	\N	\N
8	FIF Group	FIF	\N	\N	\N	\N	\N	\N
9	Mitra Pinasthika Mustika Finance	MPMF	\N	\N	\N	\N	\N	\N
10	Top Finance Company	TFC	\N	\N	\N	\N	\N	\N
11	Kredit Plus	KP+	\N	\N	\N	\N	\N	\N
12	WOM Finance	WOMF	\N	\N	\N	\N	\N	\N
13	MEGAPARA	MPR	\N	\N	\N	\N	\N	\N
14	Clipan Karawang\n	CLIP K	\N	\N	\N	\N	\N	\N
15	Clipan Palembang	CLIP P	\N	\N	\N	\N	\N	\N
16	Safron Finance Karawang	SFI K	\N	\N	\N	\N	\N	\N
17	BFI Finance	BFI	\N	\N	\N	\N	\N	\N
18	Mandiri Tunas Finance Semarang	MTF S	\N	\N	\N	\N	\N	\N
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

COPY public.invoices (id, invoice_at, payment_term, due_at, salesman, finance_id, subtotal, ppn, tax, total, account_id, memo, token) FROM stdin;
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
47	000000047	2022-03-16	2022-03-16	9000000.00	11.11	8000000.00	Mastur	\N	15	3	t	0.00	9000000.00	'1623':7 '2006':1 'bg':6 'clip':4 'honda':2 'jazz':3 'p':5 'pf':8 'r4':9
46	000000046	2022-03-14	2022-03-14	23500000.00	10.64	21000000.00	Mastur	\N	16	3	t	0.00	23500000.00	'-7':4 '1052':8 '2020':1 'k':6 'r4':10 'sfi':5 'suzuk':2 't':7 'ul':9 'xl':3
45	000000045	2022-03-14	2022-03-14	8000000.00	20.00	6400000.00	Mastur	\N	2	3	t	0.00	8000000.00	'2017':1 '8013':6 'adira':4 'e':5 'mitsubish':2 'pickup':3 'qa':7 'r4':8
44	000000044	2022-03-12	2022-03-12	8000000.00	20.00	6400000.00	Mastur	\N	17	3	t	0.00	8000000.00	'2014':1 '8903':6 'bfi':4 'carry':3 'e':5 'pp':7 'r4':8 'suzuk':2
43	000000043	2022-03-09	2022-03-09	21000000.00	9.52	19000000.00	Mastur	\N	16	3	t	0.00	21000000.00	'2020':1 '8060':7 'carry':3 'eg':8 'k':5 'r4':9 'sfi':4 'suzuk':2 't':6
42	000000042	2022-02-16	2022-02-16	15000000.00	12.00	13200000.00	Mastur	\N	18	3	t	0.00	15000000.00	'1000':4 '2020':1 '9049':8 'brio':3 'h':7 'honda':2 'mtf':5 'r4':10 's':6 'se':9
41	000000041	2022-01-25	2022-01-25	26620000.00	9.09	24200000.00	Mastur	\N	14	3	t	0.00	26620000.00	'1788':7 '2017':1 'bc':8 'clip':4 'honda':2 'k':5 'mobilio':3 'r4':9 't':6
39	000000039	2022-01-14	2022-01-14	26000000.00	6.92	24200000.00	Mastur	\N	18	3	t	0.00	26000000.00	'1000':4 '2016':1 '8715':8 'brio':3 'gp':9 'h':7 'honda':2 'mtf':5 'r4':10 's':6
38	000000038	2021-12-29	2021-12-29	13000000.00	7.69	12000000.00	Mastur	\N	14	3	t	0.00	13000000.00	'1412':7 '2004':1 'carry':3 'clip':4 'k':5 'km':8 'r4':9 'suzuk':2 't':6
37	000000037	2021-12-02	2021-12-02	20000000.00	15.00	17000000.00	Mastur	\N	14	3	t	0.00	20000000.00	'2006':1 '8936':7 'b':6 'clip':4 'honda':2 'jazz':3 'k':5 'no':8 'r4':9
36	000000036	2022-03-18	2022-03-18	1500000.00	20.00	1200000.00	Mastur	\N	1	1	t	0.00	1500000.00	'2018':1 '5713':7 'baf':5 'e':6 'm3':4 'mio':3 'pav':8 'r2':9 'yamaha':2
35	000000035	2022-03-18	2022-03-18	1800000.00	20.00	1440000.00	Mastur	\N	2	1	t	0.00	1800000.00	'-15':4 '2017':1 '2110':7 'adira':5 'r':3 'r2':9 't':6 'yamaha':2 'yv':8
34	000000034	2022-03-17	2022-03-17	900000.00	20.00	720000.00	Mastur	\N	2	1	t	0.00	900000.00	'2012':1 '4593':6 'adira':4 'e':5 'jupiter':3 'r2':8 'tq':7 'yamaha':2
33	000000033	2022-03-16	2022-03-16	1500000.00	20.00	1200000.00	Mastur	\N	2	1	t	0.00	1500000.00	'2019':1 '5856':6 'adira':4 'genio':3 'honda':2 'r2':8 't':5 'zt':7
32	000000032	2022-03-14	2022-03-14	1200000.00	20.00	960000.00	Mastur	\N	2	1	t	0.00	1200000.00	'2017':1 '2191':6 'adira':4 'beat':3 'honda':2 'r2':8 't':5 'ys':7
31	000000031	2022-03-12	2022-03-12	1800000.00	20.00	1440000.00	Mastur	\N	2	1	t	0.00	1800000.00	'-15':4 '2017':1 '2391':7 'adira':5 'e':6 'jm':8 'r':3 'r2':9 'yamaha':2
30	000000030	2022-03-12	2022-03-12	1450000.00	20.00	1160000.00	Mastur	\N	2	1	t	0.00	1450000.00	'2018':1 '3615':6 'adira':4 'honda':2 'r2':8 't':5 'verza':3 'zd':7
29	000000029	2022-03-09	2022-03-09	1450000.00	20.00	1160000.00	Mastur	\N	13	1	t	0.00	1450000.00	'2015':1 '4544':6 'e':5 'jd':7 'mio':3 'mpr':4 'r2':8 'yamaha':2
28	000000028	2022-03-09	2022-03-09	1300000.00	20.00	1040000.00	Mastur	\N	2	1	t	0.00	1300000.00	'2017':1 '4487':7 'adira':5 'mio':3 'pj':8 'r2':9 't':6 'yamaha':2 'z':4
27	000000027	2022-03-09	2022-03-09	700000.00	20.00	560000.00	Mastur	\N	2	1	t	0.00	700000.00	'2014':1 '6819':6 'adira':4 'b':5 'pzi':7 'r2':8 'xeon':3 'yamaha':2
26	000000026	2022-03-07	2022-03-07	850000.00	20.00	680000.00	Mastur	\N	11	1	t	0.00	850000.00	'150':4 '2014':1 '2891':7 'honda':2 'kp':5 'r2':9 't':6 'vario':3 'wp':8
25	000000025	2022-03-02	2022-03-02	1000000.00	20.00	800000.00	Mastur	\N	2	1	t	0.00	1000000.00	'150':4 '2015':1 '3812':7 'adira':5 'b':6 'honda':2 'r2':9 'ujy':8 'vario':3
24	000000024	2022-02-25	2022-02-25	1500000.00	20.00	1200000.00	Mastur	\N	1	1	t	0.00	1500000.00	'2018':1 '2146':7 'baf':5 'e':6 'mio':3 'qaf':8 'r2':9 's':4 'yamaha':2
23	000000023	2022-02-23	2022-02-23	0.00	20.00	40000.00	Mastur	\N	12	1	t	0.00	0.00	'2021':1 '2815':6 'e':5 'gear':3 'pbx':7 'r2':8 'womf':4 'yamaha':2
22	000000022	2022-02-23	2022-02-23	950000.00	20.00	760000.00	Mastur	\N	1	1	t	0.00	950000.00	'2015':1 '2830':7 'baf':5 'e':6 'm3':4 'mio':3 'qr':8 'r2':9 'yamaha':2
21	000000021	2022-02-21	2022-02-21	950000.00	20.00	760000.00	Mastur	\N	1	1	t	0.00	950000.00	'2015':1 '6262':7 'b':6 'baf':5 'm3':4 'mio':3 'r2':9 'vky':8 'yamaha':2
19	000000019	2022-01-26	2022-01-26	1000000.00	20.00	800000.00	Mastur	\N	2	1	t	0.00	1000000.00	'2018':1 '5638':6 'adira':4 'e':5 'honda':2 'pav':7 'r2':8 'revo':3
15	000000015	2022-01-06	2022-01-06	900000.00	20.00	900000.00	Mastur	\N	5	1	t	0.00	900000.00	'2012':1 '4146':7 'jupiter':3 'ko':8 'mx':4 'otto':5 'r2':9 't':6 'yamaha':2
13	000000013	2021-12-06	2021-12-06	1000000.00	20.00	800000.00	Mastur	\N	6	1	t	0.00	1000000.00	'2017':1 '2417':7 'col':5 'e':6 'm3':4 'mio':3 'pao':8 'r2':9 'yamaha':2
11	000000011	2021-09-27	2021-09-27	900000.00	20.00	720000.00	Mastur	\N	8	1	t	0.00	900000.00	'2012':1 '4892':6 'e':5 'fif':4 'jupiter':3 'r2':8 'tk':7 'yamaha':2
9	000000009	2022-03-18	2022-03-18	1300000.00	20.00	1040000.00	Mastur	\N	2	3	t	0.00	1300000.00	'2017':1 '6053':6 'adira':4 'beat':3 'e':5 'honda':2 'pam':7 'r2':8
8	000000008	2022-03-17	2022-03-17	1700000.00	20.00	1360000.00	Mastur	\N	2	3	t	0.00	1700000.00	'-15':4 '2018':1 '6277':7 'adira':5 'e':6 'paz':8 'r':3 'r2':9 'yamaha':2
7	000000007	2022-03-15	2022-03-15	1700000.00	20.00	1360000.00	Mastur	\N	1	3	t	0.00	1700000.00	'125':4 '2019':1 '2033':7 'baf':5 'e':6 'fino':3 'pbj':8 'r2':9 'yamaha':2
6	000000006	2022-03-02	2022-03-02	1200000.00	20.00	960000.00	Mastur	\N	2	3	t	0.00	1200000.00	'2016':1 '2633':6 'adira':4 'beat':3 'e':5 'honda':2 'pac':7 'r2':8
5	000000005	2022-02-24	2022-02-24	1500000.00	20.00	1200000.00	Mastur	\N	7	3	t	0.00	1500000.00	'125':4 '2017':1 '4096':7 'e':6 'fino':3 'muf':5 'paq':8 'r2':9 'yamaha':2
3	000000003	2022-02-07	2022-02-07	1300000.00	20.00	1040000.00	Mastur	\N	1	3	t	0.00	1300000.00	'2018':1 '5125':7 'baf':5 'e':6 'm3':4 'mio':3 'pbc':8 'r2':9 'yamaha':2
1	000000001	2021-12-15	2021-12-15	1300000.00	20.00	1040000.00	Mastur	\N	5	3	t	0.00	1300000.00	'2017':1 '5605':7 'e':6 'mio':3 'otto':5 'pas':8 'r2':9 'yamaha':2 'z':4
48	000000048	2022-03-18	2022-03-18	8000000.00	18.75	6500000.00	Mastur	\N	2	3	t	0.00	8000000.00	'2018':1 '938':6 'adira':4 'e':5 'expander':3 'mitsubish':2 'r4':8 'xy':7
40	000000040	2022-01-17	2022-01-17	26000000.00	7.69	24000000.00	Mastur	\N	14	3	t	0.00	26000000.00	'-3':4 '1164':8 '2017':1 'clip':5 'er':3 'fq':9 'k':6 'r4':10 'suzuk':2 't':7
20	000000020	2022-01-21	2022-01-21	900000.00	20.00	720000.00	Mastur	\N	11	1	t	0.00	900000.00	'125':4 '2013':1 '5253':7 'e':6 'honda':2 'kp':5 'r2':9 'ty':8 'vario':3
18	000000018	2022-01-18	2022-01-18	1000000.00	20.00	800000.00	Mastur	\N	5	1	t	0.00	1000000.00	'2015':1 '6716':7 'e':6 'ix':8 'm3':4 'mio':3 'otto':5 'r2':9 'yamaha':2
17	000000017	2022-01-14	2022-01-14	750000.00	20.00	600000.00	Mastur	\N	10	1	t	0.00	750000.00	'-125':5 '2008':1 '3828':8 'fw':9 'honda':2 'r2':10 'supra':3 't':7 'tfc':6 'x':4
16	000000016	2022-01-14	2022-01-14	1100000.00	20.00	880000.00	Mastur	\N	5	1	t	0.00	1100000.00	'2016':1 '3848':6 'beat':3 'e':5 'honda':2 'otto':4 'r2':8 'ub':7
14	000000014	2021-12-07	2021-12-07	1500000.00	20.00	1200000.00	Mastur	\N	9	1	t	0.00	1500000.00	'2012':1 '3521':7 'fu':4 'kl':8 'mpmf':5 'r2':9 'satria':3 'suzuk':2 't':6
12	000000012	2021-11-04	2021-11-04	1000000.00	20.00	800000.00	Mastur	\N	6	1	t	0.00	1000000.00	'2016':1 '3479':7 'b':6 'col':5 'm3':4 'mio':3 'r2':9 'uju':8 'yamaha':2
10	000000010	2022-03-19	2022-03-19	850000.00	20.00	680000.00	Mastur	\N	2	3	t	0.00	850000.00	'2013':1 '5474':6 'adira':4 'beat':3 'e':5 'honda':2 'q':7 'r2':8
4	000000004	2022-02-18	2022-02-18	1200000.00	20.00	960000.00	Mastur	\N	2	3	t	0.00	1200000.00	'2015':1 '5080':6 'adira':4 'br':5 'py':7 'r2':8 'vixion':3 'yamaha':2
2	000000002	2022-01-10	2022-01-10	1300000.00	20.00	1040000.00	Mastur	\N	6	3	t	0.00	1300000.00	'2016':1 '3977':6 'col':4 'e':5 'mio':3 'pac':7 'r2':8 'yamaha':2
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
164	10	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /qweqwe qweqwe	2022-03-19	Kendaraan R2 Yamaha Fino 125 , Nopol E25632FF	'/qweqwe':1 '/ref-10':14 '125':6 'auto':9 'baf':11 'bussan':8 'e25632ff':7 'finance':10 'fino':5 'jatibarang':12 'mastur':13 'order':17 'qweqwe':2 'r2':3 'trx':16 'trx-order':15 'yamaha':4
162	11	TRX-Order	Piutang jasa Bussan Auto Finance (BAF) Order SPK: /x-001254	2022-03-19	Kendaraan R2 Honda Vario 125 , Nopol E56985698	'/ref-11':13 '/x-001254':1 '125':5 'auto':8 'baf':10 'bussan':7 'e56985698':6 'finance':9 'honda':3 'jatibarang':11 'mastur':12 'order':16 'r2':2 'trx':15 'trx-order':14 'vario':4
169	104	trx-invoice	Pendapatan jasa dari Bussan Auto Finance Invoice #104	2022-03-22	\N	'/id-0':2 '104':19 '125':9 '125pendapatan':12 'auto':4,16 'baf':6 'bussan':3,15 'dari':14 'e25632ff':10 'e56985698':7 'finance':5,17 'fino':11 'invoice':18 'jasa':13 'vario':8 'wakid':1
\.


--
-- Data for Name: trx_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trx_detail (id, code_id, trx_id, debt, cred) FROM stdin;
1	4111	169	0.00	2430000.00
2	1112	169	2430000.00	0.00
1	5511	162	960000.00	0.00
2	1112	162	0.00	960000.00
1	5511	164	1200000.00	0.00
2	1112	164	0.00	1200000.00
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
24	ER-3	3	12
25	Mobilio	3	13
26	Pickup	3	1
27	XL-7	3	12
28	Expander	3	1
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
23	E 2815 PBX	2021	\N	\N	\N	17	4
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
3	GUDANG	---
4	KURANG 300	---
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

SELECT pg_catalog.setval('public.action_id_seq', 4, true);


--
-- Name: branch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.branch_id_seq', 2, true);


--
-- Name: finance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.finance_id_seq', 3, true);


--
-- Name: invoices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.invoices_id_seq', 202, true);


--
-- Name: merk_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.merk_id_seq', 14, true);


--
-- Name: order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_id_seq', 13, true);


--
-- Name: order_name_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_name_seq', 24, true);


--
-- Name: trx_detail_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_detail_seq', 1, false);


--
-- Name: trx_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trx_seq', 170, true);


--
-- Name: type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.type_id_seq', 28, true);


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
-- Name: ix_invoice_token; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX ix_invoice_token ON public.invoices USING gin (token);


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

