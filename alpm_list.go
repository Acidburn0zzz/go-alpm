package alpm

/*
#include <stdlib.h>
#include <alpm_list.h>
*/
import "C"

import (
  "unsafe"
)

type AlpmList struct {
  alpm_list_t *C.alpm_list_t
}

/* allocation */
func (v *AlpmList) Free() {
  C.alpm_list_free(v.alpm_list_t)
}

/* mutators */
func (v *AlpmList) Add(data unsafe.Pointer) *AlpmList {
  return &AlpmList{C.alpm_list_add(v.alpm_list_t, data)}
}

/* accessors */
func (v *AlpmList) First() *AlpmList {
  return &AlpmList{C.alpm_list_first(v.alpm_list_t)}
}

func (v *AlpmList) Nth(n C.int) *AlpmList {
  return &AlpmList{C.alpm_list_nth(v.alpm_list_t, n)}
}

func (v *AlpmList) Next() *AlpmList {
  return &AlpmList{C.alpm_list_next(v.alpm_list_t)}
}

func (v *AlpmList) Last() *AlpmList {
  return &AlpmList{C.alpm_list_last(v.alpm_list_t)}
}

func (v *AlpmList) GetData() interface{} {
  return v.alpm_list_t.data
}

/* misc */

