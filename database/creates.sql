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

CREATE TABLE ruleset (
 ruleset_id integer PRIMARY KEY AUTOINCREMENT,
 ruleset_uniqueid text NOT NULL,
 ruleset_param text NOT NULL,
 ruleset_value text NOT NULL
);

CREATE TABLE ruleset_node (
    ruleset_id integer PRIMARY KEY AUTOINCREMENT,
    ruleset_uniqueid text NOT NULL,
    node_uniqueid text NOT NULL
 );

 CREATE TABLE rule_note (
    ruleset_id integer PRIMARY KEY AUTOINCREMENT,
    ruleset_uniqueid text NOT NULL,
    rule_sid text NOT NULL,
    note_date text NOT NULL,
    ruleNote text NOT NULL
 );