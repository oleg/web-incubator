-- for main
CREATE ROLE mainworder WITH CREATEDB LOGIN PASSWORD 'mainpass';
CREATE DATABASE worder OWNER mainworder;

-- for test
CREATE ROLE testworder WITH CREATEDB LOGIN PASSWORD 'testpass';
CREATE DATABASE worder_test OWNER testworder;
