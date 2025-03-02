CREATE TABLE reminders (
                           id SERIAL PRIMARY KEY,
                           event_id INT NOT NULL,
                           reminder_time TIMESTAMP NOT NULL,
                           notification_method VARCHAR(50) NOT NULL,
                           FOREIGN KEY (event_id) REFERENCES events(id)
);