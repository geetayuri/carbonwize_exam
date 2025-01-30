carbonwize-test

ข้อ 2.
/api/carbon/footprint/calculate as POST Method
db config inside .env 
- import db.sql dbname: carbonwize_DB
- หรือทำข้อ 3: run create and insert virtual data ก่อนลองยิงในข้อ 2 นี้

ข้อ 3.
CREATE TABLE emission_factors (
    id SERIAL PRIMARY KEY, 
    activity_type VARCHAR(50),
    name character varying(50),
    type character varying(50),
    unit character varying(50),
    emission_factor numeric,
    source character varying(200),
    year character varying(50)
);

INSERT INTO emission_factors (activity_group, name, type, unit, emission_factor, source, year) VALUES
('transportation', 'gasoline', 'car', 'km', 0.17553, 'UK Department for Business, Energy & Industrial Strategy (BEIS)', '2022'),
('transportation', 'diesel', 'bus', 'km', 1.17772, 'UK Department for Business, Energy & Industrial Strategy (BEIS)', '2022'),
('energy consumption', '', 'electricity', 'kWh', 0.500, 'IPCC', '2023'),
('energy consumption', '', 'electricity', 'kWh', 0.450, 'DEFRA', '2022');

SELECT id, activity_group, name, type, unit, emission_factor, source, year
	FROM public.emission_factors;

ข้อ 4.
CREATE TABLE emission_data (
    id SERIAL PRIMARY KEY,
    activity_type VARCHAR(50),
    input_value NUMERIC,
    emission_factor NUMERIC,
    total_emission NUMERIC GENERATED ALWAYS AS (input_value * emission_factor) STORED
);

เทสด้วย:
INSERT INTO public.emission_data(
	activity_type, input_value, emission_factor)
	VALUES ('energy consumption', 150, 0.450);

SELECT id, activity_type, input_value, emission_factor, total_emission
	FROM public.emission_data;

ข้อ 5.
สร้าง func ใหม่ เพื่อรับค่าตามโจทย์ที่ ./virtualfuncforfifth
fmt show ค่าก่อน run server ด้านบน

ีunit test func นี้ที่ ./tests/fifth
go test -v ./tests