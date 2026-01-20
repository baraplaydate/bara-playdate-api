CREATE SCHEMA bara_playdate;

CREATE TABLE bara_playdate.mst_param_global (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  param_code varchar(191),
  param_name varchar(191),
  param_group varchar(191) NOT NULL,
  param_value varchar(191) ,
  param_type varchar(191),
  description TEXT,
  version INT NOT NULL DEFAULT 1, 
  sort INT,
  created_by varchar(191) NOT NULL,
  updated_by varchar(191),
  is_active varchar(20) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at timestamp(0)
  --UNIQUE (param_code, param_group)
);



CREATE TABLE bara_playdate.mst_user (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  role_id varchar(191) NOT NULL, --ambil data dari table mst_acc_role
  username varchar(191) NOT NULL,
  email varchar(191) NOT NULL,
  fullname varchar(191) NOT NULL,
  is_gender varchar(1), --ambil data dari table mst_param_global (param_group = 'GENDER_TYPE')
  address  TEXT,
  hp_number varchar(20),
  password varchar(191) NOT NULL,
  url_signature varchar(191),
  date_activation timestamp(0),
  email_verified_at timestamp(0),
  birth_date DATE,
  hire_date DATE,
  created_by varchar(191) NOT NULL,
  updated_by varchar(191),
  is_active varchar(20) NOT NULL, --ambil data dari table mst_param_global (param_group = 'EMPL_STATUS')
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at timestamp(0)
)
;
ALTER TABLE bara_playdate.mst_user ADD CONSTRAINT mst_user_unique UNIQUE ("email");
ALTER TABLE bara_playdate.mst_user ADD CONSTRAINT mst_user_username_unique UNIQUE ("username");

CREATE TABLE bara_playdate.mst_acc_role (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  role_name VARCHAR(100) NOT NULL UNIQUE,
  description TEXT,
  created_by varchar(191) NOT NULL,
  updated_by varchar(191),
  is_active varchar(20) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at timestamp(0)
);


INSERT INTO bara_playdate.mst_user 
(username,role_id,email,fullname,is_gender,password,created_by,is_active)
VALUES 
('superadmin', 'd40d2e7b-3d8e-4c4c-b010-4cb9798f1bd0','superadmin@gmail.com','Superadmin' ,'L', '$2a$10$EuXvsCws9GKL6JTOujpXy.GI51WvYYxFED0HFAZsyXnY/1tplyAxu', 'SYSTEM', 'ACTIVE') ;



INSERT INTO bara_playdate.mst_acc_role
(role_name, created_by, is_active)
VALUES
('Superadmin', 'system', 'ACTIVE');