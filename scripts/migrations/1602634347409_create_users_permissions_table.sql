CREATE TABLE users_permissions
	(
		id SERIAL NOT NULL PRIMARY KEY,
		user_id int NOT NULL,
		user_name varchar(255) NOT NULL,
		user_email varchar(255) NOT NULL,
		permission_id int NOT NULL,
		permission_role varchar(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP WITH TIME ZONE,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (permission_id) REFERENCES permissions(id)
	);