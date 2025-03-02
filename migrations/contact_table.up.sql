CREATE TABLE contacts (
                          id SERIAL PRIMARY KEY,
                          user_id INT NOT NULL,
                          name VARCHAR(255) NOT NULL,
                          phone VARCHAR(20),
                          email VARCHAR(255),
                          FOREIGN KEY (user_id) REFERENCES users(id)
);