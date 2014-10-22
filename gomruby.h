#ifndef _GOMRUBY_H_INCLUDED
#define _GOMRUBY_H_INCLUDED

#include <stdlib.h>
#include <stdio.h>
#include <errno.h>

#include <mruby.h>
#include <mruby/compile.h>
#include <mruby/irep.h>
#include <mruby/dump.h>
#include <mruby/proc.h>
#include <mruby/string.h>

// Sets ctx->no_exec to avoid unwanted executions
static void
__cxt_no_exec(mrbc_context *ctx)
{
  ctx->no_exec = 1;
}

// Sets ctx->filename is necessary to dump the bytecodes
static void
__cxt_flename(mrb_state *mrb, mrbc_context *ctx, const char *filename)
{
  if (filename) {
    int len = strlen(filename);
    char *p = (char *)mrb_alloca(mrb, len + 1);

    memcpy(p, filename, len + 1);
    ctx->filename = p;
  }
}

// Getting the irep, go can't handle `->`
static mrb_irep *
__ptr_irep(mrb_value b)
{
  return mrb_proc_ptr(b)->body.irep;
}

// From error.c's exc_to_s,
// This function returns a mrb_value which holds the mruby's exception's string,
// if there was no exception, it returns an mruby nil value.
char*
__mrb_exc_cstr(mrb_state *mrb)
{
  if (mrb->exc == NULL) {
      errno = 1;
      return NULL;
  }
  mrb_value exc = mrb_obj_value(mrb->exc);
  mrb_value mesg = mrb_attr_get(mrb, exc, mrb_intern_lit(mrb, "mesg"));

  if (mrb_nil_p(mesg)) return mrb_str_to_cstr(mrb, mrb_str_new_cstr(mrb, mrb_obj_classname(mrb, exc)));
  return mrb_str_to_cstr(mrb, mesg);
}

#endif
