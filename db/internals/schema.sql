--
-- PostgreSQL database dump
--

-- Dumped from database version 12.22 (Debian 12.22-1.pgdg120+1)
-- Dumped by pg_dump version 17.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: admin
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO admin;

--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: drills; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.drills (
    drill_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    drill_name character(50) NOT NULL,
    category character varying(50),
    difficulty character varying(25)
);


ALTER TABLE public.drills OWNER TO admin;

--
-- Name: goal_categories; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.goal_categories (
    goal_category_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    category character varying(50) NOT NULL
);


ALTER TABLE public.goal_categories OWNER TO admin;

--
-- Name: goals; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.goals (
    goal_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    goal_type character varying(50) NOT NULL,
    goal_name character varying(50) NOT NULL
);


ALTER TABLE public.goals OWNER TO admin;

--
-- Name: player_goals; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.player_goals (
    player_goal_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    player_id uuid NOT NULL,
    drill_id uuid NOT NULL,
    current_value integer,
    goal_value integer NOT NULL,
    goal_category_id uuid NOT NULL,
    goal_name character varying(50) NOT NULL,
    goal_description character varying(50),
    completed boolean DEFAULT false NOT NULL
);


ALTER TABLE public.player_goals OWNER TO admin;

--
-- Name: player_performances; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.player_performances (
    player_performance_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    player_id uuid NOT NULL,
    drill_id uuid NOT NULL,
    date date NOT NULL,
    attempts integer,
    successful integer
);


ALTER TABLE public.player_performances OWNER TO admin;

--
-- Name: session_performances; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.session_performances (
    session_performance_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    session_id uuid NOT NULL,
    player_performance_id uuid NOT NULL
);


ALTER TABLE public.session_performances OWNER TO admin;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.sessions (
    session_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    session_type character varying(25) NOT NULL,
    date date NOT NULL,
    location character varying(50),
    user_id uuid NOT NULL,
    session_name character varying(50) NOT NULL
);


ALTER TABLE public.sessions OWNER TO admin;

--
-- Name: users; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.users (
    user_id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    username character varying(100) NOT NULL,
    passhash character varying(500) NOT NULL,
    email character varying(50) NOT NULL,
    firstname character varying(25),
    lastname character varying(25)
);


ALTER TABLE public.users OWNER TO admin;

--
-- Name: drills drills_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.drills
    ADD CONSTRAINT drills_pkey PRIMARY KEY (drill_id);


--
-- Name: goal_categories goal_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.goal_categories
    ADD CONSTRAINT goal_categories_pkey PRIMARY KEY (goal_category_id);


--
-- Name: goals goals_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.goals
    ADD CONSTRAINT goals_pkey PRIMARY KEY (goal_id);


--
-- Name: player_goals player_goals_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.player_goals
    ADD CONSTRAINT player_goals_pkey PRIMARY KEY (player_goal_id);


--
-- Name: player_performances player_performances_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.player_performances
    ADD CONSTRAINT player_performances_pkey PRIMARY KEY (player_performance_id);


--
-- Name: session_performances session_performances_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.session_performances
    ADD CONSTRAINT session_performances_pkey PRIMARY KEY (session_performance_id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (session_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: player_performances drill_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.player_performances
    ADD CONSTRAINT drill_fk FOREIGN KEY (drill_id) REFERENCES public.drills(drill_id);


--
-- Name: player_goals drill_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.player_goals
    ADD CONSTRAINT drill_fk FOREIGN KEY (drill_id) REFERENCES public.drills(drill_id) NOT VALID;


--
-- Name: session_performances fk_player_performance; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.session_performances
    ADD CONSTRAINT fk_player_performance FOREIGN KEY (player_performance_id) REFERENCES public.player_performances(player_performance_id);


--
-- Name: session_performances fk_session; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.session_performances
    ADD CONSTRAINT fk_session FOREIGN KEY (session_id) REFERENCES public.sessions(session_id);


--
-- Name: player_goals goal_category_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.player_goals
    ADD CONSTRAINT goal_category_fk FOREIGN KEY (goal_category_id) REFERENCES public.goal_categories(goal_category_id);


--
-- Name: sessions user_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: player_performances user_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.player_performances
    ADD CONSTRAINT user_fk FOREIGN KEY (player_id) REFERENCES public.users(user_id);


--
-- Name: player_goals user_fk; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.player_goals
    ADD CONSTRAINT user_fk FOREIGN KEY (player_id) REFERENCES public.users(user_id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: admin
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

