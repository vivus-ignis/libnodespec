#include <assert.h>
#include <errno.h>
#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>
#include <sys/sysctl.h>
#include <string.h>

typedef struct kinfo_proc kinfo_proc;

/*
 *
 * https://developer.apple.com/legacy/library/qa/qa2001/qa1123.html
 *
 */

static int GetBSDProcessList(kinfo_proc **procList, size_t *procCount)
    // Returns a list of all BSD processes on the system.  This routine
    // allocates the list and puts it in *procList and a count of the
    // number of entries in *procCount.  You are responsible for freeing
    // this list (use "free" from System framework).
    // On success, the function returns 0.
    // On error, the function returns a BSD errno value.
{
    int                 err;
    kinfo_proc *        result;
    bool                done;
    int name[4];
    name[0] = CTL_KERN;
    name[1] = KERN_PROC;
    name[2] = KERN_PROC_ALL;
    name[3] = 0;

    // Declaring name as const requires us to cast it when passing it to
    // sysctl because the prototype doesn't include the const modifier.
    size_t              length;

    assert( procList != NULL);
    assert(*procList == NULL);
    assert(procCount != NULL);

    *procCount = 0;

    // We start by calling sysctl with result == NULL and length == 0.
    // That will succeed, and set length to the appropriate length.
    // We then allocate a buffer of that size and call sysctl again
    // with that buffer.  If that succeeds, we're done.  If that fails
    // with ENOMEM, we have to throw away our buffer and loop.  Note
    // that the loop causes use to call sysctl with NULL again; this
    // is necessary because the ENOMEM failure case sets length to
    // the amount of data returned, not the amount of data that
    // could have been returned.

    result = NULL;
    done = false;
    do {
        assert(result == NULL);

        // Call sysctl with a NULL buffer.

        /* printf("%d\n", *name); */

        length = 0;
        /* err = sysctl( (int *) name, (sizeof(name) / sizeof(*name)) - 1, */
        err = sysctl( name, (sizeof(name) / sizeof(*name)) - 1,
                      NULL, &length,
                      NULL, 0);
        if (err == -1) {
            err = errno;
        }

        // Allocate an appropriately sized buffer based on the results
        // from the previous call.

        if (err == 0) {
            result = malloc(length);
            if (result == NULL) {
                err = ENOMEM;
            }
        }

        // Call sysctl again with the new buffer.  If we get an ENOMEM
        // error, toss away our buffer and start again.

        if (err == 0) {
            err = sysctl( (int *) name, (sizeof(name) / sizeof(*name)) - 1,
                          result, &length,
                          NULL, 0);
            if (err == -1) {
                err = errno;
            }
            if (err == 0) {
                done = true;
            } else if (err == ENOMEM) {
                assert(result != NULL);
                free(result);
                result = NULL;
                err = 0;
            }
        }
    } while (err == 0 && ! done);

    // Clean up and establish post conditions.

    if (err != 0 && result != NULL) {
        free(result);
        result = NULL;
    }
    *procList = result;
    if (err == 0) {
        *procCount = length / sizeof(kinfo_proc);
    }

    assert( (err == 0) == (*procList != NULL) );

    return err;
}

int isBSDProcessExists(char* psname)
{
  size_t nil_i = 0;
  size_t *proc_count = &nil_i;

  kinfo_proc *nil = NULL;
  kinfo_proc **proc_list = &nil;

  int res;
  res = GetBSDProcessList(proc_list, proc_count);

  int i;
  for(i = 0; i <= *proc_count; i++) {
    // printf("%s <=> %s\n", psname, (*proc_list)[i].kp_proc.p_comm);
    if (strcmp(psname, (*proc_list)[i].kp_proc.p_comm) == 0) {
      return 0;
    }
  }

  return 1;
}
