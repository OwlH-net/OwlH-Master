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
CREATE TABLE rule_files (
    rule_id integer PRIMARY KEY AUTOINCREMENT,
    rule_uniqueid text NOT NULL,
    rule_param text NOT NULL,
    rule_value text NOT NULL
 );
CREATE TABLE scheduler (
    scheduler_id integer PRIMARY KEY AUTOINCREMENT,
    scheduler_uniqueid text NOT NULL,
    scheduler_param text NOT NULL,
    scheduler_value text NOT NULL
 );
 CREATE TABLE scheduler_log (
    log_id integer PRIMARY KEY AUTOINCREMENT,
    log_uniqueid text NOT NULL,
    log_param text NOT NULL,
    log_value text NOT NULL
 );
 CREATE TABLE plugins (
    plugin_id integer PRIMARY KEY AUTOINCREMENT,
    plugin_uniqueid text NOT NULL,
    plugin_param text NOT NULL,
    plugin_value text NOT NULL
 );
CREATE TABLE groups (
    group_id integer PRIMARY KEY AUTOINCREMENT,
    group_uniqueid text NOT NULL,
    group_param text NOT NULL,
    group_value text NOT NULL
 );
CREATE TABLE masterconfig (
    config_id integer PRIMARY KEY AUTOINCREMENT,
    config_uniqueid text NOT NULL,
    config_param text NOT NULL,
    config_value text NOT NULL
 );
CREATE TABLE dataflow (
    flow_id integer PRIMARY KEY AUTOINCREMENT,
    flow_uniqueid text NOT NULL,
    flow_param text NOT NULL,
    flow_value text NOT NULL
 );
 CREATE TABLE changerecord (
    control_id integer PRIMARY KEY AUTOINCREMENT,
    control_uniqueid text NOT NULL,
    control_param text NOT NULL,
    control_value text NOT NULL
);
CREATE TABLE incidents (
   incidents_id integer PRIMARY KEY AUTOINCREMENT,
   incidents_uniqueid text NOT NULL,
   incidents_param text NOT NULL,
   incidents_value text NOT NULL
);
CREATE TABLE groupnodes (
   gn_id integer PRIMARY KEY AUTOINCREMENT,
   gn_uniqueid text NOT NULL,
   gn_param text NOT NULL,
   gn_value text NOT NULL
);
CREATE TABLE groupcluster (
   gc_id integer PRIMARY KEY AUTOINCREMENT,
   gc_uniqueid text NOT NULL,
   gc_param text NOT NULL,
   gc_value text NOT NULL
);
CREATE TABLE users (
   user_id integer PRIMARY KEY AUTOINCREMENT,
   user_uniqueid text NOT NULL,
   user_param text NOT NULL,
   user_value text NOT NULL
);