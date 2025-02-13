PGDMP                       }            carbonwize_DB    16.4    16.4     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    49211    carbonwize_DB    DATABASE     �   CREATE DATABASE "carbonwize_DB" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Thai_Thailand.874';
    DROP DATABASE "carbonwize_DB";
                postgres    false            �            1259    49213    emission_data    TABLE     �   CREATE TABLE public.emission_data (
    id integer NOT NULL,
    activity_type character varying(50),
    input_value numeric,
    emission_factor numeric,
    total_emission numeric GENERATED ALWAYS AS ((input_value * emission_factor)) STORED
);
 !   DROP TABLE public.emission_data;
       public         heap    postgres    false            �            1259    49212    emission_data_id_seq    SEQUENCE     �   CREATE SEQUENCE public.emission_data_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.emission_data_id_seq;
       public          postgres    false    216            �           0    0    emission_data_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.emission_data_id_seq OWNED BY public.emission_data.id;
          public          postgres    false    215            �            1259    49224    emission_factors    TABLE     ,  CREATE TABLE public.emission_factors (
    id integer NOT NULL,
    activity_group character varying(50),
    name character varying(50),
    type character varying(50),
    unit character varying(50),
    emission_factor numeric,
    source character varying(200),
    year character varying(50)
);
 $   DROP TABLE public.emission_factors;
       public         heap    postgres    false            �            1259    49223    emission_factors_id_seq    SEQUENCE     �   CREATE SEQUENCE public.emission_factors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.emission_factors_id_seq;
       public          postgres    false    218            �           0    0    emission_factors_id_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.emission_factors_id_seq OWNED BY public.emission_factors.id;
          public          postgres    false    217            U           2604    49216    emission_data id    DEFAULT     t   ALTER TABLE ONLY public.emission_data ALTER COLUMN id SET DEFAULT nextval('public.emission_data_id_seq'::regclass);
 ?   ALTER TABLE public.emission_data ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    216    216            W           2604    49227    emission_factors id    DEFAULT     z   ALTER TABLE ONLY public.emission_factors ALTER COLUMN id SET DEFAULT nextval('public.emission_factors_id_seq'::regclass);
 B   ALTER TABLE public.emission_factors ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    218    218            �          0    49213    emission_data 
   TABLE DATA           X   COPY public.emission_data (id, activity_type, input_value, emission_factor) FROM stdin;
    public          postgres    false    216   a       �          0    49224    emission_factors 
   TABLE DATA           o   COPY public.emission_factors (id, activity_group, name, type, unit, emission_factor, source, year) FROM stdin;
    public          postgres    false    218   �       �           0    0    emission_data_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.emission_data_id_seq', 2, true);
          public          postgres    false    215            �           0    0    emission_factors_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.emission_factors_id_seq', 4, true);
          public          postgres    false    217            Y           2606    49221     emission_data emission_data_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.emission_data
    ADD CONSTRAINT emission_data_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.emission_data DROP CONSTRAINT emission_data_pkey;
       public            postgres    false    216            [           2606    49231 &   emission_factors emission_factors_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY public.emission_factors
    ADD CONSTRAINT emission_factors_pkey PRIMARY KEY (id);
 P   ALTER TABLE ONLY public.emission_factors DROP CONSTRAINT emission_factors_pkey;
       public            postgres    false    218            �   A   x�3�,)J�+.�/*I,����425�4�3654�2�L�K-J�TH��+.�- ���ML�b���� �#      �   �   x��ϱ
�0����DAJ�Z:[[���"..�=5�MJ.|{�����v��w��`�W^[WE���V�Dgi��a���=�/��|��#��Ҡ����W��;�:�w=��4/������o��H��y��+��������^vX�N���C{��7S!�ڮV�v�?����\�o�1�Uxo�     