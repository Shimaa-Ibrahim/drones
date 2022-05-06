TRUNCATE TABLE drones.drone_models CASCADE;
TRUNCATE TABLE drones.drone_states CASCADE;
TRUNCATE TABLE drones.medications CASCADE;

WITH models_list AS (
    SELECT '{"Lightweight", "Middleweight", "Cruiserweight", "Heavyweight"}'::TEXT[]model
)
INSERT INTO drones.drone_models (name)
SELECT model[n] FROM models_list, generate_series(1, 4) as n;

WITH states_list AS (
    SELECT '{"IDLE", "LOADING", "LOADED", "DELIVERING", "DELIVERED", "RETURNING"}'::TEXT[]stat
)
INSERT INTO drones.drone_states (name)
SELECT stat[n] FROM states_list, generate_series(1, 6) as n;


INSERT INTO drones.medications (name, weight, code, image_path)
SELECT 'medication ' || n as name, ceil(random() * 250), md5(random()::text), md5(random()::text)
FROM generate_series(1, 25) as n;