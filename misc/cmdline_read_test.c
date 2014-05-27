#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/stat.h>

#define FILEPATH "/proc/778/cmdline"

int main(int argc, char *argv[])
{
  int fd = 0;
  int rt = 0;
  struct stat st_buf;
  
  rt = stat(FILEPATH, &st_buf);
  if (rt != 0) {
    printf("Stat error code: %d\n", errno);
    char *error_str = strerror(errno);
    printf("Stat error: %s\n", error_str);
    return rt;
  }

  printf("File size: %d\n", st_buf.off_t);

  fd = open(FILEPATH, O_RDONLY);
  close(fd);
}
