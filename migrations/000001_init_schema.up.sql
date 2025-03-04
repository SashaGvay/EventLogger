CREATE TABLE IF NOT EXISTS public.weather_logs (
    id SERIAL PRIMARY KEY,
    temperature FLOAT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);