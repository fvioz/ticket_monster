CREATE DATABASE events;

\c events

CREATE TABLE base_plans (
  id INT PRIMARY KEY,
  sell_mode VARCHAR (50),
  title text
);

CREATE INDEX index_base_plans_sell_mode ON base_plans(sell_mode);

CREATE TABLE plans (
  id INT PRIMARY KEY,
  base_plan_id INT NOT NULL,
  plan_start_date TIMESTAMP,
  plan_end_date TIMESTAMP,
  sell_to TIMESTAMP,
  sold_out BOOLEAN
);

CREATE INDEX index_plans_dates ON plans(plan_start_date, plan_end_date);

CREATE TABLE zones (
  id INT PRIMARY KEY,
  plan_id INT NOT NULL,
  capacity INT,
  price FLOAT,
  name TEXT,
  sell_to TIMESTAMP,
  numbered BOOLEAN
);
