CREATE TABLE events (
                        id SERIAL PRIMARY KEY,
                        user_id INT NOT NULL,
                        title VARCHAR(255) NOT NULL,
                        date TIMESTAMP NOT NULL,
                        description TEXT,
                        status VARCHAR(50) NOT NULL,
                        FOREIGN KEY (user_id) REFERENCES users(id)
);