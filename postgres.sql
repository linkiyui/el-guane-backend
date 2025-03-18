CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ci VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role VARCHAR(255)
);

CREATE TABLE consultas (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    doctor_id UUID,
    date TIMESTAMP,
    paciente_id UUID,
    grade VARCHAR(1),
    CONSTRAINT fk_doctor FOREIGN KEY(doctor_id) REFERENCES doctors(id),
    CONSTRAINT fk_paciente FOREIGN KEY(paciente_id) REFERENCES pacientes(id)
);

CREATE TABLE doctors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255)
);

CREATE TABLE pacientes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    age INT,
    sex VARCHAR(10)
);

CREATE TABLE leves (
    consulta_id UUID PRIMARY KEY,
    diagnosis TEXT,
    analisis BOOLEAN,
    CONSTRAINT fk_consulta FOREIGN KEY(consulta_id) REFERENCES consultas(id)
);

CREATE TABLE graves (
    consulta_id UUID PRIMARY KEY,
    sintomas TEXT,
    ingreso BOOLEAN,
    temp FLOAT,
    pulse FLOAT,
    press_min INT,
    press_max INT,
    CONSTRAINT fk_consulta FOREIGN KEY(consulta_id) REFERENCES consultas(id)
);