#ifndef lcthw_List_merge_sort_node_h
#define lcthw_List_merge_sort_node_h

#include <lcthw/list.h>


typedef int (*List_compare) (const void *a, const void *b);

List *List_merge_sort_node(List * list, List_compare cmp);

#endif
