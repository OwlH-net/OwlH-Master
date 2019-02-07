
// Limpiar o eliminar
CREATE TABLE master (
 master_id integer PRIMARY KEY AUTOINCREMENT,
 master_name text NOT NULL,
 master_ip text NOT NULL,
 master_port integer NOT NULL
);

// limpiar o eliminar mejor

CREATE TABLE node (
 node_id integer PRIMARY KEY AUTOINCREMENT,
 node_name text NOT NULL,
 node_ip text NOT NULL,
 node_port integer NOT NULL,
 node_type text NOT NULL,
 noce_UUID text NOT NULL
);

CREATE TABLE aboutme (
 am_id integer PRIMARY KEY AUTOINCREMENT,
 am_param text NOT NULL, 
 am_value text NOT NULL
);

CREATE TABLE nodes (
 node_id integer PRIMARY KEY AUTOINCREMENT,
 node_uniqueid text NOT NULL,
 node_param text NOT NULL,
 node_value text NOT NULL
);