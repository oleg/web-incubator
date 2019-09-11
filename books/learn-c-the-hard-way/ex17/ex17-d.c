#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>

#define MAX_DATA 512
//#define MAX_ROWS 100

struct Address {
  int id;
  int set;
  char name[MAX_DATA];
  char email[MAX_DATA];
};

struct Database {
  size_t max_rows;
  struct Address *rows;
};

struct Connection {
  FILE *file;
  struct Database *db;
};

void die(const char *message)
{
  if (errno) {
    perror(message);
  } else {
    printf("ERROR: %s\n", message);
  }

  exit(1);
}

void Address_print(struct Address *addr)
{
  printf("%d %s %s\n", addr->id, addr->name, addr->email);
}

size_t Database_rowsSize(size_t max_rows)
{
  return sizeof(struct Address) * max_rows;
}

size_t Database_size()
{
  return sizeof(struct Database);
}

struct Connection *Connection_create()
{
  struct Connection *conn = malloc(sizeof(struct Connection));
  if (!conn)
    die("Memory error.");
  return conn;
}  

void Connection_openFile(struct Connection *conn, const char *filename, const char *mode)
{
  conn->file = fopen(filename, mode); //w or r+
  if (!conn->file)
    die("Failed to open the file.");
}

struct Connection *Database_create(const char *filename, int max_rows)
{
  struct Connection *conn = Connection_create();
  Connection_openFile(conn, filename, "w");
  conn->db = malloc(Database_size());
  conn->db->max_rows = max_rows;
  conn->db->rows = malloc(Database_rowsSize(conn->db->max_rows));
  
  int i = 0;
  for (i = 0; i < conn->db->max_rows; i++) {
    struct Address addr = { .id = i, .set = 0 };
    conn->db->rows[i] = addr;
  }
  return conn;
}

struct Connection *Database_open(const char *filename)
{
  struct Connection *conn = Connection_create();
  Connection_openFile(conn, filename, "r+");
  conn->db = malloc(Database_size());
    
  int rc;
  rc = fread(conn->db, Database_size(), 1, conn->file);
  if (rc != 1)
    die("Failed to load database.");

  conn->db->rows = malloc(Database_rowsSize(conn->db->max_rows));
  rc = fread(conn->db->rows, Database_rowsSize(conn->db->max_rows), 1, conn->file);
  if (rc != 1)
    die("Failed to load rows.");
  return conn;

}

void Database_close(struct Connection *conn)
{
  if (conn) {
    if (conn->file)
      fclose(conn->file);
    
    if (conn->db && conn->db->rows)
      free(conn->db->rows);

    if (conn->db)
      free(conn->db);

    free(conn);
  }
}

void Database_write(struct Connection *conn)
{
  rewind(conn->file);

  int rc = fwrite(conn->db, Database_size(), 1, conn->file);
  if (rc != 1)
    die("Failed to write database.");
  
  rc = fwrite(conn->db->rows, Database_rowsSize(conn->db->max_rows), 1, conn->file);
  if (rc != 1)
    die("Failed to write addresses.");

  rc = fflush(conn->file);
  if (rc == -1)
    die("Cannot flush database.");
}

void Database_set(struct Connection *conn, int id, const char *name, const char *email)
{
  struct Address *addr = &conn->db->rows[id];
  if (addr->set)
    die("Already set, delete it first");

  addr->set = 1;
  char *res = strncpy(addr->name, name, MAX_DATA);
  if (!res)
    die("Name copy failed");

  res = strncpy(addr->email, email, MAX_DATA);
  if (!res)
    die("Email copy failed");
}

void Database_get(struct Connection *conn, int id)
{
  struct Address *addr = &conn->db->rows[id];
  if (addr->set) {
    Address_print(addr);
  } else {
    die("ID is not set");
  }
}

void Database_delete(struct Connection *conn, int id)
{
  struct Address addr = { .id = id, .set = 0};
  conn->db->rows[id] = addr;
}

void Database_list(struct Connection *conn)
{
  int i = 0;
  struct Database *db = conn->db;

  for (i = 0; i < db->max_rows; i++) {
    struct Address *curr = &db->rows[i];
    if (curr->set) {
      Address_print(curr);
    }
  }
}

int main(int argc, char *argv[])
{
  if (argc < 3)
    die("USAGE: ex17 <dbfile> <action> [action params]");

  char *filename = argv[1];
  char action = argv[2][0];
  struct Connection *conn;
  int id;

  switch(action) {
  case 'c': 
    conn = Database_create(filename, atoi(argv[3]));
    Database_write(conn);
    break;
  
  case 'g': 
    if (argc != 4)
      die("Need an id to get");
    
    conn = Database_open(filename);
    id = atoi(argv[3]);
    if (id >= conn->db->max_rows)
      die("There's not that many records.");

    Database_get(conn, id);
    break;

  case 's': 
    if (argc != 6)
      die("Need id, name, email to set");
    
    conn = Database_open(filename);
    id = atoi(argv[3]);
    if (id >= conn->db->max_rows)
      die("There's not that many records.");

    Database_set(conn, id, argv[4], argv[5]);
    Database_write(conn);
    break;
  
  case 'd': 
    if (argc != 4)
      die("Need id to delete");

    conn = Database_open(filename);
    id = atoi(argv[3]);
    if (id >= conn->db->max_rows)
      die("There's not that many records.");

    Database_delete(conn, id);
    Database_write(conn);
    break;
  
  case 'l': 
    conn = Database_open(filename);
    Database_list(conn);
    break;

  default:
    die("Invalid action: c=create, g=get, s=set, d=delete, l=list");
  }
  
  Database_close(conn);
  
  return 0;
}
