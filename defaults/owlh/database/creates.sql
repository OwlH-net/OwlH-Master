// node.db
CREATE TABLE nodes (
 node_id integer PRIMARY KEY AUTOINCREMENT,
 node_uniqueid text NOT NULL,
 node_param text NOT NULL,
 node_value text NOT NULL
);

// ruleset.db
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
CREATE TABLE groups (
    group_id integer PRIMARY KEY AUTOINCREMENT,
    group_uniqueid text NOT NULL,
    group_param text NOT NULL,
    group_value text NOT NULL
 );
CREATE TABLE rule_files (
    rule_id integer PRIMARY KEY AUTOINCREMENT,
    rule_uniqueid text NOT NULL,
    rule_param text NOT NULL,
    rule_value text NOT NULL
 );
