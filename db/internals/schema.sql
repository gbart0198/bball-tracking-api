--
-- PostgreSQL database dump
--

-- Dumped from database version 14.15 (Ubuntu 14.15-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.15 (Ubuntu 14.15-0ubuntu0.22.04.1)

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
-- Name: drills; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.drills (
    drill_id uuid DEFAULT gen_random_uuid() NOT NULL,
    drill_name character varying(50) NOT NULL,
    category character varying(50),
    difficulty character(25)
);


ALTER TABLE public.drills OWNER TO postgres;

--
-- Name: goals; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goals (
    goal_id uuid DEFAULT gen_random_uuid() NOT NULL,
    goal_type character varying(50) NOT NULL,
    goal_name character varying(50) NOT NULL
);


ALTER TABLE public.goals OWNER TO postgres;

--
-- Name: TABLE goals; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.goals IS 'lookup table for different types of goals';


--
-- Name: player_goals; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.player_goals (
    player_goal_id uuid DEFAULT gen_random_uuid() NOT NULL,
    player_id uuid NOT NULL,
    goal_id uuid NOT NULL,
    current_value integer,
    goal_value integer NOT NULL
);


ALTER TABLE public.player_goals OWNER TO postgres;

--
-- Name: TABLE player_goals; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.player_goals IS 'tracks players goals';


--
-- Name: player_performances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.player_performances (
    player_performance_id uuid DEFAULT gen_random_uuid() NOT NULL,
    player_id uuid NOT NULL,
    drill_id uuid NOT NULL,
    date date NOT NULL,
    attempts integer,
    successful integer
);


ALTER TABLE public.player_performances OWNER TO postgres;

--
-- Name: session_performances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.session_performances (
    session_performance_id uuid DEFAULT gen_random_uuid() NOT NULL,
    session_id uuid NOT NULL,
    player_performance_id uuid NOT NULL
);


ALTER TABLE public.session_performances OWNER TO postgres;

--
-- Name: TABLE session_performances; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.session_performances IS 'tracks performances on a session level';


--
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    session_id uuid DEFAULT gen_random_uuid() NOT NULL,
    session_type character varying(25) NOT NULL,
    date timestamp with time zone NOT NULL,
    location character varying(50),
    user_id uuid NOT NULL,
    session_name character varying(50) NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- Name: TABLE sessions; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.sessions IS 'stores sessions for player workouts';


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL,
    username character varying(100) NOT NULL,
    passhash character varying(500) NOT NULL,
    email character varying(100) NOT NULL,
    firstname character varying,
    lastname character varying
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: drills drills_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.drills
    ADD CONSTRAINT drills_pkey PRIMARY KEY (drill_id);


--
-- Name: goals goal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goals
    ADD CONSTRAINT goal_pkey PRIMARY KEY (goal_id);


--
-- Name: player_goals player_goal_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_goals
    ADD CONSTRAINT player_goal_pkey PRIMARY KEY (player_goal_id);


--
-- Name: player_performances player_performances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_performances
    ADD CONSTRAINT player_performances_pkey PRIMARY KEY (player_performance_id);


--
-- Name: session_performances session_performance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.session_performances
    ADD CONSTRAINT session_performance_pkey PRIMARY KEY (session_performance_id);


--
-- Name: sessions session_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT session_pkey PRIMARY KEY (session_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: player_performances drill; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_performances
    ADD CONSTRAINT drill FOREIGN KEY (drill_id) REFERENCES public.drills(drill_id) ON DELETE CASCADE;


--
-- Name: player_goals goal; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_goals
    ADD CONSTRAINT goal FOREIGN KEY (goal_id) REFERENCES public.goals(goal_id);


--
-- Name: player_performances player; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_performances
    ADD CONSTRAINT player FOREIGN KEY (player_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- Name: player_goals player; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.player_goals
    ADD CONSTRAINT player FOREIGN KEY (player_id) REFERENCES public.users(user_id);


--
-- Name: session_performances player_performance; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.session_performances
    ADD CONSTRAINT player_performance FOREIGN KEY (player_performance_id) REFERENCES public.player_performances(player_performance_id);


--
-- Name: session_performances session; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.session_performances
    ADD CONSTRAINT session FOREIGN KEY (session_id) REFERENCES public.sessions(session_id);


--
-- Name: sessions session_owner; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT session_owner FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

