##tar-compare


tool to ensure only changed or missing files are updated

Supports a purge url to clean out caches as need for things like varnish/phpopache/etc


### Parameters 

Source Tar file to use

Destination  Path to compare

Purgeurl url to send purge

###Example

```bash
vamage@docket:~$ /usr/bin/time -v  ./tar-compare -source master.tar.gz -destination linux-master
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fnamespace.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2FMakefile
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Futil.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fsyscall.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fshm.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fmsg.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fipc_sysctl.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fmqueue.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fcompat.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Futil.h
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fmsgutil.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fmq_sysctl.c
http://127.0.0.1/opcache.php?file=%2Fhome%2Fvamage%2Flinux-master%2Fipc2%2Fipc%2Fsem.c
master.tar.gz , linux-master, 61676The hashes false
Command exited with non-zero status 12
        Command being timed: "./tar-compare -source master.tar.gz -destination linux-master"
        User time (seconds): 10.68
        System time (seconds): 0.99
        Percent of CPU this job got: 64%
        Elapsed (wall clock) time (h:mm:ss or m:ss): 0:18.09
        Average shared text size (kbytes): 0
        Average unshared data size (kbytes): 0
        Average stack size (kbytes): 0
        Average total size (kbytes): 0
        Maximum resident set size (kbytes): 98732
        Average resident set size (kbytes): 0
        Major (requiring I/O) page faults: 0
        Minor (reclaiming a frame) page faults: 24332
        Voluntary context switches: 31132
        Involuntary context switches: 3217
        Swaps: 0
        File system inputs: 315592
        File system outputs: 0
        Socket messages sent: 0
        Socket messages received: 0
        Signals delivered: 0
        Page size (bytes): 4096
        Exit status: 12
```


