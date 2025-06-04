CREATE DATABASE events;

\c events

CREATE TABLE base_plans (
  id INT PRIMARY KEY,
  sell_mode VARCHAR (50),
  queue_limit INT DEFAULT 1000,
  queue_enabled BOOLEAN DEFAULT FALSE,
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

CREATE INDEX index_plan_id ON plans(base_plan_id);
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

CREATE INDEX index_plan_id ON zones(plan_id);

CREATE TABLE clients_queue (
  plan_id INT NOT NULL,
  uuid INT PRIMARY KEY,
  expiration_at TIMESTAMP SET DEFAULT (NOW() + INTERVAL '1 hour'),
  PRIMARY KEY (plan_id, uuid)
);
