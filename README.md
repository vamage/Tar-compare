##tar-vfs


tool to ensure only changed or missing files are updated. This will limit writes and limit changes to file system bursting cache but peserving the the real path. 



### Parameters 

Source Tar file to pull from

deploy Path to set up symlink paths

Cache   Path to store raw files

###Example

Time to verify changes for linux kernel. 

```vamage@docket:~$ /usr/bin/time -v  ./tar-compare -source master.tar.gz -cache linux-master2
1: /realpath/linux-master
 2:1
        Command being timed: "./tar-compare -source master.tar.gz -cache linux-master22222"
        User time (seconds): 8.38
        System time (seconds): 3.83
        Percent of CPU this job got: 19%
        Elapsed (wall clock) time (h:mm:ss or m:ss): 1:02.34
        Average shared text size (kbytes): 0
        Average unshared data size (kbytes): 0
        Average stack size (kbytes): 0
        Average total size (kbytes): 0
        Maximum resident set size (kbytes): 65656
        Average resident set size (kbytes): 0
        Major (requiring I/O) page faults: 0
        Minor (reclaiming a frame) page faults: 15989
        Voluntary context switches: 5448
        Involuntary context switches: 4741
        Swaps: 0
        File system inputs: 16
        File system outputs: 1825320
        Socket messages sent: 0
        Socket messages received: 0
        Signals delivered: 0
        Page size (bytes): 4096
        Exit status: 0 ``` ```
