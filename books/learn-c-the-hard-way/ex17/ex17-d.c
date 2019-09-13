
#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <errno.h>
#include <string.h>


struct Address {
  int id;
  int set;
  char *name;
  char *email;
};

struct Database {
  size_t max_data;
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

size_t Database_dataSize(size_t max_data)
{
  return sizeof(char) * max_data;
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
  conn->file = fopen(filename, mode);
  if (!conn->file)
    die("Failed to open the file.");
}

void Address_create(struct Database *db)
{
  db->rows = malloc(Database_rowsSize(db->max_rows));
  int i = 0;
  for (; i < db->max_rows; i++) {
    
    struct Address addr = {.id=i,.set=0};
    addr.name = malloc(Database_dataSize(db->max_data));
    addr.email = malloc(Database_dataSize(db->max_data));
    
    db->rows[i] = addr;
  }
}

struct Connection *Database_create(const char *filename, int max_rows, int max_data)
{
  struct Connection *conn = Connection_create();
  Connection_openFile(conn, filename, "w");
  conn->db = malloc(Database_size());
  conn->db->max_rows = max_rows;
  conn->db->max_data = max_data;

  Address_create(conn->db);
  return conn;
}

struct Connection *Database_open(const char *filename)
{
  int rc;
  
  struct Connection *conn = Connection_create();
  Connection_openFile(conn, filename, "r+");
  
  conn->db = malloc(Database_size());
  rc = fread(conn->db, Database_size(), 1, conn->file);
  if (rc != 1)
    die("Failed to load database.");

  size_t data_size = Database_dataSize(conn->db->max_data);
  conn->db->rows = malloc(Database_rowsSize(conn->db->max_rows)); 
  int i = 0;
  for (i = 0; i < conn->db->max_rows; i++) {
    struct Address *addr = &conn->db->rows[i];
    rc = fread(addr, sizeof(struct Address), 1, conn->file);
    if (rc != 1)
      die("Failed to load rows.");
    
    addr->name = malloc(data_size);
    rc = fread(addr->name, data_size, 1, conn->file);
    if (rc != 1)
      die("Failed to load rows.");

    addr->email = malloc(data_size);
    rc = fread(addr->email, data_size, 1, conn->file);
    if (rc != 1)
      die("Failed to load rows.");
  }

  if (rc != 1)
    die("Failed to load rows.");
  
  return conn;
}

void Database_write(struct Connection *conn)
{
  rewind(conn->file);

  int rc;
  
  rc = fwrite(conn->db, Database_size(), 1, conn->file);
  if (rc != 1)
    die("Failed to write database.");

  size_t data_size = Database_dataSize(conn->db->max_data);
  int i = 0;
  for (i = 0; i < conn->db->max_rows; i++) {
    
    struct Address addr =  conn->db->rows[i];

    rc = fwrite(&addr, sizeof(struct Address), 1, conn->file);
    if (rc != 1)
      die("Failed to write database.");
    
    rc = fwrite(addr.name, data_size, 1, conn->file);
    if (rc != 1)
      die("Failed to write database.");

    rc = fwrite(addr.email, data_size, 1, conn->file);
    if (rc != 1)
      die("Failed to write database.");
  }

  rc = fflush(conn->file);
  if (rc == -1)
    die("Cannot flush database.");
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

void Database_set(struct Connection *conn, int id, const char *name, const char *email)
{
  struct Address *addr = &conn->db->rows[id];
  if (addr->set)
    die("Already set, delete it first");

  addr->set = 1;
  char *res;
  res = strncpy(addr->name, name, conn->db->max_data);

  if (!res)
    die("Name copy failed");

  res = strncpy(addr->email, email, conn->db->max_data);
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
  struct Address addr =
    {
     .id = id,
     .set = 0,
     .name = malloc(Database_dataSize(conn->db->max_data)),
     .email = malloc(Database_dataSize(conn->db->max_data))
    };
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
    conn = Database_create(filename, atoi(argv[3]), atoi(argv[4]));
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
