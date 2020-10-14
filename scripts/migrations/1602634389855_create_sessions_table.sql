CREATE TABLE sessions 
	(
		id SERIAL NOT NULL PRIMARY KEY,
		user_id int NOT NULL,
		user_name varchar(255) NOT NULL,
		user_email varchar(255) NOT NULL,
		permission_id int NOT NULL,
		permission_role varchar(255) NOT NULL,
		agent varchar(255) NOT NULL,
		remote_address varchar(255) NOT NULL,
		local_address varchar(255) NOT NULL,
		local_port varchar(255) NOT NULL,
		access_token varchar(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP WITH TIME ZONE
	);