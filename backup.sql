--
-- PostgreSQL database dump
--

-- Dumped from database version 12.8
-- Dumped by pg_dump version 13.2

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

--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: categories; Type: TYPE; Schema: public; Owner: root
--

CREATE TYPE public.categories AS ENUM (
    'men',
    'women',
    'sneaker',
    'hat',
    'jacket',
    'nothing'
);


ALTER TYPE public.categories OWNER TO root;

--
-- Name: role; Type: TYPE; Schema: public; Owner: root
--

CREATE TYPE public.role AS ENUM (
    'admin',
    'pro',
    'user'
);


ALTER TYPE public.role OWNER TO root;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: files; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.files (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    name text,
    url text,
    mime text,
    size bigint
);


ALTER TABLE public.files OWNER TO root;

--
-- Name: products; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.products (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    name text NOT NULL,
    category public.categories DEFAULT 'nothing'::public.categories NOT NULL,
    cover text NOT NULL,
    price double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.products OWNER TO root;

--
-- Name: refresh_token; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.refresh_token (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    token text NOT NULL,
    expir_on timestamp with time zone NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE public.refresh_token OWNER TO root;

--
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    lastname text NOT NULL,
    firstname text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    role public.role DEFAULT 'user'::public.role NOT NULL,
    birthday text,
    phone text,
    firebase_id_token text,
    firebase_uid text,
    firebase_provider text,
    password_confirm_code text
);


ALTER TABLE public.users OWNER TO root;

--
-- Data for Name: files; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.products VALUES ('a5efa38b-311f-4c67-946d-1b8d0e6e02a9', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', NULL, 'Brown Brim', 'hat', 'https://i.ibb.co/ZYW3VTp/brown-brim.png', 25);
INSERT INTO public.products VALUES ('52ae7981-78a3-4d29-ad5b-34088f15fd7b', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', NULL, 'Brown Cowboy', 'hat', 'https://i.ibb.co/QdJwgmp/brown-cowboy.png', 35);
INSERT INTO public.products VALUES ('384b1373-062a-4334-a7c2-bc8ad60d223d', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', NULL, 'Green Beanie', 'hat', 'https://i.ibb.co/YTjW3vF/green-beanie.png', 18);
INSERT INTO public.products VALUES ('95652bbf-96b2-4bbc-b031-206ee540ff28', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', NULL, 'Palm Tree Cap', 'hat', 'https://i.ibb.co/rKBDvJX/palm-tree-cap.png', 14);
INSERT INTO public.products VALUES ('d46acb3d-2bc6-426e-bb12-7f5d307941f0', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', NULL, 'Red Beanie', 'hat', 'https://i.ibb.co/bLB646Z/red-beanie.png', 18);
INSERT INTO public.products VALUES ('4d338dd9-1449-47f6-8ead-d983f3bf6678', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', NULL, 'Wolf Cap', 'hat', 'https://i.ibb.co/1f2nWMM/wolf-cap.png', 14);
INSERT INTO public.products VALUES ('e723d8a8-79ad-4bd8-8f32-eb8b14b6e5dc', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', NULL, 'Blue Snapback', 'hat', 'https://i.ibb.co/X2VJP2W/blue-snapback.png', 16);
INSERT INTO public.products VALUES ('96ef481c-f5e5-498f-bd4d-15a39d9a5a55', '2021-11-27 09:37:06.451223+00', '2021-11-27 09:37:06.451223+00', NULL, 'Adidas NMD', 'sneaker', 'https://i.ibb.co/0s3pdnc/adidas-nmd.png', 220);
INSERT INTO public.products VALUES ('e4c2cc24-bfdb-4785-8da0-ba98c1f401cc', '2021-11-27 09:37:06.451223+00', '2021-11-27 09:37:06.451223+00', NULL, 'Adidas Yeezy', 'sneaker', 'https://i.ibb.co/dJbG1cT/yeezy.png', 280);
INSERT INTO public.products VALUES ('89165f4b-7028-43f3-b2ce-b9f59d478b0c', '2021-11-27 09:37:06.451223+00', '2021-11-27 09:37:06.451223+00', NULL, 'Black Converse', 'sneaker', 'https://i.ibb.co/bPmVXyP/black-converse.png', 110);
INSERT INTO public.products VALUES ('4232a42f-dfed-4d3e-b4bc-5a7b048ea3d2', '2021-11-27 09:37:06.451223+00', '2021-11-27 09:37:06.451223+00', NULL, 'Nike White AirForce', 'sneaker', 'https://i.ibb.co/1RcFPk0/white-nike-high-tops.png', 160);
INSERT INTO public.products VALUES ('00db990a-ec75-4001-966c-16cc624eecbc', '2021-11-27 09:37:06.451223+00', '2021-11-27 09:37:06.451223+00', NULL, 'Nike Red High Tops', 'sneaker', 'https://i.ibb.co/QcvzydB/nikes-red.png', 160);
INSERT INTO public.products VALUES ('c2b6942e-d799-475c-b1db-8897bf113dd1', '2021-11-27 09:37:06.451223+00', '2021-11-27 09:37:06.451223+00', NULL, 'Nike Brown High Tops', 'sneaker', 'https://i.ibb.co/fMTV342/nike-brown.png', 160);
INSERT INTO public.products VALUES ('95284142-7b2a-4d18-9c15-9cd29c2290cd', '2021-11-27 09:37:06.451223+00', '2021-11-27 09:37:06.451223+00', NULL, 'Air Jordan Limited', 'sneaker', 'https://i.ibb.co/w4k6Ws9/nike-funky.png', 190);
INSERT INTO public.products VALUES ('347272ef-586a-4347-bf67-1882916a5fcc', '2021-11-27 09:37:06.451223+00', '2021-11-27 09:37:06.451223+00', NULL, 'Timberlands', 'sneaker', 'https://i.ibb.co/Mhh6wBg/timberlands.png', 200);
INSERT INTO public.products VALUES ('4c741b3e-623b-43aa-8047-18158a851223', '2021-11-27 09:38:35.974382+00', '2021-11-27 09:38:35.974382+00', NULL, 'Black Jean Shearling', 'jacket', 'https://i.ibb.co/XzcwL5s/black-shearling.png', 125);
INSERT INTO public.products VALUES ('9349a4db-2e87-45c5-806c-bb4d8cf21bee', '2021-11-27 09:38:35.974382+00', '2021-11-27 09:38:35.974382+00', NULL, 'Blue Jean Jacket', 'jacket', 'https://i.ibb.co/mJS6vz0/blue-jean-jacket.png', 90);
INSERT INTO public.products VALUES ('ac105483-0240-457d-bcf1-62a7ced613f1', '2021-11-27 09:38:35.974382+00', '2021-11-27 09:38:35.974382+00', NULL, 'Grey Jean Jacket', 'jacket', 'https://i.ibb.co/N71k1ML/grey-jean-jacket.png', 90);
INSERT INTO public.products VALUES ('b3d3f20c-6262-4e5e-b6f0-6f6592d0fb55', '2021-11-27 09:38:35.974382+00', '2021-11-27 09:38:35.974382+00', NULL, 'Brown Shearling', 'jacket', 'https://i.ibb.co/s96FpdP/brown-shearling.png', 165);
INSERT INTO public.products VALUES ('ee375c91-ef05-4fb6-ab54-f044cbc62226', '2021-11-27 09:38:35.974382+00', '2021-11-27 09:38:35.974382+00', NULL, 'Tan Trench', 'jacket', 'https://i.ibb.co/M6hHc3F/brown-trench.png', 185);
INSERT INTO public.products VALUES ('69212359-2aa2-4faf-8bc0-07c089e77c7e', '2021-11-27 09:40:36.456722+00', '2021-11-27 09:40:36.456722+00', NULL, 'Blue Tanktop', 'women', 'https://i.ibb.co/7CQVJNm/blue-tank.png', 25);
INSERT INTO public.products VALUES ('e6dc0d77-9e21-47f0-95fc-67a74fb62934', '2021-11-27 09:40:36.456722+00', '2021-11-27 09:40:36.456722+00', NULL, 'Floral Blouse', 'women', 'https://i.ibb.co/4W2DGKm/floral-blouse.png', 20);
INSERT INTO public.products VALUES ('f30431c6-6739-4cb9-92be-6b7c18012ec0', '2021-11-27 09:40:36.456722+00', '2021-11-27 09:40:36.456722+00', NULL, 'Floral Dress', 'women', 'https://i.ibb.co/KV18Ysr/floral-skirt.png', 80);
INSERT INTO public.products VALUES ('8cf59c4c-8c39-4a15-8622-9f1de6e1a7ad', '2021-11-27 09:40:36.456722+00', '2021-11-27 09:40:36.456722+00', NULL, 'Red Dots Dress', 'women', 'https://i.ibb.co/N3BN1bh/red-polka-dot-dress.png', 80);
INSERT INTO public.products VALUES ('bfb86f23-e631-4138-8f96-767592675baa', '2021-11-27 09:40:36.456722+00', '2021-11-27 09:40:36.456722+00', NULL, 'Striped Sweater', 'women', 'https://i.ibb.co/KmSkMbH/striped-sweater.png', 45);
INSERT INTO public.products VALUES ('90e9718a-a46f-4703-a18d-caeb6c95244b', '2021-11-27 09:40:36.456722+00', '2021-11-27 09:40:36.456722+00', NULL, 'Yellow Track Suit', 'women', 'https://i.ibb.co/v1cvwNf/yellow-track-suit.png', 135);
INSERT INTO public.products VALUES ('cec80a4c-6b76-4381-b34f-180094d40cc7', '2021-11-27 09:40:36.456722+00', '2021-11-27 09:40:36.456722+00', NULL, 'White Blouse', 'women', 'https://i.ibb.co/qBcrsJg/white-vest.png', 20);
INSERT INTO public.products VALUES ('3c0a0cde-ae0a-4faa-8a5f-a960cd948af7', '2021-11-27 09:41:27.500774+00', '2021-11-27 09:41:27.500774+00', NULL, 'Camo Down Vest', 'men', 'https://i.ibb.co/xJS0T3Y/camo-vest.png', 325);
INSERT INTO public.products VALUES ('19fc1004-2055-4ed0-9a19-caac865eba9f', '2021-11-27 09:41:27.500774+00', '2021-11-27 09:41:27.500774+00', NULL, 'Floral T-shirt', 'men', 'https://i.ibb.co/qMQ75QZ/floral-shirt.png', 20);
INSERT INTO public.products VALUES ('fcc48445-cb4d-4f87-a909-7eaa0a6abe10', '2021-11-27 09:41:27.500774+00', '2021-11-27 09:41:27.500774+00', NULL, 'Black & White Longsleeve', 'men', 'https://i.ibb.co/55z32tw/long-sleeve.png', 25);
INSERT INTO public.products VALUES ('e0dab840-4cca-453f-94c6-b392c8b9562a', '2021-11-27 09:41:27.500774+00', '2021-11-27 09:41:27.500774+00', NULL, 'Pink T-shirt', 'men', 'https://i.ibb.co/RvwnBL8/pink-shirt.png', 25);
INSERT INTO public.products VALUES ('5043dc81-7a0a-4539-880f-f03a449465a2', '2021-11-27 09:41:27.500774+00', '2021-11-27 09:41:27.500774+00', NULL, 'Jean Long Sleeve', 'men', 'https://i.ibb.co/VpW4x5t/roll-up-jean-shirt.png', 40);
INSERT INTO public.products VALUES ('f60f0dc2-7570-4f27-824f-173bc76ed58e', '2021-11-27 09:41:27.500774+00', '2021-11-27 09:41:27.500774+00', NULL, 'Burgundy T-shirt', 'men', 'https://i.ibb.co/mh3VM1f/polka-dot-shirt.png', 25);
INSERT INTO public.products VALUES ('bcfb388b-0c62-48d7-bc27-b81f97619f26', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', '2021-11-27 17:34:41.369778+00', 'Blue Beanie', 'hat', 'https://i.ibb.co/ypkgK0X/blue-beanie.png', 18);
INSERT INTO public.products VALUES ('8bf07181-7cb8-4572-9f21-407d0d3614e5', '2021-11-27 09:35:49.460578+00', '2021-11-27 09:35:49.460578+00', '2021-12-11 12:25:19.514596+00', 'Grey Brim', 'hat', 'https://i.ibb.co/RjBLWxB/grey-brim.png', 25);


--
-- Data for Name: refresh_token; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.refresh_token VALUES ('0665e318-f97f-4a83-b75a-d21ef9b5e8ce', '2021-11-27 09:44:31.302058+00', '2021-11-27 09:44:31.302058+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU3ODIyNzEsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6InVzZXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0.6xlvUYi-6gnai0hSemwI94ASVvH1sl6CUUL6FfesZrU', '2022-02-25 09:44:31.301561+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('cf9063cb-88bb-4f5d-9940-561864c96c06', '2021-11-27 10:24:18.513288+00', '2021-11-27 10:24:18.513288+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU3ODQ2NTgsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.Cr2Aw0VbXXPxY_-xDumMpl12p1UdVcIORQLMqgLxu6Q', '2022-02-25 10:24:18.512962+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('3fa00fad-f2c4-4c26-9bd9-c17e16550109', '2021-11-27 10:25:45.665008+00', '2021-11-27 10:25:45.665008+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU3ODQ3NDUsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.YbsitFBHUoIjv-YJ3FGYPsZ83rKMEy8RiPY0DO6txyU', '2022-02-25 10:25:45.664731+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('09d09468-89bf-4856-b7c9-6abddca5302e', '2021-11-27 10:29:38.299029+00', '2021-11-27 10:29:38.299029+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU3ODQ5NzgsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.iWSnsUwuYYetnlkgIqD8k8arkIElG-R6MRfhc9SU-Mo', '2022-02-25 10:29:38.298673+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('738ec0b4-354a-4083-aa97-80512db86f88', '2021-11-27 16:27:52.214963+00', '2021-11-27 16:27:52.214963+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU4MDY0NzIsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.YtxfaVb4ffogD5xvmc7h-cpQ9pCE4hXFqeq0Dg9luxU', '2022-02-25 16:27:52.214529+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('7098c649-a3d9-4f8a-96e9-434c7cf7758a', '2021-11-28 05:30:44.057816+00', '2021-11-28 05:30:44.057816+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU4NTM0NDQsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.wWWoK9qUEKmzy10GCxfA_r3UXy_uSKag0dXWy-oBeMY', '2022-02-26 05:30:44.05755+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('752af9a6-4264-48da-9020-7068e67580ac', '2021-11-28 07:18:48.747744+00', '2021-11-28 07:18:48.747744+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU4NTk5MjgsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.--emtrfy9J-elU0Iq9I3piaSZp_Hy-9PKRpsFAYVx80', '2022-02-26 07:18:48.747118+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('5fbd9465-de85-4165-b8b9-701030f7c213', '2021-11-28 10:12:30.206069+00', '2021-11-28 10:12:30.206069+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU4NzAzNTAsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.ogBJuEWqkWru1IFLZu8we_SOrR6fop1WUfM7pMf1-qY', '2022-02-26 10:12:30.205566+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('6354fe89-b634-4b66-9662-7a387af32fb1', '2021-11-28 10:12:48.509796+00', '2021-11-28 10:12:48.509796+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU4NzAzNjgsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.4gNDaj02nOieH-RaDYTw2T_SPXhgfv3dLJecPGMWgAM', '2022-02-26 10:12:48.50945+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('8bb72f3a-8709-4ddf-93a0-e345e9931ce6', '2021-11-28 10:22:52.71406+00', '2021-11-28 10:22:52.71406+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU4NzA5NzIsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.FeVk7PtM9Ps8BuwN7jFuJh0numhpxqkRHivHyTXO07E', '2022-02-26 10:22:52.713637+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('ce0a8cac-0db5-40a1-b3f9-c02361edf8e2', '2021-11-28 10:41:13.509705+00', '2021-11-28 10:41:13.509705+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU4NzIwNzMsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.4nSikwSCwPKN7aiwARy0z2DmgpwyNc6o8BbbfBCzook', '2022-02-26 10:41:13.509346+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('e038176b-6c5a-46c5-abe5-72875e037ef5', '2021-11-28 10:49:03.940957+00', '2021-11-28 10:49:03.940957+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU4NzI1NDMsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.t50HIfYhqlB-4yoS6jlInxuMsbAKy03gXoSG6Ll9gew', '2022-02-26 10:49:03.94063+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('0ae38259-b22d-403a-96f8-f8b46d8a4156', '2021-11-29 08:56:43.459328+00', '2021-11-29 08:56:43.459328+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU5NTIyMDMsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.tw_5SEbdN4ErA-nA6YrFnbLOloNdBpAVBelCju7XHjM', '2022-02-27 08:56:43.458834+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('60d64aec-e688-4b80-9531-dba9c3af26fd', '2021-12-04 09:52:29.013183+00', '2021-12-04 09:52:29.013183+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDYzODc1NDksImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19._8I3fU-3zl2X3NHPsi1T0WW1et-l-a43rkk1jkAhaI4', '2022-03-04 09:52:29.012607+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('1bd70a71-27e1-4638-a2df-68603bb46e16', '2021-12-04 10:01:17.080865+00', '2021-12-04 10:01:17.080865+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDYzODgwNzcsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.FeKlVkkNsW0RU2dOVaJBaOvg_AS48XtO0HfcI-idp7c', '2022-03-04 10:01:17.080447+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('ed878df6-e916-4181-9151-718b8136e6e9', '2021-12-04 13:19:19.450196+00', '2021-12-04 13:19:19.450196+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDYzOTk5NTksImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.vf9AfJ6IXZADAZsi7R1ETdPy4onRjRWbI-OawkS9Qn0', '2022-03-04 13:19:19.449839+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('eb9c2298-bbd5-4db0-9b61-b0e7043ea184', '2021-12-04 13:22:35.949303+00', '2021-12-04 13:22:35.949303+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY0MDAxNTUsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.wFTJ460B_-2h34VfhsfxboywyGQlCaSBOF1Lj3T-Mq8', '2022-03-04 13:22:35.948976+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('e22c6d55-97de-4aba-bee7-a4d2efce8418', '2021-12-04 17:39:23.776874+00', '2021-12-04 17:39:23.776874+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY0MTU1NjMsImlkIjoiNDkxY2NiMWYtODMxNS00YjYxLWE4NDMtMDVkYmE2YTU1MmE4Iiwicm9sZSI6InVzZXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0.bHsqEl8_aDA6Irbl51kfcbKo4WdF2iCYdwA6EfOAfWk', '2022-03-04 17:39:23.775785+00', '491ccb1f-8315-4b61-a843-05dba6a552a8');
INSERT INTO public.refresh_token VALUES ('28e0098b-12ac-4c1f-b74c-572b8c2f08f9', '2021-12-04 17:39:47.709283+00', '2021-12-04 17:39:47.709283+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY0MTU1ODcsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.9biBtMJ6mLqkke7kGI2riw4y87J-hRrYG9f9hl_oKUM', '2022-03-04 17:39:47.709043+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('3631c76e-1626-4d9f-866d-98129025d83e', '2021-12-04 18:06:22.722579+00', '2021-12-04 18:06:22.722579+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY0MTcxODIsImlkIjoiNDE2MzkyNjgtNWJlZC00NGY2LTk0NTUtYmVmMjBhODBiZjdmIiwicm9sZSI6InVzZXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0.1P8Utf1wpsQL84nOOojU6dc2QstvNumSbIR4rZpCRkk', '2022-03-04 18:06:22.722253+00', '41639268-5bed-44f6-9455-bef20a80bf7f');
INSERT INTO public.refresh_token VALUES ('30cc7eb1-98d8-4998-986f-949a25a2a0de', '2021-12-04 18:06:33.004795+00', '2021-12-04 18:06:33.004795+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY0MTcxOTMsImlkIjoiNDhmOTJiODMtNWI5Yi00OTJjLThlZDQtYWIwOTc2NWU5ZjFkIiwicm9sZSI6InVzZXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0.1kWeco6pVMVslHLmuQlMIs4xmeGBj4ZX2tfkmS2NpHs', '2022-03-04 18:06:33.004496+00', '48f92b83-5b9b-492c-8ed4-ab09765e9f1d');
INSERT INTO public.refresh_token VALUES ('eda43d4d-edc8-49fa-8cfc-c6a11b824507', '2021-12-04 18:14:57.902709+00', '2021-12-04 18:14:57.902709+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY0MTc2OTcsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.prKcqaG1X3hrc3A5LsZTP3Eb-AzdhURgjTOEI9vM-r8', '2022-03-04 18:14:57.902484+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('bc969675-3ac3-4e4f-a151-bb2ffb6d0ad7', '2021-12-06 08:03:41.411017+00', '2021-12-06 08:03:41.411017+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY1NTM4MjEsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.pa0rqcKSN2XQjT3nV2ZdgBUM__u5CyAkpDVH_N6RRPs', '2022-03-06 08:03:41.410716+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('60eb5e18-d3ec-42d3-af3d-2c953868d212', '2021-12-06 08:10:04.907236+00', '2021-12-06 08:10:04.907236+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY1NTQyMDQsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.kuvtvJ3HvUcjmE-Lh1LjdfTxr34GnbB76Ob4qTBgkdI', '2022-03-06 08:10:04.906706+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('51191605-45fb-4cd3-a69e-f8db4af72626', '2021-12-06 08:13:49.588623+00', '2021-12-06 08:13:49.588623+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY1NTQ0MjksImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.AXg2gXhW1gW2Ojv8cfR4C5a7LoqGbO9uSqHZjmsmYhY', '2022-03-06 08:13:49.588246+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('546faea1-3232-42d9-82a6-f893f64059a0', '2021-12-07 12:55:49.144562+00', '2021-12-07 12:55:49.144562+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY2NTc3NDksImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.RmZ8bvAydN8pBUF8mGO_1AGzt3DeG-ZsxYLCYyzOBUo', '2022-03-07 12:55:49.144102+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('f574a11d-f88d-4b57-a1af-34fe7786e2fb', '2021-12-08 05:58:18.258354+00', '2021-12-08 05:58:18.258354+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY3MTkwOTgsImlkIjoiNTI2MWI1MzktMTljNi00ZDZiLWE0M2ItZGQ3YmYzM2IwM2UxIiwicm9sZSI6InVzZXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0.sm_pie4zOO1sVUe6SfFgk4bx_hWcJ9J1MNgc74wk5YE', '2022-03-08 05:58:18.257867+00', '5261b539-19c6-4d6b-a43b-dd7bf33b03e1');
INSERT INTO public.refresh_token VALUES ('1ae924c4-ff50-48e9-a490-b4e22e16b093', '2021-12-08 06:05:38.230361+00', '2021-12-08 06:05:38.230361+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY3MTk1MzgsImlkIjoiNTI2MWI1MzktMTljNi00ZDZiLWE0M2ItZGQ3YmYzM2IwM2UxIiwicm9sZSI6InVzZXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0.doWX5PzNYo2YHFfG4IrffIAWHvezq4ZpJx2MFiL1auY', '2022-03-08 06:05:38.22992+00', '5261b539-19c6-4d6b-a43b-dd7bf33b03e1');
INSERT INTO public.refresh_token VALUES ('04ed3181-7993-4f2d-b8f6-7a95f6ffe317', '2021-12-08 06:11:57.135937+00', '2021-12-08 06:11:57.135937+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY3MTk5MTcsImlkIjoiNTI2MWI1MzktMTljNi00ZDZiLWE0M2ItZGQ3YmYzM2IwM2UxIiwicm9sZSI6InVzZXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0.BIukzWeVsVk5TaosXisxyFSZwcO4ECtMnSxaTk1uMxs', '2022-03-08 06:11:57.135653+00', '5261b539-19c6-4d6b-a43b-dd7bf33b03e1');
INSERT INTO public.refresh_token VALUES ('2158fa3e-e0f5-48a2-8d48-1de44f1f4ecc', '2021-12-08 07:31:41.722802+00', '2021-12-08 07:31:41.722802+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY3MjQ3MDEsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.LxhmZcQqeWZKoHxAwvw33arOG-kgLmc9Tj4S_0Z7FSw', '2022-03-08 07:31:41.722447+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('31571110-37f3-450a-beca-91fdacf18075', '2021-12-08 10:36:52.797528+00', '2021-12-08 10:36:52.797528+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY3MzU4MTIsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.4ScbTPfcVA6Ayjs_eUEqgrRQcyIbLORTfpU8491V3OE', '2022-03-08 10:36:52.796927+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('e43decc6-fe98-4543-a0f7-03e70e527dea', '2021-12-08 11:13:30.91612+00', '2021-12-08 11:13:30.91612+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY3MzgwMTAsImlkIjoiNTI2MWI1MzktMTljNi00ZDZiLWE0M2ItZGQ3YmYzM2IwM2UxIiwicm9sZSI6InVzZXIiLCJzY29wZXMiOlsiYXBpOnJlYWQiLCJhcGk6d3JpdGUiXX0.Du7pXfhzdeRkFrDaJaVd_EJaWPE76IRmdYn01Q2OJtg', '2022-03-08 11:13:30.91579+00', '5261b539-19c6-4d6b-a43b-dd7bf33b03e1');
INSERT INTO public.refresh_token VALUES ('8e099d9b-7764-4974-8ec8-b61adce0695e', '2021-12-11 07:27:08.531216+00', '2021-12-11 07:27:08.531216+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY5ODM2MjgsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.94p8U22qvTIQEBOm40ESX44DDAHy_YMIcSSWrJAgbxQ', '2022-03-11 07:27:08.530933+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('912cb62c-601e-46bf-8d3a-f60a0babffef', '2021-12-11 08:09:09.113024+00', '2021-12-11 08:09:09.113024+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY5ODYxNDksImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.ie7zWJY_FYI0KHCceY3ki6hcKX3Zd9dY1_BNjenxsVM', '2022-03-11 08:09:09.112444+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('be67cd27-89f9-40ea-8ba4-9eac305e0c60', '2021-12-11 11:16:33.367887+00', '2021-12-11 11:16:33.367887+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY5OTczOTMsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.8x5fFxLKzTglcWLOPnrg2Kp99RcKdCzwCyGOQeXMGWI', '2022-03-11 11:16:33.367355+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('009ac771-cd51-4122-97d1-7a3738e4e73d', '2021-12-11 11:50:54.220534+00', '2021-12-11 11:50:54.220534+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDY5OTk0NTQsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.TqKoOx7_7zDLyZozqKfVDM3GZ59sXfy0DGt4jXU2Tao', '2022-03-11 11:50:54.219962+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('444c22c5-efa7-4163-83e9-fee4f002eaef', '2021-12-18 13:23:45.073066+00', '2021-12-18 13:23:45.073066+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDc2MDk4MjUsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.Ou6cJlICAAB1v2co1oQy0FtswQ3AL45KX034Z7ypI_8', '2022-03-18 13:23:45.072752+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('cd2b97d8-4523-443e-83be-1aa80780a115', '2021-12-19 17:30:04.087092+00', '2021-12-19 17:30:04.087092+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDc3MTEwMDQsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.yuhjioFIifi6abVXDXdpcUl-H2SwzKSIaObo9J6scU0', '2022-03-19 17:30:04.086046+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('38d99318-1cb0-4c9a-b746-5ee7fcc521e2', '2021-12-20 06:37:16.329179+00', '2021-12-20 06:37:16.329179+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDc3NTgyMzYsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.32__uAlavvlXXojeTIOSr_ilKatb57QRZRCMiI75d9U', '2022-03-20 06:37:16.328872+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');
INSERT INTO public.refresh_token VALUES ('83c3170e-1cf4-4d12-9511-84d516c62bd1', '2021-12-20 06:37:23.569167+00', '2021-12-20 06:37:23.569167+00', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDc3NTgyNDMsImlkIjoiY2M3NjMzMGItYTYzYS00OTI3LWJiZjctN2I5OGFhNDljYmZlIiwicm9sZSI6ImFkbWluIiwic2NvcGVzIjpbImFwaTpyZWFkIiwiYXBpOndyaXRlIl19.JBHzRkVCWs1-DShP9CdggLV1BCqhAeEZlOv8OUP9ku8', '2022-03-20 06:37:23.56894+00', 'cc76330b-a63a-4927-bbf7-7b98aa49cbfe');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.users VALUES ('48f92b83-5b9b-492c-8ed4-ab09765e9f1d', '2021-12-04 19:06:33.000352', '2021-12-08 11:48:02.671859', NULL, 'MORIN', 'Guillaume', 'guillaumemm@gmail.com', '$2a$06$kAiNCUBlyttiE7jetouX6.KeFjG7JG.OQaLXqBPqTG2sxqG.fKIa.', 'user', NULL, NULL, NULL, NULL, NULL, 'QTGP6');
INSERT INTO public.users VALUES ('3edc936f-bfd7-4cec-a837-13fd07b480cb', '2021-11-27 12:02:19.032877', '2021-11-27 12:09:29.466585', '2021-11-27 17:28:12.267583', 'MORIN1', 'Guillaume', 'guillaumemm@gmail.com', '$2a$06$A5a9kEFu9A..6RK01pAUDeRbt68vibAo.r8GVnOWsDZ8q.EVdA58e', 'pro', '2021-11-04', '++262974824', NULL, NULL, NULL, NULL);
INSERT INTO public.users VALUES ('5261b539-19c6-4d6b-a43b-dd7bf33b03e1', '2021-12-08 06:58:18.248986', '2021-12-08 12:13:20.237095', NULL, 'Morin', 'Guillaume', 'guillaume.morin@epitech.eu', '$2a$06$uBrsoKSsivztmVqX0MX9o.cR1V1DdlvmM7qZlxUWXjVBnyKdKcgTW', 'user', NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.users VALUES ('491ccb1f-8315-4b61-a843-05dba6a552a8', '2021-12-04 18:39:21.169518', '2021-12-04 18:39:21.169518', '2021-12-20 07:35:54.262468', 'Morin', 'Guillaume', 'guillaume.morin974@gmail.com', '$2a$06$GXPCxl9JG1TlwxNA9KM4c.KPNCjzFkOfRqXwaSnK.jLivmSJ7ZiJO', 'user', NULL, NULL, 'eyJhbGciOiJSUzI1NiIsImtpZCI6IjgwNTg1Zjk5MjExMmZmODgxMTEzOTlhMzY5NzU2MTc1YWExYjRjZjkiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiR3VpbGxhdW1lIE1vcmluIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hLS9BT2gxNEdncUhrTnB1SHlNSXo1SC1oMFFWWEdpSml2TkRxZVhHSG05eHgzeEVnPXM5Ni1jIiwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL3N0YXJ0ZXJwYWNrLTg1YWJjIiwiYXVkIjoic3RhcnRlcnBhY2stODVhYmMiLCJhdXRoX3RpbWUiOjE2Mzg2Mzk1NjAsInVzZXJfaWQiOiJnSEQ0c1UyVEtTV1ZsbklJdHhvd1R0NnRLNDIzIiwic3ViIjoiZ0hENHNVMlRLU1dWbG5JSXR4b3dUdDZ0SzQyMyIsImlhdCI6MTYzODYzOTU2MCwiZXhwIjoxNjM4NjQzMTYwLCJlbWFpbCI6Imd1aWxsYXVtZS5tb3Jpbjk3NEBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJnb29nbGUuY29tIjpbIjEwNTQ5MjYyNTE4NzU4NDY5MjMwMyJdLCJlbWFpbCI6WyJndWlsbGF1bWUubW9yaW45NzRAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.i7mpG1iLgrXp8dEdQ3yGDKekcUwR3Py40DYNar9kwoj-3I7NcFsIi0dLUyBXDbDdE6gMWCvEkyWcIqFf3lFAYmttcwNNCrQvyKGPk5Nwsn8rprcWUHNi3ErWgRJRoh0XVtBhDkiRUMypqmu5i3yE1w2YAnHIUHHiAVNWVOcmok7p4AjosDP_s-aN9b2JVn0M6up49XDbM1betxazqlQO6K8VkD5kEx-buMAOnmlhEcL9wa3jkcO82oPiBYaSo-dxqC3jruH2iWRkfDWFA9WXG2KlI5eXE9i59QGrGq4tUVLhWrf5mCcINs1iiCN4-r8vqWNutCrt1UF-IpRbpG3PxA', 'gHD4sU2TKSWVlnIItxowTt6tK423', 'google.com', NULL);
INSERT INTO public.users VALUES ('dc4cb9b3-d8bf-4dd9-abfa-586e794ec491', '2021-11-27 11:41:03.944383', '2021-11-27 18:28:19.085716', NULL, 'Grondin', 'Jeanne', 'jeanne@gmail.com', '$2a$06$yPVg.PJ4AuZNS3IG2AgHwuCn1ihFC2GmwBUh1jSEbVbPQtC7prsja', 'pro', '2002-05-30', NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.users VALUES ('927057cc-0ddb-4921-b0ef-d2b58b9823d1', '2021-12-04 11:19:08.554853', '2021-12-04 11:19:08.554853', '2021-12-04 11:19:19.515721', 'MORIN', 'Guillaume', 'guillaumemm@gmail.com', '$2a$06$a4S3imHAsXz5oY4K7TSYp.DJQYAmKkkmJr2FEMplXBs4accHZLDY.', 'pro', NULL, '+11111111111', NULL, NULL, NULL, NULL);
INSERT INTO public.users VALUES ('cf48b61a-ec6f-4378-9961-e55fc91276f1', '2021-12-04 11:26:15.802568', '2021-12-04 12:11:21.783009', '2021-12-04 12:11:32.615432', 'MORIN', 'Guillaume', 'guillaumemm@gmail.com', '$2a$06$zCz3/ykYe1DVMSI9uQq52OyaEkig73xUaPZ2j4fJoPlmq06eMTZxy', 'user', NULL, '11111111111', NULL, NULL, NULL, NULL);
INSERT INTO public.users VALUES ('cc76330b-a63a-4927-bbf7-7b98aa49cbfe', '2021-11-27 09:43:11.306186', '2021-12-04 12:21:40.84465', NULL, 'Morin', 'Guillaume', 'guillaume@gmail.com', '$2a$06$JQoe1nXKDu5GSJ/cZYUpxuFB3JbaPea//.T4qUv.TOOWgCgeZdwZu', 'admin', '2021-11-03', '693472824', NULL, NULL, NULL, NULL);
INSERT INTO public.users VALUES ('41639268-5bed-44f6-9455-bef20a80bf7f', '2021-12-04 19:06:22.717648', '2021-12-04 19:06:22.717648', '2021-12-06 09:10:13.98962', 'MORIN', 'Guillaume', 'guillaume1@gmail.com', '$2a$06$Ht9F/uO4tA7MVyAX68VXqO1/2NYv8aw5RZ9QS4KEtmjAkWzEU0QbO', 'user', NULL, NULL, NULL, NULL, NULL, NULL);


--
-- Name: files files_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: refresh_token refresh_token_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.refresh_token
    ADD CONSTRAINT refresh_token_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: refresh_token refresh_token_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.refresh_token
    ADD CONSTRAINT refresh_token_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

